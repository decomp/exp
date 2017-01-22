// ll2dot is a tool which generates control flow graphs from LLVM IR assembly
// files (e.g. *.ll -> *.dot). The output is a set of Graphviz DOT files, each
// representing the control flow graph of a function using one node per basic
// block.
//
// For a source file "foo.ll" containing the functions "bar" and "baz" the
// following DOT files will be generated:
//
//    * foo_graphs/bar.dot
//    * foo_graphs/baz.dot
//
// Usage
//
//    ll2dot [OPTION]... FILE...
//
// If FILE is -, read standard input.
//
// Flags
//
//    -f    force overwrite existing graph directories
//    -funcs string
//          comma separated list of functions to parse (e.g. "foo,bar")
//    -img
//          generate an image representation of the CFG
//    -q    suppress non-error messages
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/graphism/dot/ast"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/mewkiz/pkg/pathutil"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// dbg represents a logger with the "ll2dot:" prefix, which logs debug messages
// to standard error.
var dbg = log.New(os.Stderr, term.Blue("ll2dot:")+" ", 0)

// usage prints a usage message to standard error.
func usage() {
	const use = `
Usage: ll2dot [OPTION]... FILE...
Generate control flow graphs from LLVM IR assembly files (e.g. *.ll -> *.dot).

If FILE is -, read standard input.

Flags:`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line flags.
	var (
		// force specifies whether to force overwrite existing graph directories.
		force bool
		// funcs specifies a comma separated list of functions to parse (e.g.
		// "foo,bar").
		funcs string
		// genimg specifies whether to generate an image representation of the
		// CFG.
		genimg bool
		// quiet specifies whether to suppress non-error messages.
		quiet bool
	)
	flag.BoolVar(&force, "f", false, "force overwrite existing graph directories")
	flag.StringVar(&funcs, "funcs", "", `comma separated list of functions to parse (e.g. "foo,bar")`)
	flag.BoolVar(&genimg, "img", false, "generate an image representation of the CFG")
	flag.BoolVar(&quiet, "q", false, "suppress non-error messages")
	flag.Usage = usage
	flag.Parse()

	// Mute debug messages if `-q` is set.
	if quiet {
		dbg = log.New(ioutil.Discard, "", 0)
	}

	// Get function names from the comma-separated `-funcs` list.
	funcNames := make(map[string]bool)
	for _, funcName := range strings.Split(funcs, ",") {
		if len(funcName) == 0 {
			continue
		}
		funcNames[funcName] = true
	}

	for _, llPath := range flag.Args() {
		if err := ll2dot(llPath, funcNames, force, genimg); err != nil {
			log.Fatal(err)
		}
	}
}

