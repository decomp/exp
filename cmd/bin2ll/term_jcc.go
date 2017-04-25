package main

import (
	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// Jump if Condition Is Met
//
//    (CF=0 and ZF=0)    JA      Jump if above.
//    (CF=0 and ZF=0)    JNBE    Jump if not below or equal.     PSEUDO-instruction
//    (CF=0)             JAE     Jump if above or equal.
//    (CF=0)             JNB     Jump if not below.              PSEUDO-instruction
//    (CF=0)             JNC     Jump if not carry.              PSEUDO-instruction
//    (CF=1 or ZF=1)     JBE     Jump if below or equal.
//    (CF=1 or ZF=1)     JNA     Jump if not above.              PSEUDO-instruction
//    (CF=1)             JB      Jump if below.
//    (CF=1)             JC      Jump if carry.                  PSEUDO-instruction
//    (CF=1)             JNAE    Jump if not above or equal.     PSEUDO-instruction
//    (CX=0)             JCXZ    Jump if CX register is zero.
//    (ECX=0)            JECXZ   Jump if ECX register is zero.
//    (OF=0)             JNO     Jump if not overflow.
//    (OF=1)             JO      Jump if overflow.
//    (PF=0)             JNP     Jump if not parity.
//    (PF=0)             JPO     Jump if parity odd.             PSEUDO-instruction
//    (PF=1)             JP      Jump if parity.
//    (PF=1)             JPE     Jump if parity even.            PSEUDO-instruction
//    (RCX=0)            JRCXZ   Jump if RCX register is zero.
//    (SF=0)             JNS     Jump if not sign.
//    (SF=1)             JS      Jump if sign.
//    (SF=OF)            JGE     Jump if greater or equal.
//    (SF=OF)            JNL     Jump if not less.               PSEUDO-instruction
//    (SF≠OF)            JL      Jump if less.
//    (SF≠OF)            JNGE    Jump if not greater or equal.   PSEUDO-instruction
//    (ZF=0 and SF=OF)   JG      Jump if greater.
//    (ZF=0 and SF=OF)   JNLE    Jump if not less or equal.      PSEUDO-instruction
//    (ZF=0)             JNE     Jump if not equal.
//    (ZF=0)             JNZ     Jump if not zero.               PSEUDO-instruction
//    (ZF=1 or SF≠OF)    JLE     Jump if less or equal.
//    (ZF=1 or SF≠OF)    JNG     Jump if not greater.            PSEUDO-instruction
//    (ZF=1)             JE      Jump if equal.
//    (ZF=1)             JZ      Jump if zero.                   PSEUDO-instruction
//
// ref: $ 3.2 Jcc - Jump if Condition Is Met, Intel 64 and IA-32
// Architectures Software Developer's Manual

// --- [ JA ] ------------------------------------------------------------------

// emitTermJA translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJA(term *Inst) error {
	// Jump if above.
	//    (CF=0 and ZF=0)
	cf := f.useStatus(CF)
	zf := f.useStatus(ZF)
	cond1 := f.cur.NewICmp(ir.IntEQ, cf, constant.False)
	cond2 := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	cond := f.cur.NewAnd(cond1, cond2)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JAE ] -----------------------------------------------------------------

// emitTermJAE translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJAE(term *Inst) error {
	// Jump if above or equal.
	//    (CF=0)
	cf := f.useStatus(CF)
	cond := f.cur.NewICmp(ir.IntEQ, cf, constant.False)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JBE ] -----------------------------------------------------------------

// emitTermJBE translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJBE(term *Inst) error {
	// Jump if below or equal.
	//    (CF=1 or ZF=1)
	cf := f.useStatus(CF)
	zf := f.useStatus(ZF)
	cond1 := cf
	cond2 := zf
	cond := f.cur.NewOr(cond1, cond2)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JB ] ------------------------------------------------------------------

// emitTermJB translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJB(term *Inst) error {
	// Jump if below.
	//    (CF=1)
	cf := f.useStatus(CF)
	cond := cf
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JCXZ ] ----------------------------------------------------------------

// emitTermJCXZ translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJCXZ(term *Inst) error {
	// Jump if CX register is zero.
	//    (CX=0)
	pretty.Println("term:", term)
	panic("emitTermJCXZ: not yet implemented")
}

// --- [ JECXZ ] ---------------------------------------------------------------

// emitTermJECXZ translates the given x86 JA terminator to LLVM IR, emitting
// code to f.
func (f *Func) emitTermJECXZ(term *Inst) error {
	// Jump if ECX register is zero.
	//    (ECX=0)
	pretty.Println("term:", term)
	panic("emitTermJECXZ: not yet implemented")
}

