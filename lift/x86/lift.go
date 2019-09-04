// Package x86 implements x86 to LLVM IR lifting.
package x86

import (
	"fmt"
	"log"
	"os"

	"github.com/decomp/exp/bin"
	"github.com/decomp/exp/disasm/x86"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
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
	// Type definitions.
	TypeDefs []types.Type
	// Functions.
	Funcs map[bin.Address]*Func
	// Map from function name to function. May also contain external functions
	// without associated virtual addresses (e.g. loaded using GetProcAddress).
	FuncByName map[string]*ir.Func
	// Global variables.
	Globals map[bin.Address]*ir.Global
}

// NewLifter creates a new Lifter for accessing the assembly instructions of the
// given binary executable, and the information contained within associated JSON
// and LLVM IR files.
//
// Associated files of the generic disassembler.
//
//    funcs.json
//    blocks.json
//    tables.json
//    chunks.json
//    data.json
//
// Associated files of the x86 disassembler.
//
//    contexts.json
//
// Associated files of the x86 to LLVM IR lifter.
//
//    info.ll
func NewLifter(file *bin.File) (*Lifter, error) {
	// Prepare x86 to LLVM IR lifter.
	dis, err := x86.NewDisasm(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	l := &Lifter{
		Disasm:     dis,
		Funcs:      make(map[bin.Address]*Func),
		FuncByName: make(map[string]*ir.Func),
		Globals:    make(map[bin.Address]*ir.Global),
	}

	// Parse associated LLVM IR information.
	llPath := "info.ll"
	module, err := parseModule(llPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Parse types.
	l.TypeDefs = module.TypeDefs

	// Parse globals.
	for _, g := range module.Globals {
		node, ok := findMetadataAttachment(g.Metadata, "addr")
		if !ok {
			return nil, errors.Errorf(`unable to locate "addr" metadata for global variable %q`, g.Ident())
		}
		addr, err := parseMetadataAddr(node)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		l.Globals[addr] = g
	}

	// Parse function signatures.
	for _, f := range module.Funcs {
		l.FuncByName[f.Name()] = f
		node, ok := findMetadataAttachment(f.Metadata, "addr")
		if !ok {
			warn.Printf(`unable to locate "addr" metadata for function %q; potentially external function without associated virtual addresses (e.g. loaded with GetProcAddress)`, f.Ident())
			continue
		}
		entry, err := parseMetadataAddr(node)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fn := &Func{
			Func: f,
		}
		l.Funcs[entry] = fn
	}

	// Parse imports.
	addFunc := func(entry bin.Address, name string) {
		// TODO: Mark function signature as unknown (using metadata), so that type
		// analysis may replace it.
		name = fmt.Sprintf("_imp_%s", name)
		sig := types.NewFunc(types.Void)
		typ := types.NewPointer(sig)
		f := &ir.Func{
			Typ: typ,
			Sig: sig,
		}
		f.SetName(name)
		md := &metadata.Attachment{
			Name: "addr",
			Node: &metadata.Tuple{
				Fields: []metadata.Field{&metadata.String{Value: entry.String()}},
			},
		}
		f.Metadata = append(f.Metadata, md)
		fn := &Func{
			Func: f,
		}
		l.Funcs[entry] = fn
	}
	for entry, fname := range l.File.Imports {
		if _, ok := l.Funcs[entry]; ok {
			// Skip import if already specified through function signature.
			continue
		}
		dbg.Printf("function import at %v: %v\n", entry, fname)
		addFunc(entry, fname)
	}

	// Parse exports.
	for entry, fname := range dis.File.Exports {
		if _, ok := l.Funcs[entry]; ok {
			// Skip export if already specified through function signature.
			continue
		}
		addFunc(entry, fname)
	}

	return l, nil
}

// ### [ Helper functions ] ####################################################

// parseModule parses and returns the given LLVM IR module.
func parseModule(llPath string) (*ir.Module, error) {
	if !osutil.Exists(llPath) {
		warn.Printf("unable to locate LLVM IR file %q", llPath)
		return &ir.Module{}, nil
	}
	return asm.ParseFile(llPath)
}

// findMetadataAttachment locates the metadata node of the given metadata
// attachment. The boolean return value indicates success.
func findMetadataAttachment(mds []*metadata.Attachment, name string) (metadata.MDNode, bool) {
	for _, md := range mds {
		if md.Name == name {
			return md.Node, true
		}
	}
	return nil, false
}

// parseMetadataAddr returns the address corresponding to the given "addr"
// metadata node.
func parseMetadataAddr(node metadata.MDNode) (bin.Address, error) {
	switch node := node.(type) {
	case *metadata.Tuple:
		if len(node.Fields) != 1 {
			return 0, errors.Errorf(`invalid number of fields in "addr" metadata node, expected 1, got %d`, len(node.Fields))
		}
		field, ok := node.Fields[0].(*metadata.String)
		if !ok {
			panic(fmt.Errorf("invalid metadata field type; expected *metadata.String, got %T", node.Fields[0]))
		}
		var addr bin.Address
		if err := addr.Set(field.Value); err != nil {
			return 0, errors.WithStack(err)
		}
		return addr, nil
	default:
		panic(fmt.Errorf("support for metadata node %T not yet implemented", node))
	}
}
