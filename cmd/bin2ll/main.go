// The bin2ll tool converts binary executables to equivalent LLVM IR assembly
// (*.exe -> *.ll).
package main

import (
	"bufio"
	"debug/pe"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir/metadata"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// dbg represents a logger with the "bin2ll:" prefix, which logs debug messages
// to standard error.
var dbg = log.New(os.Stderr, term.MagentaBold("bin2ll:")+" ", 0)

func usage() {
	const use = `
Convert binary executables to equivalent LLVM IR assembly (*.exe -> *.ll).

Usage:

	bin2ll [OPTION]... FILE.ll

Flags:
`
	fmt.Fprint(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line arguments.
	var (
		// blockAddr specifies a basic block address to decompile.
		blockAddr bin.Address
		// TODO: Remove -first flag and firstAddr.
		// firstAddr specifies the first function address to decompile.
		firstAddr bin.Address
		// funcAddr specifies a function address to decompile.
		funcAddr bin.Address
		// quiet specifies whether to suppress non-error messages.
		quiet bool
	)
	flag.Usage = usage
	flag.Var(&blockAddr, "block", "basic block address to decompile")
	flag.Var(&firstAddr, "first", "first function address to decompile")
	flag.Var(&funcAddr, "func", "function address to decompile")
	flag.BoolVar(&quiet, "q", false, "suppress non-error messages")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	binPath := flag.Arg(0)
	// Mute debug messages if `-q` is set.
	if quiet {
		dbg.SetOutput(ioutil.Discard)
	}

	// Convert binary into LLVM IR assembly.
	d, err := parseFile(binPath)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer d.file.Close()

	// TODO: Remove -block. Used for debugging.
	if blockAddr != 0 {
		block, err := d.decodeBlock(blockAddr)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		printBlock(block)
		//fmt.Println("targets from basic block address:", block.addr)
		//targets := d.targets(block.term)
		//for _, target := range targets {
		//	fmt.Println(target)
		//}
		return
	}

	// Translate functions from x86 machine code to LLVM IR assembly.
	funcAddrs := d.funcAddrs
	if funcAddr != 0 {
		funcAddrs = []bin.Address{funcAddr}
	}
	for _, funcAddr := range funcAddrs {
		if firstAddr != 0 && funcAddr < firstAddr {
			// skip functions with lower address than the first function.
			continue
		}
		f, err := d.decodeFunc(funcAddr)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		if err := d.translateFunc(f); err != nil {
			log.Fatalf("%+v", err)
		}
		printFunc(f)
		fmt.Println(f.Function)
	}
}

// A disassembler tracks information required to disassemble x86 executables.
type disassembler struct {
	// PE file.
	file *pe.File
	// Processor mode (16, 32 or 64).
	mode int
	// Image base address.
	imageBase uint64
	// .text section base address.
	codeBase uint64
	// .text section size.
	codeSize uint64
	// .idata section base address.
	idataBase uint64
	// .idata section size.
	idataSize uint64
	// Entry address.
	entry bin.Address
	// Function addresses.
	funcAddrs []bin.Address
	// Basic block addresses.
	blockAddrs []bin.Address
	// Map from jump table address to target addresses.
	tables map[bin.Address][]bin.Address
	// Chunks of bytes.
	chunks []Chunk
	// Functions.
	funcs map[bin.Address]*function
	// Map from basic block address (function chunk) to function address, to
	// which the basic block belongs.
	chunkFunc map[bin.Address]bin.Address
	// TODO: Remove.
	decodedBlock map[bin.Address]bool
}

// parseFile parses the given PE file and associated JSON files, containing
// information required to disassemble the x86 executables.
func parseFile(binPath string) (*disassembler, error) {
	// Parse PE executable.
	file, err := pe.Open(binPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	d := &disassembler{
		file:         file,
		tables:       make(map[bin.Address][]bin.Address),
		chunkFunc:    make(map[bin.Address]bin.Address),
		decodedBlock: make(map[bin.Address]bool),
	}
	switch opt := file.OptionalHeader.(type) {
	case *pe.OptionalHeader32:
		d.mode = 32
		d.imageBase = uint64(opt.ImageBase)
		d.codeBase = uint64(opt.BaseOfCode)
		d.codeSize = uint64(opt.SizeOfCode)
		d.idataBase = uint64(opt.DataDirectory[12].VirtualAddress)
		d.idataSize = uint64(opt.DataDirectory[12].Size)
		d.entry = bin.Address(opt.AddressOfEntryPoint)
	case *pe.OptionalHeader64:
		d.mode = 64
		d.imageBase = opt.ImageBase
		d.codeBase = uint64(opt.BaseOfCode)
		d.codeSize = uint64(opt.SizeOfCode)
		d.idataBase = uint64(opt.DataDirectory[12].VirtualAddress)
		d.idataSize = uint64(opt.DataDirectory[12].Size)
		d.entry = bin.Address(opt.AddressOfEntryPoint)
	default:
		panic(fmt.Errorf("support for optional header type %T not yet implemented", opt))
	}
	dbg.Println("executable entry address:", d.entry)

	// Parse function addresses.
	funcAddrs, err := parseAddrs("funcs.json")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sort.Sort(bin.Addresses(funcAddrs))
	d.funcAddrs = funcAddrs

	// Parse basic block addresses.
	blockAddrs, err := parseAddrs("blocks.json")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sort.Sort(bin.Addresses(blockAddrs))
	d.blockAddrs = blockAddrs

	// Parse data addresses (e.g. jump tables).
	dataAddrs, err := parseAddrs("data.json")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sort.Sort(bin.Addresses(dataAddrs))

	// Parse jump table targets.
	if err := parseJSON("tables.json", &d.tables); err != nil {
		return nil, errors.WithStack(err)
	}

	// Parse function chunks.
	if err := parseJSON("chunks.json", &d.chunkFunc); err != nil {
		return nil, errors.WithStack(err)
	}

	// Append basic blocks as code chunks.
	for _, blockAddr := range blockAddrs {
		chunk := Chunk{
			kind: kindCode,
			addr: blockAddr,
		}
		d.chunks = append(d.chunks, chunk)
	}

	// Append data as data chunks.
	for _, dataAddr := range dataAddrs {
		chunk := Chunk{
			kind: kindData,
			addr: dataAddr,
		}
		d.chunks = append(d.chunks, chunk)
	}
	less := func(i, j int) bool {
		return d.chunks[i].addr < d.chunks[j].addr
	}
	sort.Slice(d.chunks, less)

	// Functions.
	d.funcs = make(map[bin.Address]*function)
	module, err := asm.ParseFile("funcs.ll")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, f := range module.Funcs {
		node, ok := f.Metadata["addr"]
		if !ok {
			return nil, errors.Errorf(`unable to locate "addr" metadata for function %q`, f.Name)
		}
		var entry bin.Address
		if err := metadata.Unmarshal(node, &entry); err != nil {
			return nil, errors.WithStack(err)
		}
		fn := &function{
			Function: f,
			entry:    entry,
			blocks:   make(map[bin.Address]*basicBlock),
		}
		d.funcs[entry] = fn
	}

	return d, nil
}

// vaddr returns the virtual address for the specified offset from the image
// base.
func (d *disassembler) vaddr(offset uint64) bin.Address {
	return bin.Address(d.imageBase + offset)
}

// data returns access to the data of the executable starting at the given
// address.
func (d *disassembler) data(addr bin.Address) ([]byte, error) {
	for _, section := range d.file.Sections {
		start := d.vaddr(uint64(section.VirtualAddress))
		end := start + bin.Address(section.Size)
		if start <= addr && addr < end {
			offset := uint64(addr - start)
			data, err := section.Data()
			if err != nil {
				return nil, errors.Errorf("unable to access data of section %q; %v", section.Name, err)
			}
			return data[offset:], nil
		}
	}
	return nil, errors.Errorf("unable to locate section for address %v", addr)
}

// Chunk represents a chunk of bytes.
type Chunk struct {
	// Chunk kind.
	kind kind
	// Chunk address.
	addr bin.Address
}

// kind represents the set of chunk kinds.
type kind uint

// Chunk kinds.
const (
	kindNone kind = iota
	kindCode
	kindData
)

// ### [ Helper functions ] ####################################################

// parseAddrs parses the given JSON file and returns the list of addresses
// contained within.
func parseAddrs(jsonPath string) ([]bin.Address, error) {
	var addrs []bin.Address
	if err := parseJSON(jsonPath, &addrs); err != nil {
		return nil, errors.WithStack(err)
	}
	return addrs, nil
}

// parseJSON parses the given JSON file and stores the result into v.
func parseJSON(jsonPath string, v interface{}) error {
	f, err := os.Open(jsonPath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	dec := json.NewDecoder(br)
	return dec.Decode(v)
}
