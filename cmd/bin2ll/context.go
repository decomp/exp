package main

import (
	"github.com/decomp/exp/bin"
)

// Contexts tracks the CPU context at various addresses of the executable.
type Contexts map[bin.Address]Context

// Context tracks the CPU context at a specific address of the executable.
type Context struct {
	// Register constraints.
	Regs map[Register]ValueContext `json:"regs"`
	// Instruction argument constraints.
	Args map[int]ValueContext `json:"args"`
}

// ValueContext defines constraints on a value used at a specific address.
//
// The following keys are defined.
//
//    addr         virtual address.
//
//    val          value.
//
//    min          minimum value.
//    max          maximum value.
//
//    Mem.offset   memory reference offset.
type ValueContext map[string]bin.Uint64
