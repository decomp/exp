package x86

import (
	"fmt"

	"golang.org/x/arch/x86/x86asm"
)

// A Register is a single register.
type Register x86asm.Reg

// String returns the string representation of reg.
func (reg Register) String() string {
	r := x86asm.Reg(reg)
	// Pretty-print pseudo-registers.
	m := map[x86asm.Reg]string{
		x86asm_EDX_EAX: "EDX:EAX",
	}
	if s, ok := m[r]; ok {
		return s
	}
	return r.String()
}

// Set sets reg to the register represented by s.
func (reg *Register) Set(s string) error {
	*reg = Register(parseReg(s))
	return nil
}

// UnmarshalText unmarshals the text into reg.
func (reg *Register) UnmarshalText(text []byte) error {
	return reg.Set(string(text))
}

// MarshalText returns the textual representation of reg.
func (reg Register) MarshalText() ([]byte, error) {
	return []byte(reg.String()), nil
}

// parseReg returns the x86 register corresponding to the given string.
func parseReg(s string) x86asm.Reg {
	m := map[string]x86asm.Reg{
		// 8-bit
		"AL":   x86asm.AL,
		"CL":   x86asm.CL,
		"DL":   x86asm.DL,
		"BL":   x86asm.BL,
		"AH":   x86asm.AH,
		"CH":   x86asm.CH,
		"DH":   x86asm.DH,
		"BH":   x86asm.BH,
		"SPB":  x86asm.SPB,
		"BPB":  x86asm.BPB,
		"SIB":  x86asm.SIB,
		"DIB":  x86asm.DIB,
		"R8B":  x86asm.R8B,
		"R9B":  x86asm.R9B,
		"R10B": x86asm.R10B,
		"R11B": x86asm.R11B,
		"R12B": x86asm.R12B,
		"R13B": x86asm.R13B,
		"R14B": x86asm.R14B,
		"R15B": x86asm.R15B,
		// 16-bit
		"AX":   x86asm.AX,
		"CX":   x86asm.CX,
		"DX":   x86asm.DX,
		"BX":   x86asm.BX,
		"SP":   x86asm.SP,
		"BP":   x86asm.BP,
		"SI":   x86asm.SI,
		"DI":   x86asm.DI,
		"R8W":  x86asm.R8W,
		"R9W":  x86asm.R9W,
		"R10W": x86asm.R10W,
		"R11W": x86asm.R11W,
		"R12W": x86asm.R12W,
		"R13W": x86asm.R13W,
		"R14W": x86asm.R14W,
		"R15W": x86asm.R15W,
		// 32-bit
		"EAX":  x86asm.EAX,
		"ECX":  x86asm.ECX,
		"EDX":  x86asm.EDX,
		"EBX":  x86asm.EBX,
		"ESP":  x86asm.ESP,
		"EBP":  x86asm.EBP,
		"ESI":  x86asm.ESI,
		"EDI":  x86asm.EDI,
		"R8L":  x86asm.R8L,
		"R9L":  x86asm.R9L,
		"R10L": x86asm.R10L,
		"R11L": x86asm.R11L,
		"R12L": x86asm.R12L,
		"R13L": x86asm.R13L,
		"R14L": x86asm.R14L,
		"R15L": x86asm.R15L,
		// 64-bit
		"RAX": x86asm.RAX,
		"RCX": x86asm.RCX,
		"RDX": x86asm.RDX,
		"RBX": x86asm.RBX,
		"RSP": x86asm.RSP,
		"RBP": x86asm.RBP,
		"RSI": x86asm.RSI,
		"RDI": x86asm.RDI,
		"R8":  x86asm.R8,
		"R9":  x86asm.R9,
		"R10": x86asm.R10,
		"R11": x86asm.R11,
		"R12": x86asm.R12,
		"R13": x86asm.R13,
		"R14": x86asm.R14,
		"R15": x86asm.R15,
		// Instruction pointers.
		"IP":  x86asm.IP,
		"EIP": x86asm.EIP,
		"RIP": x86asm.RIP,
		// 387 floating point registers.
		"F0": x86asm.F0,
		"F1": x86asm.F1,
		"F2": x86asm.F2,
		"F3": x86asm.F3,
		"F4": x86asm.F4,
		"F5": x86asm.F5,
		"F6": x86asm.F6,
		"F7": x86asm.F7,
		// MMX registers.
		"M0": x86asm.M0,
		"M1": x86asm.M1,
		"M2": x86asm.M2,
		"M3": x86asm.M3,
		"M4": x86asm.M4,
		"M5": x86asm.M5,
		"M6": x86asm.M6,
		"M7": x86asm.M7,
		// XMM registers.
		"X0":  x86asm.X0,
		"X1":  x86asm.X1,
		"X2":  x86asm.X2,
		"X3":  x86asm.X3,
		"X4":  x86asm.X4,
		"X5":  x86asm.X5,
		"X6":  x86asm.X6,
		"X7":  x86asm.X7,
		"X8":  x86asm.X8,
		"X9":  x86asm.X9,
		"X10": x86asm.X10,
		"X11": x86asm.X11,
		"X12": x86asm.X12,
		"X13": x86asm.X13,
		"X14": x86asm.X14,
		"X15": x86asm.X15,
		// Segment registers.
		"ES": x86asm.ES,
		"CS": x86asm.CS,
		"SS": x86asm.SS,
		"DS": x86asm.DS,
		"FS": x86asm.FS,
		"GS": x86asm.GS,
		// System registers.
		"GDTR": x86asm.GDTR,
		"IDTR": x86asm.IDTR,
		"LDTR": x86asm.LDTR,
		"MSW":  x86asm.MSW,
		"TASK": x86asm.TASK,
		// Control registers.
		"CR0":  x86asm.CR0,
		"CR1":  x86asm.CR1,
		"CR2":  x86asm.CR2,
		"CR3":  x86asm.CR3,
		"CR4":  x86asm.CR4,
		"CR5":  x86asm.CR5,
		"CR6":  x86asm.CR6,
		"CR7":  x86asm.CR7,
		"CR8":  x86asm.CR8,
		"CR9":  x86asm.CR9,
		"CR10": x86asm.CR10,
		"CR11": x86asm.CR11,
		"CR12": x86asm.CR12,
		"CR13": x86asm.CR13,
		"CR14": x86asm.CR14,
		"CR15": x86asm.CR15,
		// Debug registers.
		"DR0":  x86asm.DR0,
		"DR1":  x86asm.DR1,
		"DR2":  x86asm.DR2,
		"DR3":  x86asm.DR3,
		"DR4":  x86asm.DR4,
		"DR5":  x86asm.DR5,
		"DR6":  x86asm.DR6,
		"DR7":  x86asm.DR7,
		"DR8":  x86asm.DR8,
		"DR9":  x86asm.DR9,
		"DR10": x86asm.DR10,
		"DR11": x86asm.DR11,
		"DR12": x86asm.DR12,
		"DR13": x86asm.DR13,
		"DR14": x86asm.DR14,
		"DR15": x86asm.DR15,
		// Task registers.
		"TR0": x86asm.TR0,
		"TR1": x86asm.TR1,
		"TR2": x86asm.TR2,
		"TR3": x86asm.TR3,
		"TR4": x86asm.TR4,
		"TR5": x86asm.TR5,
		"TR6": x86asm.TR6,
		"TR7": x86asm.TR7,
		// PSEUDO-registers.
		"EDX:EAX": x86asm_EDX_EAX,
	}
	if reg, ok := m[s]; ok {
		return reg
	}
	panic(fmt.Errorf("support for register %q not yet implemented", s))
}

// PSEUDO-registers.
const (
	firstReg = x86asm.AL
	// CL
	// ...
	// TR7

	// EDX:EAX (used in idiv)
	x86asm_EDX_EAX = x86asm.TR7 + 1

	lastReg = x86asm_EDX_EAX
)
