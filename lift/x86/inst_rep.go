package x86

import (
	"fmt"

	"github.com/decomp/exp/disasm/x86"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
	"golang.org/x/arch/x86/x86asm"
)

// Repeat Prefixes
//
//    Repeat prefix      Terminating condition 1      Terminating condition 2
//
//    REP                RCX or (E)CX = 0             None
//    REPE               RCX or (E)CX = 0             ZF = 0
//    REPNE              RCX or (E)CX = 0             ZF = 1
//
// ref: $ 4.2 REP/REPE/REPZ/REPNE/REPNZ - Repeat String Operation Prefix, Intel
// 64 and IA-32 Architectures Software Developer's Manual

// liftREPInst lifs the given REP prefixed x86 instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftREPInst(inst *x86.Inst) error {
	loop := &ir.Block{}
	body := &ir.Block{}
	exit := &ir.Block{}
	f.Blocks = append(f.Blocks, loop)
	f.Blocks = append(f.Blocks, body)
	f.Blocks = append(f.Blocks, exit)
	f.cur.NewBr(loop)
	// Generate loop basic block.
	f.cur = loop
	ecx := f.useReg(x86.ECX)
	zero := constant.NewInt(types.I32, 0)
	cond := f.cur.NewICmp(enum.IPredNE, ecx, zero)
	f.cur.NewCondBr(cond, body, exit)
	// Generate body basic block.
	f.cur = body
	switch inst.Op {
	case x86asm.MOVSB:
		if err := f.liftInstMOVSB(inst); err != nil {
			return errors.WithStack(err)
		}
	case x86asm.MOVSD:
		if err := f.liftInstMOVSD(inst); err != nil {
			return errors.WithStack(err)
		}
	case x86asm.STOSB:
		if err := f.liftInstSTOSB(inst); err != nil {
			return errors.WithStack(err)
		}
	case x86asm.STOSW:
		if err := f.liftInstSTOSW(inst); err != nil {
			return errors.WithStack(err)
		}
	case x86asm.STOSD:
		if err := f.liftInstSTOSD(inst); err != nil {
			return errors.WithStack(err)
		}
	default:
		panic(fmt.Errorf("support for REP prefixed %v instruction not yet implemented", inst.Op))
	}
	ecx = f.useReg(x86.ECX)
	one := constant.NewInt(types.I32, 1)
	tmp := f.cur.NewSub(ecx, one)
	f.defReg(x86.ECX, tmp)
	f.cur.NewBr(loop)
	// Generate exit block.
	f.cur = exit
	return nil
}

// liftREPNInst lifts the given REPN prefixed x86 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftREPNInst(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitREPNInst: not yet implemented")
}
