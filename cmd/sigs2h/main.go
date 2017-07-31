// The sigs2h tool converts function signatures to C headers (*.json -> *.h).
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"text/template"

	"github.com/decomp/exp/bin"
	"github.com/mewkiz/pkg/errutil"
	"github.com/pkg/errors"
)

func usage() {
	const use = `
Convert function signatures to empty C headers (*.json -> *.h).

Usage:

	sigs2h [OPTION]... FILE.json

Flags:
`
	fmt.Fprint(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line flags.
	var (
		// output specifies the output path.
		output string
	)
	flag.StringVar(&output, "o", "", "output path")
	flag.Parse()
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	jsonPath := flag.Arg(0)

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
	if err := convert(w, jsonPath); err != nil {
		log.Fatalf("%+v", err)
	}
}

// convert converts the the given function signatures to C headers.
func convert(w io.Writer, jsonPath string) error {
	// Parse file.
	sigs := make(map[bin.Address]FuncSig)
	if err := parseJSON(jsonPath, &sigs); err != nil {
		return errors.WithStack(err)
	}
	var funcAddrs []bin.Address
	for funcAddr := range sigs {
		funcAddrs = append(funcAddrs, funcAddr)
	}
	sort.Sort(bin.Addresses(funcAddrs))

	// Convert function signatures to C.
	const source = `
#include <stdint.h> // int8_t, ...
#include <stdarg.h> // va_list
#if __WORDSIZE == 64
	typedef uint64_t size_t;
#else
	typedef uint32_t size_t;
#endif

#include "types.h"

{{ range . }}
// {{ .Addr }}
{{ .Sig }} {}
{{ end }}
`
	var funcSigs []Signature
	for _, funcAddr := range funcAddrs {
		sig := sigs[funcAddr]
		s := sig.Sig
		if len(s) == 0 {
			s = fmt.Sprintf("void %s() /* signature missing */", sig.Name)
		}
		funcSig := Signature{
			Addr: funcAddr,
			Sig:  s,
		}
		funcSigs = append(funcSigs, funcSig)
	}
	t := template.New("signatures")
	if _, err := t.Parse(source[1:]); err != nil {
		return errors.WithStack(err)
	}
	if err := t.Execute(w, funcSigs); err != nil {
		return errutil.Err(err)
	}
	return nil
}

// Signature represents a function signature.
type Signature struct {
	// Function address.
	Addr bin.Address
	// Function signature.
	Sig string
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
