// The lst2json tool extracts basic block addresses from IDA assembly listings
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
	"github.com/pkg/errors"
)

func usage() {
	const use = `
Extract basic block addresses from IDA assembly listings (*.lst -> *.json).

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

// extract extracts basic block addresses from the given assembly listing.
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
		regFallthrough = `[ \t]+(ja|jb|jbe|jecxz|jg|jge|jl|jle|jnb|jns|jnz|jp|js|jz)[ \t]+[^\n]*\n[.]text[:]00([0-9a-fA-F]+)`
		regTarget      = `[.]text[:]00([0-9a-fA-F]+)[ \t][$@_a-zA-Z][$@_a-zA-Z0-9]+:`
		// Data.
		regJumpTable     = `[.]text[:]00([0-9a-fA-F]+)[^\n]*;[ \t]jump[ \t]table`
		regIndirectTable = `[.]text[:]00([0-9a-fA-F]+)[^\n]*;[ \t]indirect[ \t]table`
		regJumpPastData  = `[ \t]+jmp[ \t]+[^\n]*\n[.]text[:]00([0-9a-fA-F]+)[ \t]+; ---------------------------------------------------------------------------[\n][.]text[:]00([0-9a-fA-F]+)[ \t]+`
		regAlign         = `; ---------------------------------------------------------------------------[\n][.]text[:]00([0-9a-fA-F]+)[ \t]+align[ \t]+`
	)

	// Function, basic block and data addresses.
	var (
		funcAddrs  []bin.Address
		blockAddrs []bin.Address
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

	// Store JSON files.
	if err := storeJSON("funcs.json", funcAddrs); err != nil {
		return errors.WithStack(err)
	}
	if err := storeJSON("blocks.json", blockAddrs); err != nil {
		return errors.WithStack(err)
	}
	if err := storeJSON("data.json", dataAddrs); err != nil {
		return errors.WithStack(err)
	}
	if err := storeJSON("tables.json", tables); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// locateTargets locates the targets of jump tables within the IDA assembly
// listing.
func locateTargets(input []byte, tableAddrs map[bin.Address]bool) (map[bin.Address][]bin.Address, error) {
	tables := make(map[bin.Address][]bin.Address)
	for tableAddr := range tableAddrs {
		present := make(map[bin.Address]bool)
		s := fmt.Sprintf("%06X", uint64(tableAddr))
		regTargets := `[.]text[:]00` + s + `[^\n]*? dd (([^\n]*?offset[ \t]loc_([0-9a-fA-F]+))+)`
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

// locateAddrs locates addresses in input based on the given regular expression.
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
