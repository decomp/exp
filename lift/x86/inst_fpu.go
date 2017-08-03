// x87 FPU Instructions.
//
// ref: $ 5.2 X87 FPU INSTRUCTIONS, Intel 64 and IA-32 architectures software
// developer's manual volume 1: Basic architecture.

package x86

import (
	"fmt"
	"math"

	"github.com/decomp/exp/disasm/x86"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
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
	panic("liftInstFIST: not yet implemented")
}

// --- [ FISTP ] ---------------------------------------------------------------

// liftInstFISTP lifts the given x87 FISTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFISTP(inst *x86.Inst) error {
	// FISTP - Store integer and pop.
	pretty.Println("inst:", inst)
	panic("liftInstFISTP: not yet implemented")
}

// --- [ FBLD ] ----------------------------------------------------------------

// liftInstFBLD lifts the given x87 FBLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFBLD(inst *x86.Inst) error {
	// FBLD - Load BCD.
	pretty.Println("inst:", inst)
	panic("liftInstFBLD: not yet implemented")
}

// --- [ FBSTP ] ---------------------------------------------------------------

// liftInstFBSTP lifts the given x87 FBSTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFBSTP(inst *x86.Inst) error {
	// FBSTP - Store BCD and pop.
	pretty.Println("inst:", inst)
	panic("liftInstFBSTP: not yet implemented")
}

// --- [ FXCH ] ----------------------------------------------------------------

// liftInstFXCH lifts the given x87 FXCH instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFXCH(inst *x86.Inst) error {
	// FXCH - Exchange registers.
	pretty.Println("inst:", inst)
	panic("liftInstFXCH: not yet implemented")
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
	panic("liftInstFCMOVE: not yet implemented")
}

// --- [ FCMOVNE ] -------------------------------------------------------------

// liftInstFCMOVNE lifts the given x87 FCMOVNE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVNE(inst *x86.Inst) error {
	// FCMOVNE - Floating-point conditional move if not equal.
	pretty.Println("inst:", inst)
	panic("liftInstFCMOVNE: not yet implemented")
}

// --- [ FCMOVB ] --------------------------------------------------------------

// liftInstFCMOVB lifts the given x87 FCMOVB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVB(inst *x86.Inst) error {
	// FCMOVB - Floating-point conditional move if below.
	pretty.Println("inst:", inst)
	panic("liftInstFCMOVB: not yet implemented")
}

// --- [ FCMOVBE ] -------------------------------------------------------------

// liftInstFCMOVBE lifts the given x87 FCMOVBE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVBE(inst *x86.Inst) error {
	// FCMOVBE - Floating-point conditional move if below or equal.
	pretty.Println("inst:", inst)
	panic("liftInstFCMOVBE: not yet implemented")
}

// --- [ FCMOVNB ] -------------------------------------------------------------

// liftInstFCMOVNB lifts the given x87 FCMOVNB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVNB(inst *x86.Inst) error {
	// FCMOVNB - Floating-point conditional move if not below.
	pretty.Println("inst:", inst)
	panic("liftInstFCMOVNB: not yet implemented")
}

// --- [ FCMOVNBE ] ------------------------------------------------------------

// liftInstFCMOVNBE lifts the given x87 FCMOVNBE instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstFCMOVNBE(inst *x86.Inst) error {
	// FCMOVNBE - Floating-point conditional move if not below or equal.
	pretty.Println("inst:", inst)
	panic("liftInstFCMOVNBE: not yet implemented")
}

// --- [ FCMOVU ] --------------------------------------------------------------

// liftInstFCMOVU lifts the given x87 FCMOVU instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVU(inst *x86.Inst) error {
	// FCMOVU - Floating-point conditional move if unordered.
	pretty.Println("inst:", inst)
	panic("liftInstFCMOVU: not yet implemented")
}

// --- [ FCMOVNU ] -------------------------------------------------------------

// liftInstFCMOVNU lifts the given x87 FCMOVNU instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCMOVNU(inst *x86.Inst) error {
	// FCMOVNU - Floating-point conditional move if not unordered.
	pretty.Println("inst:", inst)
	panic("liftInstFCMOVNU: not yet implemented")
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
		// Two-operand form.
		dst := f.useArg(inst.Arg(0))
		src := f.useArg(inst.Arg(1))
		result := f.cur.NewFAdd(dst, src)
		f.defArg(inst.Arg(0), result)
		return nil
	}
	// One-operand form.
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
	if inst.Args[1] != nil {
		// Two-operand form.
		dst := f.useArg(inst.Arg(0))
		src := f.useArg(inst.Arg(1))
		result := f.cur.NewFAdd(dst, src)
		f.defArg(inst.Arg(0), result)
		return nil
	}
	// Zero-operand form.

	// TODO: Figure out how to handle F1, directly or through abstraction since
	// the underlying register of F1 changes as ST is updated.
	st0 := f.useReg(x86.NewReg(x86asm.F0, inst))
	st1 := f.useReg(x86.NewReg(x86asm.F1, inst))
	result := f.cur.NewFAdd(st0, st1)
	f.defReg(x86.NewReg(x86asm.F1, inst), result)

	f.pop()
	return nil
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
	panic("liftInstFIADD: not yet implemented")
}

