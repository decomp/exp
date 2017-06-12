// x87 FPU Instructions.
//
// ref: $ 5.2 X87 FPU INSTRUCTIONS, Intel 64 and IA-32 architectures software
// developer's manual volume 1: Basic architecture.

package lift

import (
	"fmt"
	"math"

	"github.com/decomp/exp/disasm/x86"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"golang.org/x/arch/x86/x86asm"
)

// === [ x87 FPU Data Transfer Instructions ] ==================================

// --- [ FLD ] -----------------------------------------------------------------

// liftInstFLD lifts the given x87 FLD instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstFLD(inst *x86.Inst) error {
	// FLD - Load floating-point value.
	//
	//    FLD m32fp           Push m32fp onto the FPU register stack.
	//    FLD m64fp           Push m64fp onto the FPU register stack.
	//    FLD m80fp           Push m80fp onto the FPU register stack.
	//    FLD ST(i)           Push ST(i) onto the FPU register stack.
	//
	// Pushes the source operand onto the FPU register stack. If the source
	// operand is in single-precision or double-precision floating-point format,
	// it is automatically converted to the double extended-precision floating-
	// point format before being pushed on the stack.
	src := f.useArg(inst.Arg(0))
	// TODO: Verify that FLD ST(i) is handled correctly.
	if !types.Equal(src.Type(), types.X86_FP80) {
		src = f.cur.NewFPExt(src, types.X86_FP80)
	}
	f.fpush(src)
	return nil
}

// --- [ FST ] -----------------------------------------------------------------

// liftInstFST lifts the given x87 FST instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstFST(inst *x86.Inst) error {
	// FST - Store floating-point value.
	//
	//    FST m32fp      Copy ST(0) to m32fp.
	//    FST m64fp      Copy ST(0) to m64fp.
	//    FST ST(i)      Copy ST(0) to ST(i).
	//
	// Copies the value in the ST(0) register to the destination operand.
	src := f.fload()
	switch arg := inst.Args[0].(type) {
	case x86asm.Reg:
		// no type conversion needed.
	case x86asm.Mem:
		var typ types.Type
		switch inst.MemBytes {
		case 4:
			typ = types.Float
		case 8:
			typ = types.Double
		default:
			panic(fmt.Errorf("support for memory argument with byte size %d not yet implemented", inst.MemBytes))
		}
		src = f.cur.NewFPTrunc(src, typ)
	default:
		panic(fmt.Errorf("support for operand type %T not yet implemented", arg))
	}
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ FSTP ] ----------------------------------------------------------------

// liftInstFSTP lifts the given x87 FSTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSTP(inst *x86.Inst) error {
	// FSTP - Store floating-point value and pop.
	//
	//    FSTP m32fp          Copy ST(0) to m32fp and pop register stack.
	//    FSTP m64fp          Copy ST(0) to m64fp and pop register stack.
	//    FSTP m80fp          Copy ST(0) to m80fp and pop register stack.
	//    FSTP ST(i)          Copy ST(0) to ST(i) and pop register stack.
	//
	// Copies the value in the ST(0) register to the destination operand.
	src := f.fload()
	switch arg := inst.Args[0].(type) {
	case x86asm.Reg:
		// no type conversion needed.
	case x86asm.Mem:
		switch inst.MemBytes {
		case 4:
			src = f.cur.NewFPTrunc(src, types.Float)
		case 8:
			src = f.cur.NewFPTrunc(src, types.Double)
		case 10:
			// no type conversion needed.
		default:
			panic(fmt.Errorf("support for memory argument with byte size %d not yet implemented", inst.MemBytes))
		}
	default:
		panic(fmt.Errorf("support for operand type %T not yet implemented", arg))
	}
	f.defArg(inst.Arg(0), src)
	f.pop()
	return nil
}

// --- [ FILD ] ----------------------------------------------------------------

// liftInstFILD lifts the given x87 FILD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFILD(inst *x86.Inst) error {
	// FILD - Load integer.
	//
	//    FILD m16int         Push m16int onto the FPU register stack.
	//    FILD m32int         Push m32int onto the FPU register stack.
	//    FILD m64int         Push m64int onto the FPU register stack.
	//
	// Converts the signed-integer source operand into double extended-precision
	// floating-point format and pushes the value onto the FPU register stack.
	arg := f.useArg(inst.Arg(0))
	src := f.cur.NewSIToFP(arg, types.X86_FP80)
	f.fpush(src)
	return nil
}

// --- [ FIST ] ----------------------------------------------------------------

// liftInstFIST lifts the given x87 FIST instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFIST(inst *x86.Inst) error {
	// FIST - Store integer.
	pretty.Println("inst:", inst)
	panic("emitInstFIST: not yet implemented")
}

// --- [ FISTP ] ---------------------------------------------------------------

// liftInstFISTP lifts the given x87 FISTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFISTP(inst *x86.Inst) error {
	// FISTP - Store integer and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFISTP: not yet implemented")
}

// --- [ FBLD ] ----------------------------------------------------------------

