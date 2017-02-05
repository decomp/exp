// ll2go is a tool which decompiles LLVM IR assembly files to Go source code
// (e.g. *.ll -> *.go).
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	xprimitive "github.com/decomp/decomp/graphs/primitive"
	"github.com/kr/pretty"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/mewkiz/pkg/osutil"
	"github.com/mewkiz/pkg/pathutil"
	"github.com/mewspring/dot"
	"github.com/pkg/errors"
)

var (
	// When flagForce is true, force overwrite existing Go source code.
	flagForce bool
	// flagFuncs specifies a comma separated list of functions to decompile (e.g.
	// "foo,bar").
	flagFuncs string
	// flagPkgName specifies the package name if non-empty.
	flagPkgName string
	// When flagQuiet is true, suppress non-error messages.
	flagQuiet bool
	// When flagQuiet is true, enable verbose output.
	flagVerbose bool
)

func init() {
	flag.BoolVar(&flagForce, "f", false, "Force overwrite existing Go source code.")
	flag.StringVar(&flagFuncs, "funcs", "", `Comma separated list of functions to decompile (e.g. "foo,bar").`)
	flag.StringVar(&flagPkgName, "pkgname", "", "Package name.")
	flag.BoolVar(&flagQuiet, "q", false, "Suppress non-error messages.")
	flag.BoolVar(&flagVerbose, "v", false, "Enable verbose output.")
	flag.Usage = usage
}

const use = `
Usage: ll2go [OPTION]... FILE...
Decompile LLVM IR assembly files to Go source code (e.g. *.ll -> *.go).

Flags:`

