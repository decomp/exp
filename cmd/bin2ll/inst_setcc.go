//+build ignore

package main

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Set Byte on Condition
//
//    (CF=0 and ZF=0)     SETA     Set byte if above.
//    (CF=0 and ZF=0)     SETNBE   Set byte if not below or equal.     PRESUDO-instruction
//    (CF=0)              SETAE    Set byte if above or equal.
//    (CF=0)              SETNB    Set byte if not below.              PRESUDO-instruction
//    (CF=0)              SETNC    Set byte if not carry.              PRESUDO-instruction
//    (CF=1 or ZF=1)      SETBE    Set byte if below or equal.
//    (CF=1 or ZF=1)      SETNA    Set byte if not above.              PRESUDO-instruction
//    (CF=1)              SETB     Set byte if below.
//    (CF=1)              SETC     Set byte if carry.                  PRESUDO-instruction
//    (CF=1)              SETNAE   Set byte if not above or equal.     PRESUDO-instruction
//    (OF=0)              SETNO    Set byte if not overflow.
//    (OF=1)              SETO     Set byte if overflow.
//    (PF=0)              SETNP    Set byte if not parity.
//    (PF=0)              SETPO    Set byte if parity odd.             PRESUDO-instruction
//    (PF=1)              SETP     Set byte if parity.
//    (PF=1)              SETPE    Set byte if parity even.            PRESUDO-instruction
//    (SF=0)              SETNS    Set byte if not sign.
//    (SF=1)              SETS     Set byte if sign.
//    (SF=OF)             SETGE    Set byte if greater or equal.
//    (SF=OF)             SETNL    Set byte if not less.               PRESUDO-instruction
//    (SF≠OF)             SETL     Set byte if less.
//    (SF≠OF)             SETNGE   Set byte if not greater or equal.   PRESUDO-instruction
//    (ZF=0 and SF=OF)    SETG     Set byte if greater.
//    (ZF=0 and SF=OF)    SETNLE   Set byte if not less or equal.      PRESUDO-instruction
//    (ZF=0)              SETNE    Set byte if not equal.
//    (ZF=0)              SETNZ    Set byte if not zero.               PRESUDO-instruction
//    (ZF=1 or SF≠OF)     SETLE    Set byte if less or equal.
//    (ZF=1 or SF≠OF)     SETNG    Set byte if not greater.            PRESUDO-instruction
//    (ZF=1)              SETE     Set byte if equal.
//    (ZF=1)              SETZ     Set byte if zero.                   PRESUDO-instruction
//
// ref: $ 4.2 SETcc - Set Byte on Condition, Intel 64 and IA-32 Architectures
// Software Developer's Manual

// --- [ SETA ] ----------------------------------------------------------------

// emitInstSETA translates the given x86 SETA instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETA(inst *Inst) error {
	// Set byte if above.
	//    (CF=0 and ZF=0)
	cf := f.useStatus(CF)
	zf := f.useStatus(ZF)
	cond1 := f.cur.NewICmp(ir.IntEQ, cf, constant.False)
	cond2 := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	cond := f.cur.NewAnd(cond1, cond2)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETAE ] ---------------------------------------------------------------

// emitInstSETAE translates the given x86 SETAE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETAE(inst *Inst) error {
	// Set byte if above or equal.
	//    (CF=0)
	cf := f.useStatus(CF)
	cond := f.cur.NewICmp(ir.IntEQ, cf, constant.False)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETBE ] ---------------------------------------------------------------

// emitInstSETBE translates the given x86 SETBE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETBE(inst *Inst) error {
	// Set byte if below or equal.
	//    (CF=1 or ZF=1)
	cf := f.useStatus(CF)
	zf := f.useStatus(ZF)
	cond1 := cf
	cond2 := zf
	cond := f.cur.NewOr(cond1, cond2)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETB ] ----------------------------------------------------------------

// emitInstSETB translates the given x86 SETB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETB(inst *Inst) error {
	// Set byte if below.
	//    (CF=1)
	cf := f.useStatus(CF)
	cond := cf
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETNO ] ---------------------------------------------------------------

