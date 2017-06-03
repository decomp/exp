package x86

import (
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/pkg/errors"
	"golang.org/x/arch/x86/x86asm"
)

// A Func is a function.
type Func struct {
	// Address of the function.
	Addr bin.Address
	// Basic blocks of the function.
	Blocks map[bin.Address]*BasicBlock
}

// A BasicBlock is a basic block; a sequence of non-branching instructions
// terminated by a branching instruction.
type BasicBlock struct {
	// Address of the basic block.
	Addr bin.Address
	// Sequence of non-branching instructions.
	Insts []*Inst
	// Terminating instruction.
	Term *Inst
}

// An Inst is a single instruction.
type Inst struct {
	// Address of the instruction.
	Addr bin.Address
	// x86 instruction.
	x86asm.Inst
}

// DecodeFunc decodes and returns the function at the given address.
func (dis *Disasm) DecodeFunc(entry bin.Address) (*Func, error) {
	dbg.Printf("decoding function at %v", entry)
	f := &Func{
		Addr:   entry,
		Blocks: make(map[bin.Address]*BasicBlock),
	}
	queue := newQueue()
	queue.push(entry)
	for !queue.empty() {
		blockAddr := queue.pop()
		if _, ok := f.Blocks[blockAddr]; ok {
			// skip basic block if already decoded.
			continue
		}
		block, err := dis.DecodeBlock(blockAddr)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		f.Blocks[blockAddr] = block
		// Add block targets to queue.
		targets := dis.Targets(block.Term, entry)
		for _, target := range targets {
			dbg.Printf("adding basic block address %v to queue", target)
			queue.push(target)
		}
	}
	return f, nil
}

// DecodeBlock decodes and returns the basic block at the given address.
func (dis *Disasm) DecodeBlock(entry bin.Address) (*BasicBlock, error) {
	dbg.Printf("decoding basic block at %v", entry)
	// Compute end address of the basic block.
	maxLen := dis.maxBlockLen(entry)
	addr := entry
	end := entry + bin.Address(maxLen)
	// Decode instructions.
	block := &BasicBlock{
		Addr: entry,
	}
	for addr < end {
		inst, err := dis.DecodeInst(addr)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		dbg.Printf("   instruction at %v: %v", addr, inst)
		addr += bin.Address(inst.Len)
		if inst.isTerm() {
			block.Term = inst
			break
		}
		block.Insts = append(block.Insts, inst)
	}
	// Sanity check.
	if addr != end {
		warn.Printf("unexpected end address of basic block at %v; expected %v, got %v", entry, end, addr)
	}
	// Add dummy terminator for fallthrough basic blocks.
	if block.Term == nil {
		block.Term = &Inst{
			Addr: end,
		}
	}
	return block, nil
}

// DecodeInst decodes and returns the instruction at the given address.
func (dis *Disasm) DecodeInst(addr bin.Address) (*Inst, error) {
	code := dis.File.Code(addr)
	i, err := x86asm.Decode(code, dis.Mode)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	inst := &Inst{
		Addr: addr,
		Inst: i,
	}
	return inst, nil
}

// maxBlockLen returns the maximum length of the given basic block.
func (dis *Disasm) maxBlockLen(blockAddr bin.Address) int64 {
	less := func(i int) bool {
		return blockAddr < dis.Frags[i].Addr
	}
	index := sort.Search(len(dis.Frags), less)
	if 0 <= index && index < len(dis.Frags) {
		return int64(dis.Frags[index].Addr - blockAddr)
	}
	return int64(dis.codeEnd() - blockAddr)
}

// codeStart returns the start address of the first code section.
func (dis *Disasm) codeStart() bin.Address {
	var min bin.Address
	for _, sect := range dis.File.Sections {
		if sect.Perm&bin.PermX != 0 {
			start := sect.Addr
			if min == 0 || min > start {
				min = start
			}
		}
	}
	if min == 0 {
		panic("unable to locate start address of first code section")
	}
	return min
}

// codeEnd returns the end address of the last code section.
func (dis *Disasm) codeEnd() bin.Address {
	var max bin.Address
	for _, sect := range dis.File.Sections {
		if sect.Perm&bin.PermX != 0 {
			end := sect.Addr + bin.Address(len(sect.Data))
			if max < end {
				max = end
			}
		}
	}
	if max == 0 {
		panic("unable to locate end address of last code section")
	}
	return max
}
