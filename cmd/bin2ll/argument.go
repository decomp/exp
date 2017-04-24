package main

import (
	"fmt"

	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/ir/value"
	"golang.org/x/arch/x86/x86asm"
)

// An Arg is a single x86 instruction argument.
type Arg struct {
	// x86 instruction argument.
	x86asm.Arg
	// Address of the parent instruction; used for relative offset arguments.
	addr bin.Address
}

// Arg returns the i:th argument of the instruction.
func (inst *Inst) Arg(i int) *Arg {
	return &Arg{
		Arg:  inst.Args[i],
		addr: inst.addr,
	}
}

// useArg returns the value held by the argument, emitting code to f.
func (f *Func) useArg(arg *Arg) value.Value {
	switch arg := arg.Arg.(type) {
	//case x86asm.Reg:
	//case x86asm.Mem:
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// defArg stores the value to the argument, emitting code to f.
func (f *Func) defArg(arg *Arg, v value.Value) {
	switch arg := arg.Arg.(type) {
	//case x86asm.Reg:
	//case x86asm.Mem:
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}
