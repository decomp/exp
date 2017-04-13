package main

import (
	"fmt"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
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
	insts []*instruction
	term  *instruction
}

type instruction struct {
	x86asm.Inst
	addr bin.Address
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
	queue := newQueue()
	for queue.push(entry); !queue.empty(); {
		addr := queue.pop()
		block, err := d.decodeBlock(addr)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		f.blocks[addr] = block
		// Add terminators to queue, if not already decoded.
		targets := d.targets(block.term)
		for _, target := range targets {
			if _, ok := f.blocks[target]; ok {
				// ignore block if already decoded.
				continue
			}
			// add block to queue.
			dbg.Printf("adding basic block address %v to queue", target)
			queue.push(target)
		}
	}
	return f, nil
}

// decodeBlock decodes the x86 machine code of the basic block at the given
// address.
func (d *disassembler) decodeBlock(addr bin.Address) (*basicBlock, error) {
	// Access the data of the executable at the given address.
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
	maxLen := int64(end - addr)

	// Decode instructions.
	block := &basicBlock{
		addr: addr,
	}
	for j := int64(0); j < maxLen; {
		inst, err := x86asm.Decode(src, d.mode)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		i := &instruction{
			Inst: inst,
			addr: addr,
		}
		block.insts = append(block.insts, i)
		j += int64(inst.Len)
		src = src[inst.Len:]
		addr += bin.Address(inst.Len)
	}
	lastInst := block.insts[len(block.insts)-1]
	if lastInst.isTerm() {
		block.insts = block.insts[:len(block.insts)-1]
		block.term = lastInst
	} else {
		// TODO: Figure out a better representation for dummy terminators.

		// dummy terminator denoted the zero value for x86asm.Inst and address of
		// fallthrough basic block.
		block.term = &instruction{
			addr: lastInst.addr + bin.Address(lastInst.Len),
		}
	}
	return block, nil
}

// isTerm reports whether the given instruction is a terminating instruction.
func (inst *instruction) isTerm() bool {
	switch inst.Op {
	case x86asm.JA, x86asm.JAE, x86asm.JB, x86asm.JBE, x86asm.JCXZ, x86asm.JE, x86asm.JECXZ, x86asm.JG, x86asm.JGE, x86asm.JL, x86asm.JLE, x86asm.JMP, x86asm.JNE, x86asm.JNO, x86asm.JNP, x86asm.JNS, x86asm.JO, x86asm.JP, x86asm.JRCXZ, x86asm.JS:
		return true
	case x86asm.RET:
		return true
	}
	return false
}

// isDummyTerm reports whether the given instruction is a dummy terminating
// instruction. Dummy terminators are used when a basic block is missing a
// terminator and falls through into the succeeding basic block, the address of
// which is denoted by term.addr.
func (inst *instruction) isDummyTerm() bool {
	zero := x86asm.Inst{}
	return inst.Inst == zero
}

// targets returns the target addresses of the given terminator.
func (d *disassembler) targets(term *instruction) []bin.Address {
	// dummy terminator denoted with x86asm.Inst zero value.
	if term.isDummyTerm() {
		return []bin.Address{term.addr}
	}
	switch term.Op {
	case x86asm.JA, x86asm.JAE, x86asm.JB, x86asm.JBE, x86asm.JCXZ, x86asm.JE, x86asm.JECXZ, x86asm.JG, x86asm.JGE, x86asm.JL, x86asm.JLE, x86asm.JNE, x86asm.JNO, x86asm.JNP, x86asm.JNS, x86asm.JO, x86asm.JP, x86asm.JRCXZ, x86asm.JS:
		// target branch of conditional branching instruction.
		base := term.addr + bin.Address(term.Len)
		targetTrue := d.getAddr(base, term.Args[0])
		// fallthrough branch of conditional branching instruction.
		targetFalse := base
		return []bin.Address{targetFalse, targetTrue}
	case x86asm.JMP:
		// target branch of JMP instruction.
		target := d.getAddr(term.addr, term.Args[0])
		return []bin.Address{target}
	case x86asm.RET:
		// no target branches.
		return nil
	default:
		panic(fmt.Errorf("support for terminator opcode %v not yet implemented", term.Op))
	}
}

// getAddr returns the address specified by the original address and given
// argument.
func (d *disassembler) getAddr(base bin.Address, arg x86asm.Arg) bin.Address {
	switch arg := arg.(type) {
	case x86asm.Rel:
		return base + bin.Address(arg)
	default:
		fmt.Println("arg:", arg)
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// ### [ Helper functions ] ####################################################

// queue represents a queue of addresses.
type queue struct {
	// Addresses in the queue.
	addrs map[bin.Address]bool
}

// newQueue returns a new queue.
func newQueue() *queue {
	return &queue{
		addrs: make(map[bin.Address]bool),
	}
}

// push pushes the given address to the queue.
func (q *queue) push(addr bin.Address) {
	q.addrs[addr] = true
}

// pop pops an address from the queue.
func (q *queue) pop() bin.Address {
	for addr := range q.addrs {
		delete(q.addrs, addr)
		return addr
	}
	panic("invalid call to pop; empty queue")
}

// empty reports whether the queue is empty.
func (q *queue) empty() bool {
	return len(q.addrs) == 0
}

// printFunc pretty-prints the given function.
func printFunc(f *function) {
	var blockAddrs []bin.Address
	for blockAddr := range f.blocks {
		blockAddrs = append(blockAddrs, blockAddr)
	}
	sort.Sort(bin.Addresses(blockAddrs))
	for _, blockAddr := range blockAddrs {
		block := f.blocks[blockAddr]
		printBlock(block)
	}
}

// printBlock pretty-prints the given basic block.
func printBlock(block *basicBlock) {
	for _, inst := range block.insts {
		fmt.Println(inst)
	}
	if block.term.isDummyTerm() {
		fmt.Println("; dummy terminator")
	} else {
		fmt.Println(block.term)
	}
}
