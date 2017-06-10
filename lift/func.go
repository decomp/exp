package lift

import (
	"fmt"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/decomp/exp/disasm/x86"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"golang.org/x/arch/x86/x86asm"
)

// A Func is a function lifter.
type Func struct {
	// Output LLVM IR of the function.
	*ir.Function
	// Input assembly of the function.
	AsmFunc *x86.Func
	// Current basic block being generated.
	cur *ir.BasicBlock
	// LLVM IR basic blocks of the function.
	blocks map[bin.Address]*ir.BasicBlock
	// Registers used within the function.
	regs map[x86asm.Reg]*ir.InstAlloca
	// Status flags used within the function.
	statusFlags map[StatusFlag]*ir.InstAlloca
	// Local varialbes used within the function.
	locals map[string]*ir.InstAlloca
	// usesEDX_EAX specifies whether any instruction of the function uses
	// EDX:EAX.
	usesEDX_EAX bool
	// usesFPU specifies whether any instruction of the function uses the FPU.
	usesFPU bool

	// TODO: Move espDisp from Func to BasicBlock, and propagate symbolic
	// execution information through context.json.

	// ESP disposition; used for shadow stack.
	espDisp int64

	// FPU register stack top; integer value in range [0, 7].
	st *ir.InstAlloca

	// Read-only global lifter state.
	l *Lifter
}

// NewFunc returns a new function lifter based on the input assembly of the
// function.
func (l *Lifter) NewFunc(asmFunc *x86.Func) *Func {
	entry := asmFunc.Addr
	f, ok := l.Funcs[entry]
	if !ok {
		// TODO: Add proper support for type signatures once type analysis has
		// been conducted.
		name := fmt.Sprintf("f_%06X", uint64(entry))
		sig := types.NewFunc(types.Void)
		typ := types.NewPointer(sig)
		f = &Func{
			Function: &ir.Function{
				Name: name,
				Typ:  typ,
				Sig:  sig,
				Metadata: map[string]*metadata.Metadata{
					"addr": {
						Nodes: []metadata.Node{&metadata.String{Val: entry.String()}},
					},
				},
			},
		}
	}
	f.AsmFunc = asmFunc
	f.blocks = make(map[bin.Address]*ir.BasicBlock)
	f.regs = make(map[x86asm.Reg]*ir.InstAlloca)
	f.statusFlags = make(map[StatusFlag]*ir.InstAlloca)
	f.locals = make(map[string]*ir.InstAlloca)
	f.l = l
	// Prepare output LLVM IR basic blocks.
	for addr := range asmFunc.Blocks {
		label := fmt.Sprintf("block_%06X", uint64(addr))
		block := &ir.BasicBlock{
			Name: label,
		}
		f.blocks[addr] = block
	}
	// Preprocess the function to assess if any instruction makes use of EDX:EAX
	// (e.g. IDIV).
	for _, bb := range asmFunc.Blocks {
		for _, inst := range bb.Insts {
			switch inst.Op {
			// TODO: Identify more instructions which makes use of the FPU register
			// stack.
			case x86asm.F2XM1, x86asm.FABS, x86asm.FADD, x86asm.FADDP, x86asm.FBLD,
				x86asm.FBSTP, x86asm.FCHS, x86asm.FCMOVB, x86asm.FCMOVBE,
				x86asm.FCMOVE, x86asm.FCMOVNB, x86asm.FCMOVNBE, x86asm.FCMOVNE,
				x86asm.FCMOVNU, x86asm.FCMOVU, x86asm.FCOM, x86asm.FCOMI,
				x86asm.FCOMIP, x86asm.FCOMP, x86asm.FCOMPP, x86asm.FCOS,
				x86asm.FDECSTP, x86asm.FDIV, x86asm.FDIVP, x86asm.FDIVR, x86asm.FDIVRP,
				x86asm.FFREE, x86asm.FFREEP, x86asm.FIADD, x86asm.FICOM, x86asm.FICOMP,
				x86asm.FIDIV, x86asm.FIDIVR, x86asm.FILD, x86asm.FIMUL, x86asm.FINCSTP,
				x86asm.FIST, x86asm.FISTP, x86asm.FISTTP, x86asm.FISUB, x86asm.FISUBR,
				x86asm.FLD, x86asm.FLD1, x86asm.FLDCW, x86asm.FLDENV, x86asm.FLDL2E,
				x86asm.FLDL2T, x86asm.FLDLG2, x86asm.FLDPI, x86asm.FMUL, x86asm.FMULP,
				x86asm.FNCLEX, x86asm.FNINIT, x86asm.FNOP, x86asm.FNSAVE,
				x86asm.FNSTCW, x86asm.FNSTENV, x86asm.FNSTSW, x86asm.FPATAN,
				x86asm.FPREM, x86asm.FPREM1, x86asm.FPTAN, x86asm.FRNDINT,
				x86asm.FRSTOR, x86asm.FSCALE, x86asm.FSIN, x86asm.FSINCOS,
				x86asm.FSQRT, x86asm.FST, x86asm.FSTP, x86asm.FSUB, x86asm.FSUBP,
				x86asm.FSUBR, x86asm.FSUBRP, x86asm.FTST, x86asm.FUCOM, x86asm.FUCOMI,
				x86asm.FUCOMIP, x86asm.FUCOMP, x86asm.FUCOMPP, x86asm.FWAIT,
				x86asm.FXAM, x86asm.FXCH, x86asm.FXRSTOR, x86asm.FXRSTOR64,
				x86asm.FXSAVE, x86asm.FXSAVE64, x86asm.FXTRACT, x86asm.FYL2X,
				x86asm.FYL2XP1:
				f.usesFPU = true
			// TODO: Identify more instructions which makes use of EDX:EAX.
			case x86asm.IDIV:
				f.usesEDX_EAX = true
			}
		}
	}
	return f
}

