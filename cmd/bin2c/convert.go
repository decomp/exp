package main

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"

	"github.com/mewkiz/pkg/errutil"
	"rsc.io/x86/x86asm"
)

// convert converts the given binary excutable to equivalent C source code.
func convert(text []byte) error {
	panic("not yet implemented")
}

// convertFunc converts the function at the given offset in text to C source
// code.
func convertFunc(text []byte, offset int) error {
	fn := &ast.FuncDecl{
		Name: getLabel("sub", offset),
		Type: &ast.FuncType{},
		Body: new(ast.BlockStmt),
	}

	for {
		// Decode instruction.
		inst, err := x86asm.Decode(text[offset:], 32)
		if err != nil {
			return errutil.Err(err)
		}
		log.Println("==================================")
		log.Println("inst:", inst)

		// Parse instruction.
		stmt, err := parseInst(inst, offset)
		if err != nil {
			return errutil.Err(err)
		}
		if stmt != nil {
			// Access the underlying statement list of the block statement.
			if block, ok := stmt.(*ast.BlockStmt); ok {
				label := getLabel("loc", offset)
				stmt = block.List[0]
				stmt = &ast.LabeledStmt{Label: label, Stmt: stmt}
				block.List[0] = stmt
				fn.Body.List = append(fn.Body.List, block.List...)
				for _, stmt := range block.List {
					log.Println("stmt:", stmt)
					//ast.Print(token.NewFileSet(), stmt)
					printer.Fprint(os.Stderr, token.NewFileSet(), stmt)
					log.Println()
				}
			} else {
				label := getLabel("loc", offset)
				stmt = &ast.LabeledStmt{Label: label, Stmt: stmt}
				fn.Body.List = append(fn.Body.List, stmt)
				log.Println("stmt:", stmt)
				//ast.Print(token.NewFileSet(), stmt)
				printer.Fprint(os.Stderr, token.NewFileSet(), stmt)
				log.Println()
			}
		}
		log.Println()

		// Next.
		offset += inst.Len
		if inst.Op == x86asm.RET {
			break
		}
	}

	printer.Fprint(os.Stdout, token.NewFileSet(), fn)

	return nil
}

// labels maps from offset to label identifiers.
var labels = map[int]*ast.Ident{}

// getLabel returns the label with the given prefix and at the given offset.
func getLabel(prefix string, offset int) *ast.Ident {
	if label, ok := labels[offset]; ok {
		return label
	}
	addr := baseAddr + offset
	name := fmt.Sprintf("%s_%X", prefix, addr)
	label := ast.NewIdent(name)
	labels[offset] = label
	return label
}
