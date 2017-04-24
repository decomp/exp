package main

import (
	"fmt"

	"github.com/decomp/exp/bin"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"golang.org/x/arch/x86/x86asm"
)

// An Inst is a single x86 instruction.
type Inst struct {
	// x86 instruction.
	x86asm.Inst
	// Address of the instruction.
	addr bin.Address
}

// emitInst translates the given x86 instruction to LLVM IR, emitting code to f.
func (f *Func) emitInst(inst *Inst) error {
	dbg.Println("lifting instruction:", inst.Inst)

	// Check if prefix is present.
	for _, prefix := range inst.Prefix[:] {
		// The first zero in the array marks the end of the prefixes.
		if prefix == 0 {
			break
		}
		switch prefix {
		case x86asm.PrefixData16, x86asm.PrefixData16 | x86asm.PrefixImplicit:
			// prefix already supported.
		default:
			pretty.Println("instruction with prefix:", inst)
			panic(fmt.Errorf("support for %v instruction with prefix not yet implemented", inst.Op))
		}
	}

	// Translate instruction.
	switch inst.Op {
	case x86asm.AAA:
		return f.emitInstAAA(inst)
	case x86asm.AAD:
		return f.emitInstAAD(inst)
	case x86asm.AAM:
		return f.emitInstAAM(inst)
	case x86asm.AAS:
		return f.emitInstAAS(inst)
	case x86asm.ADC:
		return f.emitInstADC(inst)
	case x86asm.ADD:
		return f.emitInstADD(inst)
	case x86asm.ADDPD:
		return f.emitInstADDPD(inst)
	case x86asm.ADDPS:
		return f.emitInstADDPS(inst)
	case x86asm.ADDSD:
		return f.emitInstADDSD(inst)
	case x86asm.ADDSS:
		return f.emitInstADDSS(inst)
	case x86asm.ADDSUBPD:
		return f.emitInstADDSUBPD(inst)
	case x86asm.ADDSUBPS:
		return f.emitInstADDSUBPS(inst)
	case x86asm.AESDEC:
		return f.emitInstAESDEC(inst)
	case x86asm.AESDECLAST:
		return f.emitInstAESDECLAST(inst)
	case x86asm.AESENC:
		return f.emitInstAESENC(inst)
	case x86asm.AESENCLAST:
		return f.emitInstAESENCLAST(inst)
	case x86asm.AESIMC:
		return f.emitInstAESIMC(inst)
	case x86asm.AESKEYGENASSIST:
		return f.emitInstAESKEYGENASSIST(inst)
	case x86asm.AND:
		return f.emitInstAND(inst)
	case x86asm.ANDNPD:
		return f.emitInstANDNPD(inst)
	case x86asm.ANDNPS:
		return f.emitInstANDNPS(inst)
	case x86asm.ANDPD:
		return f.emitInstANDPD(inst)
	case x86asm.ANDPS:
		return f.emitInstANDPS(inst)
	case x86asm.ARPL:
		return f.emitInstARPL(inst)
	case x86asm.BLENDPD:
		return f.emitInstBLENDPD(inst)
	case x86asm.BLENDPS:
		return f.emitInstBLENDPS(inst)
	case x86asm.BLENDVPD:
		return f.emitInstBLENDVPD(inst)
	case x86asm.BLENDVPS:
		return f.emitInstBLENDVPS(inst)
	case x86asm.BOUND:
		return f.emitInstBOUND(inst)
	case x86asm.BSF:
		return f.emitInstBSF(inst)
	case x86asm.BSR:
		return f.emitInstBSR(inst)
	case x86asm.BSWAP:
		return f.emitInstBSWAP(inst)
	case x86asm.BT:
		return f.emitInstBT(inst)
	case x86asm.BTC:
		return f.emitInstBTC(inst)
	case x86asm.BTR:
		return f.emitInstBTR(inst)
	case x86asm.BTS:
		return f.emitInstBTS(inst)
	case x86asm.CALL:
		return f.emitInstCALL(inst)
	case x86asm.CBW:
		return f.emitInstCBW(inst)
	case x86asm.CDQ:
		return f.emitInstCDQ(inst)
	case x86asm.CDQE:
		return f.emitInstCDQE(inst)
	case x86asm.CLC:
		return f.emitInstCLC(inst)
	case x86asm.CLD:
		return f.emitInstCLD(inst)
	case x86asm.CLFLUSH:
		return f.emitInstCLFLUSH(inst)
	case x86asm.CLI:
		return f.emitInstCLI(inst)
	case x86asm.CLTS:
		return f.emitInstCLTS(inst)
	case x86asm.CMC:
		return f.emitInstCMC(inst)
	case x86asm.CMOVA:
		return f.emitInstCMOVA(inst)
	case x86asm.CMOVAE:
		return f.emitInstCMOVAE(inst)
	case x86asm.CMOVB:
		return f.emitInstCMOVB(inst)
	case x86asm.CMOVBE:
		return f.emitInstCMOVBE(inst)
	case x86asm.CMOVE:
		return f.emitInstCMOVE(inst)
	case x86asm.CMOVG:
		return f.emitInstCMOVG(inst)
	case x86asm.CMOVGE:
		return f.emitInstCMOVGE(inst)
	case x86asm.CMOVL:
		return f.emitInstCMOVL(inst)
	case x86asm.CMOVLE:
		return f.emitInstCMOVLE(inst)
	case x86asm.CMOVNE:
		return f.emitInstCMOVNE(inst)
	case x86asm.CMOVNO:
		return f.emitInstCMOVNO(inst)
	case x86asm.CMOVNP:
		return f.emitInstCMOVNP(inst)
	case x86asm.CMOVNS:
		return f.emitInstCMOVNS(inst)
	case x86asm.CMOVO:
		return f.emitInstCMOVO(inst)
	case x86asm.CMOVP:
		return f.emitInstCMOVP(inst)
	case x86asm.CMOVS:
		return f.emitInstCMOVS(inst)
	case x86asm.CMP:
		return f.emitInstCMP(inst)
	case x86asm.CMPPD:
		return f.emitInstCMPPD(inst)
	case x86asm.CMPPS:
		return f.emitInstCMPPS(inst)
	case x86asm.CMPSB:
		return f.emitInstCMPSB(inst)
	case x86asm.CMPSD:
		return f.emitInstCMPSD(inst)
	case x86asm.CMPSD_XMM:
		return f.emitInstCMPSD_XMM(inst)
	case x86asm.CMPSQ:
		return f.emitInstCMPSQ(inst)
	case x86asm.CMPSS:
		return f.emitInstCMPSS(inst)
	case x86asm.CMPSW:
		return f.emitInstCMPSW(inst)
	case x86asm.CMPXCHG:
		return f.emitInstCMPXCHG(inst)
	case x86asm.CMPXCHG16B:
		return f.emitInstCMPXCHG16B(inst)
	case x86asm.CMPXCHG8B:
		return f.emitInstCMPXCHG8B(inst)
	case x86asm.COMISD:
		return f.emitInstCOMISD(inst)
	case x86asm.COMISS:
		return f.emitInstCOMISS(inst)
	case x86asm.CPUID:
		return f.emitInstCPUID(inst)
	case x86asm.CQO:
		return f.emitInstCQO(inst)
	case x86asm.CRC32:
		return f.emitInstCRC32(inst)
	case x86asm.CVTDQ2PD:
		return f.emitInstCVTDQ2PD(inst)
	case x86asm.CVTDQ2PS:
		return f.emitInstCVTDQ2PS(inst)
	case x86asm.CVTPD2DQ:
		return f.emitInstCVTPD2DQ(inst)
	case x86asm.CVTPD2PI:
		return f.emitInstCVTPD2PI(inst)
	case x86asm.CVTPD2PS:
		return f.emitInstCVTPD2PS(inst)
	case x86asm.CVTPI2PD:
		return f.emitInstCVTPI2PD(inst)
	case x86asm.CVTPI2PS:
		return f.emitInstCVTPI2PS(inst)
	case x86asm.CVTPS2DQ:
		return f.emitInstCVTPS2DQ(inst)
	case x86asm.CVTPS2PD:
		return f.emitInstCVTPS2PD(inst)
	case x86asm.CVTPS2PI:
		return f.emitInstCVTPS2PI(inst)
	case x86asm.CVTSD2SI:
		return f.emitInstCVTSD2SI(inst)
	case x86asm.CVTSD2SS:
		return f.emitInstCVTSD2SS(inst)
	case x86asm.CVTSI2SD:
		return f.emitInstCVTSI2SD(inst)
	case x86asm.CVTSI2SS:
		return f.emitInstCVTSI2SS(inst)
	case x86asm.CVTSS2SD:
		return f.emitInstCVTSS2SD(inst)
	case x86asm.CVTSS2SI:
		return f.emitInstCVTSS2SI(inst)
	case x86asm.CVTTPD2DQ:
		return f.emitInstCVTTPD2DQ(inst)
	case x86asm.CVTTPD2PI:
		return f.emitInstCVTTPD2PI(inst)
	case x86asm.CVTTPS2DQ:
		return f.emitInstCVTTPS2DQ(inst)
	case x86asm.CVTTPS2PI:
		return f.emitInstCVTTPS2PI(inst)
	case x86asm.CVTTSD2SI:
		return f.emitInstCVTTSD2SI(inst)
	case x86asm.CVTTSS2SI:
		return f.emitInstCVTTSS2SI(inst)
	case x86asm.CWD:
		return f.emitInstCWD(inst)
	case x86asm.CWDE:
		return f.emitInstCWDE(inst)
	case x86asm.DAA:
		return f.emitInstDAA(inst)
	case x86asm.DAS:
		return f.emitInstDAS(inst)
	case x86asm.DEC:
		return f.emitInstDEC(inst)
	case x86asm.DIV:
		return f.emitInstDIV(inst)
	case x86asm.DIVPD:
		return f.emitInstDIVPD(inst)
	case x86asm.DIVPS:
		return f.emitInstDIVPS(inst)
	case x86asm.DIVSD:
		return f.emitInstDIVSD(inst)
	case x86asm.DIVSS:
		return f.emitInstDIVSS(inst)
	case x86asm.DPPD:
		return f.emitInstDPPD(inst)
	case x86asm.DPPS:
		return f.emitInstDPPS(inst)
	case x86asm.EMMS:
		return f.emitInstEMMS(inst)
	case x86asm.ENTER:
		return f.emitInstENTER(inst)
	case x86asm.EXTRACTPS:
		return f.emitInstEXTRACTPS(inst)
	case x86asm.F2XM1:
		return f.emitInstF2XM1(inst)
	case x86asm.FABS:
		return f.emitInstFABS(inst)
	case x86asm.FADD:
		return f.emitInstFADD(inst)
	case x86asm.FADDP:
		return f.emitInstFADDP(inst)
	case x86asm.FBLD:
		return f.emitInstFBLD(inst)
	case x86asm.FBSTP:
		return f.emitInstFBSTP(inst)
	case x86asm.FCHS:
		return f.emitInstFCHS(inst)
	case x86asm.FCMOVB:
		return f.emitInstFCMOVB(inst)
	case x86asm.FCMOVBE:
		return f.emitInstFCMOVBE(inst)
	case x86asm.FCMOVE:
		return f.emitInstFCMOVE(inst)
	case x86asm.FCMOVNB:
		return f.emitInstFCMOVNB(inst)
	case x86asm.FCMOVNBE:
		return f.emitInstFCMOVNBE(inst)
	case x86asm.FCMOVNE:
		return f.emitInstFCMOVNE(inst)
	case x86asm.FCMOVNU:
		return f.emitInstFCMOVNU(inst)
	case x86asm.FCMOVU:
		return f.emitInstFCMOVU(inst)
	case x86asm.FCOM:
		return f.emitInstFCOM(inst)
	case x86asm.FCOMI:
		return f.emitInstFCOMI(inst)
	case x86asm.FCOMIP:
		return f.emitInstFCOMIP(inst)
	case x86asm.FCOMP:
		return f.emitInstFCOMP(inst)
	case x86asm.FCOMPP:
		return f.emitInstFCOMPP(inst)
	case x86asm.FCOS:
		return f.emitInstFCOS(inst)
	case x86asm.FDECSTP:
		return f.emitInstFDECSTP(inst)
	case x86asm.FDIV:
		return f.emitInstFDIV(inst)
	case x86asm.FDIVP:
		return f.emitInstFDIVP(inst)
	case x86asm.FDIVR:
		return f.emitInstFDIVR(inst)
	case x86asm.FDIVRP:
		return f.emitInstFDIVRP(inst)
	case x86asm.FFREE:
		return f.emitInstFFREE(inst)
	case x86asm.FFREEP:
		return f.emitInstFFREEP(inst)
	case x86asm.FIADD:
		return f.emitInstFIADD(inst)
	case x86asm.FICOM:
		return f.emitInstFICOM(inst)
	case x86asm.FICOMP:
		return f.emitInstFICOMP(inst)
	case x86asm.FIDIV:
		return f.emitInstFIDIV(inst)
	case x86asm.FIDIVR:
		return f.emitInstFIDIVR(inst)
	case x86asm.FILD:
		return f.emitInstFILD(inst)
	case x86asm.FIMUL:
		return f.emitInstFIMUL(inst)
	case x86asm.FINCSTP:
		return f.emitInstFINCSTP(inst)
	case x86asm.FIST:
		return f.emitInstFIST(inst)
	case x86asm.FISTP:
		return f.emitInstFISTP(inst)
	case x86asm.FISTTP:
		return f.emitInstFISTTP(inst)
	case x86asm.FISUB:
		return f.emitInstFISUB(inst)
	case x86asm.FISUBR:
		return f.emitInstFISUBR(inst)
	case x86asm.FLD:
		return f.emitInstFLD(inst)
	case x86asm.FLD1:
		return f.emitInstFLD1(inst)
	case x86asm.FLDCW:
		return f.emitInstFLDCW(inst)
	case x86asm.FLDENV:
		return f.emitInstFLDENV(inst)
	case x86asm.FLDL2E:
		return f.emitInstFLDL2E(inst)
	case x86asm.FLDL2T:
		return f.emitInstFLDL2T(inst)
	case x86asm.FLDLG2:
		return f.emitInstFLDLG2(inst)
	case x86asm.FLDLN2:
		return f.emitInstFLDLN2(inst)
	case x86asm.FLDPI:
		return f.emitInstFLDPI(inst)
	case x86asm.FLDZ:
		return f.emitInstFLDZ(inst)
	case x86asm.FMUL:
		return f.emitInstFMUL(inst)
	case x86asm.FMULP:
		return f.emitInstFMULP(inst)
	case x86asm.FNCLEX:
		return f.emitInstFNCLEX(inst)
	case x86asm.FNINIT:
		return f.emitInstFNINIT(inst)
	case x86asm.FNOP:
		return f.emitInstFNOP(inst)
	case x86asm.FNSAVE:
		return f.emitInstFNSAVE(inst)
	case x86asm.FNSTCW:
		return f.emitInstFNSTCW(inst)
	case x86asm.FNSTENV:
		return f.emitInstFNSTENV(inst)
	case x86asm.FNSTSW:
		return f.emitInstFNSTSW(inst)
	case x86asm.FPATAN:
		return f.emitInstFPATAN(inst)
	case x86asm.FPREM:
		return f.emitInstFPREM(inst)
	case x86asm.FPREM1:
		return f.emitInstFPREM1(inst)
	case x86asm.FPTAN:
		return f.emitInstFPTAN(inst)
	case x86asm.FRNDINT:
		return f.emitInstFRNDINT(inst)
	case x86asm.FRSTOR:
		return f.emitInstFRSTOR(inst)
	case x86asm.FSCALE:
		return f.emitInstFSCALE(inst)
	case x86asm.FSIN:
		return f.emitInstFSIN(inst)
	case x86asm.FSINCOS:
		return f.emitInstFSINCOS(inst)
	case x86asm.FSQRT:
		return f.emitInstFSQRT(inst)
	case x86asm.FST:
		return f.emitInstFST(inst)
	case x86asm.FSTP:
		return f.emitInstFSTP(inst)
	case x86asm.FSUB:
		return f.emitInstFSUB(inst)
	case x86asm.FSUBP:
		return f.emitInstFSUBP(inst)
	case x86asm.FSUBR:
		return f.emitInstFSUBR(inst)
	case x86asm.FSUBRP:
		return f.emitInstFSUBRP(inst)
	case x86asm.FTST:
		return f.emitInstFTST(inst)
	case x86asm.FUCOM:
		return f.emitInstFUCOM(inst)
	case x86asm.FUCOMI:
		return f.emitInstFUCOMI(inst)
	case x86asm.FUCOMIP:
		return f.emitInstFUCOMIP(inst)
	case x86asm.FUCOMP:
		return f.emitInstFUCOMP(inst)
	case x86asm.FUCOMPP:
		return f.emitInstFUCOMPP(inst)
	case x86asm.FWAIT:
		return f.emitInstFWAIT(inst)
	case x86asm.FXAM:
		return f.emitInstFXAM(inst)
	case x86asm.FXCH:
		return f.emitInstFXCH(inst)
	case x86asm.FXRSTOR:
		return f.emitInstFXRSTOR(inst)
	case x86asm.FXRSTOR64:
		return f.emitInstFXRSTOR64(inst)
	case x86asm.FXSAVE:
		return f.emitInstFXSAVE(inst)
	case x86asm.FXSAVE64:
		return f.emitInstFXSAVE64(inst)
	case x86asm.FXTRACT:
		return f.emitInstFXTRACT(inst)
	case x86asm.FYL2X:
		return f.emitInstFYL2X(inst)
	case x86asm.FYL2XP1:
		return f.emitInstFYL2XP1(inst)
	case x86asm.HADDPD:
		return f.emitInstHADDPD(inst)
	case x86asm.HADDPS:
		return f.emitInstHADDPS(inst)
	case x86asm.HLT:
		return f.emitInstHLT(inst)
	case x86asm.HSUBPD:
		return f.emitInstHSUBPD(inst)
	case x86asm.HSUBPS:
		return f.emitInstHSUBPS(inst)
	case x86asm.ICEBP:
		return f.emitInstICEBP(inst)
	case x86asm.IDIV:
		return f.emitInstIDIV(inst)
	case x86asm.IMUL:
		return f.emitInstIMUL(inst)
	case x86asm.IN:
		return f.emitInstIN(inst)
	case x86asm.INC:
		return f.emitInstINC(inst)
	case x86asm.INSB:
		return f.emitInstINSB(inst)
	case x86asm.INSD:
		return f.emitInstINSD(inst)
	case x86asm.INSERTPS:
		return f.emitInstINSERTPS(inst)
	case x86asm.INSW:
		return f.emitInstINSW(inst)
	case x86asm.INT:
		return f.emitInstINT(inst)
	case x86asm.INTO:
		return f.emitInstINTO(inst)
	case x86asm.INVD:
		return f.emitInstINVD(inst)
	case x86asm.INVLPG:
		return f.emitInstINVLPG(inst)
	case x86asm.INVPCID:
		return f.emitInstINVPCID(inst)
	case x86asm.IRET:
		return f.emitInstIRET(inst)
	case x86asm.IRETD:
		return f.emitInstIRETD(inst)
	case x86asm.IRETQ:
		return f.emitInstIRETQ(inst)
	case x86asm.JA:
		return f.emitInstJA(inst)
	case x86asm.JAE:
		return f.emitInstJAE(inst)
	case x86asm.JB:
		return f.emitInstJB(inst)
	case x86asm.JBE:
		return f.emitInstJBE(inst)
	case x86asm.JCXZ:
		return f.emitInstJCXZ(inst)
	case x86asm.JE:
		return f.emitInstJE(inst)
	case x86asm.JECXZ:
		return f.emitInstJECXZ(inst)
	case x86asm.JG:
		return f.emitInstJG(inst)
	case x86asm.JGE:
		return f.emitInstJGE(inst)
	case x86asm.JL:
		return f.emitInstJL(inst)
	case x86asm.JLE:
		return f.emitInstJLE(inst)
	case x86asm.JMP:
		return f.emitInstJMP(inst)
	case x86asm.JNE:
		return f.emitInstJNE(inst)
	case x86asm.JNO:
		return f.emitInstJNO(inst)
	case x86asm.JNP:
		return f.emitInstJNP(inst)
	case x86asm.JNS:
		return f.emitInstJNS(inst)
	case x86asm.JO:
		return f.emitInstJO(inst)
	case x86asm.JP:
		return f.emitInstJP(inst)
	case x86asm.JRCXZ:
		return f.emitInstJRCXZ(inst)
	case x86asm.JS:
		return f.emitInstJS(inst)
	case x86asm.LAHF:
		return f.emitInstLAHF(inst)
	case x86asm.LAR:
		return f.emitInstLAR(inst)
	case x86asm.LCALL:
		return f.emitInstLCALL(inst)
	case x86asm.LDDQU:
		return f.emitInstLDDQU(inst)
	case x86asm.LDMXCSR:
		return f.emitInstLDMXCSR(inst)
	case x86asm.LDS:
		return f.emitInstLDS(inst)
	case x86asm.LEA:
		return f.emitInstLEA(inst)
	case x86asm.LEAVE:
		return f.emitInstLEAVE(inst)
	case x86asm.LES:
		return f.emitInstLES(inst)
	case x86asm.LFENCE:
		return f.emitInstLFENCE(inst)
	case x86asm.LFS:
		return f.emitInstLFS(inst)
	case x86asm.LGDT:
		return f.emitInstLGDT(inst)
	case x86asm.LGS:
		return f.emitInstLGS(inst)
	case x86asm.LIDT:
		return f.emitInstLIDT(inst)
	case x86asm.LJMP:
		return f.emitInstLJMP(inst)
	case x86asm.LLDT:
		return f.emitInstLLDT(inst)
	case x86asm.LMSW:
		return f.emitInstLMSW(inst)
	case x86asm.LODSB:
		return f.emitInstLODSB(inst)
	case x86asm.LODSD:
		return f.emitInstLODSD(inst)
	case x86asm.LODSQ:
		return f.emitInstLODSQ(inst)
	case x86asm.LODSW:
		return f.emitInstLODSW(inst)
	case x86asm.LOOP:
		return f.emitInstLOOP(inst)
	case x86asm.LOOPE:
		return f.emitInstLOOPE(inst)
	case x86asm.LOOPNE:
		return f.emitInstLOOPNE(inst)
	case x86asm.LRET:
		return f.emitInstLRET(inst)
	case x86asm.LSL:
		return f.emitInstLSL(inst)
	case x86asm.LSS:
		return f.emitInstLSS(inst)
	case x86asm.LTR:
		return f.emitInstLTR(inst)
	case x86asm.LZCNT:
		return f.emitInstLZCNT(inst)
	case x86asm.MASKMOVDQU:
		return f.emitInstMASKMOVDQU(inst)
	case x86asm.MASKMOVQ:
		return f.emitInstMASKMOVQ(inst)
	case x86asm.MAXPD:
		return f.emitInstMAXPD(inst)
	case x86asm.MAXPS:
		return f.emitInstMAXPS(inst)
	case x86asm.MAXSD:
		return f.emitInstMAXSD(inst)
	case x86asm.MAXSS:
		return f.emitInstMAXSS(inst)
	case x86asm.MFENCE:
		return f.emitInstMFENCE(inst)
	case x86asm.MINPD:
		return f.emitInstMINPD(inst)
	case x86asm.MINPS:
		return f.emitInstMINPS(inst)
	case x86asm.MINSD:
		return f.emitInstMINSD(inst)
	case x86asm.MINSS:
		return f.emitInstMINSS(inst)
	case x86asm.MONITOR:
		return f.emitInstMONITOR(inst)
	case x86asm.MOV:
		return f.emitInstMOV(inst)
	case x86asm.MOVAPD:
		return f.emitInstMOVAPD(inst)
	case x86asm.MOVAPS:
		return f.emitInstMOVAPS(inst)
	case x86asm.MOVBE:
		return f.emitInstMOVBE(inst)
	case x86asm.MOVD:
		return f.emitInstMOVD(inst)
	case x86asm.MOVDDUP:
		return f.emitInstMOVDDUP(inst)
	case x86asm.MOVDQ2Q:
		return f.emitInstMOVDQ2Q(inst)
	case x86asm.MOVDQA:
		return f.emitInstMOVDQA(inst)
	case x86asm.MOVDQU:
		return f.emitInstMOVDQU(inst)
	case x86asm.MOVHLPS:
		return f.emitInstMOVHLPS(inst)
	case x86asm.MOVHPD:
		return f.emitInstMOVHPD(inst)
	case x86asm.MOVHPS:
		return f.emitInstMOVHPS(inst)
	case x86asm.MOVLHPS:
		return f.emitInstMOVLHPS(inst)
	case x86asm.MOVLPD:
		return f.emitInstMOVLPD(inst)
	case x86asm.MOVLPS:
		return f.emitInstMOVLPS(inst)
	case x86asm.MOVMSKPD:
		return f.emitInstMOVMSKPD(inst)
	case x86asm.MOVMSKPS:
		return f.emitInstMOVMSKPS(inst)
	case x86asm.MOVNTDQ:
		return f.emitInstMOVNTDQ(inst)
	case x86asm.MOVNTDQA:
		return f.emitInstMOVNTDQA(inst)
	case x86asm.MOVNTI:
		return f.emitInstMOVNTI(inst)
	case x86asm.MOVNTPD:
		return f.emitInstMOVNTPD(inst)
	case x86asm.MOVNTPS:
		return f.emitInstMOVNTPS(inst)
	case x86asm.MOVNTQ:
		return f.emitInstMOVNTQ(inst)
	case x86asm.MOVNTSD:
		return f.emitInstMOVNTSD(inst)
	case x86asm.MOVNTSS:
		return f.emitInstMOVNTSS(inst)
	case x86asm.MOVQ:
		return f.emitInstMOVQ(inst)
	case x86asm.MOVQ2DQ:
		return f.emitInstMOVQ2DQ(inst)
	case x86asm.MOVSB:
		return f.emitInstMOVSB(inst)
	case x86asm.MOVSD:
		return f.emitInstMOVSD(inst)
	case x86asm.MOVSD_XMM:
		return f.emitInstMOVSD_XMM(inst)
	case x86asm.MOVSHDUP:
		return f.emitInstMOVSHDUP(inst)
	case x86asm.MOVSLDUP:
		return f.emitInstMOVSLDUP(inst)
	case x86asm.MOVSQ:
		return f.emitInstMOVSQ(inst)
	case x86asm.MOVSS:
		return f.emitInstMOVSS(inst)
	case x86asm.MOVSW:
		return f.emitInstMOVSW(inst)
	case x86asm.MOVSX:
		return f.emitInstMOVSX(inst)
	case x86asm.MOVSXD:
		return f.emitInstMOVSXD(inst)
	case x86asm.MOVUPD:
		return f.emitInstMOVUPD(inst)
	case x86asm.MOVUPS:
		return f.emitInstMOVUPS(inst)
	case x86asm.MOVZX:
		return f.emitInstMOVZX(inst)
	case x86asm.MPSADBW:
		return f.emitInstMPSADBW(inst)
	case x86asm.MUL:
		return f.emitInstMUL(inst)
	case x86asm.MULPD:
		return f.emitInstMULPD(inst)
	case x86asm.MULPS:
		return f.emitInstMULPS(inst)
	case x86asm.MULSD:
		return f.emitInstMULSD(inst)
	case x86asm.MULSS:
		return f.emitInstMULSS(inst)
	case x86asm.MWAIT:
		return f.emitInstMWAIT(inst)
	case x86asm.NEG:
		return f.emitInstNEG(inst)
	case x86asm.NOP:
		return f.emitInstNOP(inst)
	case x86asm.NOT:
		return f.emitInstNOT(inst)
	case x86asm.OR:
		return f.emitInstOR(inst)
	case x86asm.ORPD:
		return f.emitInstORPD(inst)
	case x86asm.ORPS:
		return f.emitInstORPS(inst)
	case x86asm.OUT:
		return f.emitInstOUT(inst)
	case x86asm.OUTSB:
		return f.emitInstOUTSB(inst)
	case x86asm.OUTSD:
		return f.emitInstOUTSD(inst)
	case x86asm.OUTSW:
		return f.emitInstOUTSW(inst)
	case x86asm.PABSB:
		return f.emitInstPABSB(inst)
	case x86asm.PABSD:
		return f.emitInstPABSD(inst)
	case x86asm.PABSW:
		return f.emitInstPABSW(inst)
	case x86asm.PACKSSDW:
		return f.emitInstPACKSSDW(inst)
	case x86asm.PACKSSWB:
		return f.emitInstPACKSSWB(inst)
	case x86asm.PACKUSDW:
		return f.emitInstPACKUSDW(inst)
	case x86asm.PACKUSWB:
		return f.emitInstPACKUSWB(inst)
	case x86asm.PADDB:
		return f.emitInstPADDB(inst)
	case x86asm.PADDD:
		return f.emitInstPADDD(inst)
	case x86asm.PADDQ:
		return f.emitInstPADDQ(inst)
	case x86asm.PADDSB:
		return f.emitInstPADDSB(inst)
	case x86asm.PADDSW:
		return f.emitInstPADDSW(inst)
	case x86asm.PADDUSB:
		return f.emitInstPADDUSB(inst)
	case x86asm.PADDUSW:
		return f.emitInstPADDUSW(inst)
	case x86asm.PADDW:
		return f.emitInstPADDW(inst)
	case x86asm.PALIGNR:
		return f.emitInstPALIGNR(inst)
	case x86asm.PAND:
		return f.emitInstPAND(inst)
	case x86asm.PANDN:
		return f.emitInstPANDN(inst)
	case x86asm.PAUSE:
		return f.emitInstPAUSE(inst)
	case x86asm.PAVGB:
		return f.emitInstPAVGB(inst)
	case x86asm.PAVGW:
		return f.emitInstPAVGW(inst)
	case x86asm.PBLENDVB:
		return f.emitInstPBLENDVB(inst)
	case x86asm.PBLENDW:
		return f.emitInstPBLENDW(inst)
	case x86asm.PCLMULQDQ:
		return f.emitInstPCLMULQDQ(inst)
	case x86asm.PCMPEQB:
		return f.emitInstPCMPEQB(inst)
	case x86asm.PCMPEQD:
		return f.emitInstPCMPEQD(inst)
	case x86asm.PCMPEQQ:
		return f.emitInstPCMPEQQ(inst)
	case x86asm.PCMPEQW:
		return f.emitInstPCMPEQW(inst)
	case x86asm.PCMPESTRI:
		return f.emitInstPCMPESTRI(inst)
	case x86asm.PCMPESTRM:
		return f.emitInstPCMPESTRM(inst)
	case x86asm.PCMPGTB:
		return f.emitInstPCMPGTB(inst)
	case x86asm.PCMPGTD:
		return f.emitInstPCMPGTD(inst)
	case x86asm.PCMPGTQ:
		return f.emitInstPCMPGTQ(inst)
	case x86asm.PCMPGTW:
		return f.emitInstPCMPGTW(inst)
	case x86asm.PCMPISTRI:
		return f.emitInstPCMPISTRI(inst)
	case x86asm.PCMPISTRM:
		return f.emitInstPCMPISTRM(inst)
	case x86asm.PEXTRB:
		return f.emitInstPEXTRB(inst)
	case x86asm.PEXTRD:
		return f.emitInstPEXTRD(inst)
	case x86asm.PEXTRQ:
		return f.emitInstPEXTRQ(inst)
	case x86asm.PEXTRW:
		return f.emitInstPEXTRW(inst)
	case x86asm.PHADDD:
		return f.emitInstPHADDD(inst)
	case x86asm.PHADDSW:
		return f.emitInstPHADDSW(inst)
	case x86asm.PHADDW:
		return f.emitInstPHADDW(inst)
	case x86asm.PHMINPOSUW:
		return f.emitInstPHMINPOSUW(inst)
	case x86asm.PHSUBD:
		return f.emitInstPHSUBD(inst)
	case x86asm.PHSUBSW:
		return f.emitInstPHSUBSW(inst)
	case x86asm.PHSUBW:
		return f.emitInstPHSUBW(inst)
	case x86asm.PINSRB:
		return f.emitInstPINSRB(inst)
	case x86asm.PINSRD:
		return f.emitInstPINSRD(inst)
	case x86asm.PINSRQ:
		return f.emitInstPINSRQ(inst)
	case x86asm.PINSRW:
		return f.emitInstPINSRW(inst)
	case x86asm.PMADDUBSW:
		return f.emitInstPMADDUBSW(inst)
	case x86asm.PMADDWD:
		return f.emitInstPMADDWD(inst)
	case x86asm.PMAXSB:
		return f.emitInstPMAXSB(inst)
	case x86asm.PMAXSD:
		return f.emitInstPMAXSD(inst)
	case x86asm.PMAXSW:
		return f.emitInstPMAXSW(inst)
	case x86asm.PMAXUB:
		return f.emitInstPMAXUB(inst)
	case x86asm.PMAXUD:
		return f.emitInstPMAXUD(inst)
	case x86asm.PMAXUW:
		return f.emitInstPMAXUW(inst)
	case x86asm.PMINSB:
		return f.emitInstPMINSB(inst)
	case x86asm.PMINSD:
		return f.emitInstPMINSD(inst)
	case x86asm.PMINSW:
		return f.emitInstPMINSW(inst)
	case x86asm.PMINUB:
		return f.emitInstPMINUB(inst)
	case x86asm.PMINUD:
		return f.emitInstPMINUD(inst)
	case x86asm.PMINUW:
		return f.emitInstPMINUW(inst)
	case x86asm.PMOVMSKB:
		return f.emitInstPMOVMSKB(inst)
	case x86asm.PMOVSXBD:
		return f.emitInstPMOVSXBD(inst)
	case x86asm.PMOVSXBQ:
		return f.emitInstPMOVSXBQ(inst)
	case x86asm.PMOVSXBW:
		return f.emitInstPMOVSXBW(inst)
	case x86asm.PMOVSXDQ:
		return f.emitInstPMOVSXDQ(inst)
	case x86asm.PMOVSXWD:
		return f.emitInstPMOVSXWD(inst)
	case x86asm.PMOVSXWQ:
		return f.emitInstPMOVSXWQ(inst)
	case x86asm.PMOVZXBD:
		return f.emitInstPMOVZXBD(inst)
	case x86asm.PMOVZXBQ:
		return f.emitInstPMOVZXBQ(inst)
	case x86asm.PMOVZXBW:
		return f.emitInstPMOVZXBW(inst)
	case x86asm.PMOVZXDQ:
		return f.emitInstPMOVZXDQ(inst)
	case x86asm.PMOVZXWD:
		return f.emitInstPMOVZXWD(inst)
	case x86asm.PMOVZXWQ:
		return f.emitInstPMOVZXWQ(inst)
	case x86asm.PMULDQ:
		return f.emitInstPMULDQ(inst)
	case x86asm.PMULHRSW:
		return f.emitInstPMULHRSW(inst)
	case x86asm.PMULHUW:
		return f.emitInstPMULHUW(inst)
	case x86asm.PMULHW:
		return f.emitInstPMULHW(inst)
	case x86asm.PMULLD:
		return f.emitInstPMULLD(inst)
	case x86asm.PMULLW:
		return f.emitInstPMULLW(inst)
	case x86asm.PMULUDQ:
		return f.emitInstPMULUDQ(inst)
	case x86asm.POP:
		return f.emitInstPOP(inst)
	case x86asm.POPA:
		return f.emitInstPOPA(inst)
	case x86asm.POPAD:
		return f.emitInstPOPAD(inst)
	case x86asm.POPCNT:
		return f.emitInstPOPCNT(inst)
	case x86asm.POPF:
		return f.emitInstPOPF(inst)
	case x86asm.POPFD:
		return f.emitInstPOPFD(inst)
	case x86asm.POPFQ:
		return f.emitInstPOPFQ(inst)
	case x86asm.POR:
		return f.emitInstPOR(inst)
	case x86asm.PREFETCHNTA:
		return f.emitInstPREFETCHNTA(inst)
	case x86asm.PREFETCHT0:
		return f.emitInstPREFETCHT0(inst)
	case x86asm.PREFETCHT1:
		return f.emitInstPREFETCHT1(inst)
	case x86asm.PREFETCHT2:
		return f.emitInstPREFETCHT2(inst)
	case x86asm.PREFETCHW:
		return f.emitInstPREFETCHW(inst)
	case x86asm.PSADBW:
		return f.emitInstPSADBW(inst)
	case x86asm.PSHUFB:
		return f.emitInstPSHUFB(inst)
	case x86asm.PSHUFD:
		return f.emitInstPSHUFD(inst)
	case x86asm.PSHUFHW:
		return f.emitInstPSHUFHW(inst)
	case x86asm.PSHUFLW:
		return f.emitInstPSHUFLW(inst)
	case x86asm.PSHUFW:
		return f.emitInstPSHUFW(inst)
	case x86asm.PSIGNB:
		return f.emitInstPSIGNB(inst)
	case x86asm.PSIGND:
		return f.emitInstPSIGND(inst)
	case x86asm.PSIGNW:
		return f.emitInstPSIGNW(inst)
	case x86asm.PSLLD:
		return f.emitInstPSLLD(inst)
	case x86asm.PSLLDQ:
		return f.emitInstPSLLDQ(inst)
	case x86asm.PSLLQ:
		return f.emitInstPSLLQ(inst)
	case x86asm.PSLLW:
		return f.emitInstPSLLW(inst)
	case x86asm.PSRAD:
		return f.emitInstPSRAD(inst)
	case x86asm.PSRAW:
		return f.emitInstPSRAW(inst)
	case x86asm.PSRLD:
		return f.emitInstPSRLD(inst)
	case x86asm.PSRLDQ:
		return f.emitInstPSRLDQ(inst)
	case x86asm.PSRLQ:
		return f.emitInstPSRLQ(inst)
	case x86asm.PSRLW:
		return f.emitInstPSRLW(inst)
	case x86asm.PSUBB:
		return f.emitInstPSUBB(inst)
	case x86asm.PSUBD:
		return f.emitInstPSUBD(inst)
	case x86asm.PSUBQ:
		return f.emitInstPSUBQ(inst)
	case x86asm.PSUBSB:
		return f.emitInstPSUBSB(inst)
	case x86asm.PSUBSW:
		return f.emitInstPSUBSW(inst)
	case x86asm.PSUBUSB:
		return f.emitInstPSUBUSB(inst)
	case x86asm.PSUBUSW:
		return f.emitInstPSUBUSW(inst)
	case x86asm.PSUBW:
		return f.emitInstPSUBW(inst)
	case x86asm.PTEST:
		return f.emitInstPTEST(inst)
	case x86asm.PUNPCKHBW:
		return f.emitInstPUNPCKHBW(inst)
	case x86asm.PUNPCKHDQ:
		return f.emitInstPUNPCKHDQ(inst)
	case x86asm.PUNPCKHQDQ:
		return f.emitInstPUNPCKHQDQ(inst)
	case x86asm.PUNPCKHWD:
		return f.emitInstPUNPCKHWD(inst)
	case x86asm.PUNPCKLBW:
		return f.emitInstPUNPCKLBW(inst)
	case x86asm.PUNPCKLDQ:
		return f.emitInstPUNPCKLDQ(inst)
	case x86asm.PUNPCKLQDQ:
		return f.emitInstPUNPCKLQDQ(inst)
	case x86asm.PUNPCKLWD:
		return f.emitInstPUNPCKLWD(inst)
	case x86asm.PUSH:
		return f.emitInstPUSH(inst)
	case x86asm.PUSHA:
		return f.emitInstPUSHA(inst)
	case x86asm.PUSHAD:
		return f.emitInstPUSHAD(inst)
	case x86asm.PUSHF:
		return f.emitInstPUSHF(inst)
	case x86asm.PUSHFD:
		return f.emitInstPUSHFD(inst)
	case x86asm.PUSHFQ:
		return f.emitInstPUSHFQ(inst)
	case x86asm.PXOR:
		return f.emitInstPXOR(inst)
	case x86asm.RCL:
		return f.emitInstRCL(inst)
	case x86asm.RCPPS:
		return f.emitInstRCPPS(inst)
	case x86asm.RCPSS:
		return f.emitInstRCPSS(inst)
	case x86asm.RCR:
		return f.emitInstRCR(inst)
	case x86asm.RDFSBASE:
		return f.emitInstRDFSBASE(inst)
	case x86asm.RDGSBASE:
		return f.emitInstRDGSBASE(inst)
	case x86asm.RDMSR:
		return f.emitInstRDMSR(inst)
	case x86asm.RDPMC:
		return f.emitInstRDPMC(inst)
	case x86asm.RDRAND:
		return f.emitInstRDRAND(inst)
	case x86asm.RDTSC:
		return f.emitInstRDTSC(inst)
	case x86asm.RDTSCP:
		return f.emitInstRDTSCP(inst)
	case x86asm.RET:
		return f.emitInstRET(inst)
	case x86asm.ROL:
		return f.emitInstROL(inst)
	case x86asm.ROR:
		return f.emitInstROR(inst)
	case x86asm.ROUNDPD:
		return f.emitInstROUNDPD(inst)
	case x86asm.ROUNDPS:
		return f.emitInstROUNDPS(inst)
	case x86asm.ROUNDSD:
		return f.emitInstROUNDSD(inst)
	case x86asm.ROUNDSS:
		return f.emitInstROUNDSS(inst)
	case x86asm.RSM:
		return f.emitInstRSM(inst)
	case x86asm.RSQRTPS:
		return f.emitInstRSQRTPS(inst)
	case x86asm.RSQRTSS:
		return f.emitInstRSQRTSS(inst)
	case x86asm.SAHF:
		return f.emitInstSAHF(inst)
	case x86asm.SAR:
		return f.emitInstSAR(inst)
	case x86asm.SBB:
		return f.emitInstSBB(inst)
	case x86asm.SCASB:
		return f.emitInstSCASB(inst)
	case x86asm.SCASD:
		return f.emitInstSCASD(inst)
	case x86asm.SCASQ:
		return f.emitInstSCASQ(inst)
	case x86asm.SCASW:
		return f.emitInstSCASW(inst)
	case x86asm.SETA:
		return f.emitInstSETA(inst)
	case x86asm.SETAE:
		return f.emitInstSETAE(inst)
	case x86asm.SETB:
		return f.emitInstSETB(inst)
	case x86asm.SETBE:
		return f.emitInstSETBE(inst)
	case x86asm.SETE:
		return f.emitInstSETE(inst)
	case x86asm.SETG:
		return f.emitInstSETG(inst)
	case x86asm.SETGE:
		return f.emitInstSETGE(inst)
	case x86asm.SETL:
		return f.emitInstSETL(inst)
	case x86asm.SETLE:
		return f.emitInstSETLE(inst)
	case x86asm.SETNE:
		return f.emitInstSETNE(inst)
	case x86asm.SETNO:
		return f.emitInstSETNO(inst)
	case x86asm.SETNP:
		return f.emitInstSETNP(inst)
	case x86asm.SETNS:
		return f.emitInstSETNS(inst)
	case x86asm.SETO:
		return f.emitInstSETO(inst)
	case x86asm.SETP:
		return f.emitInstSETP(inst)
	case x86asm.SETS:
		return f.emitInstSETS(inst)
	case x86asm.SFENCE:
		return f.emitInstSFENCE(inst)
	case x86asm.SGDT:
		return f.emitInstSGDT(inst)
	case x86asm.SHL:
		return f.emitInstSHL(inst)
	case x86asm.SHLD:
		return f.emitInstSHLD(inst)
	case x86asm.SHR:
		return f.emitInstSHR(inst)
	case x86asm.SHRD:
		return f.emitInstSHRD(inst)
	case x86asm.SHUFPD:
		return f.emitInstSHUFPD(inst)
	case x86asm.SHUFPS:
		return f.emitInstSHUFPS(inst)
	case x86asm.SIDT:
		return f.emitInstSIDT(inst)
	case x86asm.SLDT:
		return f.emitInstSLDT(inst)
	case x86asm.SMSW:
		return f.emitInstSMSW(inst)
	case x86asm.SQRTPD:
		return f.emitInstSQRTPD(inst)
	case x86asm.SQRTPS:
		return f.emitInstSQRTPS(inst)
	case x86asm.SQRTSD:
		return f.emitInstSQRTSD(inst)
	case x86asm.SQRTSS:
		return f.emitInstSQRTSS(inst)
	case x86asm.STC:
		return f.emitInstSTC(inst)
	case x86asm.STD:
		return f.emitInstSTD(inst)
	case x86asm.STI:
		return f.emitInstSTI(inst)
	case x86asm.STMXCSR:
		return f.emitInstSTMXCSR(inst)
	case x86asm.STOSB:
		return f.emitInstSTOSB(inst)
	case x86asm.STOSD:
		return f.emitInstSTOSD(inst)
	case x86asm.STOSQ:
		return f.emitInstSTOSQ(inst)
	case x86asm.STOSW:
		return f.emitInstSTOSW(inst)
	case x86asm.STR:
		return f.emitInstSTR(inst)
	case x86asm.SUB:
		return f.emitInstSUB(inst)
	case x86asm.SUBPD:
		return f.emitInstSUBPD(inst)
	case x86asm.SUBPS:
		return f.emitInstSUBPS(inst)
	case x86asm.SUBSD:
		return f.emitInstSUBSD(inst)
	case x86asm.SUBSS:
		return f.emitInstSUBSS(inst)
	case x86asm.SWAPGS:
		return f.emitInstSWAPGS(inst)
	case x86asm.SYSCALL:
		return f.emitInstSYSCALL(inst)
	case x86asm.SYSENTER:
		return f.emitInstSYSENTER(inst)
	case x86asm.SYSEXIT:
		return f.emitInstSYSEXIT(inst)
	case x86asm.SYSRET:
		return f.emitInstSYSRET(inst)
	case x86asm.TEST:
		return f.emitInstTEST(inst)
	case x86asm.TZCNT:
		return f.emitInstTZCNT(inst)
	case x86asm.UCOMISD:
		return f.emitInstUCOMISD(inst)
	case x86asm.UCOMISS:
		return f.emitInstUCOMISS(inst)
	case x86asm.UD1:
		return f.emitInstUD1(inst)
	case x86asm.UD2:
		return f.emitInstUD2(inst)
	case x86asm.UNPCKHPD:
		return f.emitInstUNPCKHPD(inst)
	case x86asm.UNPCKHPS:
		return f.emitInstUNPCKHPS(inst)
	case x86asm.UNPCKLPD:
		return f.emitInstUNPCKLPD(inst)
	case x86asm.UNPCKLPS:
		return f.emitInstUNPCKLPS(inst)
	case x86asm.VERR:
		return f.emitInstVERR(inst)
	case x86asm.VERW:
		return f.emitInstVERW(inst)
	case x86asm.VMOVDQA:
		return f.emitInstVMOVDQA(inst)
	case x86asm.VMOVDQU:
		return f.emitInstVMOVDQU(inst)
	case x86asm.VMOVNTDQ:
		return f.emitInstVMOVNTDQ(inst)
	case x86asm.VMOVNTDQA:
		return f.emitInstVMOVNTDQA(inst)
	case x86asm.VZEROUPPER:
		return f.emitInstVZEROUPPER(inst)
	case x86asm.WBINVD:
		return f.emitInstWBINVD(inst)
	case x86asm.WRFSBASE:
		return f.emitInstWRFSBASE(inst)
	case x86asm.WRGSBASE:
		return f.emitInstWRGSBASE(inst)
	case x86asm.WRMSR:
		return f.emitInstWRMSR(inst)
	case x86asm.XABORT:
		return f.emitInstXABORT(inst)
	case x86asm.XADD:
		return f.emitInstXADD(inst)
	case x86asm.XBEGIN:
		return f.emitInstXBEGIN(inst)
	case x86asm.XCHG:
		return f.emitInstXCHG(inst)
	case x86asm.XEND:
		return f.emitInstXEND(inst)
	case x86asm.XGETBV:
		return f.emitInstXGETBV(inst)
	case x86asm.XLATB:
		return f.emitInstXLATB(inst)
	case x86asm.XOR:
		return f.emitInstXOR(inst)
	case x86asm.XORPD:
		return f.emitInstXORPD(inst)
	case x86asm.XORPS:
		return f.emitInstXORPS(inst)
	case x86asm.XRSTOR:
		return f.emitInstXRSTOR(inst)
	case x86asm.XRSTOR64:
		return f.emitInstXRSTOR64(inst)
	case x86asm.XRSTORS:
		return f.emitInstXRSTORS(inst)
	case x86asm.XRSTORS64:
		return f.emitInstXRSTORS64(inst)
	case x86asm.XSAVE:
		return f.emitInstXSAVE(inst)
	case x86asm.XSAVE64:
		return f.emitInstXSAVE64(inst)
	case x86asm.XSAVEC:
		return f.emitInstXSAVEC(inst)
	case x86asm.XSAVEC64:
		return f.emitInstXSAVEC64(inst)
	case x86asm.XSAVEOPT:
		return f.emitInstXSAVEOPT(inst)
	case x86asm.XSAVEOPT64:
		return f.emitInstXSAVEOPT64(inst)
	case x86asm.XSAVES:
		return f.emitInstXSAVES(inst)
	case x86asm.XSAVES64:
		return f.emitInstXSAVES64(inst)
	case x86asm.XSETBV:
		return f.emitInstXSETBV(inst)
	case x86asm.XTEST:
		return f.emitInstXTEST(inst)
	default:
		panic(fmt.Errorf("support for x86 instruction opcode %v not yet implemented", inst.Op))
	}
}