// --- [ FSUB ] ----------------------------------------------------------------

// liftInstFSUB lifts the given x87 FSUB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSUB(inst *x86.Inst) error {
	// FSUB - Subtract floating-point.
	//
	//    FSUB m32fp             Subtract m32fp from ST(0) and store result in ST(0).
	//    FSUB m64fp             Subtract m64fp from ST(0) and store result in ST(0).
	//    FSUB ST(0), ST(i)      Subtract ST(i) from ST(0) and store result in ST(0).
	//    FSUB ST(i), ST(0)      Subtract ST(0) from ST(i) and store result in ST(i).
	//
	// Subtracts the source operand from the destination operand and stores the
	// difference in the destination location.
	pretty.Println("inst:", inst)
	panic("liftInstFSUB: not yet implemented")
}

// --- [ FSUBP ] ---------------------------------------------------------------

// liftInstFSUBP lifts the given x87 FSUBP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSUBP(inst *x86.Inst) error {
	// FSUBP - Subtract floating-point and pop.
	//
	//    FSUBP ST(i), ST(0)     Subtract ST(0) from ST(i), store result in ST(i), and pop register stack.
	//    FSUBP                  Subtract ST(0) from ST(1), store result in ST(1), and pop register stack.
	//
	// Subtracts the source operand from the destination operand and stores the
	// difference in the destination location.
	pretty.Println("inst:", inst)
	panic("liftInstFSUBP: not yet implemented")
}

// --- [ FISUB ] ---------------------------------------------------------------

// liftInstFISUB lifts the given x87 FISUB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFISUB(inst *x86.Inst) error {
	// FISUB - Subtract integer.
	//
	//    FISUB m16int           Subtract m16int from ST(0) and store result in ST(0).
	//    FISUB m32int           Subtract m32int from ST(0) and store result in ST(0).
	//
	// Subtracts the source operand from the destination operand and stores the
	// difference in the destination location.
	arg := f.useArg(inst.Arg(0))
	src := f.cur.NewSIToFP(arg, types.X86_FP80)
	st0 := f.fload()
	result := f.cur.NewFSub(st0, src)
	f.fstore(result)
	return nil
}

// --- [ FSUBR ] ---------------------------------------------------------------

// liftInstFSUBR lifts the given x87 FSUBR instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSUBR(inst *x86.Inst) error {
	// FSUBR - Subtract floating-point reverse.
	pretty.Println("inst:", inst)
	panic("liftInstFSUBR: not yet implemented")
}

// --- [ FSUBRP ] --------------------------------------------------------------

// liftInstFSUBRP lifts the given x87 FSUBRP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFSUBRP(inst *x86.Inst) error {
	// FSUBRP - Subtract floating-point reverse and pop.
	pretty.Println("inst:", inst)
	panic("liftInstFSUBRP: not yet implemented")
}

// --- [ FISUBR ] --------------------------------------------------------------

// liftInstFISUBR lifts the given x87 FISUBR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFISUBR(inst *x86.Inst) error {
	// FISUBR - Subtract integer reverse.
	pretty.Println("inst:", inst)
	panic("liftInstFISUBR: not yet implemented")
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
		// Two-operand form.
		dst := f.useArg(inst.Arg(0))
		src := f.useArg(inst.Arg(1))
		result := f.cur.NewFMul(dst, src)
		f.defArg(inst.Arg(0), result)
		return nil
	}
	// One-operand form.
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
	//
	//    FMULP ST(i), ST(0)     Multiply ST(i) by ST(0), store result in ST(i), and pop the register stack.
	//    FMULP                  Multiply ST(1) by ST(0), store result in ST(1), and pop the register stack.
	//
	// Multiplies the destination and source operands and stores the product in
	// the destination location.
	pretty.Println("inst:", inst)
	panic("liftInstFMULP: not yet implemented")
}

// --- [ FIMUL ] ---------------------------------------------------------------

// liftInstFIMUL lifts the given x87 FIMUL instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFIMUL(inst *x86.Inst) error {
	// FIMUL - Multiply integer.
	//
	//    FIMUL m16int           Multiply ST(0) by m16int and store result in ST(0).
	//    FIMUL m32int           Multiply ST(0) by m32int and store result in ST(0).
	//
	// Multiplies the destination and source operands and stores the product in
	// the destination location.
	arg := f.useArg(inst.Arg(0))
	src := f.cur.NewSIToFP(arg, types.X86_FP80)
	st0 := f.fload()
	result := f.cur.NewFMul(st0, src)
	f.fstore(result)
	return nil
}

// --- [ FDIV ] ----------------------------------------------------------------

