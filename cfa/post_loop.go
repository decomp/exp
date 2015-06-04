package cfa

import (
	"fmt"

	"decomp.org/x/graphs/primitive"
	"github.com/mewfork/dot"
)

// PostLoop represents a post-test loop.
//
// Pseudo-code:
//
//    do {
//    } while (A)
//    B
type PostLoop struct {
	// Condition node (A).
	Cond *dot.Node
	// Exit node (B).
	Exit *dot.Node
}

// Prim returns a representation of the high-level control flow primitive, as a
// mapping from control flow primitive node names to control flow graph node
// names.
//
// Example mapping:
//
//    "cond": "A"
//    "exit": "B"
func (prim PostLoop) Prim() *primitive.Primitive {
	cond, exit := prim.Cond.Name, prim.Exit.Name
	return &primitive.Primitive{
		Prim: "post_loop",
		Node: "post_loop_" + cond,
		Nodes: map[string]string{
			"cond": cond,
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
//    digraph post_loop {
//       cond -> cond
//       cond -> exit
//    }
func (prim PostLoop) String() string {
	cond, exit := prim.Cond, prim.Exit
	const format = `
digraph post_loop {
	%s -> %s
	%s -> %s
}`
	return fmt.Sprintf(format[1:], cond, cond, cond, exit)
}

// FindPostLoop returns the first occurrence of a post-test loop in g, and a
// boolean indicating if such a primitive was found.
func FindPostLoop(g *dot.Graph) (prim PostLoop, ok bool) {
	// Range through cond node candidates.
	for _, cond := range g.Nodes.Nodes {
		// Verify that cond has two successors (cond and exit).
		if len(cond.Succs) != 2 {
			continue
		}
		prim.Cond = cond

		// Try the first exit node candidate.
		prim.Exit = cond.Succs[0]
		if prim.IsValid(g) {
			return prim, true
		}

		// Try the second exit node candidate.
		prim.Exit = cond.Succs[1]
		if prim.IsValid(g) {
			return prim, true
		}
	}
	return PostLoop{}, false
}

// IsValid reports whether the cond and exit node candidates of prim form a
// valid post-test loop in g.
//
// Control flow graph:
//
//    cond ↘
//    ↓   ↖↲
//    ↓
//    exit
func (prim PostLoop) IsValid(g *dot.Graph) bool {
	cond, exit := prim.Cond, prim.Exit

	// Verify that cond has two successors (cond and exit).
	if len(cond.Succs) != 2 || !cond.HasSucc(cond) || !cond.HasSucc(exit) {
		return false
	}

	// Verify that exit has one predecessor (cond).
	return len(exit.Preds) == 1
}
