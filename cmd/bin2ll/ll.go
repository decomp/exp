package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewbak/x86/x86asm"
	"github.com/pkg/errors"
)

// translateFunc translates the given function from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) translateFunc(f *function) error {
	if f.Function == nil {
		f.Function = &ir.Function{}
	}
	var blocks []*basicBlock
	for _, block := range f.blocks {
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
		switch f.callconv {
		case "__fastcall":
			if ecx, ok := f.regs[x86asm.ECX]; ok {
				entry.NewStore(f.Sig.Params[0], ecx)
			}
			if edx, ok := f.regs[x86asm.EDX]; ok {
				entry.NewStore(f.Sig.Params[1], edx)
			}
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
	block.BasicBlock = &ir.BasicBlock{}
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
	case x86asm.MOV:
		return d.instMOV(f, block, inst)
	default:
		panic(fmt.Errorf("support for instruction opcode %v not yet implemented", inst.Op))
	}
}

// instAND translates the given AND instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instAND(f *function, block *basicBlock, inst *instruction) error {
	x := d.useArg(f, block, inst.Args[0])
	y := d.useArg(f, block, inst.Args[1])
	result := block.NewAnd(x, y)
	d.defArg(f, block, inst.Args[0], result)
	return nil
}

// instMOV translates the given MOV instruction from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) instMOV(f *function, block *basicBlock, inst *instruction) error {
	y := d.useArg(f, block, inst.Args[1])
	d.defArg(f, block, inst.Args[0], y)
	return nil
}

// translateTerm translates the given terminator from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) translateTerm(f *function, block *basicBlock, term *instruction) error {
	fmt.Println("term:", term)
	switch term.Op {
	case x86asm.RET:
		// TODO: Add support for non-void returns.
		block.NewRet(nil)
		return nil
	default:
		panic(fmt.Errorf("support for terminator opcode %v not yet implemented", term.Op))
	}
}

func (d *disassembler) useArg(f *function, block *basicBlock, arg x86asm.Arg) value.Value {
	fmt.Println("useArg:", arg)
	switch arg := arg.(type) {
	case x86asm.Reg:
		return d.reg(f, arg)
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
		}
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	case x86asm.Imm:
		return constant.NewInt(int64(arg), types.I32)
	//case x86asm.Rel:
	default:
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

func (d *disassembler) defArg(f *function, block *basicBlock, arg x86asm.Arg, v value.Value) {
	fmt.Println("defArg:", arg)
	switch arg := arg.(type) {
	//case x86asm.Reg:
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

func (d *disassembler) reg(f *function, reg x86asm.Reg) value.Value {
	if v, ok := f.regs[reg]; ok {
		return v
	}
	v := ir.NewAlloca(types.I32)
	v.SetName(strings.ToLower(reg.String()))
	f.regs[reg] = v
	return v
}