// Lift lifts the function from input assembly to LLVM IR.
func (f *Func) Lift() {
	dbg.Printf("lifting function at %v", f.AsmFunc.Addr)
	// Allocate a local variable for the FPU stack top used within the function.
	if f.usesFPU {
		v := ir.NewAlloca(types.I8)
		v.SetName("st")
		f.st = v
	}
	var blockAddrs bin.Addresses
	for blockAddr := range f.AsmFunc.Blocks {
		blockAddrs = append(blockAddrs, blockAddr)
	}
	sort.Sort(blockAddrs)
	if len(blockAddrs) == 0 {
		panic(fmt.Errorf("invalid function definition at %v; missing function body", f.AsmFunc.Addr))
	}
	for _, blockAddr := range blockAddrs {
		bb := f.AsmFunc.Blocks[blockAddr]
		f.liftBlock(bb)
	}
	// Add new entry basic block to define registers and status flags used within
	// the function.
	if len(f.regs) > 0 || len(f.statusFlags) > 0 || f.usesFPU {
		entry := &ir.BasicBlock{}
		// Allocate local variables for each register used within the function.
		for reg := x86.FirstReg; reg <= x86.LastReg; reg++ {
			if inst, ok := f.regs[reg]; ok {
				entry.AppendInst(inst)
			}
		}
		// Allocate local variables for the FPU register stack used within the
		// function.
		if f.usesFPU {
			entry.AppendInst(f.st)
			seven := constant.NewInt(7, types.I8)
			entry.NewStore(seven, f.st)
		}
		// Allocate local variables for each status flag used within the function.
		for status := CF; status <= OF; status++ {
			if inst, ok := f.statusFlags[status]; ok {
				entry.AppendInst(inst)
			}
		}
		// Allocate local variables for each local variable used within the
		// function.
		var names []string
		for name := range f.locals {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			inst := f.locals[name]
			entry.AppendInst(inst)
		}
		// Handle calling conventions.
		f.cur = entry
		// TODO: Initialize parameter initialization in entry block prior to basic
		// block translation. Move this code to before f.translateBlock, and remove
		// f.espDisp = 0.
		f.espDisp = 0
		for i, param := range f.Sig.Params {
			// Use parameter in register.
			switch f.CallConv {
			case ir.CallConvX86_FastCall:
				switch i {
				case 0:
					f.defReg(x86.ECX, param)
					continue
				case 1:
					f.defReg(x86.EDX, param)
					continue
				}
			default:
				// TODO: Add support for more calling conventions.
			}
			// Use parameter on stack.
			m := x86asm.Mem{
				Base: x86asm.ESP,
				Disp: 4,
			}
			mem := x86.NewMem(m, nil)
			f.defMem(mem, param)
		}
		target := f.Blocks[0]
		entry.NewBr(target)
		f.Blocks = append([]*ir.BasicBlock{entry}, f.Blocks...)
	}
}

// liftBlock lifts the basic block from input assembly to LLVM IR.
func (f *Func) liftBlock(bb *x86.BasicBlock) {
	dbg.Printf("lifting basic block at %v", bb.Addr)
	f.cur = f.blocks[bb.Addr]
	f.Blocks = append(f.Blocks, f.cur)
	for _, inst := range bb.Insts {
		f.liftInst(inst)
	}
	f.liftTerm(bb.Term)
}
