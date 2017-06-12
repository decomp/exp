package x86

import (
	"fmt"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"golang.org/x/arch/x86/x86asm"
)

// === [ arguments ] ===========================================================

// Arg returns the i:th argument of the instruction.
func (inst *Inst) Arg(i int) *Arg {
	return NewArg(inst.Args[i], inst)
}

// Reg returns the register at the i:th argument of the instruction.
func (inst *Inst) Reg(i int) *Reg {
	return NewReg(inst.Args[i], inst)
}

// Mem returns the memory reference at the i:th argument of the instruction.
func (inst *Inst) Mem(i int) *Mem {
	return NewMem(inst.Args[i], inst)
}

// --- [ argument ] ------------------------------------------------------------

// An Arg is a single x86 instruction argument.
type Arg struct {
	// x86 instruction argument.
	x86asm.Arg
	// Parent instruction; used to calculate relative offsets and retrieve
	// symbolic execution information.
	Parent *Inst
}

// NewArg returns a new x86 argument with the given parent instruction.
func NewArg(arg x86asm.Arg, parent *Inst) *Arg {
	return &Arg{
		Arg:    arg,
		Parent: parent,
	}
}

// --- [ register ] ------------------------------------------------------------

// A Reg is a single x86 register.
type Reg struct {
	// x86 register.
	x86asm.Reg
	// Parent instruction; used to retrieve symbolic execution information.
	Parent *Inst
}

// NewReg returns a new x86 register argument with the given parent instruction.
func NewReg(arg x86asm.Arg, parent *Inst) *Reg {
	reg, ok := arg.(x86asm.Reg)
	if !ok {
		panic(fmt.Errorf("invalid register argument type; expected x86asm.Reg, got %T", arg))
	}
	return &Reg{
		Reg:    reg,
		Parent: parent,
	}
}

// --- [ memory reference ] ----------------------------------------------------

// A Mem is a memory reference.
type Mem struct {
	// x86 memory reference.
	x86asm.Mem
	// Parent instruction; used to retrieve symbolic execution information.
	Parent *Inst
}

// NewMem returns a new memory reference argument with the given parent
// instruction.
func NewMem(arg x86asm.Arg, parent *Inst) *Mem {
	mem, ok := arg.(x86asm.Mem)
	if !ok {
		panic(fmt.Errorf("invalid memory reference argument type; expected x86asm.Mem, got %T", arg))
	}
	return &Mem{
		Mem:    mem,
		Parent: parent,
	}
}

// Segment returns the segment register of the memory reference.
func (mem *Mem) Segment() *Reg {
	return NewReg(mem.Mem.Segment, mem.Parent)
}

// Base returns the base register of the memory reference.
func (mem *Mem) Base() *Reg {
	return NewReg(mem.Mem.Base, mem.Parent)
}

// Index returns the index register of the memory reference.
func (mem *Mem) Index() *Reg {
	return NewReg(mem.Mem.Index, mem.Parent)
}

// ### [ Helper functions ] ####################################################

// Addrs returns the addresses specified by the given argument. Addr specifies
// the address of the terminator, and next the address of the next instruction.
func (dis *Disasm) Addrs(arg x86asm.Arg, addr, next bin.Address) []bin.Address {
	switch arg := arg.(type) {
	//case x86asm.Reg:
	case x86asm.Mem:
		// Segment:[Base+Scale*Index+Disp].

		//    Segment Reg
		//    Base    Reg
		//    Scale   uint8
		//    Index   Reg
		//    Disp    int64

		// Static target.
		disp := bin.Address(arg.Disp)
		if arg.Segment == 0 && arg.Base == 0 && arg.Index == 0 {
			return []bin.Address{disp}
		}

		// Adjust disposition based on index register value.
		if arg.Index != 0 {
			if context, ok := dis.Contexts[addr]; ok {
				if c, ok := context.Regs[Register(arg.Index)]; ok {
					if indexMin, ok := c["min"]; ok {
						disp += bin.Address(arg.Scale) * indexMin.Addr()
					}
				}
			}
		}

		// Jump table target.
		if targets, ok := dis.Tables[disp]; ok {
			if arg.Segment == 0 && arg.Base == 0 && arg.Scale == 4 && arg.Index != 0 {
				return targets
			}
		}

		// TODO: Figure out how to handle indirect jump to function pointer.

		// Target is likely a function pointer; skip for now.
		if arg.Base != 0 && disp < dis.codeStart() {
			warn.Printf("ignoring indirect targets from %v of memory reference %v", addr, arg)
			return nil
		}

		pretty.Println("mem:", arg)
		panic("x86.Disasm.Addrs: not yet implemented")
	//case x86asm.Imm:
	case x86asm.Rel:
		target := next + bin.Address(arg)
		return []bin.Address{target}
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}
