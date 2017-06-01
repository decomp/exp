// The bin2ll tool lifts binary executables to equivalent LLVM IR assembly
// (*.exe -> *.ll).
package main

import (
	"debug/pe"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/decomp/exp/bin"
	_ "github.com/decomp/exp/bin/elf" // register ELF decoder
	_ "github.com/decomp/exp/bin/pe"  // register PE decoder
	_ "github.com/decomp/exp/bin/pef" // register PEF decoder
	"github.com/decomp/exp/lift"
	"github.com/llir/llvm/ir"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// Loggers.
var (
	// dbg represents a logger with the "bin2ll:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.MagentaBold("bin2ll:")+" ", 0)
	// warn represents a logger with the "warning:" prefix, which logs warning
	// messages to standard error.
	warn = log.New(os.Stderr, term.RedBold("warning:")+" ", 0)
)

func usage() {
	const use = `
Lift binary executables to equivalent LLVM IR assembly (*.exe -> *.ll).

Usage:

	bin2ll [OPTION]... FILE

Flags:
`
	fmt.Fprint(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line arguments.
	var (
		// blockAddr specifies a basic block address to lift.
		blockAddr bin.Address
		// TODO: Remove -first flag and firstAddr.
		// firstAddr specifies the first function address to lift.
		firstAddr bin.Address
		// funcAddr specifies a function address to lift.
		funcAddr bin.Address
		// TODO: Remove -last flag and lastAddr.
		// lastAddr specifies the last function address to disassemble.
		lastAddr bin.Address
		// quiet specifies whether to suppress non-error messages.
		quiet bool
	)
	flag.Usage = usage
	flag.Var(&blockAddr, "block", "basic block address to lift")
	flag.Var(&firstAddr, "first", "first function address to lift")
	flag.Var(&funcAddr, "func", "function address to lift")
	flag.Var(&lastAddr, "last", "last function address to lift")
	flag.BoolVar(&quiet, "q", false, "suppress non-error messages")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	binPath := flag.Arg(0)
	// Mute debug and warning messages if `-q` is set.
	if quiet {
		dbg.SetOutput(ioutil.Discard)
		warn.SetOutput(ioutil.Discard)
	}

	// Prepare x86 to LLVM IR lifter for the binary executable.
	l, err := newLifter(binPath)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Lift basic block.
	if blockAddr != 0 {
		block, err := l.DecodeBlock(blockAddr)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		_ = block
		return
	}

	// Lift function specified by `-func` flag.
	funcAddrs := l.FuncAddrs
	if funcAddr != 0 {
		funcAddrs = []bin.Address{funcAddr}
	}

	// Create function lifters.
	for _, funcAddr := range funcAddrs {
		if firstAddr != 0 && funcAddr < firstAddr {
			// skip functions before first address.
			continue
		}
		if lastAddr != 0 && funcAddr >= lastAddr {
			// skip functions after last address.
			break
		}
		asmFunc, err := l.DecodeFunc(funcAddr)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		f := l.NewFunc(asmFunc)
		l.Funcs[funcAddr] = f
	}

	for _, funcAddr := range funcAddrs {
		f, ok := l.Funcs[funcAddr]
		if !ok {
			continue
		}
		f.Lift()
		fmt.Println("f:", f)
	}
}

// newLifter returns a new x86 to LLVM IR lifter for the given binary
// executable.
func newLifter(binPath string) (*lift.Lifter, error) {
	// Parse binary executable.
	file, err := bin.ParseFile(binPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return lift.NewLifter(file)
}

// A disassembler tracks information required to disassemble x86 executables.
//
// Information to this structure should be written exactly once, during
// initialization. After initialization the structure is considered in read-only
// mode to allow for concurrent decoding of functions.
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
	funcs map[bin.Address]*Func
	// Global variables.
	globals map[bin.Address]*ir.Global
	// Map from basic block address (function chunk) to function address, to
	// which the basic block belongs.
	chunkFunc map[bin.Address]bin.Address
	// CPU contexts.
	contexts Contexts
	// TODO: Remove.
	decodedBlock map[bin.Address]bool
}

func (d *disassembler) data(addr bin.Address) ([]byte, error) {
	panic("not yet implemented")
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
