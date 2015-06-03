package cfa

import (
	"fmt"

	"decomp.org/x/graphs/primitive"
	"github.com/mewfork/dot"
)

// PreLoop represents a pre-test loop.
//
// Pseudo-code:
//
//    while (A) {
//       B
//    }
//    C
type PreLoop struct {
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
func (prim PreLoop) Prim() *primitive.Primitive {
	cond, body, exit := prim.Cond.Name, prim.Body.Name, prim.Exit.Name
	return &primitive.Primitive{
		Prim: "pre_loop",
		Node: "pre_loop_" + cond,
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
//    digraph pre_loop {
//       cond -> body
//       cond -> exit
//       body -> cond
//    }
func (prim PreLoop) String() string {
	cond, body, exit := prim.Cond, prim.Body, prim.Exit
	const format = `
digraph pre_loop {
	%s -> %s
	%s -> %s
	%s -> %s
}`
	return fmt.Sprintf(format[1:], cond, body, cond, exit, body, cond)
}

// FindPreLoop returns the first occurrence of a pre-test loop in g, and a
// boolean indicating if such a primitive was found.
func FindPreLoop(g *dot.Graph) (prim PreLoop, ok bool) {
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
	return PreLoop{}, false
}

// IsValid reports whether the cond, body and exit node candidates of prim form
// a valid pre-test loop in g.
//
// Control flow graph:
//
//    cond
//    ↓  ↖↘
//    ↓   body
//    ↓
//    exit
func (prim PreLoop) IsValid(g *dot.Graph) bool {
	cond, body, exit := prim.Cond, prim.Body, prim.Exit

	// Verify that cond has two successors (body and exit).
	if len(cond.Succs) != 2 || !cond.HasSucc(body) || !cond.HasSucc(exit) {
		return false
	}

	// Verify that body has one predecessor (cond) and one successor (cond).
	if len(body.Preds) != 1 || len(body.Succs) != 1 || !body.HasSucc(cond) {
		return false
	}

	// Verify that exit has one predecessor (cond).
	return len(exit.Preds) == 1
}