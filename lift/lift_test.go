package lift

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/decomp/exp/bin"
	_ "github.com/decomp/exp/bin/elf" // register ELF decoder
	_ "github.com/decomp/exp/bin/pe"  // register PE decoder
	_ "github.com/decomp/exp/bin/pef" // register PEF decoder
	"github.com/decomp/exp/bin/raw"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

func TestLift(t *testing.T) {
	golden := []struct {
		// Path to input binary executable or object file.
		in string
		// Path to output LLVM IR assembly file.
		out string
		// Raw machine architecture; or 0 if any format other than raw.
		arch bin.Arch
	}{
		// File formats.
		//
		//   * .bin  - raw executable files
		//   * .o    - ELF object files
		//   * .so   - ELF shared object files
		//   * .out  - ELF executable files
		//   * .coff - COFF object files
		{in: "testdata/x86_32/format.bin", out: "testdata/x86_32/format_bin.ll", arch: bin.ArchX86_32},
		{in: "testdata/x86_32/format_elf.o", out: "testdata/x86_32/format_o.ll"},
		{in: "testdata/x86_32/format_elf.so", out: "testdata/x86_32/format_so.ll"},
		{in: "testdata/x86_32/format_elf.out", out: "testdata/x86_32/format_out.ll"},
		{in: "testdata/x86_64/format.bin", out: "testdata/x86_64/format_bin.ll", arch: bin.ArchX86_32},
		{in: "testdata/x86_64/format_elf.o", out: "testdata/x86_64/format_o.ll"},
		{in: "testdata/x86_64/format_elf.so", out: "testdata/x86_64/format_so.ll"},
		{in: "testdata/x86_64/format_elf.out", out: "testdata/x86_64/format_out.ll"},
		// TODO: Add support for COFF files.
		//{in: "testdata/format.coff", out: "testdata/format_coff.ll"},

		// Arithmetic instructions.
		//{in: "testdata/arithmetic.so", out: "testdata/arithmetic.ll"},
	}
	for _, g := range golden {
		log.Printf("testing: %q", g.in)
		l, err := newLifter(g.in, g.arch)
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

// newLifter returns a new x86 to LLVM IR lifter for the given binary
// executable.
func newLifter(path string, arch bin.Arch) (*Lifter, error) {
	if arch != 0 {
		file, err := raw.ParseFile(path, arch)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return NewLifter(file)
	}
	file, err := bin.ParseFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return NewLifter(file)
}
