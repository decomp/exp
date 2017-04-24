package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
	"golang.org/x/arch/x86/x86asm"
)

// translateFunc translates the given function from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) translateFunc(f *Func) error {
	for addr := range f.bbs {
		label := fmt.Sprintf("block_%06X", uint64(addr))
		block := &ir.BasicBlock{
			Name: label,
		}
		f.blocks[addr] = block
	}

	if f.Function == nil {
		// TODO: Add proper support for type signatures once type analysis has
		// been conducted.
		name := fmt.Sprintf("f_%06X", uint64(f.entry))
		sig := types.NewFunc(types.Void)
		typ := types.NewPointer(sig)
		f.Function = &ir.Function{
			Name: name,
			Typ:  typ,
			Sig:  sig,
			Metadata: map[string]*metadata.Metadata{
				"addr": &metadata.Metadata{
					Nodes: []metadata.Node{&metadata.String{Val: f.entry.String()}},
				},
			},
		}
	}
	dbg.Printf("translating function %q at %v", f.Name, f.entry)

	var blockAddrs []bin.Address
	for _, bb := range f.bbs {
		blockAddrs = append(blockAddrs, bb.addr)
	}
	sort.Sort(bin.Addresses(blockAddrs))
	if len(blockAddrs) == 0 {
		return errors.New("invalid function definition; missing function body")
	}
	for _, blockAddr := range blockAddrs {
		bb := f.bbs[blockAddr]
		if err := d.translateBlock(f, bb); err != nil {
			return errors.WithStack(err)
		}
	}

	// Add new entry basic block to define registers and status flags used within
	// the function.
	if len(f.regs) > 0 || len(f.statusFlags) > 0 {
		entry := &ir.BasicBlock{}
		// Allocate local variables for each register used within the function.
		for reg := x86asm.AL; reg <= x86asm.TR7; reg++ {
			if inst, ok := f.regs[reg]; ok {
				entry.AppendInst(inst)
			}
		}
		// Allocate local variables for each status flag used within the function.
		for status := CF; status <= OF; status++ {
			if inst, ok := f.statusFlags[status]; ok {
				entry.AppendInst(inst)
			}
		}
		// Allocate local variables for each local variable used within the
		// function.
		var names []string
		for name := range f.locals {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			inst := f.locals[name]
			entry.AppendInst(inst)
		}
		// Handle calling conventions.
		f.cur = entry
		// TODO: Initialize parameter initialization in entry block prior to basic
		// block translation. Move this code to before f.translateBlock, and remove
		// f.espDisp = 0.
		f.espDisp = 0
		for i, param := range f.Sig.Params {
			// Use parameter in register.
			switch f.CallConv {
			case ir.CallConvX86_FastCall:
				switch i {
				case 0:
					f.defReg(ECX, param)
					continue
				case 1:
					f.defReg(EDX, param)
					continue
				}
			default:
				// TODO: Add support for more calling conventions.
			}
			// Use parameter on stack.
			m := x86asm.Mem{
				Base: x86asm.ESP,
				Disp: 4,
			}
			mem := NewMem(m, nil)
			f.defMem(mem, param)
		}
		target := f.Blocks[0]
		entry.NewBr(target)
		f.Blocks = append([]*ir.BasicBlock{entry}, f.Blocks...)
	}
	return nil
}

// translateBlock translates the given basic block from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) translateBlock(f *Func, bb *BasicBlock) error {
	block, ok := f.blocks[bb.addr]
	if !ok {
		return errors.Errorf("unable to locate LLVM IR basic block at %v", bb.addr)
	}
	f.AppendBlock(block)
	f.cur = block
	dbg.Printf("translating basic block at %v", bb.addr)
	for _, inst := range bb.insts {
		if err := f.emitInst(inst); err != nil {
			return errors.WithStack(err)
		}
	}
	// Translate terminator.
	if err := d.translateTerm(f, bb, bb.term); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// instADD translates the given ADD instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instADD(f *Func, bb *BasicBlock, inst *Inst) error {
	x := d.useArg(f, bb, inst, inst.Args[0])
	y := d.useArg(f, bb, inst, inst.Args[1])
	result := f.cur.NewAdd(x, y)
	d.defArg(f, bb, inst, inst.Args[0], result)
	return nil
}

// instAND translates the given AND instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instAND(f *Func, bb *BasicBlock, inst *Inst) error {
	x := d.useArg(f, bb, inst, inst.Args[0])
	y := d.useArg(f, bb, inst, inst.Args[1])
	result := f.cur.NewAnd(x, y)
	d.defArg(f, bb, inst, inst.Args[0], result)
	return nil
}

