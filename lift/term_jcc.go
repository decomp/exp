//+build ignore

package lift

import (
	"github.com/decomp/exp/bin"
	"github.com/decomp/exp/disasm/x86"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
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

// liftTermJA lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJA(term *x86.Inst) error {
	// Jump if above.
	//    (CF=0 and ZF=0)
	cf := f.useStatus(CF)
	zf := f.useStatus(ZF)
	cond1 := f.cur.NewICmp(ir.IntEQ, cf, constant.False)
	cond2 := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	cond := f.cur.NewAnd(cond1, cond2)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JAE ] -----------------------------------------------------------------

// liftTermJAE lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJAE(term *x86.Inst) error {
	// Jump if above or equal.
	//    (CF=0)
	cf := f.useStatus(CF)
	cond := f.cur.NewICmp(ir.IntEQ, cf, constant.False)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JBE ] -----------------------------------------------------------------

// liftTermJBE lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJBE(term *x86.Inst) error {
	// Jump if below or equal.
	//    (CF=1 or ZF=1)
	cf := f.useStatus(CF)
	zf := f.useStatus(ZF)
	cond1 := cf
	cond2 := zf
	cond := f.cur.NewOr(cond1, cond2)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JB ] ------------------------------------------------------------------

// liftTermJB lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJB(term *x86.Inst) error {
	// Jump if below.
	//    (CF=1)
	cf := f.useStatus(CF)
	cond := cf
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JCXZ ] ----------------------------------------------------------------

// liftTermJCXZ lifts the given x86 JA terminator to LLVM IR, emitting code to
// f.
func (f *Func) liftTermJCXZ(term *x86.Inst) error {
	// Jump if CX register is zero.
	//    (CX=0)
	pretty.Println("term:", term)
	panic("emitTermJCXZ: not yet implemented")
}

// --- [ JECXZ ] ---------------------------------------------------------------

// liftTermJECXZ lifts the given x86 JA terminator to LLVM IR, emitting code to
// f.
func (f *Func) liftTermJECXZ(term *x86.Inst) error {
	// Jump if ECX register is zero.
	//    (ECX=0)
	ecx := f.useReg(ECX)
	zero := constant.NewInt(0, types.I32)
	cond := f.cur.NewICmp(ir.IntEQ, ecx, zero)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JNO ] -----------------------------------------------------------------

// liftTermJNO lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJNO(term *x86.Inst) error {
	// Jump if not overflow.
	//    (OF=0)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(ir.IntEQ, of, constant.False)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JO ] ------------------------------------------------------------------

// liftTermJO lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJO(term *x86.Inst) error {
	// Jump if overflow.
	//    (OF=1)
	of := f.useStatus(OF)
	cond := of
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JNP ] -----------------------------------------------------------------

// liftTermJNP lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJNP(term *x86.Inst) error {
	// Jump if not parity.
	//    (PF=0)
	pf := f.useStatus(PF)
	cond := f.cur.NewICmp(ir.IntEQ, pf, constant.False)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JP ] ------------------------------------------------------------------

// liftTermJP lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJP(term *x86.Inst) error {
	// Jump if parity.
	//    (PF=1)
	pf := f.useStatus(PF)
	cond := pf
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JRCXZ ] ---------------------------------------------------------------

// liftTermJRCXZ lifts the given x86 JA terminator to LLVM IR, emitting code to
// f.
func (f *Func) liftTermJRCXZ(term *x86.Inst) error {
	// Jump if RCX register is zero.
	//    (RCX=0)
	pretty.Println("term:", term)
	panic("emitTermJRCXZ: not yet implemented")
}

// --- [ JNS ] -----------------------------------------------------------------

// liftTermJNS lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJNS(term *x86.Inst) error {
	// Jump if not sign.
	//    (SF=0)
	sf := f.useStatus(SF)
	cond := f.cur.NewICmp(ir.IntEQ, sf, constant.False)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JS ] ------------------------------------------------------------------

// liftTermJS lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJS(term *x86.Inst) error {
	// Jump if sign.
	//    (SF=1)
	sf := f.useStatus(SF)
	cond := sf
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JGE ] -----------------------------------------------------------------

// liftTermJGE lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJGE(term *x86.Inst) error {
	// Jump if greater or equal.
	//    (SF=OF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(ir.IntEQ, sf, of)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JL ] ------------------------------------------------------------------

// liftTermJL lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJL(term *x86.Inst) error {
	// Jump if less.
	//    (SF≠OF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond := f.cur.NewICmp(ir.IntNE, sf, of)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JG ] ------------------------------------------------------------------

// liftTermJG lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJG(term *x86.Inst) error {
	// Jump if greater.
	//    (ZF=0 and SF=OF)
	zf := f.useStatus(ZF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond1 := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	cond2 := f.cur.NewICmp(ir.IntEQ, sf, of)
	cond := f.cur.NewAnd(cond1, cond2)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JNE ] -----------------------------------------------------------------

// liftTermJNE lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJNE(term *x86.Inst) error {
	// Jump if not equal.
	//    (ZF=0)
	zf := f.useStatus(ZF)
	cond := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JLE ] -----------------------------------------------------------------

// liftTermJLE lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJLE(term *x86.Inst) error {
	// Jump if less or equal.
	//    (ZF=1 or SF≠OF)
	zf := f.useStatus(ZF)
	sf := f.useStatus(SF)
	of := f.useStatus(OF)
	cond1 := zf
	cond2 := f.cur.NewICmp(ir.IntNE, sf, of)
	cond := f.cur.NewOr(cond1, cond2)
	return f.liftTermJcc(term.Arg(0), cond)
}

// --- [ JE ] ------------------------------------------------------------------

// liftTermJE lifts the given x86 JA terminator to LLVM IR, emitting code to f.
func (f *Func) liftTermJE(term *x86.Inst) error {
	// Jump if equal.
	//    (ZF=1)
	zf := f.useStatus(ZF)
	cond := zf
	return f.liftTermJcc(term.Arg(0), cond)
}

// === [ Helper functions ] ====================================================

// liftTermJcc lifts the given x86 SETcc instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftTermJcc(arg *Arg, cond value.Value) error {
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
