package main

import (
	"fmt"
	"sort"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
	"golang.org/x/arch/x86/x86asm"
)

// function tracks the information required to translate a function from x86
// machine code to LLVM IR assembly.
type function struct {
	// LLVM IR code for the function.
	*ir.Function
	// Entry address of the function.
	entry bin.Address
	// Basic blocks of the function.
	blocks map[bin.Address]*basicBlock
	// Registers used within the function.
	regs map[x86asm.Reg]*ir.InstAlloca
	// Status flags used within the function.
	status map[StatusFlag]*ir.InstAlloca
}

// basicBlock tracks the information required to translate a basic block from
// x86 machine code to LLVM IR assembly.
type basicBlock struct {
	// LLVM IR code for the basic block.
	*ir.BasicBlock
	// Entry address of the basic block.
	addr bin.Address
	// Instructions of the basic block.
	insts []*instruction
	// Terminator of the basic block.
	term *instruction
	// Additional basic blocks used when translation of single x86 basic blocks
	// require multiple LLVM IR basic blocks.
	extra []*basicBlock
}

// instruction tracks the information required to translate an instruction from
// x86 machine code to LLVM IR assembly.
type instruction struct {
	// Decoded x86 instruction.
	x86asm.Inst
	// Address of the instruction.
	addr bin.Address
}