// liftInstFDIV lifts the given x87 FDIV instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFDIV(inst *x86.Inst) error {
	// FDIV - Divide floating-point.
	//
	//    FDIV m32fp          Divide ST(0) by m32fp and store result in ST(0).
	//    FDIV m64fp          Divide ST(0) by m64fp and store result in ST(0).
	//    FDIV ST(0), ST(i)   Divide ST(0) by ST(i) and store result in ST(0).
	//    FDIV ST(i), ST(0)   Divide ST(i) by ST(0) and store result in ST(i).
	//
	// Divides the destination operand by the source operand and stores the
	// result in the destination location.
	if inst.Args[1] != nil {
		// Two-operand form.
		dst := f.useArg(inst.Arg(0))
		src := f.useArg(inst.Arg(1))
		result := f.cur.NewFDiv(dst, src)
		f.defArg(inst.Arg(0), result)
		return nil
	}
	// One-operand form.
	arg := f.useArg(inst.Arg(0))
	src := f.cur.NewFPExt(arg, types.X86_FP80)
	st0 := f.fload()
	result := f.cur.NewFDiv(st0, src)
	f.fstore(result)
	return nil
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
	// Divides the destination operand by the source operand and stores the
	// result in the destination location.
	if inst.Args[1] != nil {
		// Two-operand form.
		dst := f.useArg(inst.Arg(0))
		src := f.useArg(inst.Arg(1))
		result := f.cur.NewFDiv(dst, src)
		f.defArg(inst.Arg(0), result)
		f.fpop()
		return nil
	}
	// Zero-operand form.
	dst := f.useReg(x86.NewReg(x86asm.F1, inst))
	src := f.useReg(x86.NewReg(x86asm.F0, inst))
	result := f.cur.NewFDiv(dst, src)
	f.defReg(x86.NewReg(x86asm.F1, inst), result)
	f.pop()
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
	//
	//    FDIVR m32fp                   Divide m32fp by ST(0) and store result in ST(0).
	//    FDIVR m64fp                   Divide m64fp by ST(0) and store result in ST(0).
	//    FDIVR ST(0), ST(i)            Divide ST(i) by ST(0) and store result in ST(0).
	//    FDIVR ST(i), ST(0)            Divide ST(0) by ST(i) and store result in ST(i).
	//
	// Divides the source operand by the destination operand and stores the
	// result in the destination location.

	pretty.Println("inst:", inst)
	panic("liftInstFDIVR: not yet implemented")
}

// --- [ FDIVRP ] --------------------------------------------------------------

// liftInstFDIVRP lifts the given x87 FDIVRP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFDIVRP(inst *x86.Inst) error {
	// FDIVRP - Divide floating-point reverse and pop.
	//
	//    FDIVRP ST(i), ST(0)           Divide ST(0) by ST(i), store result in ST(i), and pop the register stack.
	//    FDIVRP                        Divide ST(0) by ST(1), store result in ST(1), and pop the register stack.
	//
	// Divides the source operand by the destination operand and stores the
	// result in the destination location.
	if inst.Args[1] != nil {
		// Two-operand form.
		dst := f.useArg(inst.Arg(0))
		src := f.useArg(inst.Arg(1))
		result := f.cur.NewFDiv(src, dst)
		f.defArg(inst.Arg(0), result)
		f.pop()
		return nil
	}
	// Zero-operand form.
	dst := f.useReg(x86.NewReg(x86asm.F1, inst))
	src := f.useReg(x86.NewReg(x86asm.F0, inst))
	result := f.cur.NewFDiv(src, dst)
	f.defReg(x86.NewReg(x86asm.F1, inst), result)
	f.pop()
	return nil
}

// --- [ FIDIVR ] --------------------------------------------------------------

// liftInstFIDIVR lifts the given x87 FIDIVR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFIDIVR(inst *x86.Inst) error {
	// FIDIVR - Divide integer reverse.
	//
	//    FIDIVR m32int       Divide m32int by ST(0) and store result in ST(0).
	//    FIDIVR m16int       Divide m16int by ST(0) and store result in ST(0).
	//
	// Divides the source operand by the destination operand and stores the
	// result in the destination location.

	pretty.Println("inst:", inst)
	panic("liftInstFIDIVR: not yet implemented")
}

// --- [ FPREM ] ---------------------------------------------------------------

// liftInstFPREM lifts the given x87 FPREM instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFPREM(inst *x86.Inst) error {
	// FPREM - Partial remainder.
	pretty.Println("inst:", inst)
	panic("liftInstFPREM: not yet implemented")
}

// --- [ FPREM1 ] --------------------------------------------------------------

// liftInstFPREM1 lifts the given x87 FPREM1 instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFPREM1(inst *x86.Inst) error {
	// FPREM1 - IEEE Partial remainder.
	pretty.Println("inst:", inst)
	panic("liftInstFPREM1: not yet implemented")
}

