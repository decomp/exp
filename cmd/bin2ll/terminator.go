package main

import (
	"fmt"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
	"golang.org/x/arch/x86/x86asm"
)

// emitTerm translates the given x86 terminator to LLVM IR, emitting code to f.
func (f *Func) emitTerm(term *Inst) error {
	// Handle implicit fallthrough terminators.
	if term.isImplicit() {
		dbg.Printf("lifting implicit terminator: JMP %v", term.addr)
		next, ok := f.blocks[term.addr]
		if !ok {
			return errors.Errorf("unable to locate basic block at %v", term.addr)
		}
		f.cur.NewBr(next)
		return nil
	}

	dbg.Println("lifting terminator:", term.Inst)

	// Check if prefix is present.
	for _, prefix := range term.Prefix[:] {
		// The first zero in the array marks the end of the prefixes.
		if prefix == 0 {
			break
		}
		switch prefix {
		case x86asm.PrefixData16, x86asm.PrefixData16 | x86asm.PrefixImplicit:
			// prefix already supported.
		default:
			pretty.Println("terminator with prefix:", term)
			panic(fmt.Errorf("support for %v terminator with prefix not yet implemented", term.Op))
		}
	}

	// Translate terminator.
	switch term.Op {
	// Loop terminators.
	case x86asm.LOOP:
		return f.emitTermLOOP(term)
	case x86asm.LOOPE:
		return f.emitTermLOOPE(term)
	case x86asm.LOOPNE:
		return f.emitTermLOOPNE(term)
	// Conditional jump terminators.
	case x86asm.JA:
		return f.emitTermJA(term)
	case x86asm.JAE:
		return f.emitTermJAE(term)
	case x86asm.JB:
		return f.emitTermJB(term)
	case x86asm.JBE:
		return f.emitTermJBE(term)
	case x86asm.JCXZ:
		return f.emitTermJCXZ(term)
	case x86asm.JE:
		return f.emitTermJE(term)
	case x86asm.JECXZ:
		return f.emitTermJECXZ(term)
	case x86asm.JG:
		return f.emitTermJG(term)
	case x86asm.JGE:
		return f.emitTermJGE(term)
	case x86asm.JL:
		return f.emitTermJL(term)
	case x86asm.JLE:
		return f.emitTermJLE(term)
	case x86asm.JNE:
		return f.emitTermJNE(term)
	case x86asm.JNO:
		return f.emitTermJNO(term)
	case x86asm.JNP:
		return f.emitTermJNP(term)
	case x86asm.JNS:
		return f.emitTermJNS(term)
	case x86asm.JO:
		return f.emitTermJO(term)
	case x86asm.JP:
		return f.emitTermJP(term)
	case x86asm.JRCXZ:
		return f.emitTermJRCXZ(term)
	case x86asm.JS:
		return f.emitTermJS(term)
	// Unconditional jump terminators.
	case x86asm.JMP:
		return f.emitTermJMP(term)
	// Return terminators.
	case x86asm.RET:
		return f.emitTermRET(term)
	default:
		panic(fmt.Errorf("support for x86 terminator opcode %v not yet implemented", term.Op))
	}
}

// --- [ JMP ] -----------------------------------------------------------------

