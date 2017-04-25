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

// === [ argument ] ============================================================

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

// Mem returns the memory reference at the i:th argument of the instruction.
func (inst *Inst) Mem(i int) *Mem {
	return NewMem(inst.Args[i], inst)
}

// Reg returns the register at the i:th argument of the instruction.
func (inst *Inst) Reg(i int) *Reg {
	return NewReg(inst.Args[i], inst)
}

// useArg returns the value held by the given argument, emitting code to f.
func (f *Func) useArg(arg *Arg) value.Value {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		reg := NewReg(a, arg.parent)
		return f.useReg(reg)
	case x86asm.Mem:
		mem := NewMem(a, arg.parent)
		return f.useMem(mem)
	case x86asm.Imm:
		return constant.NewInt(int64(a), types.I32)
	case x86asm.Rel:
		next := arg.parent.addr + bin.Address(arg.parent.Len)
		addr := next + bin.Address(a)
		return f.useAddr(addr)
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg.Arg))
	}
}

// useArgElem returns a value of the specified element type held by the given
// argument, emitting code to f.
func (f *Func) useArgElem(arg *Arg, elem types.Type) value.Value {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		reg := NewReg(a, arg.parent)
		return f.useRegElem(reg, elem)
	case x86asm.Mem:
		mem := NewMem(a, arg.parent)
		return f.useMemElem(mem, elem)
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// defArg stores the value to the given argument, emitting code to f.
func (f *Func) defArg(arg *Arg, v value.Value) {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		reg := NewReg(a, arg.parent)
		f.defReg(reg, v)
	case x86asm.Mem:
		mem := NewMem(a, arg.parent)
		f.defMem(mem, v)
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// defArgElem stores the value of the specified element type to the given
// argument, emitting code to f.
func (f *Func) defArgElem(arg *Arg, v value.Value, elem types.Type) {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		reg := NewReg(a, arg.parent)
		f.defRegElem(reg, v, elem)
	case x86asm.Mem:
		mem := NewMem(a, arg.parent)
		f.defMemElem(mem, v, elem)
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// --- [ register ] ------------------------------------------------------------

// A Reg is a single x86 register.
type Reg struct {
	// x86 register.
	x86asm.Reg
	// Parent instruction; used to retrieve symbolic execution information.
	parent *Inst
}

// NewReg returns a new x86 register argument with the given parent instruction.
func NewReg(arg x86asm.Arg, parent *Inst) *Reg {
	reg, ok := arg.(x86asm.Reg)
	if !ok {
		panic(fmt.Errorf("invalid register argument type; expected x86asm.Reg, got %T", arg))
	}
	return &Reg{
		Reg:    reg,
		parent: parent,
	}
}

// useReg loads and returns a value from the given x86 register, emitting code
// to f.
func (f *Func) useReg(reg *Reg) value.Value {
	src := f.reg(reg.Reg)
	return f.cur.NewLoad(src)
}

// useRegElem loads and returns a value of the specified element type from the
// given x86 register, emitting code to f.
func (f *Func) useRegElem(reg *Reg, elem types.Type) value.Value {
	src := f.reg(reg.Reg)
	typ := types.NewPointer(elem)
	if !typ.Equal(src.Type()) {
		src = f.cur.NewBitCast(src, typ)
	}
	return f.cur.NewLoad(src)
}

// defReg stores the value to the given x86 register, emitting code to f.
func (f *Func) defReg(reg *Reg, v value.Value) {
	dst := f.reg(reg.Reg)
	f.cur.NewStore(v, dst)
	switch reg.Reg {
	case x86asm.EAX, x86asm.EDX:
		// Redefine the PSEUDO-register EDX:EAX based on change in EAX or EDX.
		f.redefEDX_EAX()
	}
}

// defRegElem stores the value of the specified element type to the given x86
// register, emitting code to f.
func (f *Func) defRegElem(reg *Reg, v value.Value, elem types.Type) {
	dst := f.reg(reg.Reg)
	typ := types.NewPointer(elem)
	if !typ.Equal(dst.Type()) {
		dst = f.cur.NewBitCast(dst, typ)
	}
	f.cur.NewStore(v, dst)
}

// reg returns a pointer to the LLVM IR value associated with the given x86
// register.
func (f *Func) reg(reg x86asm.Reg) value.Value {
	if v, ok := f.regs[reg]; ok {
		return v
	}
	typ := regType(reg)
	v := ir.NewAlloca(typ)
	name := strings.ToLower(Register(reg).String())
	v.SetName(name)
	f.regs[reg] = v
	return v
}

// --- [ memory reference ] ----------------------------------------------------

// A Mem is a memory reference.
type Mem struct {
	// x86 memory reference.
	x86asm.Mem
	// Parent instruction; used to retrieve symbolic execution information.
	parent *Inst
}

// Segment returns the segment register of the memory reference.
func (mem *Mem) Segment() *Reg {
	return NewReg(mem.Mem.Segment, mem.parent)
}

// Base returns the base register of the memory reference.
func (mem *Mem) Base() *Reg {
	return NewReg(mem.Mem.Base, mem.parent)
}

// Index returns the index register of the memory reference.
func (mem *Mem) Index() *Reg {
	return NewReg(mem.Mem.Index, mem.parent)
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
		parent: parent,
	}
}

// useMem loads and returns the value of the given memory reference, emitting
// code to f.
func (f *Func) useMem(mem *Mem) value.Value {
	src := f.mem(mem)
	return f.cur.NewLoad(src)
}

// useMemElem loads and returns a value of the specified element type from the
// given memory reference, emitting code to f.
func (f *Func) useMemElem(mem *Mem, elem types.Type) value.Value {
	src := f.mem(mem)
	typ := types.NewPointer(elem)
	if !typ.Equal(src.Type()) {
		src = f.cur.NewBitCast(src, typ)
	}
	return f.cur.NewLoad(src)
}

// defMem stores the value to the given memory reference, emitting code to f.
func (f *Func) defMem(mem *Mem, v value.Value) {
	dst := f.mem(mem)
	f.cur.NewStore(v, dst)
}

// defMemElem stores the value of the specified element type to the given memory
// reference, emitting code to f.
func (f *Func) defMemElem(mem *Mem, v value.Value, elem types.Type) {
	dst := f.mem(mem)
	typ := types.NewPointer(elem)
	if !typ.Equal(dst.Type()) {
		dst = f.cur.NewBitCast(dst, typ)
	}
	f.cur.NewStore(v, dst)
}

// mem returns a pointer to the LLVM IR value associated with the given memory
// argument, emitting code to f.
func (f *Func) mem(mem *Mem) value.Value {
	// Segment:[Base+Scale*Index+Disp].
	var (
		segment value.Value
		base    value.Value
		index   value.Value
		disp    value.Value
	)
	if mem.Mem.Segment != 0 {
		segment = f.useReg(mem.Segment())
	}
	if mem.Mem.Base != 0 {
		base = f.useReg(mem.Base())
	}
	if mem.Mem.Index != 0 {
		index = f.useReg(mem.Index())
	}

	// TODO: Add proper support for memory references.
	//    Segment Reg
	//    Base    Reg
	//    Scale   uint8
	//    Index   Reg
	//    Disp    int64

	// Handle local variables.
	if segment == nil && index == nil {
		// Stack local memory access.
		if mem.Mem.Base == x86asm.ESP || mem.Mem.Base == x86asm.EBP {
			name := fmt.Sprintf("%s_%d", strings.ToLower(Register(mem.Mem.Base).String()), f.espDisp+mem.Disp)
			if v, ok := f.locals[name]; ok {
				return v
			}
			v := ir.NewAlloca(types.I32)
			v.SetName(name)
			f.locals[name] = v
			return v
		}
	}

	// Handle disposition.
	if mem.Disp != 0 {
		if context, ok := f.d.contexts[mem.parent.addr]; ok {
			if c, ok := context.Args[1]; ok {
				if o, ok := c["Mem.offset"]; ok {
					offset := int64(o)
					addr := bin.Address(mem.Disp - offset)
					v, ok := f.addr(addr)
					if !ok {
						panic(fmt.Errorf("unable to locate value at address %v; referenced from %v instruction at %v", addr, mem.parent.Op, mem.parent.addr))
					}
					// TODO: Figure out how to handle negative offsets.
					disp = f.getElementPtr(v, offset)
				}
			}
		}
		if disp == nil {
			addr := bin.Address(mem.Disp)
			v, ok := f.addr(addr)
			if !ok {
				warn.Printf("unable to locate value at address %v; referenced from %v instruction at %v", addr, mem.parent.Op, mem.parent.addr)
			}
			disp = v
		}
	}

	// Early return for direct memory access.
	if segment == nil && base == nil && index == nil {
		if disp == nil {
			addr := bin.Address(mem.Disp)
			panic(fmt.Errorf("unable to locate value at address %v; referenced from %v instruction at %v", addr, mem.parent.Op, mem.parent.addr))
		}
		return disp
	}

	// TODO: Handle Segment.
	src := disp
	if segment != nil {
		// Ignore segments for now, assume byte addressing.
		//pretty.Println(mem)
		//panic("support for memory reference segment not yet implemented")
	}

	// Handle Base.
	if base != nil {
		if src == nil {
			src = base
		} else {
			src = f.castToPtr(src, mem.parent)
			indices := []value.Value{base}
			src = f.cur.NewGetElementPtr(src, indices...)
		}
	}

	// TODO: Handle Scale*Index.
	if index != nil {
		// TODO: Figure out how to handle scale. If we can validate that gep
		// indexes into elements of size `scale`, the scale can be safely ignored.
		if src == nil {
			src = index
		} else {
			src = f.castToPtr(src, mem.parent)
			indices := []value.Value{index}
			src = f.cur.NewGetElementPtr(src, indices...)
		}
	}

	// Handle dynamic memory reference.
	if src == nil {
		pretty.Println(mem)
		panic("unable to locate memory reference")
	}

	// TODO: Cast into proper type, once type analysis information is available.

	// Force bitcast into pointer type.
	return f.castToPtr(src, mem.parent)
}

// castToPtr casts the given value into a pointer, where the element type is
// derrived from src and instruction prefixes, with instruction prefix takes
// precedence.
func (f *Func) castToPtr(src value.Value, parent *Inst) value.Value {
	elem := src.Type()
	for _, prefix := range parent.Prefix[:] {
		// The first zero in the array marks the end of the prefixes.
		if prefix == 0 {
			break
		}
		// Validate that implied prefixes are supported.
		implied := prefix&x86asm.PrefixImplicit != 0
		if implied {
			switch prefix & 0x0FFF {
			case x86asm.PrefixData16:
				switch parent.Op {
				case x86asm.MOV, x86asm.MOVSW:
					// supported.
				default:
					panic(fmt.Errorf("support for implied prefix %v (0x%04X) not yet implemented for %v instruction", prefix, uint16(prefix), parent.Op))
				}
			default:
				panic(fmt.Errorf("support for prefix %v (0x%04X) not yet implemented", prefix, uint16(prefix)))
			}
		} else {
			switch prefix & 0x0FFF {
			case x86asm.PrefixData16:
				elem = types.I16
				panic(fmt.Errorf("inst prefix: %v", parent))
			default:
				panic(fmt.Errorf("support for prefix %v (0x%04X) not yet implemented", prefix, uint16(prefix)))
			}
		}
	}
	if !types.IsPointer(elem) {
		return f.cur.NewBitCast(src, types.NewPointer(elem))
	}
	return src
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

// useStatus loads and returns the value of the given x86 status flag, emitting
// code to f.
func (f *Func) useStatus(status StatusFlag) value.Value {
	src := f.status(status)
	return f.cur.NewLoad(src)
}

// defStatus stores the value to the given x86 status flag, emitting code to f.
func (f *Func) defStatus(status StatusFlag, v value.Value) {
	dst := f.status(status)
	f.cur.NewStore(v, dst)
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

// useAddr loads and returns the value of the given address, emitting code to f.
func (f *Func) useAddr(addr bin.Address) value.Named {
	src, ok := f.addr(addr)
	if !ok {
		panic(fmt.Errorf("unable to locate value at address %v", addr))
	}
	return f.cur.NewLoad(src)
}

// defAddr stores the value to the given address, emitting code to f.
func (f *Func) defAddr(addr bin.Address, v value.Value) {
	dst, ok := f.addr(addr)
	if !ok {
		panic(fmt.Errorf("unable to locate value at address %v", addr))
	}
	f.cur.NewStore(v, dst)
}

// addr returns a pointer to the LLVM IR value associated with the given
// address, emitting code to f. The returned value is one of *ir.BasicBlock,
// *ir.Global and *ir.Function, and the boolean value indicates success
func (f *Func) addr(addr bin.Address) (value.Named, bool) {
	if block, ok := f.blocks[addr]; ok {
		return block, true
	}
	if g, ok := f.d.globals[addr]; ok {
		return g, true
	}
	if fn, ok := f.d.funcs[addr]; ok {
		return fn.Function, true
	}
	// TODO: Add support for lookup of more globally addressable values.
	return nil, false
}

// === [ helpers ] =============================================================

// getAddr returns the static address represented by the given argument, and a
// boolean indicating success.
func (f *Func) getAddr(arg *Arg) (bin.Address, bool) {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		if context, ok := f.d.contexts[arg.parent.addr]; ok {
			if c, ok := context.Regs[Register(a)]; ok {
				if addr, ok := c["addr"]; ok {
					return bin.Address(addr), true
				}
			}
		}
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

	// Handle function pointers in structures.
	switch a := arg.Arg.(type) {
	case x86asm.Mem:
		if context, ok := f.d.contexts[arg.parent.addr]; ok {
			if c, ok := context.Regs[Register(a.Base)]; ok {
				if addr, ok := c["addr"]; ok {
					v := f.useAddr(bin.Address(addr))
					// TODO: Figure out how to handle negative offsets.
					v = f.getElementPtr(v, a.Disp)
					v = f.cur.NewLoad(v)
					if typ, ok := v.Type().(*types.PointerType); ok {
						if sig := typ.Elem.(*types.FuncType); ok {
							// TODO: Figure out how to recover calling convention.
							// Perhaps through context.json at call sites?
							return v, sig, ir.CallConvNone, true
						}
					}
				}
			}
		}
	}

	fmt.Printf("unable to locate function for argument %v\n", arg.Arg)
	panic("not yet implemented")
}

// redefEDX_EAX redefines the 64-bit pseudo-register EDX:EAX based on the value
// of EAX and EDX.
func (f *Func) redefEDX_EAX() {
	edx := f.useReg(EDX)
	eax := f.useReg(EAX)
	tmp1 := f.cur.NewSExt(edx, types.I64)
	tmp2 := f.cur.NewZExt(eax, types.I64)
	tmp := f.cur.NewOr(tmp1, tmp2)
	f.defReg(EDX_EAX, tmp)
}