// --- [ FABS ] ----------------------------------------------------------------

// liftInstFABS lifts the given x87 FABS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFABS(inst *x86.Inst) error {
	// FABS - Absolute value.
	pretty.Println("inst:", inst)
	panic("liftInstFABS: not yet implemented")
}

// --- [ FCHS ] ----------------------------------------------------------------

// liftInstFCHS lifts the given x87 FCHS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCHS(inst *x86.Inst) error {
	// FCHS - Change sign.
	pretty.Println("inst:", inst)
	panic("liftInstFCHS: not yet implemented")
}

// --- [ FRNDINT ] -------------------------------------------------------------

// liftInstFRNDINT lifts the given x87 FRNDINT instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFRNDINT(inst *x86.Inst) error {
	// FRNDINT - Round to integer.
	pretty.Println("inst:", inst)
	panic("liftInstFRNDINT: not yet implemented")
}

// --- [ FSCALE ] --------------------------------------------------------------

// liftInstFSCALE lifts the given x87 FSCALE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFSCALE(inst *x86.Inst) error {
	// FSCALE - Scale by power of two.
	pretty.Println("inst:", inst)
	panic("liftInstFSCALE: not yet implemented")
}

// --- [ FSQRT ] ---------------------------------------------------------------

// liftInstFSQRT lifts the given x87 FSQRT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSQRT(inst *x86.Inst) error {
	// FSQRT - Square root.
	pretty.Println("inst:", inst)
	panic("liftInstFSQRT: not yet implemented")
}

// --- [ FXTRACT ] -------------------------------------------------------------

// liftInstFXTRACT lifts the given x87 FXTRACT instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFXTRACT(inst *x86.Inst) error {
	// FXTRACT - Extract exponent and significand.
	pretty.Println("inst:", inst)
	panic("liftInstFXTRACT: not yet implemented")
}

// === [ x87 FPU Comparison Instructions ] =====================================

// --- [ FCOM ] ----------------------------------------------------------------

// liftInstFCOM lifts the given x87 FCOM instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCOM(inst *x86.Inst) error {
	// FCOM - Compare floating-point.
	//
	//    FCOM m32fp          Compare ST(0) with m32fp.
	//    FCOM m64fp          Compare ST(0) with m64fp.
	//    FCOM ST(i)          Compare ST(0) with ST(i).
	//    FCOM                Compare ST(0) with ST(1).
	//
	// Compares the contents of register ST(0) and source value and sets
	// condition code flags C0, C2, and C3 in the FPU status word according to
	// the results (see the table below).
	//
	//    Condition       C3 C2 C0
	//
	//    ST(0) > SRC      0  0  0
	//    ST(0) < SRC      0  0  1
	//    ST(0) = SRC      1  0  0
	//    Unordered        1  1  1
	if inst.Args[0] == nil {
		panic(fmt.Errorf("support for zero-operand FCOM not yet implemented; instruction %v at address %v", inst, inst.Addr))
	}
	src := f.useArg(inst.Arg(0))
	if !types.Equal(src.Type(), types.X86_FP80) {
		src = f.cur.NewFPExt(src, types.X86_FP80)
	}
	st0 := f.fload()
	a := f.cur.NewFCmp(ir.FloatOGT, st0, src)
	b := f.cur.NewFCmp(ir.FloatOLT, st0, src)
	c := f.cur.NewFCmp(ir.FloatOEQ, st0, src)
	d := f.cur.NewFCmp(ir.FloatUNO, st0, src)
	end := &ir.BasicBlock{}
	targetA := &ir.BasicBlock{}
	targetA.NewBr(end)
	targetB := &ir.BasicBlock{}
	targetB.NewBr(end)
	targetC := &ir.BasicBlock{}
	targetC.NewBr(end)
	targetD := &ir.BasicBlock{}
	targetD.NewBr(end)
	next := &ir.BasicBlock{}
	f.cur.NewCondBr(a, targetA, next)
	f.cur = next
	f.AppendBlock(next)
	next = &ir.BasicBlock{}
	f.cur.NewCondBr(b, targetB, next)
	f.cur = next
	f.AppendBlock(next)
	next = &ir.BasicBlock{}
	f.cur.NewCondBr(c, targetC, next)
	f.cur = next
	f.AppendBlock(next)
	next = &ir.BasicBlock{}
	f.cur.NewCondBr(d, targetD, end)
	f.cur = targetA
	f.AppendBlock(targetA)
	f.defFStatus(C0, constant.False)
	f.defFStatus(C2, constant.False)
	f.defFStatus(C3, constant.False)
	f.cur = targetB
	f.AppendBlock(targetB)
	f.defFStatus(C0, constant.True)
	f.defFStatus(C2, constant.False)
	f.defFStatus(C3, constant.False)
	f.cur = targetC
	f.AppendBlock(targetC)
	f.defFStatus(C0, constant.False)
	f.defFStatus(C2, constant.False)
	f.defFStatus(C3, constant.True)
	f.cur = targetD
	f.AppendBlock(targetD)
	f.defFStatus(C0, constant.True)
	f.defFStatus(C2, constant.True)
	f.defFStatus(C3, constant.True)
	f.cur = end
	f.AppendBlock(end)
	return nil
}

