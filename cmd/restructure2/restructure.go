package main

import (
	"decomp.org/x/graphs/primitive"
	"github.com/mewfork/dot"
	"github.com/mewkiz/pkg/errutil"
)

// restructure attempts to recover the control flow primitives of a given
// control flow graph. It does so by repeatedly locating and merging structured
// subgraphs (graph representations of control flow primitives) into single
// nodes until the entire graph is reduced into a single node or no structured
// subgraphs may be located. The steps argument specifies whether to record the
// intermediate CFGs at each step. The returned list of primitives is ordered in
// the same sequence as they were located.
func restructure(g *dot.Graph, steps bool) (prims []*primitive.Primitive, err error) {
	// Locate control flow primitives.
	for step := 1; len(g.Nodes.Nodes) > 1; step++ {
		prim, err := findPrim(g, steps, step)
		if err != nil {
			return nil, errutil.Err(err)
		}
		prims = append(prims, prim)
	}
	return prims, nil
}

// findPrim locates a control flow primitive in the provided control flow graph
// and merges its nodes into a single node. The steps argument specifies whether
// to record the pre- and post-merge CFGs of the given step.
func findPrim(g *dot.Graph, steps bool, step int) (*primitive.Primitive, error) {
	panic("not yet implemented")
}
