// The bin2ll tool lifts binary executables to equivalent LLVM IR assembly
// (*.exe -> *.ll).
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/decomp/exp/bin"
	_ "github.com/decomp/exp/bin/elf" // register ELF decoder
	_ "github.com/decomp/exp/bin/pe"  // register PE decoder
	_ "github.com/decomp/exp/bin/pef" // register PEF decoder
	"github.com/decomp/exp/bin/raw"
	"github.com/decomp/exp/lift/x86"
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
		// output specifies the output path.
		output string
		// quiet specifies whether to suppress non-error messages.
		quiet bool
		// rawArch specifies the machine architecture of a raw binary executable.
		rawArch bin.Arch
		// rawEntry specifies the entry point of a raw binary executable.
		rawEntry bin.Address
		// rawBase specifies the base address of a raw binary executable.
		rawBase bin.Address
	)
	flag.Usage = usage
	flag.Var(&blockAddr, "block", "basic block address to lift")
	flag.Var(&firstAddr, "first", "first function address to lift")
	flag.Var(&funcAddr, "func", "function address to lift")
	flag.Var(&lastAddr, "last", "last function address to lift")
	flag.StringVar(&output, "o", "", "output path")
	flag.BoolVar(&quiet, "q", false, "suppress non-error messages")
	flag.Var(&rawArch, "raw", "machine architecture of raw binary executable (x86_32, x86_64, PowerPC_32, ...)")
	flag.Var(&rawEntry, "rawentry", "entry point of raw binary executable")
	flag.Var(&rawBase, "rawbase", "base address of raw binary executable")
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
	l, err := newLifter(binPath, rawArch, rawEntry, rawBase)
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
	var funcAddrs bin.Addresses
	if funcAddr != 0 {
		funcAddrs = []bin.Address{funcAddr}
	} else {
		for _, funcAddr := range l.FuncAddrs {
			if firstAddr != 0 && funcAddr < firstAddr {
				// skip functions before first address.
				continue
			}
			if lastAddr != 0 && funcAddr >= lastAddr {
				// skip functions after last address.
				break
			}
			funcAddrs = append(funcAddrs, funcAddr)
		}
	}

	// Create function lifters.
	for _, funcAddr := range funcAddrs {
		asmFunc, err := l.DecodeFunc(funcAddr)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		f := l.NewFunc(asmFunc)
		l.Funcs[funcAddr] = f
	}

	// Lift functions.
	for i, funcAddr := range funcAddrs {
		if i != 0 {
			dbg.Println()
		}
		f, ok := l.Funcs[funcAddr]
		if !ok {
			continue
		}
		f.Lift()
		dbg.Println(f)
	}

	// Store LLVM IR output.
	w := os.Stdout
	if len(output) > 0 {
		f, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	}
	var funcs []*ir.Function
	var allFuncAddrs bin.Addresses
	for funcAddr := range l.Funcs {
		allFuncAddrs = append(allFuncAddrs, funcAddr)
	}
	sort.Sort(allFuncAddrs)
	for _, funcAddr := range allFuncAddrs {
		f := l.Funcs[funcAddr]
		funcs = append(funcs, f.Function)
	}
	var globals []*ir.Global
	var globalAddrs bin.Addresses
	for globalAddr := range l.Globals {
		globalAddrs = append(globalAddrs, globalAddr)
	}
	sort.Sort(globalAddrs)
	for _, globalAddr := range globalAddrs {
		g := l.Globals[globalAddr]
		globals = append(globals, g)
	}
	m := &ir.Module{
		TypeDefs: l.TypeDefs,
		Globals:  globals,
		Funcs:    funcs,
	}
	if _, err := fmt.Fprintln(w, m); err != nil {
		log.Fatalf("%+v", err)
	}

	// Create call graph.
	//if err := genCallGraph(l.Funcs); err != nil {
	//	log.Fatalf("%+v", err)
	//}
}

// newLifter returns a new x86 to LLVM IR lifter for the given binary
// executable.
func newLifter(binPath string, rawArch bin.Arch, rawEntry, rawBase bin.Address) (*x86.Lifter, error) {
	// Parse raw binary executable.
	if rawArch != 0 {
		file, err := raw.ParseFile(binPath, rawArch)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		file.Entry = rawEntry
		file.Sections[0].Addr = rawBase
		return x86.NewLifter(file)
	}
	// Parse binary executable.
	file, err := bin.ParseFile(binPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return x86.NewLifter(file)
}
