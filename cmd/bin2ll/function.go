package main

import (
	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/ir"
	"golang.org/x/arch/x86/x86asm"
)

// A Func tracks the information required to translate a function from x86
// machine code to LLVM IR assembly.
type Func struct {
	// LLVM IR code for the function.
	*ir.Function
	// Entry address of the function.
	entry bin.Address
	// Current basic block being generated.
	cur *ir.BasicBlock
	// x86 basic blocks of the function.
	bbs map[bin.Address]*BasicBlock
	// LLVM IR basic blocks of the function.
	blocks map[bin.Address]*ir.BasicBlock
	// Registers used within the function.
	regs map[x86asm.Reg]*ir.InstAlloca
	// Status flags used within the function.
	statusFlags map[StatusFlag]*ir.InstAlloca
	// Read-only global disassembler state.
	d *disassembler
}
