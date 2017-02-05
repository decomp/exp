package main

import (
	"go/ast"

	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// BasicBlock represents a conceptual basic block. If one statement of the basic
// block is executed all statements of the basic block are executed until the
// terminating instruction is reached which transfers control to another basic
// block.
type BasicBlock interface {
	// Name returns the name of the basic block.
	Name() string
	// Stmts returns the statements of the basic block.
	Stmts() []ast.Stmt
	// SetStmts sets the statements of the basic block.
	SetStmts(stmts []ast.Stmt)
	// Term returns the terminator instruction of the basic block.
	Term() ir.Terminator
}

// basicBlock represents a basic block in which the instructions have been
// translated to Go AST statement nodes but the terminator instruction is an
// unmodified LLVM IR value.
type basicBlock struct {
	// Basic block name.
	name string
	// Basic block instructions.
	stmts []ast.Stmt
	// A map from variable name to variable definitions which represents the PHI
	// instructions of the basic block.
	phis map[string][]*definition
	// Terminator instruction.
	term ir.Terminator
}

// Name returns the name of the basic block.
func (bb *basicBlock) Name() string { return bb.name }

// Stmts returns the statements of the basic block.
func (bb *basicBlock) Stmts() []ast.Stmt { return bb.stmts }

// SetStmts sets the statements of the basic block.
func (bb *basicBlock) SetStmts(stmts []ast.Stmt) { bb.stmts = stmts }

// Term returns the terminator instruction of the basic block.
func (bb *basicBlock) Term() ir.Terminator { return bb.term }

// parseBasicBlock converts the provided LLVM IR basic block into a basic block
// in which the instructions have been translated to Go AST statement nodes but
// the terminator instruction is an unmodified LLVM IR value.
func parseBasicBlock(llBB *ir.BasicBlock) (bb *basicBlock, err error) {
	name := llBB.Name
	bb = &basicBlock{name: name, phis: make(map[string][]*definition)}
	for _, inst := range llBB.Insts {
		switch inst := inst.(type) {
		case *ir.InstPhi:
			// Handle PHI instructions.
			ident, def, err := parsePHIInst(inst)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			bb.phis[ident] = def
		default:
			// Handle non-terminator instructions.
			stmt, err := parseInst(inst)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			bb.stmts = append(bb.stmts, stmt)
		}
	}

	// Handle terminator instruction.
	if err := bb.addTerm(llBB.Term); err != nil {
		return nil, errors.WithStack(err)
	}
	return bb, nil
}

// addTerm adds the provided terminator instruction to the basic block. If the
// terminator instruction doesn't have a target basic block (e.g. ret) it is
// parsed and added to the statements list of the basic block instead.
func (bb *basicBlock) addTerm(term ir.Terminator) error {
	// TODO: Check why there is no opcode in the llvm library for the resume
	// terminator instruction.
	switch term := term.(type) {
	case *ir.TermRet:
		// The return instruction doesn't have any target basic blocks so treat it
		// like a regular instruction and append it to the list of statements.
		ret, err := parseRetInst(term)
		if err != nil {
			return errors.WithStack(err)
		}
		bb.stmts = append(bb.stmts, ret)
	case *ir.TermBr:
		// Parse the terminator instruction during the control flow analysis.
		bb.term = term
	case *ir.TermCondBr:
		// Parse the terminator instruction during the control flow analysis.
		bb.term = term
	case *ir.TermSwitch:
		// Parse the terminator instruction during the control flow analysis.
		bb.term = term
	case *ir.TermUnreachable:
		// Parse the terminator instruction during the control flow analysis.
		bb.term = term
	default:
		panic(errors.Errorf("support for terminator %T not yet implemented", term))
	}
	return nil
}