// instCALL translates the given CALL instruction from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) instCALL(f *Func, bb *BasicBlock, inst *Inst) error {
	var callee *Func
	if addr, ok := d.getAddr(f, bb, inst, inst.Args[0]); ok {
		if c, ok := d.funcs[addr]; ok {
			callee = c
		} else {
			return errors.Errorf("unable to locate function at %v", addr)
		}
	} else {
		panic(fmt.Errorf("unknown callee address from argument `%v` in call instruction at %v", inst.Args[0], inst.addr))
		c := d.useArg(f, bb, inst, inst.Args[0])
		// TODO: Add support for value.Named callees. Using *ir.Function for now, to
		// gain access to the calling convention of the function. Data flow and type
		// analysis will provide this information in the future also for local
		// function pointer callees.
		var ok bool
		callee, ok = c.(*Func)
		if !ok {
			return errors.Errorf("invalid callee type; expected *main.function, got %T", c)
		}
	}
	var args []value.Value
	switch callee.CallConv {
	case ir.CallConvX86_FastCall:
		params := callee.Sig.Params
		fmt.Println("params:", params)
		if len(params) > 0 {
			arg := d.useArg(f, bb, nil, x86asm.ECX)
			args = append(args, arg)
		}
		if len(params) > 1 {
			arg := d.useArg(f, bb, nil, x86asm.EDX)
			args = append(args, arg)
		}
	default:
		// TODO: Handle call arguments.
	}
	result := f.cur.NewCall(callee, args...)
	// Handle return values of non-void callees (passed through EAX).
	fmt.Println("call result type:", callee.Sig.Ret)
	if !types.Equal(callee.Sig.Ret, types.Void) {
		d.defArg(f, bb, nil, x86asm.EAX, result)
	}
	return nil
}

// instCDQ translates the given CDQ instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instCDQ(f *Func, bb *BasicBlock, inst *Inst) error {
	// EDX:EAX = sign-extend of EAX.
	eax := d.useReg(f, x86asm.EAX)
	tmp := f.cur.NewLShr(eax, constant.NewInt(31, types.I32))
	cond := f.cur.NewTrunc(tmp, types.I1)
	targetTrue := &ir.BasicBlock{}
	targetFalse := &ir.BasicBlock{}
	exit := &ir.BasicBlock{}
	f.AppendBlock(targetTrue)
	f.AppendBlock(targetFalse)
	f.AppendBlock(exit)
	f.cur.NewCondBr(cond, targetTrue, targetFalse)
	f.cur = targetTrue
	d.defReg(f, x86asm.EDX, constant.NewInt(0xFFFFFFFF, types.I32))
	f.cur = targetFalse
	d.defReg(f, x86asm.EDX, constant.NewInt(0, types.I32))
	targetTrue.NewBr(exit)
	targetFalse.NewBr(exit)
	f.cur = exit
	return nil
}

// instCMP translates the given CMP instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instCMP(f *Func, bb *BasicBlock, inst *Inst) error {
	x := d.useArg(f, bb, inst, inst.Args[0])
	y := d.useArg(f, bb, inst, inst.Args[1])
	// Set the status flags according to the result.

	// TODO: Fix calculation of status flags. Pass SUB from CMP instruction and
	// AND from TEST instruction.
	return d.updateStatusFlags(f, bb, x, y)
}

// instDEC translates the given DEC instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instDEC(f *Func, bb *BasicBlock, inst *Inst) error {
	x := d.useArg(f, bb, inst, inst.Args[0])
	one := constant.NewInt(1, types.I32)
	result := f.cur.NewSub(x, one)
	d.defArg(f, bb, inst, inst.Args[0], result)
	return nil
}

// instIMUL translates the given IMUL instruction from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) instIMUL(f *Func, bb *BasicBlock, inst *Inst) error {
	x := d.useArg(f, bb, inst, inst.Args[1])
	y := d.useArg(f, bb, inst, inst.Args[2])
	result := f.cur.NewMul(x, y)
	d.defArg(f, bb, inst, inst.Args[0], result)
	return nil
}

// instINC translates the given INC instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instINC(f *Func, bb *BasicBlock, inst *Inst) error {
	x := d.useArg(f, bb, inst, inst.Args[0])
	one := constant.NewInt(1, types.I32)
	result := f.cur.NewAdd(x, one)
	d.defArg(f, bb, inst, inst.Args[0], result)
	return nil
}

// instLEA translates the given LEA instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instLEA(f *Func, bb *BasicBlock, inst *Inst) error {
	y, ok := inst.Args[1].(x86asm.Mem)
	if !ok {
		return errors.Errorf("invalid LEA operand type; expected x86asm.Mem, got %T", inst.Args[1])
	}
	result := d.mem(f, bb, inst, y)
	d.defArg(f, bb, inst, inst.Args[0], result)
	return nil
}

// instMOV translates the given MOV instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instMOV(f *Func, bb *BasicBlock, inst *Inst) error {
	y := d.useArg(f, bb, inst, inst.Args[1])
	d.defArg(f, bb, inst, inst.Args[0], y)
	return nil
}

