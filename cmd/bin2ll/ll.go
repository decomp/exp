package main

import (
	"fmt"

	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
	"github.com/mewbak/x86/x86asm"
	"github.com/pkg/errors"
)

// translateFunc translates the given function from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) translateFunc(f *function) error {
	if f.Function == nil {
		f.Function = &ir.Function{}
	}
	for _, block := range f.blocks {
		if err := d.translateBlock(f, block); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// translateBlock translates the given basic block from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) translateBlock(f *function, block *basicBlock) error {
	block.BasicBlock = &ir.BasicBlock{}
	for _, inst := range block.insts {
		if err := d.translateInst(f, block, inst); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// translateInst translates the given instruction from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) translateInst(f *function, block *basicBlock, inst *instruction) error {
	switch inst.Op {
	case x86asm.AND:
		return d.instAND(f, block, inst)
	default:
		panic(fmt.Errorf("support for instruction opcode %v not yet implemented", inst.Op))
	}
}

// instAND translates the given AND instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instAND(f *function, block *basicBlock, inst *instruction) error {
	x := d.useArg(inst.Args[0])
	y := d.useArg(inst.Args[1])
	result := ir.NewAnd(x, y)
	d.defArg(inst.Args[0], result)
	return nil
}

func (d *disassembler) useArg(arg x86asm.Arg) value.Value {
	switch arg := arg.(type) {
	//case x86asm.Reg:
	//case x86asm.Mem:
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		fmt.Println("arg:", arg)
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

func (d *disassembler) defArg(arg x86asm.Arg, v value.Value) value.Value {
	switch arg := arg.(type) {
	//case x86asm.Reg:
	//case x86asm.Mem:
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		fmt.Println("arg:", arg)
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}
