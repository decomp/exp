package lift

import (
	"fmt"

	"github.com/decomp/exp/disasm/x86"
)

// liftTerm lifts the terminator from input assembly to LLVM IR.
func (f *Func) liftTerm(term *x86.Inst) {
	switch term.Op {
	default:
		panic(fmt.Errorf("support for terminator %v not yet implemented", term.Op))
	}
}