// --- [ JNO ] -----------------------------------------------------------------

// emitTermJNO translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJNO(term *Inst) error {
	// Jump if not overflow.
	//    (OF=0)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(ir.IntEQ, of, constant.False)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JO ] ------------------------------------------------------------------

// emitTermJO translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJO(term *Inst) error {
	// Jump if overflow.
	//    (OF=1)
	of := f.useStatus(OF)
	cond := of
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JNP ] -----------------------------------------------------------------

// emitTermJNP translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJNP(term *Inst) error {
	// Jump if not parity.
	//    (PF=0)
	pf := f.useStatus(PF)
	cond := f.cur.NewICmp(ir.IntEQ, pf, constant.False)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JP ] ------------------------------------------------------------------

// emitTermJP translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJP(term *Inst) error {
	// Jump if parity.
	//    (PF=1)
	pf := f.useStatus(PF)
	cond := pf
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JRCXZ ] ---------------------------------------------------------------

// emitTermJRCXZ translates the given x86 JA terminator to LLVM IR, emitting
// code to f.
func (f *Func) emitTermJRCXZ(term *Inst) error {
	// Jump if RCX register is zero.
	//    (RCX=0)
	pretty.Println("term:", term)
	panic("emitTermJRCXZ: not yet implemented")
}

// --- [ JNS ] -----------------------------------------------------------------

// emitTermJNS translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJNS(term *Inst) error {
	// Jump if not sign.
	//    (SF=0)
	sf := f.useStatus(SF)
	cond := f.cur.NewICmp(ir.IntEQ, sf, constant.False)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JS ] ------------------------------------------------------------------

// emitTermJS translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJS(term *Inst) error {
	// Jump if sign.
	//    (SF=1)
	sf := f.useStatus(SF)
	cond := sf
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JGE ] -----------------------------------------------------------------

// emitTermJGE translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJGE(term *Inst) error {
	// Jump if greater or equal.
	//    (SF=OF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(ir.IntEQ, sf, of)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JL ] ------------------------------------------------------------------

// emitTermJL translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJL(term *Inst) error {
	// Jump if less.
	//    (SF≠OF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(ir.IntNE, sf, of)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JG ] ------------------------------------------------------------------

// emitTermJG translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJG(term *Inst) error {
	// Jump if greater.
	//    (ZF=0 and SF=OF)
	zf := f.useStatus(ZF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond1 := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	cond2 := f.cur.NewICmp(ir.IntEQ, sf, of)
	cond := f.cur.NewAnd(cond1, cond2)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JNE ] -----------------------------------------------------------------

// emitTermJNE translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJNE(term *Inst) error {
	// Jump if not equal.
	//    (ZF=0)
	zf := f.useStatus(ZF)
	cond := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JLE ] -----------------------------------------------------------------

// emitTermJLE translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJLE(term *Inst) error {
	// Jump if less or equal.
	//    (ZF=1 or SF≠OF)
	zf := f.useStatus(ZF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond1 := zf
	cond2 := f.cur.NewICmp(ir.IntNE, sf, of)
	cond := f.cur.NewOr(cond1, cond2)
	return f.emitTermJcc(term.Arg(0), cond)
}

// --- [ JE ] ------------------------------------------------------------------

// emitTermJE translates the given x86 JA terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJE(term *Inst) error {
	// Jump if equal.
	//    (ZF=1)
	zf := f.useStatus(ZF)
	cond := zf
	return f.emitTermJcc(term.Arg(0), cond)
}

// === [ Helper functions ] ====================================================

// emitTermJcc translates the given x86 SETcc instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitTermJcc(arg *Arg, cond value.Value) error {
	// Target branch of conditional jump.
	nextAddr := arg.parent.addr + bin.Address(arg.parent.Len)
	targetAddr, ok := f.getAddr(arg)
	if !ok {
		return errors.Errorf("unable to locate address for terminator argument %v", arg)
	}
	target, ok := f.blocks[targetAddr]
	if !ok {
		return errors.Errorf("unable to locate target basic block at %v", targetAddr)
	}
	// Fallthrough branch of conditional jump.
	next, ok := f.blocks[nextAddr]
	if !ok {
		return errors.Errorf("unable to locate fallthrough basic block at %v", nextAddr)
	}
	f.cur.NewCondBr(cond, target, next)
	f.cur = next
	return nil
}
