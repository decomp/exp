package lift

import (
	"fmt"
	"sort"
	"strings"

	"github.com/decomp/exp/bin"
	"github.com/decomp/exp/disasm/x86"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"golang.org/x/arch/x86/x86asm"
)

// === [ argument ] ============================================================

// useArg returns the value held by the given argument, emitting code to f.
func (f *Func) useArg(arg *x86.Arg) value.Value {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		reg := x86.NewReg(a, arg.Parent)
		return f.useReg(reg)
	case x86asm.Mem:
		mem := x86.NewMem(a, arg.Parent)
		return f.useMem(mem)
	case x86asm.Imm:
		return constant.NewInt(int64(a), types.I32)
	case x86asm.Rel:
		next := arg.Parent.Addr + bin.Address(arg.Parent.Len)
		addr := next + bin.Address(a)
		return f.useAddr(addr)
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg.Arg))
	}
}

// useArgElem returns a value of the specified element type held by the given
// argument, emitting code to f.
func (f *Func) useArgElem(arg *x86.Arg, elem types.Type) value.Value {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		reg := x86.NewReg(a, arg.Parent)
		return f.useRegElem(reg, elem)
	case x86asm.Mem:
		mem := x86.NewMem(a, arg.Parent)
		return f.useMemElem(mem, elem)
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// defArg stores the value to the given argument, emitting code to f.
func (f *Func) defArg(arg *x86.Arg, v value.Value) {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		reg := x86.NewReg(a, arg.Parent)
		f.defReg(reg, v)
	case x86asm.Mem:
		mem := x86.NewMem(a, arg.Parent)
		f.defMem(mem, v)
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// defArgElem stores the value of the specified element type to the given
// argument, emitting code to f.
func (f *Func) defArgElem(arg *x86.Arg, v value.Value, elem types.Type) {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		reg := x86.NewReg(a, arg.Parent)
		f.defRegElem(reg, v, elem)
	case x86asm.Mem:
		mem := x86.NewMem(a, arg.Parent)
		f.defMemElem(mem, v, elem)
	//case x86asm.Imm:
	//case x86asm.Rel:
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// === [ register ] ============================================================

// useReg loads and returns a value from the given x86 register, emitting code
// to f.
func (f *Func) useReg(reg *x86.Reg) value.Value {
	src := f.reg(reg.Reg)
	return f.cur.NewLoad(src)
}

// useRegElem loads and returns a value of the specified element type from the
// given x86 register, emitting code to f.
func (f *Func) useRegElem(reg *x86.Reg, elem types.Type) value.Value {
	src := f.reg(reg.Reg)
	typ := types.NewPointer(elem)
	if !typ.Equal(src.Type()) {
		src = f.cur.NewBitCast(src, typ)
	}
	return f.cur.NewLoad(src)
}

// defReg stores the value to the given x86 register, emitting code to f.
func (f *Func) defReg(reg *x86.Reg, v value.Value) {
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
func (f *Func) defRegElem(reg *x86.Reg, v value.Value, elem types.Type) {
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
	name := strings.ToLower(x86.Register(reg).String())
	v.SetName(name)
	f.regs[reg] = v
	return v
}

// === [ memory reference ] ====================================================

// useMem loads and returns the value of the given memory reference, emitting
// code to f.
func (f *Func) useMem(mem *x86.Mem) value.Value {
	src := f.mem(mem)
	return f.cur.NewLoad(src)
}

// useMemElem loads and returns a value of the specified element type from the
// given memory reference, emitting code to f.
func (f *Func) useMemElem(mem *x86.Mem, elem types.Type) value.Value {
	src := f.mem(mem)
	typ := types.NewPointer(elem)
	if !typ.Equal(src.Type()) {
		src = f.cur.NewBitCast(src, typ)
	}
	return f.cur.NewLoad(src)
}

// defMem stores the value to the given memory reference, emitting code to f.
func (f *Func) defMem(mem *x86.Mem, v value.Value) {
	dst := f.mem(mem)
	// Bitcast pointer to appropriate size.
	dst = f.castToPtr(dst, mem.Parent)
	f.cur.NewStore(v, dst)
}

// defMemElem stores the value of the specified element type to the given memory
// reference, emitting code to f.
func (f *Func) defMemElem(mem *x86.Mem, v value.Value, elem types.Type) {
	dst := f.mem(mem)
	typ := types.NewPointer(elem)
	if !typ.Equal(dst.Type()) {
		dst = f.cur.NewBitCast(dst, typ)
	}
	f.cur.NewStore(v, dst)
}

// mem returns a pointer to the LLVM IR value associated with the given memory
// argument, emitting code to f.
func (f *Func) mem(mem *x86.Mem) value.Value {
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
			name := fmt.Sprintf("%s_%d", strings.ToLower(x86.Register(mem.Mem.Base).String()), f.espDisp+mem.Disp)
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
		if context, ok := f.l.Contexts[mem.Parent.Addr]; ok {
			if c, ok := context.Args[1]; ok {
				if o, ok := c["Mem.offset"]; ok {
					offset := int64(o)
					addr := bin.Address(mem.Disp - offset)
					v, ok := f.addr(addr)
					if !ok {
						panic(fmt.Errorf("unable to locate value at address %v; referenced from %v instruction at %v", addr, mem.Parent.Op, mem.Parent.Addr))
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
				warn.Printf("unable to locate value at address %v; referenced from %v instruction at %v", addr, mem.Parent.Op, mem.Parent.Addr)
			}
			disp = v
		}
	}

	// Early return for direct memory access.
	if segment == nil && base == nil && index == nil {
		if disp == nil {
			addr := bin.Address(mem.Disp)
			panic(fmt.Errorf("unable to locate value at address %v; referenced from %v instruction at %v", addr, mem.Parent.Op, mem.Parent.Addr))
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
			src = f.castToPtr(src, mem.Parent)
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
			src = f.castToPtr(src, mem.Parent)
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
	return f.castToPtr(src, mem.Parent)
}

// castToPtr casts the given value into a pointer, where the element type is
// derrived from src and instruction prefixes, with instruction prefix takes
// precedence.
func (f *Func) castToPtr(src value.Value, parent *x86.Inst) value.Value {
	elem := src.Type()
	var preBits int
	if typ, ok := src.Type().(*types.PointerType); ok {
		elem = typ.Elem
		if elem, ok := elem.(*types.IntType); ok {
			preBits = elem.Size
		}
	}
	// Derive element size from the parent instruction.
	var bits int
	if parent != nil {
		if parent.MemBytes != 0 {
			bits = parent.MemBytes * 8
		}
		for _, prefix := range parent.Prefix[:] {
			// The first zero in the array marks the end of the prefixes.
			if prefix == 0 {
				break
			}
			switch prefix &^ x86asm.PrefixImplicit {
			case x86asm.PrefixData16:
				bits = 16
			case x86asm.PrefixREP, x86asm.PrefixREPN:
				// nothing to do.
			default:
				panic(fmt.Errorf("support for prefix %v (0x%04X) not yet implemented", prefix, uint16(prefix)))
			}
		}
	}
	if bits != 0 {
		elem = types.NewInt(bits)
	}
	needCast := !types.IsPointer(src.Type())
	if bits != 0 && preBits != 0 && bits != preBits {
		needCast = true
	}
	if needCast {
		typ := types.NewPointer(elem)
		var s string
		if v, ok := src.(value.Named); ok {
			if name := v.GetName(); len(name) > 0 {
				s = fmt.Sprintf(" %q", name)
			}
		}
		dbg.Printf("casting%s to pointer type: %v", s, typ)
		return f.cur.NewBitCast(src, typ)
	}
	return src
}

// === [ status flag ] =========================================================

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

// === [ address ] =============================================================

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
	// Direct or indirect access to global variable.
	if g, ok := f.global(addr); ok {
		return g, true
	}
	if fn, ok := f.l.Funcs[addr]; ok {
		return fn.Function, true
	}
	// TODO: Add support for lookup of more globally addressable values.
	return nil, false
}

// global returns a pointer to the LLVM IR value associated with the given
// global variable address, and a boolean value indicating success.
func (f *Func) global(addr bin.Address) (value.Named, bool) {
	// Early return if direct access to global variable.
	if src, ok := f.l.Globals[addr]; ok {
		return src, true
	}

	// Use binary search if indirect access to global variable (e.g. struct
	// field, array element).
	var globalAddrs []bin.Address
	for globalAddr := range f.l.Globals {
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
		g := f.l.Globals[start]
		size := f.l.sizeOfType(g.Typ.Elem)
		end := start + bin.Address(size)
		if start <= addr && addr < end {
			offset := int64(addr - start)
			return f.getElementPtr(g, offset), true
		}
	}
	return nil, false
}

// ### [ helpers ] #############################################################

// getAddr returns the static address represented by the given argument, and a
// boolean indicating success.
func (f *Func) getAddr(arg *x86.Arg) (bin.Address, bool) {
	switch a := arg.Arg.(type) {
	case x86asm.Reg:
		if context, ok := f.l.Contexts[arg.Parent.Addr]; ok {
			if c, ok := context.Regs[x86.Register(a)]; ok {
				if addr, ok := c["addr"]; ok {
					return bin.Address(addr), true
				}
			}
		}
	case x86asm.Rel:
		next := arg.Parent.Addr + bin.Address(arg.Parent.Len)
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
func (f *Func) getFunc(arg *x86.Arg) (value.Named, *types.FuncType, ir.CallConv, bool) {
	if addr, ok := f.getAddr(arg); ok {
		if fn, ok := f.l.Funcs[addr]; ok {
			v := fn.Function
			return v, v.Sig, v.CallConv, true
		}
	}

	// Handle function pointers in structures.
	switch a := arg.Arg.(type) {
	case x86asm.Mem:
		if a.Base != 0 {
			context, ok := f.l.Contexts[arg.Parent.Addr]
			if !ok {
				pretty.Println(arg.Arg)
				panic(fmt.Errorf("unable to locate context for %v register used at %v", a.Base, arg.Parent.Addr))
			}
			if c, ok := context.Regs[x86.Register(a.Base)]; ok {
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
					panic(fmt.Errorf("invalid callee type; expected pointer to function type, got %v", v.Type()))
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
	if !f.usesEDX_EAX {
		return
	}
	edx := f.useReg(x86.EDX)
	eax := f.useReg(x86.EAX)
	tmp1 := f.cur.NewSExt(edx, types.I64)
	tmp2 := f.cur.NewZExt(eax, types.I64)
	tmp := f.cur.NewOr(tmp1, tmp2)
	f.defReg(x86.EDX_EAX, tmp)
}