package main

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"golang.org/x/arch/x86/x86asm"
)

// Loop According to ECX Counter
//
//    (ECX≠0)            LOOP     Loop if ECX register is not zero.
//    (ECX≠0 and ZF=1)   LOOPE    Loop if equal and ECX register is not zero.
//    (ECX≠0 and ZF=0)   LOOPNE   Loop if not equal and ECX register is not zero.
//
// ref: $ 3.2 LOOP/LOOPcc - Loop According to ECX Counter, Intel 64 and IA-32
// Architectures Software Developer's Manual

// --- [ LOOP ] ----------------------------------------------------------------

// emitTermLOOP translates the given x86 LOOP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitTermLOOP(term *Inst) error {
	// Loop if ECX register is not zero.
	//    (ECX≠0)
	ecx := f.dec(NewArg(x86asm.ECX, term))
	zero := constant.NewInt(0, types.I32)
	cond := f.cur.NewICmp(ir.IntNE, ecx, zero)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ LOOPE ] ---------------------------------------------------------------

// emitTermLOOPE translates the given x86 LOOPE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitTermLOOPE(term *Inst) error {
	// Loop if equal and ECX register is not zero.
	//    (ECX≠0 and ZF=1)
	ecx := f.dec(NewArg(x86asm.ECX, term))
	zero := constant.NewInt(0, types.I32)
	zf := f.useStatus(ZF)
	cond1 := f.cur.NewICmp(ir.IntNE, ecx, zero)
	cond2 := zf
	cond := f.cur.NewAnd(cond1, cond2)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ LOOPNE ] --------------------------------------------------------------

// emitTermLOOPNE translates the given x86 LOOPNE instruction to LLVM IR,
// emitting code to f.
func (f *Func) emitTermLOOPNE(term *Inst) error {
	// Loop if not equal and ECX register is not zero.
	//    (ECX≠0 and ZF=0)
	ecx := f.dec(NewArg(x86asm.ECX, term))
	zero := constant.NewInt(0, types.I32)
	zf := f.useStatus(ZF)
	cond1 := f.cur.NewICmp(ir.IntNE, ecx, zero)
	cond2 := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	cond := f.cur.NewAnd(cond1, cond2)
	return f.emitTermJcc(term.Arg(0), cond)
}
