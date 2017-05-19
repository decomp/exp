package lift

import (
	"github.com/decomp/exp/disasm/x86"
	"github.com/llir/llvm/ir"
)

// A Func is a function.
type Func struct {
	// Output LLVM IR of the function.
	*ir.Function
	// Input assembly of the function.
	Old *x86.Func
	// Read-only global lifter state.
	l *Lifter
}