// --- [ FCOMP ] ---------------------------------------------------------------

// liftInstFCOMP lifts the given x87 FCOMP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCOMP(inst *x86.Inst) error {
	// FCOMP - Compare floating-point and pop.
	//
	//    FCOMP m32fp         Compare ST(0) with m32fp and pop register stack.
	//    FCOMP m64fp         Compare ST(0) with m64fp and pop register stack.
	//    FCOMP ST(i)         Compare ST(0) with ST(i) and pop register stack.
	//    FCOMP               Compare ST(0) with ST(1) and pop register stack.
	//
	// Compares the contents of register ST(0) and source value and sets
	// condition code flags C0, C2, and C3 in the FPU status word according to
	// the results (see the table below).
	//
	//    Condition       C3 C2 C0
	//
	//    ST(0) > SRC      0  0  0
	//    ST(0) < SRC      0  0  1
	//    ST(0) = SRC      1  0  0
	//    Unordered        1  1  1
	if err := f.liftInstFCOM(inst); err != nil {
		return errors.WithStack(err)
	}
	f.pop()
	return nil
}

// --- [ FCOMPP ] --------------------------------------------------------------

// liftInstFCOMPP lifts the given x87 FCOMPP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCOMPP(inst *x86.Inst) error {
	// FCOMPP - Compare floating-point and pop twice.
	//
	//    FCOMPP              Compare ST(0) with ST(1) and pop register stack twice.
	//
	// Compares the contents of register ST(0) and source value and sets
	// condition code flags C0, C2, and C3 in the FPU status word according to
	// the results (see the table below).
	//
	//    Condition       C3 C2 C0
	//
	//    ST(0) > SRC      0  0  0
	//    ST(0) < SRC      0  0  1
	//    ST(0) = SRC      1  0  0
	//    Unordered        1  1  1
	if err := f.liftInstFCOM(inst); err != nil {
		return errors.WithStack(err)
	}
	f.pop()
	f.pop()
	return nil
}

// --- [ FUCOM ] ---------------------------------------------------------------

// liftInstFUCOM lifts the given x87 FUCOM instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFUCOM(inst *x86.Inst) error {
	// FUCOM - Unordered compare floating-point.
	pretty.Println("inst:", inst)
	panic("liftInstFUCOM: not yet implemented")
}

// --- [ FUCOMP ] --------------------------------------------------------------

// liftInstFUCOMP lifts the given x87 FUCOMP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFUCOMP(inst *x86.Inst) error {
	// FUCOMP - Unordered compare floating-point and pop.
	pretty.Println("inst:", inst)
	panic("liftInstFUCOMP: not yet implemented")
}

// --- [ FUCOMPP ] -------------------------------------------------------------

// liftInstFUCOMPP lifts the given x87 FUCOMPP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFUCOMPP(inst *x86.Inst) error {
	// FUCOMPP - Unordered compare floating-point and pop twice.
	pretty.Println("inst:", inst)
	panic("liftInstFUCOMPP: not yet implemented")
}

// --- [ FICOM ] ---------------------------------------------------------------

// liftInstFICOM lifts the given x87 FICOM instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFICOM(inst *x86.Inst) error {
	// FICOM - Compare integer.
	pretty.Println("inst:", inst)
	panic("liftInstFICOM: not yet implemented")
}

// --- [ FICOMP ] --------------------------------------------------------------

// liftInstFICOMP lifts the given x87 FICOMP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFICOMP(inst *x86.Inst) error {
	// FICOMP - Compare integer and pop.
	pretty.Println("inst:", inst)
	panic("liftInstFICOMP: not yet implemented")
}

// --- [ FCOMI ] ---------------------------------------------------------------

// liftInstFCOMI lifts the given x87 FCOMI instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCOMI(inst *x86.Inst) error {
	// FCOMI - Compare floating-point and set EFLAGS.
	pretty.Println("inst:", inst)
	panic("liftInstFCOMI: not yet implemented")
}

// --- [ FUCOMI ] --------------------------------------------------------------

// liftInstFUCOMI lifts the given x87 FUCOMI instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFUCOMI(inst *x86.Inst) error {
	// FUCOMI - Unordered compare floating-point and set EFLAGS.
	pretty.Println("inst:", inst)
	panic("liftInstFUCOMI: not yet implemented")
}

// --- [ FCOMIP ] --------------------------------------------------------------

// liftInstFCOMIP lifts the given x87 FCOMIP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFCOMIP(inst *x86.Inst) error {
	// FCOMIP - Compare floating-point, set EFLAGS, and pop.
	pretty.Println("inst:", inst)
	panic("liftInstFCOMIP: not yet implemented")
}

