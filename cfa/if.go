package cfa

import (
	"fmt"

	"github.com/decomp/decomp/graphs/primitive"
	"github.com/mewspring/dot"
)

// If represents a 1-way conditional statement.
//
// Pseudo-code:
//
//    if (A) {
//       B
//    }
//    C
type If struct {
	// Condition node (A).
	Cond *dot.Node
	// Body node (B).
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
func (prim If) Prim() *primitive.Primitive {
	cond, body, exit := prim.Cond.Name, prim.Body.Name, prim.Exit.Name
	return &primitive.Primitive{
		Prim: "if",
		Node: "if_" + cond,
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
//    digraph if {
//       cond -> body
//       cond -> exit
//       body -> exit
//    }
func (prim If) String() string {
	cond, body, exit := prim.Cond, prim.Body, prim.Exit
	const format = `
digraph if {
	%s -> %s
	%s -> %s
	%s -> %s
}`
	return fmt.Sprintf(format[1:], cond, body, cond, exit, body, exit)
}

// FindIf returns the first occurrence of a 1-way conditional statement in g,
// and a boolean indicating if such a primitive was found.
func FindIf(g *dot.Graph) (prim If, ok bool) {
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
	return If{}, false
}

// IsValid reports whether the cond, body and exit node candidates of prim form
// a valid 1-way conditional statement in g.
//
// Control flow graph:
//
//    cond
//    ↓   ↘
//    ↓    body
//    ↓   ↙
//    exit
func (prim If) IsValid(g *dot.Graph) bool {
	cond, body, exit := prim.Cond, prim.Body, prim.Exit

	// Dominator sanity check.
	if !cond.Dominates(body) || !cond.Dominates(exit) {
		return false
	}

	// Verify that cond has two successors (body and exit).
	if len(cond.Succs) != 2 || !cond.HasSucc(body) || !cond.HasSucc(exit) {
		return false
	}

	// Verify that body has one predecessor (cond) and one successor (exit).
	if len(body.Preds) != 1 || len(body.Succs) != 1 || !body.HasSucc(exit) {
		return false
	}

	// Verify that exit has two predecessors (cond and body).
	return len(exit.Preds) == 2
}
