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
		fmt.Println("inst:", inst)

		// Parse instruction.
		stmt, err := parseInst(inst)
		if err != nil {
			return errutil.Err(err)
		}
		if stmt != nil {
			label := ast.NewIdent(fmt.Sprintf("loc_%X", base+offset))
			stmt = &ast.LabeledStmt{Label: label, Stmt: stmt}
			fmt.Println("stmt:", stmt)
			//ast.Print(token.NewFileSet(), stmt)
			printer.Fprint(os.Stdout, token.NewFileSet(), stmt)
			fmt.Println()
		}

		// Next.
		offset += inst.Len
		if inst.Op == x86asm.RET {
			break
		}
	}
	return nil
}

// parseInst parses the given assembly instruction and returns a corresponding
// Go statement.
func parseInst(inst x86asm.Inst) (ast.Stmt, error) {
	switch inst.Op {
	case x86asm.RET:
		return &ast.ReturnStmt{}, nil
	case x86asm.XOR:
		return parseBinaryInst(inst)
	case x86asm.PUSH, x86asm.POP:
		// ignore for now.
	default:
		return nil, errutil.Newf("support for opcode %v not yet implemented", inst.Op)
	}
	return nil, nil
}

// regNames maps from register names to their corresponding Go identifiers.
var regNames = map[string]*ast.Ident{
	"EAX": ast.NewIdent("eax"),
}

// parseArg parses the given assembly instruction argument and returns a
// corresponding Go expression.
func parseArg(arg x86asm.Arg) (ast.Expr, error) {
	switch arg := arg.(type) {
	case x86asm.Reg:
		reg, ok := regNames[arg.String()]
		if !ok {
			return nil, errutil.Newf("unable to lookup identifer for register %q", arg)
		}
		return reg, nil
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	case x86asm.Mem:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	case x86asm.Imm:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	case x86asm.Rel:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	default:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	}
}

// parseBinaryInst parses the given binary instruction and returns a
// corresponding Go statement.
func parseBinaryInst(inst x86asm.Inst) (ast.Stmt, error) {
	x, err := parseArg(inst.Args[0])
	if err != nil {
		return nil, errutil.Err(err)
	}
	y, err := parseArg(inst.Args[1])
	if err != nil {
		return nil, errutil.Err(err)
	}
	var op token.Token
	switch inst.Op {
	case x86asm.XOR:
		op = token.XOR
	default:
		return nil, errutil.Newf("support for opcode %v not yet implemented", inst.Op)
	}
	rhs := &ast.BinaryExpr{X: x, Op: op, Y: y}
	lhs := x
	assign := &ast.AssignStmt{
		Lhs: []ast.Expr{lhs},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{rhs},
	}
	return assign, nil
}
