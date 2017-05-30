package lift

import (
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/decomp/exp/disasm/x86"
	"github.com/llir/llvm/ir"
)

// A Func is a function lifter.
type Func struct {
	// Output LLVM IR of the function.
	*ir.Function
	// Input assembly of the function.
	AsmFunc *x86.Func
	// Read-only global lifter state.
	l *Lifter
}

// NewFunc returns a new function lifter based on the input assembly of the
// function.
func (l *Lifter) NewFunc(asmFunc *x86.Func) *Func {
	return &Func{
		Function: &ir.Function{},
		AsmFunc:  asmFunc,
		l:        l,
	}
}

// Lift lifts the function from input assembly to LLVM IR.
func (f *Func) Lift() {
	dbg.Printf("lifting function at %v", f.AsmFunc.Addr)
	var blockAddrs bin.Addresses
	for blockAddr := range f.AsmFunc.Blocks {
		blockAddrs = append(blockAddrs, blockAddr)
	}
	sort.Sort(blockAddrs)
	for _, blockAddr := range blockAddrs {
		block := f.AsmFunc.Blocks[blockAddr]
		f.liftBlock(block)
	}
}

// liftBlock lifts the basic block from input assembly to LLVM IR.
func (f *Func) liftBlock(block *x86.BasicBlock) {
	dbg.Printf("lifting basic block at %v", block.Addr)
	for _, inst := range block.Insts {
		f.liftInst(inst)
	}
	f.liftTerm(block.Term)
}