// --- [ AAA ] -----------------------------------------------------------------

// emitInst translates the given x86 AAA instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstAAA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAAA: not yet implemented")
}

// --- [ AAD ] -----------------------------------------------------------------

// emitInst translates the given x86 AAD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstAAD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAAD: not yet implemented")
}

// --- [ AAM ] -----------------------------------------------------------------

// emitInst translates the given x86 AAM instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstAAM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAAM: not yet implemented")
}

// --- [ AAS ] -----------------------------------------------------------------

// emitInst translates the given x86 AAS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstAAS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAAS: not yet implemented")
}

// --- [ ADC ] -----------------------------------------------------------------

// emitInst translates the given x86 ADC instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstADC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADC: not yet implemented")
}

// --- [ ADD ] -----------------------------------------------------------------

// emitInst translates the given x86 ADD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstADD(inst *Inst) error {
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewAdd(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ ADDPD ] ---------------------------------------------------------------

// emitInst translates the given x86 ADDPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstADDPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDPD: not yet implemented")
}

// --- [ ADDPS ] ---------------------------------------------------------------

// emitInst translates the given x86 ADDPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstADDPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDPS: not yet implemented")
}

// --- [ ADDSD ] ---------------------------------------------------------------

// emitInst translates the given x86 ADDSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstADDSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDSD: not yet implemented")
}

