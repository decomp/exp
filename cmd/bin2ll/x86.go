package main

import (
	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/ir"
	"github.com/mewbak/x86/x86asm"
)

type function struct {
	*ir.Function
	entry  bin.Address
	blocks []*basicBlock
}

type basicBlock struct {
	*ir.BasicBlock
	addr  bin.Address
	insts []x86asm.Inst
}

// decodeFunc decodes the x86 machine code of the function at the given address.
func (d *disassembler) decodeFunc(addr bin.Address) (*function, error) {
	panic("not yet implemented")
}

// decodeBlock decodes the x86 machine code of the basic block at the given
// address.
func (d *disassembler) decodeBlock(addr bin.Address) (*basicBlock, error) {
	panic("not yet implemented")
}