// ll2dot parses the provided LLVM IR assembly file and generates a control flow
// graph for each of its defined functions using one node per basic block.
func ll2dot(llPath string, funcNames map[string]bool, force, genimg bool) error {
	// Parse LLVM IR assembly file.
	module, err := asm.ParseFile(llPath)
	if err != nil {
		return errors.WithStack(err)
	}

	// Get functions set by `-funcs` or all functions if `-funcs` not used.
	var funcs []*ir.Function
	for _, f := range module.Funcs {
		if len(funcNames) == 0 || funcNames[f.Name] {
			funcs = append(funcs, f)
		}
	}

	// Generate a control flow graph for each function.
	dotDir, err := createDotDir(llPath, force)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, f := range funcs {
		// Skip function declarations.
		if len(f.Blocks) == 0 {
			continue
		}

		// Generate control flow graph.
		dbg.Printf("Parsing function %q.", f.Name)
		graph, err := createCFG(f)
		if err != nil {
			return errors.WithStack(err)
		}

		// Store DOT graph.
		if err := dumpCFG(dotDir, f.Name, graph, genimg); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// dumpCFG stores the given control flow graph as a DOT file. If `-img` is set,
// it also stores an image representation of the CFG.
//
// For a source file "foo.ll" containing the functions "bar" and "baz" the
// following DOT files will be created:
//
//    foo_graphs/bar.dot
//    foo_graphs/baz.dot
func dumpCFG(dotDir, funcName string, graph *ast.Graph, genimg bool) error {
	dotName := funcName + ".dot"
	dotPath := filepath.Join(dotDir, dotName)
	dbg.Printf("Creating: %q.", dotPath)
	buf := []byte(graph.String())
	if err := ioutil.WriteFile(dotPath, buf, 0644); err != nil {
		return errors.WithStack(err)
	}

	// Store an image representation of the CFG if `-img` is set.
	if genimg {
		pngName := funcName + ".png"
		pngPath := filepath.Join(dotDir, pngName)
		dbg.Printf("Creating: %q.", pngPath)
		cmd := exec.Command("dot", "-Tpng", "-o", pngPath, dotPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// createCFG generates a control flow graph for the given function using one
// node per basic block.
func createCFG(f *ir.Function) (*ast.Graph, error) {
	// Create a new directed graph.
	funcName := f.Name
	graph := &ast.Graph{
		ID:       funcName,
		Directed: true,
	}

	// Populate the graph with one node per basic block.
	for i, block := range f.Blocks {
		// Add a node for the given basic block to the graph.
		blockName := block.Name
		var attrs []*ast.Attr
		if i == 0 {
			attrs = []*ast.Attr{
				{Key: "label", Val: "entry"},
			}
		}
		node := &ast.Node{
			ID: blockName,
		}
		nodeStmt := &ast.NodeStmt{
			Node:  node,
			Attrs: attrs,
		}
		graph.Stmts = append(graph.Stmts, nodeStmt)

		// Add edges from the node to the target basic blocks.
		term := block.Term
		switch term := term.(type) {
		case *ir.TermRet:
			// Return instruction.
			//    ret
			//    ret Type Value
			//
			// Exit node with no target basic blocks.
		case *ir.TermUnreachable:
			// Unreachable instruction.
			//    unreachable
			//
			// Exit node with no target basic blocks.
		case *ir.TermBr:
			// Unconditional branch instruction.
			//    br label TargetBranch
			//
			// Add target branch.
			from := &ast.Node{ID: blockName}
			to := &ast.Edge{
				Directed: true,
				Vertex:   &ast.Node{ID: term.Target.Name},
			}
			edge := &ast.EdgeStmt{
				From: from,
				To:   to,
			}
			graph.Stmts = append(graph.Stmts, edge)
		case *ir.TermCondBr:
			// Conditional branching instruction.
			//    br i1 Cond, label TrueBranch, label FalseBranch
			//
			// Add true target branch.
			from := &ast.Node{ID: blockName}
			to := &ast.Edge{
				Directed: true,
				Vertex:   &ast.Node{ID: term.TargetTrue.Name},
			}
			edge := &ast.EdgeStmt{
				From: from,
				To:   to,
				Attrs: []*ast.Attr{
					{Key: "label", Val: "true"},
				},
			}
			graph.Stmts = append(graph.Stmts, edge)
			// Add false target branch.
			from = &ast.Node{ID: blockName}
			to = &ast.Edge{
				Directed: true,
				Vertex:   &ast.Node{ID: term.TargetFalse.Name},
			}
			edge = &ast.EdgeStmt{
				From: from,
				To:   to,
				Attrs: []*ast.Attr{
					{Key: "label", Val: "false"},
				},
			}
			graph.Stmts = append(graph.Stmts, edge)
		default:
			panic(fmt.Sprintf("support for terminator %T not yet implemented", term))
		}
	}
	return graph, nil
}

// createDotDir creates and returns an output directory based on the path of the
// given LLVM IR file.
//
// For a source file "/foo/bar.ll" the output directory "/foo/bar_graphs/" is
// created. If the `-force` flag is set, existing graph directories are
// overwritten by force.
func createDotDir(llPath string, force bool) (string, error) {
	dotDir := pathutil.TrimExt(llPath) + "_graphs"

	// Force overwrite existing graph directories.
	if force {
		if err := os.RemoveAll(dotDir); err != nil {
			return "", errors.WithStack(err)
		}
	}

	if err := os.Mkdir(dotDir, 0755); err != nil {
		return "", errors.WithStack(err)
	}
	return dotDir, nil
}
