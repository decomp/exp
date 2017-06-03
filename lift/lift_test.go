package lift

import (
	"io/ioutil"
	"testing"

	"github.com/decomp/exp/bin"
	_ "github.com/decomp/exp/bin/elf" // register ELF decoder
	_ "github.com/decomp/exp/bin/pe"  // register PE decoder
	_ "github.com/decomp/exp/bin/pef" // register PEF decoder
	"github.com/llir/llvm/ir"
)

func TestLift(t *testing.T) {
	golden := []struct {
		// Path to input binary executable or object file.
		in string
		// Path to output LLVM IR assembly file.
		out string
	}{
		{in: "testdata/arithmetic.so", out: "testdata/arithmetic.ll"},
	}
	for _, g := range golden {
		file, err := bin.ParseFile(g.in)
		if err != nil {
			t.Errorf("%q: unable to parse file; %v", g.in, err)
			continue
		}
		l, err := NewLifter(file)
		if err != nil {
			t.Errorf("%q: unable to prepare lifter; %v", g.in, err)
			continue
		}

		// Create function lifters.
		for _, funcAddr := range l.FuncAddrs {
			asmFunc, err := l.DecodeFunc(funcAddr)
			if err != nil {
				t.Errorf("%q: unable to decode function; %v", g.in, err)
				continue
			}
			f := l.NewFunc(asmFunc)
			l.Funcs[funcAddr] = f
		}

		// Lift functions.
		module := &ir.Module{}
		for _, funcAddr := range l.FuncAddrs {
			f, ok := l.Funcs[funcAddr]
			if !ok {
				continue
			}
			f.Lift()
			module.Funcs = append(module.Funcs, f.Function)
		}
		buf, err := ioutil.ReadFile(g.out)
		if err != nil {
			t.Errorf("%q: unable to read file: %v", g.in, err)
			continue
		}
		got := module.String()
		want := string(buf)
		if got != want {
			t.Errorf("%q: module mismatch; expected `%v`, got `%v`", g.in, want, got)
			continue
		}
	}
}
