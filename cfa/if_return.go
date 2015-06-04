package cfa

import (
	"fmt"
	"log"

	"decomp.org/x/graphs/primitive"
	"github.com/mewfork/dot"
)

// IfReturn represents a 1-way conditional with a body return statement.
//
// Pseudo-code:
//
//    if (A) {
//       B
//       return
//    }
//    C
type IfReturn struct {
	// Condition node (A).
	Cond *dot.Node
	// Body node with return statement (B).
	Body *dot.Node
	// Exit node (C).
	Exit *dot.Node
}

// Prim returns a representation of the high-level control flow primitive, as a
// mapping from control flow primitive node names to control flow graph node
// names.
//
// Example mapping:
//
//    "cond": "A"
//    "body": "B"
//    "exit": "C"
func (prim IfReturn) Prim() *primitive.Primitive {
	cond, body, exit := prim.Cond.Name, prim.Body.Name, prim.Exit.Name
	return &primitive.Primitive{
		Prim: "if_return",
		Node: "if_return_" + cond,
		Nodes: map[string]string{
			"cond": cond,
			"body": body,
			"exit": exit,
		},
		Entry: cond,
		Exit:  exit,
	}
}

// String returns a string representation of prim in DOT format.
//
// Example output:
//
//    digraph if_return {
//       cond -> body
//       cond -> exit
//    }
func (prim IfReturn) String() string {
	cond, body, exit := prim.Cond, prim.Body, prim.Exit
	const format = `
digraph if_return {
	%s -> %s
	%s -> %s
}`
	return fmt.Sprintf(format[1:], cond, body, cond, exit)
}

// FindIfReturn returns the first occurrence of a 1-way conditional with a body
// return statement in g, and a boolean indicating if such a primitive was
// found.
func FindIfReturn(g *dot.Graph) (prim IfReturn, ok bool) {
	// Range through cond node candidates.
	for _, cond := range g.Nodes.Nodes {
		// Verify that cond has two successors (body and exit).
		if len(cond.Succs) != 2 {
			continue
		}
		prim.Cond = cond

		// Select body and exit node candidates.
		prim.Body, prim.Exit = cond.Succs[0], cond.Succs[1]
		if prim.IsValid(g) {
			return prim, true
		}

		// Swap body and exit node candidates and try again.
		prim.Body, prim.Exit = prim.Exit, prim.Body
		if prim.IsValid(g) {
			return prim, true
		}
	}
	return IfReturn{}, false
}

// IsValid reports whether the cond, body and exit node candidates of prim form
// a valid 1-way conditional with a body return statement in g.
//
// Control flow graph:
//
//    cond
//    ↓   ↘
//    ↓    body
//    ↓
//    exit
func (prim IfReturn) IsValid(g *dot.Graph) bool {
	cond, body, exit := prim.Cond, prim.Body, prim.Exit

	// Dominator sanity check.
	if !cond.Dominates(body) {
		// TODO: Remove debug output.
		log.Printf("IfReturn: cond %q does not dominate body %q", cond, body)
		return false
	}
	if !cond.Dominates(exit) {
		// TODO: Remove debug output.
		log.Printf("IfReturn: cond %q does not dominate exit %q", cond, exit)
		return false
	}

	// Verify that cond has two successors (body and exit).
	if len(cond.Succs) != 2 || !cond.HasSucc(body) || !cond.HasSucc(exit) {
		return false
	}

	// Verify that body has one predecessor (cond) and zero successors.
	if len(body.Preds) != 1 || len(body.Succs) != 0 {
		return false
	}

	// Verify that exit has one predecessor (cond).
	if len(exit.Preds) != 1 {
		return false
	}

	// Verify that the entry node (cond) has no predecessors dominated by cond,
	// as that would indicate a loop construct.
	//
	//       cond
	//     ↗ ↓   ↘
	//    ↑  ↓    body
	//    ↑  ↓
	//    ↑  exit
	//     ↖ ↓
	//       A
	for _, pred := range cond.Preds {
		if cond.Dominates(pred) {
			return false
		}
	}

	return true
}
