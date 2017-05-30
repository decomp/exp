package lift

import (
	"fmt"

	"github.com/decomp/exp/disasm/x86"
)

// liftInst lifts the instruction from input assembly to LLVM IR.
func (f *Func) liftInst(inst *x86.Inst) {
	dbg.Printf("lifting instruction at %v", inst.Addr)
	switch inst.Op {
	default:
		panic(fmt.Errorf("support for x86 instruction opcode %v not yet implemented", inst.Op))
	}
}