// --- [ FUCOMIP ] -------------------------------------------------------------

// liftInstFUCOMIP lifts the given x87 FUCOMIP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFUCOMIP(inst *x86.Inst) error {
	// FUCOMIP - Unordered compare floating-point, set EFLAGS, and pop.
	pretty.Println("inst:", inst)
	panic("liftInstFUCOMIP: not yet implemented")
}

// --- [ FTST ] ----------------------------------------------------------------

// liftInstFTST lifts the given x87 FTST instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFTST(inst *x86.Inst) error {
	// FTST - Test floating-point (compare with 0.0).
	pretty.Println("inst:", inst)
	panic("liftInstFTST: not yet implemented")
}

// --- [ FXAM ] ----------------------------------------------------------------

// liftInstFXAM lifts the given x87 FXAM instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFXAM(inst *x86.Inst) error {
	// FXAM - Examine floating-point.
	pretty.Println("inst:", inst)
	panic("liftInstFXAM: not yet implemented")
}

// === [ x87 FPU Transcendental Instructions ] =================================

// --- [ FSIN ] ----------------------------------------------------------------

// liftInstFSIN lifts the given x87 FSIN instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSIN(inst *x86.Inst) error {
	// FSIN - Sine.
	pretty.Println("inst:", inst)
	panic("liftInstFSIN: not yet implemented")
}

// --- [ FCOS ] ----------------------------------------------------------------

// liftInstFCOS lifts the given x87 FCOS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCOS(inst *x86.Inst) error {
	// FCOS - Cosine.
	pretty.Println("inst:", inst)
	panic("liftInstFCOS: not yet implemented")
}

// --- [ FSINCOS ] -------------------------------------------------------------

// liftInstFSINCOS lifts the given x87 FSINCOS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFSINCOS(inst *x86.Inst) error {
	// FSINCOS - Sine and cosine.
	pretty.Println("inst:", inst)
	panic("liftInstFSINCOS: not yet implemented")
}

// --- [ FPTAN ] ---------------------------------------------------------------

// liftInstFPTAN lifts the given x87 FPTAN instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFPTAN(inst *x86.Inst) error {
	// FPTAN - Partial tangent.
	pretty.Println("inst:", inst)
	panic("liftInstFPTAN: not yet implemented")
}

// --- [ FPATAN ] --------------------------------------------------------------

// liftInstFPATAN lifts the given x87 FPATAN instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFPATAN(inst *x86.Inst) error {
	// FPATAN - Partial arctangent.
	pretty.Println("inst:", inst)
	panic("liftInstFPATAN: not yet implemented")
}

// --- [ F2XM1 ] ---------------------------------------------------------------

// liftInstF2XM1 lifts the given x87 F2XM1 instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstF2XM1(inst *x86.Inst) error {
	// F2XM1 - 2^x - 1.
	pretty.Println("inst:", inst)
	panic("liftInstF2XM1: not yet implemented")
}

// --- [ FYL2X ] ---------------------------------------------------------------

// liftInstFYL2X lifts the given x87 FYL2X instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFYL2X(inst *x86.Inst) error {
	// FYL2X - y*log_2(x).
	pretty.Println("inst:", inst)
	panic("liftInstFYL2X: not yet implemented")
}

// --- [ FYL2XP1 ] -------------------------------------------------------------

// liftInstFYL2XP1 lifts the given x87 FYL2XP1 instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFYL2XP1(inst *x86.Inst) error {
	// FYL2XP1 - y*log_2(x+1).
	pretty.Println("inst:", inst)
	panic("liftInstFYL2XP1: not yet implemented")
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
	panic("liftInstFINCSTP: not yet implemented")
}

// --- [ FDECSTP ] -------------------------------------------------------------

// liftInstFDECSTP lifts the given x87 FDECSTP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFDECSTP(inst *x86.Inst) error {
	// FDECSTP - Decrement FPU register stack pointer.
	pretty.Println("inst:", inst)
	panic("liftInstFDECSTP: not yet implemented")
}

// --- [ FFREE ] ---------------------------------------------------------------

// liftInstFFREE lifts the given x87 FFREE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFFREE(inst *x86.Inst) error {
	// FFREE - Free floating-point register.
	pretty.Println("inst:", inst)
	panic("liftInstFFREE: not yet implemented")
}

// --- [ FINIT ] ---------------------------------------------------------------

// liftInstFINIT lifts the given x87 FINIT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFINIT(inst *x86.Inst) error {
	// FINIT - Initialize FPU after checking error conditions.
	pretty.Println("inst:", inst)
	panic("liftInstFINIT: not yet implemented")
}

// --- [ FNINIT ] --------------------------------------------------------------

// liftInstFNINIT lifts the given x87 FNINIT instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNINIT(inst *x86.Inst) error {
	// FNINIT - Initialize FPU without checking error conditions.
	pretty.Println("inst:", inst)
	panic("liftInstFNINIT: not yet implemented")
}

