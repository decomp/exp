package main

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strconv"

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
			label := ast.NewIdent(fmt.Sprintf("loc_%X", baseAddr+offset))
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
	case x86asm.LEA:
		return parseLEA(inst)
	case x86asm.XOR:
		return parseBinaryInst(inst)
	case x86asm.PUSH, x86asm.POP:
		// ignore for now.
	default:
		fmt.Printf("%#v\n", inst)
		return nil, errutil.Newf("support for opcode %v not yet implemented", inst.Op)
	}
	return nil, nil
}

// parseBinaryInst parses the given LEA instruction and returns a corresponding
// Go statement.
func parseLEA(inst x86asm.Inst) (ast.Stmt, error) {
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

// getArg converts arg into a corresponding Go expression.
func getArg(arg x86asm.Arg) ast.Expr {
	switch arg := arg.(type) {
	case x86asm.Reg:
		return getReg(arg)
	case x86asm.Mem:
		// The general memory reference form is:
		//    Segment:[Base+Scale*Index+Disp]

		// ... + Disp
		expr := &ast.BinaryExpr{}
		if arg.Disp != 0 {
			disp := getExpr(arg.Disp)
			expr.Op = token.ADD
			expr.Y = disp
		}

		// ... + (Scale*Index) + ...
		if arg.Scale != 0 && arg.Index != 0 {
			scale := getExpr(arg.Scale)
			index := getReg(arg.Index)
			product := &ast.BinaryExpr{
				X:  scale,
				Op: token.MUL,
				Y:  index,
			}
			switch {
			case expr.Y == nil:
				// ... + (Scale*Index)
				expr.Op = token.ADD
				expr.Y = product
			default:
				// ... + (Scale*Index) + Disp
				expr.X = product
				expr.Op = token.ADD
			}
		}

		// ... + Base + ...
		if arg.Base != 0 {
			base := getReg(arg.Base)
			switch {
			case expr.X == nil:
				// Base + (Scale*Index)
				// or
				// Base + Disp
				expr.X = base
				expr.Op = token.ADD
			case expr.Y == nil:
				// ... + Base
				expr.Op = token.ADD
				expr.Y = base
			default:
				sum := &ast.BinaryExpr{
					X:  expr.X,
					Op: token.ADD,
					Y:  expr.Y,
				}
				expr.X = base
				expr.Op = token.ADD
				expr.Y = sum
			}
		}

		// TODO: Figure out how the calculation is affected by segment in:
		//    Segment:[Base+Scale*Index+Disp]
		if arg.Segment != 0 {
			segment := getReg(arg.Segment)
			_ = segment
			fmt.Printf("%#v\n", arg)
			log.Fatal(errutil.Newf("support for Mem.Segment not yet implemented"))
		}

		switch {
		case expr.X == nil && expr.Y == nil:
			fmt.Printf("%#v\n", arg)
			log.Fatal(errutil.New("support for memory reference to address zero not yet implemented"))
		case expr.X == nil && expr.Y != nil:
			return expr.Y
		case expr.X != nil && expr.Y == nil:
			return expr.X
		case expr.X != nil && expr.Y != nil:
			return expr
		}
	case x86asm.Imm:
		// TODO: Implement support for immediate values.
	case x86asm.Rel:
		// TODO: Implement support for relative addresses.
	}
	fmt.Printf("%#v\n", arg)
	log.Fatal(errutil.Newf("support for type %T not yet implemented", arg))
	panic("unreachable")
}

// getExpr converts x into a corresponding Go expression.
func getExpr(x interface{}) ast.Expr {
	switch x := x.(type) {
	case uint8:
		s := strconv.FormatUint(uint64(x), 10)
		return &ast.BasicLit{Kind: token.INT, Value: s}
	case int64:
		s := strconv.FormatInt(x, 10)
		return &ast.BasicLit{Kind: token.INT, Value: s}
	}
	log.Fatal(errutil.Newf("support for type %T not yet implemented", x))
	panic("unreachable")
}

// getReg converts reg into a corresponding Go expression.
func getReg(reg x86asm.Reg) ast.Expr {
	// regNames maps from register names to their corresponding Go identifiers.
	var regNames = map[string]*ast.Ident{
		"EAX": ast.NewIdent("eax"),
		"EDI": ast.NewIdent("edi"),
	}
	if expr, ok := regNames[reg.String()]; ok {
		return expr
	}
	log.Fatal(errutil.Newf("unable to lookup identifer for register %q", reg))
	panic("unreachable")
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
