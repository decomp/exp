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
		// funcAddr specifies a function address to decompile.
		funcAddr bin.Address
		// quiet specifies whether to suppress non-error messages.
		quiet bool
	)
	flag.Usage = usage
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

	// Translate functions from x86 machine code to LLVM IR assembly.
	funcAddrs := d.funcAddrs
	if funcAddr != 0 {
		funcAddrs = []bin.Address{funcAddr}
	}
	for _, funcAddr := range funcAddrs {
		f, err := d.decodeFunc(funcAddr)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		printFunc(f)
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
	// Entry address.
	entry bin.Address
	// Function addresses.
	funcAddrs []bin.Address
	// Basic block addresses.
	blockAddrs []bin.Address
	// Chunks of bytes.
	chunks []Chunk
	// Functions.
	funcs map[bin.Address]*function
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
		file: file,
	}
	switch opt := file.OptionalHeader.(type) {
	case *pe.OptionalHeader32:
		d.mode = 32
		d.imageBase = uint64(opt.ImageBase)
		d.entry = bin.Address(opt.AddressOfEntryPoint)
	case *pe.OptionalHeader64:
		d.mode = 64
		d.imageBase = opt.ImageBase
		d.entry = bin.Address(opt.AddressOfEntryPoint)
	default:
		panic(fmt.Errorf("support for optional header type %T not yet implemented", opt))
	}
	fmt.Println("executable entry address:", d.entry)

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
