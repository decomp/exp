package main

import (
	"fmt"

	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
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

// useArg returns the value held by the given argument, emitting code to f.
func (f *Func) useArg(arg *Arg) value.Value {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		return f.useReg(a)
	case x86asm.Mem:
		return f.useMem(a)
	case x86asm.Imm:
		return constant.NewInt(int64(a), types.I32)
	case x86asm.Rel:
		addr := arg.addr + bin.Address(a)
		return f.useGlobal(addr)
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// useReg loads and returns the value of the given x86 register, emitting code
// to f.
func (f *Func) useReg(mem x86asm.Reg) value.Value {
	panic("not yet implemented")
}

// useMem loads and returns the value of the given memory argument, emitting
// code to f.
func (f *Func) useMem(mem x86asm.Mem) value.Value {
	panic("not yet implemented")
}

// useStatus loads and returns the value of the given x86 status flag, emitting
// code to f.
func (f *Func) useStatus(status StatusFlag) value.Value {
	panic("not yet implemented")
}

// useGlobal loads and returns the value of the given global address, emitting
// code to f.
func (f *Func) useGlobal(addr bin.Address) value.Value {
	panic("not yet implemented")
}

// defArg stores the value to the given argument, emitting code to f.
func (f *Func) defArg(arg *Arg, v value.Value) {
	switch arg := arg.Arg.(type) {
	case x86asm.Reg:
	case x86asm.Mem:
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// defReg stores the value to the given x86 register, emitting code to f.
func (f *Func) defReg(mem x86asm.Reg, v value.Value) {
	panic("not yet implemented")
}

// defMem stores the value to the given memory argument, emitting code to f.
func (f *Func) defMem(mem x86asm.Mem, v value.Value) {
	panic("not yet implemented")
}

// defStatus stores the value to the given x86 status flag, emitting code to f.
func (f *Func) defStatus(status StatusFlag, v value.Value) {
	panic("not yet implemented")
}

// defGlobal stores the value to the given global address, emitting code to f.
func (f *Func) defGlobal(addr bin.Address, v value.Value) {
	panic("not yet implemented")
}

// reg returns a pointer to the LLVM IR value associated with the given x86
// register.
func (f *Func) reg(reg x86asm.Reg) value.Value {
	panic("not yet implemented")
}

// status returns a pointer to the LLVM IR value associated with the given x86
// status flag.
func (f *Func) status(status StatusFlag) value.Value {
	panic("not yet implemented")
}

// global returns a pointer to the LLVM IR value associated with the given
// global address, and a boolean value indicating success.
func (f *Func) global(addr bin.Address) (value.Value, bool) {
	panic("not yet implemented")
}