// --- [ FCLEX ] ---------------------------------------------------------------

// liftInstFCLEX lifts the given x87 FCLEX instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFCLEX(inst *x86.Inst) error {
	// FCLEX - Clear floating-point exception flags after checking for error
	// conditions.
	pretty.Println("inst:", inst)
	panic("liftInstFCLEX: not yet implemented")
}

// --- [ FNCLEX ] --------------------------------------------------------------

// liftInstFNCLEX lifts the given x87 FNCLEX instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNCLEX(inst *x86.Inst) error {
	// FNCLEX - Clear floating-point exception flags without checking for error
	// conditions.
	pretty.Println("inst:", inst)
	panic("liftInstFNCLEX: not yet implemented")
}

// --- [ FSTCW ] ---------------------------------------------------------------

// liftInstFSTCW lifts the given x87 FSTCW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSTCW(inst *x86.Inst) error {
	// FSTCW - Store FPU control word after checking error conditions.
	pretty.Println("inst:", inst)
	panic("liftInstFSTCW: not yet implemented")
}

// --- [ FNSTCW ] --------------------------------------------------------------

// liftInstFNSTCW lifts the given x87 FNSTCW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNSTCW(inst *x86.Inst) error {
	// FNSTCW - Store FPU control word without checking error conditions.
	pretty.Println("inst:", inst)
	panic("liftInstFNSTCW: not yet implemented")
}

// --- [ FLDCW ] ---------------------------------------------------------------

// liftInstFLDCW lifts the given x87 FLDCW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFLDCW(inst *x86.Inst) error {
	// FLDCW - Load FPU control word.
	pretty.Println("inst:", inst)
	panic("liftInstFLDCW: not yet implemented")
}

// --- [ FSTENV ] --------------------------------------------------------------

// liftInstFSTENV lifts the given x87 FSTENV instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFSTENV(inst *x86.Inst) error {
	// FSTENV - Store FPU environment after checking error conditions.
	pretty.Println("inst:", inst)
	panic("liftInstFSTENV: not yet implemented")
}

// --- [ FNSTENV ] -------------------------------------------------------------

// liftInstFNSTENV lifts the given x87 FNSTENV instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNSTENV(inst *x86.Inst) error {
	// FNSTENV - Store FPU environment without checking error conditions.
	pretty.Println("inst:", inst)
	panic("liftInstFNSTENV: not yet implemented")
}

// --- [ FLDENV ] --------------------------------------------------------------

// liftInstFLDENV lifts the given x87 FLDENV instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFLDENV(inst *x86.Inst) error {
	// FLDENV - Load FPU environment.
	pretty.Println("inst:", inst)
	panic("liftInstFLDENV: not yet implemented")
}

// --- [ FSAVE ] ---------------------------------------------------------------

// liftInstFSAVE lifts the given x87 FSAVE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSAVE(inst *x86.Inst) error {
	// FSAVE - Save FPU state after checking error conditions.
	pretty.Println("inst:", inst)
	panic("liftInstFSAVE: not yet implemented")
}

// --- [ FNSAVE ] --------------------------------------------------------------

// liftInstFNSAVE lifts the given x87 FNSAVE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNSAVE(inst *x86.Inst) error {
	// FNSAVE - Save FPU state without checking error conditions.
	pretty.Println("inst:", inst)
	panic("liftInstFNSAVE: not yet implemented")
}

// --- [ FRSTOR ] --------------------------------------------------------------

// liftInstFRSTOR lifts the given x87 FRSTOR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFRSTOR(inst *x86.Inst) error {
	// FRSTOR - Restore FPU state.
	pretty.Println("inst:", inst)
	panic("liftInstFRSTOR: not yet implemented")
}

// --- [ FSTSW ] ---------------------------------------------------------------

// liftInstFSTSW lifts the given x87 FSTSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFSTSW(inst *x86.Inst) error {
	// FSTSW - Store FPU status word after checking error conditions.
	//
	//    FSTSW m16           Store FPU status word at m16 after checking for pending unmasked floating-point exceptions.
	//    FSTSW AX            Store FPU status word in AX register after checking for pending unmasked floating-point exceptions.
	//
	// Stores the current value of the x87 FPU status word in the destination
	// location.
	if err := f.liftInstFNSTSW(inst); err != nil {
		return errors.WithStack(err)
	}
	// TODO: Check FPU error condition.
	return nil
}

// --- [ FNSTSW ] --------------------------------------------------------------

