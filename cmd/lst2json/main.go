// The lst2json tool extracts information for decomp from IDA assembly listings
// (*.lst -> *.json).
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/decomp/exp/bin"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// dbg represents a logger with the "lst2json:" prefix, which logs debug
// messages to standard error.
var dbg = log.New(os.Stderr, term.RedBold("lst2json:")+" ", 0)

func usage() {
	const use = `
Extract information for decomp from IDA assembly listings (*.lst -> *.json).

Usage:

	lst2json [OPTION]... FILE.lst

Flags:
`
	fmt.Fprint(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line flags.
	flag.Parse()
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	lstPath := flag.Arg(0)

	if err := extract(lstPath); err != nil {
		log.Fatalf("%+v", err)
	}
}

// extract extracts information for decomp from the given IDA assembly listing.
func extract(lstPath string) error {
	// Read file.
	input, err := ioutil.ReadFile(lstPath)
	if err != nil {
		return errors.WithStack(err)
	}

	// Regular expressions for locating addresses.
	const (
		// Functions (and basic blocks).
		regFunc = `[.]text[:]00([0-9a-fA-F]+)[^\n]+proc[ \t]near`
		// Basic blocks.
		regFallthrough = `[ \t]+(loop|loope|loopne|ja|jb|jbe|jecxz|jg|jge|jl|jle|jnb|jns|jnz|jp|js|jz)[ \t]+[^\n]*\n[.]text[:]00([0-9a-fA-F]+)`
		regTarget      = `[.]text[:]00([0-9a-fA-F]+)[ \t][$@_a-zA-Z][$@_a-zA-Z0-9]+:`
		// Instructions.
		regInst     = `[.]text[:]00([0-9a-fA-F]+)[\t]{2}[ ]{7}[a-z]`
		regTextData = `[.]text[:]00([0-9a-fA-F]+)[\t]{2}[ ]{7}(?:db|dw|dd|dq|align|assume|include|public)[ ]`
		// Data.
		regJumpTable     = `[a-zA-Z]+[:]00([0-9a-fA-F]+)[^\n]*;[ \t]jump[ \t]table`
		regIndirectTable = `[a-zA-Z]+[:]00([0-9a-fA-F]+)[^\n]*;[ \t]indirect[ \t]table`
		regJumpPastData  = `[ \t]+jmp[ \t]+[^\n]*\n[.][a-zA-Z]+[a-zA-Z]+[:]00([0-9a-fA-F]+)[ \t]+; ---------------------------------------------------------------------------[\n][.][a-zA-Z]+[:]00([0-9a-fA-F]+)[ \t]+`
		regAlign         = `; ---------------------------------------------------------------------------[\n][.][a-zA-Z]+[:]00([0-9a-fA-F]+)[ \t]+align[ \t]+`
	)

	// Function, basic block, instruction and data addresses.
	var (
		funcAddrs  []bin.Address
		blockAddrs []bin.Address
		instAddrs  []bin.Address
		dataAddrs  []bin.Address
	)

	// Locate function addresses.
	m := make(map[bin.Address]bool)
	if err := locateAddrs(input, m, regFunc); err != nil {
		return errors.WithStack(err)
	}
	for funcAddr := range m {
		funcAddrs = append(funcAddrs, funcAddr)
	}
	sort.Sort(bin.Addresses(funcAddrs))

	// Locate basic block addresses.
	//
	// Don't reset m, since the address of each function is the address of its
	// entry basic block.
	if err := locateAddrs(input, m, regFallthrough); err != nil {
		return errors.WithStack(err)
	}
	if err := locateAddrs(input, m, regTarget); err != nil {
		return errors.WithStack(err)
	}
	for blockAddr := range m {
		blockAddrs = append(blockAddrs, blockAddr)
	}
	sort.Sort(bin.Addresses(blockAddrs))

	// Locate instruction addresses.
	instAddrSet := make(map[bin.Address]bool)
	if err := locateAddrs(input, instAddrSet, regInst); err != nil {
		return errors.WithStack(err)
	}
	textDataAddrSet := make(map[bin.Address]bool)
	if err := locateAddrs(input, textDataAddrSet, regTextData); err != nil {
		return errors.WithStack(err)
	}
	// Remove data directives from instructions (e.g. "dd 0x00").
	for dataAddr := range textDataAddrSet {
		delete(instAddrSet, dataAddr)
	}
	for instAddr := range instAddrSet {
		instAddrs = append(instAddrs, instAddr)
	}
	sort.Sort(bin.Addresses(instAddrs))

	// Locate data addresses.
	tableAddrs := make(map[bin.Address]bool)
	if err := locateAddrs(input, tableAddrs, regJumpTable); err != nil {
		return errors.WithStack(err)
	}
	for dataAddr := range tableAddrs {
		dataAddrs = append(dataAddrs, dataAddr)
	}
	// Reset m.
	m = make(map[bin.Address]bool)
	if err := locateAddrs(input, m, regIndirectTable); err != nil {
		return errors.WithStack(err)
	}
	if err := locateAddrs(input, m, regJumpPastData); err != nil {
		return errors.WithStack(err)
	}
	if err := locateAddrs(input, m, regAlign); err != nil {
		return errors.WithStack(err)
	}
	for dataAddr := range m {
		dataAddrs = append(dataAddrs, dataAddr)
	}
	sort.Sort(bin.Addresses(dataAddrs))

	// Locate targets of jump tables.
	tables, err := locateTargets(input, tableAddrs)
	if err != nil {
		return errors.WithStack(err)
	}

	// Locate function signatures.
	sigs, err := locateFuncSigs(input)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, funcAddr := range funcAddrs {
		if _, ok := sigs[funcAddr]; !ok {
			dbg.Printf("WARNING: unable to locate function signature for function at %v", funcAddr)
		}
	}

	// Locate imports.
	imports, err := locateImports(input)
	if err != nil {
		return errors.WithStack(err)
	}

	// Locate function chunks.
	chunks, err := locateFuncChunks(input)
	if err != nil {
		return errors.WithStack(err)
	}

	// Store JSON files.
	if err := storeJSON("funcs.json", funcAddrs); err != nil {
		return errors.WithStack(err)
	}
	if err := storeJSON("blocks.json", blockAddrs); err != nil {
		return errors.WithStack(err)
	}
	if err := storeJSON("insts.json", instAddrs); err != nil {
		return errors.WithStack(err)
	}
	if err := storeJSON("data.json", dataAddrs); err != nil {
		return errors.WithStack(err)
	}
	if err := storeJSON("tables.json", tables); err != nil {
		return errors.WithStack(err)
	}
	if err := storeJSON("sigs.json", sigs); err != nil {
		return errors.WithStack(err)
	}
	if err := storeJSON("imports.json", imports); err != nil {
		return errors.WithStack(err)
	}
	if err := storeJSON("chunks.json", chunks); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// FuncSig represents a function signature.
type FuncSig struct {
	// Function name.
	Name string `json:"name"`
	// Function signature.
	Sig string `json:"sig"`
}

// locateFuncSigs locates function signatures in the input IDA assembly listing.
func locateFuncSigs(input []byte) (map[bin.Address]FuncSig, error) {
	const regFuncSig = `(;[ \t]*([^\n]+))?[\n][.]text[:]00([0-9a-fA-F]+)[ \t]+([a-zA-Z0-9_?@$]+)[ \t]+proc[ \t]near`
	re, err := regexp.Compile(regFuncSig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	subs := re.FindAllSubmatch(input, -1)
	sigs := make(map[bin.Address]FuncSig)
	for _, sub := range subs {
		var sig FuncSig
		// parse function signature.
		sig.Sig = string(sub[2])
		// parse address.
		s := string(sub[3])
		x, err := strconv.ParseUint(s, 16, 64)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		addr := bin.Address(x)
		// parse function name.
		sig.Name = string(sub[4])
		sigs[addr] = sig
	}
	return sigs, nil
}

// locateImports locates imports in the input IDA assembly listing.
func locateImports(input []byte) (map[bin.Address]FuncSig, error) {
	const regImport = `([.]idata[:]00[0-9a-fA-F]+[ \t];[ \t]*([^\n]+))?[\n][.]idata[:]00([0-9a-fA-F]+)[ \t]+extrn[ \t]+([a-zA-Z0-9_?@$]+)`
	re, err := regexp.Compile(regImport)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	subs := re.FindAllSubmatch(input, -1)
	sigs := make(map[bin.Address]FuncSig)
	for _, sub := range subs {
		var sig FuncSig
		// parse function signature.
		sig.Sig = string(sub[2])
		// parse address.
		s := string(sub[3])
		x, err := strconv.ParseUint(s, 16, 64)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		addr := bin.Address(x)
		// parse function name.
		sig.Name = string(sub[4])
		sigs[addr] = sig
	}
	return sigs, nil
}

// locateFuncChunks locates addresses of function chunks belonging to parent
// functions.
func locateFuncChunks(input []byte) (map[bin.Address]map[bin.Address]bool, error) {
	const regFuncChunk = `[.]text[:]00([0-9a-fA-F]+)[ \t];[ \t]FUNCTION[ \t]CHUNK[ \t]AT[ \t][.]text[:]00([0-9a-fA-F]+)`
	re, err := regexp.Compile(regFuncChunk)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	chunks := make(map[bin.Address]map[bin.Address]bool)
	subs := re.FindAllSubmatch(input, -1)
	for _, sub := range subs {
		// Parent function address.
		var parent bin.Address
		// Function chunk address.
		var chunk bin.Address
		if err := parent.Set(fmt.Sprintf("0x%s", sub[1])); err != nil {
			return nil, errors.WithStack(err)
		}
		if err := chunk.Set(fmt.Sprintf("0x%s", sub[2])); err != nil {
			return nil, errors.WithStack(err)
		}
		if _, ok := chunks[chunk]; !ok {
			chunks[chunk] = make(map[bin.Address]bool)
		}
		chunks[chunk][parent] = true
	}
	return chunks, nil
}

// locateTargets locates the targets of jump tables in the input IDA assembly
// listing.
func locateTargets(input []byte, tableAddrs map[bin.Address]bool) (map[bin.Address][]bin.Address, error) {
	tables := make(map[bin.Address][]bin.Address)
	for tableAddr := range tableAddrs {
		present := make(map[bin.Address]bool)
		s := fmt.Sprintf("%06X", uint64(tableAddr))
		regTargets := `[.][a-zA-Z]+[:]00` + s + `[^\n]*? dd (([^\n]*?offset[ \t]loc_([0-9a-fA-F]+))+)`
		re, err := regexp.Compile(regTargets)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		subs := re.FindAllSubmatch(input, -1)
		for _, sub := range subs {
			line := sub[1]
			// line contains data formatted as follows.
			//
			//    offset loc_422F0B, offset loc_422F0B, offset loc_422F1B
			re, err := regexp.Compile("loc_([0-9a-fA-F]+)")
			if err != nil {
				return nil, errors.WithStack(err)
			}
			subs := re.FindAllSubmatch(line, -1)
			for _, sub := range subs {
				var target bin.Address
				s := "0x" + string(sub[1])
				if err := target.Set(s); err != nil {
					return nil, errors.WithStack(err)
				}
				if present[target] {
					// skip if target already present.
					continue
				}
				tables[tableAddr] = append(tables[tableAddr], target)
				present[target] = true
			}
		}
	}
	return tables, nil
}

// locateAddrs locates addresses in the input IDA assembly listing based on the
// given regular expression.
func locateAddrs(input []byte, m map[bin.Address]bool, reg string) error {
	re, err := regexp.Compile(reg)
	if err != nil {
		return errors.WithStack(err)
	}
	subs := re.FindAllSubmatch(input, -1)
	for _, sub := range subs {
		s := string(sub[len(sub)-1])
		x, err := strconv.ParseUint(s, 16, 64)
		if err != nil {
			return errors.WithStack(err)
		}
		addr := bin.Address(x)
		m[addr] = true
	}
	return nil
}

// storeJSON stores a JSON encoded representation of the addresses to the given
// file.
func storeJSON(path string, v interface{}) error {
	buf, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return errors.WithStack(err)
	}
	buf = append(buf, '\n')
	if err := ioutil.WriteFile(path, buf, 0644); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
