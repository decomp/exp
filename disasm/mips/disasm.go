// Package mips implements a disassembler for the MIPS architecture.
package mips

import (
	"fmt"
	"log"
	"os"

	"github.com/decomp/exp/bin"
	"github.com/decomp/exp/disasm"
	"github.com/mewkiz/pkg/jsonutil"
	"github.com/mewkiz/pkg/osutil"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// TODO: Remove loggers once the library matures.

// Loggers.
var (
	// dbg represents a logger with the "mips:" prefix, which logs debug messages
	// to standard error.
	dbg = log.New(os.Stderr, term.BlueBold("mips:")+" ", 0)
	// warn represents a logger with the "warning:" prefix, which logs warning
	// messages to standard error.
	warn = log.New(os.Stderr, term.RedBold("warning:")+" ", 0)
)

// A Disasm tracks information required to disassemble a binary executable.
//
// Data should only be written to this structure during initialization. After
// initialization the structure is considered in read-only mode to allow for
// concurrent decoding of functions.
type Disasm struct {
	*disasm.Disasm
	// Processor mode.
	Mode int
}

// NewDisasm creates a new Disasm for accessing the assembly instructions of the
// given binary executable.
//
// Associated files of the generic disassembler.
//
//    funcs.json
//    blocks.json
//    tables.json
//    chunks.json
//    data.json
//
// Associated files of the MIPS disassembler.
//
//    contexts.json
func NewDisasm(file *bin.File) (*Disasm, error) {
	// Prepare MIPS disassembler.
	d, err := disasm.New(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dis := &Disasm{
		Disasm: d,
	}

	// Parse processor mode.
	switch dis.File.Arch {
	case bin.ArchMIPS_32:
		dis.Mode = 32
	default:
		panic(fmt.Errorf("support for machine architecture %v not yet implemented", dis.File.Arch))
	}

	// Parse CPU contexts.
	//if err := parseJSON("contexts.json", &dis.Contexts); err != nil {
	//	return nil, errors.WithStack(err)
	//}

	return dis, nil
}

// ### [ Helper functions ] ####################################################

// parseJSON parses the given JSON file and stores the result into v.
func parseJSON(jsonPath string, v interface{}) error {
	if !osutil.Exists(jsonPath) {
		warn.Printf("unable to locate JSON file %q", jsonPath)
		return nil
	}
	return jsonutil.ParseFile(jsonPath, v)
}
