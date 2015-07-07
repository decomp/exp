package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"strconv"

	"github.com/mewkiz/pkg/errutil"
	"rsc.io/x86/x86asm"
)

// getArg converts arg into a corresponding Go expression.
func getArg(arg x86asm.Arg) ast.Expr {
	switch arg := arg.(type) {
	case x86asm.Reg:
		return getReg(arg)
	case x86asm.Mem:
		return getMem(arg)
	case x86asm.Imm:
		return createExpr(int64(arg))
	case x86asm.Rel:
		// TODO: Implement support for relative addresses.
	}
	fmt.Printf("%#v\n", arg)
	log.Fatal(errutil.Newf("support for type %T not yet implemented", arg))
	panic("unreachable")
}

// regs maps register names to their corresponding Go identifiers.
var regs = map[string]*ast.Ident{
	// 8-bit
	"AL":   ast.NewIdent("al"),
	"CL":   ast.NewIdent("cl"),
	"DL":   ast.NewIdent("dl"),
	"BL":   ast.NewIdent("bl"),
	"AH":   ast.NewIdent("ah"),
	"CH":   ast.NewIdent("ch"),
	"DH":   ast.NewIdent("dh"),
	"BH":   ast.NewIdent("bh"),
	"SPB":  ast.NewIdent("spb"),
	"BPB":  ast.NewIdent("bpb"),
	"SIB":  ast.NewIdent("sib"),
	"DIB":  ast.NewIdent("dib"),
	"R8B":  ast.NewIdent("r8b"),
	"R9B":  ast.NewIdent("r9b"),
	"R10B": ast.NewIdent("r10b"),
	"R11B": ast.NewIdent("r11b"),
	"R12B": ast.NewIdent("r12b"),
	"R13B": ast.NewIdent("r13b"),
	"R14B": ast.NewIdent("r14b"),
	"R15B": ast.NewIdent("r15b"),

	// 16-bit
	"AX":   ast.NewIdent("ax"),
	"CX":   ast.NewIdent("cx"),
	"DX":   ast.NewIdent("dx"),
	"BX":   ast.NewIdent("bx"),
	"SP":   ast.NewIdent("sp"),
	"BP":   ast.NewIdent("bp"),
	"SI":   ast.NewIdent("si"),
	"DI":   ast.NewIdent("di"),
	"R8W":  ast.NewIdent("r8w"),
	"R9W":  ast.NewIdent("r9w"),
	"R10W": ast.NewIdent("r10w"),
	"R11W": ast.NewIdent("r11w"),
	"R12W": ast.NewIdent("r12w"),
	"R13W": ast.NewIdent("r13w"),
	"R14W": ast.NewIdent("r14w"),
	"R15W": ast.NewIdent("r15w"),

	// 32-bit
	"EAX":  ast.NewIdent("eax"),
	"ECX":  ast.NewIdent("ecx"),
	"EDX":  ast.NewIdent("edx"),
	"EBX":  ast.NewIdent("ebx"),
	"ESP":  ast.NewIdent("esp"),
	"EBP":  ast.NewIdent("ebp"),
	"ESI":  ast.NewIdent("esi"),
	"EDI":  ast.NewIdent("edi"),
	"R8L":  ast.NewIdent("r8l"),
	"R9L":  ast.NewIdent("r9l"),
	"R10L": ast.NewIdent("r10l"),
	"R11L": ast.NewIdent("r11l"),
	"R12L": ast.NewIdent("r12l"),
	"R13L": ast.NewIdent("r13l"),
	"R14L": ast.NewIdent("r14l"),
	"R15L": ast.NewIdent("r15l"),

	// 64-bit
	"RAX": ast.NewIdent("rax"),
	"RCX": ast.NewIdent("rcx"),
	"RDX": ast.NewIdent("rdx"),
	"RBX": ast.NewIdent("rbx"),
	"RSP": ast.NewIdent("rsp"),
	"RBP": ast.NewIdent("rbp"),
	"RSI": ast.NewIdent("rsi"),
	"RDI": ast.NewIdent("rdi"),
	"R8":  ast.NewIdent("r8"),
	"R9":  ast.NewIdent("r9"),
	"R10": ast.NewIdent("r10"),
	"R11": ast.NewIdent("r11"),
	"R12": ast.NewIdent("r12"),
	"R13": ast.NewIdent("r13"),
	"R14": ast.NewIdent("r14"),
	"R15": ast.NewIdent("r15"),

	// Instruction pointer.
	"IP":  ast.NewIdent("ip"),  // 16-bit
	"EIP": ast.NewIdent("eip"), // 32-bit
	"RIP": ast.NewIdent("rip"), // 64-bit

	// 387 floating point registers.
	"F0": ast.NewIdent("f0"),
	"F1": ast.NewIdent("f1"),
	"F2": ast.NewIdent("f2"),
	"F3": ast.NewIdent("f3"),
	"F4": ast.NewIdent("f4"),
	"F5": ast.NewIdent("f5"),
	"F6": ast.NewIdent("f6"),
	"F7": ast.NewIdent("f7"),

	// MMX registers.
	"M0": ast.NewIdent("m0"),
	"M1": ast.NewIdent("m1"),
	"M2": ast.NewIdent("m2"),
	"M3": ast.NewIdent("m3"),
	"M4": ast.NewIdent("m4"),
	"M5": ast.NewIdent("m5"),
	"M6": ast.NewIdent("m6"),
	"M7": ast.NewIdent("m7"),

	// XMM registers.
	"X0":  ast.NewIdent("x0"),
	"X1":  ast.NewIdent("x1"),
	"X2":  ast.NewIdent("x2"),
	"X3":  ast.NewIdent("x3"),
	"X4":  ast.NewIdent("x4"),
	"X5":  ast.NewIdent("x5"),
	"X6":  ast.NewIdent("x6"),
	"X7":  ast.NewIdent("x7"),
	"X8":  ast.NewIdent("x8"),
	"X9":  ast.NewIdent("x9"),
	"X10": ast.NewIdent("x10"),
	"X11": ast.NewIdent("x11"),
	"X12": ast.NewIdent("x12"),
	"X13": ast.NewIdent("x13"),
	"X14": ast.NewIdent("x14"),
	"X15": ast.NewIdent("x15"),

	// Segment registers.
	"ES": ast.NewIdent("es"),
	"CS": ast.NewIdent("cs"),
	"SS": ast.NewIdent("ss"),
	"DS": ast.NewIdent("ds"),
	"FS": ast.NewIdent("fs"),
	"GS": ast.NewIdent("gs"),

	// System registers.
	"GDTR": ast.NewIdent("gdtr"),
	"IDTR": ast.NewIdent("idtr"),
	"LDTR": ast.NewIdent("ldtr"),
	"MSW":  ast.NewIdent("msw"),
	"TASK": ast.NewIdent("task"),

	// Control registers.
	"CR0":  ast.NewIdent("cr0"),
	"CR1":  ast.NewIdent("cr1"),
	"CR2":  ast.NewIdent("cr2"),
	"CR3":  ast.NewIdent("cr3"),
	"CR4":  ast.NewIdent("cr4"),
	"CR5":  ast.NewIdent("cr5"),
	"CR6":  ast.NewIdent("cr6"),
	"CR7":  ast.NewIdent("cr7"),
	"CR8":  ast.NewIdent("cr8"),
	"CR9":  ast.NewIdent("cr9"),
	"CR10": ast.NewIdent("cr10"),
	"CR11": ast.NewIdent("cr11"),
	"CR12": ast.NewIdent("cr12"),
	"CR13": ast.NewIdent("cr13"),
	"CR14": ast.NewIdent("cr14"),
	"CR15": ast.NewIdent("cr15"),

	// Debug registers.
	"DR0":  ast.NewIdent("dr0"),
	"DR1":  ast.NewIdent("dr1"),
	"DR2":  ast.NewIdent("dr2"),
	"DR3":  ast.NewIdent("dr3"),
	"DR4":  ast.NewIdent("dr4"),
	"DR5":  ast.NewIdent("dr5"),
	"DR6":  ast.NewIdent("dr6"),
	"DR7":  ast.NewIdent("dr7"),
	"DR8":  ast.NewIdent("dr8"),
	"DR9":  ast.NewIdent("dr9"),
	"DR10": ast.NewIdent("dr10"),
	"DR11": ast.NewIdent("dr11"),
	"DR12": ast.NewIdent("dr12"),
	"DR13": ast.NewIdent("dr13"),
	"DR14": ast.NewIdent("dr14"),
	"DR15": ast.NewIdent("dr15"),

	// Task registers.
	"TR0": ast.NewIdent("tr0"),
	"TR1": ast.NewIdent("tr1"),
	"TR2": ast.NewIdent("tr2"),
	"TR3": ast.NewIdent("tr3"),
	"TR4": ast.NewIdent("tr4"),
	"TR5": ast.NewIdent("tr5"),
	"TR6": ast.NewIdent("tr6"),
	"TR7": ast.NewIdent("tr7"),
}