// liftInstFBLD lifts the given x87 FBLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFBLD(inst *x86.Inst) error {
	// FBLD - Load BCD.
	pretty.Println("inst:", inst)
	panic("emitInstFBLD: not yet implemented")
}

// --- [ FBSTP ] ---------------------------------------------------------------

// liftInstFBSTP lifts the given x87 FBSTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFBSTP(inst *x86.Inst) error {
	// FBSTP - Store BCD and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFBSTP: not yet implemented")
}

// --- [ FXCH ] ----------------------------------------------------------------

// liftInstFXCH lifts the given x87 FXCH instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFXCH(inst *x86.Inst) error {
	// FXCH - Exchange registers.
	pretty.Println("inst:", inst)
	panic("emitInstFXCH: not yet implemented")
}

// ___ [ FCMOVcc - Floating-Point Conditional Move Instructions ] ______________
//
//    Instruction Mnemonic   Status Flag States   Condition Description
//
//    FCMOVB                 CF=1                 Below
//    FCMOVNB                CF=0                 Not below
//    FCMOVE                 ZF=1                 Equal
//    FCMOVNE                ZF=0                 Not equal
//    FCMOVBE                CF=1 or ZF=1         Below or equal
//    FCMOVNBE               CF=0 or ZF=0         Not below nor equal
//    FCMOVU                 PF=1                 Unordered
//    FCMOVNU                PF=0                 Not unordered
//
// ref: $ 8.3.3 Data Transfer Instructions, Table 8-5, Floating-Point
// Conditional Move Instructions, Intel 64 and IA-32 Architectures Software
// Developer's Manual: Basic architecture.

// --- [ FCMOVE ] --------------------------------------------------------------

// liftInstFCMOVE lifts the given x87 FCMOVE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVE(inst *x86.Inst) error {
	// FCMOVE - Floating-point conditional move if equal.
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVE: not yet implemented")
}

// --- [ FCMOVNE ] -------------------------------------------------------------

// liftInstFCMOVNE lifts the given x87 FCMOVNE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVNE(inst *x86.Inst) error {
	// FCMOVNE - Floating-point conditional move if not equal.
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVNE: not yet implemented")
}

// --- [ FCMOVB ] --------------------------------------------------------------

// liftInstFCMOVB lifts the given x87 FCMOVB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVB(inst *x86.Inst) error {
	// FCMOVB - Floating-point conditional move if below.
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVB: not yet implemented")
}

// --- [ FCMOVBE ] -------------------------------------------------------------

// liftInstFCMOVBE lifts the given x87 FCMOVBE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVBE(inst *x86.Inst) error {
	// FCMOVBE - Floating-point conditional move if below or equal.
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVBE: not yet implemented")
}

// --- [ FCMOVNB ] -------------------------------------------------------------

// liftInstFCMOVNB lifts the given x87 FCMOVNB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVNB(inst *x86.Inst) error {
	// FCMOVNB - Floating-point conditional move if not below.
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVNB: not yet implemented")
}

// --- [ FCMOVNBE ] ------------------------------------------------------------

// liftInstFCMOVNBE lifts the given x87 FCMOVNBE instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstFCMOVNBE(inst *x86.Inst) error {
	// FCMOVNBE - Floating-point conditional move if not below or equal.
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVNBE: not yet implemented")
}

// --- [ FCMOVU ] --------------------------------------------------------------

// liftInstFCMOVU lifts the given x87 FCMOVU instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVU(inst *x86.Inst) error {
	// FCMOVU - Floating-point conditional move if unordered.
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVU: not yet implemented")
}

// --- [ FCMOVNU ] -------------------------------------------------------------

// liftInstFCMOVNU lifts the given x87 FCMOVNU instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVNU(inst *x86.Inst) error {
	// FCMOVNU - Floating-point conditional move if not unordered.
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVNU: not yet implemented")
}

// === [ x87 FPU Basic Arithmetic Instructions ] ===============================

// --- [ FADD ] ----------------------------------------------------------------

// liftInstFADD lifts the given x87 FADD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFADD(inst *x86.Inst) error {
	// FADD - Add floating-point.
	//
	//    FADD m32fp          Add m32fp to ST(0) and store result in ST(0).
	//    FADD m64fp          Add m64fp to ST(0) and store result in ST(0).
	//    FADD ST(0), ST(i)   Add ST(0) to ST(i) and store result in ST(0).
	//    FADD ST(i), ST(0)   Add ST(i) to ST(0) and store result in ST(i).
	//
	// Adds the destination and source operands and stores the sum in the
	// destination location.
	if inst.Args[1] != nil {
		panic(fmt.Errorf("support for two-operand FADD instruction not yet implemented; instruction %v at address %v", inst, inst.Addr))
	}
	src := f.useArg(inst.Arg(0))
	v := f.cur.NewFPExt(src, types.X86_FP80)
	st0 := f.fload()
	result := f.cur.NewFAdd(st0, v)
	f.fstore(result)
	return nil
}

// --- [ FADDP ] ---------------------------------------------------------------

