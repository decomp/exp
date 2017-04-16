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
	"github.com/mewbak/x86/x86asm"
	"github.com/pkg/errors"
)

// translateFunc translates the given function from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) translateFunc(f *function) error {
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

	var blocks []*basicBlock
	var blockAddrs []bin.Address
	for _, block := range f.blocks {
		blockAddrs = append(blockAddrs, block.addr)
	}
	sort.Sort(bin.Addresses(blockAddrs))
	for _, blockAddr := range blockAddrs {
		block := f.blocks[blockAddr]
		if err := d.translateBlock(f, block); err != nil {
			return errors.WithStack(err)
		}
		blocks = append(blocks, block)
	}
	if len(blocks) == 0 {
		return errors.New("invalid function definition; missing function body")
	}
	less := func(i, j int) bool {
		return blocks[i].addr < blocks[j].addr
	}
	sort.Slice(blocks, less)

	// Add new entry basic block to define registers used within the function.
	if len(f.regs) > 0 {
		entry := &basicBlock{
			BasicBlock: &ir.BasicBlock{},
		}
		// Allocate local variables for each register used within the function.
		for reg := x86asm.AL; reg <= x86asm.TR7; reg++ {
			if inst, ok := f.regs[reg]; ok {
				entry.AppendInst(inst)
			}
		}
		// Handle calling conventions.
		switch f.CallConv {
		case ir.CallConvX86FastCall:
			params := f.Sig.Params
			if len(params) > 0 {
				if ecx, ok := f.regs[x86asm.ECX]; ok {
					entry.NewStore(f.Sig.Params[0], ecx)
				}
			}
			if len(params) > 1 {
				if edx, ok := f.regs[x86asm.EDX]; ok {
					entry.NewStore(f.Sig.Params[1], edx)
				}
			}
		default:
			// TODO: Add support for additional calling conventions.
		}
		target := blocks[0].BasicBlock
		entry.NewBr(target)
		blocks = append([]*basicBlock{entry}, blocks...)
	}

	for _, block := range blocks {
		f.AppendBlock(block.BasicBlock)
	}

	return nil
}

// translateBlock translates the given basic block from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) translateBlock(f *function, block *basicBlock) error {
	dbg.Printf("translating basic block at %v", block.addr)
	for _, inst := range block.insts {
		if err := d.translateInst(f, block, inst); err != nil {
			return errors.WithStack(err)
		}
	}
	// Translate terminator.
	if err := d.translateTerm(f, block, block.term); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// translateInst translates the given instruction from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) translateInst(f *function, block *basicBlock, inst *instruction) error {
	fmt.Println("inst:", inst)
	switch inst.Op {
	case x86asm.AND:
		return d.instAND(f, block, inst)
	case x86asm.CALL:
		return d.instCALL(f, block, inst)
	case x86asm.CMP:
		return d.instCMP(f, block, inst)
	case x86asm.IMUL:
		return d.instIMUL(f, block, inst)
	case x86asm.INC:
		return d.instINC(f, block, inst)
	case x86asm.MOV:
		return d.instMOV(f, block, inst)
	case x86asm.PUSH, x86asm.POP:
		// TODO: Figure out how to handle push and pop.
		return nil
	case x86asm.XOR:
		return d.instXOR(f, block, inst)
	default:
		panic(fmt.Errorf("support for instruction opcode %v not yet implemented", inst.Op))
	}
}

// instAND translates the given AND instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instAND(f *function, block *basicBlock, inst *instruction) error {
	x := d.useArg(f, block, inst, inst.Args[0])
	y := d.useArg(f, block, inst, inst.Args[1])
	result := block.NewAnd(x, y)
	d.defArg(f, block, inst, inst.Args[0], result)
	return nil
}

// instCALL translates the given CALL instruction from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) instCALL(f *function, block *basicBlock, inst *instruction) error {
	c := d.useArg(f, block, inst, inst.Args[0])
	// TODO: Add support for value.Named callees. Using *ir.Function for now, to
	// gain access to the calling convention of the function. Data flow and type
	// analysis will provide this information in the future also for local
	// function pointer callees.
	callee, ok := c.(*function)
	if !ok {
		return errors.Errorf("invalid callee type; expected *main.function, got %T", c)
	}
	var args []value.Value
	switch callee.CallConv {
	case ir.CallConvX86FastCall:
		params := callee.Sig.Params
		fmt.Println("params:", params)
		if len(params) > 0 {
			arg := d.useArg(f, block, nil, x86asm.ECX)
			args = append(args, arg)
		}
		if len(params) > 1 {
			arg := d.useArg(f, block, nil, x86asm.EDX)
			args = append(args, arg)
		}
	default:
		// TODO: Handle call arguments.
	}
	result := block.NewCall(callee, args...)
	// Handle return values of non-void callees (passed through EAX).
	fmt.Println("call result type:", callee.Sig.Ret)
	if !types.Equal(callee.Sig.Ret, types.Void) {
		d.defArg(f, block, nil, x86asm.EAX, result)
	}
	return nil
}

