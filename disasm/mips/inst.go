package mips

import (
	"fmt"
	"sort"

	"github.com/decomp/exp/bin"
)

// MIPS instruction length in bytes.
const mipsInstLen = 4

// MIPS $ra register index.
const mipsRegRA = 31

// String returns the string representation of the instruction.
func (inst *Inst) String() string {
	if inst.IsDummyTerm() {
		return fmt.Sprintf("; fallthrough %v", inst.Addr)
	}
	line, err := inst.Render()
	if err != nil {
		panic(fmt.Errorf("unable to pretty-print instruction; %v", err))
	}
	return line.String()
}

// isTerm reports whether the given instruction is a terminating instruction.
func (inst *Inst) isTerm() bool {
	switch inst.Name {
	// Conditional branch instructions.
	case "BEQ", "BGEZ", "BGTZ", "BLEZ", "BLTZ", "BNE":
		return true
	// Unconditional jump instructions.
	case "J", "JAL":
		return true
	// Unconditional indirect jump instructions.
	case "JALR", "JR":
		return true
	}
	return false
}

// IsDummyTerm reports whether the given instruction is a dummy terminating
// instruction. Dummy terminators are used when a basic block is missing a
// terminator and falls through into the succeeding basic block, the address of
// which is denoted by inst.Addr.
func (inst *Inst) IsDummyTerm() bool {
	return inst.Instruction == nil
}

// Targets returns the targets of the given terminator instruction. Entry
// denotes the entry address of the function containing the terminator
// instruction.
func (dis *Disasm) Targets(term *Inst, funcEntry bin.Address) []bin.Address {
	if term.IsDummyTerm() {
		// Dummy terminator; fall through into the succeeding basic block, the
		// address of which is denoted by term.Addr.
		return []bin.Address{term.Addr}
	}
	next := term.Addr + mipsInstLen
	switch term.Name {
	// Conditional branch instructions.
	case "BEQ", "BGEZ", "BGTZ", "BLEZ", "BLTZ", "BNE":
		var targets []bin.Address
		cp := term.CodePointer
		if cp.IsSymbol {
			// TODO: Add support for symbol code pointers.
			panic(fmt.Errorf("support for terminators with symbol code pointers not yet implemented; %v", term))
		}
		if cp.Absolute {
			target := term.Addr&0xF0000000 | bin.Address(cp.Constant)
			targets = append(targets, target)
		} else {
			// Relative target.
			var target bin.Address
			switch dis.Mode {
			case 32:
				target = bin.Address(uint32(term.Addr) + cp.Constant + mipsInstLen)
			case 64:
				target = term.Addr + bin.Address(cp.Constant) + mipsInstLen
			default:
				panic(fmt.Errorf("support for CPU mode %d not yet implemented", dis.Mode))
			}
			targets = append(targets, target)
		}
		targets = append(targets, next)
		return targets
	// Unconditional jump instructions.
	case "J", "JAL":
		var targets []bin.Address
		cp := term.CodePointer
		if cp.IsSymbol {
			// TODO: Add support for symbol code pointers.
			panic(fmt.Errorf("support for terminators with symbol code pointers not yet implemented; %v", term))
			// Return terminators.
			//if cp.Symbol == "$ra" {
			//	// no targets.
			//	return nil
			//}
		}
		if cp.Absolute {
			target := term.Addr&0xF0000000 | bin.Address(cp.Constant)
			targets = append(targets, target)
		} else {
			// Relative target.
			target := term.Addr + bin.Address(cp.Constant) + mipsInstLen
			targets = append(targets, target)
		}
		return targets
	// Unconditional indirect jump instructions.
	case "JALR", "JR":
		reg := term.Registers[len(term.Registers)-1]
		fmt.Println("term:", term)
		fmt.Println("   reg:", reg)
		if reg == mipsRegRA {
			return nil
		}
		// TODO: Handle indirect jumps. Note, the $ra return register has no
		// targets. Other registers need context information.
		//panic(fmt.Errorf("support for indirect jump to register %v not yet implemented; %v", reg, term))
		return nil // TODO: Remove, once context is implemented.
	}
	panic(fmt.Errorf("support for terminator instruction %v not yet implemented", term.Name))
}

// isTailCall reports whether the given JMP instruction is a tail call
// instruction.
func (dis *Disasm) isTailCall(funcEntry bin.Address, target bin.Address) bool {
	funcEnd := dis.funcEnd(funcEntry)
	if funcEntry <= target && target < funcEnd {
		// Target inside function body.
		return false
	}
	if funcAddr, ok := dis.Chunks[target]; ok {
		if funcAddr == funcEntry {
			// Target part of function chunk.
			return false
		}
	}
	if !dis.IsFunc(target) {
		panic(fmt.Errorf("tail call to non-function address %v from function at %v.\n\ttip: %v may be a function chunk of the parent function at %v\n\tadd to lst:  FUNCTION CHUNK AT .text:%08X\n\tadd to json: %q: %q,", target, funcEntry, target, funcEntry, uint64(target), target, funcEntry))
	}
	// Target is a tail call.
	return true
}

// funcEnd returns the end address of the function, under the assumption that
// the function is continuous.
func (dis *Disasm) funcEnd(funcEntry bin.Address) bin.Address {
	less := func(i int) bool {
		return funcEntry < dis.FuncAddrs[i]
	}
	index := sort.Search(len(dis.FuncAddrs), less)
	if 0 <= index && index < len(dis.FuncAddrs) {
		return dis.FuncAddrs[index]
	}
	return dis.codeEnd()
}
