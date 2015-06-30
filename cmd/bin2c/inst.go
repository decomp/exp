// TODO: Handle flags for all instructions.

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
func parseInst(inst x86asm.Inst, offset int) (ast.Stmt, error) {
	switch inst.Op {
	case x86asm.ADD:
		return parseBinaryInst(inst, token.ADD)
	case x86asm.CMP:
		return parseCMP(inst)
	case x86asm.JNE:
		return parseJNE(inst, offset)
	case x86asm.LEA:
		return parseLEA(inst)
	case x86asm.MOV:
		return parseMOV(inst)
	case x86asm.RET:
		return parseRET(inst)
	case x86asm.XOR:
		return parseBinaryInst(inst, token.XOR)
	case x86asm.PUSH, x86asm.POP:
		// ignore for now.
		return nil, nil
	default:
		fmt.Printf("%#v\n", inst)
		return nil, errutil.Newf("support for opcode %v not yet implemented", inst.Op)
	}
}

// parseCMP parses the given CMP instruction and returns a corresponding Go
// statement.
func parseCMP(inst x86asm.Inst) (ast.Stmt, error) {
	// Parse arguments.
	x := getArg(inst.Args[0])
	y := getArg(inst.Args[1])

	// Create statement.
	//    zf = x == y
	lhs := getFlag(ZF)
	rhs := &ast.BinaryExpr{
		X:  x,
		Op: token.EQL,
		Y:  y,
	}
	return getAssign(lhs, rhs), nil
}

// parseJNE parses the given JNE instruction and returns a corresponding Go
// statement.
func parseJNE(inst x86asm.Inst, offset int) (ast.Stmt, error) {
	// Parse arguments.
	arg := inst.Args[0]
	switch arg := arg.(type) {
	case x86asm.Rel:
		offset += inst.Len + int(arg)
	default:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	}

	// Create statement.
	//    goto x
	label := getLabel(offset)
	stmt := &ast.BranchStmt{
		Tok:   token.GOTO,
		Label: label,
	}
	return stmt, nil
}

// parseLEA parses the given LEA instruction and returns a corresponding Go
// statement.
func parseLEA(inst x86asm.Inst) (ast.Stmt, error) {
	// Parse arguments.
	x := getArg(inst.Args[0])
	y := getArg(inst.Args[1])
	star, ok := y.(*ast.StarExpr)
	if !ok {
		return nil, errutil.Newf("invalid argument type; expected *ast.StarExpr, got %T", y)
	}
	paren, ok := star.X.(*ast.ParenExpr)
	if !ok {
		return nil, errutil.Newf("invalid argument type; expected *ast.ParenExpr, got %T", star.X)
	}

	// Create statement.
	//    x = &y
	lhs := x
	rhs := paren.X
	return getAssign(lhs, rhs), nil
}

// parseMOV parses the given MOV instruction and returns a corresponding Go
// statement.
func parseMOV(inst x86asm.Inst) (ast.Stmt, error) {
	// Parse arguments.
	x := getArg(inst.Args[0])
	y := getArg(inst.Args[1])

	// Create statement.
	//    x = y
	lhs := x
	rhs := y
	return getAssign(lhs, rhs), nil
}

// parseRET parses the given RET instruction and returns a corresponding Go
// statement.
func parseRET(inst x86asm.Inst) (ast.Stmt, error) {
	// TODO: Handle pops; e.g.
	//    ret 0xC

	// Create statement.
	//    return
	return &ast.ReturnStmt{}, nil
}

// parseBinaryInst parses the given binary instruction and returns a
// corresponding Go statement.
func parseBinaryInst(inst x86asm.Inst, op token.Token) (ast.Stmt, error) {
	// Parse arguments.
	x := getArg(inst.Args[0])
	y := getArg(inst.Args[1])

	// Create statement.
	//    x = x OP y
	lhs := x
	rhs := &ast.BinaryExpr{
		X:  x,
		Op: op,
		Y:  y,
	}
	return getAssign(lhs, rhs), nil
}

// getAssign returns an assignment statement with the given left- and right-hand
// sides.
func getAssign(lhs, rhs ast.Expr) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{lhs},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{rhs},
	}
}
