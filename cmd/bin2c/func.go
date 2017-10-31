package main

import (
	"fmt"
	dbg "fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"

	"github.com/mewkiz/pkg/errutil"
	"golang.org/x/arch/x86/x86asm"
)

// parseFunc parses the given function and returns a corresponding Go function.
func parseFunc(text []byte, offset int) (*ast.FuncDecl, error) {
	fn := &ast.FuncDecl{
		Name: getLabel("sub", offset),
		Type: &ast.FuncType{},
		Body: new(ast.BlockStmt),
	}

	for {
		// Decode instruction.
		inst, err := x86asm.Decode(text[offset:], 32)
		if err != nil {
			return nil, errutil.Err(err)
		}
		dbg.Println("==================================")
		dbg.Println("inst:", inst)

		// Parse instruction.
		stmt, err := parseInst(inst, offset)
		if err != nil {
			return nil, errutil.Err(err)
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
					dbg.Println("stmt:", stmt)
					//ast.Print(token.NewFileSet(), stmt)
					printer.Fprint(os.Stderr, token.NewFileSet(), stmt)
					dbg.Println()
				}
			} else {
				label := getLabel("loc", offset)
				stmt = &ast.LabeledStmt{Label: label, Stmt: stmt}
				fn.Body.List = append(fn.Body.List, stmt)
				dbg.Println("stmt:", stmt)
				//ast.Print(token.NewFileSet(), stmt)
				printer.Fprint(os.Stderr, token.NewFileSet(), stmt)
				dbg.Println()
			}
		}
		dbg.Println()

		// Next.
		offset += inst.Len
		if inst.Op == x86asm.RET {
			break
		}
	}

	return fn, nil
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