// instMOVSB translates the given MOVSB instruction from x86 machine code to
// LLVM IR assembly.
func (d *disassembler) instMOVSB(f *Func, bb *BasicBlock, inst *Inst) error {
	y := d.useArg(f, bb, inst, inst.Args[1])
	y = f.cur.NewBitCast(y, types.NewPointer(types.I8))
	d.defArg(f, bb, inst, inst.Args[0], y)
	return nil
}

// instMOVSD translates the given MOVSD instruction from x86 machine code to
// LLVM IR assembly.
func (d *disassembler) instMOVSD(f *Func, bb *BasicBlock, inst *Inst) error {
	y := d.useArg(f, bb, inst, inst.Args[1])
	y = f.cur.NewBitCast(y, types.NewPointer(types.I32))
	d.defArg(f, bb, inst, inst.Args[0], y)
	return nil
}

// instMOVSW translates the given MOVSW instruction from x86 machine code to
// LLVM IR assembly.
func (d *disassembler) instMOVSW(f *Func, bb *BasicBlock, inst *Inst) error {
	y := d.useArg(f, bb, inst, inst.Args[1])
	y = f.cur.NewBitCast(y, types.NewPointer(types.I16))
	d.defArg(f, bb, inst, inst.Args[0], y)
	return nil
}

// instMOVZX translates the given MOVZX instruction from x86 machine code to
// LLVM IR assembly.
func (d *disassembler) instMOVZX(f *Func, bb *BasicBlock, inst *Inst) error {
	y := d.useArg(f, bb, inst, inst.Args[1])
	d.defArg(f, bb, inst, inst.Args[0], y)
	return nil
}

// instSAR translates the given SAR instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instSAR(f *Func, bb *BasicBlock, inst *Inst) error {
	// shift arithmetic right (SAR)
	x := d.useArg(f, bb, inst, inst.Args[0])
	y := d.useArg(f, bb, inst, inst.Args[1])
	result := f.cur.NewAShr(x, y)
	d.defArg(f, bb, inst, inst.Args[0], result)
	return nil
}

// instSUB translates the given SUB instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instSUB(f *Func, bb *BasicBlock, inst *Inst) error {
	x := d.useArg(f, bb, inst, inst.Args[0])
	y := d.useArg(f, bb, inst, inst.Args[1])
	result := f.cur.NewSub(x, y)
	d.defArg(f, bb, inst, inst.Args[0], result)
	return nil
}

// instTEST translates the given TEST instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instTEST(f *Func, bb *BasicBlock, inst *Inst) error {
	x := d.useArg(f, bb, inst, inst.Args[0])
	y := d.useArg(f, bb, inst, inst.Args[1])
	// Set the status flags according to the result.

	// TODO: Fix calculation of status flags. Pass SUB from CMP instruction and
	// AND from TEST instruction.
	return d.updateStatusFlags(f, bb, x, y)
}

// instXOR translates the given XOR instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instXOR(f *Func, bb *BasicBlock, inst *Inst) error {
	x := d.useArg(f, bb, inst, inst.Args[0])
	y := d.useArg(f, bb, inst, inst.Args[1])
	result := f.cur.NewXor(x, y)
	d.defArg(f, bb, inst, inst.Args[0], result)
	return nil
}

// translateTerm translates the given terminator from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) translateTerm(f *Func, bb *BasicBlock, term *Inst) error {
	if term.isDummyTerm() {
		target, ok := f.blocks[term.addr]
		if !ok {
			return errors.Errorf("unable to locate basic block at %v", term.addr)
		}
		f.cur.NewBr(target)
		return nil
	}
	fmt.Println("term:", term)
	switch term.Op {
	case x86asm.JA, x86asm.JAE, x86asm.JB, x86asm.JBE, x86asm.JCXZ, x86asm.JE, x86asm.JECXZ, x86asm.JG, x86asm.JGE, x86asm.JL, x86asm.JLE, x86asm.JNE, x86asm.JNO, x86asm.JNP, x86asm.JNS, x86asm.JO, x86asm.JP, x86asm.JRCXZ, x86asm.JS:
		return d.termCondBranch(f, bb, term)
	case x86asm.JMP:
		return d.termJMP(f, bb, term)
	case x86asm.RET:
		return d.termRET(f, bb, term)
	default:
		panic(fmt.Errorf("support for terminator opcode %v not yet implemented", term.Op))
	}
}