// emitInstSETNO translates the given x86 SETNO instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETNO(inst *Inst) error {
	// Set byte if not overflow.
	//    (OF=0)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(ir.IntEQ, of, constant.False)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETO ] ----------------------------------------------------------------

// emitInstSETO translates the given x86 SETO instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETO(inst *Inst) error {
	// Set byte if overflow.
	//    (OF=1)
	of := f.useStatus(OF)
	cond := of
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETNP ] ---------------------------------------------------------------

// emitInstSETNP translates the given x86 SETNP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETNP(inst *Inst) error {
	// Set byte if not parity.
	//    (PF=0)
	pf := f.useStatus(PF)
	cond := f.cur.NewICmp(ir.IntEQ, pf, constant.False)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETP ] ----------------------------------------------------------------

// emitInstSETP translates the given x86 SETP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETP(inst *Inst) error {
	// Set byte if parity.
	//    (PF=1)
	pf := f.useStatus(PF)
	cond := pf
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETNS ] ---------------------------------------------------------------

// emitInstSETNS translates the given x86 SETNS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETNS(inst *Inst) error {
	// Set byte if not sign.
	//    (SF=0)
	sf := f.useStatus(SF)
	cond := f.cur.NewICmp(ir.IntEQ, sf, constant.False)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETS ] ----------------------------------------------------------------

// emitInstSETS translates the given x86 SETS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETS(inst *Inst) error {
	// Set byte if sign.
	//    (SF=1)
	sf := f.useStatus(SF)
	cond := sf
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETGE ] ---------------------------------------------------------------

// emitInstSETGE translates the given x86 SETGE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETGE(inst *Inst) error {
	// Set byte if greater or equal.
	//    (SF=OF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(ir.IntEQ, sf, of)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETL ] ----------------------------------------------------------------

// emitInstSETL translates the given x86 SETL instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETL(inst *Inst) error {
	// Set byte if less.
	//    (SF≠OF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(ir.IntNE, sf, of)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETG ] ----------------------------------------------------------------

// emitInstSETG translates the given x86 SETG instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETG(inst *Inst) error {
	// Set byte if greater.
	//    (ZF=0 and SF=OF)
	zf := f.useStatus(ZF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond1 := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	cond2 := f.cur.NewICmp(ir.IntEQ, sf, of)
	cond := f.cur.NewAnd(cond1, cond2)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETNE ] ---------------------------------------------------------------

// emitInstSETNE translates the given x86 SETNE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETNE(inst *Inst) error {
	// Set byte if not equal.
	//    (ZF=0)
	zf := f.useStatus(ZF)
	cond := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETLE ] ---------------------------------------------------------------

// emitInstSETLE translates the given x86 SETLE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETLE(inst *Inst) error {
	// Set byte if less or equal.
	//    (ZF=1 or SF≠OF)
	zf := f.useStatus(ZF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond1 := zf
	cond2 := f.cur.NewICmp(ir.IntNE, sf, of)
	cond := f.cur.NewOr(cond1, cond2)
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// --- [ SETE ] ----------------------------------------------------------------

// emitInstSETE translates the given x86 SETE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETE(inst *Inst) error {
	// Set byte if equal.
	//    (ZF=1)
	zf := f.useStatus(ZF)
	cond := zf
	return f.emitInstSETcc(inst.Arg(0), cond)
}

// === [ Helper functions ] ====================================================

// emitInstSETcc translates the given x86 SETcc instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSETcc(arg *Arg, cond value.Value) error {
	targetTrue := &ir.BasicBlock{}
	exit := &ir.BasicBlock{}
	f.AppendBlock(targetTrue)
	f.AppendBlock(exit)
	f.cur.NewCondBr(cond, targetTrue, exit)
	f.cur = targetTrue
	one := constant.NewInt(1, types.I8)
	f.defArgElem(arg, one, types.I8)
	targetTrue.NewBr(exit)
	f.cur = exit
	return nil
}
