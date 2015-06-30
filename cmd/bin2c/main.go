// bin2c is a tool which converts binary executables to equivalent C source
// code.
package main

import (
	"debug/pe"
	"flag"
	"fmt"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mewkiz/pkg/errutil"
)

func usage() {
	const use = `
Usage: bin2c -addr ADDRESS [OPTION]... FILE
Convert binary executables to equivalent C source code.

Flags:
`
	fmt.Fprint(os.Stderr, use[1:])
	flag.PrintDefaults()
}

// Command line flags.
var (
	// flagVerbose specifies whether verbose output is enabled.
	flagVerbose bool
)

// Base address of the ".text" section.
const baseAddr = 0x00401000

func main() {
	// Parse command line arguments.
	var (
		// addr specifies the address to decompile.
		addr address
	)
	flag.BoolVar(&flagVerbose, "v", false, "Enable verbose output.")
	flag.Var(&addr, "addr", "Address of function to decompile.")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	path := flag.Arg(0)
	if addr == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Parse ".text" section.
	text, err := parseText(path)
	if err != nil {
		log.Fatal(err)
	}

	// Sanity check.
	offset := int(addr - baseAddr)
	if offset < 0 || offset >= len(text) {
		log.Fatalf("invalid address; expected >= 0x%X and < 0x%X, got 0x%X", baseAddr, baseAddr+len(text), addr)
	}

	// Convert the given function to C source code.
	fn, err := parseFunc(text, offset)
	if err != nil {
		log.Fatal(err)
	}
	printer.Fprint(os.Stdout, token.NewFileSet(), fn)
}

// parseText parses and returns the ".text" section of the given binary
// executable.
func parseText(path string) (text []byte, err error) {
	f, err := pe.Open(path)
	if err != nil {
		return nil, errutil.Err(err)
	}
	defer f.Close()

	sect := f.Section(".text")
	if sect == nil {
		return nil, errutil.Newf(`unable to locate ".text" section in %q`, path)
	}
	return sect.Data()
}

// address implements the flag.Value interface and allows addresses to be
// specified in hexadecimal format.
type address uint64

// String returns the hexadecimal string representation of v.
func (v *address) String() string {
	return fmt.Sprintf("0x%X", uint64(*v))
}

// Set sets v to the numberic value represented by s.
func (v *address) Set(s string) error {
	base := 10
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		s = s[2:]
		base = 16
	}
	x, err := strconv.ParseUint(s, base, 64)
	if err != nil {
		return errutil.Err(err)
	}
	*v = address(x)
	return nil
}