// termCondBranch translates the given conditional branch terminator from x86
// machine code to LLVM IR assembly.
func (d *disassembler) termCondBranch(f *Func, bb *BasicBlock, term *Inst) error {
	// target branch of conditional branching instruction.
	next := term.addr + bin.Address(term.Len)
	targetTrueAddrs := d.getAddrs(next, term.Args[0])
	if len(targetTrueAddrs) != 1 {
		return errors.Errorf("invalid number of true branches; expected 1, got %d", len(targetTrueAddrs))
	}
	targetTrueAddr := targetTrueAddrs[0]
	targetTrue, ok := f.blocks[targetTrueAddr]
	if !ok {
		return errors.Errorf("unable to locate basic block at %v", targetTrueAddr)
	}
	// fallthrough branch of conditional branching instruction.
	targetFalseAddr := next
	targetFalse, ok := f.blocks[targetFalseAddr]
	if !ok {
		return errors.Errorf("unable to locate basic block at %v", targetTrueAddr)
	}
	// Compute conditional value.
	//
	//    (CF=0 and ZF=0)    JA      Jump if above.
	//    (CF=0 and ZF=0)    JNBE    Jump if not below or equal.     PSEUDO-instruction
	//    (CF=0)             JAE     Jump if above or equal.
	//    (CF=0)             JNB     Jump if not below.              PSEUDO-instruction
	//    (CF=0)             JNC     Jump if not carry.              PSEUDO-instruction
	//    (CF=1 or ZF=1)     JBE     Jump if below or equal.
	//    (CF=1 or ZF=1)     JNA     Jump if not above.              PSEUDO-instruction
	//    (CF=1)             JB      Jump if below.
	//    (CF=1)             JC      Jump if carry.                  PSEUDO-instruction
	//    (CF=1)             JNAE    Jump if not above or equal.     PSEUDO-instruction
	//    (CX=0)             JCXZ    Jump if CX register is zero.
	//    (ECX=0)            JECXZ   Jump if ECX register is zero.
	//    (OF=0)             JNO     Jump if not overflow.
	//    (OF=1)             JO      Jump if overflow.
	//    (PF=0)             JNP     Jump if not parity.
	//    (PF=0)             JPO     Jump if parity odd.             PSEUDO-instruction
	//    (PF=1)             JP      Jump if parity.
	//    (PF=1)             JPE     Jump if parity even.            PSEUDO-instruction
	//    (RCX=0)            JRCXZ   Jump if RCX register is zero.
	//    (SF=0)             JNS     Jump if not sign.
	//    (SF=1)             JS      Jump if sign.
	//    (SF=OF)            JGE     Jump if greater or equal.
	//    (SF=OF)            JNL     Jump if not less.               PSEUDO-instruction
	//    (SF≠OF)            JL      Jump if less.
	//    (SF≠OF)            JNGE    Jump if not greater or equal.   PSEUDO-instruction
	//    (ZF=0 and SF=OF)   JG      Jump if greater.
	//    (ZF=0 and SF=OF)   JNLE    Jump if not less or equal.      PSEUDO-instruction
	//    (ZF=0)             JNE     Jump if not equal.
	//    (ZF=0)             JNZ     Jump if not zero.               PSEUDO-instruction
	//    (ZF=1 or SF≠OF)    JLE     Jump if less or equal.
	//    (ZF=1 or SF≠OF)    JNG     Jump if not greater.            PSEUDO-instruction
	//    (ZF=1)             JE      Jump if equal.
	//    (ZF=1)             JZ      Jump if zero.                   PSEUDO-instruction
	//
	// ref: $ 3.2 Jcc - Jump if Condition Is Met, Intel 64 and IA-32
	// Architectures Software Developer's Manual
	var cond value.Value
	switch term.Op {
	case x86asm.JA:
		// Jump if above.
		//
		//    CF=0 and ZF=0
		cf := d.useStatus(f, bb, CF)
		zf := d.useStatus(f, bb, ZF)
		cond1 := f.cur.NewICmp(ir.IntEQ, cf, constant.False)
		cond2 := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
		cond = f.cur.NewAnd(cond1, cond2)
	case x86asm.JAE:
		// Jump if above or equal.
		//
		//    CF=0
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JBE:
		// Jump if below or equal.
		//
		//    CF=1 or ZF=1
		cf := d.useStatus(f, bb, CF)
		zf := d.useStatus(f, bb, ZF)
		cond = f.cur.NewOr(cf, zf)
	case x86asm.JB:
		// Jump if below.
		//
		//    CF=1
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JCXZ:
		// Jump if CX register is zero.
		//
		//    CX=0
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JECXZ:
		// Jump if ECX register is zero.
		//
		//    ECX=0
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JNO:
		// Jump if not overflow.
		//
		//    OF=0
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JO:
		// Jump if overflow.
		//
		//    OF=1
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JNP:
		// Jump if not parity.
		//
		//    PF=0
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JP:
		// Jump if parity.
		//
		//    PF=1
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JRCXZ:
		// Jump if RCX register is zero.
		//
		//    RCX=0
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JNS:
		// Jump if not sign.
		//
		//    SF=0
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JS:
		// Jump if sign.
		//
		//    SF=1
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JGE:
		// Jump if greater or equal.
		//
		//    SF=OF
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JL:
		// Jump if less.
		//
		//    SF≠OF
		sf := d.useStatus(f, bb, SF)
		of := d.useStatus(f, bb, OF)
		cond = f.cur.NewICmp(ir.IntNE, sf, of)
	case x86asm.JG:
		// Jump if greater.
		//
		//    ZF=0 and SF=OF
		sf := d.useStatus(f, bb, SF)
		of := d.useStatus(f, bb, OF)
		zf := d.useStatus(f, bb, ZF)
		cond1 := f.cur.NewICmp(ir.IntEQ, zf, constant.False)
		cond2 := f.cur.NewICmp(ir.IntEQ, sf, of)
		cond = f.cur.NewAnd(cond1, cond2)
	case x86asm.JNE:
		// Jump if not equal.
		//
		//    ZF=0
		zf := d.useStatus(f, bb, ZF)
		cond = f.cur.NewICmp(ir.IntEQ, zf, constant.False)
	case x86asm.JLE:
		// Jump if less or equal.
		//
		//    ZF=1 or SF≠OF
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	case x86asm.JE:
		// Jump if equal.
		//
		//    ZF=1
		zf := d.useStatus(f, bb, ZF)
		cond = zf
	default:
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	}
	f.cur.NewCondBr(cond, targetTrue, targetFalse)
	return nil
}

