package lift

import (
	"github.com/decomp/exp/disasm/x86"
	"github.com/llir/llvm/ir"
)

type Func struct {
	*ir.Function
	Old *x86.Func
}
