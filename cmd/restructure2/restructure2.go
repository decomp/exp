// restructure is a tool which recovers high-level control flow primitives from
// control flow graphs (e.g. *.dot -> *.json). It takes an unstructured CFG (in
// the Graphviz DOT file format) as input and produces a structured CFG (in JSON
// format), which describes how the high-level control flow primitives relate to
// the nodes of the CFG.
//
// Usage:
//     restructure [OPTION]... [CFG.dot]
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

	"decomp.org/x/graphs/primitive"
	"github.com/mewfork/dot"
	"github.com/mewkiz/pkg/errutil"
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

func main() {
	// Parse command line arguments.
	var (
		// Specifies whether to indent JSON output.
		indent bool
		// Output path.
		output string
	)
	flag.BoolVar(&indent, "indent", false, "Indent JSON output.")
	flag.StringVar(&output, "o", "", "Output path.")
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
	prims, err := restructure(g)
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
// subgraphs may be located. The list of primitives is ordered in the same
// sequence as they were located.
func restructure(g *dot.Graph) ([]*primitive.Primitive, error) {
	panic("not yet implemented")
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
