package disasm

import "github.com/decomp/exp/bin"

// Function represents an assembly function.
type Function interface {
	// Addr returns the address of the function entry point.
	Addr() bin.Address
	// Blocks returns the basic blocks of the function.
	Blocks() map[bin.Address]Block
}

// Block represents a basic block, which consists of a sequence of non-
// branching instructions terminated by a control flow instruction.
type Block interface {
	// Addr returns the address of the basic block.
	Addr() bin.Address
	// Insts returns the non-branching instructions of the basic block.
	Insts() []Instruction
	// Term returns the terminating instruction of the basic block.
	Term() Terminator
}

// Instruction represents a non-branching instruction.
type Instruction interface {
	// Addr returns the address of the instruction.
	Addr() bin.Address
	// Parent returns the parent basic block of the instruction.
	Parent() Block
}

// Terminator represents a terminating instruction.
type Terminator interface {
	Instruction
	// Successors returns the successor basic blocks of the terminating
	// instruction.
	Successors() []Block
}