// getReg converts reg into a corresponding Go expression.
func getReg(reg x86asm.Reg) ast.Expr {
	if expr, ok := regs[reg.String()]; ok {
		return expr
	}
	log.Fatal(errutil.Newf("unable to lookup identifer for register %q", reg))
	panic("unreachable")
}

// getMem converts mem into a corresponding Go expression.
func getMem(mem x86asm.Mem) ast.Expr {
	// The general memory reference form is:
	//    Segment:[Base+Scale*Index+Disp]

	// ... + Disp
	expr := &ast.BinaryExpr{}
	if mem.Disp != 0 {
		disp := createExpr(mem.Disp)
		expr.Op = token.ADD
		expr.Y = disp
	}

	// ... + (Scale*Index) + ...
	if mem.Scale != 0 && mem.Index != 0 {
		scale := createExpr(mem.Scale)
		index := getReg(mem.Index)
		product := &ast.BinaryExpr{
			X:  scale,
			Op: token.MUL,
			Y:  index,
		}
		switch {
		case expr.Y == nil:
			// ... + (Scale*Index)
			expr.Op = token.ADD
			expr.Y = product
		default:
			// ... + (Scale*Index) + Disp
			expr.X = product
			expr.Op = token.ADD
		}
	}

	// ... + Base + ...
	if mem.Base != 0 {
		base := getReg(mem.Base)
		switch {
		case expr.X == nil:
			// Base + (Scale*Index)
			// or
			// Base + Disp
			expr.X = base
			expr.Op = token.ADD
		case expr.Y == nil:
			// ... + Base
			expr.Op = token.ADD
			expr.Y = base
		default:
			sum := &ast.BinaryExpr{
				X:  expr.X,
				Op: token.ADD,
				Y:  expr.Y,
			}
			expr.X = base
			expr.Op = token.ADD
			expr.Y = sum
		}
	}

	// TODO: Figure out how the calculation is affected by segment in:
	//    Segment:[Base+Scale*Index+Disp]
	if mem.Segment != 0 {
		segment := getReg(mem.Segment)
		_ = segment
		fmt.Printf("%#v\n", mem)
		log.Fatal(errutil.Newf("support for Mem.Segment not yet implemented"))
	}

	switch {
	case expr.X == nil && expr.Y == nil:
		fmt.Printf("%#v\n", mem)
		log.Fatal(errutil.New("support for memory reference to address zero not yet implemented"))
		panic("unreachable")
	case expr.X == nil && expr.Y != nil:
		return createPtrDeref(expr.Y)
	case expr.X != nil && expr.Y == nil:
		return createPtrDeref(expr.X)
	default:
		return createPtrDeref(expr)
	}
}

// createPtrDeref returns a pointer dereference expression of addr.
func createPtrDeref(addr ast.Expr) ast.Expr {
	return &ast.StarExpr{X: &ast.ParenExpr{X: addr}}
}

// createExpr converts x into a corresponding Go expression.
func createExpr(x interface{}) ast.Expr {
	switch x := x.(type) {
	case int:
		s := strconv.FormatInt(int64(x), 10)
		return &ast.BasicLit{Kind: token.INT, Value: s}
	case int64:
		s := strconv.FormatInt(x, 10)
		return &ast.BasicLit{Kind: token.INT, Value: s}
	case uint8:
		s := strconv.FormatUint(uint64(x), 10)
		return &ast.BasicLit{Kind: token.INT, Value: s}
	}
	log.Fatal(errutil.Newf("support for type %T not yet implemented", x))
	panic("unreachable")
}
