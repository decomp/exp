package x86

import (
	"fmt"
	"sort"

	"github.com/decomp/exp/bin"
	"golang.org/x/arch/x86/x86asm"
)

// String returns the string representation of the instruction.
func (inst *Inst) String() string {
	if inst.isDummyTerm() {
		return fmt.Sprintf("; fallthrough %v", inst.Addr)
	}
	return inst.Inst.String()
}

// isTerm reports whether the given instruction is a terminating instruction.
func (term *Inst) isTerm() bool {
	switch term.Op {
	// Loop terminators.
	case x86asm.LOOP, x86asm.LOOPE, x86asm.LOOPNE:
		return true
	// Conditional jump terminators.
	case x86asm.JA, x86asm.JAE, x86asm.JB, x86asm.JBE, x86asm.JCXZ, x86asm.JE, x86asm.JECXZ, x86asm.JG, x86asm.JGE, x86asm.JL, x86asm.JLE, x86asm.JNE, x86asm.JNO, x86asm.JNP, x86asm.JNS, x86asm.JO, x86asm.JP, x86asm.JRCXZ, x86asm.JS:
		return true
	// Unconditional jump terminators.
	case x86asm.JMP:
		return true
	// Return terminators.
	case x86asm.RET:
		return true
	}
	return false
}

// isDummyTerm reports whether the given instruction is a dummy terminating
// instruction. Dummy terminators are used when a basic block is missing a
// terminator and falls through into the succeeding basic block, the address of
// which is denoted by term.Addr.
func (term *Inst) isDummyTerm() bool {
	zero := x86asm.Inst{}
	return term.Inst == zero
}

// Targets returns the targets of the given terminator instruction. Entry
// denotes the entry address of the function containing the terminator
// instruction.
func (dis *Disasm) Targets(term *Inst, funcEntry bin.Address) []bin.Address {
	if term.isDummyTerm() {
		// Dummy terminator; fall through into the succeeding basic block, the
		// address of which is denoted by term.Addr.
		return []bin.Address{term.Addr}
	}
	next := term.Addr + bin.Address(term.Len)
	switch term.Op {
	// Loop terminators.
	case x86asm.LOOP, x86asm.LOOPE, x86asm.LOOPNE:
		targets := dis.Addrs(term.Args[0], term.Addr, next)
		return append(targets, next)
	// Conditional jump terminators.
	case x86asm.JA, x86asm.JAE, x86asm.JB, x86asm.JBE, x86asm.JCXZ, x86asm.JE, x86asm.JECXZ, x86asm.JG, x86asm.JGE, x86asm.JL, x86asm.JLE, x86asm.JNE, x86asm.JNO, x86asm.JNP, x86asm.JNS, x86asm.JO, x86asm.JP, x86asm.JRCXZ, x86asm.JS:
		targets := dis.Addrs(term.Args[0], term.Addr, next)
		return append(targets, next)
	// Unconditional jump terminators.
	case x86asm.JMP:
		preTargets := dis.Addrs(term.Args[0], term.Addr, next)
		var targets []bin.Address
		for _, target := range preTargets {
			if dis.isTailCall(funcEntry, target) {
				dbg.Printf("tail call at %v", term.Addr)
			} else {
				// Append target if not part of a tail call.
				targets = append(targets, target)
			}
		}
		return targets
	// Return terminators.
	case x86asm.RET:
		// no targets.
		return nil
	}
	panic(fmt.Errorf("support for terminator instruction %v not yet implemented", term.Op))
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