// emitTermJMP translates the given x86 JMP terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermJMP(term *Inst) error {
	// Handle tail calls.
	if f.isTailCall(term) {
		// Hack: interpret the JMP instruction as a CALL instruction. This works
		// since emitInstCALL only interprets inst.Args[0], which is the same in
		// both JMP and CALL instructions.
		if err := f.emitInstCALL(term); err != nil {
			return errors.WithStack(err)
		}
		// Handle return values.
		if !types.Equal(f.Sig.Ret, types.Void) {
			// Non-void functions, pass return value in EAX.
			result := f.useReg(EAX)
			f.cur.NewRet(result)
			return nil
		}
		f.cur.NewRet(nil)
		return nil
	}

	// Handle static jump.
	arg := term.Arg(0)
	if targetAddr, ok := f.getAddr(arg); ok {
		target, ok := f.blocks[targetAddr]
		if !ok {
			return errors.Errorf("unable to locate target basic block at %v", targetAddr)
		}
		f.cur.NewBr(target)
		return nil
	}
	// Handle jump tables.
	if _, ok := arg.Arg.(x86asm.Mem); ok {
		mem := term.Mem(0)
		if targetAddrs, ok := f.d.tables[bin.Address(mem.Disp)]; ok {
			// TODO: Implement proper support for jump table translation. The
			// current implementation makes a range of assumptions, which do not
			// hold true in the general case; e.g. assuming that mem.Base == 0 && mem.Scale == 4.
			if mem.Mem.Base != 0 {
				panic("support for jump table memory reference with base register not yet implemented")
			}
			if mem.Scale != 4 {
				panic(fmt.Errorf("support for jump table memory reference with scale %d not yet implemented", mem.Scale))
			}

			// TODO: Locate default target using information from symbolic
			// execution and predecessor basic blocks.

			// At this stage of recovery, the assumption is `index` is always
			// within the range of the jump table offsets. Thus, the default branch
			// is always unreachable.
			//
			// This assumption will be validated and revisited when information
			// from symbolic execution is available.

			// TODO: Add support for indirect jump tables; i.e.
			//
			//    targets[values[index]]
			index := f.useReg(mem.Index())
			unreachable := &ir.BasicBlock{}
			unreachable.NewUnreachable()
			f.AppendBlock(unreachable)
			targetDefault := unreachable
			var cases []*ir.Case
			for i, targetAddr := range targetAddrs {
				target, ok := f.blocks[targetAddr]
				if !ok {
					return errors.Errorf("unable to locate basic block at %v", targetAddr)
				}
				ii := constant.NewInt(int64(i), index.Type())
				c := ir.NewCase(ii, target)
				cases = append(cases, c)
			}
			f.cur.NewSwitch(index, targetDefault, cases...)
			return nil
		}
	}
	pretty.Println("term:", term)
	panic("emitTermJMP: not yet implemented")
}

// --- [ RET ] -----------------------------------------------------------------

// emitTermRET translates the given x86 RET terminator to LLVM IR, emitting code
// to f.
func (f *Func) emitTermRET(term *Inst) error {
	// Handle return values of non-void functions (passed through EAX).
	if !types.Equal(f.Sig.Ret, types.Void) {
		result := f.useReg(EAX)
		f.cur.NewRet(result)
		return nil
	}
	f.cur.NewRet(nil)
	return nil
}

// === [ Helper functions ] ====================================================

// isImplicit reports whether term is an implicit terminator. Implicit
// terminators are used when a basic block is missing a terminator and falls
// through into the succeeding basic block, the address of which is denoted by
// term.addr.
func (term *Inst) isImplicit() bool {
	zero := x86asm.Inst{}
	return term.Inst == zero
}

// isTailCall reports whether the given instruction is a tail call instruction.
func (f *Func) isTailCall(inst *Inst) bool {
	arg := inst.Arg(0)
	if target, ok := f.getAddr(arg); ok {
		if f.contains(target) {
			return false
		}
		if !f.d.isFunc(target) {
			fmt.Println("arg:", arg)
			pretty.Println(arg)
			panic(fmt.Errorf("tail call to non-function address %v", target))
		}
		return true
	}
	// Target read from jump table (e.g. switch statement).
	if mem, ok := arg.Arg.(x86asm.Mem); ok {
		addr := bin.Address(mem.Disp)
		if targets, ok := f.d.tables[addr]; ok {
			for _, target := range targets {
				if !f.contains(target) {
					if !f.d.isFunc(target) {
						fmt.Println("arg:", arg)
						pretty.Println(arg)
						panic(fmt.Errorf("tail call to non-function address %v", target))
					}
					return true
				}
			}
			return false
		}
	}
	fmt.Println("arg:", arg)
	pretty.Println(arg)
	panic("not yet implemented")
}

// contains reports whether the target address is part of the address space of
// the function.
func (f *Func) contains(target bin.Address) bool {
	// Target inside function address range.
	funcEnd := f.d.getFuncEndAddr(f.entry)
	if f.entry <= target && target < funcEnd {
		return true
	}
	// Target inside function chunk.
	if funcAddr, ok := f.d.chunkFunc[target]; ok {
		if funcAddr == f.entry {
			return true
		}
	}
	// Target is an imported function.
	if f.d.isImport(target) {
		return false
	}
	return false
}
