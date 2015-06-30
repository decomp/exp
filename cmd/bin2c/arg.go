package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"strconv"

	"github.com/mewkiz/pkg/errutil"
	"rsc.io/x86/x86asm"
)

// getArg converts arg into a corresponding Go expression.
func getArg(arg x86asm.Arg) ast.Expr {
	switch arg := arg.(type) {
	case x86asm.Reg:
		return getReg(arg)
	case x86asm.Mem:
		return getMem(arg)
	case x86asm.Imm:
		// TODO: Implement support for immediate values.
	case x86asm.Rel:
		// TODO: Implement support for relative addresses.
	}
	fmt.Printf("%#v\n", arg)
	log.Fatal(errutil.Newf("support for type %T not yet implemented", arg))
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

// getMem converts mem into a corresponding Go expression.
func getMem(mem x86asm.Mem) ast.Expr {
	// The general memory reference form is:
	//    Segment:[Base+Scale*Index+Disp]

	// ... + Disp
	expr := &ast.BinaryExpr{}
	if mem.Disp != 0 {
		disp := getExpr(mem.Disp)
		expr.Op = token.ADD
		expr.Y = disp
	}

	// ... + (Scale*Index) + ...
	if mem.Scale != 0 && mem.Index != 0 {
		scale := getExpr(mem.Scale)
		index := getReg(mem.Index)
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
	if mem.Base != 0 {
		base := getReg(mem.Base)
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
	if mem.Segment != 0 {
		segment := getReg(mem.Segment)
		_ = segment
		fmt.Printf("%#v\n", mem)
		log.Fatal(errutil.Newf("support for Mem.Segment not yet implemented"))
	}

	switch {
	case expr.X == nil && expr.Y == nil:
		fmt.Printf("%#v\n", mem)
		log.Fatal(errutil.New("support for memory reference to address zero not yet implemented"))
		panic("unreachable")
	case expr.X == nil && expr.Y != nil:
		return expr.Y
	case expr.X != nil && expr.Y == nil:
		return expr.X
	default:
		return expr
	}
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
