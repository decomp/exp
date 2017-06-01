package lift

import (
	"sort"

	"golang.org/x/arch/x86/x86asm"

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
	// LLVM IR basic blocks of the function.
	blocks map[bin.Address]*ir.BasicBlock
	// Current basic block being generated.
	cur *ir.BasicBlock
	// Registers used within the function.
	regs map[x86asm.Reg]*ir.InstAlloca
	// Status flags used within the function.
	statusFlags map[StatusFlag]*ir.InstAlloca
	// Local varialbes used within the function.
	locals map[string]*ir.InstAlloca
	// usesEDX_EAX specifies whether any instruction of the function uses
	// EDX:EAX.
	usesEDX_EAX bool

	// TODO: Move espDisp from Func to BasicBlock, and propagate symbolic
	// execution information through context.json.

	// ESP disposition; used for shadow stack.
	espDisp int64

	// Read-only global lifter state.
	l *Lifter
}

// NewFunc returns a new function lifter based on the input assembly of the
// function.
func (l *Lifter) NewFunc(asmFunc *x86.Func) *Func {
	f, ok := l.Funcs[asmFunc.Addr]
	if !ok {
		f.Function = &ir.Function{}
	}
	f.AsmFunc = asmFunc
	f.cur = &ir.BasicBlock{}
	f.Blocks = append(f.Blocks, f.cur)
	f.regs = make(map[x86asm.Reg]*ir.InstAlloca)
	f.statusFlags = make(map[StatusFlag]*ir.InstAlloca)
	f.locals = make(map[string]*ir.InstAlloca)
	f.l = l
	return f
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
