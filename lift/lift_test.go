package lift

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
		// Base directory; which may contain decomp JSON files.
		dir string
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
		{dir: "testdata/x86_32/format", in: "format.bin", out: "format_bin.ll", arch: bin.ArchX86_32},
		{dir: "testdata/x86_32/format", in: "format_elf.o", out: "format_o.ll"},
		{dir: "testdata/x86_32/format", in: "format_elf.so", out: "format_so.ll"},
		{dir: "testdata/x86_32/format", in: "format_elf.out", out: "format_out.ll"},
		{dir: "testdata/x86_64/format", in: "format.bin", out: "format_bin.ll", arch: bin.ArchX86_32},
		{dir: "testdata/x86_64/format", in: "format_elf.o", out: "format_o.ll"},
		{dir: "testdata/x86_64/format", in: "format_elf.so", out: "format_so.ll"},
		{dir: "testdata/x86_64/format", in: "format_elf.out", out: "format_out.ll"},
		// TODO: Add support for COFF files.
		//{in: "testdata/format.coff", out: "testdata/format_coff.ll"},

		// Arithmetic instructions.
		{dir: "testdata/x86_32/arithmetic", in: "arithmetic.so", out: "arithmetic.ll"},
		{dir: "testdata/x86_64/arithmetic", in: "arithmetic.so", out: "arithmetic.ll"},
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("unable to retrieve current working directory; %+v", err)
	}
	for _, g := range golden {
		if err := os.Chdir(wd); err != nil {
			t.Errorf("%q: unable to change working directory; %+v", g.in, err)
			continue
		}
		if err := os.Chdir(g.dir); err != nil {
			t.Errorf("%q: unable to change working directory; %+v", g.in, err)
			continue
		}
		in := filepath.Join(g.dir, g.in)
		log.Printf("testing: %q", in)
		l, err := newLifter(g.in, g.arch)
		if err != nil {
			t.Errorf("%q: unable to prepare lifter; %+v", g.in, err)
			continue
		}

		// Create function lifters.
		for _, funcAddr := range l.FuncAddrs {
			asmFunc, err := l.DecodeFunc(funcAddr)
			if err != nil {
				t.Errorf("%q: unable to decode function; %+v", g.in, err)
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
			t.Errorf("%q: unable to read file: %+v", g.in, err)
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
