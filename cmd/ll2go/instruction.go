package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
	"unicode"

	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// parseInst converts the provided LLVM IR instruction into an equivalent Go AST
// node (a statement).
func parseInst(inst ir.Instruction) (ast.Stmt, error) {
	// TODO: Remove debug output.
	if flagVerbose {
		fmt.Println("parseInst:")
		fmt.Println()
		pretty.Println(inst)
		fmt.Println()
	}

	// Assignment operation.
	//    %foo = ...
	// Binary Operations
	switch inst := inst.(type) {
	case *ir.InstAdd:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.ADD)
	case *ir.InstFAdd:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.ADD)
	case *ir.InstSub:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.SUB)
	case *ir.InstFSub:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.SUB)
	case *ir.InstMul:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.MUL)
	case *ir.InstFMul:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.MUL)
	case *ir.InstUDiv:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.QUO)
	case *ir.InstSDiv:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.QUO)
	case *ir.InstFDiv:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.QUO)
	case *ir.InstURem:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.REM)
	case *ir.InstSRem:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.REM)
	case *ir.InstFRem:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.REM)

	// Bitwise Binary Operations
	case *ir.InstShl:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.SHL)
	case *ir.InstLShr:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.SHR)
	case *ir.InstAShr:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.SHR)
	case *ir.InstAnd:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.AND)
	case *ir.InstOr:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.OR)
	case *ir.InstXor:
		return parseBinOp(inst.Name, inst.X, inst.Y, token.XOR)

	// Other Operators
	case *ir.InstICmp:
		cond := getICmpPred(inst.Cond)
		return parseBinOp(inst.Name, inst.X, inst.Y, cond)
	case *ir.InstFCmp:
		cond := getFCmpPred(inst.Cond)
		return parseBinOp(inst.Name, inst.X, inst.Y, cond)

	default:
		panic(errors.Errorf("support for instruction %T not yet implemented", inst))

	}
}

// parseBinOp converts the provided LLVM IR binary operation into an equivalent
// Go AST node (an assignment statement with a binary expression on the right-
// hand side).
//
// Syntax:
//    <result> add <type> <op1>, <op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#binary-operations
func parseBinOp(result string, x, y value.Value, op token.Token) (ast.Stmt, error) {
	xx, err := parseOperand(x)
	if err != nil {
		return nil, err
	}
	yy, err := parseOperand(y)
	if err != nil {
		return nil, err
	}
	res := newIdent(result)
	lhs := []ast.Expr{res}
	rhs := []ast.Expr{&ast.BinaryExpr{X: xx, Op: op, Y: yy}}
	// TODO: Use "=" instead of ":=" and let go-post and grind handle the ":=" to
	// "=" propagation.
	return &ast.AssignStmt{Lhs: lhs, Tok: token.DEFINE, Rhs: rhs}, nil
}

// parseOperand converts the provided LLVM IR operand into an equivalent Go AST
// expression node (a basic literal, a composite literal or an identifier).
//
// Syntax:
//    i32 1
//    %foo = ...
func parseOperand(op value.Value) (ast.Expr, error) {
	// TODO: Support *BasicLit, *CompositeLit.

	// TODO: Add support for operand of other types than int.
	// TODO: Parse type.

	// Create and return a constant operand.
	//    i32 42
	switch op := op.(type) {
	case *constant.Int:
		return &ast.BasicLit{Kind: token.INT, Value: op.X.String()}, nil
	case value.Named:
		return newIdent(op.GetName()), nil
	default:
		panic(errors.Errorf("support for operand %T not yet implemented", op))
	}
}

// parseRetInst converts the provided LLVM IR ret instruction into an equivalent
// Go return statement.
//
// Syntax:
//    ret void
//    ret <type> <val>
func parseRetInst(term *ir.TermRet) (*ast.ReturnStmt, error) {
	// TODO: Make more robust by using proper parsing instead of relying on
	// tokens. The current approach is used for a proof of concept and would fail
	// for composite literals. This TODO applies to the use of tokens in all
	// functions.

	// Create and return a void return statement.
	if term.X == nil {
		return &ast.ReturnStmt{}, nil
	}

	// Create and return a return statement.
	val, err := parseOperand(term.X)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ret := &ast.ReturnStmt{
		Results: []ast.Expr{val},
	}
	return ret, nil
}

// A definition captures the semantics of a PHI instruction's right-hand side,
// i.e. it specifies a variable definition expression in relation to its source
// basic block.
type definition struct {
	// Source basic block of the variable definition.
	bb string
	// Variable definition expression.
	expr ast.Expr
}

// parsePHIInst converts the provided LLVM IR phi instruction into an equivalent
// variable definition mapping.
//
// Syntax:
//    %foo = phi i32 [ 42, %2 ], [ %bar, %3 ]
func parsePHIInst(inst *ir.InstPhi) (ident string, defs []*definition, err error) {
	// Parse operands.
	for _, inc := range inst.Incs {
		x, err := parseOperand(inc.X)
		if err != nil {
			return "", nil, errors.WithStack(err)
		}
		pred := inc.Pred.Name
		def := &definition{bb: pred, expr: x}
		defs = append(defs, def)
	}

	return inst.Name, defs, nil
}