// --- [ ADDSS ] ---------------------------------------------------------------

// emitInst translates the given x86 ADDSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstADDSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDSS: not yet implemented")
}

// --- [ ADDSUBPD ] ------------------------------------------------------------

// emitInst translates the given x86 ADDSUBPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstADDSUBPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDSUBPD: not yet implemented")
}

// --- [ ADDSUBPS ] ------------------------------------------------------------

// emitInst translates the given x86 ADDSUBPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstADDSUBPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDSUBPS: not yet implemented")
}

// --- [ AESDEC ] --------------------------------------------------------------

// emitInst translates the given x86 AESDEC instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstAESDEC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESDEC: not yet implemented")
}

// --- [ AESDECLAST ] ----------------------------------------------------------

// emitInst translates the given x86 AESDECLAST instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstAESDECLAST(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESDECLAST: not yet implemented")
}

// --- [ AESENC ] --------------------------------------------------------------

// emitInst translates the given x86 AESENC instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstAESENC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESENC: not yet implemented")
}

// --- [ AESENCLAST ] ----------------------------------------------------------

// emitInst translates the given x86 AESENCLAST instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstAESENCLAST(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESENCLAST: not yet implemented")
}

// --- [ AESIMC ] --------------------------------------------------------------

// emitInst translates the given x86 AESIMC instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstAESIMC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESIMC: not yet implemented")
}

// --- [ AESKEYGENASSIST ] -----------------------------------------------------

// emitInst translates the given x86 AESKEYGENASSIST instruction to LLVM IR,
// emitting code to f.
func (f *Func) emitInstAESKEYGENASSIST(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESKEYGENASSIST: not yet implemented")
}

// --- [ AND ] -----------------------------------------------------------------

// emitInst translates the given x86 AND instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstAND(inst *Inst) error {
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewAnd(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ ANDNPD ] --------------------------------------------------------------

// emitInst translates the given x86 ANDNPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstANDNPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstANDNPD: not yet implemented")
}

// --- [ ANDNPS ] --------------------------------------------------------------

// emitInst translates the given x86 ANDNPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstANDNPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstANDNPS: not yet implemented")
}

// --- [ ANDPD ] ---------------------------------------------------------------

// emitInst translates the given x86 ANDPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstANDPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstANDPD: not yet implemented")
}

// --- [ ANDPS ] ---------------------------------------------------------------

// emitInst translates the given x86 ANDPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstANDPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstANDPS: not yet implemented")
}

// --- [ ARPL ] ----------------------------------------------------------------

// emitInst translates the given x86 ARPL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstARPL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstARPL: not yet implemented")
}

// --- [ BLENDPD ] -------------------------------------------------------------

// emitInst translates the given x86 BLENDPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstBLENDPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBLENDPD: not yet implemented")
}

// --- [ BLENDPS ] -------------------------------------------------------------

// emitInst translates the given x86 BLENDPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstBLENDPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBLENDPS: not yet implemented")
}

// --- [ BLENDVPD ] ------------------------------------------------------------

// emitInst translates the given x86 BLENDVPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstBLENDVPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBLENDVPD: not yet implemented")
}

// --- [ BLENDVPS ] ------------------------------------------------------------

// emitInst translates the given x86 BLENDVPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstBLENDVPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBLENDVPS: not yet implemented")
}

// --- [ BOUND ] ---------------------------------------------------------------

// emitInst translates the given x86 BOUND instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstBOUND(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBOUND: not yet implemented")
}

