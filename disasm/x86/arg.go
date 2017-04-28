package x86

import (
	"fmt"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"golang.org/x/arch/x86/x86asm"
)

// Addrs returns the addresses specified by the given argument. Addr specifies
// the address of the terminator, and next the address of the next instruction.
func (dis *Disasm) Addrs(arg x86asm.Arg, addr, next bin.Address) []bin.Address {
	switch arg := arg.(type) {
	//case x86asm.Reg:
	case x86asm.Mem:
		// Segment:[Base+Scale*Index+Disp].

		//    Segment Reg
		//    Base    Reg
		//    Scale   uint8
		//    Index   Reg
		//    Disp    int64

		// Static target.
		disp := bin.Address(arg.Disp)
		if arg.Segment == 0 && arg.Base == 0 && arg.Index == 0 {
			return []bin.Address{disp}
		}

		// Adjust disposition based on index register value.
		if arg.Index != 0 {
			if context, ok := dis.Contexts[addr]; ok {
				if c, ok := context.Regs[Register(arg.Index)]; ok {
					if indexMin, ok := c["min"]; ok {
						disp += bin.Address(arg.Scale) * bin.Address(indexMin)
					}
				}
			}
		}

		// Jump table target.
		if targets, ok := dis.Tables[disp]; ok {
			if arg.Segment == 0 && arg.Base == 0 && arg.Scale == 4 && arg.Index != 0 {
				return targets
			}
		}

		// TODO: Figure out how to handle indirect jump to function pointer.

		// Target is likely a function pointer; skip for now.
		if arg.Base != 0 && disp < dis.codeStart() {
			warn.Printf("ignoring indirect targets from %v of memory reference %v", addr, arg)
			return nil
		}

		pretty.Println("mem:", arg)
		panic("x86.Disasm.Addrs: not yet implemented")
	//case x86asm.Imm:
	case x86asm.Rel:
		target := next + bin.Address(arg)
		return []bin.Address{target}
	default:
		panic(fmt.Errorf("support for argument type %T not yet implemented", arg))
	}
}
