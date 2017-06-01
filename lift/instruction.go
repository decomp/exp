package lift

import (
	"fmt"

	"golang.org/x/arch/x86/x86asm"

	"github.com/decomp/exp/disasm/x86"
)

// liftInst lifts the instruction from input assembly to LLVM IR.
func (f *Func) liftInst(inst *x86.Inst) {
	dbg.Printf("lifting instruction at %v", inst.Addr)
	switch inst.Op {
	case x86asm.AND:
		f.liftInstAND(inst)
	case x86asm.MOV:
		f.liftInstMOV(inst)
	default:
		panic(fmt.Errorf("support for x86 instruction opcode %v not yet implemented", inst.Op))
	}
}

// --- [ AND ] -----------------------------------------------------------------

// liftInstAND lifts the given x86 AND instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstAND(inst *x86.Inst) {
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewAnd(x, y)
	f.defArg(inst.Arg(0), result)
}

// --- [ MOV ] -----------------------------------------------------------------

// liftInstMOV lifts the given x86 MOV instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstMOV(inst *x86.Inst) {
	src := f.useArg(inst.Arg(1))
	f.defArg(inst.Arg(0), src)
}