// --- [ BSF ] -----------------------------------------------------------------

// emitInst translates the given x86 BSF instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstBSF(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBSF: not yet implemented")
}

// --- [ BSR ] -----------------------------------------------------------------

// emitInst translates the given x86 BSR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstBSR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBSR: not yet implemented")
}

// --- [ BSWAP ] ---------------------------------------------------------------

// emitInst translates the given x86 BSWAP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstBSWAP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBSWAP: not yet implemented")
}

// --- [ BT ] ------------------------------------------------------------------

// emitInst translates the given x86 BT instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstBT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBT: not yet implemented")
}

// --- [ BTC ] -----------------------------------------------------------------

// emitInst translates the given x86 BTC instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstBTC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBTC: not yet implemented")
}

// --- [ BTR ] -----------------------------------------------------------------

// emitInst translates the given x86 BTR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstBTR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBTR: not yet implemented")
}

// --- [ BTS ] -----------------------------------------------------------------

// emitInst translates the given x86 BTS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstBTS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBTS: not yet implemented")
}

// --- [ CALL ] ----------------------------------------------------------------

// emitInst translates the given x86 CALL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCALL(inst *Inst) error {
	// Locate callee information.
	callee, sig, callconv, ok := f.getFunc(inst.Arg(0))
	if !ok {
		panic(fmt.Errorf("unable to locate function for argument %v", inst.Arg(0)))
	}

	// Handle function arguments.
	var args []value.Value
	purge := int64(0)
	for i := range sig.Params {
		// Pass argument in register.
		switch callconv {
		case ir.CallConvX86_FastCall:
			switch i {
			case 0:
				arg := f.useReg(ECX)
				args = append(args, arg)
				continue
			case 1:
				arg := f.useReg(EDX)
				args = append(args, arg)
				continue
			}
		default:
			// TODO: Add support for more calling conventions.
		}
		// Pass argument on stack.
		arg := f.pop()
		args = append(args, arg)
		switch callconv {
		case ir.CallConvX86_FastCall, ir.CallConvX86_StdCall:
			// callee purge.
			purge += 4
		case ir.CallConvC:
			// caller purge; nothing to do.
		default:
			// TODO: Add support for more calling conventions.
		}
	}

	// Emit call instruction.
	result := f.cur.NewCall(callee, args...)

	// Handle purged arguments by callee.
	f.espDisp += purge

	// Handle return value.
	if !types.Equal(f.Sig.Ret, types.Void) {
		f.defReg(EAX, result)
	}
	return nil
}

// --- [ CBW ] -----------------------------------------------------------------

// emitInst translates the given x86 CBW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCBW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCBW: not yet implemented")
}

// --- [ CDQ ] -----------------------------------------------------------------

// emitInst translates the given x86 CDQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCDQ(inst *Inst) error {
	// EDX:EAX = sign-extend of EAX.
	eax := f.useReg(EAX)
	tmp := f.cur.NewLShr(eax, constant.NewInt(31, types.I32))
	cond := f.cur.NewTrunc(tmp, types.I1)
	targetTrue := &ir.BasicBlock{}
	targetFalse := &ir.BasicBlock{}
	exit := &ir.BasicBlock{}
	f.AppendBlock(targetTrue)
	f.AppendBlock(targetFalse)
	f.AppendBlock(exit)
	f.cur.NewCondBr(cond, targetTrue, targetFalse)
	f.cur = targetTrue
	f.defReg(EDX, constant.NewInt(0xFFFFFFFF, types.I32))
	f.cur = targetFalse
	f.defReg(EDX, constant.NewInt(0, types.I32))
	targetTrue.NewBr(exit)
	targetFalse.NewBr(exit)
	f.cur = exit
	return nil
}

// --- [ CDQE ] ----------------------------------------------------------------

// emitInst translates the given x86 CDQE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCDQE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCDQE: not yet implemented")
}

// --- [ CLC ] -----------------------------------------------------------------

// emitInst translates the given x86 CLC instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCLC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCLC: not yet implemented")
}

// --- [ CLD ] -----------------------------------------------------------------

// emitInst translates the given x86 CLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCLD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCLD: not yet implemented")
}

// --- [ CLFLUSH ] -------------------------------------------------------------

// emitInst translates the given x86 CLFLUSH instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCLFLUSH(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCLFLUSH: not yet implemented")
}

// --- [ CLI ] -----------------------------------------------------------------

// emitInst translates the given x86 CLI instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCLI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCLI: not yet implemented")
}

// --- [ CLTS ] ----------------------------------------------------------------

// emitInst translates the given x86 CLTS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCLTS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCLTS: not yet implemented")
}

// --- [ CMC ] -----------------------------------------------------------------

// emitInst translates the given x86 CMC instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMC: not yet implemented")
}

// --- [ CMOVA ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVA instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMOVA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVA: not yet implemented")
}

// --- [ CMOVAE ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVAE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMOVAE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVAE: not yet implemented")
}

// --- [ CMOVB ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMOVB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVB: not yet implemented")
}

// --- [ CMOVBE ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVBE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMOVBE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVBE: not yet implemented")
}

// --- [ CMOVE ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMOVE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVE: not yet implemented")
}

// --- [ CMOVG ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVG instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMOVG(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVG: not yet implemented")
}

// --- [ CMOVGE ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVGE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMOVGE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVGE: not yet implemented")
}

// --- [ CMOVL ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMOVL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVL: not yet implemented")
}

// --- [ CMOVLE ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVLE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMOVLE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVLE: not yet implemented")
}

// --- [ CMOVNE ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVNE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMOVNE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVNE: not yet implemented")
}

// --- [ CMOVNO ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVNO instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMOVNO(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVNO: not yet implemented")
}

// --- [ CMOVNP ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVNP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMOVNP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVNP: not yet implemented")
}

// --- [ CMOVNS ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVNS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMOVNS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVNS: not yet implemented")
}

// --- [ CMOVO ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVO instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMOVO(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVO: not yet implemented")
}

// --- [ CMOVP ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMOVP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVP: not yet implemented")
}

// --- [ CMOVS ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMOVS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVS: not yet implemented")
}

// --- [ CMP ] -----------------------------------------------------------------

// emitInst translates the given x86 CMP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMP(inst *Inst) error {
	// result = x SUB y; set CF, PF, AF, ZF, SF, and OF according to result.
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewSub(x, y)

	// CF (bit 0) Carry flag - Set if an arithmetic operation generates a carry
	// or a borrow out of the most- significant bit of the result; cleared
	// otherwise. This flag indicates an overflow condition for unsigned-integer
	// arithmetic. It is also used in multiple-precision arithmetic.

	// TODO: Add support for the CF status flag.

	// PF (bit 2) Parity flag - Set if the least-significant byte of the result
	// contains an even number of 1 bits; cleared otherwise.

	// TODO: Add support for the PF status flag.

	// AF (bit 4) Auxiliary Carry flag - Set if an arithmetic operation generates
	// a carry or a borrow out of bit 3 of the result; cleared otherwise. This
	// flag is used in binary-coded decimal (BCD) arithmetic.

	// TODO: Add support for the AF status flag.

	// ZF (bit 6) Zero flag - Set if the result is zero; cleared otherwise.
	zero := constant.NewInt(0, types.I32)
	zf := f.cur.NewICmp(ir.IntEQ, result, zero)
	f.defStatus(ZF, zf)

	// SF (bit 7) Sign flag - Set equal to the most-significant bit of the
	// result, which is the sign bit of a signed integer. (0 indicates a positive
	// value and 1 indicates a negative value.)

	// TODO: Add support for SF flag.

	// OF (bit 11) Overflow flag - Set if the integer result is too large a
	// positive number or too small a negative number (excluding the sign-bit) to
	// fit in the destination operand; cleared otherwise. This flag indicates an
	// overflow condition for signed-integer (two's complement) arithmetic.

	// TODO: Add support for the OF status flag.

	return nil
}

// --- [ CMPPD ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMPPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPPD: not yet implemented")
}

// --- [ CMPPS ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMPPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPPS: not yet implemented")
}

// --- [ CMPSB ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMPSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSB: not yet implemented")
}

// --- [ CMPSD ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMPSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSD: not yet implemented")
}

// --- [ CMPSD_XMM ] -----------------------------------------------------------

// emitInst translates the given x86 CMPSD_XMM instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMPSD_XMM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSD_XMM: not yet implemented")
}

// --- [ CMPSQ ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPSQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMPSQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSQ: not yet implemented")
}

// --- [ CMPSS ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMPSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSS: not yet implemented")
}

// --- [ CMPSW ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCMPSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSW: not yet implemented")
}

// --- [ CMPXCHG ] -------------------------------------------------------------

// emitInst translates the given x86 CMPXCHG instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMPXCHG(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPXCHG: not yet implemented")
}

// --- [ CMPXCHG16B ] ----------------------------------------------------------

// emitInst translates the given x86 CMPXCHG16B instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMPXCHG16B(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPXCHG16B: not yet implemented")
}

// --- [ CMPXCHG8B ] -----------------------------------------------------------

// emitInst translates the given x86 CMPXCHG8B instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCMPXCHG8B(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPXCHG8B: not yet implemented")
}

// --- [ COMISD ] --------------------------------------------------------------

// emitInst translates the given x86 COMISD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCOMISD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCOMISD: not yet implemented")
}

// --- [ COMISS ] --------------------------------------------------------------

// emitInst translates the given x86 COMISS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCOMISS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCOMISS: not yet implemented")
}

// --- [ CPUID ] ---------------------------------------------------------------

// emitInst translates the given x86 CPUID instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCPUID(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCPUID: not yet implemented")
}

// --- [ CQO ] -----------------------------------------------------------------

// emitInst translates the given x86 CQO instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCQO(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCQO: not yet implemented")
}

// --- [ CRC32 ] ---------------------------------------------------------------

// emitInst translates the given x86 CRC32 instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCRC32(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCRC32: not yet implemented")
}

// --- [ CVTDQ2PD ] ------------------------------------------------------------

// emitInst translates the given x86 CVTDQ2PD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTDQ2PD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTDQ2PD: not yet implemented")
}

// --- [ CVTDQ2PS ] ------------------------------------------------------------

// emitInst translates the given x86 CVTDQ2PS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTDQ2PS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTDQ2PS: not yet implemented")
}

// --- [ CVTPD2DQ ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPD2DQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTPD2DQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPD2DQ: not yet implemented")
}

// --- [ CVTPD2PI ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPD2PI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTPD2PI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPD2PI: not yet implemented")
}

// --- [ CVTPD2PS ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPD2PS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTPD2PS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPD2PS: not yet implemented")
}

// --- [ CVTPI2PD ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPI2PD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTPI2PD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPI2PD: not yet implemented")
}

// --- [ CVTPI2PS ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPI2PS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTPI2PS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPI2PS: not yet implemented")
}

// --- [ CVTPS2DQ ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPS2DQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTPS2DQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPS2DQ: not yet implemented")
}

// --- [ CVTPS2PD ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPS2PD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTPS2PD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPS2PD: not yet implemented")
}

// --- [ CVTPS2PI ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPS2PI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTPS2PI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPS2PI: not yet implemented")
}

// --- [ CVTSD2SI ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSD2SI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTSD2SI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSD2SI: not yet implemented")
}

// --- [ CVTSD2SS ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSD2SS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTSD2SS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSD2SS: not yet implemented")
}

// --- [ CVTSI2SD ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSI2SD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTSI2SD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSI2SD: not yet implemented")
}

// --- [ CVTSI2SS ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSI2SS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTSI2SS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSI2SS: not yet implemented")
}

// --- [ CVTSS2SD ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSS2SD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTSS2SD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSS2SD: not yet implemented")
}

// --- [ CVTSS2SI ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSS2SI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTSS2SI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSS2SI: not yet implemented")
}

// --- [ CVTTPD2DQ ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTPD2DQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTTPD2DQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTPD2DQ: not yet implemented")
}

// --- [ CVTTPD2PI ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTPD2PI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTTPD2PI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTPD2PI: not yet implemented")
}

// --- [ CVTTPS2DQ ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTPS2DQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTTPS2DQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTPS2DQ: not yet implemented")
}

// --- [ CVTTPS2PI ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTPS2PI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTTPS2PI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTPS2PI: not yet implemented")
}

// --- [ CVTTSD2SI ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTSD2SI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTTSD2SI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTSD2SI: not yet implemented")
}

// --- [ CVTTSS2SI ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTSS2SI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstCVTTSS2SI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTSS2SI: not yet implemented")
}

// --- [ CWD ] -----------------------------------------------------------------

// emitInst translates the given x86 CWD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCWD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCWD: not yet implemented")
}

// --- [ CWDE ] ----------------------------------------------------------------

// emitInst translates the given x86 CWDE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstCWDE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCWDE: not yet implemented")
}

// --- [ DAA ] -----------------------------------------------------------------

// emitInst translates the given x86 DAA instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstDAA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDAA: not yet implemented")
}

// --- [ DAS ] -----------------------------------------------------------------

// emitInst translates the given x86 DAS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstDAS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDAS: not yet implemented")
}

// --- [ DEC ] -----------------------------------------------------------------

// emitInst translates the given x86 DEC instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstDEC(inst *Inst) error {
	x := f.useArg(inst.Arg(0))
	one := constant.NewInt(1, types.I32)
	result := f.cur.NewSub(x, one)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ DIV ] -----------------------------------------------------------------

// emitInst translates the given x86 DIV instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstDIV(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDIV: not yet implemented")
}

// --- [ DIVPD ] ---------------------------------------------------------------

// emitInst translates the given x86 DIVPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstDIVPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDIVPD: not yet implemented")
}

// --- [ DIVPS ] ---------------------------------------------------------------

// emitInst translates the given x86 DIVPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstDIVPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDIVPS: not yet implemented")
}

// --- [ DIVSD ] ---------------------------------------------------------------

// emitInst translates the given x86 DIVSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstDIVSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDIVSD: not yet implemented")
}

// --- [ DIVSS ] ---------------------------------------------------------------

// emitInst translates the given x86 DIVSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstDIVSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDIVSS: not yet implemented")
}

// --- [ DPPD ] ----------------------------------------------------------------

// emitInst translates the given x86 DPPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstDPPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDPPD: not yet implemented")
}

// --- [ DPPS ] ----------------------------------------------------------------

// emitInst translates the given x86 DPPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstDPPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDPPS: not yet implemented")
}

// --- [ EMMS ] ----------------------------------------------------------------

// emitInst translates the given x86 EMMS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstEMMS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstEMMS: not yet implemented")
}

// --- [ ENTER ] ---------------------------------------------------------------

// emitInst translates the given x86 ENTER instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstENTER(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstENTER: not yet implemented")
}

// --- [ EXTRACTPS ] -----------------------------------------------------------

// emitInst translates the given x86 EXTRACTPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstEXTRACTPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstEXTRACTPS: not yet implemented")
}

// --- [ F2XM1 ] ---------------------------------------------------------------

// emitInst translates the given x86 F2XM1 instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstF2XM1(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstF2XM1: not yet implemented")
}

// --- [ FABS ] ----------------------------------------------------------------

// emitInst translates the given x86 FABS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFABS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFABS: not yet implemented")
}

// --- [ FADD ] ----------------------------------------------------------------

// emitInst translates the given x86 FADD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFADD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFADD: not yet implemented")
}

// --- [ FADDP ] ---------------------------------------------------------------

// emitInst translates the given x86 FADDP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFADDP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFADDP: not yet implemented")
}

// --- [ FBLD ] ----------------------------------------------------------------

// emitInst translates the given x86 FBLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFBLD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFBLD: not yet implemented")
}

// --- [ FBSTP ] ---------------------------------------------------------------

// emitInst translates the given x86 FBSTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFBSTP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFBSTP: not yet implemented")
}

// --- [ FCHS ] ----------------------------------------------------------------

// emitInst translates the given x86 FCHS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFCHS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCHS: not yet implemented")
}

// --- [ FCMOVB ] --------------------------------------------------------------

// emitInst translates the given x86 FCMOVB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFCMOVB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVB: not yet implemented")
}

// --- [ FCMOVBE ] -------------------------------------------------------------

// emitInst translates the given x86 FCMOVBE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFCMOVBE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVBE: not yet implemented")
}

// --- [ FCMOVE ] --------------------------------------------------------------

// emitInst translates the given x86 FCMOVE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFCMOVE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVE: not yet implemented")
}

// --- [ FCMOVNB ] -------------------------------------------------------------

// emitInst translates the given x86 FCMOVNB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFCMOVNB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVNB: not yet implemented")
}

// --- [ FCMOVNBE ] ------------------------------------------------------------

// emitInst translates the given x86 FCMOVNBE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFCMOVNBE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVNBE: not yet implemented")
}

// --- [ FCMOVNE ] -------------------------------------------------------------

// emitInst translates the given x86 FCMOVNE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFCMOVNE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVNE: not yet implemented")
}

// --- [ FCMOVNU ] -------------------------------------------------------------

// emitInst translates the given x86 FCMOVNU instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFCMOVNU(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVNU: not yet implemented")
}

// --- [ FCMOVU ] --------------------------------------------------------------

// emitInst translates the given x86 FCMOVU instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFCMOVU(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCMOVU: not yet implemented")
}

// --- [ FCOM ] ----------------------------------------------------------------

// emitInst translates the given x86 FCOM instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFCOM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCOM: not yet implemented")
}

// --- [ FCOMI ] ---------------------------------------------------------------

// emitInst translates the given x86 FCOMI instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFCOMI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCOMI: not yet implemented")
}

// --- [ FCOMIP ] --------------------------------------------------------------

// emitInst translates the given x86 FCOMIP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFCOMIP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCOMIP: not yet implemented")
}

// --- [ FCOMP ] ---------------------------------------------------------------

// emitInst translates the given x86 FCOMP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFCOMP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCOMP: not yet implemented")
}

// --- [ FCOMPP ] --------------------------------------------------------------

// emitInst translates the given x86 FCOMPP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFCOMPP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCOMPP: not yet implemented")
}

// --- [ FCOS ] ----------------------------------------------------------------

// emitInst translates the given x86 FCOS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFCOS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFCOS: not yet implemented")
}