// liftInstFNSTSW lifts the given x87 FNSTSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFNSTSW(inst *x86.Inst) error {
	// FNSTSW - Store FPU status word without checking error conditions.
	//
	//    FNSTSW m16          Store FPU status word at m16 without checking for pending unmasked floating-point exceptions.
	//    FNSTSW AX           Store FPU status word in AX register without checking for pending unmasked floating-point exceptions.
	//
	// Stores the current value of the x87 FPU status word in the destination
	// location.

	// Load FPU status flags.
	b := f.useFStatus(Busy)
	c3 := f.useFStatus(C3)
	st := f.fload()
	c2 := f.useFStatus(C2)
	c1 := f.useFStatus(C1)
	c0 := f.useFStatus(C0)
	es := f.useFStatus(ES)
	sf := f.useFStatus(StackFault)
	pe := f.useFStatus(PE)
	ue := f.useFStatus(UE)
	oe := f.useFStatus(OE)
	ze := f.useFStatus(ZE)
	de := f.useFStatus(DE)
	ie := f.useFStatus(IE)

	// Extend to 16-bit.
	b = f.cur.NewZExt(b, types.I16)
	c3 = f.cur.NewZExt(b, types.I16)
	st = f.cur.NewZExt(b, types.I16)
	c2 = f.cur.NewZExt(b, types.I16)
	c1 = f.cur.NewZExt(b, types.I16)
	c0 = f.cur.NewZExt(b, types.I16)
	es = f.cur.NewZExt(b, types.I16)
	sf = f.cur.NewZExt(b, types.I16)
	pe = f.cur.NewZExt(b, types.I16)
	ue = f.cur.NewZExt(b, types.I16)
	oe = f.cur.NewZExt(b, types.I16)
	ze = f.cur.NewZExt(b, types.I16)
	de = f.cur.NewZExt(b, types.I16)
	ie = f.cur.NewZExt(b, types.I16)

	// x87 FPU Status Word
	//
	//    15    - B, FPU Busy
	//    14    - C3, Conditional Code 3
	//    11-13 - TOP, Top of Stack Pointer
	//    10    - C2, Conditional Code 2
	//     9    - C1, Conditional Code 1
	//     8    - C0, Conditional Code 0
	//     7    - ES, Exception Summary Status
	//     6    - SF, Stack Fault
	//     5    - PE, Precision
	//     4    - UE, Underflow
	//     3    - OE, Overflow
	//     2    - ZE, Zero Divide
	//     1    - DE, Denormalized Operand
	//     0    - IE, Invalid Operation
	//
	// ref: 8.1.3 x87 FPU Status Register, Intel 64 and IA-32 architectures
	// software developer's manual volume 1: Basic architecture.
	b = f.cur.NewShl(b, constant.NewInt(15, types.I64))
	c3 = f.cur.NewShl(c3, constant.NewInt(14, types.I64))
	st = f.cur.NewShl(st, constant.NewInt(11, types.I64))
	c2 = f.cur.NewShl(c2, constant.NewInt(10, types.I64))
	c1 = f.cur.NewShl(c1, constant.NewInt(9, types.I64))
	c0 = f.cur.NewShl(c0, constant.NewInt(8, types.I64))
	es = f.cur.NewShl(es, constant.NewInt(7, types.I64))
	sf = f.cur.NewShl(sf, constant.NewInt(6, types.I64))
	pe = f.cur.NewShl(pe, constant.NewInt(5, types.I64))
	ue = f.cur.NewShl(ue, constant.NewInt(4, types.I64))
	oe = f.cur.NewShl(oe, constant.NewInt(3, types.I64))
	ze = f.cur.NewShl(ze, constant.NewInt(2, types.I64))
	de = f.cur.NewShl(de, constant.NewInt(1, types.I64))
	//ie = f.cur.NewShl(ie, constant.NewInt(0, types.I64))

	tmp := f.cur.NewOr(b, c3)
	tmp = f.cur.NewOr(tmp, st)
	tmp = f.cur.NewOr(tmp, c2)
	tmp = f.cur.NewOr(tmp, c1)
	tmp = f.cur.NewOr(tmp, c0)
	tmp = f.cur.NewOr(tmp, es)
	tmp = f.cur.NewOr(tmp, sf)
	tmp = f.cur.NewOr(tmp, pe)
	tmp = f.cur.NewOr(tmp, ue)
	tmp = f.cur.NewOr(tmp, oe)
	tmp = f.cur.NewOr(tmp, ze)
	tmp = f.cur.NewOr(tmp, de)
	result := f.cur.NewOr(tmp, ie)

	// Store FPU status flags.
	f.defArg(inst.Arg(0), result)

	return nil
}

// --- [ FWAIT ] ---------------------------------------------------------------

// liftInstWAIT_FWAIT lifts the given x87 WAIT_FWAIT instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstFWAIT(inst *x86.Inst) error {
	// FWAIT - Wait for FPU.
	pretty.Println("inst:", inst)
	panic("liftInstFWAIT: not yet implemented")
}

// --- [ FNOP ] ----------------------------------------------------------------

// liftInstFNOP lifts the given x87 FNOP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstFNOP(inst *x86.Inst) error {
	// FNOP - FPU no operation.
	pretty.Println("inst:", inst)
	panic("liftInstFNOP: not yet implemented")
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
