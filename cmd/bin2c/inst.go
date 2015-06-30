package main

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/mewkiz/pkg/errutil"
	"rsc.io/x86/x86asm"
)

// parseInst parses the given assembly instruction and returns a corresponding
// Go statement.
func parseInst(inst x86asm.Inst) (ast.Stmt, error) {
	switch inst.Op {
	case x86asm.RET:
		return &ast.ReturnStmt{}, nil
	case x86asm.LEA, x86asm.MOV:
		return parseAssignInst(inst)
	case x86asm.XOR:
		return parseBinaryInst(inst)
	case x86asm.PUSH, x86asm.POP:
		// ignore for now.
		return nil, nil
	default:
		fmt.Printf("%#v\n", inst)
		return nil, errutil.Newf("support for opcode %v not yet implemented", inst.Op)
	}
}

// parseAssignInst parses the given assignment instruction and returns a
// corresponding Go statement.
func parseAssignInst(inst x86asm.Inst) (ast.Stmt, error) {
	x := getArg(inst.Args[0])
	y := getArg(inst.Args[1])
	lhs := x
	rhs := y
	assign := &ast.AssignStmt{
		Lhs: []ast.Expr{lhs},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{rhs},
	}
	return assign, nil
}

// parseBinaryInst parses the given binary instruction and returns a
// corresponding Go statement.
func parseBinaryInst(inst x86asm.Inst) (ast.Stmt, error) {
	x := getArg(inst.Args[0])
	y := getArg(inst.Args[1])
	var op token.Token
	switch inst.Op {
	case x86asm.XOR:
		op = token.XOR
	default:
		return nil, errutil.Newf("support for opcode %v not yet implemented", inst.Op)
	}
	lhs := x
	rhs := &ast.BinaryExpr{X: x, Op: op, Y: y}
	assign := &ast.AssignStmt{
		Lhs: []ast.Expr{lhs},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{rhs},
	}
	return assign, nil
}