// --- [ FDECSTP ] -------------------------------------------------------------

// emitInst translates the given x86 FDECSTP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFDECSTP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFDECSTP: not yet implemented")
}

// --- [ FDIV ] ----------------------------------------------------------------

// emitInst translates the given x86 FDIV instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFDIV(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFDIV: not yet implemented")
}

// --- [ FDIVP ] ---------------------------------------------------------------

// emitInst translates the given x86 FDIVP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFDIVP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFDIVP: not yet implemented")
}

// --- [ FDIVR ] ---------------------------------------------------------------

// emitInst translates the given x86 FDIVR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFDIVR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFDIVR: not yet implemented")
}

// --- [ FDIVRP ] --------------------------------------------------------------

// emitInst translates the given x86 FDIVRP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFDIVRP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFDIVRP: not yet implemented")
}

// --- [ FFREE ] ---------------------------------------------------------------

// emitInst translates the given x86 FFREE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFFREE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFFREE: not yet implemented")
}

// --- [ FFREEP ] --------------------------------------------------------------

// emitInst translates the given x86 FFREEP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFFREEP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFFREEP: not yet implemented")
}

// --- [ FIADD ] ---------------------------------------------------------------

// emitInst translates the given x86 FIADD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFIADD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFIADD: not yet implemented")
}

// --- [ FICOM ] ---------------------------------------------------------------

// emitInst translates the given x86 FICOM instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFICOM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFICOM: not yet implemented")
}

// --- [ FICOMP ] --------------------------------------------------------------

// emitInst translates the given x86 FICOMP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFICOMP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFICOMP: not yet implemented")
}

// --- [ FIDIV ] ---------------------------------------------------------------

// emitInst translates the given x86 FIDIV instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFIDIV(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFIDIV: not yet implemented")
}

// --- [ FIDIVR ] --------------------------------------------------------------

// emitInst translates the given x86 FIDIVR instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFIDIVR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFIDIVR: not yet implemented")
}

// --- [ FILD ] ----------------------------------------------------------------

// emitInst translates the given x86 FILD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFILD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFILD: not yet implemented")
}

// --- [ FIMUL ] ---------------------------------------------------------------

// emitInst translates the given x86 FIMUL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFIMUL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFIMUL: not yet implemented")
}

// --- [ FINCSTP ] -------------------------------------------------------------

// emitInst translates the given x86 FINCSTP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFINCSTP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFINCSTP: not yet implemented")
}

// --- [ FIST ] ----------------------------------------------------------------

// emitInst translates the given x86 FIST instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFIST(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFIST: not yet implemented")
}

// --- [ FISTP ] ---------------------------------------------------------------

// emitInst translates the given x86 FISTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFISTP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFISTP: not yet implemented")
}

// --- [ FISTTP ] --------------------------------------------------------------

// emitInst translates the given x86 FISTTP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFISTTP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFISTTP: not yet implemented")
}

// --- [ FISUB ] ---------------------------------------------------------------

// emitInst translates the given x86 FISUB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFISUB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFISUB: not yet implemented")
}

// --- [ FISUBR ] --------------------------------------------------------------

// emitInst translates the given x86 FISUBR instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFISUBR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFISUBR: not yet implemented")
}

// --- [ FLD ] -----------------------------------------------------------------

// emitInst translates the given x86 FLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFLD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFLD: not yet implemented")
}

// --- [ FLD1 ] ----------------------------------------------------------------

// emitInst translates the given x86 FLD1 instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFLD1(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFLD1: not yet implemented")
}

// --- [ FLDCW ] ---------------------------------------------------------------

// emitInst translates the given x86 FLDCW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFLDCW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFLDCW: not yet implemented")
}

// --- [ FLDENV ] --------------------------------------------------------------

// emitInst translates the given x86 FLDENV instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFLDENV(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFLDENV: not yet implemented")
}

// --- [ FLDL2E ] --------------------------------------------------------------

// emitInst translates the given x86 FLDL2E instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFLDL2E(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFLDL2E: not yet implemented")
}

// --- [ FLDL2T ] --------------------------------------------------------------

// emitInst translates the given x86 FLDL2T instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFLDL2T(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFLDL2T: not yet implemented")
}

// --- [ FLDLG2 ] --------------------------------------------------------------

// emitInst translates the given x86 FLDLG2 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFLDLG2(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFLDLG2: not yet implemented")
}

// --- [ FLDLN2 ] --------------------------------------------------------------

// emitInst translates the given x86 FLDLN2 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFLDLN2(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFLDLN2: not yet implemented")
}

// --- [ FLDPI ] ---------------------------------------------------------------

// emitInst translates the given x86 FLDPI instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFLDPI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFLDPI: not yet implemented")
}

// --- [ FLDZ ] ----------------------------------------------------------------

// emitInst translates the given x86 FLDZ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFLDZ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFLDZ: not yet implemented")
}

// --- [ FMUL ] ----------------------------------------------------------------

// emitInst translates the given x86 FMUL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFMUL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFMUL: not yet implemented")
}

// --- [ FMULP ] ---------------------------------------------------------------

// emitInst translates the given x86 FMULP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFMULP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFMULP: not yet implemented")
}

// --- [ FNCLEX ] --------------------------------------------------------------

// emitInst translates the given x86 FNCLEX instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFNCLEX(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFNCLEX: not yet implemented")
}

// --- [ FNINIT ] --------------------------------------------------------------

// emitInst translates the given x86 FNINIT instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFNINIT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFNINIT: not yet implemented")
}

// --- [ FNOP ] ----------------------------------------------------------------

// emitInst translates the given x86 FNOP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFNOP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFNOP: not yet implemented")
}

// --- [ FNSAVE ] --------------------------------------------------------------

// emitInst translates the given x86 FNSAVE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFNSAVE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFNSAVE: not yet implemented")
}

// --- [ FNSTCW ] --------------------------------------------------------------

// emitInst translates the given x86 FNSTCW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFNSTCW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFNSTCW: not yet implemented")
}

// --- [ FNSTENV ] -------------------------------------------------------------

// emitInst translates the given x86 FNSTENV instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFNSTENV(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFNSTENV: not yet implemented")
}

// --- [ FNSTSW ] --------------------------------------------------------------

// emitInst translates the given x86 FNSTSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFNSTSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFNSTSW: not yet implemented")
}

// --- [ FPATAN ] --------------------------------------------------------------

// emitInst translates the given x86 FPATAN instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFPATAN(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFPATAN: not yet implemented")
}

// --- [ FPREM ] ---------------------------------------------------------------

// emitInst translates the given x86 FPREM instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFPREM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFPREM: not yet implemented")
}

// --- [ FPREM1 ] --------------------------------------------------------------

// emitInst translates the given x86 FPREM1 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFPREM1(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFPREM1: not yet implemented")
}

// --- [ FPTAN ] ---------------------------------------------------------------

// emitInst translates the given x86 FPTAN instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFPTAN(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFPTAN: not yet implemented")
}

// --- [ FRNDINT ] -------------------------------------------------------------

// emitInst translates the given x86 FRNDINT instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFRNDINT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFRNDINT: not yet implemented")
}

// --- [ FRSTOR ] --------------------------------------------------------------

// emitInst translates the given x86 FRSTOR instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFRSTOR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFRSTOR: not yet implemented")
}

// --- [ FSCALE ] --------------------------------------------------------------

// emitInst translates the given x86 FSCALE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFSCALE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFSCALE: not yet implemented")
}

// --- [ FSIN ] ----------------------------------------------------------------

// emitInst translates the given x86 FSIN instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFSIN(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFSIN: not yet implemented")
}

// --- [ FSINCOS ] -------------------------------------------------------------

// emitInst translates the given x86 FSINCOS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFSINCOS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFSINCOS: not yet implemented")
}

// --- [ FSQRT ] ---------------------------------------------------------------

// emitInst translates the given x86 FSQRT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFSQRT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFSQRT: not yet implemented")
}

// --- [ FST ] -----------------------------------------------------------------

// emitInst translates the given x86 FST instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFST(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFST: not yet implemented")
}

// --- [ FSTP ] ----------------------------------------------------------------

// emitInst translates the given x86 FSTP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFSTP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFSTP: not yet implemented")
}

// --- [ FSUB ] ----------------------------------------------------------------

// emitInst translates the given x86 FSUB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFSUB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFSUB: not yet implemented")
}

// --- [ FSUBP ] ---------------------------------------------------------------

// emitInst translates the given x86 FSUBP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFSUBP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFSUBP: not yet implemented")
}

// --- [ FSUBR ] ---------------------------------------------------------------

// emitInst translates the given x86 FSUBR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFSUBR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFSUBR: not yet implemented")
}

// --- [ FSUBRP ] --------------------------------------------------------------

// emitInst translates the given x86 FSUBRP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFSUBRP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFSUBRP: not yet implemented")
}

// --- [ FTST ] ----------------------------------------------------------------

// emitInst translates the given x86 FTST instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFTST(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFTST: not yet implemented")
}

// --- [ FUCOM ] ---------------------------------------------------------------

// emitInst translates the given x86 FUCOM instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFUCOM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFUCOM: not yet implemented")
}

// --- [ FUCOMI ] --------------------------------------------------------------

// emitInst translates the given x86 FUCOMI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFUCOMI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFUCOMI: not yet implemented")
}

// --- [ FUCOMIP ] -------------------------------------------------------------

// emitInst translates the given x86 FUCOMIP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFUCOMIP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFUCOMIP: not yet implemented")
}

// --- [ FUCOMP ] --------------------------------------------------------------

// emitInst translates the given x86 FUCOMP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFUCOMP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFUCOMP: not yet implemented")
}

// --- [ FUCOMPP ] -------------------------------------------------------------

// emitInst translates the given x86 FUCOMPP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFUCOMPP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFUCOMPP: not yet implemented")
}

// --- [ FWAIT ] ---------------------------------------------------------------

// emitInst translates the given x86 FWAIT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFWAIT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFWAIT: not yet implemented")
}

// --- [ FXAM ] ----------------------------------------------------------------

// emitInst translates the given x86 FXAM instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFXAM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXAM: not yet implemented")
}

// --- [ FXCH ] ----------------------------------------------------------------

// emitInst translates the given x86 FXCH instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFXCH(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXCH: not yet implemented")
}

// --- [ FXRSTOR ] -------------------------------------------------------------

// emitInst translates the given x86 FXRSTOR instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFXRSTOR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXRSTOR: not yet implemented")
}

// --- [ FXRSTOR64 ] -----------------------------------------------------------

// emitInst translates the given x86 FXRSTOR64 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFXRSTOR64(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXRSTOR64: not yet implemented")
}

// --- [ FXSAVE ] --------------------------------------------------------------

// emitInst translates the given x86 FXSAVE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFXSAVE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXSAVE: not yet implemented")
}

// --- [ FXSAVE64 ] ------------------------------------------------------------

// emitInst translates the given x86 FXSAVE64 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFXSAVE64(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXSAVE64: not yet implemented")
}

// --- [ FXTRACT ] -------------------------------------------------------------

// emitInst translates the given x86 FXTRACT instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFXTRACT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXTRACT: not yet implemented")
}

// --- [ FYL2X ] ---------------------------------------------------------------

// emitInst translates the given x86 FYL2X instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstFYL2X(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFYL2X: not yet implemented")
}

// --- [ FYL2XP1 ] -------------------------------------------------------------

// emitInst translates the given x86 FYL2XP1 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstFYL2XP1(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFYL2XP1: not yet implemented")
}

// --- [ HADDPD ] --------------------------------------------------------------

// emitInst translates the given x86 HADDPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstHADDPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstHADDPD: not yet implemented")
}

// --- [ HADDPS ] --------------------------------------------------------------

// emitInst translates the given x86 HADDPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstHADDPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstHADDPS: not yet implemented")
}

// --- [ HLT ] -----------------------------------------------------------------

// emitInst translates the given x86 HLT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstHLT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstHLT: not yet implemented")
}

// --- [ HSUBPD ] --------------------------------------------------------------

// emitInst translates the given x86 HSUBPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstHSUBPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstHSUBPD: not yet implemented")
}

// --- [ HSUBPS ] --------------------------------------------------------------

// emitInst translates the given x86 HSUBPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstHSUBPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstHSUBPS: not yet implemented")
}

// --- [ ICEBP ] ---------------------------------------------------------------

// emitInst translates the given x86 ICEBP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstICEBP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstICEBP: not yet implemented")
}

// --- [ IDIV ] ----------------------------------------------------------------

// emitInst translates the given x86 IDIV instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstIDIV(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstIDIV: not yet implemented")
}

// --- [ IMUL ] ----------------------------------------------------------------

// emitInst translates the given x86 IMUL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstIMUL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstIMUL: not yet implemented")
}

// --- [ IN ] ------------------------------------------------------------------

// emitInst translates the given x86 IN instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstIN(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstIN: not yet implemented")
}

// --- [ INC ] -----------------------------------------------------------------

// emitInst translates the given x86 INC instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstINC(inst *Inst) error {
	x := f.useArg(inst.Arg(0))
	one := constant.NewInt(1, types.I32)
	result := f.cur.NewAdd(x, one)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ INSB ] ----------------------------------------------------------------

// emitInst translates the given x86 INSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstINSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINSB: not yet implemented")
}

// --- [ INSD ] ----------------------------------------------------------------

// emitInst translates the given x86 INSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstINSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINSD: not yet implemented")
}

// --- [ INSERTPS ] ------------------------------------------------------------

// emitInst translates the given x86 INSERTPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstINSERTPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINSERTPS: not yet implemented")
}

// --- [ INSW ] ----------------------------------------------------------------

// emitInst translates the given x86 INSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstINSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINSW: not yet implemented")
}

// --- [ INT ] -----------------------------------------------------------------

// emitInst translates the given x86 INT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstINT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINT: not yet implemented")
}

// --- [ INTO ] ----------------------------------------------------------------

// emitInst translates the given x86 INTO instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstINTO(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINTO: not yet implemented")
}

// --- [ INVD ] ----------------------------------------------------------------

// emitInst translates the given x86 INVD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstINVD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINVD: not yet implemented")
}

// --- [ INVLPG ] --------------------------------------------------------------

// emitInst translates the given x86 INVLPG instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstINVLPG(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINVLPG: not yet implemented")
}

// --- [ INVPCID ] -------------------------------------------------------------

// emitInst translates the given x86 INVPCID instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstINVPCID(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINVPCID: not yet implemented")
}

// --- [ IRET ] ----------------------------------------------------------------

// emitInst translates the given x86 IRET instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstIRET(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstIRET: not yet implemented")
}

// --- [ IRETD ] ---------------------------------------------------------------

// emitInst translates the given x86 IRETD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstIRETD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstIRETD: not yet implemented")
}

// --- [ IRETQ ] ---------------------------------------------------------------

// emitInst translates the given x86 IRETQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstIRETQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstIRETQ: not yet implemented")
}

// --- [ JA ] ------------------------------------------------------------------

// emitInst translates the given x86 JA instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstJA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJA: not yet implemented")
}

// --- [ JAE ] -----------------------------------------------------------------

// emitInst translates the given x86 JAE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJAE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJAE: not yet implemented")
}

// --- [ JB ] ------------------------------------------------------------------

// emitInst translates the given x86 JB instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstJB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJB: not yet implemented")
}

// --- [ JBE ] -----------------------------------------------------------------

// emitInst translates the given x86 JBE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJBE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJBE: not yet implemented")
}

// --- [ JCXZ ] ----------------------------------------------------------------

// emitInst translates the given x86 JCXZ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJCXZ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJCXZ: not yet implemented")
}

// --- [ JE ] ------------------------------------------------------------------

// emitInst translates the given x86 JE instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstJE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJE: not yet implemented")
}

// --- [ JECXZ ] ---------------------------------------------------------------

// emitInst translates the given x86 JECXZ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJECXZ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJECXZ: not yet implemented")
}

// --- [ JG ] ------------------------------------------------------------------

// emitInst translates the given x86 JG instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstJG(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJG: not yet implemented")
}

// --- [ JGE ] -----------------------------------------------------------------

// emitInst translates the given x86 JGE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJGE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJGE: not yet implemented")
}

// --- [ JL ] ------------------------------------------------------------------

// emitInst translates the given x86 JL instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstJL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJL: not yet implemented")
}

// --- [ JLE ] -----------------------------------------------------------------

// emitInst translates the given x86 JLE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJLE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJLE: not yet implemented")
}

// --- [ JMP ] -----------------------------------------------------------------

// emitInst translates the given x86 JMP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJMP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJMP: not yet implemented")
}

// --- [ JNE ] -----------------------------------------------------------------

// emitInst translates the given x86 JNE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJNE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJNE: not yet implemented")
}

// --- [ JNO ] -----------------------------------------------------------------

// emitInst translates the given x86 JNO instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJNO(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJNO: not yet implemented")
}

// --- [ JNP ] -----------------------------------------------------------------

// emitInst translates the given x86 JNP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJNP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJNP: not yet implemented")
}

// --- [ JNS ] -----------------------------------------------------------------

// emitInst translates the given x86 JNS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJNS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJNS: not yet implemented")
}

// --- [ JO ] ------------------------------------------------------------------

// emitInst translates the given x86 JO instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstJO(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJO: not yet implemented")
}

// --- [ JP ] ------------------------------------------------------------------

// emitInst translates the given x86 JP instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstJP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJP: not yet implemented")
}

// --- [ JRCXZ ] ---------------------------------------------------------------

// emitInst translates the given x86 JRCXZ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstJRCXZ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJRCXZ: not yet implemented")
}

// --- [ JS ] ------------------------------------------------------------------

// emitInst translates the given x86 JS instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstJS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstJS: not yet implemented")
}

// --- [ LAHF ] ----------------------------------------------------------------

// emitInst translates the given x86 LAHF instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLAHF(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLAHF: not yet implemented")
}

// --- [ LAR ] -----------------------------------------------------------------

// emitInst translates the given x86 LAR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLAR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLAR: not yet implemented")
}

// --- [ LCALL ] ---------------------------------------------------------------

// emitInst translates the given x86 LCALL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLCALL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLCALL: not yet implemented")
}