// liftInstFADDP lifts the given x87 FADDP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFADDP(inst *x86.Inst) error {
	// FADDP - Add floating-point and pop.
	//
	//    FADDP ST(i), ST(0)            Add ST(0) to ST(i), store result in ST(i), and pop the register stack.
	//    FADDP                         Add ST(0) to ST(1), store result in ST(1), and pop the register stack.
	//
	// Adds the destination and source operands and stores the sum in the
	// destination location.
	pretty.Println("inst:", inst)
	panic("emitInstFADDP: not yet implemented")
}

// --- [ FIADD ] ---------------------------------------------------------------

// liftInstFIADD lifts the given x87 FIADD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFIADD(inst *x86.Inst) error {
	// FIADD - Add integer.
	//
	//    FIADD m32int        Add m32int to ST(0) and store result in ST(0).
	//    FIADD m16int        Add m16int to ST(0) and store result in ST(0).
	//
	// Adds the destination and source operands and stores the sum in the
	// destination location.
	pretty.Println("inst:", inst)
	panic("emitInstFIADD: not yet implemented")
}

// --- [ FSUB ] ----------------------------------------------------------------

// liftInstFSUB lifts the given x87 FSUB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSUB(inst *x86.Inst) error {
	// FSUB - Subtract floating-point.
	pretty.Println("inst:", inst)
	panic("emitInstFSUB: not yet implemented")
}

// --- [ FSUBP ] ---------------------------------------------------------------

// liftInstFSUBP lifts the given x87 FSUBP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSUBP(inst *x86.Inst) error {
	// FSUBP - Subtract floating-point and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFSUBP: not yet implemented")
}

// --- [ FISUB ] ---------------------------------------------------------------

// liftInstFISUB lifts the given x87 FISUB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFISUB(inst *x86.Inst) error {
	// FISUB - Subtract integer.
	pretty.Println("inst:", inst)
	panic("emitInstFISUB: not yet implemented")
}

// --- [ FSUBR ] ---------------------------------------------------------------

// liftInstFSUBR lifts the given x87 FSUBR instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSUBR(inst *x86.Inst) error {
	// FSUBR - Subtract floating-point reverse.
	pretty.Println("inst:", inst)
	panic("emitInstFSUBR: not yet implemented")
}

// --- [ FSUBRP ] --------------------------------------------------------------

// liftInstFSUBRP lifts the given x87 FSUBRP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFSUBRP(inst *x86.Inst) error {
	// FSUBRP - Subtract floating-point reverse and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFSUBRP: not yet implemented")
}

// --- [ FISUBR ] --------------------------------------------------------------

// liftInstFISUBR lifts the given x87 FISUBR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFISUBR(inst *x86.Inst) error {
	// FISUBR - Subtract integer reverse.
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FMUL ] ----------------------------------------------------------------

// liftInstFMUL lifts the given x87 FMUL instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFMUL(inst *x86.Inst) error {
	// FMUL - Multiply floating-point.
	//
	//    FMUL m32fp          Multiply ST(0) by m32fp and store result in ST(0).
	//    FMUL m64fp          Multiply ST(0) by m64fp and store result in ST(0).
	//    FMUL ST(0), ST(i)   Multiply ST(0) by ST(i) and store result in ST(0).
	//    FMUL ST(i), ST(0)   Multiply ST(i) by ST(0) and store result in ST(i).
	//
	// Multiplies the destination and source operands and stores the product in
	// the destination location.
	if inst.Args[1] != nil {
		panic(fmt.Errorf("support for two-operand form of FMUL not yet implemented; %v", inst))
	}
	arg := f.useArg(inst.Arg(0))
	src := f.cur.NewFPExt(arg, types.X86_FP80)
	st0 := f.fload()
	result := f.cur.NewFMul(st0, src)
	f.fstore(result)
	return nil
}

// --- [ FMULP ] ---------------------------------------------------------------

// liftInstFMULP lifts the given x87 FMULP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFMULP(inst *x86.Inst) error {
	// FMULP - Multiply floating-point and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFMULP: not yet implemented")
}

// --- [ FIMUL ] ---------------------------------------------------------------

// liftInstFIMUL lifts the given x87 FIMUL instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFIMUL(inst *x86.Inst) error {
	// FIMUL - Multiply integer.
	pretty.Println("inst:", inst)
	panic("emitInstFIMUL: not yet implemented")
}

// --- [ FDIV ] ----------------------------------------------------------------

// liftInstFDIV lifts the given x87 FDIV instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFDIV(inst *x86.Inst) error {
	// FDIV - Divide floating-point.
	pretty.Println("inst:", inst)
	panic("emitInstFDIV: not yet implemented")
}

// --- [ FDIVP ] ---------------------------------------------------------------

