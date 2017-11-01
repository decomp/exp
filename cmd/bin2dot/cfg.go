package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/decomp/exp/bin"
	"github.com/decomp/exp/disasm/x86"
	"github.com/graphism/simple"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
)

// dumpCFG dumps the control flow graph of the given function.
func dumpCFG(dis *x86.Disasm, f *x86.Func) (graph.Directed, error) {
	// Index functions, basic blocks and instructions.
	g := simple.NewDirectedGraph()
	nodes := make(map[bin.Address]*Node)
	for _, block := range f.Blocks {
		id := strconv.Quote(block.Addr.String())
		n := &Node{
			Node:  g.NewNode(),
			id:    id,
			Attrs: make(Attrs),
		}
		if block.Addr == f.Addr {
			n.Attrs["label"] = "entry"
		}
		nodes[block.Addr] = n
		g.AddNode(n)
	}
	for _, block := range f.Blocks {
		targets := dis.Targets(block.Term, f.Addr)
		fmt.Println("block.Addr:", block.Addr)
		fmt.Println("block.Term:", block.Term)
		fmt.Println("targets:", targets)
		fmt.Println()

		from, ok := nodes[block.Addr]
		if !ok {
			panic(errors.Errorf("unable to locate basic block at %v", block.Addr))
		}
		for _, target := range targets {
			to, ok := nodes[target]
			if !ok {
				return nil, errors.Errorf("unable to locate target basic block at %v from %v in function at %v", target, block.Addr, f.Addr)
			}
			e := simple.Edge{
				F: from,
				T: to,
			}
			g.SetEdge(e)
		}
	}
	return g, nil
}

// Node is a basic block of a function.
type Node struct {
	graph.Node
	// DOTID of node.
	id string
	// DOT attributes.
	Attrs
}

// DOTID returns the DOTID of the node.
func (n Node) DOTID() string {
	return n.id
}

// ### [ Helper functions ] ####################################################

// Attrs specifies a set of DOT attributes as key-value pairs.
type Attrs map[string]string

// --- [ encoding.Attributer ] -------------------------------------------------

// Attributes returns the DOT attributes of a node or edge.
func (a Attrs) Attributes() []encoding.Attribute {
	var keys []string
	for key := range a {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var attrs []encoding.Attribute
	for _, key := range keys {
		attr := encoding.Attribute{
			Key:   key,
			Value: a[key],
		}
		attrs = append(attrs, attr)
	}
	return attrs
}
