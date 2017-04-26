package main

import (
	"fmt"

	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
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

// emitREPInst translates the given REP prefixed x86 instruction to LLVM IR,
// emitting code to f.
func (f *Func) emitREPInst(inst *Inst) error {
	loop := &ir.BasicBlock{}
	body := &ir.BasicBlock{}
	exit := &ir.BasicBlock{}
	f.AppendBlock(loop)
	f.AppendBlock(body)
	f.AppendBlock(exit)
	f.cur.NewBr(loop)
	// Generate loop basic block.
	f.cur = loop
	ecx := f.useReg(ECX)
	zero := constant.NewInt(0, types.I32)
	cond := f.cur.NewICmp(ir.IntNE, ecx, zero)
	f.cur.NewCondBr(cond, body, exit)
	// Generate body basic block.
	f.cur = body
	switch inst.Op {
	case x86asm.MOVSD:
		if err := f.emitInstMOVSD(inst); err != nil {
			return errors.WithStack(err)
		}
	default:
		panic(fmt.Errorf("support for REP prefixed %v instruction not yet implemented", inst.Op))
	}
	ecx = f.useReg(ECX)
	one := constant.NewInt(1, types.I32)
	tmp := f.cur.NewSub(ecx, one)
	f.defReg(ECX, tmp)
	f.cur.NewBr(loop)
	// Generate exit block.
	f.cur = exit
	return nil
}

// emitREPNInst translates the given REPN prefixed x86 instruction to LLVM IR,
// emitting code to f.
func (f *Func) emitREPNInst(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitREPNInst: not yet implemented")
}
