package main

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
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
	for {
		// Decode instruction.
		inst, err := x86asm.Decode(text[offset:], 32)
		if err != nil {
			return errutil.Err(err)
		}
		fmt.Println("==================================")
		fmt.Println("inst:", inst)

		// Parse instruction.
		stmt, err := parseInst(inst)
		if err != nil {
			return errutil.Err(err)
		}
		if stmt != nil {
			label := ast.NewIdent(fmt.Sprintf("loc_%X", baseAddr+offset))
			stmt = &ast.LabeledStmt{Label: label, Stmt: stmt}
			fmt.Println("stmt:", stmt)
			//ast.Print(token.NewFileSet(), stmt)
			printer.Fprint(os.Stdout, token.NewFileSet(), stmt)
			fmt.Println()
		}
		fmt.Println()

		// Next.
		offset += inst.Len
		if inst.Op == x86asm.RET {
			break
		}
	}
	return nil
}