// termJMP translates the given JMP terminator from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) termJMP(f *Func, bb *BasicBlock, term *Inst) error {
	if d.isTailCall(f.entry, term) {
		// Handle tail call terminator instructions.

		// Hack: interpret JMP instruction as CALL instruction. Works since
		// instCALL only interprets inst.Args[0], which is the same in both
		// call and jmp instructions.
		if err := d.instCALL(f, bb, term); err != nil {
			return errors.WithStack(err)
		}
		// Add return statement.
		// Handle return values of non-void functions (passed through EAX).
		if !types.Equal(f.Sig.Ret, types.Void) {
			result := d.useArg(f, bb, nil, x86asm.EAX)
			f.cur.NewRet(result)
			return nil
		}
		f.cur.NewRet(nil)
		return nil
	}
	if addr, ok := d.getAddr(f, bb, term, term.Args[0]); ok {
		if target, ok := f.blocks[addr]; ok {
			f.cur.NewBr(target)
			return nil
		}
		return errors.Errorf("unable to locate basic block at %v", addr)
	}
	if arg, ok := term.Args[0].(x86asm.Mem); ok {
		if targetAddrs, ok := d.tables[bin.Address(arg.Disp)]; ok {
			// TODO: Implement proper support for switch table jmp translation.

			// TODO: Locate default target using information from symbolic
			// execution and predecessor basic blocks.

			// At this stage of recovery, the assumption is `index` is always
			// within the range of the jump table offsets. Thus, the default branch
			// is always unreachable.
			//
			// This assumption will be validated and revisited when information
			// from symbolic execution is available.
			index := d.useReg(f, arg.Index)
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
				c := ir.NewCase(constant.NewInt(int64(i), index.Type()), target)
				cases = append(cases, c)
			}
			// TODO: Add support for indirect switch statements.
			f.cur.NewSwitch(index, targetDefault, cases...)
			return nil
		}
	}
	// TODO: Add proper support for JMP terminators.
	fmt.Println("termJMP arg:", term.Args[0])
	panic(fmt.Errorf("support for terminator opcode %v not yet implemented", term.Op))
}

// termRET translates the given RET terminator from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) termRET(f *Func, bb *BasicBlock, term *Inst) error {
	// Handle return values of non-void functions (passed through EAX).
	if !types.Equal(f.Sig.Ret, types.Void) {
		result := d.useArg(f, bb, nil, x86asm.EAX)
		f.cur.NewRet(result)
		return nil
	}
	f.cur.NewRet(nil)
	return nil
}

