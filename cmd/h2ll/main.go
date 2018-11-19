// The h2ll tool converts C headers to LLVM IR function declarations  (*.h ->
// *.ll).
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

func usage() {
	const use = `
Convert C headers to LLVM IR function declarations (*.h -> *.ll).

Usage:

	h2ll [OPTION]... FILE.h

Flags:
`
	fmt.Fprint(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line flags.
	var (
		// jsonPath specifies the path to a JSON file with function signatures.
		jsonPath string
		// output specifies the output path.
		output string
	)
	flag.StringVar(&jsonPath, "sigs", "sigs.json", "JSON file with function signatures")
	flag.StringVar(&output, "o", "", "output path")
	flag.Parse()
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	hPath := flag.Arg(0)

	// Parse JSON file containing function signatures.
	sigs, funcAddrs, err := parseSigs(jsonPath)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// Convert function signatures to LLVM IR.
	buf, err := ioutil.ReadFile(hPath)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	input := string(buf)
	old, err := compile(input)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	funcs, err := llFuncSigs(old, sigs, funcAddrs)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// Store C header output.
	w := os.Stdout
	if len(output) > 0 {
		f, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	}
	module := ir.NewModule()
	module.TypeDefs = old.TypeDefs
	module.Funcs = funcs
	if _, err := w.WriteString(module.String()); err != nil {
		log.Fatalf("%+v", err)
	}
}

// parseSigs parses the given JSON file containing function signatures.
func parseSigs(jsonPath string) (map[bin.Address]FuncSig, []bin.Address, error) {
	// Parse file.
	sigs := make(map[bin.Address]FuncSig)
	if err := parseJSON(jsonPath, &sigs); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	var funcAddrs []bin.Address
	for funcAddr := range sigs {
		funcAddrs = append(funcAddrs, funcAddr)
	}
	sort.Sort(bin.Addresses(funcAddrs))
	return sigs, funcAddrs, nil
}

// FuncSig represents a function signature.
type FuncSig struct {
	// Function name.
	Name string `json:"name"`
	// Function signature.
	Sig string `json:"sig"`
}

// ### [ Helper functions ] ####################################################

// parseJSON parses the given JSON file and stores the result into v.
func parseJSON(jsonPath string, v interface{}) error {
	f, err := os.Open(jsonPath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	dec := json.NewDecoder(br)
	return dec.Decode(v)
}
