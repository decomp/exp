package lift

import (
	"fmt"

	"github.com/decomp/exp/disasm/x86"
)

// liftInst lifts the instruction from input assembly to LLVM IR.
func (f *Func) liftInst(inst *x86.Inst) {
	switch inst.Op {
	default:
		panic(fmt.Errorf("support for instruction %v not yet implemented", inst.Op))
	}
}