func usage() {
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
	for _, llPath := range flag.Args() {
		err := ll2go(llPath)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// ll2go parses the provided LLVM IR assembly file and decompiles it to Go
// source code.
func ll2go(llPath string) error {
	// File name and file path without extension.
	baseName := pathutil.FileName(llPath)
	basePath := pathutil.TrimExt(llPath)

	// TODO: Create graphs in /tmp/xxx_graphs/*.dot

	// Create temporary foo.dot file, e.g.
	//
	//    foo.ll -> foo_graphs/*.dot
	dotDir := basePath + "_graphs"
	if ok, _ := osutil.Exists(dotDir); !ok {
		if !flagQuiet {
			log.Printf("Creating control flow graphs for %q.\n", filepath.Base(llPath))
		}
		cmd := exec.Command("ll2dot", "-q", "-funcs", flagFuncs, "-f", llPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Parse foo.ll
	module, err := asm.ParseFile(llPath)
	if err != nil {
		return errors.WithStack(err)
	}

	// Force generate local IDs for unnamed local variables.
	_ = module.String()

	// Get function names.
	var funcNames []string
	if len(flagFuncs) > 0 {
		// Get function names from command line flag:
		//
		//    -funcs="foo,bar"
		funcNames = strings.Split(flagFuncs, ",")
	} else {
		// Get all function names.
		for _, llFunc := range module.Funcs {
			if len(llFunc.Blocks) == 0 {
				// Ignore function declarations (e.g. functions without bodies).
				continue
			}
			funcNames = append(funcNames, llFunc.Name)
		}
	}

	// Locate package name.
	pkgName := flagPkgName
	if len(flagPkgName) == 0 {
		pkgName = baseName
		for _, funcName := range funcNames {
			if funcName == "main" {
				pkgName = "main"
				break
			}
		}
	}

	// Create foo.go.
	file := &ast.File{
		Name: newIdent(pkgName),
	}

	// TODO: Implement support for global variables.

	// Parse each function.
	for _, funcName := range funcNames {
		if !flagQuiet {
			log.Printf("Parsing function: %q\n", funcName)
		}
		graph, err := parseCFG(basePath, funcName)
		if err != nil {
			return errors.WithStack(err)
		}

		// Structure the CFG.
		dotDir := basePath + "_graphs"
		dotName := funcName + ".dot"
		dotPath := path.Join(dotDir, dotName)
		jsonName := funcName + ".json"
		jsonPath := path.Join(dotDir, jsonName)
		if ok, _ := osutil.Exists(jsonPath); !ok {
			cmd := exec.Command("restructure", "-o", jsonPath, dotPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if !flagQuiet {
				log.Printf("Structuring function: %q\n", funcName)
			}
			err = cmd.Run()
			if err != nil {
				return errors.WithStack(err)
			}
		}
		var hprims []*xprimitive.Primitive
		fr, err := os.Open(jsonPath)
		if err != nil {
			return errors.WithStack(err)
		}
		defer fr.Close()
		dec := json.NewDecoder(fr)
		err = dec.Decode(&hprims)
		if err != nil {
			return errors.WithStack(err)
		}

		f, err := parseFunc(graph, module, funcName, hprims)
		if err != nil {
			return errors.WithStack(err)
		}
		file.Decls = append(file.Decls, f)
		if flagVerbose && !flagQuiet {
			printFunc(f)
		}
	}

	// Store Go source code to file.
	goPath := basePath + ".go"
	if !flagQuiet {
		log.Printf("Creating: %q\n", goPath)
	}
	return storeFile(goPath, file)
}

// parseCFG parses the control flow graph of the function.
//
// For a source file "foo.ll" containing the functions "bar" and "baz" the
// following DOT files will be created:
//
//    foo_graphs/bar.dot
//    foo_graphs/baz.dot
func parseCFG(basePath, funcName string) (graph *dot.Graph, err error) {
	dotDir := basePath + "_graphs"
	dotName := funcName + ".dot"
	dotPath := fmt.Sprintf("%s/%s", dotDir, dotName)
	return dot.ParseFile(dotPath)
}

// locateFunc locates and returns the specified function in given module. The
// boolean return variable indicates success.
func locateFunc(module *ir.Module, funcName string) (*ir.Function, bool) {
	for _, f := range module.Funcs {
		if f.Name == funcName {
			return f, true
		}
	}
	return nil, false
}

// parseFunc parses the given function and attempts to construct an equivalent
// Go function declaration AST node.
func parseFunc(graph *dot.Graph, module *ir.Module, funcName string, hprims []*xprimitive.Primitive) (*ast.FuncDecl, error) {
	llFunc, ok := locateFunc(module, funcName)
	if !ok {
		return nil, errors.Errorf("unable to locate function %q", funcName)
	}
	if len(llFunc.Blocks) == 0 {
		return nil, errors.Errorf("unable to create AST for %q; expected function definition, got function declaration (e.g. no body)", funcName)
	}

	// Parse each basic block.
	bbs := make(map[string]BasicBlock)
	for _, llBB := range llFunc.Blocks {
		bb, err := parseBasicBlock(llBB)
		if err != nil {
			return nil, err
		}
		bbs[bb.Name()] = bb
		if flagVerbose && !flagQuiet {
			printBB(bb)
		}
	}

	// Replace PHI instructions with assignment statements in the appropriate
	// basic blocks.
	for _, bb := range bbs {
		block, ok := bb.(*basicBlock)
		if !ok {
			return nil, errors.Errorf("invalid basic block type; expected *basicBlock, got %T", bb)
		}
		for ident, defs := range block.phis {
			for _, def := range defs {
				assign := &ast.AssignStmt{
					Lhs: []ast.Expr{newIdent(ident)},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{def.expr},
				}
				bbSrc := bbs[def.bb]
				stmts := bbSrc.Stmts()
				stmts = append(stmts, assign)
				bbSrc.SetStmts(stmts)
			}
		}
	}

	// Perform control flow analysis.
	body, err := restructure(graph, bbs, hprims)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sig := &ast.FuncType{
		Params: &ast.FieldList{},
	}
	if funcName != "main" {
		// TODO: Implement parsing of function signature.
	}
	return createFunc(funcName, sig, body)
}

// createFunc creates and returns a Go function declaration based on the
// provided function name, function signature and basic block.
func createFunc(name string, sig *ast.FuncType, body *ast.BlockStmt) (*ast.FuncDecl, error) {
	f := &ast.FuncDecl{
		Name: newIdent(name),
		Type: sig,
		Body: body,
	}
	return f, nil
}

// storeFile stores the given Go source code to the provided file path.
func storeFile(goPath string, file *ast.File) error {
	// Don't force overwrite Go output file.
	if !flagForce {
		if ok, _ := osutil.Exists(goPath); ok {
			return errors.Errorf("output file %q already exists", goPath)
		}
	}
	f, err := os.Create(goPath)
	if err != nil {
		return err
	}
	defer f.Close()
	fset := token.NewFileSet()
	return printer.Fprint(f, fset, file)
}

// printBB pretty-prints the basic block to stdout.
func printBB(bb BasicBlock) {
	fset := token.NewFileSet()
	fmt.Printf("--- [ basic block %q ] ---\n", bb.Name())
	printer.Fprint(os.Stdout, fset, bb.Stmts())
	fmt.Println()
	if term := bb.Term(); term != nil {
		pretty.Println(term)
	}
	fmt.Println()
}

// printFunc pretty-prints the function to stdout.
func printFunc(f *ast.FuncDecl) {
	fset := token.NewFileSet()
	fmt.Printf("--- [ function %q ] ---\n", f.Name)
	printer.Fprint(os.Stdout, fset, f)
	fmt.Println()
}