// liftInstFDIVP lifts the given x87 FDIVP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFDIVP(inst *x86.Inst) error {
	// FDIVP - Divide floating-point and pop.
	//
	//    FDIVP ST(i), ST(0)            Divide ST(i) by ST(0), store result in ST(i), and pop the register stack.
	//    FDIVP                         Divide ST(1) by ST(0), store result in ST(1), and pop the register stack.
	//
	// Convert an integer source operand to double extended-precision floating-
	// point format before performing the division.
	if inst.Args[0] == nil {
		panic(fmt.Errorf("support for zero-operand FDIVP instruction not yet implemented; instruction %v at address %v", inst, inst.Addr))
	}
	dst := f.useArg(inst.Arg(0))
	src := f.useArg(inst.Arg(1))
	result := f.cur.NewFDiv(dst, src)
	f.defArg(inst.Arg(0), result)
	f.fpop()
	return nil
}

// --- [ FIDIV ] ---------------------------------------------------------------

// liftInstFIDIV lifts the given x87 FIDIV instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFIDIV(inst *x86.Inst) error {
	// FIDIV - Divide integer.
	//
	//    FIDIV m16int        Divide ST(0) by m16int and store result in ST(0).
	//    FIDIV m32int        Divide ST(0) by m32int and store result in ST(0).
	//
	// Convert an integer source operand to double extended-precision floating-
	// point format before performing the division.
	arg := f.useArg(inst.Arg(0))
	src := f.cur.NewSIToFP(arg, types.X86_FP80)
	st0 := f.fload()
	result := f.cur.NewFDiv(st0, src)
	f.fstore(result)
	return nil
}

// --- [ FDIVR ] ---------------------------------------------------------------

// liftInstFDIVR lifts the given x87 FDIVR instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFDIVR(inst *x86.Inst) error {
	// FDIVR - Divide floating-point reverse.
	pretty.Println("inst:", inst)
	panic("emitInstFDIVR: not yet implemented")
}

// --- [ FDIVRP ] --------------------------------------------------------------

// liftInstFDIVRP lifts the given x87 FDIVRP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFDIVRP(inst *x86.Inst) error {
	// FDIVRP - Divide floating-point reverse and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFDIVRP: not yet implemented")
}

// --- [ FIDIVR ] --------------------------------------------------------------

// liftInstFIDIVR lifts the given x87 FIDIVR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFIDIVR(inst *x86.Inst) error {
	// FIDIVR - Divide integer reverse.
	pretty.Println("inst:", inst)
	panic("emitInstFIDIVR: not yet implemented")
}

// --- [ FPREM ] ---------------------------------------------------------------

// liftInstFPREM lifts the given x87 FPREM instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFPREM(inst *x86.Inst) error {
	// FPREM - Partial remainder.
	pretty.Println("inst:", inst)
	panic("emitInstFPREM: not yet implemented")
}

// --- [ FPREM1 ] --------------------------------------------------------------

// liftInstFPREM1 lifts the given x87 FPREM1 instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFPREM1(inst *x86.Inst) error {
	// FPREM1 - IEEE Partial remainder.
	pretty.Println("inst:", inst)
	panic("emitInstFPREM1: not yet implemented")
}

// --- [ FABS ] ----------------------------------------------------------------

// liftInstFABS lifts the given x87 FABS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFABS(inst *x86.Inst) error {
	// FABS - Absolute value.
	pretty.Println("inst:", inst)
	panic("emitInstFABS: not yet implemented")
}

// --- [ FCHS ] ----------------------------------------------------------------

// liftInstFCHS lifts the given x87 FCHS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCHS(inst *x86.Inst) error {
	// FCHS - Change sign.
	pretty.Println("inst:", inst)
	panic("emitInstFCHS: not yet implemented")
}

// --- [ FRNDINT ] -------------------------------------------------------------

// liftInstFRNDINT lifts the given x87 FRNDINT instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFRNDINT(inst *x86.Inst) error {
	// FRNDINT - Round to integer.
	pretty.Println("inst:", inst)
	panic("emitInstFRNDINT: not yet implemented")
}

// --- [ FSCALE ] --------------------------------------------------------------

// liftInstFSCALE lifts the given x87 FSCALE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFSCALE(inst *x86.Inst) error {
	// FSCALE - Scale by power of two.
	pretty.Println("inst:", inst)
	panic("emitInstFSCALE: not yet implemented")
}

// --- [ FSQRT ] ---------------------------------------------------------------

// liftInstFSQRT lifts the given x87 FSQRT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSQRT(inst *x86.Inst) error {
	// FSQRT - Square root.
	pretty.Println("inst:", inst)
	panic("emitInstFSQRT: not yet implemented")
}

// --- [ FXTRACT ] -------------------------------------------------------------

// liftInstFXTRACT lifts the given x87 FXTRACT instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFXTRACT(inst *x86.Inst) error {
	// FXTRACT - Extract exponent and significand.
	pretty.Println("inst:", inst)
	panic("emitInstFXTRACT: not yet implemented")
}

// === [ x87 FPU Comparison Instructions ] =====================================

// --- [ FCOM ] ----------------------------------------------------------------

// liftInstFCOM lifts the given x87 FCOM instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCOM(inst *x86.Inst) error {
	// FCOM - Compare floating-point.
	pretty.Println("inst:", inst)
	panic("emitInstFCOM: not yet implemented")
}

// --- [ FCOMP ] ---------------------------------------------------------------

