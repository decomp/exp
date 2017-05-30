package lift

import (
	"fmt"

	"github.com/decomp/exp/disasm/x86"
)

// liftTerm lifts the terminator from input assembly to LLVM IR.
func (f *Func) liftTerm(term *x86.Inst) {
	dbg.Printf("lifting terminator at %v", term.Addr)

	// Lift terminator.
	switch term.Op {
	/*
		// Loop terminators.
		case x86asm.LOOP:
			return f.liftTermLOOP(term)
		case x86asm.LOOPE:
			return f.liftTermLOOPE(term)
		case x86asm.LOOPNE:
			return f.liftTermLOOPNE(term)
		// Conditional jump terminators.
		case x86asm.JA:
			return f.liftTermJA(term)
		case x86asm.JAE:
			return f.liftTermJAE(term)
		case x86asm.JB:
			return f.liftTermJB(term)
		case x86asm.JBE:
			return f.liftTermJBE(term)
		case x86asm.JCXZ:
			return f.liftTermJCXZ(term)
		case x86asm.JE:
			return f.liftTermJE(term)
		case x86asm.JECXZ:
			return f.liftTermJECXZ(term)
		case x86asm.JG:
			return f.liftTermJG(term)
		case x86asm.JGE:
			return f.liftTermJGE(term)
		case x86asm.JL:
			return f.liftTermJL(term)
		case x86asm.JLE:
			return f.liftTermJLE(term)
		case x86asm.JNE:
			return f.liftTermJNE(term)
		case x86asm.JNO:
			return f.liftTermJNO(term)
		case x86asm.JNP:
			return f.liftTermJNP(term)
		case x86asm.JNS:
			return f.liftTermJNS(term)
		case x86asm.JO:
			return f.liftTermJO(term)
		case x86asm.JP:
			return f.liftTermJP(term)
		case x86asm.JRCXZ:
			return f.liftTermJRCXZ(term)
		case x86asm.JS:
			return f.liftTermJS(term)
		// Unconditional jump terminators.
		case x86asm.JMP:
			return f.liftTermJMP(term)
		// Return terminators.
		case x86asm.RET:
			return f.liftTermRET(term)
	*/
	default:
		panic(fmt.Errorf("support for x86 terminator opcode %v not yet implemented", term.Op))
	}
}
