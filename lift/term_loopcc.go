package lift

import (
	"github.com/decomp/exp/disasm/x86"
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

// liftTermLOOP lifts the given x86 LOOP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftTermLOOP(term *x86.Inst) error {
	// Loop if ECX register is not zero.
	//    (ECX≠0)
	ecx := f.dec(x86.NewArg(x86asm.ECX, term))
	zero := constant.NewInt(0, types.I32)
	cond := f.cur.NewICmp(ir.IntNE, ecx, zero)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ LOOPE ] ---------------------------------------------------------------

// liftTermLOOPE lifts the given x86 LOOPE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftTermLOOPE(term *x86.Inst) error {
	// Loop if equal and ECX register is not zero.
	//    (ECX≠0 and ZF=1)
	ecx := f.dec(x86.NewArg(x86asm.ECX, term))
	zero := constant.NewInt(0, types.I32)
	zf := f.useStatus(ZF)
	cond1 := f.cur.NewICmp(ir.IntNE, ecx, zero)
	cond2 := zf
	cond := f.cur.NewAnd(cond1, cond2)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ LOOPNE ] --------------------------------------------------------------

// liftTermLOOPNE lifts the given x86 LOOPNE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftTermLOOPNE(term *x86.Inst) error {
	// Loop if not equal and ECX register is not zero.
	//    (ECX≠0 and ZF=0)
	ecx := f.dec(x86.NewArg(x86asm.ECX, term))
	zero := constant.NewInt(0, types.I32)
	zf := f.useStatus(ZF)
	cond1 := f.cur.NewICmp(ir.IntNE, ecx, zero)
	cond2 := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	cond := f.cur.NewAnd(cond1, cond2)
	return f.liftTermJcc(term.Arg(0), cond)
}
