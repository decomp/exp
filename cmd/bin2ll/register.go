package main

import (
	"fmt"

	"github.com/llir/llvm/ir/types"

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
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
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
	case x86asm_EDX_EAX:
		return types.I64
	default:
		panic(fmt.Errorf("support for register %v not yet implemented", reg))
	}
}

// Registers.
var (
	// 8-bit
	AL   = NewReg(x86asm.AL, nil)
	CL   = NewReg(x86asm.CL, nil)
	DL   = NewReg(x86asm.DL, nil)
	BL   = NewReg(x86asm.BL, nil)
	AH   = NewReg(x86asm.AH, nil)
	CH   = NewReg(x86asm.CH, nil)
	DH   = NewReg(x86asm.DH, nil)
	BH   = NewReg(x86asm.BH, nil)
	SPB  = NewReg(x86asm.SPB, nil)
	BPB  = NewReg(x86asm.BPB, nil)
	SIB  = NewReg(x86asm.SIB, nil)
	DIB  = NewReg(x86asm.DIB, nil)
	R8B  = NewReg(x86asm.R8B, nil)
	R9B  = NewReg(x86asm.R9B, nil)
	R10B = NewReg(x86asm.R10B, nil)
	R11B = NewReg(x86asm.R11B, nil)
	R12B = NewReg(x86asm.R12B, nil)
	R13B = NewReg(x86asm.R13B, nil)
	R14B = NewReg(x86asm.R14B, nil)
	R15B = NewReg(x86asm.R15B, nil)
	// 16-bit
	AX   = NewReg(x86asm.AX, nil)
	CX   = NewReg(x86asm.CX, nil)
	DX   = NewReg(x86asm.DX, nil)
	BX   = NewReg(x86asm.BX, nil)
	SP   = NewReg(x86asm.SP, nil)
	BP   = NewReg(x86asm.BP, nil)
	SI   = NewReg(x86asm.SI, nil)
	DI   = NewReg(x86asm.DI, nil)
	R8W  = NewReg(x86asm.R8W, nil)
	R9W  = NewReg(x86asm.R9W, nil)
	R10W = NewReg(x86asm.R10W, nil)
	R11W = NewReg(x86asm.R11W, nil)
	R12W = NewReg(x86asm.R12W, nil)
	R13W = NewReg(x86asm.R13W, nil)
	R14W = NewReg(x86asm.R14W, nil)
	R15W = NewReg(x86asm.R15W, nil)
	// 32-bit
	EAX  = NewReg(x86asm.EAX, nil)
	ECX  = NewReg(x86asm.ECX, nil)
	EDX  = NewReg(x86asm.EDX, nil)
	EBX  = NewReg(x86asm.EBX, nil)
	ESP  = NewReg(x86asm.ESP, nil)
	EBP  = NewReg(x86asm.EBP, nil)
	ESI  = NewReg(x86asm.ESI, nil)
	EDI  = NewReg(x86asm.EDI, nil)
	R8L  = NewReg(x86asm.R8L, nil)
	R9L  = NewReg(x86asm.R9L, nil)
	R10L = NewReg(x86asm.R10L, nil)
	R11L = NewReg(x86asm.R11L, nil)
	R12L = NewReg(x86asm.R12L, nil)
	R13L = NewReg(x86asm.R13L, nil)
	R14L = NewReg(x86asm.R14L, nil)
	R15L = NewReg(x86asm.R15L, nil)
	// 64-bit
	RAX = NewReg(x86asm.RAX, nil)
	RCX = NewReg(x86asm.RCX, nil)
	RDX = NewReg(x86asm.RDX, nil)
	RBX = NewReg(x86asm.RBX, nil)
	RSP = NewReg(x86asm.RSP, nil)
	RBP = NewReg(x86asm.RBP, nil)
	RSI = NewReg(x86asm.RSI, nil)
	RDI = NewReg(x86asm.RDI, nil)
	R8  = NewReg(x86asm.R8, nil)
	R9  = NewReg(x86asm.R9, nil)
	R10 = NewReg(x86asm.R10, nil)
	R11 = NewReg(x86asm.R11, nil)
	R12 = NewReg(x86asm.R12, nil)
	R13 = NewReg(x86asm.R13, nil)
	R14 = NewReg(x86asm.R14, nil)
	R15 = NewReg(x86asm.R15, nil)
	// Instruction pointers.
	IP  = NewReg(x86asm.IP, nil)
	EIP = NewReg(x86asm.EIP, nil)
	RIP = NewReg(x86asm.RIP, nil)
	// 387 floating point registers.
	F0 = NewReg(x86asm.F0, nil)
	F1 = NewReg(x86asm.F1, nil)
	F2 = NewReg(x86asm.F2, nil)
	F3 = NewReg(x86asm.F3, nil)
	F4 = NewReg(x86asm.F4, nil)
	F5 = NewReg(x86asm.F5, nil)
	F6 = NewReg(x86asm.F6, nil)
	F7 = NewReg(x86asm.F7, nil)
	// MMX registers.
	M0 = NewReg(x86asm.M0, nil)
	M1 = NewReg(x86asm.M1, nil)
	M2 = NewReg(x86asm.M2, nil)
	M3 = NewReg(x86asm.M3, nil)
	M4 = NewReg(x86asm.M4, nil)
	M5 = NewReg(x86asm.M5, nil)
	M6 = NewReg(x86asm.M6, nil)
	M7 = NewReg(x86asm.M7, nil)
	// XMM registers.
	X0  = NewReg(x86asm.X0, nil)
	X1  = NewReg(x86asm.X1, nil)
	X2  = NewReg(x86asm.X2, nil)
	X3  = NewReg(x86asm.X3, nil)
	X4  = NewReg(x86asm.X4, nil)
	X5  = NewReg(x86asm.X5, nil)
	X6  = NewReg(x86asm.X6, nil)
	X7  = NewReg(x86asm.X7, nil)
	X8  = NewReg(x86asm.X8, nil)
	X9  = NewReg(x86asm.X9, nil)
	X10 = NewReg(x86asm.X10, nil)
	X11 = NewReg(x86asm.X11, nil)
	X12 = NewReg(x86asm.X12, nil)
	X13 = NewReg(x86asm.X13, nil)
	X14 = NewReg(x86asm.X14, nil)
	X15 = NewReg(x86asm.X15, nil)
	// Segment registers.
	ES = NewReg(x86asm.ES, nil)
	CS = NewReg(x86asm.CS, nil)
	SS = NewReg(x86asm.SS, nil)
	DS = NewReg(x86asm.DS, nil)
	FS = NewReg(x86asm.FS, nil)
	GS = NewReg(x86asm.GS, nil)
	// System registers.
	GDTR = NewReg(x86asm.GDTR, nil)
	IDTR = NewReg(x86asm.IDTR, nil)
	LDTR = NewReg(x86asm.LDTR, nil)
	MSW  = NewReg(x86asm.MSW, nil)
	TASK = NewReg(x86asm.TASK, nil)
	// Control registers.
	CR0  = NewReg(x86asm.CR0, nil)
	CR1  = NewReg(x86asm.CR1, nil)
	CR2  = NewReg(x86asm.CR2, nil)
	CR3  = NewReg(x86asm.CR3, nil)
	CR4  = NewReg(x86asm.CR4, nil)
	CR5  = NewReg(x86asm.CR5, nil)
	CR6  = NewReg(x86asm.CR6, nil)
	CR7  = NewReg(x86asm.CR7, nil)
	CR8  = NewReg(x86asm.CR8, nil)
	CR9  = NewReg(x86asm.CR9, nil)
	CR10 = NewReg(x86asm.CR10, nil)
	CR11 = NewReg(x86asm.CR11, nil)
	CR12 = NewReg(x86asm.CR12, nil)
	CR13 = NewReg(x86asm.CR13, nil)
	CR14 = NewReg(x86asm.CR14, nil)
	CR15 = NewReg(x86asm.CR15, nil)
	// Debug registers.
	DR0  = NewReg(x86asm.DR0, nil)
	DR1  = NewReg(x86asm.DR1, nil)
	DR2  = NewReg(x86asm.DR2, nil)
	DR3  = NewReg(x86asm.DR3, nil)
	DR4  = NewReg(x86asm.DR4, nil)
	DR5  = NewReg(x86asm.DR5, nil)
	DR6  = NewReg(x86asm.DR6, nil)
	DR7  = NewReg(x86asm.DR7, nil)
	DR8  = NewReg(x86asm.DR8, nil)
	DR9  = NewReg(x86asm.DR9, nil)
	DR10 = NewReg(x86asm.DR10, nil)
	DR11 = NewReg(x86asm.DR11, nil)
	DR12 = NewReg(x86asm.DR12, nil)
	DR13 = NewReg(x86asm.DR13, nil)
	DR14 = NewReg(x86asm.DR14, nil)
	DR15 = NewReg(x86asm.DR15, nil)
	// Task registers.
	TR0 = NewReg(x86asm.TR0, nil)
	TR1 = NewReg(x86asm.TR1, nil)
	TR2 = NewReg(x86asm.TR2, nil)
	TR3 = NewReg(x86asm.TR3, nil)
	TR4 = NewReg(x86asm.TR4, nil)
	TR5 = NewReg(x86asm.TR5, nil)
	TR6 = NewReg(x86asm.TR6, nil)
	TR7 = NewReg(x86asm.TR7, nil)
	// PSEUDO-registers.
	EDX_EAX = NewReg(x86asm_EDX_EAX, nil)
)

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
