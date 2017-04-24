package main

import (
	"fmt"
	"strings"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"golang.org/x/arch/x86/x86asm"
)

// An Arg is a single x86 instruction argument.
type Arg struct {
	// x86 instruction argument.
	x86asm.Arg
	// Parent instruction; used to calculate relative offsets and retrieve
	// symbolic execution information.
	parent *Inst
}

// Arg returns the i:th argument of the instruction.
func (inst *Inst) Arg(i int) *Arg {
	return &Arg{
		Arg:    inst.Args[i],
		parent: inst,
	}
}

// === [ usage ] ===============================================================

// useArg returns the value held by the given argument, emitting code to f.
func (f *Func) useArg(arg *Arg) value.Value {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		return f.useReg(a)
	case x86asm.Mem:
		return f.useMem(a)
	case x86asm.Imm:
		return constant.NewInt(int64(a), types.I32)
	case x86asm.Rel:
		next := arg.parent.addr + bin.Address(arg.parent.Len)
		addr := next + bin.Address(a)
		return f.useAddr(addr)
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// --- [ register ] ------------------------------------------------------------

// useReg loads and returns the value of the given x86 register, emitting code
// to f.
func (f *Func) useReg(reg x86asm.Reg) value.Value {
	src := f.reg(reg)
	return f.cur.NewLoad(src)
}

// --- [ memory access ] -------------------------------------------------------

// useMem loads and returns the value of the given memory argument, emitting
// code to f.
func (f *Func) useMem(mem x86asm.Mem) value.Value {
	src := f.mem(mem)
	return f.cur.NewLoad(src)
}

// --- [ status flag ] ---------------------------------------------------------

// useStatus loads and returns the value of the given x86 status flag, emitting
// code to f.
func (f *Func) useStatus(status StatusFlag) value.Value {
	src := f.status(status)
	return f.cur.NewLoad(src)
}

// --- [ address ] -------------------------------------------------------------

// useAddr loads and returns the value of the given address, emitting code to f.
func (f *Func) useAddr(addr bin.Address) value.Value {
	src, ok := f.addr(addr)
	if !ok {
		panic(fmt.Errorf("unable to locate value at address %v", addr))
	}
	return f.cur.NewLoad(src)
}

// === [ definition ] ==========================================================

// defArg stores the value to the given argument, emitting code to f.
func (f *Func) defArg(arg *Arg, v value.Value) {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		f.defReg(a, v)
	case x86asm.Mem:
		f.defMem(a, v)
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// --- [ register ] ------------------------------------------------------------

// defReg stores the value to the given x86 register, emitting code to f.
func (f *Func) defReg(reg x86asm.Reg, v value.Value) {
	dst := f.reg(reg)
	f.cur.NewStore(v, dst)
}

// --- [ memory access ] -------------------------------------------------------

// defMem stores the value to the given memory argument, emitting code to f.
func (f *Func) defMem(mem x86asm.Mem, v value.Value) {
	dst := f.mem(mem)
	f.cur.NewStore(v, dst)
}

// --- [ status flag ] ---------------------------------------------------------

// defStatus stores the value to the given x86 status flag, emitting code to f.
func (f *Func) defStatus(status StatusFlag, v value.Value) {
	dst := f.status(status)
	f.cur.NewStore(v, dst)
}

// --- [ address ] -------------------------------------------------------------

// defAddr stores the value to the given address, emitting code to f.
func (f *Func) defAddr(addr bin.Address, v value.Value) {
	dst, ok := f.addr(addr)
	if !ok {
		panic(fmt.Errorf("unable to locate value at address %v", addr))
	}
	f.cur.NewStore(v, dst)
}

// === [ pointer to value ] ====================================================

// --- [ register ] ------------------------------------------------------------

// reg returns a pointer to the LLVM IR value associated with the given x86
// register.
func (f *Func) reg(reg x86asm.Reg) value.Value {
	if v, ok := f.regs[reg]; ok {
		return v
	}
	typ := regType(reg)
	v := ir.NewAlloca(typ)
	name := strings.ToLower(reg.String())
	v.SetName(name)
	f.regs[reg] = v
	return v
}

// regType returns the LLVM IR type of the given register.
func regType(reg x86asm.Reg) types.Type {
	switch reg {
	// 8-bit
	case x86asm.AL, x86asm.CL, x86asm.DL, x86asm.BL, x86asm.AH, x86asm.CH, x86asm.DH, x86asm.BH, x86asm.SPB, x86asm.BPB, x86asm.SIB, x86asm.DIB, x86asm.R8B, x86asm.R9B, x86asm.R10B, x86asm.R11B, x86asm.R12B, x86asm.R13B, x86asm.R14B, x86asm.R15B:
		return types.I8
	// 16-bit
	case x86asm.AX, x86asm.CX, x86asm.DX, x86asm.BX, x86asm.SP, x86asm.BP, x86asm.SI, x86asm.DI, x86asm.R8W, x86asm.R9W, x86asm.R10W, x86asm.R11W, x86asm.R12W, x86asm.R13W, x86asm.R14W, x86asm.R15W:
		return types.I16
	// 32-bit
	case x86asm.EAX, x86asm.ECX, x86asm.EDX, x86asm.EBX, x86asm.ESP, x86asm.EBP, x86asm.ESI, x86asm.EDI, x86asm.R8L, x86asm.R9L, x86asm.R10L, x86asm.R11L, x86asm.R12L, x86asm.R13L, x86asm.R14L, x86asm.R15L:
		return types.I32
	// 64-bit
	case x86asm.RAX, x86asm.RCX, x86asm.RDX, x86asm.RBX, x86asm.RSP, x86asm.RBP, x86asm.RSI, x86asm.RDI, x86asm.R8, x86asm.R9, x86asm.R10, x86asm.R11, x86asm.R12, x86asm.R13, x86asm.R14, x86asm.R15:
		return types.I64
	// Instruction pointer.
	case x86asm.IP: // 16-bit
		return types.I16
	case x86asm.EIP: // 32-bit
		return types.I32
	case x86asm.RIP: // 64-bit
		return types.I64
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
}

// --- [ memory access ] -------------------------------------------------------

// mem returns a pointer to the LLVM IR value associated with the given memory
// argument, emitting code to f.
func (f *Func) mem(mem x86asm.Mem) value.Value {
	// Segment:[Base+Scale*Index+Disp].

	// TODO: Add proper support for memory arguments.
	//    Segment Reg
	//    Base    Reg
	//    Scale   uint8
	//    Index   Reg
	//    Disp    int64

	// Handle direct memory access.
	segment := mem.Segment
	base := mem.Base
	scale := mem.Scale
	index := mem.Index
	disp := mem.Disp
	if segment == 0 && base == 0 && scale == 0 && index == 0 {
		addr := bin.Address(disp)
		v, ok := f.addr(addr)
		if !ok {
			panic(fmt.Errorf("unable to locate value at address %v", addr))
		}
		return v
	}

	// Handle local variables.
	if segment == 0 && index == 0 {
		// Stack local memory access.
		if base == x86asm.ESP || base == x86asm.EBP {
			name := fmt.Sprintf("%s_%d", strings.ToLower(base.String()), f.espDisp+disp)
			if v, ok := f.locals[name]; ok {
				return v
			}
			v := ir.NewAlloca(types.I32)
			v.SetName(name)
			f.locals[name] = v
			return v
		}
	}

	pretty.Println(mem)
	panic("not yet implemented")
}

// --- [ status flag ] ---------------------------------------------------------

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

// status returns a pointer to the LLVM IR value associated with the given x86
// status flag.
func (f *Func) status(status StatusFlag) value.Value {
	if v, ok := f.statusFlags[status]; ok {
		return v
	}
	v := ir.NewAlloca(types.I1)
	name := strings.ToLower(status.String())
	v.SetName(name)
	f.statusFlags[status] = v
	return v
}

// --- [ address ] -------------------------------------------------------------

// addr returns a pointer to the LLVM IR value associated with the given
// address, emitting code to f. The returned value is one of *ir.BasicBlock,
// *ir.Global and *ir.Function, and the boolean value indicates success
func (f *Func) addr(addr bin.Address) (value.Value, bool) {
	if block, ok := f.blocks[addr]; ok {
		return block, true
	}
	if g, ok := f.d.globals[addr]; ok {
		return g, true
	}
	if fn, ok := f.d.funcs[addr]; ok {
		return fn.Function, true
	}
	fmt.Printf("unable to locate value at address %v\n", addr)
	panic("not yet implemented")
}

// === [ helpers ] =============================================================

// getAddr returns the static address represented by the given argument, and a
// boolean indicating success.
func (f *Func) getAddr(arg *Arg) (bin.Address, bool) {
	switch a := arg.Arg.(type) {
	case x86asm.Rel:
		next := arg.parent.addr + bin.Address(arg.parent.Len)
		addr := next + bin.Address(a)
		return addr, true
	case x86asm.Mem:
		if a.Segment == 0 && a.Base == 0 && a.Scale == 0 && a.Index == 0 {
			return bin.Address(a.Disp), true
		}
	}
	return 0, false
}

// getFunc resolves the function, function type, and calling convention of the
// given argument. The boolean return value indicates success.
func (f *Func) getFunc(arg *Arg) (value.Named, *types.FuncType, ir.CallConv, bool) {
	if addr, ok := f.getAddr(arg); ok {
		if fn, ok := f.d.funcs[addr]; ok {
			v := fn.Function
			return v, v.Sig, v.CallConv, true
		}
	}
	fmt.Printf("unable to locate function for argument %v\n", arg.Arg)
	panic("not yet implemented")
}
