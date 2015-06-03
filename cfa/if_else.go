package cfa

import (
	"fmt"

	"decomp.org/x/graphs/primitive"
	"github.com/mewfork/dot"
)

// IfElse represents a 2-way conditional statement.
//
// Pseudo-code:
//
//    if (A) {
//       B
//    } else {
//       C
//    }
//    D
type IfElse struct {
	// Condition node (A).
	Cond *dot.Node
	// Body node of the true branch (B).
	BodyTrue *dot.Node
	// Body node of the false branch (C).
	BodyFalse *dot.Node
	// Exit node (D).
	Exit *dot.Node
}

// Prim returns a representation of the high-level control flow primitive, as a
// mapping from control flow primitive node names to control flow graph node
// names.
//
// Example mapping:
//
//    "cond":       "A"
//    "body_true":  "B"
//    "body_false": "C"
//    "exit":       "D"
func (prim IfElse) Prim() *primitive.Primitive {
	cond, bodyTrue, bodyFalse, exit := prim.Cond.Name, prim.BodyTrue.Name, prim.BodyFalse.Name, prim.Exit.Name
	return &primitive.Primitive{
		Prim: "if_else",
		Node: "if_else_" + cond,
		Nodes: map[string]string{
			"cond":       cond,
			"body_true":  bodyTrue,
			"body_false": bodyFalse,
			"exit":       exit,
		},
		Entry: cond,
		Exit:  exit,
	}
}

// String returns a string representation of prim in DOT format.
//
// Example output:
//
//    digraph if_else {
//       cond -> body_true
//       cond -> body_false
//       body_true -> exit
//       body_false -> exit
//    }
func (prim IfElse) String() string {
	cond, bodyTrue, bodyFalse, exit := prim.Cond, prim.BodyTrue, prim.BodyFalse, prim.Exit
	const format = `
digraph if_else {
	%s -> %s
	%s -> %s
	%s -> %s
	%s -> %s
}`
	return fmt.Sprintf(format[1:], cond, bodyTrue, cond, bodyFalse, bodyTrue, exit, bodyFalse, exit)
}

// FindIfElse returns the first occurrence of a 2-way conditional statement in
// g, and a boolean indicating if such a primitive was found.
func FindIfElse(g *dot.Graph) (prim IfElse, ok bool) {
	// Range through cond node candidates.
	for _, cond := range g.Nodes.Nodes {
		// Verify that cond has two successors (body_true and body_false).
		if len(cond.Succs) != 2 {
			continue
		}
		prim.Cond = cond

		// Select body_true and body_false node candidates.
		prim.BodyTrue, prim.BodyFalse = cond.Succs[0], cond.Succs[1]

		// Verify that body_true has one successor (exit).
		if len(prim.BodyTrue.Succs) != 1 {
			continue
		}

		// Select exit node candidate.
		prim.Exit = prim.BodyTrue.Succs[0]
		if prim.IsValid(g) {
			return prim, true
		}
	}
	return IfElse{}, false
}

// IsValid reports whether the cond, body_true, body_false and exit node
// candidates of prim form a valid 2-way conditional statement in g.
//
// Control flow graph:
//
//              cond
//             ↙    ↘
//    body_true      body_false
//             ↘    ↙
//              exit
func (prim IfElse) IsValid(g *dot.Graph) bool {
	cond, bodyTrue, bodyFalse, exit := prim.Cond, prim.BodyTrue, prim.BodyFalse, prim.Exit

	// Verify that cond has two successors (body_true and body_false).
	if len(cond.Succs) != 2 || !cond.HasSucc(bodyTrue) || !cond.HasSucc(bodyFalse) {
		return false
	}

	// Verify that body_true has one predecessor (cond) and one successor (exit).
	if len(bodyTrue.Preds) != 1 || len(bodyTrue.Succs) != 1 || !bodyTrue.HasSucc(exit) {
		return false
	}

	// Verify that body_false has one predecessor (cond) and one successor (exit).
	if len(bodyFalse.Preds) != 1 || len(bodyFalse.Succs) != 1 || !bodyFalse.HasSucc(exit) {
		return false
	}

	// Verify that exit has two predecessor (body_true and body_false).
	return len(exit.Preds) == 2
}
