package main

import (
	"fmt"
	"strings"

	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"golang.org/x/arch/x86/x86asm"
)

// An Arg is a single x86 instruction argument.
type Arg struct {
	// x86 instruction argument.
	x86asm.Arg
	// Address of the instruction pointer; used for relative offset arguments.
	addr bin.Address
}

// Arg returns the i:th argument of the instruction.
func (inst *Inst) Arg(i int) *Arg {
	return &Arg{
		Arg:  inst.Args[i],
		addr: inst.addr + bin.Address(inst.Len),
	}
}

// --- [ usage ] ---------------------------------------------------------------

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
		return f.useAddr(addr)
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// useReg loads and returns the value of the given x86 register, emitting code
// to f.
func (f *Func) useReg(reg x86asm.Reg) value.Value {
	src := f.reg(reg)
	return f.cur.NewLoad(src)
}

// useMem loads and returns the value of the given memory argument, emitting
// code to f.
func (f *Func) useMem(mem x86asm.Mem) value.Value {
	src := f.mem(mem)
	return f.cur.NewLoad(src)
}

// useStatus loads and returns the value of the given x86 status flag, emitting
// code to f.
func (f *Func) useStatus(status StatusFlag) value.Value {
	src := f.status(status)
	return f.cur.NewLoad(src)
}

// useAddr loads and returns the value of the given address, emitting code to f.
func (f *Func) useAddr(addr bin.Address) value.Value {
	src, ok := f.addr(addr)
	if !ok {
		panic(fmt.Errorf("unable to locate value at address %v", addr))
	}
	return f.cur.NewLoad(src)
}

// --- [ definition ] ----------------------------------------------------------

// defArg stores the value to the given argument, emitting code to f.
func (f *Func) defArg(arg *Arg, v value.Value) {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		f.defReg(a, v)
	case x86asm.Mem:
		f.defMem(a, v)
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// defReg stores the value to the given x86 register, emitting code to f.
func (f *Func) defReg(reg x86asm.Reg, v value.Value) {
	dst := f.reg(reg)
	f.cur.NewStore(v, dst)
}

// defMem stores the value to the given memory argument, emitting code to f.
func (f *Func) defMem(mem x86asm.Mem, v value.Value) {
	dst := f.mem(mem)
	f.cur.NewStore(v, dst)
}

// defStatus stores the value to the given x86 status flag, emitting code to f.
func (f *Func) defStatus(status StatusFlag, v value.Value) {
	dst := f.status(status)
	f.cur.NewStore(v, dst)
}

// defAddr stores the value to the given address, emitting code to f.
func (f *Func) defAddr(addr bin.Address, v value.Value) {
	dst, ok := f.addr(addr)
	if !ok {
		panic(fmt.Errorf("unable to locate value at address %v", addr))
	}
	f.cur.NewStore(v, dst)
}

// --- [ pointer to value ] ----------------------------------------------------

// reg returns a pointer to the LLVM IR value associated with the given x86
// register.
func (f *Func) reg(reg x86asm.Reg) value.Value {
	if v, ok := f.regs[reg]; ok {
		return v
	}
	typ := regType(reg)
	v := ir.NewAlloca(typ)
	name := strings.ToLower(reg.String())
	v.SetName(name)
	f.regs[reg] = v
	return v
}

// mem returns a pointer to the LLVM IR value associated with the given memory
// argument, emitting code to f.
func (f *Func) mem(mem x86asm.Mem) value.Value {
	// Segment:[Base+Scale*Index+Disp].

	if mem.Segment == 0 && mem.Base == 0 && mem.Scale == 0 && mem.Index == 0 {
		addr := bin.Address(mem.Disp)
		v, ok := f.addr(addr)
		if !ok {
			panic(fmt.Errorf("unable to locate value at address %v", addr))
		}
		return v
	}

	// TODO: Add proper support for memory arguments.
	//    Segment Reg
	//    Base    Reg
	//    Scale   uint8
	//    Index   Reg
	//    Disp    int64

	panic("not yet implemented")
}

// status returns a pointer to the LLVM IR value associated with the given x86
// status flag.
func (f *Func) status(status StatusFlag) value.Value {
	if v, ok := f.statusFlags[status]; ok {
		return v
	}
	panic(fmt.Errorf("unable to locate status flag %v", status))
}

// addr returns a pointer to the LLVM IR value associated with the given
// address, emitting code to f. The returned value is one of *ir.BasicBlock,
// *ir.Global and *ir.Function, and the boolean value indicates success
func (f *Func) addr(addr bin.Address) (value.Value, bool) {
	if block, ok := f.blocks[addr]; ok {
		return block, true
	}
	if g, ok := f.d.globals[addr]; ok {
		return g, true
	}
	fmt.Printf("unable to locate value at address %v\n", addr)
	panic("not yet implemented")
}