// liftInstFCOMP lifts the given x87 FCOMP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCOMP(inst *x86.Inst) error {
	// FCOMP - Compare floating-point and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFCOMP: not yet implemented")
}

// --- [ FCOMPP ] --------------------------------------------------------------

// liftInstFCOMPP lifts the given x87 FCOMPP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCOMPP(inst *x86.Inst) error {
	// FCOMPP - Compare floating-point and pop twice.
	pretty.Println("inst:", inst)
	panic("emitInstFCOMPP: not yet implemented")
}

// --- [ FUCOM ] ---------------------------------------------------------------

// liftInstFUCOM lifts the given x87 FUCOM instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFUCOM(inst *x86.Inst) error {
	// FUCOM - Unordered compare floating-point.
	pretty.Println("inst:", inst)
	panic("emitInstFUCOM: not yet implemented")
}

// --- [ FUCOMP ] --------------------------------------------------------------

// liftInstFUCOMP lifts the given x87 FUCOMP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFUCOMP(inst *x86.Inst) error {
	// FUCOMP - Unordered compare floating-point and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFUCOMP: not yet implemented")
}

// --- [ FUCOMPP ] -------------------------------------------------------------

// liftInstFUCOMPP lifts the given x87 FUCOMPP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFUCOMPP(inst *x86.Inst) error {
	// FUCOMPP - Unordered compare floating-point and pop twice.
	pretty.Println("inst:", inst)
	panic("emitInstFUCOMPP: not yet implemented")
}

// --- [ FICOM ] ---------------------------------------------------------------

// liftInstFICOM lifts the given x87 FICOM instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFICOM(inst *x86.Inst) error {
	// FICOM - Compare integer.
	pretty.Println("inst:", inst)
	panic("emitInstFICOM: not yet implemented")
}

// --- [ FICOMP ] --------------------------------------------------------------

// liftInstFICOMP lifts the given x87 FICOMP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFICOMP(inst *x86.Inst) error {
	// FICOMP - Compare integer and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFICOMP: not yet implemented")
}

// --- [ FCOMI ] ---------------------------------------------------------------

// liftInstFCOMI lifts the given x87 FCOMI instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCOMI(inst *x86.Inst) error {
	// FCOMI - Compare floating-point and set EFLAGS.
	pretty.Println("inst:", inst)
	panic("emitInstFCOMI: not yet implemented")
}

// --- [ FUCOMI ] --------------------------------------------------------------

// liftInstFUCOMI lifts the given x87 FUCOMI instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFUCOMI(inst *x86.Inst) error {
	// FUCOMI - Unordered compare floating-point and set EFLAGS.
	pretty.Println("inst:", inst)
	panic("emitInstFUCOMI: not yet implemented")
}

// --- [ FCOMIP ] --------------------------------------------------------------

// liftInstFCOMIP lifts the given x87 FCOMIP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCOMIP(inst *x86.Inst) error {
	// FCOMIP - Compare floating-point, set EFLAGS, and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFCOMIP: not yet implemented")
}

// --- [ FUCOMIP ] -------------------------------------------------------------

// liftInstFUCOMIP lifts the given x87 FUCOMIP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFUCOMIP(inst *x86.Inst) error {
	// FUCOMIP - Unordered compare floating-point, set EFLAGS, and pop.
	pretty.Println("inst:", inst)
	panic("emitInstFUCOMIP: not yet implemented")
}

// --- [ FTST ] ----------------------------------------------------------------

// liftInstFTST lifts the given x87 FTST instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFTST(inst *x86.Inst) error {
	// FTST - Test floating-point (compare with 0.0).
	pretty.Println("inst:", inst)
	panic("emitInstFTST: not yet implemented")
}

// --- [ FXAM ] ----------------------------------------------------------------

// liftInstFXAM lifts the given x87 FXAM instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFXAM(inst *x86.Inst) error {
	// FXAM - Examine floating-point.
	pretty.Println("inst:", inst)
	panic("emitInstFXAM: not yet implemented")
}

// === [ x87 FPU Transcendental Instructions ] =================================

// --- [ FSIN ] ----------------------------------------------------------------

// liftInstFSIN lifts the given x87 FSIN instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSIN(inst *x86.Inst) error {
	// FSIN - Sine.
	pretty.Println("inst:", inst)
	panic("emitInstFSIN: not yet implemented")
}

// --- [ FCOS ] ----------------------------------------------------------------

// liftInstFCOS lifts the given x87 FCOS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCOS(inst *x86.Inst) error {
	// FCOS - Cosine.
	pretty.Println("inst:", inst)
	panic("emitInstFCOS: not yet implemented")
}

// --- [ FSINCOS ] -------------------------------------------------------------

// liftInstFSINCOS lifts the given x87 FSINCOS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFSINCOS(inst *x86.Inst) error {
	// FSINCOS - Sine and cosine.
	pretty.Println("inst:", inst)
	panic("emitInstFSINCOS: not yet implemented")
}

// --- [ FPTAN ] ---------------------------------------------------------------

