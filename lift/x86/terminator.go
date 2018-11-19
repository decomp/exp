package x86

import (
	"fmt"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/decomp/exp/disasm/x86"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
	"golang.org/x/arch/x86/x86asm"
)

// liftTerm lifts the given x86 terminator to LLVM IR, emitting code to f.
func (f *Func) liftTerm(term *x86.Inst) error {
	// Handle implicit fallthrough terminators.
	if term.IsDummyTerm() {
		dbg.Printf("lifting implicit terminator: JMP %v", term.Addr)
		next, ok := f.blocks[term.Addr]
		if !ok {
			return errors.Errorf("unable to locate basic block at %v", term.Addr)
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
		return f.liftTermLOOP(term)
	case x86asm.LOOPE:
		return f.liftTermLOOPE(term)
	case x86asm.LOOPNE:
		return f.liftTermLOOPNE(term)
	// Conditional jump terminators.
	case x86asm.JA:
		return f.liftTermJA(term)
	case x86asm.JAE:
		return f.liftTermJAE(term)
	case x86asm.JB:
		return f.liftTermJB(term)
	case x86asm.JBE:
		return f.liftTermJBE(term)
	case x86asm.JCXZ:
		return f.liftTermJCXZ(term)
	case x86asm.JE:
		return f.liftTermJE(term)
	case x86asm.JECXZ:
		return f.liftTermJECXZ(term)
	case x86asm.JG:
		return f.liftTermJG(term)
	case x86asm.JGE:
		return f.liftTermJGE(term)
	case x86asm.JL:
		return f.liftTermJL(term)
	case x86asm.JLE:
		return f.liftTermJLE(term)
	case x86asm.JNE:
		return f.liftTermJNE(term)
	case x86asm.JNO:
		return f.liftTermJNO(term)
	case x86asm.JNP:
		return f.liftTermJNP(term)
	case x86asm.JNS:
		return f.liftTermJNS(term)
	case x86asm.JO:
		return f.liftTermJO(term)
	case x86asm.JP:
		return f.liftTermJP(term)
	case x86asm.JRCXZ:
		return f.liftTermJRCXZ(term)
	case x86asm.JS:
		return f.liftTermJS(term)
	// Unconditional jump terminators.
	case x86asm.JMP:
		return f.liftTermJMP(term)
	// Return terminators.
	case x86asm.RET:
		return f.liftTermRET(term)
	default:
		panic(fmt.Errorf("support for x86 terminator opcode %v not yet implemented", term.Op))
	}
}

// --- [ JMP ] -----------------------------------------------------------------

// liftTermJMP lifts the given x86 JMP terminator to LLVM IR, emitting code to
// f.
func (f *Func) liftTermJMP(term *x86.Inst) error {
	// Handle tail calls.
	if f.isTailCall(term) {
		// Hack: interpret the JMP instruction as a CALL instruction. This works
		// since emitInstCALL only interprets inst.Args[0], which is the same in
		// both JMP and CALL instructions.
		if err := f.liftInstCALL(term); err != nil {
			return errors.WithStack(err)
		}
		// Handle return values.
		if !types.Equal(f.Sig.RetType, types.Void) {
			// Non-void functions, pass return value in EAX.
			result := f.useReg(x86.EAX)
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
		if targetAddrs, ok := f.l.Tables[bin.Address(mem.Disp)]; ok {
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
			f.Blocks = append(f.Blocks, unreachable)
			targetDefault := unreachable
			var cases []*ir.Case
			for i, targetAddr := range targetAddrs {
				target, ok := f.blocks[targetAddr]
				if !ok {
					return errors.Errorf("unable to locate basic block at %v", targetAddr)
				}
				ii := constant.NewInt(index.Type().(*types.IntType), int64(i))
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

// liftTermRET lifts the given x86 RET terminator to LLVM IR, emitting code to
// f.
func (f *Func) liftTermRET(term *x86.Inst) error {
	// Handle return values of non-void functions (passed through EAX).
	if !types.Equal(f.Sig.RetType, types.Void) {
		result := f.useReg(x86.EAX)
		f.cur.NewRet(result)
		return nil
	}
	f.cur.NewRet(nil)
	return nil
}

// === [ Helper functions ] ====================================================

// isTailCall reports whether the given instruction is a tail call instruction.
func (f *Func) isTailCall(inst *x86.Inst) bool {
	arg := inst.Arg(0)
	if target, ok := f.getAddr(arg); ok {
		if f.contains(target) {
			return false
		}
		if !f.l.IsFunc(target) {
			fmt.Println("arg:", arg)
			pretty.Println(arg)
			panic(fmt.Errorf("tail call to non-function address %v", target))
		}
		return true
	}
	// Target read from jump table (e.g. switch statement).
	if mem, ok := arg.Arg.(x86asm.Mem); ok {
		addr := bin.Address(mem.Disp)
		if targets, ok := f.l.Tables[addr]; ok {
			for _, target := range targets {
				if !f.contains(target) {
					if !f.l.IsFunc(target) {
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

	// TODO: Find a prettier solution for handling indirect jumps to potential
	// tail call functions at register relative memory locations; e.g.
	//    JMP [EAX+0x8]

	// HACK: set the current basic block to a dummy basic block so that we may
	// invoke f.getFunc (which may emit load instructions) to figure out if we
	// are jumping to a function.
	cur := f.cur
	dummy := &ir.BasicBlock{}
	f.cur = dummy
	_, _, _, ok := f.getFunc(arg)
	f.cur = cur
	if ok {
		return true
	}

	fmt.Println("arg:", arg)
	pretty.Println(arg)
	panic("not yet implemented")
}

// contains reports whether the target address is part of the address space of
// the function.
func (f *Func) contains(target bin.Address) bool {
	// Target inside function address range.
	entry := f.AsmFunc.Addr
	funcEnd := f.l.getFuncEndAddr(entry)
	if entry <= target && target < funcEnd {
		return true
	}
	// Target inside function chunk.
	if funcAddr, ok := f.l.Chunks[target]; ok {
		if funcAddr[entry] {
			return true
		}
	}
	// Target is an imported function.
	if _, ok := f.l.File.Imports[target]; ok {
		return false
	}
	return false
}

// getFuncEndAddr returns the end address of the given function.
func (l *Lifter) getFuncEndAddr(entry bin.Address) bin.Address {
	less := func(i int) bool {
		return entry < l.FuncAddrs[i]
	}
	index := sort.Search(len(l.FuncAddrs), less)
	if index < len(l.FuncAddrs) {
		return l.FuncAddrs[index]
	}
	return l.getCodeEnd()
}

//// getCodeStart returns the start address of the code section.
//func (l *Lifter) getCodeStart() bin.Address {
//	return bin.Address(d.imageBase + d.codeBase)
//}

// getCodeEnd returns the end address of the code section.
func (l *Lifter) getCodeEnd() bin.Address {
	var max bin.Address
	for _, sect := range l.File.Sections {
		if sect.Perm&bin.PermX != 0 {
			end := sect.Addr + bin.Address(len(sect.Data))
			if max < end {
				max = end
			}
		}
	}
	if max == 0 {
		panic("unable to locate end of code segment")
	}
	return max
}