// getAddr returns the address specified by the given argument, and a boolean
// value indicating success.
func (d *disassembler) getAddr(f *Func, bb *BasicBlock, inst *Inst, arg x86asm.Arg) (bin.Address, bool) {
	switch arg := arg.(type) {
	case x86asm.Reg:
		fmt.Println("arg:", arg)
		pretty.Println(arg)
		if context, ok := d.contexts[inst.addr]; ok {
			if c, ok := context.Regs[Register(arg)]; ok {
				if addr, ok := c["addr"]; ok {
					return bin.Address(addr), true
				}
			}
		}
	case x86asm.Mem:
		// Segment:[Base+Scale*Index+Disp].

		// TODO: Add proper support for memory arguments.
		//
		//    Segment Reg
		//    Base    Reg
		//    Scale   uint8
		//    Index   Reg
		if arg.Segment == 0 && arg.Base == 0 && arg.Scale == 0 && arg.Index == 0 {
			return bin.Address(arg.Disp), true
		}
		if arg.Disp > 0 {
			fmt.Printf("unable to locate memory at address %v\n", bin.Address(arg.Disp))
		}

	//case x86asm.Imm:
	case x86asm.Rel:
		next := inst.addr + bin.Address(inst.Len)
		return next + bin.Address(arg), true
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
	return 0, false
}

// mem returns a pointer to the LLVM IR value associated with the given memory
// argument.
func (d *disassembler) mem(f *Func, bb *BasicBlock, inst *Inst, arg x86asm.Mem) value.Value {
	// Early return if constant address.
	if arg.Segment == 0 && arg.Base == 0 && arg.Scale == 0 && arg.Index == 0 {
		addr := bin.Address(arg.Disp)
		if fn, ok := d.funcs[addr]; ok {
			return fn
		}
		if g, ok := d.global(f, bb, addr); ok {
			return g
		}
		panic(fmt.Errorf("unable to locate function or global at %v", addr))
	}

	// TODO: Figure out how to handle Segment.

	// Segment:[Base+Scale*Index+Disp].
	//
	//    Segment Reg
	//    Base    Reg
	//    Scale   uint8
	//    Index   Reg
	//    Disp    int64
	var result value.Value
	var base value.Value
	if arg.Base != 0 {
		base = d.useArg(f, bb, inst, arg.Base)
	}
	var scaledIndex value.Value
	if arg.Index != 0 {
		if arg.Scale == 0 {
			panic(fmt.Errorf("invalid scale; zero scale used with non-zero index %v", arg.Index))
		}
		index := d.useArg(f, bb, inst, arg.Index)
		if arg.Scale == 1 {
			scaledIndex = index
		} else {
			scale := constant.NewInt(int64(arg.Scale), types.I32)
			scaledIndex = f.cur.NewMul(scale, index)
		}
	}
	var disp value.Value
	if arg.Disp != 0 {
		disp = constant.NewInt(arg.Disp, types.I64)
	}
	if base != nil {
		result = base
	}
	if scaledIndex != nil {
		if result == nil {
			result = scaledIndex
		} else {
			result = f.cur.NewAdd(result, scaledIndex)
		}
	}
	if disp != nil {
		if result == nil {
			result = disp
		} else {
			result = f.cur.NewAdd(result, disp)
		}
	}
	if result == nil {
		result = constant.NewInt(0, types.I64)
	}
	// TODO: Fix type once type analysis information is available.
	return f.cur.NewBitCast(result, types.NewPointer(result.Type()))
}

func (d *disassembler) useArg(f *Func, bb *BasicBlock, inst *Inst, arg x86asm.Arg) value.Value {
	fmt.Println("useArg:", arg)
	switch arg := arg.(type) {
	case x86asm.Reg:
		return d.useReg(f, arg)
	case x86asm.Mem:
		src := d.mem(f, bb, inst, arg)
		return f.cur.NewLoad(src)
	case x86asm.Imm:
		return constant.NewInt(int64(arg), types.I32)
	case x86asm.Rel:
		addr := inst.addr + bin.Address(inst.Len) + bin.Address(arg)
		if v, ok := d.funcs[addr]; ok {
			return v
		}
		if g, ok := d.global(f, bb, addr); ok {
			fmt.Println("inst:", inst)
			fmt.Println("arg:", arg)
			panic("not yet implemented")
			// TODO: Verify if the global variable should be loaded, or used as
			// pointer.
			return f.cur.NewLoad(g)
		}
		panic(fmt.Errorf("unable to locate global or function at %v", addr))
	default:
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

func (d *disassembler) useReg(f *Func, reg x86asm.Reg) value.Value {
	src := d.reg(f, reg)
	return f.cur.NewLoad(src)
}

func (d *disassembler) defReg(f *Func, reg x86asm.Reg, v value.Value) {
	dst := d.reg(f, reg)
	f.cur.NewStore(v, dst)
}

func (d *disassembler) defArg(f *Func, bb *BasicBlock, inst *Inst, arg x86asm.Arg, v value.Value) {
	fmt.Println("defArg:", arg)
	switch arg := arg.(type) {
	case x86asm.Reg:
		dst := d.reg(f, arg)
		f.cur.NewStore(v, dst)
	case x86asm.Mem:
		dst := d.mem(f, bb, inst, arg)
		f.cur.NewStore(v, dst)
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// reg returns the LLVM IR value associated with the given x86 register.
func (d *disassembler) reg(f *Func, reg x86asm.Reg) value.Value {
	if v, ok := f.regs[reg]; ok {
		return v
	}
	var typ types.Type
	switch reg {
	// 8-bit
	case x86asm.AL, x86asm.CL, x86asm.DL, x86asm.BL, x86asm.AH, x86asm.CH, x86asm.DH, x86asm.BH, x86asm.SPB, x86asm.BPB, x86asm.SIB, x86asm.DIB, x86asm.R8B, x86asm.R9B, x86asm.R10B, x86asm.R11B, x86asm.R12B, x86asm.R13B, x86asm.R14B, x86asm.R15B:
		typ = types.I8
	// 16-bit
	case x86asm.AX, x86asm.CX, x86asm.DX, x86asm.BX, x86asm.SP, x86asm.BP, x86asm.SI, x86asm.DI, x86asm.R8W, x86asm.R9W, x86asm.R10W, x86asm.R11W, x86asm.R12W, x86asm.R13W, x86asm.R14W, x86asm.R15W:
		typ = types.I16
	// 32-bit
	case x86asm.EAX, x86asm.ECX, x86asm.EDX, x86asm.EBX, x86asm.ESP, x86asm.EBP, x86asm.ESI, x86asm.EDI, x86asm.R8L, x86asm.R9L, x86asm.R10L, x86asm.R11L, x86asm.R12L, x86asm.R13L, x86asm.R14L, x86asm.R15L:
		typ = types.I32
	// 64-bit
	case x86asm.RAX, x86asm.RCX, x86asm.RDX, x86asm.RBX, x86asm.RSP, x86asm.RBP, x86asm.RSI, x86asm.RDI, x86asm.R8, x86asm.R9, x86asm.R10, x86asm.R11, x86asm.R12, x86asm.R13, x86asm.R14, x86asm.R15:
		typ = types.I64
	// Instruction pointer.
	case x86asm.IP: // 16-bit
		typ = types.I16
	case x86asm.EIP: // 32-bit
		typ = types.I32
	case x86asm.RIP: // 64-bit
		typ = types.I64
	// 387 floating point registers.
	case x86asm.F0, x86asm.F1, x86asm.F2, x86asm.F3, x86asm.F4, x86asm.F5, x86asm.F6, x86asm.F7:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// MMX registers.
	case x86asm.M0, x86asm.M1, x86asm.M2, x86asm.M3, x86asm.M4, x86asm.M5, x86asm.M6, x86asm.M7:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// XMM registers.
	case x86asm.X0, x86asm.X1, x86asm.X2, x86asm.X3, x86asm.X4, x86asm.X5, x86asm.X6, x86asm.X7, x86asm.X8, x86asm.X9, x86asm.X10, x86asm.X11, x86asm.X12, x86asm.X13, x86asm.X14, x86asm.X15:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// Segment registers.
	case x86asm.ES, x86asm.CS, x86asm.SS, x86asm.DS, x86asm.FS, x86asm.GS:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// System registers.
	case x86asm.GDTR:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	case x86asm.IDTR:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	case x86asm.LDTR:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	case x86asm.MSW:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	case x86asm.TASK:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// Control registers.
	case x86asm.CR0, x86asm.CR1, x86asm.CR2, x86asm.CR3, x86asm.CR4, x86asm.CR5, x86asm.CR6, x86asm.CR7, x86asm.CR8, x86asm.CR9, x86asm.CR10, x86asm.CR11, x86asm.CR12, x86asm.CR13, x86asm.CR14, x86asm.CR15:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// Debug registers.
	case x86asm.DR0, x86asm.DR1, x86asm.DR2, x86asm.DR3, x86asm.DR4, x86asm.DR5, x86asm.DR6, x86asm.DR7, x86asm.DR8, x86asm.DR9, x86asm.DR10, x86asm.DR11, x86asm.DR12, x86asm.DR13, x86asm.DR14, x86asm.DR15:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// Task registers.
	case x86asm.TR0, x86asm.TR1, x86asm.TR2, x86asm.TR3, x86asm.TR4, x86asm.TR5, x86asm.TR6, x86asm.TR7:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	default:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	}
	v := ir.NewAlloca(typ)
	v.SetName(strings.ToLower(reg.String()))
	f.regs[reg] = v
	return v
}

// updateStatusFlags updates the status flags based on the result of an
// arithmetic instruction, emitting LLVM IR code to the given basic block.
//
// Status flags
//
//    CF (bit 0)    Carry Flag
//    PF (bit 2)    Parity Flag
//    AF (bit 4)    Auxiliary Carry Flag
//    ZF (bit 6)    Zero Flag
//    SF (bit 7)    Sign Flag
//    OF (bit 11)   Overflow Flag
//
// ref: $ 3.4.3.1 Status Flags, Intel 64 and IA-32 Architectures Software
// Developer's Manual
func (d *disassembler) updateStatusFlags(f *Func, bb *BasicBlock, x, y value.Value) error {
	// CF (bit 0) Carry flag - Set if an arithmetic operation generates a carry
	// or a borrow out of the most- significant bit of the result; cleared
	// otherwise. This flag indicates an overflow condition for unsigned-integer
	// arithmetic. It is also used in multiple-precision arithmetic.

	// TODO: Add support for the CF status flag.

	// PF (bit 2) Parity flag - Set if the least-significant byte of the result
	// contains an even number of 1 bits; cleared otherwise.

	// TODO: Add support for the PF status flag.

	// AF (bit 4) Auxiliary Carry flag - Set if an arithmetic operation generates
	// a carry or a borrow out of bit 3 of the result; cleared otherwise. This
	// flag is used in binary-coded decimal (BCD) arithmetic.

	// TODO: Add support for the AF status flag.

	// ZF (bit 6) Zero flag - Set if the result is zero; cleared otherwise.
	zf := f.cur.NewICmp(ir.IntEQ, x, y)
	d.defStatus(f, bb, ZF, zf)

	// SF (bit 7) Sign flag - Set equal to the most-significant bit of the
	// result, which is the sign bit of a signed integer. (0 indicates a positive
	// value and 1 indicates a negative value.)
	sf := f.cur.NewICmp(ir.IntSLT, x, y)
	d.defStatus(f, bb, SF, sf)

	// OF (bit 11) Overflow flag - Set if the integer result is too large a
	// positive number or too small a negative number (excluding the sign-bit) to
	// fit in the destination operand; cleared otherwise. This flag indicates an
	// overflow condition for signed-integer (two's complement) arithmetic.

	// TODO: Add support for the OF status flag.

	return nil
}

// status returns a pointer to the LLVM IR value associated with the given
// status flag.
func (d *disassembler) status(f *Func, status StatusFlag) value.Value {
	if v, ok := f.statusFlags[status]; ok {
		return v
	}
	v := ir.NewAlloca(types.I1)
	v.SetName(strings.ToLower(status.String()))
	f.statusFlags[status] = v
	return v
}

// useStatus loads and returns the LLVM IR value associated with the given
// status flag.
func (d *disassembler) useStatus(f *Func, bb *BasicBlock, status StatusFlag) value.Value {
	src := d.status(f, status)
	return f.cur.NewLoad(src)
}

// defStatus stores the given value to the LLVM IR value associated with the
// given status flag.
func (d *disassembler) defStatus(f *Func, bb *BasicBlock, status StatusFlag, v value.Value) {
	dst := d.status(f, status)
	f.cur.NewStore(v, dst)
}

// useGlobal loads and returns the LLVM IR value associated with the given
// global variable address.
func (d *disassembler) useGlobal(f *Func, bb *BasicBlock, addr bin.Address) value.Value {
	src, ok := d.global(f, bb, addr)
	if !ok {
		panic(fmt.Sprintf("unable to locate global variable at %v", addr))
	}
	return f.cur.NewLoad(src)
}

// defGlobal stores the given value to the LLVM IR value associated with the
// given global variable address.
func (d *disassembler) defGlobal(f *Func, bb *BasicBlock, addr bin.Address, v value.Value) {
	dst, ok := d.global(f, bb, addr)
	if !ok {
		panic(fmt.Sprintf("unable to locate global variable at %v", addr))
	}
	f.cur.NewStore(v, dst)
}

// global returns a pointer to the LLVM IR value associated with the given
// global variable, and a boolean value indicating success.
func (d *disassembler) global(f *Func, bb *BasicBlock, addr bin.Address) (value.Value, bool) {
	// Early return if direct access to global variable.
	if src, ok := d.globals[addr]; ok {
		return src, true
	}

	// Use binary search if indirect access to global variable (e.g. struct
	// field, array element).
	var globalAddrs []bin.Address
	for globalAddr := range d.globals {
		globalAddrs = append(globalAddrs, globalAddr)
	}
	sort.Sort(bin.Addresses(globalAddrs))
	less := func(i int) bool {
		return addr < globalAddrs[i]
	}
	index := sort.Search(len(globalAddrs), less)
	index--
	if 0 <= index && index < len(globalAddrs) {
		start := globalAddrs[index]
		g := d.globals[start]
		size := d.sizeOfType(g.Typ.Elem)
		end := start + bin.Address(size)
		if start <= addr && addr < end {
			offset := int64(addr - start)
			return d.getElementPtr(f, bb, g, offset), true
		}
	}
	return nil, false
}

// getElementPtr returns a pointer to the given offset into the source value.
func (d *disassembler) getElementPtr(f *Func, bb *BasicBlock, src value.Value, offset int64) *ir.InstGetElementPtr {
	srcType, ok := src.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
	}
	elem := srcType.Elem
	e := elem
	total := int64(0)
	var indices []value.Value
	for i := 0; total < offset; i++ {
		if i == 0 {
			// Ignore checking the 0th index as it simply follows the pointer of
			// src.
			//
			// ref: http://llvm.org/docs/GetElementPtr.html#why-is-the-extra-0-index-required
			index := constant.NewInt(0, types.I64)
			indices = append(indices, index)
			continue
		}
		switch t := e.(type) {
		case *types.PointerType:
			// ref: http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep
			panic("unable to index into element of pointer type; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep")
		case *types.ArrayType:
			elemSize := d.sizeOfType(t.Elem)
			j := int64(0)
			for ; j < t.Len; j++ {
				if total+elemSize > offset {
					break
				}
				total += elemSize
			}
			index := constant.NewInt(j, types.I64)
			indices = append(indices, index)
			e = t.Elem
		case *types.StructType:
			j := int64(0)
			for ; j < int64(len(t.Fields)); j++ {
				fieldSize := d.sizeOfType(t.Fields[j])
				if total+fieldSize > offset {
					break
				}
				total += fieldSize
			}
			index := constant.NewInt(j, types.I64)
			indices = append(indices, index)
			e = t.Fields[j]
		default:
			panic(fmt.Errorf("support for indexing element type %T not yet implemented", e))
		}
	}
	return f.cur.NewGetElementPtr(src, indices...)
}
