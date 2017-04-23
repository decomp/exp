package main

import (
	"github.com/decomp/exp/bin"
	"golang.org/x/arch/x86/x86asm"
)

// Context maps from CPU register to the value it holds.
type Context map[x86asm.Reg]uint64

// Contexts tracks the CPU context at various addresses of the executable.
type Contexts map[bin.Address]Context