// getICmpPred returns a Go token equivalent of the given integer comparison
// predicate.
func getICmpPred(pred ir.IntPred) token.Token {
	switch pred {
	// Int predicates.
	case ir.IntEQ: // eq: equal
		return token.EQL // ==
	case ir.IntNE: // ne: not equal
		return token.NEQ // !=
	case ir.IntUGT: // ugt: unsigned greater than
		return token.GTR // >
	case ir.IntUGE: // uge: unsigned greater or equal
		return token.GEQ // >=
	case ir.IntULT: // ult: unsigned less than
		return token.LSS // <
	case ir.IntULE: // ule: unsigned less or equal
		return token.LEQ // <=
	case ir.IntSGT: // sgt: signed greater than
		return token.GTR // >
	case ir.IntSGE: // sge: signed greater or equal
		return token.GEQ // >=
	case ir.IntSLT: // slt: signed less than
		return token.LSS // <
	case ir.IntSLE: // sle: signed less or equal
		return token.LEQ // <=

	default:
		panic(errors.Errorf("support for integer comparison predicate %v not yet implemented", pred))
	}
}

// getFCmpPred returns a Go token equivalent of the given floating-point
// comparison predicate.
func getFCmpPred(pred ir.FloatPred) token.Token {
	switch pred {
	// Float predicates.
	case ir.FloatFalse: // false: no comparison, always returns false
		panic(errors.Errorf(`support for the floating point comparison predicate "false" not yet implemented`))
	case ir.FloatOEQ: // oeq: ordered and equal
		return token.EQL
	case ir.FloatOGT: // ogt: ordered and greater than
		return token.GTR
	case ir.FloatOGE: // oge: ordered and greater than or equal
		return token.GEQ
	case ir.FloatOLT: // olt: ordered and less than
		return token.LSS
	case ir.FloatOLE: // ole: ordered and less than or equal
		return token.LEQ
	case ir.FloatONE: // one: ordered and not equal
		return token.NEQ
	case ir.FloatORD: // ord: ordered (no nans)
		panic(errors.Errorf(`support for the floating point comparison predicate "ord" not yet implemented`))
	case ir.FloatUEQ: // ueq: unordered or equal
		return token.EQL
	case ir.FloatUGT: // ugt: unordered or greater than
		return token.GTR
	case ir.FloatUGE: // uge: unordered or greater than or equal
		return token.GEQ
	case ir.FloatULT: // ult: unordered or less than
		return token.LSS
	case ir.FloatULE: // ule: unordered or less than or equal
		return token.LEQ
	case ir.FloatUNE: // une: unordered or not equal
		return token.NEQ
	case ir.FloatUNO: // uno: unordered (either nans)
		panic(errors.Errorf(`support for the floating point comparison predicate "uno" not yet implemented`))
	case ir.FloatTrue: // true: no comparison, always returns true
		panic(errors.Errorf(`support for the floating point comparison predicate "true" not yet implemented`))

	default:
		panic(errors.Errorf("support for floating-point comparison predicate %v not yet implemented", pred))
	}

}

// getBrCond parses the provided branch instruction and returns its condition.
//
// Syntax:
//    br i1 <cond>, label <target_true>, label <target_false>
func getBrCond(term ir.Terminator) (cond ast.Expr, targetTrue, targetFalse string, err error) {
	t, ok := term.(*ir.TermCondBr)
	if !ok {
		return nil, "", "", errors.Errorf("invalid conditional branch terminator type; expected *ir.TermCondBr, got %T", term)
	}
	// Create and return the condition.
	cond, err = parseOperand(t.Cond)
	if err != nil {
		return nil, "", "", errors.WithStack(err)
	}
	return cond, t.TargetTrue.Name, t.TargetFalse.Name, nil
}

// getResult returns the result identifier of the provided assignment operation.
//
// Syntax:
//    %foo = ...
func getResult(inst ir.Instruction) (result ast.Expr, err error) {
	v, ok := inst.(value.Named)
	if !ok {
		return nil, errors.Errorf("unable to get result identifier from instruction; expected value.Named, got %T", inst)
	}
	if types.Equal(v.Type(), types.Void) {
		return nil, errors.Errorf("unable to get result from void type instruction")
	}
	return newIdent(v.GetName()), nil
}

// newIdent returns a new identifier based on the given string after replacing
// any illegal characters with underscore and dropping any numeric suffixes
// (e.g. "i.0" and "i.1" => "i").
func newIdent(s string) *ast.Ident {
	// Drop numeric suffix.
	//if pos := strings.Index(s, "."); pos != -1 {
	//	s = s[:pos]
	//}

	// Prefix with "_" for local variables.
	s = "_" + s

	f := func(r rune) rune {
		switch {
		case unicode.IsLetter(r), unicode.IsNumber(r):
			// valid rune in identifier.
			return r
		}
		return '_'
	}
	return ast.NewIdent(strings.Map(f, s))
}