// liftInstFPTAN lifts the given x87 FPTAN instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFPTAN(inst *x86.Inst) error {
	// FPTAN - Partial tangent.
	pretty.Println("inst:", inst)
	panic("emitInstFPTAN: not yet implemented")
}

// --- [ FPATAN ] --------------------------------------------------------------

// liftInstFPATAN lifts the given x87 FPATAN instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFPATAN(inst *x86.Inst) error {
	// FPATAN - Partial arctangent.
	pretty.Println("inst:", inst)
	panic("emitInstFPATAN: not yet implemented")
}

// --- [ F2XM1 ] ---------------------------------------------------------------

// liftInstF2XM1 lifts the given x87 F2XM1 instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstF2XM1(inst *x86.Inst) error {
	// F2XM1 - 2^x - 1.
	pretty.Println("inst:", inst)
	panic("emitInstF2XM1: not yet implemented")
}

// --- [ FYL2X ] ---------------------------------------------------------------

// liftInstFYL2X lifts the given x87 FYL2X instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFYL2X(inst *x86.Inst) error {
	// FYL2X - y*log_2(x).
	pretty.Println("inst:", inst)
	panic("emitInstFYL2X: not yet implemented")
}

// --- [ FYL2XP1 ] -------------------------------------------------------------

// liftInstFYL2XP1 lifts the given x87 FYL2XP1 instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFYL2XP1(inst *x86.Inst) error {
	// FYL2XP1 - y*log_2(x+1).
	pretty.Println("inst:", inst)
	panic("emitInstFYL2XP1: not yet implemented")
}

// === [ x87 FPU Load Constants Instructions ] =================================

// --- [ FLD1 ] ----------------------------------------------------------------

// liftInstFLD1 lifts the given x87 FLD1 instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFLD1(inst *x86.Inst) error {
	// FLD1 - Load +1.0.
	//
	//    FLD1                Push +1.0 onto the FPU register stack.
	//
	// Push one of seven commonly used constants (in double extended-precision
	// floating-point format) onto the FPU register stack.
	src := constant.NewFloat(1, types.X86_FP80)
	f.fpush(src)
	return nil
}

// --- [ FLDZ ] ----------------------------------------------------------------

// liftInstFLDZ lifts the given x87 FLDZ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFLDZ(inst *x86.Inst) error {
	// FLDZ - Load +0.0.
	//
	//    FLDZ                Push +0.0 onto the FPU register stack.
	//
	// Push one of seven commonly used constants (in double extended-precision
	// floating-point format) onto the FPU register stack.
	src := constant.NewFloat(0, types.X86_FP80)
	f.fpush(src)
	return nil
}

// --- [ FLDPI ] ---------------------------------------------------------------

// liftInstFLDPI lifts the given x87 FLDPI instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFLDPI(inst *x86.Inst) error {
	// FLDPI - Load π.
	//
	//    FLDPI               Push π onto the FPU register stack.
	//
	// Push one of seven commonly used constants (in double extended-precision
	// floating-point format) onto the FPU register stack.
	src := constant.NewFloat(math.Pi, types.X86_FP80)
	f.fpush(src)
	return nil
}

// --- [ FLDL2E ] --------------------------------------------------------------

// liftInstFLDL2E lifts the given x87 FLDL2E instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFLDL2E(inst *x86.Inst) error {
	// FLDL2E - Load log_2(e).
	//
	//    FLDL2E              Push log_2(e) onto the FPU register stack.
	//
	// Push one of seven commonly used constants (in double extended-precision
	// floating-point format) onto the FPU register stack.
	src := constant.NewFloat(math.Log2E, types.X86_FP80)
	f.fpush(src)
	return nil
}

// --- [ FLDLN2 ] --------------------------------------------------------------

// liftInstFLDLN2 lifts the given x87 FLDLN2 instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFLDLN2(inst *x86.Inst) error {
	// FLDLN2 - Load log_e(2).
	//
	//    FLDLN2              Push log_e(2) onto the FPU register stack.
	//
	// Push one of seven commonly used constants (in double extended-precision
	// floating-point format) onto the FPU register stack.
	src := constant.NewFloat(math.Ln2, types.X86_FP80)
	f.fpush(src)
	return nil
}

// --- [ FLDL2T ] --------------------------------------------------------------

// liftInstFLDL2T lifts the given x87 FLDL2T instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFLDL2T(inst *x86.Inst) error {
	// FLDL2T - Load log_2(10).
	//
	//    FLDL2T              Push log_2(10) onto the FPU register stack.
	//
	// Push one of seven commonly used constants (in double extended-precision
	// floating-point format) onto the FPU register stack.
	src := constant.NewFloat(math.Log2(10), types.X86_FP80)
	f.fpush(src)
	return nil
}

// --- [ FLDLG2 ] --------------------------------------------------------------

// liftInstFLDLG2 lifts the given x87 FLDLG2 instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFLDLG2(inst *x86.Inst) error {
	// FLDLG2 - Load log_10(2).
	//
	//    FLDLG2              Push log_10(2) onto the FPU register stack.
	//
	// Push one of seven commonly used constants (in double extended-precision
	// floating-point format) onto the FPU register stack.
	src := constant.NewFloat(math.Log10(2), types.X86_FP80)
	f.fpush(src)
	return nil
}

