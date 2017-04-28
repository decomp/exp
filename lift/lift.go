package lift

import (
	"log"
	"os"

	"github.com/decomp/exp/bin"
	"github.com/decomp/exp/disasm/x86"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/metadata"
	"github.com/mewkiz/pkg/osutil"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// TODO: Remove loggers once the library matures.

// Loggers.
var (
	// dbg represents a logger with the "lift:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.CyanBold("lift:")+" ", 0)
	// warn represents a logger with the "warning:" prefix, which logs warning
	// messages to standard error.
	warn = log.New(os.Stderr, term.RedBold("warning:")+" ", 0)
)

// A Lifter tracks information required to lift the assembly of a binary
// executable.
//
// Data should only be written to this structure during initialization. After
// initialization the structure is considered in read-only mode to allow for
// concurrent lifting of functions.
type Lifter struct {
	*x86.Disasm
	// Functions.
	Funcs map[bin.Address]*Func
	// Global variables.
	Globals map[bin.Address]*ir.Global
}

// NewLifter creates a new Lifter for accessing the assembly instructions of the
// given binary executable, and the information contained within associated JSON
// and LLVM IR files.
func NewLifter(file *bin.File) (*Lifter, error) {
	// Prepare lifter.
	dis, err := x86.NewDisasm(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	l := &Lifter{
		Disasm:  dis,
		Funcs:   make(map[bin.Address]*Func),
		Globals: make(map[bin.Address]*ir.Global),
	}

	// Parse associated LLVM IR information.
	llPath := "info.ll"
	if !osutil.Exists(llPath) {
		warn.Printf("unable to locate LLVM IR file %q", llPath)
		return l, nil
	}
	module, err := asm.ParseFile(llPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Parse globals.
	for _, g := range module.Globals {
		node, ok := g.Metadata["addr"]
		if !ok {
			return nil, errors.Errorf(`unable to locate "addr" metadata for global variable %q`, g.Name)
		}
		var addr bin.Address
		if err := metadata.Unmarshal(node, &addr); err != nil {
			return nil, errors.WithStack(err)
		}
		l.Globals[addr] = g
	}

	// Parse function signatures.
	for _, f := range module.Funcs {
		node, ok := f.Metadata["addr"]
		if !ok {
			return nil, errors.Errorf(`unable to locate "addr" metadata for function %q`, f.Name)
		}
		var entry bin.Address
		if err := metadata.Unmarshal(node, &entry); err != nil {
			return nil, errors.WithStack(err)
		}
		fn := &Func{
			Function: f,
		}
		l.Funcs[entry] = fn
	}

	return l, nil
}

type Func struct {
	*ir.Function
}