// decodeFunc decodes the x86 machine code of the function at the given address.
func (d *disassembler) decodeFunc(entry bin.Address) (*function, error) {
	dbg.Printf("decoding function at %v", entry)
	f, ok := d.funcs[entry]
	if !ok {
		f = &function{
			entry:  entry,
			blocks: make(map[bin.Address]*basicBlock),
			regs:   make(map[x86asm.Reg]*ir.InstAlloca),
			status: make(map[StatusFlag]*ir.InstAlloca),
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
		targets := d.targets(entry, block.term)
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
	if d.decodedBlock[addr] {
		panic(fmt.Errorf("decoded basic block at %v twice", addr))
	}
	d.decodedBlock[addr] = true
	dbg.Printf("decoding basic block at %v", addr)
	// Access the data of the executable at the given address.
	src, err := d.data(addr)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Calculate maximum basic block length, based on the address of the
	// succeeding chunk.
	maxLen := d.getMaxBlockLen(addr)

	// Decode instructions.
	label := fmt.Sprintf("block_%06X", uint64(addr))
	block := &basicBlock{
		BasicBlock: &ir.BasicBlock{
			Name: label,
		},
		addr: addr,
	}
	for j := int64(0); j < maxLen; {
		inst, err := x86asm.Decode(src, d.mode)
		if err != nil {
			// TODO: Remove debug info when the disassembler matures.
			printBlock(block)
			fmt.Println("addr:", addr)
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
		if i.isTerm() {
			if j != maxLen {
				panic(fmt.Errorf("basic block length mismatch; expected %d, got %d", maxLen, j))
			}
			break
		}
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

// targets returns the target addresses of the given terminator. Entry specifies
// the entry address of the function to which the terminator belongs.
func (d *disassembler) targets(entry bin.Address, term *instruction) []bin.Address {
	// dummy terminator denoted with x86asm.Inst zero value.
	if term.isDummyTerm() {
		return []bin.Address{term.addr}
	}
	switch term.Op {
	case x86asm.JA, x86asm.JAE, x86asm.JB, x86asm.JBE, x86asm.JCXZ, x86asm.JE, x86asm.JECXZ, x86asm.JG, x86asm.JGE, x86asm.JL, x86asm.JLE, x86asm.JNE, x86asm.JNO, x86asm.JNP, x86asm.JNS, x86asm.JO, x86asm.JP, x86asm.JRCXZ, x86asm.JS:
		// target branch of conditional branching instruction.
		next := term.addr + bin.Address(term.Len)
		targetsTrue := d.getAddrs(next, term.Args[0])
		// fallthrough branch of conditional branching instruction.
		targetFalse := next
		return append(targetsTrue, targetFalse)
	case x86asm.JMP:
		if d.isTailCall(entry, term) {
			dbg.Printf("tail call at %v", term.addr)
			// no target branches for tail calls.
			return nil
		}
		// target branch of JMP instruction.
		next := term.addr + bin.Address(term.Len)
		targets := d.getAddrs(next, term.Args[0])
		return targets
	case x86asm.RET:
		// no target branches.
		return nil
	default:
		panic(fmt.Errorf("support for terminator opcode %v not yet implemented", term.Op))
	}
}

// isTailCall reports whether the given instruction is a tail call instruction.
func (d *disassembler) isTailCall(funcEntry bin.Address, inst *instruction) bool {
	funcEnd := d.getFuncEndAddr(funcEntry)
	next := inst.addr + bin.Address(inst.Len)
	switch arg := inst.Args[0].(type) {
	//case x86asm.Reg:
	case x86asm.Mem:
		target := bin.Address(arg.Disp)
		if _, ok := d.tables[target]; ok {
			// Target read from jump table (e.g. switch statement).
			return false
		}
		if funcEntry <= target && target < funcEnd {
			// Target inside function.
			return false
		}
		if funcAddr, ok := d.chunkFunc[target]; ok && funcAddr == funcEntry {
			// Target inside function chunk.
			return false
		}
		if target < d.getCodeStart() && arg.Base != 0 {
			// Target function pointer; read from memory [base reg + disp imm].
			//
			// Note, this may be a false assumption.
			//
			// TODO: Validate this assumption once type analysis information is
			// available.
			return true
		}
		if d.isImport(target) {
			// Call to imported function.
			return true
		}
		if !d.isFunc(target) {
			fmt.Println("arg:", arg)
			pretty.Println(arg)
			panic(fmt.Errorf("tail call to non-function address %v", target))
		}
		return true
	//case x86asm.Imm:
	case x86asm.Rel:
		target := next + bin.Address(arg)
		if funcEntry <= target && target < funcEnd {
			// Target inside function.
			return false
		}
		if funcAddr, ok := d.chunkFunc[target]; ok && funcAddr == funcEntry {
			// Target inside function chunk.
			return false
		}
		if d.isImport(target) {
			// Call to imported function.
			return true
		}
		if !d.isFunc(target) {
			fmt.Println("arg:", arg)
			pretty.Println(arg)
			panic(fmt.Errorf("tail call to non-function address %v", target))
		}
		return true
	default:
		fmt.Println("arg:", arg)
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// isImport reports whether the given address is part of the `.idata` section.
func (d *disassembler) isImport(addr bin.Address) bool {
	start := bin.Address(d.imageBase + d.idataBase)
	end := start + bin.Address(d.idataSize)
	return start <= addr && addr < end
}

// isFunc reports whether the given address is the entry address of a function.
func (d *disassembler) isFunc(addr bin.Address) bool {
	less := func(i int) bool {
		return addr <= d.funcAddrs[i]
	}
	index := sort.Search(len(d.funcAddrs), less)
	if index < len(d.funcAddrs) {
		return d.funcAddrs[index] == addr
	}
	return false
}

// getAddrs returns the addresses specified given argument.
func (d *disassembler) getAddrs(next bin.Address, arg x86asm.Arg) []bin.Address {
	switch arg := arg.(type) {
	//case x86asm.Reg:
	case x86asm.Mem:
		disp := bin.Address(arg.Disp)
		if targets, ok := d.tables[disp]; ok {
			return targets
		}
		fmt.Println("arg:", arg)
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	//case x86asm.Imm:
	case x86asm.Rel:
		return []bin.Address{next + bin.Address(arg)}
	default:
		fmt.Println("arg:", arg)
		pretty.Println(arg)
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}

// getMaxBlockLen returns the maximum length of the given basic block.
func (d *disassembler) getMaxBlockLen(blockAddr bin.Address) int64 {
	less := func(i int) bool {
		return blockAddr < d.chunks[i].addr
	}
	index := sort.Search(len(d.chunks), less)
	if index < len(d.chunks) {
		return int64(d.chunks[index].addr - blockAddr)
	}
	return int64(d.getCodeEnd() - blockAddr)
}

// getFuncEndAddr returns the end address of the given function.
func (d *disassembler) getFuncEndAddr(entry bin.Address) bin.Address {
	less := func(i int) bool {
		return entry < d.funcAddrs[i]
	}
	index := sort.Search(len(d.funcAddrs), less)
	if index < len(d.funcAddrs) {
		return d.funcAddrs[index]
	}
	return d.getCodeEnd()
}

// getCodeStart returns the start address of the code section.
func (d *disassembler) getCodeStart() bin.Address {
	return bin.Address(d.imageBase + d.codeBase)
}

// getCodeEnd returns the end address of the code section.
func (d *disassembler) getCodeEnd() bin.Address {
	return d.getCodeStart() + bin.Address(d.codeSize)
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
	if len(q.addrs) == 0 {
		panic("invalid call to pop; empty queue")
	}
	var min bin.Address
	for addr := range q.addrs {
		if min == 0 || addr < min {
			min = addr
		}
	}
	delete(q.addrs, min)
	return min
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
	if block.term == nil {
		fmt.Println("; ### terminator missing in basic block")
	} else if block.term.isDummyTerm() {
		fmt.Println("; dummy terminator")
	} else {
		fmt.Println(block.term)
	}
}
