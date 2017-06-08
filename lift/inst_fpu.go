package lift

import (
	"github.com/decomp/exp/disasm/x86"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// ref: $ 5.2 X87 FPU INSTRUCTIONS, Intel 64 and IA-32 architectures software
// developer's manual volume 1: Basic architecture.

// === [ x87 FPU Data Transfer Instructions ] ==================================

// --- [ FLD ] -----------------------------------------------------------------

// liftInstFLD lifts the given x87 FLD instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstFLD(inst *x86.Inst) error {
	// Load floating-point value.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FST ] -----------------------------------------------------------------

// liftInstFST lifts the given x87 FST instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstFST(inst *x86.Inst) error {
	// Store floating-point value.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FSTP ] ----------------------------------------------------------------

// liftInstFSTP lifts the given x87 FSTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSTP(inst *x86.Inst) error {
	// Store floating-point value and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FILD ] ----------------------------------------------------------------

// liftInstFILD lifts the given x87 FILD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFILD(inst *x86.Inst) error {
	// FILD - Load integer.
	//
	//    FILD arg
	//
	// Converts the signed-integer source operand into double extended-precision
	// floating-point format and pushes the value onto the FPU register stack.
	end := &ir.BasicBlock{}
	arg := f.useArg(inst.Arg(0))
	src := f.cur.NewSIToFP(arg, types.X86_FP80)
	cur := f.cur
	var cases []*ir.Case
	for i, v := range f.fpuStack[:] {
		block := &ir.BasicBlock{}
		block.NewBr(end)
		f.cur = block
		f.AppendBlock(block)
		f.cur.NewStore(src, v)
		c := ir.NewCase(constant.NewInt(int64(i), types.I8), block)
		cases = append(cases, c)
	}
	f.cur = cur
	st := f.cur.NewLoad(f.st)
	defaultTarget := &ir.BasicBlock{}
	defaultTarget.NewUnreachable()
	f.AppendBlock(defaultTarget)
	f.cur.NewSwitch(st, defaultTarget, cases...)
	f.cur = end
	f.AppendBlock(end)
	return nil
}

// --- [ FIST ] ----------------------------------------------------------------

// liftInstFIST lifts the given x87 FIST instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFIST(inst *x86.Inst) error {
	// Store integer.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FISTP ] ---------------------------------------------------------------

// liftInstFISTP lifts the given x87 FISTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFISTP(inst *x86.Inst) error {
	// Store integer and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FBLD ] ----------------------------------------------------------------

// liftInstFBLD lifts the given x87 FBLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFBLD(inst *x86.Inst) error {
	// Load BCD.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FBSTP ] ---------------------------------------------------------------

// liftInstFBSTP lifts the given x87 FBSTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFBSTP(inst *x86.Inst) error {
	// Store BCD and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FXCH ] ----------------------------------------------------------------

// liftInstFXCH lifts the given x87 FXCH instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFXCH(inst *x86.Inst) error {
	// Exchange registers.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FCMOVE ] --------------------------------------------------------------

// liftInstFCMOVE lifts the given x87 FCMOVE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVE(inst *x86.Inst) error {
	// Floating-point conditional move if equal.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FCMOVNE ] -------------------------------------------------------------

// liftInstFCMOVNE lifts the given x87 FCMOVNE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVNE(inst *x86.Inst) error {
	// Floating-point conditional move if not equal.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FCMOVB ] --------------------------------------------------------------

// liftInstFCMOVB lifts the given x87 FCMOVB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVB(inst *x86.Inst) error {
	// Floating-point conditional move if below.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FCMOVBE ] -------------------------------------------------------------

// liftInstFCMOVBE lifts the given x87 FCMOVBE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVBE(inst *x86.Inst) error {
	// Floating-point conditional move if below or equal.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FCMOVNB ] -------------------------------------------------------------

// liftInstFCMOVNB lifts the given x87 FCMOVNB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVNB(inst *x86.Inst) error {
	// Floating-point conditional move if not below.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FCMOVNBE ] ------------------------------------------------------------

// liftInstFCMOVNBE lifts the given x87 FCMOVNBE instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstFCMOVNBE(inst *x86.Inst) error {
	// Floating-point conditional move if not below or equal.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FCMOVU ] --------------------------------------------------------------

// liftInstFCMOVU lifts the given x87 FCMOVU instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVU(inst *x86.Inst) error {
	// Floating-point conditional move if unordered.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FCMOVNU ] -------------------------------------------------------------

// liftInstFCMOVNU lifts the given x87 FCMOVNU instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVNU(inst *x86.Inst) error {
	// Floating-point conditional move if not unordered.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// === [ x87 FPU Basic Arithmetic Instructions ] ===============================

// === [ x87 FPU Comparison Instructions ] =====================================

// === [ x87 FPU Transcendental Instructions ] =================================

// === [ x87 FPU Load Constants Instructions ] =================================

// === [ x87 FPU Control Instructions ] ========================================
