// restructure is a tool which recovers high-level control flow primitives from
// control flow graphs (e.g. *.dot -> *.json). It takes an unstructured CFG (in
// Graphviz DOT format) as input and produces a structured CFG (in JSON format),
// which describes how the high-level control flow primitives relate to the
// nodes of the CFG.
//
// Usage:
//     restructure [OPTION]... [CFG.dot]
//
//     Flags:
//       -indent
//             Indent JSON output.
//       -o string
//             Output path.
//       -q    Suppress non-error messages.
//       -steps
//             Output intermediate CFGs at each step.
//       -v    Verbose output.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/decomp/decomp/graphs/primitive"
	"github.com/decomp/exp/cfa"
	"github.com/mewkiz/pkg/errutil"
	"github.com/mewspring/dot"
)

const use = `
Usage: restructure [OPTION]... [CFG.dot]
Recover control flow primitives from control flow graphs (e.g. *.dot -> *.json).
`

// usage prints the usage information of restructure to stderr.
func usage() {
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

// Global command line flags.
var (
	// When flagQuiet is true, suppress non-error messages.
	flagQuiet bool
	// When flagVerbose is true, enable verbose output.
	flagVerbose bool
)

func main() {
	// Parse command line arguments.
	var (
		// Specifies whether to indent JSON output.
		indent bool
		// Output path.
		output string
		// Specifies whether to output intermediate CFGs at each step.
		steps bool
	)
	flag.BoolVar(&indent, "indent", false, "Indent JSON output.")
	flag.StringVar(&output, "o", "", "Output path.")
	flag.BoolVar(&flagQuiet, "q", false, "Suppress non-error messages.")
	flag.BoolVar(&steps, "steps", false, "Output intermediate CFGs at each step.")
	flag.BoolVar(&flagVerbose, "v", false, "Verbose output.")
	flag.Usage = usage
	flag.Parse()

	// Parse input graph.
	var dotPath string
	switch flag.NArg() {
	case 0:
		// Read from stdin.
		dotPath = "-"
	case 1:
		// Read from FILE.
		dotPath = flag.Arg(0)
	default:
		flag.Usage()
		os.Exit(1)
	}
	g, err := parseGraph(dotPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create a structured CFG from the unstructured CFG.
	prims, err := restructure(g, steps)
	if err != nil {
		log.Fatal(err)
	}

	// Print the JSON output to stdout or the path specified by the "-o" flag.
	w := os.Stdout
	if len(output) > 0 {
		f, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	}
	if err := writeJSON(w, prims, indent); err != nil {
		log.Fatal(err)
	}
}

// parseGraph parses the provided DOT graph.
func parseGraph(dotPath string) (g *dot.Graph, err error) {
	// Parse DOT graph.
	if dotPath == "-" {
		// Read from stdin.
		buf, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, errutil.Err(err)
		}
		g, err = dot.Read(buf)
		if err != nil {
			return nil, errutil.Err(err)
		}
	} else {
		// Read from FILE.
		g, err = dot.ParseFile(dotPath)
		if err != nil {
			return nil, errutil.Err(err)
		}
	}

	// Validate the parsed graph.
	if len(g.Nodes.Nodes) == 0 {
		return nil, errutil.Newf("unable to restructure empty graph %q", dotPath)
	}

	return g, nil
}

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
		// Locate primitive.
		prim, err := findPrim(g)
		if err != nil {
			return nil, errutil.Err(err)
		}
		prims = append(prims, prim)

		// Pretty-print located primitive.
		if flagVerbose && !flagQuiet {
			printPrim(prim)
		}

		// Output pre-merge intermediate CFG.
		if steps {
			path := fmt.Sprintf("%s_%da.dot", g.Name, step)
			var highlight []string
			for _, node := range prim.Nodes {
				highlight = append(highlight, node)
			}
			if err := dumpStep(g, path, highlight); err != nil {
				return nil, errutil.Err(err)
			}
		}

		// Merge the nodes of the primitive into a single node.
		if err := merge(g, prim); err != nil {
			return nil, errutil.Err(err)
		}

		// Output post-merge intermediate CFG.
		if steps {
			path := fmt.Sprintf("%s_%db.dot", g.Name, step)
			highlight := []string{prim.Node}
			if err := dumpStep(g, path, highlight); err != nil {
				return nil, errutil.Err(err)
			}
		}
	}
	return prims, nil
}