// --- [ LDDQU ] ---------------------------------------------------------------

// emitInst translates the given x86 LDDQU instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLDDQU(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLDDQU: not yet implemented")
}

// --- [ LDMXCSR ] -------------------------------------------------------------

// emitInst translates the given x86 LDMXCSR instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstLDMXCSR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLDMXCSR: not yet implemented")
}

// --- [ LDS ] -----------------------------------------------------------------

// emitInst translates the given x86 LDS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLDS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLDS: not yet implemented")
}

// --- [ LEA ] -----------------------------------------------------------------

// emitInst translates the given x86 LEA instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLEA(inst *Inst) error {
	y := f.mem(inst.Mem(1))
	f.defArg(inst.Arg(0), y)
	return nil
}

// --- [ LEAVE ] ---------------------------------------------------------------

// emitInst translates the given x86 LEAVE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLEAVE(inst *Inst) error {
	// Pseudo-instruction for:
	//
	//    mov esp, ebp
	//    pop ebp

	//    mov esp, ebp
	ebp := f.useReg(EBP)
	f.defReg(ESP, ebp)
	// TODO: Explicitly setting espDisp to -4 should not be needed once espDisp
	// is stored per basic block and its changes tracked through the CFG. Remove
	// when handling of espDisp has matured.
	f.espDisp = -4

	//    pop ebp
	ebp = f.pop()
	f.defReg(EBP, ebp)

	return nil
}

// --- [ LES ] -----------------------------------------------------------------

// emitInst translates the given x86 LES instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLES(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLES: not yet implemented")
}

// --- [ LFENCE ] --------------------------------------------------------------

// emitInst translates the given x86 LFENCE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstLFENCE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLFENCE: not yet implemented")
}

// --- [ LFS ] -----------------------------------------------------------------

// emitInst translates the given x86 LFS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLFS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLFS: not yet implemented")
}

// --- [ LGDT ] ----------------------------------------------------------------

// emitInst translates the given x86 LGDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLGDT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLGDT: not yet implemented")
}

// --- [ LGS ] -----------------------------------------------------------------

// emitInst translates the given x86 LGS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLGS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLGS: not yet implemented")
}

// --- [ LIDT ] ----------------------------------------------------------------

// emitInst translates the given x86 LIDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLIDT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLIDT: not yet implemented")
}

// --- [ LJMP ] ----------------------------------------------------------------

// emitInst translates the given x86 LJMP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLJMP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLJMP: not yet implemented")
}

// --- [ LLDT ] ----------------------------------------------------------------

// emitInst translates the given x86 LLDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLLDT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLLDT: not yet implemented")
}

// --- [ LMSW ] ----------------------------------------------------------------

// emitInst translates the given x86 LMSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLMSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLMSW: not yet implemented")
}

// --- [ LODSB ] ---------------------------------------------------------------

// emitInst translates the given x86 LODSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLODSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLODSB: not yet implemented")
}

// --- [ LODSD ] ---------------------------------------------------------------

// emitInst translates the given x86 LODSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLODSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLODSD: not yet implemented")
}

// --- [ LODSQ ] ---------------------------------------------------------------

// emitInst translates the given x86 LODSQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLODSQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLODSQ: not yet implemented")
}

// --- [ LODSW ] ---------------------------------------------------------------

// emitInst translates the given x86 LODSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLODSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLODSW: not yet implemented")
}

// --- [ LOOP ] ----------------------------------------------------------------

// emitInst translates the given x86 LOOP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLOOP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLOOP: not yet implemented")
}

// --- [ LOOPE ] ---------------------------------------------------------------

// emitInst translates the given x86 LOOPE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLOOPE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLOOPE: not yet implemented")
}

// --- [ LOOPNE ] --------------------------------------------------------------

// emitInst translates the given x86 LOOPNE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstLOOPNE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLOOPNE: not yet implemented")
}

// --- [ LRET ] ----------------------------------------------------------------

// emitInst translates the given x86 LRET instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLRET(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLRET: not yet implemented")
}

// --- [ LSL ] -----------------------------------------------------------------

// emitInst translates the given x86 LSL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLSL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLSL: not yet implemented")
}

// --- [ LSS ] -----------------------------------------------------------------

// emitInst translates the given x86 LSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLSS: not yet implemented")
}

// --- [ LTR ] -----------------------------------------------------------------

// emitInst translates the given x86 LTR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLTR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLTR: not yet implemented")
}

// --- [ LZCNT ] ---------------------------------------------------------------

// emitInst translates the given x86 LZCNT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstLZCNT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLZCNT: not yet implemented")
}

// --- [ MASKMOVDQU ] ----------------------------------------------------------

// emitInst translates the given x86 MASKMOVDQU instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMASKMOVDQU(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMASKMOVDQU: not yet implemented")
}

// --- [ MASKMOVQ ] ------------------------------------------------------------

// emitInst translates the given x86 MASKMOVQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMASKMOVQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMASKMOVQ: not yet implemented")
}

// --- [ MAXPD ] ---------------------------------------------------------------

// emitInst translates the given x86 MAXPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMAXPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMAXPD: not yet implemented")
}

// --- [ MAXPS ] ---------------------------------------------------------------

// emitInst translates the given x86 MAXPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMAXPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMAXPS: not yet implemented")
}

// --- [ MAXSD ] ---------------------------------------------------------------

// emitInst translates the given x86 MAXSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMAXSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMAXSD: not yet implemented")
}

// --- [ MAXSS ] ---------------------------------------------------------------

// emitInst translates the given x86 MAXSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMAXSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMAXSS: not yet implemented")
}

// --- [ MFENCE ] --------------------------------------------------------------

// emitInst translates the given x86 MFENCE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMFENCE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMFENCE: not yet implemented")
}

// --- [ MINPD ] ---------------------------------------------------------------

// emitInst translates the given x86 MINPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMINPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMINPD: not yet implemented")
}

// --- [ MINPS ] ---------------------------------------------------------------

// emitInst translates the given x86 MINPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMINPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMINPS: not yet implemented")
}

// --- [ MINSD ] ---------------------------------------------------------------

// emitInst translates the given x86 MINSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMINSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMINSD: not yet implemented")
}

// --- [ MINSS ] ---------------------------------------------------------------

// emitInst translates the given x86 MINSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMINSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMINSS: not yet implemented")
}

// --- [ MONITOR ] -------------------------------------------------------------

// emitInst translates the given x86 MONITOR instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMONITOR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMONITOR: not yet implemented")
}

// --- [ MOV ] -----------------------------------------------------------------

// emitInst translates the given x86 MOV instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOV(inst *Inst) error {
	src := f.useArg(inst.Arg(1))
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MOVAPD ] --------------------------------------------------------------

// emitInst translates the given x86 MOVAPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVAPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVAPD: not yet implemented")
}

// --- [ MOVAPS ] --------------------------------------------------------------

// emitInst translates the given x86 MOVAPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVAPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVAPS: not yet implemented")
}

// --- [ MOVBE ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVBE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOVBE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVBE: not yet implemented")
}

// --- [ MOVD ] ----------------------------------------------------------------

// emitInst translates the given x86 MOVD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOVD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVD: not yet implemented")
}

// --- [ MOVDDUP ] -------------------------------------------------------------

// emitInst translates the given x86 MOVDDUP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVDDUP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVDDUP: not yet implemented")
}

// --- [ MOVDQ2Q ] -------------------------------------------------------------

// emitInst translates the given x86 MOVDQ2Q instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVDQ2Q(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVDQ2Q: not yet implemented")
}

// --- [ MOVDQA ] --------------------------------------------------------------

// emitInst translates the given x86 MOVDQA instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVDQA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVDQA: not yet implemented")
}

// --- [ MOVDQU ] --------------------------------------------------------------

// emitInst translates the given x86 MOVDQU instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVDQU(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVDQU: not yet implemented")
}

// --- [ MOVHLPS ] -------------------------------------------------------------

// emitInst translates the given x86 MOVHLPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVHLPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVHLPS: not yet implemented")
}

// --- [ MOVHPD ] --------------------------------------------------------------

// emitInst translates the given x86 MOVHPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVHPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVHPD: not yet implemented")
}

// --- [ MOVHPS ] --------------------------------------------------------------

// emitInst translates the given x86 MOVHPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVHPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVHPS: not yet implemented")
}

// --- [ MOVLHPS ] -------------------------------------------------------------

// emitInst translates the given x86 MOVLHPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVLHPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVLHPS: not yet implemented")
}

// --- [ MOVLPD ] --------------------------------------------------------------

// emitInst translates the given x86 MOVLPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVLPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVLPD: not yet implemented")
}

// --- [ MOVLPS ] --------------------------------------------------------------

// emitInst translates the given x86 MOVLPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVLPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVLPS: not yet implemented")
}

// --- [ MOVMSKPD ] ------------------------------------------------------------

// emitInst translates the given x86 MOVMSKPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVMSKPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVMSKPD: not yet implemented")
}

// --- [ MOVMSKPS ] ------------------------------------------------------------

// emitInst translates the given x86 MOVMSKPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVMSKPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVMSKPS: not yet implemented")
}

// --- [ MOVNTDQ ] -------------------------------------------------------------

// emitInst translates the given x86 MOVNTDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVNTDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTDQ: not yet implemented")
}

// --- [ MOVNTDQA ] ------------------------------------------------------------

// emitInst translates the given x86 MOVNTDQA instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVNTDQA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTDQA: not yet implemented")
}

// --- [ MOVNTI ] --------------------------------------------------------------

// emitInst translates the given x86 MOVNTI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVNTI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTI: not yet implemented")
}

// --- [ MOVNTPD ] -------------------------------------------------------------

// emitInst translates the given x86 MOVNTPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVNTPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTPD: not yet implemented")
}

// --- [ MOVNTPS ] -------------------------------------------------------------

// emitInst translates the given x86 MOVNTPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVNTPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTPS: not yet implemented")
}

// --- [ MOVNTQ ] --------------------------------------------------------------

// emitInst translates the given x86 MOVNTQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVNTQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTQ: not yet implemented")
}

// --- [ MOVNTSD ] -------------------------------------------------------------

// emitInst translates the given x86 MOVNTSD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVNTSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTSD: not yet implemented")
}

// --- [ MOVNTSS ] -------------------------------------------------------------

// emitInst translates the given x86 MOVNTSS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVNTSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTSS: not yet implemented")
}

// --- [ MOVQ ] ----------------------------------------------------------------

// emitInst translates the given x86 MOVQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOVQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVQ: not yet implemented")
}

// --- [ MOVQ2DQ ] -------------------------------------------------------------

// emitInst translates the given x86 MOVQ2DQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVQ2DQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVQ2DQ: not yet implemented")
}

// --- [ MOVSB ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOVSB(inst *Inst) error {
	src := f.useArgElem(inst.Arg(1), types.I8)
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MOVSD ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOVSD(inst *Inst) error {
	src := f.useArgElem(inst.Arg(1), types.I32)
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MOVSD_XMM ] -----------------------------------------------------------

// emitInst translates the given x86 MOVSD_XMM instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVSD_XMM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSD_XMM: not yet implemented")
}

// --- [ MOVSHDUP ] ------------------------------------------------------------

// emitInst translates the given x86 MOVSHDUP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVSHDUP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSHDUP: not yet implemented")
}

// --- [ MOVSLDUP ] ------------------------------------------------------------

// emitInst translates the given x86 MOVSLDUP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVSLDUP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSLDUP: not yet implemented")
}

// --- [ MOVSQ ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOVSQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSQ: not yet implemented")
}

// --- [ MOVSS ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOVSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSS: not yet implemented")
}

// --- [ MOVSW ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOVSW(inst *Inst) error {
	src := f.useArgElem(inst.Arg(1), types.I16)
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MOVSX ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSX instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOVSX(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSX: not yet implemented")
}

// --- [ MOVSXD ] --------------------------------------------------------------

// emitInst translates the given x86 MOVSXD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVSXD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSXD: not yet implemented")
}

// --- [ MOVUPD ] --------------------------------------------------------------

// emitInst translates the given x86 MOVUPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVUPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVUPD: not yet implemented")
}

// --- [ MOVUPS ] --------------------------------------------------------------

// emitInst translates the given x86 MOVUPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMOVUPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVUPS: not yet implemented")
}

// --- [ MOVZX ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVZX instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMOVZX(inst *Inst) error {
	size := inst.MemBytes * 8
	elem := types.NewInt(size)
	src := f.useArgElem(inst.Arg(1), elem)
	// TODO: Handle dst type dynamically.
	src = f.cur.NewZExt(src, types.I32)
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MPSADBW ] -------------------------------------------------------------

// emitInst translates the given x86 MPSADBW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstMPSADBW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMPSADBW: not yet implemented")
}

// --- [ MUL ] -----------------------------------------------------------------

// emitInst translates the given x86 MUL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMUL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMUL: not yet implemented")
}

// --- [ MULPD ] ---------------------------------------------------------------

// emitInst translates the given x86 MULPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMULPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMULPD: not yet implemented")
}

// --- [ MULPS ] ---------------------------------------------------------------

// emitInst translates the given x86 MULPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMULPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMULPS: not yet implemented")
}

// --- [ MULSD ] ---------------------------------------------------------------

// emitInst translates the given x86 MULSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMULSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMULSD: not yet implemented")
}

// --- [ MULSS ] ---------------------------------------------------------------

// emitInst translates the given x86 MULSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMULSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMULSS: not yet implemented")
}

// --- [ MWAIT ] ---------------------------------------------------------------

// emitInst translates the given x86 MWAIT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstMWAIT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMWAIT: not yet implemented")
}

// --- [ NEG ] -----------------------------------------------------------------

// emitInst translates the given x86 NEG instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstNEG(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstNEG: not yet implemented")
}

// --- [ NOP ] -----------------------------------------------------------------

// emitInst translates the given x86 NOP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstNOP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstNOP: not yet implemented")
}

// --- [ NOT ] -----------------------------------------------------------------

// emitInst translates the given x86 NOT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstNOT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstNOT: not yet implemented")
}

// --- [ OR ] ------------------------------------------------------------------

// emitInst translates the given x86 OR instruction to LLVM IR, emitting code to
// f.
func (f *Func) emitInstOR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstOR: not yet implemented")
}

// --- [ ORPD ] ----------------------------------------------------------------

// emitInst translates the given x86 ORPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstORPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstORPD: not yet implemented")
}

// --- [ ORPS ] ----------------------------------------------------------------

// emitInst translates the given x86 ORPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstORPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstORPS: not yet implemented")
}

// --- [ OUT ] -----------------------------------------------------------------

// emitInst translates the given x86 OUT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstOUT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstOUT: not yet implemented")
}

// --- [ OUTSB ] ---------------------------------------------------------------

// emitInst translates the given x86 OUTSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstOUTSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstOUTSB: not yet implemented")
}

// --- [ OUTSD ] ---------------------------------------------------------------

// emitInst translates the given x86 OUTSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstOUTSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstOUTSD: not yet implemented")
}

// --- [ OUTSW ] ---------------------------------------------------------------

// emitInst translates the given x86 OUTSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstOUTSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstOUTSW: not yet implemented")
}

// --- [ PABSB ] ---------------------------------------------------------------

// emitInst translates the given x86 PABSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPABSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPABSB: not yet implemented")
}

// --- [ PABSD ] ---------------------------------------------------------------

// emitInst translates the given x86 PABSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPABSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPABSD: not yet implemented")
}

// --- [ PABSW ] ---------------------------------------------------------------

// emitInst translates the given x86 PABSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPABSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPABSW: not yet implemented")
}

// --- [ PACKSSDW ] ------------------------------------------------------------

// emitInst translates the given x86 PACKSSDW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPACKSSDW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPACKSSDW: not yet implemented")
}

// --- [ PACKSSWB ] ------------------------------------------------------------

// emitInst translates the given x86 PACKSSWB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPACKSSWB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPACKSSWB: not yet implemented")
}

// --- [ PACKUSDW ] ------------------------------------------------------------

// emitInst translates the given x86 PACKUSDW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPACKUSDW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPACKUSDW: not yet implemented")
}

// --- [ PACKUSWB ] ------------------------------------------------------------

// emitInst translates the given x86 PACKUSWB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPACKUSWB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPACKUSWB: not yet implemented")
}

// --- [ PADDB ] ---------------------------------------------------------------

// emitInst translates the given x86 PADDB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPADDB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDB: not yet implemented")
}

// --- [ PADDD ] ---------------------------------------------------------------

// emitInst translates the given x86 PADDD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPADDD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDD: not yet implemented")
}

// --- [ PADDQ ] ---------------------------------------------------------------

// emitInst translates the given x86 PADDQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPADDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDQ: not yet implemented")
}

// --- [ PADDSB ] --------------------------------------------------------------

// emitInst translates the given x86 PADDSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPADDSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDSB: not yet implemented")
}

// --- [ PADDSW ] --------------------------------------------------------------

// emitInst translates the given x86 PADDSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPADDSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDSW: not yet implemented")
}

// --- [ PADDUSB ] -------------------------------------------------------------

// emitInst translates the given x86 PADDUSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPADDUSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDUSB: not yet implemented")
}

// --- [ PADDUSW ] -------------------------------------------------------------

// emitInst translates the given x86 PADDUSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPADDUSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDUSW: not yet implemented")
}

// --- [ PADDW ] ---------------------------------------------------------------

// emitInst translates the given x86 PADDW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPADDW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDW: not yet implemented")
}

// --- [ PALIGNR ] -------------------------------------------------------------

// emitInst translates the given x86 PALIGNR instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPALIGNR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPALIGNR: not yet implemented")
}

// --- [ PAND ] ----------------------------------------------------------------

// emitInst translates the given x86 PAND instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPAND(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPAND: not yet implemented")
}

// --- [ PANDN ] ---------------------------------------------------------------

// emitInst translates the given x86 PANDN instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPANDN(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPANDN: not yet implemented")
}

// --- [ PAUSE ] ---------------------------------------------------------------

// emitInst translates the given x86 PAUSE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPAUSE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPAUSE: not yet implemented")
}

// --- [ PAVGB ] ---------------------------------------------------------------

// emitInst translates the given x86 PAVGB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPAVGB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPAVGB: not yet implemented")
}

