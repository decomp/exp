package main

import "github.com/decomp/exp/bin"

// A BasicBlock is an x86 basic block.
type BasicBlock struct {
	// Entry address of the basic block.
	addr bin.Address
	// Instructions of the basic block.
	insts []*Inst
	// Terminator of the basic block.
	term *Inst
}