// findPrim locates a control flow primitive in the provided control flow graph
// and merges its nodes into a single node.
func findPrim(g *dot.Graph) (*primitive.Primitive, error) {
	// Locate pre-test loops.
	if prim, ok := cfa.FindPreLoop(g); ok {
		return prim.Prim(), nil
	}

	// Locate post-test loops.
	if prim, ok := cfa.FindPostLoop(g); ok {
		return prim.Prim(), nil
	}

	// Locate 1-way conditionals.
	if prim, ok := cfa.FindIf(g); ok {
		return prim.Prim(), nil
	}

	// Locate 1-way conditionals with a body return statements.
	if prim, ok := cfa.FindIfReturn(g); ok {
		return prim.Prim(), nil
	}

	// Locate 2-way conditionals.
	if prim, ok := cfa.FindIfElse(g); ok {
		return prim.Prim(), nil
	}

	// Locate sequences of two statements.
	if prim, ok := cfa.FindSeq(g); ok {
		return prim.Prim(), nil
	}

	return nil, errutil.New("unable to locate control flow primitive")
}

// printPrim pretty-prints the given primitive to stderr.
func printPrim(prim *primitive.Primitive) {
	// Sort primitive nodes.
	var keys []string
	for key := range prim.Nodes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Print mapping between primitive nodes and graph nodes.
	fmt.Fprintf(os.Stderr, "Located %q primitive at node %q:\n", prim.Prim, prim.Entry)
	w := tabwriter.NewWriter(os.Stderr, 0, 1, 1, ' ', 0)
	for _, primNode := range keys {
		fmt.Fprintf(w, "  %q:\t%q\n", primNode, prim.Nodes[primNode])
	}
	w.Flush()
}

// dumpStep stores a DOT representation of g to path with the specified nodes
// highlighted in red.
func dumpStep(g *dot.Graph, path string, highlight []string) error {
	// Highlight the specified nodes in red.
	for _, nodeName := range highlight {
		node, ok := g.Nodes.Lookup[nodeName]
		if !ok {
			return errutil.Newf("unable to locate node %q", nodeName)
		}
		if node.Attrs == nil {
			node.Attrs = dot.NewAttrs()
		}
		node.Attrs["fillcolor"] = "red"
		node.Attrs["style"] = "filled"
	}

	// Store DOT graph.
	if !flagQuiet {
		log.Printf("Creating: %q", path)
	}
	buf := []byte(g.String())
	if err := ioutil.WriteFile(path, buf, 0644); err != nil {
		return errutil.Err(err)
	}

	// Restore node colour.
	for _, nodeName := range highlight {
		node, ok := g.Nodes.Lookup[nodeName]
		if !ok {
			return errutil.Newf("unable to locate node %q", nodeName)
		}

		delete(node.Attrs, "fillcolor")
		delete(node.Attrs, "style")
	}

	return nil
}

// merge merges the nodes of the primitive into a single node.
func merge(g *dot.Graph, prim *primitive.Primitive) error {
	var nodes []*dot.Node
	for _, nodeName := range prim.Nodes {
		node, ok := g.Nodes.Lookup[nodeName]
		if !ok {
			return errutil.Newf("unable to locate pre-merge node %q", nodeName)
		}
		nodes = append(nodes, node)
	}

	// Locate entry node.
	entry, ok := g.Nodes.Lookup[prim.Entry]
	if !ok {
		return errutil.Newf("unable to locate entry node %q", prim.Entry)
	}

	// TODO: Implement support for single-entry/no-exit primitives.

	// Locate exit node.
	exit, ok := g.Nodes.Lookup[prim.Exit]
	if !ok {
		return errutil.Newf("unable to locate exit node %q", prim.Exit)
	}

	// Merge nodes.
	if err := g.Replace(nodes, prim.Node, entry, exit); err != nil {
		return errutil.Err(err)
	}

	return nil
}

// writeJSON writes the primitives in JSON format to w.
func writeJSON(w io.Writer, prims []*primitive.Primitive, indent bool) error {
	// Output indented JSON.
	if indent {
		buf, err := json.MarshalIndent(prims, "", "\t")
		if err != nil {
			return errutil.Err(err)
		}
		if _, err = io.Copy(w, bytes.NewReader(buf)); err != nil {
			return errutil.Err(err)
		}
		return nil
	}

	// Output JSON.
	enc := json.NewEncoder(w)
	if err := enc.Encode(prims); err != nil {
		return errutil.Err(err)
	}
	return nil
}