// instCMP translates the given CMP instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instCMP(f *function, block *basicBlock, inst *instruction) error {
	x := d.useArg(f, block, inst, inst.Args[0])
	y := d.useArg(f, block, inst, inst.Args[1])
	result := block.NewSub(x, y)
	// Set the status flags according to the result.
	return d.updateStatusFlags(f, block, result)
}

// instIMUL translates the given IMUL instruction from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) instIMUL(f *function, block *basicBlock, inst *instruction) error {
	x := d.useArg(f, block, inst, inst.Args[1])
	y := d.useArg(f, block, inst, inst.Args[2])
	result := block.NewMul(x, y)
	d.defArg(f, block, inst, inst.Args[0], result)
	return nil
}

// instINC translates the given INC instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instINC(f *function, block *basicBlock, inst *instruction) error {
	x := d.useArg(f, block, inst, inst.Args[0])
	one := constant.NewInt(1, types.I32)
	result := block.NewAdd(x, one)
	d.defArg(f, block, inst, inst.Args[0], result)
	return nil
}

// instMOV translates the given MOV instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instMOV(f *function, block *basicBlock, inst *instruction) error {
	y := d.useArg(f, block, inst, inst.Args[1])
	d.defArg(f, block, inst, inst.Args[0], y)
	return nil
}

// instXOR translates the given XOR instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instXOR(f *function, block *basicBlock, inst *instruction) error {
	x := d.useArg(f, block, inst, inst.Args[0])
	y := d.useArg(f, block, inst, inst.Args[1])
	result := block.NewXor(x, y)
	d.defArg(f, block, inst, inst.Args[0], result)
	return nil
}

// translateTerm translates the given terminator from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) translateTerm(f *function, block *basicBlock, term *instruction) error {
	fmt.Println("term:", term)
	switch term.Op {
	case x86asm.JA, x86asm.JAE, x86asm.JB, x86asm.JBE, x86asm.JCXZ, x86asm.JE, x86asm.JECXZ, x86asm.JG, x86asm.JGE, x86asm.JL, x86asm.JLE, x86asm.JNE, x86asm.JNO, x86asm.JNP, x86asm.JNS, x86asm.JO, x86asm.JP, x86asm.JRCXZ, x86asm.JS:
		return d.termCondBranch(f, block, term)
	case x86asm.JMP:
		return d.termJMP(f, block, term)
	case x86asm.RET:
		return d.termRET(f, block, term)
	default:
		panic(fmt.Errorf("support for terminator opcode %v not yet implemented", term.Op))
	}
}

// termJMP translates the given JMP terminator from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) termJMP(f *function, block *basicBlock, term *instruction) error {
	if d.isTailCall(f.entry, term) {
		// Handle tail call terminator instructions.

		// Hack: interpret JMP instruction as CALL instruction. Works since
		// instCALL only interprets inst.Args[0], which is the same in both
		// call and jmp instructions.
		if err := d.instCALL(f, block, term); err != nil {
			return errors.WithStack(err)
		}

		// Add return statement.
		// Handle return values of non-void functions (passed through EAX).
		if !types.Equal(f.Sig.Ret, types.Void) {
			result := d.useArg(f, block, nil, x86asm.EAX)
			block.NewRet(result)
			return nil
		}
		block.NewRet(nil)
		return nil
	}
	// TODO: Add proper support for JMP terminators.
	panic(fmt.Errorf("support for terminator opcode %v not yet implemented", term.Op))
}

// termCondBranch translates the given conditional branch terminator from x86
// machine code to LLVM IR assembly.
func (d *disassembler) termCondBranch(f *function, block *basicBlock, term *instruction) error {
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
	var cond value.Value
	switch term.Op {
	case x86asm.JNE:
		zf := block.NewLoad(d.status(f, ZF))
		cond = block.NewICmp(ir.IntNE, zf, constant.True)
	default:
		panic(fmt.Sprintf("support for conditional branch instruction with opcode %v not yet implemented", term.Op))
	}
	block.NewCondBr(cond, targetTrue.BasicBlock, targetFalse.BasicBlock)
	return nil
}

