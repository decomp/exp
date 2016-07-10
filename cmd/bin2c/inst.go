// TODO: Handle flags for all instructions.

package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"

	"github.com/mewkiz/pkg/errutil"
	"golang.org/x/arch/x86/x86asm"
)

// parseInst parses the given assembly instruction and returns a corresponding
// Go statement.
func parseInst(inst x86asm.Inst, offset int) (ast.Stmt, error) {
	switch inst.Op {
	case x86asm.ADD:
		return parseBinaryInst(inst, token.ADD)
	case x86asm.AND:
		return parseBinaryInst(inst, token.AND)
	case x86asm.CALL:
		return parseCALL(inst, offset)
	case x86asm.CMP:
		return parseCMP(inst)
	case x86asm.DEC:
		return parseDEC(inst)
	case x86asm.IMUL:
		return parseIMUL(inst)
	case x86asm.INC:
		return parseINC(inst)
	case x86asm.JL:
		return parseJL(inst, offset)
	case x86asm.JLE:
		return parseJLE(inst, offset)
	case x86asm.JMP:
		return parseJMP(inst, offset)
	case x86asm.JNE:
		return parseJNE(inst, offset)
	case x86asm.LEA:
		return parseLEA(inst)
	case x86asm.MOV:
		return parseMOV(inst)
	case x86asm.RET:
		return parseRET(inst)
	case x86asm.SUB:
		return parseBinaryInst(inst, token.SUB)
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

// parseCALL parses the given CALL instruction and returns a corresponding Go
// statement.
func parseCALL(inst x86asm.Inst, offset int) (ast.Stmt, error) {
	// Parse arguments.
	arg := inst.Args[0]
	switch arg := arg.(type) {
	case x86asm.Rel:
		offset += inst.Len + int(arg)
	default:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	}
	log.Println("target:", offset)

	// TODO: Figure out how to identify the calling convention.
	return nil, errutil.New("not yet implemented")
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
	rhs := createBinaryExpr(x, y, token.EQL)
	stmt1 := createAssign(lhs, rhs)

	// Create statement.
	//    cf = x < y
	lhs = getFlag(CF)
	rhs = createBinaryExpr(x, y, token.LSS)
	stmt2 := createAssign(lhs, rhs)

	// Create block statement.
	stmt := &ast.BlockStmt{
		List: []ast.Stmt{stmt1, stmt2},
	}
	return stmt, nil
}

// parseDEC parses the given DEC instruction and returns a corresponding Go
// statement.
func parseDEC(inst x86asm.Inst) (ast.Stmt, error) {
	// Parse arguments.
	x := getArg(inst.Args[0])

	// Create statement.
	//    x--
	stmt1 := &ast.IncDecStmt{
		X:   x,
		Tok: token.DEC,
	}

	// Create statement.
	//    zf = x == 0
	lhs := getFlag(ZF)
	rhs := createBinaryExpr(x, createExpr(0), token.EQL)
	stmt2 := createAssign(lhs, rhs)

	// TODO: Find a better solution for multiple statement than block statement.

	// Create block statement.
	stmt := &ast.BlockStmt{
		List: []ast.Stmt{stmt1, stmt2},
	}
	return stmt, nil
}

// parseIMUL parses the given IMUL instruction and returns a corresponding Go
// statement.
func parseIMUL(inst x86asm.Inst) (ast.Stmt, error) {
	// Parse arguments.
	x := getArg(inst.Args[0])
	y := getArg(inst.Args[1])
	z := getArg(inst.Args[2])

	// Create statement.
	//    x = x OP y
	lhs := x
	rhs := createBinaryExpr(y, z, token.MUL)
	return createAssign(lhs, rhs), nil
}

// parseINC parses the given INC instruction and returns a corresponding Go
// statement.
func parseINC(inst x86asm.Inst) (ast.Stmt, error) {
	// Parse arguments.
	x := getArg(inst.Args[0])

	// Create statement.
	//    x++
	stmt := &ast.IncDecStmt{
		X:   x,
		Tok: token.INC,
	}
	return stmt, nil
}

// parseJL parses the given JL instruction and returns a corresponding Go
// statement.
func parseJL(inst x86asm.Inst, offset int) (ast.Stmt, error) {
	// Parse arguments.
	arg := inst.Args[0]
	switch arg := arg.(type) {
	case x86asm.Rel:
		offset += inst.Len + int(arg)
	default:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	}

	// Create statement.
	//    if cf {
	//       goto x
	//    }
	cond := getFlag(CF)
	label := getLabel("loc", offset)
	body := &ast.BranchStmt{
		Tok:   token.GOTO,
		Label: label,
	}
	stmt := &ast.IfStmt{
		Cond: cond,
		Body: &ast.BlockStmt{List: []ast.Stmt{body}},
	}
	return stmt, nil
}

// parseJLE parses the given JLE instruction and returns a corresponding Go
// statement.
func parseJLE(inst x86asm.Inst, offset int) (ast.Stmt, error) {
	// Parse arguments.
	arg := inst.Args[0]
	switch arg := arg.(type) {
	case x86asm.Rel:
		offset += inst.Len + int(arg)
	default:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	}

	// Create statement.
	//    if cf || zf {
	//       goto x
	//    }
	cond := &ast.BinaryExpr{
		X:  getFlag(CF),
		Op: token.LOR,
		Y:  getFlag(ZF),
	}
	label := getLabel("loc", offset)
	body := &ast.BranchStmt{
		Tok:   token.GOTO,
		Label: label,
	}
	stmt := &ast.IfStmt{
		Cond: cond,
		Body: &ast.BlockStmt{List: []ast.Stmt{body}},
	}
	return stmt, nil
}

// parseJMP parses the given JMP instruction and returns a corresponding Go
// statement.
func parseJMP(inst x86asm.Inst, offset int) (ast.Stmt, error) {
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
	label := getLabel("loc", offset)
	stmt := &ast.BranchStmt{
		Tok:   token.GOTO,
		Label: label,
	}
	return stmt, nil
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
	//    if !zf {
	//       goto x
	//    }
	cond := &ast.UnaryExpr{
		Op: token.NOT,
		X:  getFlag(ZF),
	}
	label := getLabel("loc", offset)
	body := &ast.BranchStmt{
		Tok:   token.GOTO,
		Label: label,
	}
	stmt := &ast.IfStmt{
		Cond: cond,
		Body: &ast.BlockStmt{List: []ast.Stmt{body}},
	}
	return stmt, nil
}

// parseLEA parses the given LEA instruction and returns a corresponding Go
// statement.
func parseLEA(inst x86asm.Inst) (ast.Stmt, error) {
	// Parse arguments.
	x := getArg(inst.Args[0])
	y := getArg(inst.Args[1])

	// Create statement.
	//    x = &y
	lhs := x
	rhs, err := unstar(y)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return createAssign(lhs, rhs), nil
}

// unstar returns the underlying expression from the given parenthesised star
// expression.
func unstar(expr ast.Expr) (ast.Expr, error) {
	star, ok := expr.(*ast.StarExpr)
	if !ok {
		return nil, errutil.Newf("invalid argument type; expected *ast.StarExpr, got %T", expr)
	}
	paren, ok := star.X.(*ast.ParenExpr)
	if !ok {
		return nil, errutil.Newf("invalid argument type; expected *ast.ParenExpr, got %T", star.X)
	}
	return paren.X, nil
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
	return createAssign(lhs, rhs), nil
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

	// Simplify instructions when x == y.
	if inst.Args[0] == inst.Args[1] {
		switch op {
		case token.SUB, token.XOR:
			// Before:
			//    x = x - x
			//    x = x ^ x
			// After:
			//    x = 0
			return createAssign(x, createExpr(0)), nil
		}
	}

	// Create statement.
	//    x = x OP y
	lhs := x
	rhs := createBinaryExpr(x, y, op)
	return createAssign(lhs, rhs), nil
}

// createBinaryExpr returns a binary expression with the given operands and
// operation. Special consideration is taken with regards to sub-registers (e.g.
// al, ah, ax).
func createBinaryExpr(x, y ast.Expr, op token.Token) ast.Expr {
	// Handle sub-registers (e.g. al, ah, ax).
	x = fromSubReg(x)
	y = fromSubReg(y)

	return &ast.BinaryExpr{
		X:  x,
		Op: op,
		Y:  y,
	}
}

// createAssign returns an assignment statement with the given left- and right-
// hand sides. Special consideration is taken with regards to sub-registers.
// (e.g. al, ah, ax).
func createAssign(lhs, rhs ast.Expr) ast.Stmt {
	rhs = fromSubReg(rhs)
	// TODO: Handle sub-registers (al, ah, ax)
	return &ast.AssignStmt{
		Lhs: []ast.Expr{lhs},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{rhs},
	}
}
