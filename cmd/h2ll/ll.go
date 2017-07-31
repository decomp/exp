package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/metadata"
	"github.com/pkg/errors"
)

// llFuncSigs translates the given function signatures from C to LLVM IR.
func llFuncSigs(module *ir.Module, sigs map[bin.Address]FuncSig, funcAddrs []bin.Address) ([]*ir.Function, error) {
	var funcs []*ir.Function
	nameToFunc := make(map[string]*ir.Function)
	for _, f := range module.Funcs {
		if _, ok := nameToFunc[f.Name]; ok {
			return nil, errors.Errorf("function name %q already present", f.Name)
		}
		nameToFunc[f.Name] = f
	}
	for _, funcAddr := range funcAddrs {
		sig := sigs[funcAddr]
		f, ok := locateFunc(sig.Name, nameToFunc)
		if !ok {
			return nil, errors.Errorf("unable to locate function %q", sig.Name)
		}
		f.Parent = nil
		f.Blocks = nil
		f.Metadata = map[string]*metadata.Metadata{
			"addr": &metadata.Metadata{
				Nodes: []metadata.Node{&metadata.String{Val: funcAddr.String()}},
			},
		}
		funcs = append(funcs, f)
	}
	return funcs, nil
}

// locateFunc locates the named function.
func locateFunc(funcName string, nameToFunc map[string]*ir.Function) (*ir.Function, bool) {
	// IDA may include _imp prefix to imports.
	//
	//    a: ExitProcess
	//    b: __imp_ExitProcess
	if strings.HasPrefix(funcName, "__imp_") {
		if f, ok := nameToFunc[funcName[len("__imp_"):]]; ok {
			return f, true
		}
	}

	// IDA seems to ignore a leading underscore in function names when printing
	// the function signature.
	//
	//    a: _crt_cpp_init
	//    b: crt_cpp_init
	if strings.HasPrefix(funcName, "_") {
		if f, ok := nameToFunc[funcName[len("_"):]]; ok {
			return f, true
		}
	}
	if f, ok := nameToFunc[funcName]; ok {
		return f, true
	}
	// TODO: Fix handling of constructors and destructors.
	switch funcName {
	case "??1type_info@@UAE@XZ":
		return locateFunc("type_info_create", nameToFunc)
	case "??_Gtype_info@@UAEPAXI@Z":
		return locateFunc("type_info_delete", nameToFunc)
	}
	//    a: WinMain
	//    b: _WinMain@16
	if pos := strings.IndexAny(funcName, "@"); pos != -1 {
		return locateFunc(funcName[:pos], nameToFunc)
	}
	return nil, false
}

// compile compiles the given C source into LLVM IR.
func compile(input string) (*ir.Module, error) {
	out := &bytes.Buffer{}
	cmd := exec.Command("clang", "-m32", "-S", "-emit-llvm", "-x", "c", "-Wno-return-type", "-Wno-invalid-noreturn", "-o", "-", "-")
	cmd.Stdin = strings.NewReader(input)
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, errors.WithStack(err)
	}
	module, err := asm.ParseBytes(out.Bytes())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return module, nil
}