// termRET translates the given RET terminator from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) termRET(f *function, block *basicBlock, term *instruction) error {
	// Handle return values of non-void functions (passed through EAX).
	if !types.Equal(f.Sig.Ret, types.Void) {
		result := d.useArg(f, block, nil, x86asm.EAX)
		block.NewRet(result)
		return nil
	}
	block.NewRet(nil)
	return nil
}

func (d *disassembler) useArg(f *function, block *basicBlock, inst *instruction, arg x86asm.Arg) value.Value {
	fmt.Println("useArg:", arg)
	switch arg := arg.(type) {
	case x86asm.Reg:
		src := d.reg(f, arg)
		return block.NewLoad(src)
	case x86asm.Mem:
		// Segment:[Base+Scale*Index+Disp].

		// TODO: Add proper support for memory arguments.
		//
		//    Segment Reg
		//    Base    Reg
		//    Scale   uint8
		//    Index   Reg
		if g, ok := d.globals[bin.Address(arg.Disp)]; ok {
			return block.NewLoad(g)
		} else if arg.Disp > 0 {
			fmt.Printf("unable to locate memory at address %v\n", bin.Address(arg.Disp))
		}
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	case x86asm.Imm:
		return constant.NewInt(int64(arg), types.I32)
	case x86asm.Rel:
		addr := inst.addr + bin.Address(inst.Len) + bin.Address(arg)
		if v, ok := d.funcs[addr]; ok {
			return v
		}
		if v, ok := d.globals[addr]; ok {
			return v
		}
		panic(fmt.Errorf("unable to locate value at address %v", addr))
	default:
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

func (d *disassembler) defArg(f *function, block *basicBlock, inst *instruction, arg x86asm.Arg, v value.Value) {
	fmt.Println("defArg:", arg)
	switch arg := arg.(type) {
	case x86asm.Reg:
		dst := d.reg(f, arg)
		block.NewStore(v, dst)
	case x86asm.Mem:
		// Segment:[Base+Scale*Index+Disp].

		// TODO: Add proper support for memory arguments.
		//
		//    Segment Reg
		//    Base    Reg
		//    Scale   uint8
		//    Index   Reg
		if dst, ok := d.globals[bin.Address(arg.Disp)]; ok {
			block.NewStore(v, dst)
			return
		} else if arg.Disp > 0 {
			fmt.Printf("unable to locate memory at address %v\n", bin.Address(arg.Disp))
		}
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// reg returns the LLVM IR value associated with the given x86 register.
func (d *disassembler) reg(f *function, reg x86asm.Reg) value.Value {
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
func (d *disassembler) updateStatusFlags(f *function, block *basicBlock, result value.Value) error {
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

	zf := block.NewICmp(ir.IntEQ, result, constant.NewInt(0, types.I32))
	dst := d.status(f, ZF)
	block.NewStore(zf, dst)

	// SF (bit 7) Sign flag - Set equal to the most-significant bit of the
	// result, which is the sign bit of a signed integer. (0 indicates a positive
	// value and 1 indicates a negative value.)

	// TODO: Add support for the SF status flag.

	// OF (bit 11) Overflow flag - Set if the integer result is too large a
	// positive number or too small a nega- tive number (excluding the sign-bit)
	// to fit in the destination operand; cleared otherwise. This flag indicates
	// an overflow condition for signed-integer (two’s complement) arithmetic.

	// TODO: Add support for the OF status flag.

	return nil
}

// StatusFlag represents the set of status flags.
type StatusFlag uint

// Status flags.
const (
	CF StatusFlag = iota // Carry Flag
	PF                   // Parity Flag
	AF                   // Auxiliary Carry Flag
	ZF                   // Zero Flag
	SF                   // Sign Flag
	OF                   // Overflow Flag
)

// String returns the string representation of the status flag.
func (status StatusFlag) String() string {
	m := map[StatusFlag]string{
		CF: "CF",
		PF: "PF",
		AF: "AF",
		ZF: "ZF",
		SF: "SF",
		OF: "OF",
	}
	if s, ok := m[status]; ok {
		return s
	}
	return fmt.Sprintf("unknown status flag %d", uint(status))
}

// status returns the LLVM IR value associated with the given status flag.
func (d *disassembler) status(f *function, status StatusFlag) value.Value {
	if v, ok := f.status[status]; ok {
		return v
	}
	v := ir.NewAlloca(types.I1)
	v.SetName(strings.ToLower(status.String()))
	f.status[status] = v
	return v
}