// === [ x87 FPU Control Instructions ] ========================================

// --- [ FINCSTP ] -------------------------------------------------------------

// liftInstFINCSTP lifts the given x87 FINCSTP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFINCSTP(inst *x86.Inst) error {
	// FINCSTP - Increment FPU register stack pointer.
	pretty.Println("inst:", inst)
	panic("emitInstFINCSTP: not yet implemented")
}

// --- [ FDECSTP ] -------------------------------------------------------------

// liftInstFDECSTP lifts the given x87 FDECSTP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFDECSTP(inst *x86.Inst) error {
	// FDECSTP - Decrement FPU register stack pointer.
	pretty.Println("inst:", inst)
	panic("emitInstFDECSTP: not yet implemented")
}

// --- [ FFREE ] ---------------------------------------------------------------

// liftInstFFREE lifts the given x87 FFREE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFFREE(inst *x86.Inst) error {
	// FFREE - Free floating-point register.
	pretty.Println("inst:", inst)
	panic("emitInstFFREE: not yet implemented")
}

// --- [ FINIT ] ---------------------------------------------------------------

// liftInstFINIT lifts the given x87 FINIT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFINIT(inst *x86.Inst) error {
	// FINIT - Initialize FPU after checking error conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFINIT: not yet implemented")
}

// --- [ FNINIT ] --------------------------------------------------------------

// liftInstFNINIT lifts the given x87 FNINIT instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNINIT(inst *x86.Inst) error {
	// FNINIT - Initialize FPU without checking error conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFNINIT: not yet implemented")
}

// --- [ FCLEX ] ---------------------------------------------------------------

// liftInstFCLEX lifts the given x87 FCLEX instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCLEX(inst *x86.Inst) error {
	// FCLEX - Clear floating-point exception flags after checking for error
	// conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFCLEX: not yet implemented")
}

// --- [ FNCLEX ] --------------------------------------------------------------

// liftInstFNCLEX lifts the given x87 FNCLEX instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNCLEX(inst *x86.Inst) error {
	// FNCLEX - Clear floating-point exception flags without checking for error
	// conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFNCLEX: not yet implemented")
}

// --- [ FSTCW ] ---------------------------------------------------------------

// liftInstFSTCW lifts the given x87 FSTCW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSTCW(inst *x86.Inst) error {
	// FSTCW - Store FPU control word after checking error conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFSTCW: not yet implemented")
}

// --- [ FNSTCW ] --------------------------------------------------------------

// liftInstFNSTCW lifts the given x87 FNSTCW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNSTCW(inst *x86.Inst) error {
	// FNSTCW - Store FPU control word without checking error conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFNSTCW: not yet implemented")
}

// --- [ FLDCW ] ---------------------------------------------------------------

// liftInstFLDCW lifts the given x87 FLDCW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFLDCW(inst *x86.Inst) error {
	// FLDCW - Load FPU control word.
	pretty.Println("inst:", inst)
	panic("emitInstFLDCW: not yet implemented")
}

// --- [ FSTENV ] --------------------------------------------------------------

// liftInstFSTENV lifts the given x87 FSTENV instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFSTENV(inst *x86.Inst) error {
	// FSTENV - Store FPU environment after checking error conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFSTENV: not yet implemented")
}

// --- [ FNSTENV ] -------------------------------------------------------------

// liftInstFNSTENV lifts the given x87 FNSTENV instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNSTENV(inst *x86.Inst) error {
	// FNSTENV - Store FPU environment without checking error conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFNSTENV: not yet implemented")
}

// --- [ FLDENV ] --------------------------------------------------------------

// liftInstFLDENV lifts the given x87 FLDENV instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFLDENV(inst *x86.Inst) error {
	// FLDENV - Load FPU environment.
	pretty.Println("inst:", inst)
	panic("emitInstFLDENV: not yet implemented")
}

// --- [ FSAVE ] ---------------------------------------------------------------

// liftInstFSAVE lifts the given x87 FSAVE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSAVE(inst *x86.Inst) error {
	// FSAVE - Save FPU state after checking error conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFSAVE: not yet implemented")
}

// --- [ FNSAVE ] --------------------------------------------------------------

// liftInstFNSAVE lifts the given x87 FNSAVE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNSAVE(inst *x86.Inst) error {
	// FNSAVE - Save FPU state without checking error conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFNSAVE: not yet implemented")
}

// --- [ FRSTOR ] --------------------------------------------------------------

// liftInstFRSTOR lifts the given x87 FRSTOR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFRSTOR(inst *x86.Inst) error {
	// FRSTOR - Restore FPU state.
	pretty.Println("inst:", inst)
	panic("emitInstFRSTOR: not yet implemented")
}

// --- [ FSTSW ] ---------------------------------------------------------------

