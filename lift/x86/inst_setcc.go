package x86

import (
	"github.com/decomp/exp/disasm/x86"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
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

// liftInstSETA lifts the given x86 SETA instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETA(inst *x86.Inst) error {
	// Set byte if above.
	//    (CF=0 and ZF=0)
	cf := f.useStatus(CF)
	zf := f.useStatus(ZF)
	cond1 := f.cur.NewICmp(enum.IPredEQ, cf, constant.False)
	cond2 := f.cur.NewICmp(enum.IPredEQ, zf, constant.False)
	cond := f.cur.NewAnd(cond1, cond2)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETAE ] ---------------------------------------------------------------

// liftInstSETAE lifts the given x86 SETAE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETAE(inst *x86.Inst) error {
	// Set byte if above or equal.
	//    (CF=0)
	cf := f.useStatus(CF)
	cond := f.cur.NewICmp(enum.IPredEQ, cf, constant.False)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETBE ] ---------------------------------------------------------------

// liftInstSETBE lifts the given x86 SETBE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETBE(inst *x86.Inst) error {
	// Set byte if below or equal.
	//    (CF=1 or ZF=1)
	cf := f.useStatus(CF)
	zf := f.useStatus(ZF)
	cond1 := cf
	cond2 := zf
	cond := f.cur.NewOr(cond1, cond2)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETB ] ----------------------------------------------------------------

// liftInstSETB lifts the given x86 SETB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETB(inst *x86.Inst) error {
	// Set byte if below.
	//    (CF=1)
	cf := f.useStatus(CF)
	cond := cf
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETNO ] ---------------------------------------------------------------

// liftInstSETNO lifts the given x86 SETNO instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETNO(inst *x86.Inst) error {
	// Set byte if not overflow.
	//    (OF=0)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(enum.IPredEQ, of, constant.False)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETO ] ----------------------------------------------------------------

// liftInstSETO lifts the given x86 SETO instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETO(inst *x86.Inst) error {
	// Set byte if overflow.
	//    (OF=1)
	of := f.useStatus(OF)
	cond := of
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETNP ] ---------------------------------------------------------------

// liftInstSETNP lifts the given x86 SETNP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETNP(inst *x86.Inst) error {
	// Set byte if not parity.
	//    (PF=0)
	pf := f.useStatus(PF)
	cond := f.cur.NewICmp(enum.IPredEQ, pf, constant.False)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETP ] ----------------------------------------------------------------

// liftInstSETP lifts the given x86 SETP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETP(inst *x86.Inst) error {
	// Set byte if parity.
	//    (PF=1)
	pf := f.useStatus(PF)
	cond := pf
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETNS ] ---------------------------------------------------------------

// liftInstSETNS lifts the given x86 SETNS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETNS(inst *x86.Inst) error {
	// Set byte if not sign.
	//    (SF=0)
	sf := f.useStatus(SF)
	cond := f.cur.NewICmp(enum.IPredEQ, sf, constant.False)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETS ] ----------------------------------------------------------------

// liftInstSETS lifts the given x86 SETS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETS(inst *x86.Inst) error {
	// Set byte if sign.
	//    (SF=1)
	sf := f.useStatus(SF)
	cond := sf
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETGE ] ---------------------------------------------------------------

// liftInstSETGE lifts the given x86 SETGE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETGE(inst *x86.Inst) error {
	// Set byte if greater or equal.
	//    (SF=OF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(enum.IPredEQ, sf, of)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETL ] ----------------------------------------------------------------

// liftInstSETL lifts the given x86 SETL instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETL(inst *x86.Inst) error {
	// Set byte if less.
	//    (SF≠OF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(enum.IPredNE, sf, of)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETG ] ----------------------------------------------------------------

// liftInstSETG lifts the given x86 SETG instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETG(inst *x86.Inst) error {
	// Set byte if greater.
	//    (ZF=0 and SF=OF)
	zf := f.useStatus(ZF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond1 := f.cur.NewICmp(enum.IPredEQ, zf, constant.False)
	cond2 := f.cur.NewICmp(enum.IPredEQ, sf, of)
	cond := f.cur.NewAnd(cond1, cond2)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETNE ] ---------------------------------------------------------------

// liftInstSETNE lifts the given x86 SETNE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETNE(inst *x86.Inst) error {
	// Set byte if not equal.
	//    (ZF=0)
	zf := f.useStatus(ZF)
	cond := f.cur.NewICmp(enum.IPredEQ, zf, constant.False)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETLE ] ---------------------------------------------------------------

// liftInstSETLE lifts the given x86 SETLE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETLE(inst *x86.Inst) error {
	// Set byte if less or equal.
	//    (ZF=1 or SF≠OF)
	zf := f.useStatus(ZF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond1 := zf
	cond2 := f.cur.NewICmp(enum.IPredNE, sf, of)
	cond := f.cur.NewOr(cond1, cond2)
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// --- [ SETE ] ----------------------------------------------------------------

// liftInstSETE lifts the given x86 SETE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETE(inst *x86.Inst) error {
	// Set byte if equal.
	//    (ZF=1)
	zf := f.useStatus(ZF)
	cond := zf
	return f.liftInstSETcc(inst.Arg(0), cond)
}

// === [ Helper functions ] ====================================================

// liftInstSETcc lifts the given x86 SETcc instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSETcc(arg *x86.Arg, cond value.Value) error {
	targetTrue := &ir.BasicBlock{}
	exit := &ir.BasicBlock{}
	f.Blocks = append(f.Blocks, targetTrue)
	f.Blocks = append(f.Blocks, exit)
	f.cur.NewCondBr(cond, targetTrue, exit)
	f.cur = targetTrue
	one := constant.NewInt(types.I8, 1)
	f.defArgElem(arg, one, types.I8)
	targetTrue.NewBr(exit)
	f.cur = exit
	return nil
}
