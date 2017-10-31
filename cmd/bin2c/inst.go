// TODO: Handle flags for all instructions.

package main

import (
	"fmt"
	"go/ast"
	"go/token"

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
	case x86asm.JE:
		return parseJE(inst, offset)
	case x86asm.JG:
		return parseJG(inst, offset)
	case x86asm.JGE:
		return parseJGE(inst, offset)
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
	case x86asm.MOVSX:
		return parseMOVSX(inst)
	case x86asm.MOVZX:
		return parseMOVZX(inst)
	case x86asm.OR:
		return parseBinaryInst(inst, token.OR)
	case x86asm.RET:
		return parseRET(inst)
	case x86asm.SUB:
		return parseBinaryInst(inst, token.SUB)
	case x86asm.TEST:
		return parseTEST(inst)
	case x86asm.XOR:
		return parseBinaryInst(inst, token.XOR)
	case x86asm.PUSH, x86asm.POP, x86asm.LEAVE:
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
	target := baseAddr + offset
	lhs := getReg(x86asm.EAX)
	// TODO: Figure out how to identify the calling convention.
	rhs := &ast.CallExpr{
		Fun: ast.NewIdent(fmt.Sprintf("sub_%08X", target)),
	}
	stmt := createAssign(lhs, rhs)
	return stmt, nil
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
	var x, y, z ast.Expr
	x = getArg(inst.Args[0])
	y = getArg(inst.Args[1])
	if inst.Args[2] != nil {
		z = getArg(inst.Args[2])
	} else {
		z = y
	}

	// Create statement.
	//    x = y * z
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

// From http://faydoc.tripod.com/cpu/jge.htm
//
//    * [ ] JA       Jump if above (CF=0 and ZF=0)
//    * [ ] JAE      Jump if above or equal (CF=0)
//    * [ ] JB       Jump if below (CF=1)
//    * [ ] JBE      Jump if below or equal (CF=1 or ZF=1)
//    * [ ] JC       Jump if carry (CF=1)
//    * [ ] JCXZ     Jump if CX register is 0
//    * [x] JE       Jump if equal (ZF=1)
//    * [ ] JECXZ    Jump if ECX register is 0
//    * [x] JG       Jump if greater (ZF=0 and SF=OF)
//    * [x] JGE      Jump if greater or equal (SF=OF)
//    * [x] JL       Jump if less (SF!=OF)
//    * [x] JLE      Jump if less or equal (ZF=1 or SF!=OF)
//    * [ ] JNA      Jump if not above (CF=1 or ZF=1)
//    * [ ] JNAE     Jump if not above or equal (CF=1)
//    * [ ] JNB      Jump if not below (CF=0)
//    * [ ] JNBE     Jump if not below or equal (CF=0 and ZF=0)
//    * [ ] JNC      Jump if not carry (CF=0)
//    * [x] JNE      Jump if not equal (ZF=0)
//    * [ ] JNG      Jump if not greater (ZF=1 or SF!=OF)
//    * [ ] JNGE     Jump if not greater or equal (SF!=OF)
//    * [ ] JNL      Jump if not less (SF=OF)
//    * [ ] JNLE     Jump if not less or equal (ZF=0 and SF=OF)
//    * [ ] JNO      Jump if not overflow (OF=0)
//    * [ ] JNP      Jump if not parity (PF=0)
//    * [ ] JNS      Jump if not sign (SF=0)
//    * [ ] JNZ      Jump if not zero (ZF=0)
//    * [ ] JO       Jump if overflow (OF=1)
//    * [ ] JP       Jump if parity (PF=1)
//    * [ ] JPE      Jump if parity even (PF=1)
//    * [ ] JPO      Jump if parity odd (PF=0)
//    * [ ] JS       Jump if sign (SF=1)
//    * [ ] JZ       Jump if zero (ZF=1)

// parseJE parses the given JE instruction and returns a corresponding Go
// statement.
func parseJE(inst x86asm.Inst, offset int) (ast.Stmt, error) {
	// Parse arguments.
	arg := inst.Args[0]
	switch arg := arg.(type) {
	case x86asm.Rel:
		offset += inst.Len + int(arg)
	default:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	}

	// JE       Jump if equal (ZF=1)

	// Create statement.
	//    if zf {
	//       goto x
	//    }
	cond := getFlag(ZF)
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

// parseJG parses the given JG instruction and returns a corresponding Go
// statement.
func parseJG(inst x86asm.Inst, offset int) (ast.Stmt, error) {
	// Parse arguments.
	arg := inst.Args[0]
	switch arg := arg.(type) {
	case x86asm.Rel:
		offset += inst.Len + int(arg)
	default:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	}

	// JG      Jump if greater (ZF=0 and SF=OF)

	// Create statement.
	//    if !zf && sf == of
	//       goto x
	//    }
	expr := &ast.BinaryExpr{
		X:  getFlag(SF),
		Op: token.EQL,
		Y:  getFlag(OF),
	}
	cond := &ast.BinaryExpr{
		X: &ast.UnaryExpr{
			Op: token.NOT,
			X:  getFlag(ZF),
		},
		Op: token.LOR,
		Y:  expr,
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

// parseJGE parses the given JGE instruction and returns a corresponding Go
// statement.
func parseJGE(inst x86asm.Inst, offset int) (ast.Stmt, error) {
	// Parse arguments.
	arg := inst.Args[0]
	switch arg := arg.(type) {
	case x86asm.Rel:
		offset += inst.Len + int(arg)
	default:
		return nil, errutil.Newf("support for type %T not yet implemented", arg)
	}

	// JGE      Jump if greater or equal (SF=OF)

	// Create statement.
	//    if sf == of {
	//       goto x
	//    }
	cond := &ast.BinaryExpr{
		X:  getFlag(SF),
		Op: token.EQL,
		Y:  getFlag(OF),
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

	// JL       Jump if less (SF!=OF)

	// Create statement.
	//    if sf != of {
	//       goto x
	//    }
	cond := &ast.BinaryExpr{
		X:  getFlag(SF),
		Op: token.NEQ,
		Y:  getFlag(OF),
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

	// JLE      Jump if less or equal (ZF=1 or SF!=OF)

	// Create statement.
	//    if zf || sf != of
	//       goto x
	//    }
	expr := &ast.BinaryExpr{
		X:  getFlag(SF),
		Op: token.NEQ,
		Y:  getFlag(OF),
	}
	cond := &ast.BinaryExpr{
		X:  getFlag(ZF),
		Op: token.LOR,
		Y:  expr,
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

	// JNE      Jump if not equal (ZF=0)

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

// parseMOVSX parses the given MOVSX instruction and returns a corresponding Go
// statement.
func parseMOVSX(inst x86asm.Inst) (ast.Stmt, error) {
	// Parse arguments.
	x := getArg(inst.Args[0])
	y := getArg(inst.Args[1])

	// Create statement.
	//    x = y
	lhs := x
	rhs := y
	return createAssign(lhs, rhs), nil
}

// parseMOVZX parses the given MOVZX instruction and returns a corresponding Go
// statement.
func parseMOVZX(inst x86asm.Inst) (ast.Stmt, error) {
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

// parseTEST parses the given TEST instruction and returns a corresponding Go
// statement.
func parseTEST(inst x86asm.Inst) (ast.Stmt, error) {
	// Parse arguments.
	x := getArg(inst.Args[0])
	y := getArg(inst.Args[1])

	// Create statement.
	//    zf = (x&y) == 0
	lhs := getFlag(ZF)
	expr := createBinaryExpr(x, y, token.AND)
	zero := &ast.BasicLit{Kind: token.INT, Value: "0"}
	rhs := createBinaryExpr(expr, zero, token.EQL)
	stmt1 := createAssign(lhs, rhs)

	// Create statement.
	//    sf = (x&y)>>31 == 1
	lhs = getFlag(SF)
	expr = createBinaryExpr(x, y, token.AND)
	thirtyone := &ast.BasicLit{Kind: token.INT, Value: "31"}
	expr = createBinaryExpr(expr, thirtyone, token.SHR)
	one := &ast.BasicLit{Kind: token.INT, Value: "1"}
	rhs = createBinaryExpr(expr, one, token.EQL)
	stmt2 := createAssign(lhs, rhs)

	// Create statement.
	//    of = 0
	lhs = getFlag(OF)
	rhs = zero
	stmt3 := createAssign(lhs, rhs)

	// Create statement.
	//    cf = 0
	lhs = getFlag(CF)
	rhs = zero
	stmt4 := createAssign(lhs, rhs)

	// TODO: Set remaining flags.

	// Create block statement.
	stmt := &ast.BlockStmt{
		List: []ast.Stmt{stmt1, stmt2, stmt3, stmt4},
	}
	return stmt, nil
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