// --- [ PAVGW ] ---------------------------------------------------------------

// emitInst translates the given x86 PAVGW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPAVGW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPAVGW: not yet implemented")
}

// --- [ PBLENDVB ] ------------------------------------------------------------

// emitInst translates the given x86 PBLENDVB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPBLENDVB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPBLENDVB: not yet implemented")
}

// --- [ PBLENDW ] -------------------------------------------------------------

// emitInst translates the given x86 PBLENDW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPBLENDW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPBLENDW: not yet implemented")
}

// --- [ PCLMULQDQ ] -----------------------------------------------------------

// emitInst translates the given x86 PCLMULQDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCLMULQDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCLMULQDQ: not yet implemented")
}

// --- [ PCMPEQB ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPEQB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPEQB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPEQB: not yet implemented")
}

// --- [ PCMPEQD ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPEQD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPEQD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPEQD: not yet implemented")
}

// --- [ PCMPEQQ ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPEQQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPEQQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPEQQ: not yet implemented")
}

// --- [ PCMPEQW ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPEQW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPEQW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPEQW: not yet implemented")
}

// --- [ PCMPESTRI ] -----------------------------------------------------------

// emitInst translates the given x86 PCMPESTRI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPESTRI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPESTRI: not yet implemented")
}

// --- [ PCMPESTRM ] -----------------------------------------------------------

// emitInst translates the given x86 PCMPESTRM instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPESTRM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPESTRM: not yet implemented")
}

// --- [ PCMPGTB ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPGTB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPGTB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPGTB: not yet implemented")
}

// --- [ PCMPGTD ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPGTD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPGTD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPGTD: not yet implemented")
}

// --- [ PCMPGTQ ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPGTQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPGTQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPGTQ: not yet implemented")
}

// --- [ PCMPGTW ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPGTW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPGTW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPGTW: not yet implemented")
}

// --- [ PCMPISTRI ] -----------------------------------------------------------

// emitInst translates the given x86 PCMPISTRI instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPISTRI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPISTRI: not yet implemented")
}

// --- [ PCMPISTRM ] -----------------------------------------------------------

// emitInst translates the given x86 PCMPISTRM instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPCMPISTRM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPISTRM: not yet implemented")
}

// --- [ PEXTRB ] --------------------------------------------------------------

// emitInst translates the given x86 PEXTRB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPEXTRB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPEXTRB: not yet implemented")
}

// --- [ PEXTRD ] --------------------------------------------------------------

// emitInst translates the given x86 PEXTRD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPEXTRD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPEXTRD: not yet implemented")
}

// --- [ PEXTRQ ] --------------------------------------------------------------

// emitInst translates the given x86 PEXTRQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPEXTRQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPEXTRQ: not yet implemented")
}

// --- [ PEXTRW ] --------------------------------------------------------------

// emitInst translates the given x86 PEXTRW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPEXTRW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPEXTRW: not yet implemented")
}

// --- [ PHADDD ] --------------------------------------------------------------

// emitInst translates the given x86 PHADDD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPHADDD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHADDD: not yet implemented")
}

// --- [ PHADDSW ] -------------------------------------------------------------

// emitInst translates the given x86 PHADDSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPHADDSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHADDSW: not yet implemented")
}

// --- [ PHADDW ] --------------------------------------------------------------

// emitInst translates the given x86 PHADDW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPHADDW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHADDW: not yet implemented")
}

// --- [ PHMINPOSUW ] ----------------------------------------------------------

// emitInst translates the given x86 PHMINPOSUW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPHMINPOSUW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHMINPOSUW: not yet implemented")
}

// --- [ PHSUBD ] --------------------------------------------------------------

// emitInst translates the given x86 PHSUBD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPHSUBD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHSUBD: not yet implemented")
}

// --- [ PHSUBSW ] -------------------------------------------------------------

// emitInst translates the given x86 PHSUBSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPHSUBSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHSUBSW: not yet implemented")
}

// --- [ PHSUBW ] --------------------------------------------------------------

// emitInst translates the given x86 PHSUBW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPHSUBW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHSUBW: not yet implemented")
}

// --- [ PINSRB ] --------------------------------------------------------------

// emitInst translates the given x86 PINSRB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPINSRB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPINSRB: not yet implemented")
}

// --- [ PINSRD ] --------------------------------------------------------------

// emitInst translates the given x86 PINSRD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPINSRD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPINSRD: not yet implemented")
}

// --- [ PINSRQ ] --------------------------------------------------------------

// emitInst translates the given x86 PINSRQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPINSRQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPINSRQ: not yet implemented")
}

// --- [ PINSRW ] --------------------------------------------------------------

// emitInst translates the given x86 PINSRW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPINSRW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPINSRW: not yet implemented")
}

// --- [ PMADDUBSW ] -----------------------------------------------------------

// emitInst translates the given x86 PMADDUBSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMADDUBSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMADDUBSW: not yet implemented")
}

// --- [ PMADDWD ] -------------------------------------------------------------

// emitInst translates the given x86 PMADDWD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMADDWD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMADDWD: not yet implemented")
}

// --- [ PMAXSB ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMAXSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXSB: not yet implemented")
}

// --- [ PMAXSD ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXSD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMAXSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXSD: not yet implemented")
}

// --- [ PMAXSW ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMAXSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXSW: not yet implemented")
}

// --- [ PMAXUB ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXUB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMAXUB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXUB: not yet implemented")
}

// --- [ PMAXUD ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXUD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMAXUD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXUD: not yet implemented")
}

// --- [ PMAXUW ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXUW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMAXUW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXUW: not yet implemented")
}

// --- [ PMINSB ] --------------------------------------------------------------

// emitInst translates the given x86 PMINSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMINSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINSB: not yet implemented")
}

// --- [ PMINSD ] --------------------------------------------------------------

// emitInst translates the given x86 PMINSD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMINSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINSD: not yet implemented")
}

// --- [ PMINSW ] --------------------------------------------------------------

// emitInst translates the given x86 PMINSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMINSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINSW: not yet implemented")
}

// --- [ PMINUB ] --------------------------------------------------------------

// emitInst translates the given x86 PMINUB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMINUB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINUB: not yet implemented")
}

// --- [ PMINUD ] --------------------------------------------------------------

// emitInst translates the given x86 PMINUD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMINUD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINUD: not yet implemented")
}

// --- [ PMINUW ] --------------------------------------------------------------

// emitInst translates the given x86 PMINUW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMINUW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINUW: not yet implemented")
}

// --- [ PMOVMSKB ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVMSKB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVMSKB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVMSKB: not yet implemented")
}

// --- [ PMOVSXBD ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXBD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVSXBD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXBD: not yet implemented")
}

// --- [ PMOVSXBQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXBQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVSXBQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXBQ: not yet implemented")
}

// --- [ PMOVSXBW ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXBW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVSXBW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXBW: not yet implemented")
}

// --- [ PMOVSXDQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVSXDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXDQ: not yet implemented")
}

// --- [ PMOVSXWD ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXWD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVSXWD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXWD: not yet implemented")
}

// --- [ PMOVSXWQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXWQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVSXWQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXWQ: not yet implemented")
}

// --- [ PMOVZXBD ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXBD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVZXBD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXBD: not yet implemented")
}

// --- [ PMOVZXBQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXBQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVZXBQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXBQ: not yet implemented")
}

// --- [ PMOVZXBW ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXBW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVZXBW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXBW: not yet implemented")
}

// --- [ PMOVZXDQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVZXDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXDQ: not yet implemented")
}

// --- [ PMOVZXWD ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXWD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVZXWD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXWD: not yet implemented")
}

// --- [ PMOVZXWQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXWQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMOVZXWQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXWQ: not yet implemented")
}

// --- [ PMULDQ ] --------------------------------------------------------------

// emitInst translates the given x86 PMULDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMULDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULDQ: not yet implemented")
}

// --- [ PMULHRSW ] ------------------------------------------------------------

// emitInst translates the given x86 PMULHRSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMULHRSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULHRSW: not yet implemented")
}

// --- [ PMULHUW ] -------------------------------------------------------------

// emitInst translates the given x86 PMULHUW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMULHUW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULHUW: not yet implemented")
}

// --- [ PMULHW ] --------------------------------------------------------------

// emitInst translates the given x86 PMULHW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMULHW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULHW: not yet implemented")
}

// --- [ PMULLD ] --------------------------------------------------------------

// emitInst translates the given x86 PMULLD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMULLD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULLD: not yet implemented")
}

// --- [ PMULLW ] --------------------------------------------------------------

// emitInst translates the given x86 PMULLW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMULLW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULLW: not yet implemented")
}

// --- [ PMULUDQ ] -------------------------------------------------------------

// emitInst translates the given x86 PMULUDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPMULUDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULUDQ: not yet implemented")
}

// --- [ POP ] -----------------------------------------------------------------

// emitInst translates the given x86 POP instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPOP(inst *Inst) error {
	v := f.pop()
	f.defArg(inst.Arg(0), v)
	return nil
}

// pop pops a value from the top of the stack of the function, emitting code to
// f.
func (f *Func) pop() value.Value {
	m := x86asm.Mem{
		Base: x86asm.ESP,
	}
	mem := NewMem(m, nil)
	v := f.useMem(mem)
	f.espDisp += 4
	return v
}

// --- [ POPA ] ----------------------------------------------------------------

// emitInst translates the given x86 POPA instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPOPA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPA: not yet implemented")
}

// --- [ POPAD ] ---------------------------------------------------------------

// emitInst translates the given x86 POPAD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPOPAD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPAD: not yet implemented")
}

// --- [ POPCNT ] --------------------------------------------------------------

// emitInst translates the given x86 POPCNT instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPOPCNT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPCNT: not yet implemented")
}

// --- [ POPF ] ----------------------------------------------------------------

// emitInst translates the given x86 POPF instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPOPF(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPF: not yet implemented")
}

// --- [ POPFD ] ---------------------------------------------------------------

// emitInst translates the given x86 POPFD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPOPFD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPFD: not yet implemented")
}

// --- [ POPFQ ] ---------------------------------------------------------------

// emitInst translates the given x86 POPFQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPOPFQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPFQ: not yet implemented")
}

// --- [ POR ] -----------------------------------------------------------------

// emitInst translates the given x86 POR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPOR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOR: not yet implemented")
}

// --- [ PREFETCHNTA ] ---------------------------------------------------------

// emitInst translates the given x86 PREFETCHNTA instruction to LLVM IR,
// emitting code to f.
func (f *Func) emitInstPREFETCHNTA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPREFETCHNTA: not yet implemented")
}

// --- [ PREFETCHT0 ] ----------------------------------------------------------

// emitInst translates the given x86 PREFETCHT0 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPREFETCHT0(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPREFETCHT0: not yet implemented")
}

// --- [ PREFETCHT1 ] ----------------------------------------------------------

// emitInst translates the given x86 PREFETCHT1 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPREFETCHT1(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPREFETCHT1: not yet implemented")
}

// --- [ PREFETCHT2 ] ----------------------------------------------------------

// emitInst translates the given x86 PREFETCHT2 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPREFETCHT2(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPREFETCHT2: not yet implemented")
}

// --- [ PREFETCHW ] -----------------------------------------------------------

// emitInst translates the given x86 PREFETCHW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPREFETCHW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPREFETCHW: not yet implemented")
}

// --- [ PSADBW ] --------------------------------------------------------------

// emitInst translates the given x86 PSADBW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSADBW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSADBW: not yet implemented")
}

// --- [ PSHUFB ] --------------------------------------------------------------

// emitInst translates the given x86 PSHUFB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSHUFB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSHUFB: not yet implemented")
}

// --- [ PSHUFD ] --------------------------------------------------------------

// emitInst translates the given x86 PSHUFD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSHUFD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSHUFD: not yet implemented")
}

// --- [ PSHUFHW ] -------------------------------------------------------------

// emitInst translates the given x86 PSHUFHW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSHUFHW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSHUFHW: not yet implemented")
}

// --- [ PSHUFLW ] -------------------------------------------------------------

// emitInst translates the given x86 PSHUFLW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSHUFLW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSHUFLW: not yet implemented")
}

// --- [ PSHUFW ] --------------------------------------------------------------

// emitInst translates the given x86 PSHUFW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSHUFW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSHUFW: not yet implemented")
}

// --- [ PSIGNB ] --------------------------------------------------------------

// emitInst translates the given x86 PSIGNB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSIGNB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSIGNB: not yet implemented")
}

// --- [ PSIGND ] --------------------------------------------------------------

// emitInst translates the given x86 PSIGND instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSIGND(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSIGND: not yet implemented")
}

// --- [ PSIGNW ] --------------------------------------------------------------

// emitInst translates the given x86 PSIGNW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSIGNW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSIGNW: not yet implemented")
}

// --- [ PSLLD ] ---------------------------------------------------------------

// emitInst translates the given x86 PSLLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSLLD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSLLD: not yet implemented")
}

// --- [ PSLLDQ ] --------------------------------------------------------------

// emitInst translates the given x86 PSLLDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSLLDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSLLDQ: not yet implemented")
}

// --- [ PSLLQ ] ---------------------------------------------------------------

// emitInst translates the given x86 PSLLQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSLLQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSLLQ: not yet implemented")
}

// --- [ PSLLW ] ---------------------------------------------------------------

// emitInst translates the given x86 PSLLW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSLLW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSLLW: not yet implemented")
}

// --- [ PSRAD ] ---------------------------------------------------------------

// emitInst translates the given x86 PSRAD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSRAD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRAD: not yet implemented")
}

// --- [ PSRAW ] ---------------------------------------------------------------

// emitInst translates the given x86 PSRAW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSRAW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRAW: not yet implemented")
}

// --- [ PSRLD ] ---------------------------------------------------------------

// emitInst translates the given x86 PSRLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSRLD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRLD: not yet implemented")
}

// --- [ PSRLDQ ] --------------------------------------------------------------

// emitInst translates the given x86 PSRLDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSRLDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRLDQ: not yet implemented")
}

// --- [ PSRLQ ] ---------------------------------------------------------------

// emitInst translates the given x86 PSRLQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSRLQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRLQ: not yet implemented")
}

// --- [ PSRLW ] ---------------------------------------------------------------

// emitInst translates the given x86 PSRLW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSRLW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRLW: not yet implemented")
}

// --- [ PSUBB ] ---------------------------------------------------------------

// emitInst translates the given x86 PSUBB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSUBB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBB: not yet implemented")
}

// --- [ PSUBD ] ---------------------------------------------------------------

// emitInst translates the given x86 PSUBD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSUBD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBD: not yet implemented")
}

// --- [ PSUBQ ] ---------------------------------------------------------------

// emitInst translates the given x86 PSUBQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSUBQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBQ: not yet implemented")
}

// --- [ PSUBSB ] --------------------------------------------------------------

// emitInst translates the given x86 PSUBSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSUBSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBSB: not yet implemented")
}

// --- [ PSUBSW ] --------------------------------------------------------------

// emitInst translates the given x86 PSUBSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSUBSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBSW: not yet implemented")
}

// --- [ PSUBUSB ] -------------------------------------------------------------

// emitInst translates the given x86 PSUBUSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSUBUSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBUSB: not yet implemented")
}

// --- [ PSUBUSW ] -------------------------------------------------------------

// emitInst translates the given x86 PSUBUSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPSUBUSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBUSW: not yet implemented")
}

// --- [ PSUBW ] ---------------------------------------------------------------

// emitInst translates the given x86 PSUBW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPSUBW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBW: not yet implemented")
}

// --- [ PTEST ] ---------------------------------------------------------------

// emitInst translates the given x86 PTEST instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPTEST(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPTEST: not yet implemented")
}

// --- [ PUNPCKHBW ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKHBW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUNPCKHBW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKHBW: not yet implemented")
}

// --- [ PUNPCKHDQ ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKHDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUNPCKHDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKHDQ: not yet implemented")
}

// --- [ PUNPCKHQDQ ] ----------------------------------------------------------

// emitInst translates the given x86 PUNPCKHQDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUNPCKHQDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKHQDQ: not yet implemented")
}

// --- [ PUNPCKHWD ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKHWD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUNPCKHWD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKHWD: not yet implemented")
}

// --- [ PUNPCKLBW ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKLBW instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUNPCKLBW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKLBW: not yet implemented")
}

// --- [ PUNPCKLDQ ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKLDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUNPCKLDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKLDQ: not yet implemented")
}

// --- [ PUNPCKLQDQ ] ----------------------------------------------------------

// emitInst translates the given x86 PUNPCKLQDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUNPCKLQDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKLQDQ: not yet implemented")
}

// --- [ PUNPCKLWD ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKLWD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUNPCKLWD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKLWD: not yet implemented")
}

// --- [ PUSH ] ----------------------------------------------------------------

// emitInst translates the given x86 PUSH instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPUSH(inst *Inst) error {
	v := f.useArg(inst.Arg(0))
	f.push(v)
	return nil
}

// push pushes the given value onto the top of the stack of the function,
// emitting code to f.
func (f *Func) push(v value.Value) {
	m := x86asm.Mem{
		Base: x86asm.ESP,
		Disp: -4,
	}
	mem := NewMem(m, nil)
	f.defMem(mem, v)
	f.espDisp -= 4
}

// --- [ PUSHA ] ---------------------------------------------------------------

// emitInst translates the given x86 PUSHA instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPUSHA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUSHA: not yet implemented")
}

// --- [ PUSHAD ] --------------------------------------------------------------

// emitInst translates the given x86 PUSHAD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUSHAD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUSHAD: not yet implemented")
}

// --- [ PUSHF ] ---------------------------------------------------------------

// emitInst translates the given x86 PUSHF instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPUSHF(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUSHF: not yet implemented")
}

// --- [ PUSHFD ] --------------------------------------------------------------

// emitInst translates the given x86 PUSHFD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUSHFD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUSHFD: not yet implemented")
}

// --- [ PUSHFQ ] --------------------------------------------------------------

// emitInst translates the given x86 PUSHFQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstPUSHFQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUSHFQ: not yet implemented")
}

// --- [ PXOR ] ----------------------------------------------------------------

// emitInst translates the given x86 PXOR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstPXOR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPXOR: not yet implemented")
}

