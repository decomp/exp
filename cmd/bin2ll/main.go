// The bin2ll tool lifts binary executables to equivalent LLVM IR assembly
// (*.exe -> *.ll).
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/decomp/exp/bin"
	_ "github.com/decomp/exp/bin/elf" // register ELF decoder
	_ "github.com/decomp/exp/bin/pe"  // register PE decoder
	_ "github.com/decomp/exp/bin/pef" // register PEF decoder
	"github.com/decomp/exp/bin/raw"
	"github.com/decomp/exp/lift"
	"github.com/gonum/graph"
	"github.com/gonum/graph/encoding/dot"
	"github.com/graphism/simple"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

// Loggers.
var (
	// dbg represents a logger with the "bin2ll:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.MagentaBold("bin2ll:")+" ", 0)
	// warn represents a logger with the "warning:" prefix, which logs warning
	// messages to standard error.
	warn = log.New(os.Stderr, term.RedBold("warning:")+" ", 0)
)

func usage() {
	const use = `
Lift binary executables to equivalent LLVM IR assembly (*.exe -> *.ll).

Usage:

	bin2ll [OPTION]... FILE

Flags:
`
	fmt.Fprint(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line arguments.
	var (
		// blockAddr specifies a basic block address to lift.
		blockAddr bin.Address
		// TODO: Remove -first flag and firstAddr.
		// firstAddr specifies the first function address to lift.
		firstAddr bin.Address
		// funcAddr specifies a function address to lift.
		funcAddr bin.Address
		// TODO: Remove -last flag and lastAddr.
		// lastAddr specifies the last function address to disassemble.
		lastAddr bin.Address
		// output specifies the output path.
		output string
		// quiet specifies whether to suppress non-error messages.
		quiet bool
		// rawArch specifies the machine architecture of a raw binary executable.
		rawArch bin.Arch
		// rawEntry specifies the entry point of a raw binary executable.
		rawEntry bin.Address
		// rawBase specifies the base address of a raw binary executable.
		rawBase bin.Address
	)
	flag.Usage = usage
	flag.Var(&blockAddr, "block", "basic block address to lift")
	flag.Var(&firstAddr, "first", "first function address to lift")
	flag.Var(&funcAddr, "func", "function address to lift")
	flag.Var(&lastAddr, "last", "last function address to lift")
	flag.StringVar(&output, "o", "", "output path")
	flag.BoolVar(&quiet, "q", false, "suppress non-error messages")
	flag.Var(&rawArch, "raw", "machine architecture of raw binary executable (x86_32, x86_64, PowerPC_32, ...)")
	flag.Var(&rawEntry, "rawentry", "entry point of raw binary executable")
	flag.Var(&rawBase, "rawbase", "base address of raw binary executable")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	binPath := flag.Arg(0)
	// Mute debug and warning messages if `-q` is set.
	if quiet {
		dbg.SetOutput(ioutil.Discard)
		warn.SetOutput(ioutil.Discard)
	}

	// Prepare x86 to LLVM IR lifter for the binary executable.
	l, err := newLifter(binPath, rawArch, rawEntry, rawBase)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Lift basic block.
	if blockAddr != 0 {
		block, err := l.DecodeBlock(blockAddr)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		_ = block
		return
	}

	// Lift function specified by `-func` flag.
	var funcAddrs bin.Addresses
	if funcAddr != 0 {
		funcAddrs = []bin.Address{funcAddr}
	} else {
		for _, funcAddr := range l.FuncAddrs {
			if firstAddr != 0 && funcAddr < firstAddr {
				// skip functions before first address.
				continue
			}
			if lastAddr != 0 && funcAddr >= lastAddr {
				// skip functions after last address.
				break
			}
			funcAddrs = append(funcAddrs, funcAddr)
		}
	}

	// Create function lifters.
	for _, funcAddr := range funcAddrs {
		asmFunc, err := l.DecodeFunc(funcAddr)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		f := l.NewFunc(asmFunc)
		l.Funcs[funcAddr] = f
	}

	// Lift functions.
	for i, funcAddr := range funcAddrs {
		if i != 0 {
			fmt.Println()
		}
		f, ok := l.Funcs[funcAddr]
		if !ok {
			continue
		}
		f.Lift()
		fmt.Println(f)
	}

	// Store LLVM IR output.
	w := os.Stdout
	if len(output) > 0 {
		f, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	}
	var funcs []*ir.Function
	sort.Sort(funcAddrs)
	for _, funcAddr := range funcAddrs {
		f := l.Funcs[funcAddr]
		funcs = append(funcs, f.Function)
	}
	var globals []*ir.Global
	var globalAddrs bin.Addresses
	for globalAddr := range l.Globals {
		globalAddrs = append(globalAddrs, globalAddr)
	}
	sort.Sort(globalAddrs)
	for _, globalAddr := range globalAddrs {
		g := l.Globals[globalAddr]
		globals = append(globals, g)
	}
	m := &ir.Module{
		Types:   l.Types,
		Globals: globals,
		Funcs:   funcs,
	}
	if _, err := fmt.Fprintln(w, m); err != nil {
		log.Fatalf("%+v", err)
	}

	// Create call graph.
	if err := genCallGraph(l.Funcs); err != nil {
		log.Fatalf("%+v", err)
	}
}

type Source struct {
	Name  string
	Start bin.Address
	End   bin.Address
}

var sources = []Source{
	{Name: "_crt.cpp", Start: 0x401000, End: 0x401029},
	{Name: "appfat.cpp", Start: 0x40102A, End: 0x401DA3},
	{Name: "automap.cpp", Start: 0x401DA4, End: 0x40311A},
	{Name: "capture.cpp", Start: 0x40311B, End: 0x4034D8},
	{Name: "codec.cpp", Start: 0x4034D9, End: 0x4037D3},
	{Name: "control.cpp", Start: 0x4037D4, End: 0x407409},
	{Name: "cursor.cpp", Start: 0x40740A, End: 0x4084A5},
	{Name: "dead.cpp", Start: 0x4084A6, End: 0x4086F3},
	{Name: "debug.cpp", Start: 0x4086F4, End: 0x4087B0},
	{Name: "diablo.cpp", Start: 0x4087B1, End: 0x40ACAC},
	{Name: "doom.cpp", Start: 0x40ACAD, End: 0x40ADD5},
	{Name: "drlg_l1.cpp", Start: 0x40ADD6, End: 0x40D356},
	{Name: "drlg_l2.cpp", Start: 0x40D357, End: 0x40FF80},
	{Name: "drlg_l3.cpp", Start: 0x40FF81, End: 0x412654},
	{Name: "drlg_l4.cpp", Start: 0x412655, End: 0x415097},
	{Name: "dthread.cpp", Start: 0x415098, End: 0x415361},
	{Name: "dx.cpp", Start: 0x415362, End: 0x4158A8},
	{Name: "effects.cpp", Start: 0x4158A9, End: 0x415F42},
	{Name: "encrypt.cpp", Start: 0x415F43, End: 0x4161FB},
	{Name: "engine.cpp", Start: 0x4161FC, End: 0x41804D},
	{Name: "error.cpp", Start: 0x41804E, End: 0x4182AC},
	{Name: "exception.cpp", Start: 0x4182AD, End: 0x418865},
	{Name: "gamemenu.cpp", Start: 0x418866, End: 0x418C8A},
	{Name: "gendung.cpp", Start: 0x418C8B, End: 0x419E8A},
	{Name: "gmenu.cpp", Start: 0x419E8B, End: 0x41A552},
	{Name: "help.cpp", Start: 0x41A553, End: 0x41A7B2},
	{Name: "init.cpp", Start: 0x41A7B3, End: 0x41B18F},
	{Name: "interfac.cpp", Start: 0x41B190, End: 0x41B813},
	{Name: "inv.cpp", Start: 0x41B814, End: 0x41F095},
	{Name: "items.cpp", Start: 0x41F096, End: 0x425442},
	{Name: "lighting.cpp", Start: 0x425443, End: 0x426563},
	{Name: "loadsave.cpp", Start: 0x426564, End: 0x4279F1},
	{Name: "log.cpp", Start: 0x4279F2, End: 0x427E0D},
	{Name: "mainmenu.cpp", Start: 0x427E0E, End: 0x428055},
	{Name: "minitext.cpp", Start: 0x428056, End: 0x4283BF},
	{Name: "missiles.cpp", Start: 0x4283C0, End: 0x430FDE},
	{Name: "monster.cpp", Start: 0x430FDF, End: 0x43AD32},
	{Name: "movie.cpp", Start: 0x43AD33, End: 0x43AE8F},
	{Name: "mpqapi.cpp", Start: 0x43AE90, End: 0x43BBA3},
	{Name: "msg.cpp", Start: 0x43BBA4, End: 0x43F848},
	{Name: "msgcmd.cpp", Start: 0x43F849, End: 0x43FAC3},
	{Name: "multi.cpp", Start: 0x43FAC4, End: 0x440DAD},
	{Name: "nthread.cpp", Start: 0x440DAE, End: 0x44121C},
	{Name: "objects.cpp", Start: 0x44121D, End: 0x448754},
	{Name: "hero.cpp", Start: 0x448755, End: 0x448DF4},
	{Name: "palette.cpp", Start: 0x448DF5, End: 0x4493D3},
	{Name: "path.cpp", Start: 0x4493D4, End: 0x4498EB},
	{Name: "pfile.cpp", Start: 0x4498EC, End: 0x44A8E5},
	{Name: "player.cpp", Start: 0x44A8E6, End: 0x450D32},
	{Name: "plrmsg.cpp", Start: 0x450D33, End: 0x450FFD},
	{Name: "portal.cpp", Start: 0x450FFE, End: 0x45138D},
	{Name: "quests.cpp", Start: 0x45138E, End: 0x452830},
	{Name: "restricted.cpp", Start: 0x452831, End: 0x452974},
	{Name: "scrollrt.cpp", Start: 0x452975, End: 0x456624},
	{Name: "setmaps.cpp", Start: 0x456625, End: 0x456A15},
	{Name: "sha1.cpp", Start: 0x456A16, End: 0x456CBA},
	{Name: "sound.cpp", Start: 0x456CBB, End: 0x45744D},
	{Name: "spells.cpp", Start: 0x45744E, End: 0x457A00},
	{Name: "stores.cpp", Start: 0x457A01, End: 0x45C198},
	{Name: "sync.cpp", Start: 0x45C199, End: 0x45C86F},
	{Name: "themes.cpp", Start: 0x45C870, End: 0x45E08B},
	{Name: "tmsg.cpp", Start: 0x45E08C, End: 0x45E150},
	{Name: "town.cpp", Start: 0x45E151, End: 0x46019A},
	{Name: "towners.cpp", Start: 0x46019B, End: 0x4618A4},
	{Name: "track.cpp", Start: 0x4618A5, End: 0x4619A6},
	{Name: "trigs.cpp", Start: 0x4619A7, End: 0x462C6C},
	{Name: "wave.cpp", Start: 0x462C6D, End: 0x46305F},
	{Name: "world.cpp", Start: 0x463060, End: 0x469719},
	{Name: "_crt.cpp", Start: 0x46971A, End: 0x47746F},
	{Name: "pkware.cpp", Start: 0x477470, End: 0x478FFF},
}

type Node struct {
	simple.Node
	Name string
}

func (n Node) DOTAttributes() []dot.Attribute {
	return []dot.Attribute{
		dot.Attribute{Key: "label", Value: fmt.Sprintf("%q", n.Name)},
	}
}

func genCallGraph(funcs map[bin.Address]*lift.Func) error {
	for _, source := range sources {
		nodes := make(map[string]graph.Node)
		g := simple.NewDirectedGraph(0, 0)
		fmt.Println("source:", source.Name)
		for addr, f := range funcs {
			if !(source.Start <= addr && addr <= source.End) {
				continue
			}
			fmt.Println("   func:", addr, f.Name)
			fn, ok := nodes[f.Name]
			if !ok {
				node := Node{
					Node: simple.Node(g.NewNodeID()),
					Name: f.Name,
				}
				fn = node
				nodes[f.Name] = node
				g.AddNode(fn)
			}
			// Callers.
			names := getCallerNames(funcs, f.Name)
			for _, name := range names {
				caller, ok := nodes[name]
				if !ok {
					node := Node{
						Node: simple.Node(g.NewNodeID()),
						Name: name,
					}
					caller = node
					nodes[name] = node
					g.AddNode(caller)
				}
				e := simple.Edge{
					F: caller,
					T: fn,
				}
				g.SetEdge(e)
			}
			// Callees.
			names = getCalleeNames(f.Function)
			for _, name := range names {
				callee, ok := nodes[name]
				if !ok {
					node := Node{
						Node: simple.Node(g.NewNodeID()),
						Name: name,
					}
					callee = node
					nodes[name] = node
					g.AddNode(callee)
				}
				e := simple.Edge{
					F: fn,
					T: callee,
				}
				g.SetEdge(e)
			}
		}
		data, err := dot.Marshal(g, "", "", "\t", false)
		if err != nil {
			return errors.WithStack(err)
		}
		path := fmt.Sprintf("call_graphs/%s.dot", source.Name)
		if err := ioutil.WriteFile(path, data, 0644); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// calls returns the callers and callees of the given functions; where callers
// maps from callee to caller, and callees maps from caller to callee.
func calls(funcs map[bin.Address]*lift.Func) (callers, callees map[string]string) {
	callers = make(map[string]string)
	callees = make(map[string]string)
	for _, f := range funcs {
		caller := f.Name
		for _, block := range f.Blocks {
			for _, inst := range block.Insts {
				if inst, ok := inst.(*ir.InstCall); ok {
					if c, ok := inst.Callee.(value.Named); ok {
						callee := c.GetName()
						callers[callee] = caller
						callees[caller] = callee
					} else {
						log.Fatalf("unable to locate name of callee `%v` in function %q", inst.Callee, caller)
					}
				}
			}
		}
	}
	return callers, callees
}

func getCallerNames(funcs map[bin.Address]*lift.Func, name string) []string {
	m := make(map[string]bool)
	for _, f := range funcs {
		for _, block := range f.Blocks {
			for _, inst := range block.Insts {
				if inst, ok := inst.(*ir.InstCall); ok {
					if c, ok := inst.Callee.(value.Named); ok {
						if c.GetName() == name {
							m[f.Name] = true
						}
					} else {
						log.Fatalf("unable to locate name of callee `%v` in function %q", inst.Callee, f.Name)
					}
				}
			}
		}
	}
	var names []string
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func getCalleeNames(f *ir.Function) []string {
	m := make(map[string]bool)
	for _, block := range f.Blocks {
		for _, inst := range block.Insts {
			if inst, ok := inst.(*ir.InstCall); ok {
				if c, ok := inst.Callee.(value.Named); ok {
					m[c.GetName()] = true
				} else {
					log.Fatalf("unable to locate name of callee `%v`", inst.Callee)
				}
			}
		}
	}
	var names []string
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// newLifter returns a new x86 to LLVM IR lifter for the given binary
// executable.
func newLifter(binPath string, rawArch bin.Arch, rawEntry, rawBase bin.Address) (*lift.Lifter, error) {
	// Parse raw binary executable.
	if rawArch != 0 {
		file, err := raw.ParseFile(binPath, rawArch)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		file.Entry = rawEntry
		file.Sections[0].Addr = rawBase
		return lift.NewLifter(file)
	}
	// Parse binary executable.
	file, err := bin.ParseFile(binPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return lift.NewLifter(file)
}
