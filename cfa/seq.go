package cfa

import (
	"fmt"

	"decomp.org/decomp/graphs/primitive"
	"github.com/mewspring/dot"
)

// Seq represents a sequence of two statements.
//
// Pseudo-code:
//
//    A
//    B
type Seq struct {
	// Entry node (A).
	Entry *dot.Node
	// Exit node (B).
	Exit *dot.Node
}

// Prim returns a representation of the high-level control flow primitive, as a
// mapping from control flow primitive node names to control flow graph node
// names.
//
// Example mapping:
//
//    "entry": "A"
//    "exit":  "B"
func (prim Seq) Prim() *primitive.Primitive {
	entry, exit := prim.Entry.Name, prim.Exit.Name
	return &primitive.Primitive{
		Prim: "seq",
		Node: "seq_" + entry,
		Nodes: map[string]string{
			"entry": entry,
			"exit":  exit,
		},
		Entry: entry,
		Exit:  exit,
	}
}

// String returns a string representation of prim in DOT format.
//
// Example output:
//
//    digraph seq {
//       entry -> exit
//    }
func (prim Seq) String() string {
	entry, exit := prim.Entry, prim.Exit
	const format = `
digraph seq {
	%s -> %s
}`
	return fmt.Sprintf(format[1:], entry, exit)
}

// FindSeq returns the first occurrence of a sequence of two statements in g,
// and a boolean indicating if such a primitive was found.
func FindSeq(g *dot.Graph) (prim Seq, ok bool) {
	// Range through entry node candidates.
	for _, entry := range g.Nodes.Nodes {
		// Verify that entry has one successor (exit).
		if len(entry.Succs) != 1 {
			continue
		}
		prim.Entry = entry

		// Select exit node candidate.
		prim.Exit = entry.Succs[0]
		if prim.IsValid(g) {
			return prim, true
		}
	}
	return Seq{}, false
}

// IsValid reports whether the entry and exit node candidates of prim form a
// valid sequence of two statements in g.
//
// Control flow graph:
//
//    entry
//    ↓
//    exit
func (prim Seq) IsValid(g *dot.Graph) bool {
	entry, exit := prim.Entry, prim.Exit

	// Dominator sanity check.
	if !entry.Dominates(exit) {
		return false
	}

	// Verify that entry has one successor (exit).
	if len(entry.Succs) != 1 || !entry.HasSucc(exit) {
		return false
	}

	// Verify that exit has one predecessor (entry).
	return len(exit.Preds) == 1
}
