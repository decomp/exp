package main

import (
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/llir/llvm/ir"
	"github.com/mewbak/x86/x86asm"
	"github.com/pkg/errors"
)

type function struct {
	*ir.Function
	entry  bin.Address
	blocks map[bin.Address]*basicBlock
}

type basicBlock struct {
	*ir.BasicBlock
	addr  bin.Address
	insts []x86asm.Inst
}

// decodeFunc decodes the x86 machine code of the function at the given address.
func (d *disassembler) decodeFunc(entry bin.Address) (*function, error) {
	f, ok := d.funcs[entry]
	if !ok {
		f = &function{
			entry:  entry,
			blocks: make(map[bin.Address]*basicBlock),
		}
	}
	queue := make(map[bin.Address]bool)
	queue[entry] = true
	for addr := range queue {
		block, err := d.decodeBlock(addr)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		f.blocks[addr] = block
		delete(queue, addr)
		// Add terminators to queue, if not already decoded.
		// TODO: Implement.
	}
	return f, nil
}

// decodeBlock decodes the x86 machine code of the basic block at the given
// address.
func (d *disassembler) decodeBlock(addr bin.Address) (*basicBlock, error) {
	// Access byte slice .
	src, err := d.data(addr)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Calculate maximum basic block length, based on the address of the
	// succeeding chunk.
	less := func(i int) bool {
		return d.chunks[i].addr > addr
	}
	index := sort.Search(len(d.chunks), less)
	end := d.chunks[index].addr
	maxLen := int(end - addr)

	// Decode instructions.
	block := &basicBlock{
		addr: addr,
	}
	for i := 0; i < maxLen; {
		inst, err := x86asm.Decode(src, d.mode)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		block.insts = append(block.insts, inst)
		i += inst.Len
		src = src[inst.Len:]
	}
	return block, nil
}
