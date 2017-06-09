package lift

import (
	"fmt"

	"github.com/decomp/exp/disasm/x86"
	"github.com/llir/llvm/ir/types"
	"golang.org/x/arch/x86/x86asm"
)

// regType returns the LLVM IR type of the given register.
func regType(reg x86asm.Reg) types.Type {
	switch reg {
	// 8-bit
	case x86asm.AL, x86asm.CL, x86asm.DL, x86asm.BL, x86asm.AH, x86asm.CH, x86asm.DH, x86asm.BH, x86asm.SPB, x86asm.BPB, x86asm.SIB, x86asm.DIB, x86asm.R8B, x86asm.R9B, x86asm.R10B, x86asm.R11B, x86asm.R12B, x86asm.R13B, x86asm.R14B, x86asm.R15B:
		return types.I8
	// 16-bit
	case x86asm.AX, x86asm.CX, x86asm.DX, x86asm.BX, x86asm.SP, x86asm.BP, x86asm.SI, x86asm.DI, x86asm.R8W, x86asm.R9W, x86asm.R10W, x86asm.R11W, x86asm.R12W, x86asm.R13W, x86asm.R14W, x86asm.R15W:
		return types.I16
	// 32-bit
	case x86asm.EAX, x86asm.ECX, x86asm.EDX, x86asm.EBX, x86asm.ESP, x86asm.EBP, x86asm.ESI, x86asm.EDI, x86asm.R8L, x86asm.R9L, x86asm.R10L, x86asm.R11L, x86asm.R12L, x86asm.R13L, x86asm.R14L, x86asm.R15L:
		return types.I32
	// 64-bit
	case x86asm.RAX, x86asm.RCX, x86asm.RDX, x86asm.RBX, x86asm.RSP, x86asm.RBP, x86asm.RSI, x86asm.RDI, x86asm.R8, x86asm.R9, x86asm.R10, x86asm.R11, x86asm.R12, x86asm.R13, x86asm.R14, x86asm.R15:
		return types.I64
	// Instruction pointer.
	case x86asm.IP: // 16-bit
		return types.I16
	case x86asm.EIP: // 32-bit
		return types.I32
	case x86asm.RIP: // 64-bit
		return types.I64
	// 387 floating point registers.
	case x86asm.F0, x86asm.F1, x86asm.F2, x86asm.F3, x86asm.F4, x86asm.F5, x86asm.F6, x86asm.F7:
		return types.X86_FP80
	// MMX registers.
	case x86asm.M0, x86asm.M1, x86asm.M2, x86asm.M3, x86asm.M4, x86asm.M5, x86asm.M6, x86asm.M7:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// XMM registers.
	case x86asm.X0, x86asm.X1, x86asm.X2, x86asm.X3, x86asm.X4, x86asm.X5, x86asm.X6, x86asm.X7, x86asm.X8, x86asm.X9, x86asm.X10, x86asm.X11, x86asm.X12, x86asm.X13, x86asm.X14, x86asm.X15:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// Segment registers.
	case x86asm.ES, x86asm.CS, x86asm.SS, x86asm.DS, x86asm.FS, x86asm.GS:
		return types.I16
	// System registers.
	case x86asm.GDTR:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	case x86asm.IDTR:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	case x86asm.LDTR:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	case x86asm.MSW:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	case x86asm.TASK:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// Control registers.
	case x86asm.CR0, x86asm.CR1, x86asm.CR2, x86asm.CR3, x86asm.CR4, x86asm.CR5, x86asm.CR6, x86asm.CR7, x86asm.CR8, x86asm.CR9, x86asm.CR10, x86asm.CR11, x86asm.CR12, x86asm.CR13, x86asm.CR14, x86asm.CR15:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// Debug registers.
	case x86asm.DR0, x86asm.DR1, x86asm.DR2, x86asm.DR3, x86asm.DR4, x86asm.DR5, x86asm.DR6, x86asm.DR7, x86asm.DR8, x86asm.DR9, x86asm.DR10, x86asm.DR11, x86asm.DR12, x86asm.DR13, x86asm.DR14, x86asm.DR15:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// Task registers.
	case x86asm.TR0, x86asm.TR1, x86asm.TR2, x86asm.TR3, x86asm.TR4, x86asm.TR5, x86asm.TR6, x86asm.TR7:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	// PSEUDO-registers.
	// TODO: Figure out how to handle edx:eax registers.
	case x86.X86asm_DX_AX:
		return types.I32
	case x86.X86asm_EDX_EAX:
		return types.I64
	case x86.X86asm_RDX_RAX:
		return types.I128
	default:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	}
}