// --- [ RCL ] -----------------------------------------------------------------

// emitInst translates the given x86 RCL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstRCL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRCL: not yet implemented")
}

// --- [ RCPPS ] ---------------------------------------------------------------

// emitInst translates the given x86 RCPPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstRCPPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRCPPS: not yet implemented")
}

// --- [ RCPSS ] ---------------------------------------------------------------

// emitInst translates the given x86 RCPSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstRCPSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRCPSS: not yet implemented")
}

// --- [ RCR ] -----------------------------------------------------------------

// emitInst translates the given x86 RCR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstRCR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRCR: not yet implemented")
}

// --- [ RDFSBASE ] ------------------------------------------------------------

// emitInst translates the given x86 RDFSBASE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstRDFSBASE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDFSBASE: not yet implemented")
}

// --- [ RDGSBASE ] ------------------------------------------------------------

// emitInst translates the given x86 RDGSBASE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstRDGSBASE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDGSBASE: not yet implemented")
}

// --- [ RDMSR ] ---------------------------------------------------------------

// emitInst translates the given x86 RDMSR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstRDMSR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDMSR: not yet implemented")
}

// --- [ RDPMC ] ---------------------------------------------------------------

// emitInst translates the given x86 RDPMC instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstRDPMC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDPMC: not yet implemented")
}

// --- [ RDRAND ] --------------------------------------------------------------

// emitInst translates the given x86 RDRAND instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstRDRAND(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDRAND: not yet implemented")
}

// --- [ RDTSC ] ---------------------------------------------------------------

// emitInst translates the given x86 RDTSC instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstRDTSC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDTSC: not yet implemented")
}

// --- [ RDTSCP ] --------------------------------------------------------------

// emitInst translates the given x86 RDTSCP instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstRDTSCP(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDTSCP: not yet implemented")
}

// --- [ RET ] -----------------------------------------------------------------

// emitInst translates the given x86 RET instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstRET(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRET: not yet implemented")
}

// --- [ ROL ] -----------------------------------------------------------------

// emitInst translates the given x86 ROL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstROL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROL: not yet implemented")
}

// --- [ ROR ] -----------------------------------------------------------------

// emitInst translates the given x86 ROR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstROR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROR: not yet implemented")
}

// --- [ ROUNDPD ] -------------------------------------------------------------

// emitInst translates the given x86 ROUNDPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstROUNDPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROUNDPD: not yet implemented")
}

// --- [ ROUNDPS ] -------------------------------------------------------------

// emitInst translates the given x86 ROUNDPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstROUNDPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROUNDPS: not yet implemented")
}

// --- [ ROUNDSD ] -------------------------------------------------------------

// emitInst translates the given x86 ROUNDSD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstROUNDSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROUNDSD: not yet implemented")
}

// --- [ ROUNDSS ] -------------------------------------------------------------

// emitInst translates the given x86 ROUNDSS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstROUNDSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROUNDSS: not yet implemented")
}

// --- [ RSM ] -----------------------------------------------------------------

// emitInst translates the given x86 RSM instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstRSM(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRSM: not yet implemented")
}

// --- [ RSQRTPS ] -------------------------------------------------------------

// emitInst translates the given x86 RSQRTPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstRSQRTPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRSQRTPS: not yet implemented")
}

// --- [ RSQRTSS ] -------------------------------------------------------------

// emitInst translates the given x86 RSQRTSS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstRSQRTSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRSQRTSS: not yet implemented")
}

// --- [ SAHF ] ----------------------------------------------------------------

// emitInst translates the given x86 SAHF instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSAHF(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSAHF: not yet implemented")
}

// --- [ SAR ] -----------------------------------------------------------------

// emitInst translates the given x86 SAR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSAR(inst *Inst) error {
	// shift arithmetic right (SAR)
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewAShr(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ SBB ] -----------------------------------------------------------------

// emitInst translates the given x86 SBB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSBB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSBB: not yet implemented")
}

// --- [ SCASB ] ---------------------------------------------------------------

// emitInst translates the given x86 SCASB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSCASB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSCASB: not yet implemented")
}

// --- [ SCASD ] ---------------------------------------------------------------

// emitInst translates the given x86 SCASD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSCASD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSCASD: not yet implemented")
}

// --- [ SCASQ ] ---------------------------------------------------------------

// emitInst translates the given x86 SCASQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSCASQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSCASQ: not yet implemented")
}

// --- [ SCASW ] ---------------------------------------------------------------

// emitInst translates the given x86 SCASW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSCASW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSCASW: not yet implemented")
}

// --- [ SFENCE ] --------------------------------------------------------------

// emitInst translates the given x86 SFENCE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSFENCE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSFENCE: not yet implemented")
}

// --- [ SGDT ] ----------------------------------------------------------------

// emitInst translates the given x86 SGDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSGDT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSGDT: not yet implemented")
}

// --- [ SHL ] -----------------------------------------------------------------

// emitInst translates the given x86 SHL instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSHL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSHL: not yet implemented")
}

// --- [ SHLD ] ----------------------------------------------------------------

// emitInst translates the given x86 SHLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSHLD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSHLD: not yet implemented")
}

// --- [ SHR ] -----------------------------------------------------------------

// emitInst translates the given x86 SHR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSHR(inst *Inst) error {
	// shift logical right (SHR)
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewLShr(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ SHRD ] ----------------------------------------------------------------

// emitInst translates the given x86 SHRD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSHRD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSHRD: not yet implemented")
}

// --- [ SHUFPD ] --------------------------------------------------------------

// emitInst translates the given x86 SHUFPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSHUFPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSHUFPD: not yet implemented")
}

// --- [ SHUFPS ] --------------------------------------------------------------

// emitInst translates the given x86 SHUFPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSHUFPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSHUFPS: not yet implemented")
}

// --- [ SIDT ] ----------------------------------------------------------------

// emitInst translates the given x86 SIDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSIDT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSIDT: not yet implemented")
}

// --- [ SLDT ] ----------------------------------------------------------------

// emitInst translates the given x86 SLDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSLDT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSLDT: not yet implemented")
}

// --- [ SMSW ] ----------------------------------------------------------------

// emitInst translates the given x86 SMSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSMSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSMSW: not yet implemented")
}

// --- [ SQRTPD ] --------------------------------------------------------------

// emitInst translates the given x86 SQRTPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSQRTPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSQRTPD: not yet implemented")
}

// --- [ SQRTPS ] --------------------------------------------------------------

// emitInst translates the given x86 SQRTPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSQRTPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSQRTPS: not yet implemented")
}

// --- [ SQRTSD ] --------------------------------------------------------------

// emitInst translates the given x86 SQRTSD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSQRTSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSQRTSD: not yet implemented")
}

// --- [ SQRTSS ] --------------------------------------------------------------

// emitInst translates the given x86 SQRTSS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSQRTSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSQRTSS: not yet implemented")
}

// --- [ STC ] -----------------------------------------------------------------

// emitInst translates the given x86 STC instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSTC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTC: not yet implemented")
}

// --- [ STD ] -----------------------------------------------------------------

// emitInst translates the given x86 STD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSTD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTD: not yet implemented")
}

// --- [ STI ] -----------------------------------------------------------------

// emitInst translates the given x86 STI instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSTI(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTI: not yet implemented")
}

// --- [ STMXCSR ] -------------------------------------------------------------

// emitInst translates the given x86 STMXCSR instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSTMXCSR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTMXCSR: not yet implemented")
}

// --- [ STOSB ] ---------------------------------------------------------------

// emitInst translates the given x86 STOSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSTOSB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTOSB: not yet implemented")
}

// --- [ STOSD ] ---------------------------------------------------------------

// emitInst translates the given x86 STOSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSTOSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTOSD: not yet implemented")
}

// --- [ STOSQ ] ---------------------------------------------------------------

// emitInst translates the given x86 STOSQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSTOSQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTOSQ: not yet implemented")
}

// --- [ STOSW ] ---------------------------------------------------------------

// emitInst translates the given x86 STOSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSTOSW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTOSW: not yet implemented")
}

// --- [ STR ] -----------------------------------------------------------------

// emitInst translates the given x86 STR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSTR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTR: not yet implemented")
}

// --- [ SUB ] -----------------------------------------------------------------

// emitInst translates the given x86 SUB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSUB(inst *Inst) error {
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewSub(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ SUBPD ] ---------------------------------------------------------------

// emitInst translates the given x86 SUBPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSUBPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSUBPD: not yet implemented")
}

// --- [ SUBPS ] ---------------------------------------------------------------

// emitInst translates the given x86 SUBPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSUBPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSUBPS: not yet implemented")
}

// --- [ SUBSD ] ---------------------------------------------------------------

// emitInst translates the given x86 SUBSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSUBSD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSUBSD: not yet implemented")
}

// --- [ SUBSS ] ---------------------------------------------------------------

// emitInst translates the given x86 SUBSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstSUBSS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSUBSS: not yet implemented")
}

// --- [ SWAPGS ] --------------------------------------------------------------

// emitInst translates the given x86 SWAPGS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSWAPGS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSWAPGS: not yet implemented")
}

// --- [ SYSCALL ] -------------------------------------------------------------

// emitInst translates the given x86 SYSCALL instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSYSCALL(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSYSCALL: not yet implemented")
}

// --- [ SYSENTER ] ------------------------------------------------------------

// emitInst translates the given x86 SYSENTER instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSYSENTER(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSYSENTER: not yet implemented")
}

// --- [ SYSEXIT ] -------------------------------------------------------------

// emitInst translates the given x86 SYSEXIT instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSYSEXIT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSYSEXIT: not yet implemented")
}

// --- [ SYSRET ] --------------------------------------------------------------

// emitInst translates the given x86 SYSRET instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstSYSRET(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSYSRET: not yet implemented")
}

// --- [ TEST ] ----------------------------------------------------------------

// emitInst translates the given x86 TEST instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstTEST(inst *Inst) error {
	// result = x AND y; set PF, ZF, and SF according to result.
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewAnd(x, y)

	// PF (bit 2) Parity flag - Set if the least-significant byte of the result
	// contains an even number of 1 bits; cleared otherwise.

	// TODO: Add support for the PF status flag.

	// ZF (bit 6) Zero flag - Set if the result is zero; cleared otherwise.
	zero := constant.NewInt(0, types.I32)
	zf := f.cur.NewICmp(ir.IntEQ, result, zero)
	f.defStatus(ZF, zf)

	// SF (bit 7) Sign flag - Set equal to the most-significant bit of the
	// result, which is the sign bit of a signed integer. (0 indicates a positive
	// value and 1 indicates a negative value.)

	// TODO: Add support for the SF flag.

	return nil

}

// --- [ TZCNT ] ---------------------------------------------------------------

// emitInst translates the given x86 TZCNT instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstTZCNT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstTZCNT: not yet implemented")
}

// --- [ UCOMISD ] -------------------------------------------------------------

// emitInst translates the given x86 UCOMISD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstUCOMISD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUCOMISD: not yet implemented")
}

// --- [ UCOMISS ] -------------------------------------------------------------

// emitInst translates the given x86 UCOMISS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstUCOMISS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUCOMISS: not yet implemented")
}

// --- [ UD1 ] -----------------------------------------------------------------

// emitInst translates the given x86 UD1 instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstUD1(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUD1: not yet implemented")
}

// --- [ UD2 ] -----------------------------------------------------------------

// emitInst translates the given x86 UD2 instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstUD2(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUD2: not yet implemented")
}

// --- [ UNPCKHPD ] ------------------------------------------------------------

// emitInst translates the given x86 UNPCKHPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstUNPCKHPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUNPCKHPD: not yet implemented")
}

// --- [ UNPCKHPS ] ------------------------------------------------------------

// emitInst translates the given x86 UNPCKHPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstUNPCKHPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUNPCKHPS: not yet implemented")
}

// --- [ UNPCKLPD ] ------------------------------------------------------------

// emitInst translates the given x86 UNPCKLPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstUNPCKLPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUNPCKLPD: not yet implemented")
}

// --- [ UNPCKLPS ] ------------------------------------------------------------

// emitInst translates the given x86 UNPCKLPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstUNPCKLPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUNPCKLPS: not yet implemented")
}

// --- [ VERR ] ----------------------------------------------------------------

// emitInst translates the given x86 VERR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstVERR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVERR: not yet implemented")
}

// --- [ VERW ] ----------------------------------------------------------------

// emitInst translates the given x86 VERW instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstVERW(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVERW: not yet implemented")
}

// --- [ VMOVDQA ] -------------------------------------------------------------

// emitInst translates the given x86 VMOVDQA instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstVMOVDQA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVMOVDQA: not yet implemented")
}

// --- [ VMOVDQU ] -------------------------------------------------------------

// emitInst translates the given x86 VMOVDQU instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstVMOVDQU(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVMOVDQU: not yet implemented")
}

// --- [ VMOVNTDQ ] ------------------------------------------------------------

// emitInst translates the given x86 VMOVNTDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstVMOVNTDQ(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVMOVNTDQ: not yet implemented")
}

// --- [ VMOVNTDQA ] -----------------------------------------------------------

// emitInst translates the given x86 VMOVNTDQA instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstVMOVNTDQA(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVMOVNTDQA: not yet implemented")
}

// --- [ VZEROUPPER ] ----------------------------------------------------------

// emitInst translates the given x86 VZEROUPPER instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstVZEROUPPER(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVZEROUPPER: not yet implemented")
}

// --- [ WBINVD ] --------------------------------------------------------------

// emitInst translates the given x86 WBINVD instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstWBINVD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstWBINVD: not yet implemented")
}

// --- [ WRFSBASE ] ------------------------------------------------------------

// emitInst translates the given x86 WRFSBASE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstWRFSBASE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstWRFSBASE: not yet implemented")
}

// --- [ WRGSBASE ] ------------------------------------------------------------

// emitInst translates the given x86 WRGSBASE instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstWRGSBASE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstWRGSBASE: not yet implemented")
}

// --- [ WRMSR ] ---------------------------------------------------------------

// emitInst translates the given x86 WRMSR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstWRMSR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstWRMSR: not yet implemented")
}

// --- [ XABORT ] --------------------------------------------------------------

// emitInst translates the given x86 XABORT instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXABORT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXABORT: not yet implemented")
}

// --- [ XADD ] ----------------------------------------------------------------

// emitInst translates the given x86 XADD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstXADD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXADD: not yet implemented")
}

// --- [ XBEGIN ] --------------------------------------------------------------

// emitInst translates the given x86 XBEGIN instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXBEGIN(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXBEGIN: not yet implemented")
}

// --- [ XCHG ] ----------------------------------------------------------------

// emitInst translates the given x86 XCHG instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstXCHG(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXCHG: not yet implemented")
}

// --- [ XEND ] ----------------------------------------------------------------

// emitInst translates the given x86 XEND instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstXEND(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXEND: not yet implemented")
}

// --- [ XGETBV ] --------------------------------------------------------------

// emitInst translates the given x86 XGETBV instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXGETBV(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXGETBV: not yet implemented")
}

// --- [ XLATB ] ---------------------------------------------------------------

// emitInst translates the given x86 XLATB instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstXLATB(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXLATB: not yet implemented")
}

// --- [ XOR ] -----------------------------------------------------------------

// emitInst translates the given x86 XOR instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstXOR(inst *Inst) error {
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewXor(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ XORPD ] ---------------------------------------------------------------

// emitInst translates the given x86 XORPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstXORPD(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXORPD: not yet implemented")
}

// --- [ XORPS ] ---------------------------------------------------------------

// emitInst translates the given x86 XORPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstXORPS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXORPS: not yet implemented")
}

// --- [ XRSTOR ] --------------------------------------------------------------

// emitInst translates the given x86 XRSTOR instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXRSTOR(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXRSTOR: not yet implemented")
}

// --- [ XRSTOR64 ] ------------------------------------------------------------

// emitInst translates the given x86 XRSTOR64 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXRSTOR64(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXRSTOR64: not yet implemented")
}

// --- [ XRSTORS ] -------------------------------------------------------------

// emitInst translates the given x86 XRSTORS instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXRSTORS(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXRSTORS: not yet implemented")
}

// --- [ XRSTORS64 ] -----------------------------------------------------------

// emitInst translates the given x86 XRSTORS64 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXRSTORS64(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXRSTORS64: not yet implemented")
}

// --- [ XSAVE ] ---------------------------------------------------------------

// emitInst translates the given x86 XSAVE instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstXSAVE(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVE: not yet implemented")
}

// --- [ XSAVE64 ] -------------------------------------------------------------

// emitInst translates the given x86 XSAVE64 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXSAVE64(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVE64: not yet implemented")
}

// --- [ XSAVEC ] --------------------------------------------------------------

// emitInst translates the given x86 XSAVEC instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXSAVEC(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVEC: not yet implemented")
}

// --- [ XSAVEC64 ] ------------------------------------------------------------

// emitInst translates the given x86 XSAVEC64 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXSAVEC64(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVEC64: not yet implemented")
}

// --- [ XSAVEOPT ] ------------------------------------------------------------

// emitInst translates the given x86 XSAVEOPT instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXSAVEOPT(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVEOPT: not yet implemented")
}

// --- [ XSAVEOPT64 ] ----------------------------------------------------------

// emitInst translates the given x86 XSAVEOPT64 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXSAVEOPT64(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVEOPT64: not yet implemented")
}

// --- [ XSAVES ] --------------------------------------------------------------

// emitInst translates the given x86 XSAVES instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXSAVES(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVES: not yet implemented")
}

// --- [ XSAVES64 ] ------------------------------------------------------------

// emitInst translates the given x86 XSAVES64 instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXSAVES64(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVES64: not yet implemented")
}

// --- [ XSETBV ] --------------------------------------------------------------

// emitInst translates the given x86 XSETBV instruction to LLVM IR, emitting
// code to f.
func (f *Func) emitInstXSETBV(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSETBV: not yet implemented")
}

// --- [ XTEST ] ---------------------------------------------------------------

// emitInst translates the given x86 XTEST instruction to LLVM IR, emitting code
// to f.
func (f *Func) emitInstXTEST(inst *Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXTEST: not yet implemented")
}
