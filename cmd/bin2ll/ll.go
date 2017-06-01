//+build ignore

package main

import (
	"fmt"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
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

	// Preprocess the function to assess if any instruction makes use of EDX:EAX
	// (e.g. IDIV).
	for _, bb := range f.bbs {
		for _, inst := range bb.insts {
			switch inst.Op {
			// TODO: Identify more instructions which makes use of EDX:EAX.
			case x86asm.IDIV:
				f.usesEDX_EAX = true
			}
		}
	}

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
		for reg := firstReg; reg <= lastReg; reg++ {
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
	if err := f.emitTerm(bb.term); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