// liftInstFSTSW lifts the given x87 FSTSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSTSW(inst *x86.Inst) error {
	// FSTSW - Store FPU status word after checking error conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFSTSW: not yet implemented")
}

// --- [ FNSTSW ] --------------------------------------------------------------

// liftInstFNSTSW lifts the given x87 FNSTSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNSTSW(inst *x86.Inst) error {
	// FNSTSW - Store FPU status word without checking error conditions.
	pretty.Println("inst:", inst)
	panic("emitInstFNSTSW: not yet implemented")
}

// --- [ FWAIT ] ---------------------------------------------------------------

// liftInstWAIT_FWAIT lifts the given x87 WAIT_FWAIT instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstFWAIT(inst *x86.Inst) error {
	// FWAIT - Wait for FPU.
	pretty.Println("inst:", inst)
	panic("emitInstFWAIT: not yet implemented")
}

// --- [ FNOP ] ----------------------------------------------------------------

// liftInstFNOP lifts the given x87 FNOP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFNOP(inst *x86.Inst) error {
	// FNOP - FPU no operation.
	pretty.Println("inst:", inst)
	panic("emitInstFNOP: not yet implemented")
}

// ### [ Helper functions ] ####################################################

// fpush pushes the given value to the top of the FPU register stack, emitting
// code to f.
func (f *Func) fpush(src value.Value) {
	// Decrement st.
	tmp1 := f.cur.NewLoad(f.st)
	targetTrue := &ir.BasicBlock{}
	targetFalse := &ir.BasicBlock{}
	follow := &ir.BasicBlock{}
	targetTrue.NewBr(follow)
	targetFalse.NewBr(follow)
	zero := constant.NewInt(0, types.I8)
	cond := f.cur.NewICmp(ir.IntEQ, tmp1, zero)
	f.cur.NewCondBr(cond, targetTrue, targetFalse)
	f.cur = targetTrue
	f.AppendBlock(targetTrue)
	seven := constant.NewInt(7, types.I8)
	f.cur.NewStore(seven, f.st)
	f.cur = targetFalse
	f.AppendBlock(targetFalse)
	one := constant.NewInt(1, types.I8)
	tmp2 := f.cur.NewSub(tmp1, one)
	f.cur.NewStore(tmp2, f.st)
	f.cur = follow
	f.AppendBlock(follow)

	// Store arg at st(0).
	f.fstore(src)
}

// fpop pops and returns the top of the FPU register stack, emitting code to f.
func (f *Func) fpop() value.Value {
	// Load arg from st(0).
	v := f.fload()

	// TODO: Mark st(0) as empty before incrementing st.

	// Increment st.
	tmp1 := f.cur.NewLoad(f.st)
	targetTrue := &ir.BasicBlock{}
	targetFalse := &ir.BasicBlock{}
	follow := &ir.BasicBlock{}
	targetTrue.NewBr(follow)
	targetFalse.NewBr(follow)
	zero := constant.NewInt(7, types.I8)
	cond := f.cur.NewICmp(ir.IntEQ, tmp1, zero)
	f.cur.NewCondBr(cond, targetTrue, targetFalse)
	f.cur = targetTrue
	f.AppendBlock(targetTrue)
	seven := constant.NewInt(0, types.I8)
	f.cur.NewStore(seven, f.st)
	f.cur = targetFalse
	f.AppendBlock(targetFalse)
	one := constant.NewInt(1, types.I8)
	tmp2 := f.cur.NewAdd(tmp1, one)
	f.cur.NewStore(tmp2, f.st)
	f.cur = follow
	f.AppendBlock(follow)

	// Return arg.
	return v
}

// fstore stores the source value to the top FPU register, emitting code to f.
func (f *Func) fstore(src value.Value) {
	// Store arg at st(0).
	end := &ir.BasicBlock{}
	cur := f.cur
	var cases []*ir.Case
	regs := []x86asm.Reg{x86asm.F0, x86asm.F1, x86asm.F2, x86asm.F3, x86asm.F4, x86asm.F5, x86asm.F6, x86asm.F7}
	for i, reg := range regs {
		block := &ir.BasicBlock{}
		block.NewBr(end)
		f.cur = block
		f.AppendBlock(block)
		dst := f.reg(reg)
		f.cur.NewStore(src, dst)
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
}

// fload returns the value of the top FPU register, emitting code to f.
func (f *Func) fload() value.Value {
	// Load value from st(0).
	end := &ir.BasicBlock{}
	cur := f.cur
	var cases []*ir.Case
	regs := []x86asm.Reg{x86asm.F0, x86asm.F1, x86asm.F2, x86asm.F3, x86asm.F4, x86asm.F5, x86asm.F6, x86asm.F7}
	var incs []*ir.Incoming
	for i, reg := range regs {
		block := &ir.BasicBlock{}
		block.NewBr(end)
		f.cur = block
		f.AppendBlock(block)
		src := f.reg(reg)
		v := f.cur.NewLoad(src)
		inc := &ir.Incoming{
			X:    v,
			Pred: block,
		}
		incs = append(incs, inc)
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
	return f.cur.NewPhi(incs...)
}
