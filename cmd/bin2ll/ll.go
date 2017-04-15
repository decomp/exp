package main

// translateFunc translates the given function from x86 machine code to LLVM IR
// assembly.
func (d *disassembler) translateFunc(f *function) error {
	panic("not yet implemented")
}

// translateBlock translates the given basic block from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) translateBlock(f *function, block *basicBlock) error {
	panic("not yet implemented")
}

// translateInst translates the given instruction from x86 machine code to LLVM
// IR assembly.
func (d *disassembler) translateInst(f *function, block *basicBlock, inst *instruction) error {
	panic("not yet implemented")
}
