package lift

import (
	"fmt"

	"github.com/decomp/exp/disasm/x86"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
	"golang.org/x/arch/x86/x86asm"
)

// liftInst lifts the given x86 instruction to LLVM IR, emitting code to f.
func (f *Func) liftInst(inst *x86.Inst) error {
	dbg.Println("lifting instruction:", inst.Inst)

	// Check if prefix is present.
	var (
		hasREP  bool
		hasREPN bool
	)
	for _, prefix := range inst.Prefix[:] {
		// The first zero in the array marks the end of the prefixes.
		if prefix == 0 {
			break
		}
		switch prefix &^ x86asm.PrefixImplicit {
		case x86asm.PrefixData16:
			// prefix already supported.
		case x86asm.PrefixREP:
			hasREP = true
		case x86asm.PrefixREPN:
			hasREPN = true
		case x86asm.PrefixREX | x86asm.PrefixREXW:
			// TODO: Implement support for REX.W
		default:
			pretty.Println("instruction with prefix:", inst)
			panic(fmt.Errorf("support for %v instruction with prefix %v (0x%04X) not yet implemented", inst.Op, prefix, uint16(prefix)))
		}
	}

	// Repeat instruction.
	switch {
	case hasREP:
		return f.liftREPInst(inst)
	case hasREPN:
		return f.liftREPNInst(inst)
	}

	// Translate instruction.
	switch inst.Op {
	case x86asm.AAA:
		return f.liftInstAAA(inst)
	case x86asm.AAD:
		return f.liftInstAAD(inst)
	case x86asm.AAM:
		return f.liftInstAAM(inst)
	case x86asm.AAS:
		return f.liftInstAAS(inst)
	case x86asm.ADC:
		return f.liftInstADC(inst)
	case x86asm.ADD:
		return f.liftInstADD(inst)
	case x86asm.ADDPD:
		return f.liftInstADDPD(inst)
	case x86asm.ADDPS:
		return f.liftInstADDPS(inst)
	case x86asm.ADDSD:
		return f.liftInstADDSD(inst)
	case x86asm.ADDSS:
		return f.liftInstADDSS(inst)
	case x86asm.ADDSUBPD:
		return f.liftInstADDSUBPD(inst)
	case x86asm.ADDSUBPS:
		return f.liftInstADDSUBPS(inst)
	case x86asm.AESDEC:
		return f.liftInstAESDEC(inst)
	case x86asm.AESDECLAST:
		return f.liftInstAESDECLAST(inst)
	case x86asm.AESENC:
		return f.liftInstAESENC(inst)
	case x86asm.AESENCLAST:
		return f.liftInstAESENCLAST(inst)
	case x86asm.AESIMC:
		return f.liftInstAESIMC(inst)
	case x86asm.AESKEYGENASSIST:
		return f.liftInstAESKEYGENASSIST(inst)
	case x86asm.AND:
		return f.liftInstAND(inst)
	case x86asm.ANDNPD:
		return f.liftInstANDNPD(inst)
	case x86asm.ANDNPS:
		return f.liftInstANDNPS(inst)
	case x86asm.ANDPD:
		return f.liftInstANDPD(inst)
	case x86asm.ANDPS:
		return f.liftInstANDPS(inst)
	case x86asm.ARPL:
		return f.liftInstARPL(inst)
	case x86asm.BLENDPD:
		return f.liftInstBLENDPD(inst)
	case x86asm.BLENDPS:
		return f.liftInstBLENDPS(inst)
	case x86asm.BLENDVPD:
		return f.liftInstBLENDVPD(inst)
	case x86asm.BLENDVPS:
		return f.liftInstBLENDVPS(inst)
	case x86asm.BOUND:
		return f.liftInstBOUND(inst)
	case x86asm.BSF:
		return f.liftInstBSF(inst)
	case x86asm.BSR:
		return f.liftInstBSR(inst)
	case x86asm.BSWAP:
		return f.liftInstBSWAP(inst)
	case x86asm.BT:
		return f.liftInstBT(inst)
	case x86asm.BTC:
		return f.liftInstBTC(inst)
	case x86asm.BTR:
		return f.liftInstBTR(inst)
	case x86asm.BTS:
		return f.liftInstBTS(inst)
	case x86asm.CALL:
		return f.liftInstCALL(inst)
	case x86asm.CBW:
		return f.liftInstCBW(inst)
	case x86asm.CDQ:
		return f.liftInstCDQ(inst)
	case x86asm.CDQE:
		return f.liftInstCDQE(inst)
	case x86asm.CLC:
		return f.liftInstCLC(inst)
	case x86asm.CLD:
		return f.liftInstCLD(inst)
	case x86asm.CLFLUSH:
		return f.liftInstCLFLUSH(inst)
	case x86asm.CLI:
		return f.liftInstCLI(inst)
	case x86asm.CLTS:
		return f.liftInstCLTS(inst)
	case x86asm.CMC:
		return f.liftInstCMC(inst)
	case x86asm.CMOVA:
		return f.liftInstCMOVA(inst)
	case x86asm.CMOVAE:
		return f.liftInstCMOVAE(inst)
	case x86asm.CMOVB:
		return f.liftInstCMOVB(inst)
	case x86asm.CMOVBE:
		return f.liftInstCMOVBE(inst)
	case x86asm.CMOVE:
		return f.liftInstCMOVE(inst)
	case x86asm.CMOVG:
		return f.liftInstCMOVG(inst)
	case x86asm.CMOVGE:
		return f.liftInstCMOVGE(inst)
	case x86asm.CMOVL:
		return f.liftInstCMOVL(inst)
	case x86asm.CMOVLE:
		return f.liftInstCMOVLE(inst)
	case x86asm.CMOVNE:
		return f.liftInstCMOVNE(inst)
	case x86asm.CMOVNO:
		return f.liftInstCMOVNO(inst)
	case x86asm.CMOVNP:
		return f.liftInstCMOVNP(inst)
	case x86asm.CMOVNS:
		return f.liftInstCMOVNS(inst)
	case x86asm.CMOVO:
		return f.liftInstCMOVO(inst)
	case x86asm.CMOVP:
		return f.liftInstCMOVP(inst)
	case x86asm.CMOVS:
		return f.liftInstCMOVS(inst)
	case x86asm.CMP:
		return f.liftInstCMP(inst)
	case x86asm.CMPPD:
		return f.liftInstCMPPD(inst)
	case x86asm.CMPPS:
		return f.liftInstCMPPS(inst)
	case x86asm.CMPSB:
		return f.liftInstCMPSB(inst)
	case x86asm.CMPSD:
		return f.liftInstCMPSD(inst)
	case x86asm.CMPSD_XMM:
		return f.liftInstCMPSD_XMM(inst)
	case x86asm.CMPSQ:
		return f.liftInstCMPSQ(inst)
	case x86asm.CMPSS:
		return f.liftInstCMPSS(inst)
	case x86asm.CMPSW:
		return f.liftInstCMPSW(inst)
	case x86asm.CMPXCHG:
		return f.liftInstCMPXCHG(inst)
	case x86asm.CMPXCHG16B:
		return f.liftInstCMPXCHG16B(inst)
	case x86asm.CMPXCHG8B:
		return f.liftInstCMPXCHG8B(inst)
	case x86asm.COMISD:
		return f.liftInstCOMISD(inst)
	case x86asm.COMISS:
		return f.liftInstCOMISS(inst)
	case x86asm.CPUID:
		return f.liftInstCPUID(inst)
	case x86asm.CQO:
		return f.liftInstCQO(inst)
	case x86asm.CRC32:
		return f.liftInstCRC32(inst)
	case x86asm.CVTDQ2PD:
		return f.liftInstCVTDQ2PD(inst)
	case x86asm.CVTDQ2PS:
		return f.liftInstCVTDQ2PS(inst)
	case x86asm.CVTPD2DQ:
		return f.liftInstCVTPD2DQ(inst)
	case x86asm.CVTPD2PI:
		return f.liftInstCVTPD2PI(inst)
	case x86asm.CVTPD2PS:
		return f.liftInstCVTPD2PS(inst)
	case x86asm.CVTPI2PD:
		return f.liftInstCVTPI2PD(inst)
	case x86asm.CVTPI2PS:
		return f.liftInstCVTPI2PS(inst)
	case x86asm.CVTPS2DQ:
		return f.liftInstCVTPS2DQ(inst)
	case x86asm.CVTPS2PD:
		return f.liftInstCVTPS2PD(inst)
	case x86asm.CVTPS2PI:
		return f.liftInstCVTPS2PI(inst)
	case x86asm.CVTSD2SI:
		return f.liftInstCVTSD2SI(inst)
	case x86asm.CVTSD2SS:
		return f.liftInstCVTSD2SS(inst)
	case x86asm.CVTSI2SD:
		return f.liftInstCVTSI2SD(inst)
	case x86asm.CVTSI2SS:
		return f.liftInstCVTSI2SS(inst)
	case x86asm.CVTSS2SD:
		return f.liftInstCVTSS2SD(inst)
	case x86asm.CVTSS2SI:
		return f.liftInstCVTSS2SI(inst)
	case x86asm.CVTTPD2DQ:
		return f.liftInstCVTTPD2DQ(inst)
	case x86asm.CVTTPD2PI:
		return f.liftInstCVTTPD2PI(inst)
	case x86asm.CVTTPS2DQ:
		return f.liftInstCVTTPS2DQ(inst)
	case x86asm.CVTTPS2PI:
		return f.liftInstCVTTPS2PI(inst)
	case x86asm.CVTTSD2SI:
		return f.liftInstCVTTSD2SI(inst)
	case x86asm.CVTTSS2SI:
		return f.liftInstCVTTSS2SI(inst)
	case x86asm.CWD:
		return f.liftInstCWD(inst)
	case x86asm.CWDE:
		return f.liftInstCWDE(inst)
	case x86asm.DAA:
		return f.liftInstDAA(inst)
	case x86asm.DAS:
		return f.liftInstDAS(inst)
	case x86asm.DEC:
		return f.liftInstDEC(inst)
	case x86asm.DIV:
		return f.liftInstDIV(inst)
	case x86asm.DIVPD:
		return f.liftInstDIVPD(inst)
	case x86asm.DIVPS:
		return f.liftInstDIVPS(inst)
	case x86asm.DIVSD:
		return f.liftInstDIVSD(inst)
	case x86asm.DIVSS:
		return f.liftInstDIVSS(inst)
	case x86asm.DPPD:
		return f.liftInstDPPD(inst)
	case x86asm.DPPS:
		return f.liftInstDPPS(inst)
	case x86asm.EMMS:
		return f.liftInstEMMS(inst)
	case x86asm.ENTER:
		return f.liftInstENTER(inst)
	case x86asm.EXTRACTPS:
		return f.liftInstEXTRACTPS(inst)
	case x86asm.F2XM1:
		return f.liftInstF2XM1(inst)
	case x86asm.FABS:
		return f.liftInstFABS(inst)
	case x86asm.FADD:
		return f.liftInstFADD(inst)
	case x86asm.FADDP:
		return f.liftInstFADDP(inst)
	case x86asm.FBLD:
		return f.liftInstFBLD(inst)
	case x86asm.FBSTP:
		return f.liftInstFBSTP(inst)
	case x86asm.FCHS:
		return f.liftInstFCHS(inst)
	case x86asm.FCMOVB:
		return f.liftInstFCMOVB(inst)
	case x86asm.FCMOVBE:
		return f.liftInstFCMOVBE(inst)
	case x86asm.FCMOVE:
		return f.liftInstFCMOVE(inst)
	case x86asm.FCMOVNB:
		return f.liftInstFCMOVNB(inst)
	case x86asm.FCMOVNBE:
		return f.liftInstFCMOVNBE(inst)
	case x86asm.FCMOVNE:
		return f.liftInstFCMOVNE(inst)
	case x86asm.FCMOVNU:
		return f.liftInstFCMOVNU(inst)
	case x86asm.FCMOVU:
		return f.liftInstFCMOVU(inst)
	case x86asm.FCOM:
		return f.liftInstFCOM(inst)
	case x86asm.FCOMI:
		return f.liftInstFCOMI(inst)
	case x86asm.FCOMIP:
		return f.liftInstFCOMIP(inst)
	case x86asm.FCOMP:
		return f.liftInstFCOMP(inst)
	case x86asm.FCOMPP:
		return f.liftInstFCOMPP(inst)
	case x86asm.FCOS:
		return f.liftInstFCOS(inst)
	case x86asm.FDECSTP:
		return f.liftInstFDECSTP(inst)
	case x86asm.FDIV:
		return f.liftInstFDIV(inst)
	case x86asm.FDIVP:
		return f.liftInstFDIVP(inst)
	case x86asm.FDIVR:
		return f.liftInstFDIVR(inst)
	case x86asm.FDIVRP:
		return f.liftInstFDIVRP(inst)
	case x86asm.FFREE:
		return f.liftInstFFREE(inst)
	case x86asm.FFREEP:
		return f.liftInstFFREEP(inst)
	case x86asm.FIADD:
		return f.liftInstFIADD(inst)
	case x86asm.FICOM:
		return f.liftInstFICOM(inst)
	case x86asm.FICOMP:
		return f.liftInstFICOMP(inst)
	case x86asm.FIDIV:
		return f.liftInstFIDIV(inst)
	case x86asm.FIDIVR:
		return f.liftInstFIDIVR(inst)
	case x86asm.FILD:
		return f.liftInstFILD(inst)
	case x86asm.FIMUL:
		return f.liftInstFIMUL(inst)
	case x86asm.FINCSTP:
		return f.liftInstFINCSTP(inst)
	case x86asm.FIST:
		return f.liftInstFIST(inst)
	case x86asm.FISTP:
		return f.liftInstFISTP(inst)
	case x86asm.FISTTP:
		return f.liftInstFISTTP(inst)
	case x86asm.FISUB:
		return f.liftInstFISUB(inst)
	case x86asm.FISUBR:
		return f.liftInstFISUBR(inst)
	case x86asm.FLD:
		return f.liftInstFLD(inst)
	case x86asm.FLD1:
		return f.liftInstFLD1(inst)
	case x86asm.FLDCW:
		return f.liftInstFLDCW(inst)
	case x86asm.FLDENV:
		return f.liftInstFLDENV(inst)
	case x86asm.FLDL2E:
		return f.liftInstFLDL2E(inst)
	case x86asm.FLDL2T:
		return f.liftInstFLDL2T(inst)
	case x86asm.FLDLG2:
		return f.liftInstFLDLG2(inst)
	case x86asm.FLDLN2:
		return f.liftInstFLDLN2(inst)
	case x86asm.FLDPI:
		return f.liftInstFLDPI(inst)
	case x86asm.FLDZ:
		return f.liftInstFLDZ(inst)
	case x86asm.FMUL:
		return f.liftInstFMUL(inst)
	case x86asm.FMULP:
		return f.liftInstFMULP(inst)
	case x86asm.FNCLEX:
		return f.liftInstFNCLEX(inst)
	case x86asm.FNINIT:
		return f.liftInstFNINIT(inst)
	case x86asm.FNOP:
		return f.liftInstFNOP(inst)
	case x86asm.FNSAVE:
		return f.liftInstFNSAVE(inst)
	case x86asm.FNSTCW:
		return f.liftInstFNSTCW(inst)
	case x86asm.FNSTENV:
		return f.liftInstFNSTENV(inst)
	case x86asm.FNSTSW:
		return f.liftInstFNSTSW(inst)
	case x86asm.FPATAN:
		return f.liftInstFPATAN(inst)
	case x86asm.FPREM:
		return f.liftInstFPREM(inst)
	case x86asm.FPREM1:
		return f.liftInstFPREM1(inst)
	case x86asm.FPTAN:
		return f.liftInstFPTAN(inst)
	case x86asm.FRNDINT:
		return f.liftInstFRNDINT(inst)
	case x86asm.FRSTOR:
		return f.liftInstFRSTOR(inst)
	case x86asm.FSCALE:
		return f.liftInstFSCALE(inst)
	case x86asm.FSIN:
		return f.liftInstFSIN(inst)
	case x86asm.FSINCOS:
		return f.liftInstFSINCOS(inst)
	case x86asm.FSQRT:
		return f.liftInstFSQRT(inst)
	case x86asm.FST:
		return f.liftInstFST(inst)
	case x86asm.FSTP:
		return f.liftInstFSTP(inst)
	case x86asm.FSUB:
		return f.liftInstFSUB(inst)
	case x86asm.FSUBP:
		return f.liftInstFSUBP(inst)
	case x86asm.FSUBR:
		return f.liftInstFSUBR(inst)
	case x86asm.FSUBRP:
		return f.liftInstFSUBRP(inst)
	case x86asm.FTST:
		return f.liftInstFTST(inst)
	case x86asm.FUCOM:
		return f.liftInstFUCOM(inst)
	case x86asm.FUCOMI:
		return f.liftInstFUCOMI(inst)
	case x86asm.FUCOMIP:
		return f.liftInstFUCOMIP(inst)
	case x86asm.FUCOMP:
		return f.liftInstFUCOMP(inst)
	case x86asm.FUCOMPP:
		return f.liftInstFUCOMPP(inst)
	case x86asm.FWAIT:
		return f.liftInstFWAIT(inst)
	case x86asm.FXAM:
		return f.liftInstFXAM(inst)
	case x86asm.FXCH:
		return f.liftInstFXCH(inst)
	case x86asm.FXRSTOR:
		return f.liftInstFXRSTOR(inst)
	case x86asm.FXRSTOR64:
		return f.liftInstFXRSTOR64(inst)
	case x86asm.FXSAVE:
		return f.liftInstFXSAVE(inst)
	case x86asm.FXSAVE64:
		return f.liftInstFXSAVE64(inst)
	case x86asm.FXTRACT:
		return f.liftInstFXTRACT(inst)
	case x86asm.FYL2X:
		return f.liftInstFYL2X(inst)
	case x86asm.FYL2XP1:
		return f.liftInstFYL2XP1(inst)
	case x86asm.HADDPD:
		return f.liftInstHADDPD(inst)
	case x86asm.HADDPS:
		return f.liftInstHADDPS(inst)
	case x86asm.HLT:
		return f.liftInstHLT(inst)
	case x86asm.HSUBPD:
		return f.liftInstHSUBPD(inst)
	case x86asm.HSUBPS:
		return f.liftInstHSUBPS(inst)
	case x86asm.ICEBP:
		return f.liftInstICEBP(inst)
	case x86asm.IDIV:
		return f.liftInstIDIV(inst)
	case x86asm.IMUL:
		return f.liftInstIMUL(inst)
	case x86asm.IN:
		return f.liftInstIN(inst)
	case x86asm.INC:
		return f.liftInstINC(inst)
	case x86asm.INSB:
		return f.liftInstINSB(inst)
	case x86asm.INSD:
		return f.liftInstINSD(inst)
	case x86asm.INSERTPS:
		return f.liftInstINSERTPS(inst)
	case x86asm.INSW:
		return f.liftInstINSW(inst)
	case x86asm.INT:
		return f.liftInstINT(inst)
	case x86asm.INTO:
		return f.liftInstINTO(inst)
	case x86asm.INVD:
		return f.liftInstINVD(inst)
	case x86asm.INVLPG:
		return f.liftInstINVLPG(inst)
	case x86asm.INVPCID:
		return f.liftInstINVPCID(inst)
	case x86asm.IRET:
		return f.liftInstIRET(inst)
	case x86asm.IRETD:
		return f.liftInstIRETD(inst)
	case x86asm.IRETQ:
		return f.liftInstIRETQ(inst)
	case x86asm.LAHF:
		return f.liftInstLAHF(inst)
	case x86asm.LAR:
		return f.liftInstLAR(inst)
	case x86asm.LCALL:
		return f.liftInstLCALL(inst)
	case x86asm.LDDQU:
		return f.liftInstLDDQU(inst)
	case x86asm.LDMXCSR:
		return f.liftInstLDMXCSR(inst)
	case x86asm.LDS:
		return f.liftInstLDS(inst)
	case x86asm.LEA:
		return f.liftInstLEA(inst)
	case x86asm.LEAVE:
		return f.liftInstLEAVE(inst)
	case x86asm.LES:
		return f.liftInstLES(inst)
	case x86asm.LFENCE:
		return f.liftInstLFENCE(inst)
	case x86asm.LFS:
		return f.liftInstLFS(inst)
	case x86asm.LGDT:
		return f.liftInstLGDT(inst)
	case x86asm.LGS:
		return f.liftInstLGS(inst)
	case x86asm.LIDT:
		return f.liftInstLIDT(inst)
	case x86asm.LJMP:
		return f.liftInstLJMP(inst)
	case x86asm.LLDT:
		return f.liftInstLLDT(inst)
	case x86asm.LMSW:
		return f.liftInstLMSW(inst)
	case x86asm.LODSB:
		return f.liftInstLODSB(inst)
	case x86asm.LODSD:
		return f.liftInstLODSD(inst)
	case x86asm.LODSQ:
		return f.liftInstLODSQ(inst)
	case x86asm.LODSW:
		return f.liftInstLODSW(inst)
	case x86asm.LRET:
		return f.liftInstLRET(inst)
	case x86asm.LSL:
		return f.liftInstLSL(inst)
	case x86asm.LSS:
		return f.liftInstLSS(inst)
	case x86asm.LTR:
		return f.liftInstLTR(inst)
	case x86asm.LZCNT:
		return f.liftInstLZCNT(inst)
	case x86asm.MASKMOVDQU:
		return f.liftInstMASKMOVDQU(inst)
	case x86asm.MASKMOVQ:
		return f.liftInstMASKMOVQ(inst)
	case x86asm.MAXPD:
		return f.liftInstMAXPD(inst)
	case x86asm.MAXPS:
		return f.liftInstMAXPS(inst)
	case x86asm.MAXSD:
		return f.liftInstMAXSD(inst)
	case x86asm.MAXSS:
		return f.liftInstMAXSS(inst)
	case x86asm.MFENCE:
		return f.liftInstMFENCE(inst)
	case x86asm.MINPD:
		return f.liftInstMINPD(inst)
	case x86asm.MINPS:
		return f.liftInstMINPS(inst)
	case x86asm.MINSD:
		return f.liftInstMINSD(inst)
	case x86asm.MINSS:
		return f.liftInstMINSS(inst)
	case x86asm.MONITOR:
		return f.liftInstMONITOR(inst)
	case x86asm.MOV:
		return f.liftInstMOV(inst)
	case x86asm.MOVAPD:
		return f.liftInstMOVAPD(inst)
	case x86asm.MOVAPS:
		return f.liftInstMOVAPS(inst)
	case x86asm.MOVBE:
		return f.liftInstMOVBE(inst)
	case x86asm.MOVD:
		return f.liftInstMOVD(inst)
	case x86asm.MOVDDUP:
		return f.liftInstMOVDDUP(inst)
	case x86asm.MOVDQ2Q:
		return f.liftInstMOVDQ2Q(inst)
	case x86asm.MOVDQA:
		return f.liftInstMOVDQA(inst)
	case x86asm.MOVDQU:
		return f.liftInstMOVDQU(inst)
	case x86asm.MOVHLPS:
		return f.liftInstMOVHLPS(inst)
	case x86asm.MOVHPD:
		return f.liftInstMOVHPD(inst)
	case x86asm.MOVHPS:
		return f.liftInstMOVHPS(inst)
	case x86asm.MOVLHPS:
		return f.liftInstMOVLHPS(inst)
	case x86asm.MOVLPD:
		return f.liftInstMOVLPD(inst)
	case x86asm.MOVLPS:
		return f.liftInstMOVLPS(inst)
	case x86asm.MOVMSKPD:
		return f.liftInstMOVMSKPD(inst)
	case x86asm.MOVMSKPS:
		return f.liftInstMOVMSKPS(inst)
	case x86asm.MOVNTDQ:
		return f.liftInstMOVNTDQ(inst)
	case x86asm.MOVNTDQA:
		return f.liftInstMOVNTDQA(inst)
	case x86asm.MOVNTI:
		return f.liftInstMOVNTI(inst)
	case x86asm.MOVNTPD:
		return f.liftInstMOVNTPD(inst)
	case x86asm.MOVNTPS:
		return f.liftInstMOVNTPS(inst)
	case x86asm.MOVNTQ:
		return f.liftInstMOVNTQ(inst)
	case x86asm.MOVNTSD:
		return f.liftInstMOVNTSD(inst)
	case x86asm.MOVNTSS:
		return f.liftInstMOVNTSS(inst)
	case x86asm.MOVQ:
		return f.liftInstMOVQ(inst)
	case x86asm.MOVQ2DQ:
		return f.liftInstMOVQ2DQ(inst)
	case x86asm.MOVSB:
		return f.liftInstMOVSB(inst)
	case x86asm.MOVSD:
		return f.liftInstMOVSD(inst)
	case x86asm.MOVSD_XMM:
		return f.liftInstMOVSD_XMM(inst)
	case x86asm.MOVSHDUP:
		return f.liftInstMOVSHDUP(inst)
	case x86asm.MOVSLDUP:
		return f.liftInstMOVSLDUP(inst)
	case x86asm.MOVSQ:
		return f.liftInstMOVSQ(inst)
	case x86asm.MOVSS:
		return f.liftInstMOVSS(inst)
	case x86asm.MOVSW:
		return f.liftInstMOVSW(inst)
	case x86asm.MOVSX:
		return f.liftInstMOVSX(inst)
	case x86asm.MOVSXD:
		return f.liftInstMOVSXD(inst)
	case x86asm.MOVUPD:
		return f.liftInstMOVUPD(inst)
	case x86asm.MOVUPS:
		return f.liftInstMOVUPS(inst)
	case x86asm.MOVZX:
		return f.liftInstMOVZX(inst)
	case x86asm.MPSADBW:
		return f.liftInstMPSADBW(inst)
	case x86asm.MUL:
		return f.liftInstMUL(inst)
	case x86asm.MULPD:
		return f.liftInstMULPD(inst)
	case x86asm.MULPS:
		return f.liftInstMULPS(inst)
	case x86asm.MULSD:
		return f.liftInstMULSD(inst)
	case x86asm.MULSS:
		return f.liftInstMULSS(inst)
	case x86asm.MWAIT:
		return f.liftInstMWAIT(inst)
	case x86asm.NEG:
		return f.liftInstNEG(inst)
	case x86asm.NOP:
		return f.liftInstNOP(inst)
	case x86asm.NOT:
		return f.liftInstNOT(inst)
	case x86asm.OR:
		return f.liftInstOR(inst)
	case x86asm.ORPD:
		return f.liftInstORPD(inst)
	case x86asm.ORPS:
		return f.liftInstORPS(inst)
	case x86asm.OUT:
		return f.liftInstOUT(inst)
	case x86asm.OUTSB:
		return f.liftInstOUTSB(inst)
	case x86asm.OUTSD:
		return f.liftInstOUTSD(inst)
	case x86asm.OUTSW:
		return f.liftInstOUTSW(inst)
	case x86asm.PABSB:
		return f.liftInstPABSB(inst)
	case x86asm.PABSD:
		return f.liftInstPABSD(inst)
	case x86asm.PABSW:
		return f.liftInstPABSW(inst)
	case x86asm.PACKSSDW:
		return f.liftInstPACKSSDW(inst)
	case x86asm.PACKSSWB:
		return f.liftInstPACKSSWB(inst)
	case x86asm.PACKUSDW:
		return f.liftInstPACKUSDW(inst)
	case x86asm.PACKUSWB:
		return f.liftInstPACKUSWB(inst)
	case x86asm.PADDB:
		return f.liftInstPADDB(inst)
	case x86asm.PADDD:
		return f.liftInstPADDD(inst)
	case x86asm.PADDQ:
		return f.liftInstPADDQ(inst)
	case x86asm.PADDSB:
		return f.liftInstPADDSB(inst)
	case x86asm.PADDSW:
		return f.liftInstPADDSW(inst)
	case x86asm.PADDUSB:
		return f.liftInstPADDUSB(inst)
	case x86asm.PADDUSW:
		return f.liftInstPADDUSW(inst)
	case x86asm.PADDW:
		return f.liftInstPADDW(inst)
	case x86asm.PALIGNR:
		return f.liftInstPALIGNR(inst)
	case x86asm.PAND:
		return f.liftInstPAND(inst)
	case x86asm.PANDN:
		return f.liftInstPANDN(inst)
	case x86asm.PAUSE:
		return f.liftInstPAUSE(inst)
	case x86asm.PAVGB:
		return f.liftInstPAVGB(inst)
	case x86asm.PAVGW:
		return f.liftInstPAVGW(inst)
	case x86asm.PBLENDVB:
		return f.liftInstPBLENDVB(inst)
	case x86asm.PBLENDW:
		return f.liftInstPBLENDW(inst)
	case x86asm.PCLMULQDQ:
		return f.liftInstPCLMULQDQ(inst)
	case x86asm.PCMPEQB:
		return f.liftInstPCMPEQB(inst)
	case x86asm.PCMPEQD:
		return f.liftInstPCMPEQD(inst)
	case x86asm.PCMPEQQ:
		return f.liftInstPCMPEQQ(inst)
	case x86asm.PCMPEQW:
		return f.liftInstPCMPEQW(inst)
	case x86asm.PCMPESTRI:
		return f.liftInstPCMPESTRI(inst)
	case x86asm.PCMPESTRM:
		return f.liftInstPCMPESTRM(inst)
	case x86asm.PCMPGTB:
		return f.liftInstPCMPGTB(inst)
	case x86asm.PCMPGTD:
		return f.liftInstPCMPGTD(inst)
	case x86asm.PCMPGTQ:
		return f.liftInstPCMPGTQ(inst)
	case x86asm.PCMPGTW:
		return f.liftInstPCMPGTW(inst)
	case x86asm.PCMPISTRI:
		return f.liftInstPCMPISTRI(inst)
	case x86asm.PCMPISTRM:
		return f.liftInstPCMPISTRM(inst)
	case x86asm.PEXTRB:
		return f.liftInstPEXTRB(inst)
	case x86asm.PEXTRD:
		return f.liftInstPEXTRD(inst)
	case x86asm.PEXTRQ:
		return f.liftInstPEXTRQ(inst)
	case x86asm.PEXTRW:
		return f.liftInstPEXTRW(inst)
	case x86asm.PHADDD:
		return f.liftInstPHADDD(inst)
	case x86asm.PHADDSW:
		return f.liftInstPHADDSW(inst)
	case x86asm.PHADDW:
		return f.liftInstPHADDW(inst)
	case x86asm.PHMINPOSUW:
		return f.liftInstPHMINPOSUW(inst)
	case x86asm.PHSUBD:
		return f.liftInstPHSUBD(inst)
	case x86asm.PHSUBSW:
		return f.liftInstPHSUBSW(inst)
	case x86asm.PHSUBW:
		return f.liftInstPHSUBW(inst)
	case x86asm.PINSRB:
		return f.liftInstPINSRB(inst)
	case x86asm.PINSRD:
		return f.liftInstPINSRD(inst)
	case x86asm.PINSRQ:
		return f.liftInstPINSRQ(inst)
	case x86asm.PINSRW:
		return f.liftInstPINSRW(inst)
	case x86asm.PMADDUBSW:
		return f.liftInstPMADDUBSW(inst)
	case x86asm.PMADDWD:
		return f.liftInstPMADDWD(inst)
	case x86asm.PMAXSB:
		return f.liftInstPMAXSB(inst)
	case x86asm.PMAXSD:
		return f.liftInstPMAXSD(inst)
	case x86asm.PMAXSW:
		return f.liftInstPMAXSW(inst)
	case x86asm.PMAXUB:
		return f.liftInstPMAXUB(inst)
	case x86asm.PMAXUD:
		return f.liftInstPMAXUD(inst)
	case x86asm.PMAXUW:
		return f.liftInstPMAXUW(inst)
	case x86asm.PMINSB:
		return f.liftInstPMINSB(inst)
	case x86asm.PMINSD:
		return f.liftInstPMINSD(inst)
	case x86asm.PMINSW:
		return f.liftInstPMINSW(inst)
	case x86asm.PMINUB:
		return f.liftInstPMINUB(inst)
	case x86asm.PMINUD:
		return f.liftInstPMINUD(inst)
	case x86asm.PMINUW:
		return f.liftInstPMINUW(inst)
	case x86asm.PMOVMSKB:
		return f.liftInstPMOVMSKB(inst)
	case x86asm.PMOVSXBD:
		return f.liftInstPMOVSXBD(inst)
	case x86asm.PMOVSXBQ:
		return f.liftInstPMOVSXBQ(inst)
	case x86asm.PMOVSXBW:
		return f.liftInstPMOVSXBW(inst)
	case x86asm.PMOVSXDQ:
		return f.liftInstPMOVSXDQ(inst)
	case x86asm.PMOVSXWD:
		return f.liftInstPMOVSXWD(inst)
	case x86asm.PMOVSXWQ:
		return f.liftInstPMOVSXWQ(inst)
	case x86asm.PMOVZXBD:
		return f.liftInstPMOVZXBD(inst)
	case x86asm.PMOVZXBQ:
		return f.liftInstPMOVZXBQ(inst)
	case x86asm.PMOVZXBW:
		return f.liftInstPMOVZXBW(inst)
	case x86asm.PMOVZXDQ:
		return f.liftInstPMOVZXDQ(inst)
	case x86asm.PMOVZXWD:
		return f.liftInstPMOVZXWD(inst)
	case x86asm.PMOVZXWQ:
		return f.liftInstPMOVZXWQ(inst)
	case x86asm.PMULDQ:
		return f.liftInstPMULDQ(inst)
	case x86asm.PMULHRSW:
		return f.liftInstPMULHRSW(inst)
	case x86asm.PMULHUW:
		return f.liftInstPMULHUW(inst)
	case x86asm.PMULHW:
		return f.liftInstPMULHW(inst)
	case x86asm.PMULLD:
		return f.liftInstPMULLD(inst)
	case x86asm.PMULLW:
		return f.liftInstPMULLW(inst)
	case x86asm.PMULUDQ:
		return f.liftInstPMULUDQ(inst)
	case x86asm.POP:
		return f.liftInstPOP(inst)
	case x86asm.POPA:
		return f.liftInstPOPA(inst)
	case x86asm.POPAD:
		return f.liftInstPOPAD(inst)
	case x86asm.POPCNT:
		return f.liftInstPOPCNT(inst)
	case x86asm.POPF:
		return f.liftInstPOPF(inst)
	case x86asm.POPFD:
		return f.liftInstPOPFD(inst)
	case x86asm.POPFQ:
		return f.liftInstPOPFQ(inst)
	case x86asm.POR:
		return f.liftInstPOR(inst)
	case x86asm.PREFETCHNTA:
		return f.liftInstPREFETCHNTA(inst)
	case x86asm.PREFETCHT0:
		return f.liftInstPREFETCHT0(inst)
	case x86asm.PREFETCHT1:
		return f.liftInstPREFETCHT1(inst)
	case x86asm.PREFETCHT2:
		return f.liftInstPREFETCHT2(inst)
	case x86asm.PREFETCHW:
		return f.liftInstPREFETCHW(inst)
	case x86asm.PSADBW:
		return f.liftInstPSADBW(inst)
	case x86asm.PSHUFB:
		return f.liftInstPSHUFB(inst)
	case x86asm.PSHUFD:
		return f.liftInstPSHUFD(inst)
	case x86asm.PSHUFHW:
		return f.liftInstPSHUFHW(inst)
	case x86asm.PSHUFLW:
		return f.liftInstPSHUFLW(inst)
	case x86asm.PSHUFW:
		return f.liftInstPSHUFW(inst)
	case x86asm.PSIGNB:
		return f.liftInstPSIGNB(inst)
	case x86asm.PSIGND:
		return f.liftInstPSIGND(inst)
	case x86asm.PSIGNW:
		return f.liftInstPSIGNW(inst)
	case x86asm.PSLLD:
		return f.liftInstPSLLD(inst)
	case x86asm.PSLLDQ:
		return f.liftInstPSLLDQ(inst)
	case x86asm.PSLLQ:
		return f.liftInstPSLLQ(inst)
	case x86asm.PSLLW:
		return f.liftInstPSLLW(inst)
	case x86asm.PSRAD:
		return f.liftInstPSRAD(inst)
	case x86asm.PSRAW:
		return f.liftInstPSRAW(inst)
	case x86asm.PSRLD:
		return f.liftInstPSRLD(inst)
	case x86asm.PSRLDQ:
		return f.liftInstPSRLDQ(inst)
	case x86asm.PSRLQ:
		return f.liftInstPSRLQ(inst)
	case x86asm.PSRLW:
		return f.liftInstPSRLW(inst)
	case x86asm.PSUBB:
		return f.liftInstPSUBB(inst)
	case x86asm.PSUBD:
		return f.liftInstPSUBD(inst)
	case x86asm.PSUBQ:
		return f.liftInstPSUBQ(inst)
	case x86asm.PSUBSB:
		return f.liftInstPSUBSB(inst)
	case x86asm.PSUBSW:
		return f.liftInstPSUBSW(inst)
	case x86asm.PSUBUSB:
		return f.liftInstPSUBUSB(inst)
	case x86asm.PSUBUSW:
		return f.liftInstPSUBUSW(inst)
	case x86asm.PSUBW:
		return f.liftInstPSUBW(inst)
	case x86asm.PTEST:
		return f.liftInstPTEST(inst)
	case x86asm.PUNPCKHBW:
		return f.liftInstPUNPCKHBW(inst)
	case x86asm.PUNPCKHDQ:
		return f.liftInstPUNPCKHDQ(inst)
	case x86asm.PUNPCKHQDQ:
		return f.liftInstPUNPCKHQDQ(inst)
	case x86asm.PUNPCKHWD:
		return f.liftInstPUNPCKHWD(inst)
	case x86asm.PUNPCKLBW:
		return f.liftInstPUNPCKLBW(inst)
	case x86asm.PUNPCKLDQ:
		return f.liftInstPUNPCKLDQ(inst)
	case x86asm.PUNPCKLQDQ:
		return f.liftInstPUNPCKLQDQ(inst)
	case x86asm.PUNPCKLWD:
		return f.liftInstPUNPCKLWD(inst)
	case x86asm.PUSH:
		return f.liftInstPUSH(inst)
	case x86asm.PUSHA:
		return f.liftInstPUSHA(inst)
	case x86asm.PUSHAD:
		return f.liftInstPUSHAD(inst)
	case x86asm.PUSHF:
		return f.liftInstPUSHF(inst)
	case x86asm.PUSHFD:
		return f.liftInstPUSHFD(inst)
	case x86asm.PUSHFQ:
		return f.liftInstPUSHFQ(inst)
	case x86asm.PXOR:
		return f.liftInstPXOR(inst)
	case x86asm.RCL:
		return f.liftInstRCL(inst)
	case x86asm.RCPPS:
		return f.liftInstRCPPS(inst)
	case x86asm.RCPSS:
		return f.liftInstRCPSS(inst)
	case x86asm.RCR:
		return f.liftInstRCR(inst)
	case x86asm.RDFSBASE:
		return f.liftInstRDFSBASE(inst)
	case x86asm.RDGSBASE:
		return f.liftInstRDGSBASE(inst)
	case x86asm.RDMSR:
		return f.liftInstRDMSR(inst)
	case x86asm.RDPMC:
		return f.liftInstRDPMC(inst)
	case x86asm.RDRAND:
		return f.liftInstRDRAND(inst)
	case x86asm.RDTSC:
		return f.liftInstRDTSC(inst)
	case x86asm.RDTSCP:
		return f.liftInstRDTSCP(inst)
	case x86asm.ROL:
		return f.liftInstROL(inst)
	case x86asm.ROR:
		return f.liftInstROR(inst)
	case x86asm.ROUNDPD:
		return f.liftInstROUNDPD(inst)
	case x86asm.ROUNDPS:
		return f.liftInstROUNDPS(inst)
	case x86asm.ROUNDSD:
		return f.liftInstROUNDSD(inst)
	case x86asm.ROUNDSS:
		return f.liftInstROUNDSS(inst)
	case x86asm.RSM:
		return f.liftInstRSM(inst)
	case x86asm.RSQRTPS:
		return f.liftInstRSQRTPS(inst)
	case x86asm.RSQRTSS:
		return f.liftInstRSQRTSS(inst)
	case x86asm.SAHF:
		return f.liftInstSAHF(inst)
	case x86asm.SAR:
		return f.liftInstSAR(inst)
	case x86asm.SBB:
		return f.liftInstSBB(inst)
	case x86asm.SCASB:
		return f.liftInstSCASB(inst)
	case x86asm.SCASD:
		return f.liftInstSCASD(inst)
	case x86asm.SCASQ:
		return f.liftInstSCASQ(inst)
	case x86asm.SCASW:
		return f.liftInstSCASW(inst)
	case x86asm.SETA:
		return f.liftInstSETA(inst)
	case x86asm.SETAE:
		return f.liftInstSETAE(inst)
	case x86asm.SETB:
		return f.liftInstSETB(inst)
	case x86asm.SETBE:
		return f.liftInstSETBE(inst)
	case x86asm.SETE:
		return f.liftInstSETE(inst)
	case x86asm.SETG:
		return f.liftInstSETG(inst)
	case x86asm.SETGE:
		return f.liftInstSETGE(inst)
	case x86asm.SETL:
		return f.liftInstSETL(inst)
	case x86asm.SETLE:
		return f.liftInstSETLE(inst)
	case x86asm.SETNE:
		return f.liftInstSETNE(inst)
	case x86asm.SETNO:
		return f.liftInstSETNO(inst)
	case x86asm.SETNP:
		return f.liftInstSETNP(inst)
	case x86asm.SETNS:
		return f.liftInstSETNS(inst)
	case x86asm.SETO:
		return f.liftInstSETO(inst)
	case x86asm.SETP:
		return f.liftInstSETP(inst)
	case x86asm.SETS:
		return f.liftInstSETS(inst)
	case x86asm.SFENCE:
		return f.liftInstSFENCE(inst)
	case x86asm.SGDT:
		return f.liftInstSGDT(inst)
	case x86asm.SHL:
		return f.liftInstSHL(inst)
	case x86asm.SHLD:
		return f.liftInstSHLD(inst)
	case x86asm.SHR:
		return f.liftInstSHR(inst)
	case x86asm.SHRD:
		return f.liftInstSHRD(inst)
	case x86asm.SHUFPD:
		return f.liftInstSHUFPD(inst)
	case x86asm.SHUFPS:
		return f.liftInstSHUFPS(inst)
	case x86asm.SIDT:
		return f.liftInstSIDT(inst)
	case x86asm.SLDT:
		return f.liftInstSLDT(inst)
	case x86asm.SMSW:
		return f.liftInstSMSW(inst)
	case x86asm.SQRTPD:
		return f.liftInstSQRTPD(inst)
	case x86asm.SQRTPS:
		return f.liftInstSQRTPS(inst)
	case x86asm.SQRTSD:
		return f.liftInstSQRTSD(inst)
	case x86asm.SQRTSS:
		return f.liftInstSQRTSS(inst)
	case x86asm.STC:
		return f.liftInstSTC(inst)
	case x86asm.STD:
		return f.liftInstSTD(inst)
	case x86asm.STI:
		return f.liftInstSTI(inst)
	case x86asm.STMXCSR:
		return f.liftInstSTMXCSR(inst)
	case x86asm.STOSB:
		return f.liftInstSTOSB(inst)
	case x86asm.STOSD:
		return f.liftInstSTOSD(inst)
	case x86asm.STOSQ:
		return f.liftInstSTOSQ(inst)
	case x86asm.STOSW:
		return f.liftInstSTOSW(inst)
	case x86asm.STR:
		return f.liftInstSTR(inst)
	case x86asm.SUB:
		return f.liftInstSUB(inst)
	case x86asm.SUBPD:
		return f.liftInstSUBPD(inst)
	case x86asm.SUBPS:
		return f.liftInstSUBPS(inst)
	case x86asm.SUBSD:
		return f.liftInstSUBSD(inst)
	case x86asm.SUBSS:
		return f.liftInstSUBSS(inst)
	case x86asm.SWAPGS:
		return f.liftInstSWAPGS(inst)
	case x86asm.SYSCALL:
		return f.liftInstSYSCALL(inst)
	case x86asm.SYSENTER:
		return f.liftInstSYSENTER(inst)
	case x86asm.SYSEXIT:
		return f.liftInstSYSEXIT(inst)
	case x86asm.SYSRET:
		return f.liftInstSYSRET(inst)
	case x86asm.TEST:
		return f.liftInstTEST(inst)
	case x86asm.TZCNT:
		return f.liftInstTZCNT(inst)
	case x86asm.UCOMISD:
		return f.liftInstUCOMISD(inst)
	case x86asm.UCOMISS:
		return f.liftInstUCOMISS(inst)
	case x86asm.UD1:
		return f.liftInstUD1(inst)
	case x86asm.UD2:
		return f.liftInstUD2(inst)
	case x86asm.UNPCKHPD:
		return f.liftInstUNPCKHPD(inst)
	case x86asm.UNPCKHPS:
		return f.liftInstUNPCKHPS(inst)
	case x86asm.UNPCKLPD:
		return f.liftInstUNPCKLPD(inst)
	case x86asm.UNPCKLPS:
		return f.liftInstUNPCKLPS(inst)
	case x86asm.VERR:
		return f.liftInstVERR(inst)
	case x86asm.VERW:
		return f.liftInstVERW(inst)
	case x86asm.VMOVDQA:
		return f.liftInstVMOVDQA(inst)
	case x86asm.VMOVDQU:
		return f.liftInstVMOVDQU(inst)
	case x86asm.VMOVNTDQ:
		return f.liftInstVMOVNTDQ(inst)
	case x86asm.VMOVNTDQA:
		return f.liftInstVMOVNTDQA(inst)
	case x86asm.VZEROUPPER:
		return f.liftInstVZEROUPPER(inst)
	case x86asm.WBINVD:
		return f.liftInstWBINVD(inst)
	case x86asm.WRFSBASE:
		return f.liftInstWRFSBASE(inst)
	case x86asm.WRGSBASE:
		return f.liftInstWRGSBASE(inst)
	case x86asm.WRMSR:
		return f.liftInstWRMSR(inst)
	case x86asm.XABORT:
		return f.liftInstXABORT(inst)
	case x86asm.XADD:
		return f.liftInstXADD(inst)
	case x86asm.XBEGIN:
		return f.liftInstXBEGIN(inst)
	case x86asm.XCHG:
		return f.liftInstXCHG(inst)
	case x86asm.XEND:
		return f.liftInstXEND(inst)
	case x86asm.XGETBV:
		return f.liftInstXGETBV(inst)
	case x86asm.XLATB:
		return f.liftInstXLATB(inst)
	case x86asm.XOR:
		return f.liftInstXOR(inst)
	case x86asm.XORPD:
		return f.liftInstXORPD(inst)
	case x86asm.XORPS:
		return f.liftInstXORPS(inst)
	case x86asm.XRSTOR:
		return f.liftInstXRSTOR(inst)
	case x86asm.XRSTOR64:
		return f.liftInstXRSTOR64(inst)
	case x86asm.XRSTORS:
		return f.liftInstXRSTORS(inst)
	case x86asm.XRSTORS64:
		return f.liftInstXRSTORS64(inst)
	case x86asm.XSAVE:
		return f.liftInstXSAVE(inst)
	case x86asm.XSAVE64:
		return f.liftInstXSAVE64(inst)
	case x86asm.XSAVEC:
		return f.liftInstXSAVEC(inst)
	case x86asm.XSAVEC64:
		return f.liftInstXSAVEC64(inst)
	case x86asm.XSAVEOPT:
		return f.liftInstXSAVEOPT(inst)
	case x86asm.XSAVEOPT64:
		return f.liftInstXSAVEOPT64(inst)
	case x86asm.XSAVES:
		return f.liftInstXSAVES(inst)
	case x86asm.XSAVES64:
		return f.liftInstXSAVES64(inst)
	case x86asm.XSETBV:
		return f.liftInstXSETBV(inst)
	case x86asm.XTEST:
		return f.liftInstXTEST(inst)
	default:
		panic(fmt.Errorf("support for x86 instruction opcode %v not yet implemented", inst.Op))
	}
}

// --- [ AAA ] -----------------------------------------------------------------

// liftInstAAA lifts the given x86 AAA instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstAAA(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAAA: not yet implemented")
}

// --- [ AAD ] -----------------------------------------------------------------

// liftInstAAD lifts the given x86 AAD instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstAAD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAAD: not yet implemented")
}

// --- [ AAM ] -----------------------------------------------------------------

// liftInstAAM lifts the given x86 AAM instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstAAM(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAAM: not yet implemented")
}

// --- [ AAS ] -----------------------------------------------------------------

// liftInstAAS lifts the given x86 AAS instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstAAS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAAS: not yet implemented")
}

// --- [ ADC ] -----------------------------------------------------------------

// liftInstADC lifts the given x86 ADC instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstADC(inst *x86.Inst) error {
	// ADC - Add with Carry.
	//
	//    ADC AL, imm8        Add with carry imm8 to AL.
	//    ADC AX, imm16       Add with carry imm16 to AX.
	//    ADC EAX, imm32      Add with carry imm32 to EAX.
	//    ADC r/m8, imm8      Add with carry imm8 to r/m8.
	//    ADC r/m8, r8        Add with carry byte register to r/m8.
	//    ADC r/m8, r8        Add with carry byte register to r/m64.
	//    ADC r/m16, imm16    Add with carry imm16 to r/m16.
	//    ADC r/m16, imm8     Add with CF sign-extended imm8 to r/m16.
	//    ADC r/m16, r16      Add with carry r16 to r/m16.
	//    ADC r/m32, imm32    Add with CF imm32 to r/m32.
	//    ADC r/m32, imm8     Add with CF sign-extended imm8 into r/m32.
	//    ADC r/m32, r32      Add with CF r32 to r/m32.
	//    ADC r/m64, imm32    Add with CF imm32 sign extended to 64-bits to r/m64.
	//    ADC r/m64, imm8     Add with CF sign-extended imm8 into r/m64.
	//    ADC r/m64, r64      Add with CF r64 to r/m64.
	//    ADC r16, r/m16      Add with carry r/m16 to r16.
	//    ADC r32, r/m32      Add with CF r/m32 to r32.
	//    ADC r64, r/m64      Add with CF r/m64 to r64.
	//    ADC r8, r/m8        Add with carry r/m8 to byte register.
	//    ADC r8, r/m8        Add with carry r/m64 to byte register.
	//    ADC RAX, imm32      Add with carry imm32 sign extended to 64-bits to RAX.
	//
	// Adds the destination operand (first operand), the source operand (second
	// operand), and the carry (CF) flag and stores the result in the destination
	// operand.
	dst := f.useArg(inst.Arg(0))
	src := f.useArg(inst.Arg(1))
	cf := f.useStatus(CF)
	v := f.cur.NewAdd(src, cf)
	result := f.cur.NewAdd(dst, v)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ ADD ] -----------------------------------------------------------------

// liftInstADD lifts the given x86 ADD instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstADD(inst *x86.Inst) error {
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewAdd(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ ADDPD ] ---------------------------------------------------------------

// liftInstADDPD lifts the given x86 ADDPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstADDPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDPD: not yet implemented")
}

// --- [ ADDPS ] ---------------------------------------------------------------

// liftInstADDPS lifts the given x86 ADDPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstADDPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDPS: not yet implemented")
}

// --- [ ADDSD ] ---------------------------------------------------------------

// liftInstADDSD lifts the given x86 ADDSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstADDSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDSD: not yet implemented")
}

// --- [ ADDSS ] ---------------------------------------------------------------

// liftInstADDSS lifts the given x86 ADDSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstADDSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDSS: not yet implemented")
}

// --- [ ADDSUBPD ] ------------------------------------------------------------

// liftInstADDSUBPD lifts the given x86 ADDSUBPD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstADDSUBPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDSUBPD: not yet implemented")
}

// --- [ ADDSUBPS ] ------------------------------------------------------------

// liftInstADDSUBPS lifts the given x86 ADDSUBPS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstADDSUBPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstADDSUBPS: not yet implemented")
}

// --- [ AESDEC ] --------------------------------------------------------------

// liftInstAESDEC lifts the given x86 AESDEC instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstAESDEC(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESDEC: not yet implemented")
}

// --- [ AESDECLAST ] ----------------------------------------------------------

// liftInstAESDECLAST lifts the given x86 AESDECLAST instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstAESDECLAST(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESDECLAST: not yet implemented")
}

// --- [ AESENC ] --------------------------------------------------------------

// liftInstAESENC lifts the given x86 AESENC instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstAESENC(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESENC: not yet implemented")
}

// --- [ AESENCLAST ] ----------------------------------------------------------

// liftInstAESENCLAST lifts the given x86 AESENCLAST instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstAESENCLAST(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESENCLAST: not yet implemented")
}

// --- [ AESIMC ] --------------------------------------------------------------

// liftInstAESIMC lifts the given x86 AESIMC instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstAESIMC(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESIMC: not yet implemented")
}

// --- [ AESKEYGENASSIST ] -----------------------------------------------------

// liftInstAESKEYGENASSIST lifts the given x86 AESKEYGENASSIST instruction to
// LLVM IR, emitting code to f.
func (f *Func) liftInstAESKEYGENASSIST(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstAESKEYGENASSIST: not yet implemented")
}

// --- [ AND ] -----------------------------------------------------------------

// liftInstAND lifts the given x86 AND instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstAND(inst *x86.Inst) error {
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewAnd(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ ANDNPD ] --------------------------------------------------------------

// liftInstANDNPD lifts the given x86 ANDNPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstANDNPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstANDNPD: not yet implemented")
}

// --- [ ANDNPS ] --------------------------------------------------------------

// liftInstANDNPS lifts the given x86 ANDNPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstANDNPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstANDNPS: not yet implemented")
}

// --- [ ANDPD ] ---------------------------------------------------------------

// liftInstANDPD lifts the given x86 ANDPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstANDPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstANDPD: not yet implemented")
}

// --- [ ANDPS ] ---------------------------------------------------------------

// liftInstANDPS lifts the given x86 ANDPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstANDPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstANDPS: not yet implemented")
}

// --- [ ARPL ] ----------------------------------------------------------------

// liftInstARPL lifts the given x86 ARPL instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstARPL(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstARPL: not yet implemented")
}

// --- [ BLENDPD ] -------------------------------------------------------------

// liftInstBLENDPD lifts the given x86 BLENDPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstBLENDPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBLENDPD: not yet implemented")
}

// --- [ BLENDPS ] -------------------------------------------------------------

// liftInstBLENDPS lifts the given x86 BLENDPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstBLENDPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBLENDPS: not yet implemented")
}

// --- [ BLENDVPD ] ------------------------------------------------------------

// liftInstBLENDVPD lifts the given x86 BLENDVPD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstBLENDVPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBLENDVPD: not yet implemented")
}

// --- [ BLENDVPS ] ------------------------------------------------------------

// liftInstBLENDVPS lifts the given x86 BLENDVPS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstBLENDVPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBLENDVPS: not yet implemented")
}

// --- [ BOUND ] ---------------------------------------------------------------

// liftInstBOUND lifts the given x86 BOUND instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstBOUND(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBOUND: not yet implemented")
}

// --- [ BSF ] -----------------------------------------------------------------

// liftInstBSF lifts the given x86 BSF instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstBSF(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBSF: not yet implemented")
}

// --- [ BSR ] -----------------------------------------------------------------

// liftInstBSR lifts the given x86 BSR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstBSR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBSR: not yet implemented")
}

// --- [ BSWAP ] ---------------------------------------------------------------

// liftInstBSWAP lifts the given x86 BSWAP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstBSWAP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBSWAP: not yet implemented")
}

// --- [ BT ] ------------------------------------------------------------------

// liftInstBT lifts the given x86 BT instruction to LLVM IR, emitting code to f.
func (f *Func) liftInstBT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBT: not yet implemented")
}

// --- [ BTC ] -----------------------------------------------------------------

// liftInstBTC lifts the given x86 BTC instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstBTC(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBTC: not yet implemented")
}

// --- [ BTR ] -----------------------------------------------------------------

// liftInstBTR lifts the given x86 BTR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstBTR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBTR: not yet implemented")
}

// --- [ BTS ] -----------------------------------------------------------------

// liftInstBTS lifts the given x86 BTS instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstBTS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstBTS: not yet implemented")
}

// --- [ CALL ] ----------------------------------------------------------------

// liftInstCALL lifts the given x86 CALL instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCALL(inst *x86.Inst) error {
	// Locate callee information.
	callee, sig, callconv, ok := f.getFunc(inst.Arg(0))
	if !ok {
		panic(fmt.Errorf("unable to locate function for argument %v of instruction at address %v", inst.Arg(0), inst.Addr))
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
				arg := f.useReg(x86.ECX)
				args = append(args, arg)
				continue
			case 1:
				arg := f.useReg(x86.EDX)
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
		f.defReg(x86.EAX, result)
	}
	return nil
}

// --- [ CBW ] -----------------------------------------------------------------

// liftInstCBW lifts the given x86 CBW instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstCBW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCBW: not yet implemented")
}

// --- [ CDQ ] -----------------------------------------------------------------

// liftInstCDQ lifts the given x86 CDQ instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstCDQ(inst *x86.Inst) error {
	// EDX:EAX = sign-extend of EAX.
	eax := f.useReg(x86.EAX)
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
	f.defReg(x86.EDX, constant.NewInt(0xFFFFFFFF, types.I32))
	f.cur = targetFalse
	f.defReg(x86.EDX, constant.NewInt(0, types.I32))
	targetTrue.NewBr(exit)
	targetFalse.NewBr(exit)
	f.cur = exit
	return nil
}

// --- [ CDQE ] ----------------------------------------------------------------

// liftInstCDQE lifts the given x86 CDQE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCDQE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCDQE: not yet implemented")
}

// --- [ CLC ] -----------------------------------------------------------------

// liftInstCLC lifts the given x86 CLC instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstCLC(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCLC: not yet implemented")
}

// --- [ CLD ] -----------------------------------------------------------------

// liftInstCLD lifts the given x86 CLD instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstCLD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCLD: not yet implemented")
}

// --- [ CLFLUSH ] -------------------------------------------------------------

// liftInstCLFLUSH lifts the given x86 CLFLUSH instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCLFLUSH(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCLFLUSH: not yet implemented")
}

// --- [ CLI ] -----------------------------------------------------------------

// liftInstCLI lifts the given x86 CLI instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstCLI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCLI: not yet implemented")
}

// --- [ CLTS ] ----------------------------------------------------------------

// liftInstCLTS lifts the given x86 CLTS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCLTS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCLTS: not yet implemented")
}

// --- [ CMC ] -----------------------------------------------------------------

// liftInstCMC lifts the given x86 CMC instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstCMC(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMC: not yet implemented")
}

// --- [ CMOVA ] ---------------------------------------------------------------

// liftInstCMOVA lifts the given x86 CMOVA instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMOVA(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVA: not yet implemented")
}

// --- [ CMOVAE ] --------------------------------------------------------------

// liftInstCMOVAE lifts the given x86 CMOVAE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCMOVAE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVAE: not yet implemented")
}

// --- [ CMOVB ] ---------------------------------------------------------------

// liftInstCMOVB lifts the given x86 CMOVB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMOVB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVB: not yet implemented")
}

// --- [ CMOVBE ] --------------------------------------------------------------

// liftInstCMOVBE lifts the given x86 CMOVBE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCMOVBE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVBE: not yet implemented")
}

// --- [ CMOVE ] ---------------------------------------------------------------

// liftInstCMOVE lifts the given x86 CMOVE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMOVE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVE: not yet implemented")
}

// --- [ CMOVG ] ---------------------------------------------------------------

// liftInstCMOVG lifts the given x86 CMOVG instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMOVG(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVG: not yet implemented")
}

// --- [ CMOVGE ] --------------------------------------------------------------

// liftInstCMOVGE lifts the given x86 CMOVGE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCMOVGE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVGE: not yet implemented")
}

// --- [ CMOVL ] ---------------------------------------------------------------

// liftInstCMOVL lifts the given x86 CMOVL instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMOVL(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVL: not yet implemented")
}

// --- [ CMOVLE ] --------------------------------------------------------------

// liftInstCMOVLE lifts the given x86 CMOVLE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCMOVLE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVLE: not yet implemented")
}

// --- [ CMOVNE ] --------------------------------------------------------------

// liftInstCMOVNE lifts the given x86 CMOVNE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCMOVNE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVNE: not yet implemented")
}

// --- [ CMOVNO ] --------------------------------------------------------------

// liftInstCMOVNO lifts the given x86 CMOVNO instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCMOVNO(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVNO: not yet implemented")
}

// --- [ CMOVNP ] --------------------------------------------------------------

// liftInstCMOVNP lifts the given x86 CMOVNP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCMOVNP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVNP: not yet implemented")
}

// --- [ CMOVNS ] --------------------------------------------------------------

// liftInstCMOVNS lifts the given x86 CMOVNS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCMOVNS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVNS: not yet implemented")
}

// --- [ CMOVO ] ---------------------------------------------------------------

// liftInstCMOVO lifts the given x86 CMOVO instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMOVO(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVO: not yet implemented")
}

// --- [ CMOVP ] ---------------------------------------------------------------

// liftInstCMOVP lifts the given x86 CMOVP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMOVP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVP: not yet implemented")
}

// --- [ CMOVS ] ---------------------------------------------------------------

// liftInstCMOVS lifts the given x86 CMOVS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMOVS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMOVS: not yet implemented")
}

// --- [ CMP ] -----------------------------------------------------------------

// liftInstCMP lifts the given x86 CMP instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstCMP(inst *x86.Inst) error {
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

// liftInstCMPPD lifts the given x86 CMPPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMPPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPPD: not yet implemented")
}

// --- [ CMPPS ] ---------------------------------------------------------------

// liftInstCMPPS lifts the given x86 CMPPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMPPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPPS: not yet implemented")
}

// --- [ CMPSB ] ---------------------------------------------------------------

// liftInstCMPSB lifts the given x86 CMPSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMPSB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSB: not yet implemented")
}

// --- [ CMPSD ] ---------------------------------------------------------------

// liftInstCMPSD lifts the given x86 CMPSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMPSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSD: not yet implemented")
}

// --- [ CMPSD_XMM ] -----------------------------------------------------------

// liftInstCMPSD_XMM lifts the given x86 CMPSD_XMM instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCMPSD_XMM(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSD_XMM: not yet implemented")
}

// --- [ CMPSQ ] ---------------------------------------------------------------

// liftInstCMPSQ lifts the given x86 CMPSQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMPSQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSQ: not yet implemented")
}

// --- [ CMPSS ] ---------------------------------------------------------------

// liftInstCMPSS lifts the given x86 CMPSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMPSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSS: not yet implemented")
}

// --- [ CMPSW ] ---------------------------------------------------------------

// liftInstCMPSW lifts the given x86 CMPSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCMPSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPSW: not yet implemented")
}

// --- [ CMPXCHG ] -------------------------------------------------------------

// liftInstCMPXCHG lifts the given x86 CMPXCHG instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCMPXCHG(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPXCHG: not yet implemented")
}

// --- [ CMPXCHG16B ] ----------------------------------------------------------

// liftInstCMPXCHG16B lifts the given x86 CMPXCHG16B instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCMPXCHG16B(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPXCHG16B: not yet implemented")
}

// --- [ CMPXCHG8B ] -----------------------------------------------------------

// liftInstCMPXCHG8B lifts the given x86 CMPXCHG8B instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCMPXCHG8B(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCMPXCHG8B: not yet implemented")
}

// --- [ COMISD ] --------------------------------------------------------------

// liftInstCOMISD lifts the given x86 COMISD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCOMISD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCOMISD: not yet implemented")
}

// --- [ COMISS ] --------------------------------------------------------------

// liftInstCOMISS lifts the given x86 COMISS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstCOMISS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCOMISS: not yet implemented")
}

// --- [ CPUID ] ---------------------------------------------------------------

// liftInstCPUID lifts the given x86 CPUID instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCPUID(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCPUID: not yet implemented")
}

// --- [ CQO ] -----------------------------------------------------------------

// liftInstCQO lifts the given x86 CQO instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstCQO(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCQO: not yet implemented")
}

// --- [ CRC32 ] ---------------------------------------------------------------

// liftInstCRC32 lifts the given x86 CRC32 instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCRC32(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCRC32: not yet implemented")
}

// --- [ CVTDQ2PD ] ------------------------------------------------------------

// liftInstCVTDQ2PD lifts the given x86 CVTDQ2PD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTDQ2PD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTDQ2PD: not yet implemented")
}

// --- [ CVTDQ2PS ] ------------------------------------------------------------

// liftInstCVTDQ2PS lifts the given x86 CVTDQ2PS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTDQ2PS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTDQ2PS: not yet implemented")
}

// --- [ CVTPD2DQ ] ------------------------------------------------------------

// liftInstCVTPD2DQ lifts the given x86 CVTPD2DQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTPD2DQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPD2DQ: not yet implemented")
}

// --- [ CVTPD2PI ] ------------------------------------------------------------

// liftInstCVTPD2PI lifts the given x86 CVTPD2PI instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTPD2PI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPD2PI: not yet implemented")
}

// --- [ CVTPD2PS ] ------------------------------------------------------------

// liftInstCVTPD2PS lifts the given x86 CVTPD2PS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTPD2PS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPD2PS: not yet implemented")
}

// --- [ CVTPI2PD ] ------------------------------------------------------------

// liftInstCVTPI2PD lifts the given x86 CVTPI2PD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTPI2PD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPI2PD: not yet implemented")
}

// --- [ CVTPI2PS ] ------------------------------------------------------------

// liftInstCVTPI2PS lifts the given x86 CVTPI2PS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTPI2PS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPI2PS: not yet implemented")
}

// --- [ CVTPS2DQ ] ------------------------------------------------------------

// liftInstCVTPS2DQ lifts the given x86 CVTPS2DQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTPS2DQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPS2DQ: not yet implemented")
}

// --- [ CVTPS2PD ] ------------------------------------------------------------

// liftInstCVTPS2PD lifts the given x86 CVTPS2PD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTPS2PD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPS2PD: not yet implemented")
}

// --- [ CVTPS2PI ] ------------------------------------------------------------

// liftInstCVTPS2PI lifts the given x86 CVTPS2PI instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTPS2PI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTPS2PI: not yet implemented")
}

// --- [ CVTSD2SI ] ------------------------------------------------------------

// liftInstCVTSD2SI lifts the given x86 CVTSD2SI instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTSD2SI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSD2SI: not yet implemented")
}

// --- [ CVTSD2SS ] ------------------------------------------------------------

// liftInstCVTSD2SS lifts the given x86 CVTSD2SS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTSD2SS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSD2SS: not yet implemented")
}

// --- [ CVTSI2SD ] ------------------------------------------------------------

// liftInstCVTSI2SD lifts the given x86 CVTSI2SD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTSI2SD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSI2SD: not yet implemented")
}

// --- [ CVTSI2SS ] ------------------------------------------------------------

// liftInstCVTSI2SS lifts the given x86 CVTSI2SS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTSI2SS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSI2SS: not yet implemented")
}

// --- [ CVTSS2SD ] ------------------------------------------------------------

// liftInstCVTSS2SD lifts the given x86 CVTSS2SD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTSS2SD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSS2SD: not yet implemented")
}

// --- [ CVTSS2SI ] ------------------------------------------------------------

// liftInstCVTSS2SI lifts the given x86 CVTSS2SI instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTSS2SI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTSS2SI: not yet implemented")
}

// --- [ CVTTPD2DQ ] -----------------------------------------------------------

// liftInstCVTTPD2DQ lifts the given x86 CVTTPD2DQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTTPD2DQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTPD2DQ: not yet implemented")
}

// --- [ CVTTPD2PI ] -----------------------------------------------------------

// liftInstCVTTPD2PI lifts the given x86 CVTTPD2PI instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTTPD2PI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTPD2PI: not yet implemented")
}

// --- [ CVTTPS2DQ ] -----------------------------------------------------------

// liftInstCVTTPS2DQ lifts the given x86 CVTTPS2DQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTTPS2DQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTPS2DQ: not yet implemented")
}

// --- [ CVTTPS2PI ] -----------------------------------------------------------

// liftInstCVTTPS2PI lifts the given x86 CVTTPS2PI instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTTPS2PI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTPS2PI: not yet implemented")
}

// --- [ CVTTSD2SI ] -----------------------------------------------------------

// liftInstCVTTSD2SI lifts the given x86 CVTTSD2SI instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTTSD2SI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTSD2SI: not yet implemented")
}

// --- [ CVTTSS2SI ] -----------------------------------------------------------

// liftInstCVTTSS2SI lifts the given x86 CVTTSS2SI instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstCVTTSS2SI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCVTTSS2SI: not yet implemented")
}

// --- [ CWD ] -----------------------------------------------------------------

// liftInstCWD lifts the given x86 CWD instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstCWD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCWD: not yet implemented")
}

// --- [ CWDE ] ----------------------------------------------------------------

// liftInstCWDE lifts the given x86 CWDE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstCWDE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstCWDE: not yet implemented")
}

// --- [ DAA ] -----------------------------------------------------------------

// liftInstDAA lifts the given x86 DAA instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstDAA(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDAA: not yet implemented")
}

// --- [ DAS ] -----------------------------------------------------------------

// liftInstDAS lifts the given x86 DAS instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstDAS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDAS: not yet implemented")
}

// --- [ DEC ] -----------------------------------------------------------------

// liftInstDEC lifts the given x86 DEC instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstDEC(inst *x86.Inst) error {
	f.dec(inst.Arg(0))
	return nil
}

// dec decrements the given argument by 1, stores and returns the result.
func (f *Func) dec(arg *x86.Arg) value.Value {
	x := f.useArg(arg)
	one := constant.NewInt(1, x.Type())
	result := f.cur.NewSub(x, one)
	f.defArg(arg, result)
	return result
}

// --- [ DIV ] -----------------------------------------------------------------

// liftInstDIV lifts the given x86 DIV instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstDIV(inst *x86.Inst) error {
	// DIV - Unsigned Divide
	//
	//    div arg
	arg := f.useArg(inst.Arg(0))
	typ, ok := arg.Type().(*types.IntType)
	if !ok {
		return errors.Errorf("invalid argument type in instruction %v; expected *types.IntType, got %T", inst, arg.Type())
	}
	switch typ.Size {
	case 8:
		// Unsigned divide AX by r/m8, with result stored in:
		//
		//    AL = quotient
		//    AH = remainder
		ax := f.useReg(x86.AX)
		arg = f.cur.NewZExt(arg, types.I16)
		quo := f.cur.NewUDiv(ax, arg)
		rem := f.cur.NewURem(ax, arg)
		f.defReg(x86.AL, quo)
		f.defReg(x86.AH, rem)
	case 16:
		// Unsigned divide DX:AX by r/m16, with result stored in:
		//
		//    AX = quotient
		//    DX = remainder
		dx_ax := f.useReg(x86.DX_AX)
		arg = f.cur.NewZExt(arg, types.I32)
		quo := f.cur.NewUDiv(dx_ax, arg)
		rem := f.cur.NewURem(dx_ax, arg)
		f.defReg(x86.AX, quo)
		f.defReg(x86.DX, rem)
	case 32:
		// Unsigned divide EDX:EAX by r/m32, with result stored in:
		//
		//    EAX = quotient
		//    EDX = remainder
		edx_eax := f.useReg(x86.EDX_EAX)
		arg = f.cur.NewZExt(arg, types.I64)
		quo := f.cur.NewUDiv(edx_eax, arg)
		rem := f.cur.NewURem(edx_eax, arg)
		f.defReg(x86.EAX, quo)
		f.defReg(x86.EDX, rem)
	case 64:
		// Unsigned divide RDX:RAX by r/m64, with result stored in:
		//
		//    RAX = quotient
		//    RDX = remainder
		rdx_rax := f.useReg(x86.RDX_RAX)
		arg = f.cur.NewZExt(arg, types.I128)
		quo := f.cur.NewUDiv(rdx_rax, arg)
		rem := f.cur.NewURem(rdx_rax, arg)
		f.defReg(x86.RAX, quo)
		f.defReg(x86.RDX, rem)
	default:
		panic(fmt.Errorf("support for argument bit size %d not yet implemented", typ.Size))
	}
	return nil
}

// --- [ DIVPD ] ---------------------------------------------------------------

// liftInstDIVPD lifts the given x86 DIVPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstDIVPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDIVPD: not yet implemented")
}

// --- [ DIVPS ] ---------------------------------------------------------------

// liftInstDIVPS lifts the given x86 DIVPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstDIVPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDIVPS: not yet implemented")
}

// --- [ DIVSD ] ---------------------------------------------------------------

// liftInstDIVSD lifts the given x86 DIVSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstDIVSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDIVSD: not yet implemented")
}

// --- [ DIVSS ] ---------------------------------------------------------------

// liftInstDIVSS lifts the given x86 DIVSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstDIVSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDIVSS: not yet implemented")
}

// --- [ DPPD ] ----------------------------------------------------------------

// liftInstDPPD lifts the given x86 DPPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstDPPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDPPD: not yet implemented")
}

// --- [ DPPS ] ----------------------------------------------------------------

// liftInstDPPS lifts the given x86 DPPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstDPPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstDPPS: not yet implemented")
}

// --- [ EMMS ] ----------------------------------------------------------------

// liftInstEMMS lifts the given x86 EMMS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstEMMS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstEMMS: not yet implemented")
}

// --- [ ENTER ] ---------------------------------------------------------------

// liftInstENTER lifts the given x86 ENTER instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstENTER(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstENTER: not yet implemented")
}

// --- [ EXTRACTPS ] -----------------------------------------------------------

// liftInstEXTRACTPS lifts the given x86 EXTRACTPS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstEXTRACTPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstEXTRACTPS: not yet implemented")
}

// --- [ FFREEP ] --------------------------------------------------------------

// liftInstFFREEP lifts the given x86 FFREEP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFFREEP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFFREEP: not yet implemented")
}

// --- [ FISTTP ] --------------------------------------------------------------

// liftInstFISTTP lifts the given x86 FISTTP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFISTTP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFISTTP: not yet implemented")
}

// --- [ FXRSTOR ] -------------------------------------------------------------

// liftInstFXRSTOR lifts the given x86 FXRSTOR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFXRSTOR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXRSTOR: not yet implemented")
}

// --- [ FXRSTOR64 ] -----------------------------------------------------------

// liftInstFXRSTOR64 lifts the given x86 FXRSTOR64 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstFXRSTOR64(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXRSTOR64: not yet implemented")
}

// --- [ FXSAVE ] --------------------------------------------------------------

// liftInstFXSAVE lifts the given x86 FXSAVE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstFXSAVE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXSAVE: not yet implemented")
}

// --- [ FXSAVE64 ] ------------------------------------------------------------

// liftInstFXSAVE64 lifts the given x86 FXSAVE64 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstFXSAVE64(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstFXSAVE64: not yet implemented")
}

// --- [ HADDPD ] --------------------------------------------------------------

// liftInstHADDPD lifts the given x86 HADDPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstHADDPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstHADDPD: not yet implemented")
}

// --- [ HADDPS ] --------------------------------------------------------------

// liftInstHADDPS lifts the given x86 HADDPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstHADDPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstHADDPS: not yet implemented")
}

// --- [ HLT ] -----------------------------------------------------------------

// liftInstHLT lifts the given x86 HLT instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstHLT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstHLT: not yet implemented")
}

// --- [ HSUBPD ] --------------------------------------------------------------

// liftInstHSUBPD lifts the given x86 HSUBPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstHSUBPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstHSUBPD: not yet implemented")
}

// --- [ HSUBPS ] --------------------------------------------------------------

// liftInstHSUBPS lifts the given x86 HSUBPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstHSUBPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstHSUBPS: not yet implemented")
}

// --- [ ICEBP ] ---------------------------------------------------------------

// liftInstICEBP lifts the given x86 ICEBP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstICEBP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstICEBP: not yet implemented")
}

// --- [ IDIV ] ----------------------------------------------------------------

// liftInstIDIV lifts the given x86 IDIV instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstIDIV(inst *x86.Inst) error {
	// IDIV - Signed Divide

	// Signed divide EDX:EAX by r/m32, with result stored in:
	//
	//    EAX = Quotient
	//    EDX = Remainder
	x := f.useArg(inst.Arg(0))
	edx_eax := f.useReg(x86.EDX_EAX)
	if !types.Equal(x.Type(), types.I64) {
		x = f.cur.NewSExt(x, types.I64)
	}
	quo := f.cur.NewSDiv(edx_eax, x)
	rem := f.cur.NewSRem(edx_eax, x)
	f.defReg(x86.EAX, quo)
	f.defReg(x86.EDX, rem)
	return nil
}

// --- [ IMUL ] ----------------------------------------------------------------

// liftInstIMUL lifts the given x86 IMUL instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstIMUL(inst *x86.Inst) error {
	// IMUL - Signed Multiply
	//
	//    IMUL r/m8                     AX = AL * r/m byte.
	//    IMUL r/m16                    DX:AX = AX * r/m word.
	//    IMUL r/m32                    EDX:EAX = EAX * r/m32.
	//    IMUL r/m64                    RDX:RAX = RAX * r/m64.
	//    IMUL r16, r/m16               Word register = word register * r/m16.
	//    IMUL r32, r/m32               Doubleword register = doubleword register * r/m32.
	//    IMUL r64, r/m64               Quadword register = quadword register * r/m64.
	//    IMUL r16, r/m16, imm8         Word register = r/m16 * sign-extended immediate byte.
	//    IMUL r32, r/m32, imm8         Doubleword register = r/m32 * sign-extended immediate byte.
	//    IMUL r64, r/m64, imm8         Quadword register = r/m64 * sign-extended immediate byte.
	//    IMUL r16, r/m16, imm16        Word register = r/m16 ∗ immediate word.
	//    IMUL r32, r/m32, imm32        Doubleword register = r/m32 * immediate doubleword.
	//    IMUL r64, r/m64, imm32        Quadword register = r/m64 * immediate doubleword.
	//
	// Performs a signed multiplication of two operands.
	var x, y value.Value
	switch {
	case inst.Args[2] != nil:
		// Three-operand form.
		x, y = f.useArg(inst.Arg(1)), f.useArg(inst.Arg(2))
	case inst.Args[1] != nil:
		// Two-operand form.
		x, y = f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	case inst.Args[0] != nil:
		// One-operand form.
		y := f.useArg(inst.Arg(0))
		size := f.l.sizeOfType(y.Type())
		var dst value.Value
		switch size {
		case 1:
			x = f.useReg(x86.AL)
			dst = f.reg(x86asm.AX)
		case 2:
			x = f.useReg(x86.AX)
			dst = f.reg(x86.X86asm_DX_AX)
		case 3:
			x = f.useReg(x86.EAX)
			dst = f.reg(x86.X86asm_EDX_EAX)
		case 4:
			x = f.useReg(x86.RAX)
			dst = f.reg(x86.X86asm_RDX_RAX)
		default:
			panic(fmt.Errorf("support for operand type of byte size %d not yet implemented", size))
		}
		result := f.cur.NewMul(x, y)
		f.cur.NewStore(result, dst)
		return nil
	}
	result := f.cur.NewMul(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ IN ] ------------------------------------------------------------------

// liftInstIN lifts the given x86 IN instruction to LLVM IR, emitting code to f.
func (f *Func) liftInstIN(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstIN: not yet implemented")
}

// --- [ INC ] -----------------------------------------------------------------

// liftInstINC lifts the given x86 INC instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstINC(inst *x86.Inst) error {
	x := f.useArg(inst.Arg(0))
	one := constant.NewInt(1, types.I32)
	result := f.cur.NewAdd(x, one)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ INSB ] ----------------------------------------------------------------

// liftInstINSB lifts the given x86 INSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstINSB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINSB: not yet implemented")
}

// --- [ INSD ] ----------------------------------------------------------------

// liftInstINSD lifts the given x86 INSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstINSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINSD: not yet implemented")
}

// --- [ INSERTPS ] ------------------------------------------------------------

// liftInstINSERTPS lifts the given x86 INSERTPS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstINSERTPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINSERTPS: not yet implemented")
}

// --- [ INSW ] ----------------------------------------------------------------

// liftInstINSW lifts the given x86 INSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstINSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINSW: not yet implemented")
}

// --- [ INT ] -----------------------------------------------------------------

// liftInstINT lifts the given x86 INT instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstINT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINT: not yet implemented")
}

// --- [ INTO ] ----------------------------------------------------------------

// liftInstINTO lifts the given x86 INTO instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstINTO(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINTO: not yet implemented")
}

// --- [ INVD ] ----------------------------------------------------------------

// liftInstINVD lifts the given x86 INVD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstINVD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINVD: not yet implemented")
}

// --- [ INVLPG ] --------------------------------------------------------------

// liftInstINVLPG lifts the given x86 INVLPG instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstINVLPG(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINVLPG: not yet implemented")
}

// --- [ INVPCID ] -------------------------------------------------------------

// liftInstINVPCID lifts the given x86 INVPCID instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstINVPCID(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstINVPCID: not yet implemented")
}

// --- [ IRET ] ----------------------------------------------------------------

// liftInstIRET lifts the given x86 IRET instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstIRET(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstIRET: not yet implemented")
}

// --- [ IRETD ] ---------------------------------------------------------------

// liftInstIRETD lifts the given x86 IRETD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstIRETD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstIRETD: not yet implemented")
}

// --- [ IRETQ ] ---------------------------------------------------------------

// liftInstIRETQ lifts the given x86 IRETQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstIRETQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstIRETQ: not yet implemented")
}

// --- [ LAHF ] ----------------------------------------------------------------

// liftInstLAHF lifts the given x86 LAHF instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLAHF(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLAHF: not yet implemented")
}

// --- [ LAR ] -----------------------------------------------------------------

// liftInstLAR lifts the given x86 LAR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstLAR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLAR: not yet implemented")
}

// --- [ LCALL ] ---------------------------------------------------------------

// liftInstLCALL lifts the given x86 LCALL instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLCALL(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLCALL: not yet implemented")
}

// --- [ LDDQU ] ---------------------------------------------------------------

// liftInstLDDQU lifts the given x86 LDDQU instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLDDQU(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLDDQU: not yet implemented")
}

// --- [ LDMXCSR ] -------------------------------------------------------------

// liftInstLDMXCSR lifts the given x86 LDMXCSR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstLDMXCSR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLDMXCSR: not yet implemented")
}

// --- [ LDS ] -----------------------------------------------------------------

// liftInstLDS lifts the given x86 LDS instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstLDS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLDS: not yet implemented")
}

// --- [ LEA ] -----------------------------------------------------------------

// liftInstLEA lifts the given x86 LEA instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstLEA(inst *x86.Inst) error {
	y := f.mem(inst.Mem(1))
	f.defArg(inst.Arg(0), y)
	return nil
}

// --- [ LEAVE ] ---------------------------------------------------------------

// liftInstLEAVE lifts the given x86 LEAVE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLEAVE(inst *x86.Inst) error {
	// Pseudo-instruction for:
	//
	//    mov esp, ebp
	//    pop ebp

	//    mov esp, ebp
	ebp := f.useReg(x86.EBP)
	f.defReg(x86.ESP, ebp)
	// TODO: Explicitly setting espDisp to -4 should not be needed once espDisp
	// is stored per basic block and its changes tracked through the CFG. Remove
	// when handling of espDisp has matured.
	f.espDisp = -4

	//    pop ebp
	ebp = f.pop()
	f.defReg(x86.EBP, ebp)

	return nil
}

// --- [ LES ] -----------------------------------------------------------------

// liftInstLES lifts the given x86 LES instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstLES(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLES: not yet implemented")
}

// --- [ LFENCE ] --------------------------------------------------------------

// liftInstLFENCE lifts the given x86 LFENCE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstLFENCE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLFENCE: not yet implemented")
}

// --- [ LFS ] -----------------------------------------------------------------

// liftInstLFS lifts the given x86 LFS instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstLFS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLFS: not yet implemented")
}

// --- [ LGDT ] ----------------------------------------------------------------

// liftInstLGDT lifts the given x86 LGDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLGDT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLGDT: not yet implemented")
}

// --- [ LGS ] -----------------------------------------------------------------

// liftInstLGS lifts the given x86 LGS instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstLGS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLGS: not yet implemented")
}

// --- [ LIDT ] ----------------------------------------------------------------

// liftInstLIDT lifts the given x86 LIDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLIDT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLIDT: not yet implemented")
}

// --- [ LJMP ] ----------------------------------------------------------------

// liftInstLJMP lifts the given x86 LJMP instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLJMP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLJMP: not yet implemented")
}

// --- [ LLDT ] ----------------------------------------------------------------

// liftInstLLDT lifts the given x86 LLDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLLDT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLLDT: not yet implemented")
}

// --- [ LMSW ] ----------------------------------------------------------------

// liftInstLMSW lifts the given x86 LMSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLMSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLMSW: not yet implemented")
}

// --- [ LODSB ] ---------------------------------------------------------------

// liftInstLODSB lifts the given x86 LODSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLODSB(inst *x86.Inst) error {
	src := f.useArg(inst.Arg(1))
	f.defArgElem(inst.Arg(0), src, types.I8)
	return nil
}

// --- [ LODSD ] ---------------------------------------------------------------

// liftInstLODSD lifts the given x86 LODSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLODSD(inst *x86.Inst) error {
	src := f.useArg(inst.Arg(1))
	f.defArgElem(inst.Arg(0), src, types.I32)
	return nil
}

// --- [ LODSQ ] ---------------------------------------------------------------

// liftInstLODSQ lifts the given x86 LODSQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLODSQ(inst *x86.Inst) error {
	src := f.useArg(inst.Arg(1))
	f.defArgElem(inst.Arg(0), src, types.I64)
	return nil
}

// --- [ LODSW ] ---------------------------------------------------------------

// liftInstLODSW lifts the given x86 LODSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLODSW(inst *x86.Inst) error {
	src := f.useArg(inst.Arg(1))
	f.defArgElem(inst.Arg(0), src, types.I16)
	return nil
}

// --- [ LRET ] ----------------------------------------------------------------

// liftInstLRET lifts the given x86 LRET instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLRET(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLRET: not yet implemented")
}

// --- [ LSL ] -----------------------------------------------------------------

// liftInstLSL lifts the given x86 LSL instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstLSL(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLSL: not yet implemented")
}

// --- [ LSS ] -----------------------------------------------------------------

// liftInstLSS lifts the given x86 LSS instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstLSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLSS: not yet implemented")
}

// --- [ LTR ] -----------------------------------------------------------------

// liftInstLTR lifts the given x86 LTR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstLTR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLTR: not yet implemented")
}

// --- [ LZCNT ] ---------------------------------------------------------------

// liftInstLZCNT lifts the given x86 LZCNT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstLZCNT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstLZCNT: not yet implemented")
}

// --- [ MASKMOVDQU ] ----------------------------------------------------------

// liftInstMASKMOVDQU lifts the given x86 MASKMOVDQU instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstMASKMOVDQU(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMASKMOVDQU: not yet implemented")
}

// --- [ MASKMOVQ ] ------------------------------------------------------------

// liftInstMASKMOVQ lifts the given x86 MASKMOVQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstMASKMOVQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMASKMOVQ: not yet implemented")
}

// --- [ MAXPD ] ---------------------------------------------------------------

// liftInstMAXPD lifts the given x86 MAXPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMAXPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMAXPD: not yet implemented")
}

// --- [ MAXPS ] ---------------------------------------------------------------

// liftInstMAXPS lifts the given x86 MAXPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMAXPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMAXPS: not yet implemented")
}

// --- [ MAXSD ] ---------------------------------------------------------------

// liftInstMAXSD lifts the given x86 MAXSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMAXSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMAXSD: not yet implemented")
}

// --- [ MAXSS ] ---------------------------------------------------------------

// liftInstMAXSS lifts the given x86 MAXSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMAXSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMAXSS: not yet implemented")
}

// --- [ MFENCE ] --------------------------------------------------------------

// liftInstMFENCE lifts the given x86 MFENCE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMFENCE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMFENCE: not yet implemented")
}

// --- [ MINPD ] ---------------------------------------------------------------

// liftInstMINPD lifts the given x86 MINPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMINPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMINPD: not yet implemented")
}

// --- [ MINPS ] ---------------------------------------------------------------

// liftInstMINPS lifts the given x86 MINPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMINPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMINPS: not yet implemented")
}

// --- [ MINSD ] ---------------------------------------------------------------

// liftInstMINSD lifts the given x86 MINSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMINSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMINSD: not yet implemented")
}

// --- [ MINSS ] ---------------------------------------------------------------

// liftInstMINSS lifts the given x86 MINSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMINSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMINSS: not yet implemented")
}

// --- [ MONITOR ] -------------------------------------------------------------

// liftInstMONITOR lifts the given x86 MONITOR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMONITOR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMONITOR: not yet implemented")
}

// --- [ MOV ] -----------------------------------------------------------------

// liftInstMOV lifts the given x86 MOV instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstMOV(inst *x86.Inst) error {
	src := f.useArg(inst.Arg(1))
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MOVAPD ] --------------------------------------------------------------

// liftInstMOVAPD lifts the given x86 MOVAPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVAPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVAPD: not yet implemented")
}

// --- [ MOVAPS ] --------------------------------------------------------------

// liftInstMOVAPS lifts the given x86 MOVAPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVAPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVAPS: not yet implemented")
}

// --- [ MOVBE ] ---------------------------------------------------------------

// liftInstMOVBE lifts the given x86 MOVBE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMOVBE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVBE: not yet implemented")
}

// --- [ MOVD ] ----------------------------------------------------------------

// liftInstMOVD lifts the given x86 MOVD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMOVD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVD: not yet implemented")
}

// --- [ MOVDDUP ] -------------------------------------------------------------

// liftInstMOVDDUP lifts the given x86 MOVDDUP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVDDUP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVDDUP: not yet implemented")
}

// --- [ MOVDQ2Q ] -------------------------------------------------------------

// liftInstMOVDQ2Q lifts the given x86 MOVDQ2Q instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVDQ2Q(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVDQ2Q: not yet implemented")
}

// --- [ MOVDQA ] --------------------------------------------------------------

// liftInstMOVDQA lifts the given x86 MOVDQA instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVDQA(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVDQA: not yet implemented")
}

// --- [ MOVDQU ] --------------------------------------------------------------

// liftInstMOVDQU lifts the given x86 MOVDQU instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVDQU(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVDQU: not yet implemented")
}

// --- [ MOVHLPS ] -------------------------------------------------------------

// liftInstMOVHLPS lifts the given x86 MOVHLPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVHLPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVHLPS: not yet implemented")
}

// --- [ MOVHPD ] --------------------------------------------------------------

// liftInstMOVHPD lifts the given x86 MOVHPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVHPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVHPD: not yet implemented")
}

// --- [ MOVHPS ] --------------------------------------------------------------

// liftInstMOVHPS lifts the given x86 MOVHPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVHPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVHPS: not yet implemented")
}

// --- [ MOVLHPS ] -------------------------------------------------------------

// liftInstMOVLHPS lifts the given x86 MOVLHPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVLHPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVLHPS: not yet implemented")
}

// --- [ MOVLPD ] --------------------------------------------------------------

// liftInstMOVLPD lifts the given x86 MOVLPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVLPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVLPD: not yet implemented")
}

// --- [ MOVLPS ] --------------------------------------------------------------

// liftInstMOVLPS lifts the given x86 MOVLPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVLPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVLPS: not yet implemented")
}

// --- [ MOVMSKPD ] ------------------------------------------------------------

// liftInstMOVMSKPD lifts the given x86 MOVMSKPD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstMOVMSKPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVMSKPD: not yet implemented")
}

// --- [ MOVMSKPS ] ------------------------------------------------------------

// liftInstMOVMSKPS lifts the given x86 MOVMSKPS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstMOVMSKPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVMSKPS: not yet implemented")
}

// --- [ MOVNTDQ ] -------------------------------------------------------------

// liftInstMOVNTDQ lifts the given x86 MOVNTDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVNTDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTDQ: not yet implemented")
}

// --- [ MOVNTDQA ] ------------------------------------------------------------

// liftInstMOVNTDQA lifts the given x86 MOVNTDQA instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstMOVNTDQA(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTDQA: not yet implemented")
}

// --- [ MOVNTI ] --------------------------------------------------------------

// liftInstMOVNTI lifts the given x86 MOVNTI instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVNTI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTI: not yet implemented")
}

// --- [ MOVNTPD ] -------------------------------------------------------------

// liftInstMOVNTPD lifts the given x86 MOVNTPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVNTPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTPD: not yet implemented")
}

// --- [ MOVNTPS ] -------------------------------------------------------------

// liftInstMOVNTPS lifts the given x86 MOVNTPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVNTPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTPS: not yet implemented")
}

// --- [ MOVNTQ ] --------------------------------------------------------------

// liftInstMOVNTQ lifts the given x86 MOVNTQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVNTQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTQ: not yet implemented")
}

// --- [ MOVNTSD ] -------------------------------------------------------------

// liftInstMOVNTSD lifts the given x86 MOVNTSD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVNTSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTSD: not yet implemented")
}

// --- [ MOVNTSS ] -------------------------------------------------------------

// liftInstMOVNTSS lifts the given x86 MOVNTSS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVNTSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVNTSS: not yet implemented")
}

// --- [ MOVQ ] ----------------------------------------------------------------

// liftInstMOVQ lifts the given x86 MOVQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMOVQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVQ: not yet implemented")
}

// --- [ MOVQ2DQ ] -------------------------------------------------------------

// liftInstMOVQ2DQ lifts the given x86 MOVQ2DQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVQ2DQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVQ2DQ: not yet implemented")
}

// --- [ MOVSB ] ---------------------------------------------------------------

// liftInstMOVSB lifts the given x86 MOVSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMOVSB(inst *x86.Inst) error {
	src := f.useArgElem(inst.Arg(1), types.I8)
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MOVSD ] ---------------------------------------------------------------

// liftInstMOVSD lifts the given x86 MOVSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMOVSD(inst *x86.Inst) error {
	src := f.useArgElem(inst.Arg(1), types.I32)
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MOVSD_XMM ] -----------------------------------------------------------

// liftInstMOVSD_XMM lifts the given x86 MOVSD_XMM instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstMOVSD_XMM(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSD_XMM: not yet implemented")
}

// --- [ MOVSHDUP ] ------------------------------------------------------------

// liftInstMOVSHDUP lifts the given x86 MOVSHDUP instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstMOVSHDUP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSHDUP: not yet implemented")
}

// --- [ MOVSLDUP ] ------------------------------------------------------------

// liftInstMOVSLDUP lifts the given x86 MOVSLDUP instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstMOVSLDUP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSLDUP: not yet implemented")
}

// --- [ MOVSQ ] ---------------------------------------------------------------

// liftInstMOVSQ lifts the given x86 MOVSQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMOVSQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSQ: not yet implemented")
}

// --- [ MOVSS ] ---------------------------------------------------------------

// liftInstMOVSS lifts the given x86 MOVSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMOVSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSS: not yet implemented")
}

// --- [ MOVSW ] ---------------------------------------------------------------

// liftInstMOVSW lifts the given x86 MOVSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMOVSW(inst *x86.Inst) error {
	src := f.useArgElem(inst.Arg(1), types.I16)
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MOVSX ] ---------------------------------------------------------------

// liftInstMOVSX lifts the given x86 MOVSX instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMOVSX(inst *x86.Inst) error {
	size := inst.MemBytes * 8
	elem := types.NewInt(size)
	src := f.useArgElem(inst.Arg(1), elem)
	// TODO: Handle dst type dynamically.
	src = f.cur.NewSExt(src, types.I32)
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MOVSXD ] --------------------------------------------------------------

// liftInstMOVSXD lifts the given x86 MOVSXD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVSXD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVSXD: not yet implemented")
}

// --- [ MOVUPD ] --------------------------------------------------------------

// liftInstMOVUPD lifts the given x86 MOVUPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVUPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVUPD: not yet implemented")
}

// --- [ MOVUPS ] --------------------------------------------------------------

// liftInstMOVUPS lifts the given x86 MOVUPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMOVUPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMOVUPS: not yet implemented")
}

// --- [ MOVZX ] ---------------------------------------------------------------

// liftInstMOVZX lifts the given x86 MOVZX instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMOVZX(inst *x86.Inst) error {
	size := inst.MemBytes * 8
	elem := types.NewInt(size)
	src := f.useArgElem(inst.Arg(1), elem)
	// TODO: Handle dst type dynamically.
	src = f.cur.NewZExt(src, types.I32)
	f.defArg(inst.Arg(0), src)
	return nil
}

// --- [ MPSADBW ] -------------------------------------------------------------

// liftInstMPSADBW lifts the given x86 MPSADBW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstMPSADBW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMPSADBW: not yet implemented")
}

// --- [ MUL ] -----------------------------------------------------------------

// liftInstMUL lifts the given x86 MUL instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstMUL(inst *x86.Inst) error {
	// MUL - Unsigned Multiply
	//
	//    MUL r/m8       Unsigned multiply (AX = AL ∗ r/m8).
	//    MUL r/m8       Unsigned multiply (AX = AL ∗ r/m8).
	//    MUL r/m16      Unsigned multiply (DX:AX = AX ∗ r/m16).
	//    MUL r/m32      Unsigned multiply (EDX:EAX = EAX ∗ r/m32).
	//    MUL r/m64      Unsigned multiply (RDX:RAX = RAX ∗ r/m64).
	//
	// Performs an unsigned multiplication of the first operand (destination
	// operand) and the second operand (source operand) and stores the result in
	// the destination operand.
	// One-operand form.
	y := f.useArg(inst.Arg(0))
	size := f.l.sizeOfType(y.Type())
	var x value.Value
	var dst value.Value
	switch size {
	case 1:
		x = f.useReg(x86.AL)
		dst = f.reg(x86asm.AX)
	case 2:
		x = f.useReg(x86.AX)
		dst = f.reg(x86.X86asm_DX_AX)
	case 3:
		x = f.useReg(x86.EAX)
		dst = f.reg(x86.X86asm_EDX_EAX)
	case 4:
		x = f.useReg(x86.RAX)
		dst = f.reg(x86.X86asm_RDX_RAX)
	default:
		panic(fmt.Errorf("support for operand type of byte size %d not yet implemented", size))
	}
	result := f.cur.NewMul(x, y)
	f.cur.NewStore(result, dst)
	return nil
}

// --- [ MULPD ] ---------------------------------------------------------------

// liftInstMULPD lifts the given x86 MULPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMULPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMULPD: not yet implemented")
}

// --- [ MULPS ] ---------------------------------------------------------------

// liftInstMULPS lifts the given x86 MULPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMULPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMULPS: not yet implemented")
}

// --- [ MULSD ] ---------------------------------------------------------------

// liftInstMULSD lifts the given x86 MULSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMULSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMULSD: not yet implemented")
}

// --- [ MULSS ] ---------------------------------------------------------------

// liftInstMULSS lifts the given x86 MULSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMULSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMULSS: not yet implemented")
}

// --- [ MWAIT ] ---------------------------------------------------------------

// liftInstMWAIT lifts the given x86 MWAIT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstMWAIT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstMWAIT: not yet implemented")
}

// --- [ NEG ] -----------------------------------------------------------------

// liftInstNEG lifts the given x86 NEG instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstNEG(inst *x86.Inst) error {
	x := f.useArg(inst.Arg(0))
	zero := constant.NewInt(0, x.Type())
	result := f.cur.NewSub(zero, x)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ NOP ] -----------------------------------------------------------------

// liftInstNOP lifts the given x86 NOP instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstNOP(inst *x86.Inst) error {
	return nil
}

// --- [ NOT ] -----------------------------------------------------------------

// liftInstNOT lifts the given x86 NOT instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstNOT(inst *x86.Inst) error {
	x := f.useArg(inst.Arg(0))
	var mask value.Value
	typ, ok := x.Type().(*types.IntType)
	if !ok {
		panic(fmt.Errorf("invalid NOT operand type; expected *types.IntType, got %T", x.Type()))
	}
	switch typ.Size {
	case 8:
		mask = constant.NewInt(0xFF, types.I8)
	case 16:
		mask = constant.NewInt(0xFFFF, types.I16)
	case 32:
		mask = constant.NewInt(0xFFFFFFFF, types.I32)
	case 64:
		mask = constant.NewIntFromString("0xFFFFFFFFFFFFFFFF", types.I64)
	default:
		panic(fmt.Errorf("support for operand bit size %d not yet implemented", typ.Size))
	}
	result := f.cur.NewXor(x, mask)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ OR ] ------------------------------------------------------------------

// liftInstOR lifts the given x86 OR instruction to LLVM IR, emitting code to f.
func (f *Func) liftInstOR(inst *x86.Inst) error {
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewOr(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ ORPD ] ----------------------------------------------------------------

// liftInstORPD lifts the given x86 ORPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstORPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstORPD: not yet implemented")
}

// --- [ ORPS ] ----------------------------------------------------------------

// liftInstORPS lifts the given x86 ORPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstORPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstORPS: not yet implemented")
}

// --- [ OUT ] -----------------------------------------------------------------

// liftInstOUT lifts the given x86 OUT instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstOUT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstOUT: not yet implemented")
}

// --- [ OUTSB ] ---------------------------------------------------------------

// liftInstOUTSB lifts the given x86 OUTSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstOUTSB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstOUTSB: not yet implemented")
}

// --- [ OUTSD ] ---------------------------------------------------------------

// liftInstOUTSD lifts the given x86 OUTSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstOUTSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstOUTSD: not yet implemented")
}

// --- [ OUTSW ] ---------------------------------------------------------------

// liftInstOUTSW lifts the given x86 OUTSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstOUTSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstOUTSW: not yet implemented")
}

// --- [ PABSB ] ---------------------------------------------------------------

// liftInstPABSB lifts the given x86 PABSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPABSB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPABSB: not yet implemented")
}

// --- [ PABSD ] ---------------------------------------------------------------

// liftInstPABSD lifts the given x86 PABSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPABSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPABSD: not yet implemented")
}

// --- [ PABSW ] ---------------------------------------------------------------

// liftInstPABSW lifts the given x86 PABSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPABSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPABSW: not yet implemented")
}

// --- [ PACKSSDW ] ------------------------------------------------------------

// liftInstPACKSSDW lifts the given x86 PACKSSDW instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPACKSSDW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPACKSSDW: not yet implemented")
}

// --- [ PACKSSWB ] ------------------------------------------------------------

// liftInstPACKSSWB lifts the given x86 PACKSSWB instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPACKSSWB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPACKSSWB: not yet implemented")
}

// --- [ PACKUSDW ] ------------------------------------------------------------

// liftInstPACKUSDW lifts the given x86 PACKUSDW instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPACKUSDW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPACKUSDW: not yet implemented")
}

// --- [ PACKUSWB ] ------------------------------------------------------------

// liftInstPACKUSWB lifts the given x86 PACKUSWB instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPACKUSWB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPACKUSWB: not yet implemented")
}

// --- [ PADDB ] ---------------------------------------------------------------

// liftInstPADDB lifts the given x86 PADDB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPADDB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDB: not yet implemented")
}

// --- [ PADDD ] ---------------------------------------------------------------

// liftInstPADDD lifts the given x86 PADDD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPADDD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDD: not yet implemented")
}

// --- [ PADDQ ] ---------------------------------------------------------------

// liftInstPADDQ lifts the given x86 PADDQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPADDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDQ: not yet implemented")
}

// --- [ PADDSB ] --------------------------------------------------------------

// liftInstPADDSB lifts the given x86 PADDSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPADDSB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDSB: not yet implemented")
}

// --- [ PADDSW ] --------------------------------------------------------------

// liftInstPADDSW lifts the given x86 PADDSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPADDSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDSW: not yet implemented")
}

// --- [ PADDUSB ] -------------------------------------------------------------

// liftInstPADDUSB lifts the given x86 PADDUSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPADDUSB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDUSB: not yet implemented")
}

// --- [ PADDUSW ] -------------------------------------------------------------

// liftInstPADDUSW lifts the given x86 PADDUSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPADDUSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDUSW: not yet implemented")
}

// --- [ PADDW ] ---------------------------------------------------------------

// liftInstPADDW lifts the given x86 PADDW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPADDW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPADDW: not yet implemented")
}

// --- [ PALIGNR ] -------------------------------------------------------------

// liftInstPALIGNR lifts the given x86 PALIGNR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPALIGNR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPALIGNR: not yet implemented")
}

// --- [ PAND ] ----------------------------------------------------------------

// liftInstPAND lifts the given x86 PAND instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPAND(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPAND: not yet implemented")
}

// --- [ PANDN ] ---------------------------------------------------------------

// liftInstPANDN lifts the given x86 PANDN instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPANDN(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPANDN: not yet implemented")
}

// --- [ PAUSE ] ---------------------------------------------------------------

// liftInstPAUSE lifts the given x86 PAUSE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPAUSE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPAUSE: not yet implemented")
}

// --- [ PAVGB ] ---------------------------------------------------------------

// liftInstPAVGB lifts the given x86 PAVGB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPAVGB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPAVGB: not yet implemented")
}

// --- [ PAVGW ] ---------------------------------------------------------------

// liftInstPAVGW lifts the given x86 PAVGW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPAVGW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPAVGW: not yet implemented")
}

// --- [ PBLENDVB ] ------------------------------------------------------------

// liftInstPBLENDVB lifts the given x86 PBLENDVB instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPBLENDVB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPBLENDVB: not yet implemented")
}

// --- [ PBLENDW ] -------------------------------------------------------------

// liftInstPBLENDW lifts the given x86 PBLENDW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPBLENDW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPBLENDW: not yet implemented")
}

// --- [ PCLMULQDQ ] -----------------------------------------------------------

// liftInstPCLMULQDQ lifts the given x86 PCLMULQDQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPCLMULQDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCLMULQDQ: not yet implemented")
}

// --- [ PCMPEQB ] -------------------------------------------------------------

// liftInstPCMPEQB lifts the given x86 PCMPEQB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPCMPEQB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPEQB: not yet implemented")
}

// --- [ PCMPEQD ] -------------------------------------------------------------

// liftInstPCMPEQD lifts the given x86 PCMPEQD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPCMPEQD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPEQD: not yet implemented")
}

// --- [ PCMPEQQ ] -------------------------------------------------------------

// liftInstPCMPEQQ lifts the given x86 PCMPEQQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPCMPEQQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPEQQ: not yet implemented")
}

// --- [ PCMPEQW ] -------------------------------------------------------------

// liftInstPCMPEQW lifts the given x86 PCMPEQW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPCMPEQW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPEQW: not yet implemented")
}

// --- [ PCMPESTRI ] -----------------------------------------------------------

// liftInstPCMPESTRI lifts the given x86 PCMPESTRI instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPCMPESTRI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPESTRI: not yet implemented")
}

// --- [ PCMPESTRM ] -----------------------------------------------------------

// liftInstPCMPESTRM lifts the given x86 PCMPESTRM instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPCMPESTRM(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPESTRM: not yet implemented")
}

// --- [ PCMPGTB ] -------------------------------------------------------------

// liftInstPCMPGTB lifts the given x86 PCMPGTB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPCMPGTB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPGTB: not yet implemented")
}

// --- [ PCMPGTD ] -------------------------------------------------------------

// liftInstPCMPGTD lifts the given x86 PCMPGTD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPCMPGTD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPGTD: not yet implemented")
}

// --- [ PCMPGTQ ] -------------------------------------------------------------

// liftInstPCMPGTQ lifts the given x86 PCMPGTQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPCMPGTQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPGTQ: not yet implemented")
}

// --- [ PCMPGTW ] -------------------------------------------------------------

// liftInstPCMPGTW lifts the given x86 PCMPGTW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPCMPGTW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPGTW: not yet implemented")
}

// --- [ PCMPISTRI ] -----------------------------------------------------------

// liftInstPCMPISTRI lifts the given x86 PCMPISTRI instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPCMPISTRI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPISTRI: not yet implemented")
}

// --- [ PCMPISTRM ] -----------------------------------------------------------

// liftInstPCMPISTRM lifts the given x86 PCMPISTRM instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPCMPISTRM(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPCMPISTRM: not yet implemented")
}

// --- [ PEXTRB ] --------------------------------------------------------------

// liftInstPEXTRB lifts the given x86 PEXTRB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPEXTRB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPEXTRB: not yet implemented")
}

// --- [ PEXTRD ] --------------------------------------------------------------

// liftInstPEXTRD lifts the given x86 PEXTRD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPEXTRD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPEXTRD: not yet implemented")
}

// --- [ PEXTRQ ] --------------------------------------------------------------

// liftInstPEXTRQ lifts the given x86 PEXTRQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPEXTRQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPEXTRQ: not yet implemented")
}

// --- [ PEXTRW ] --------------------------------------------------------------

// liftInstPEXTRW lifts the given x86 PEXTRW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPEXTRW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPEXTRW: not yet implemented")
}

// --- [ PHADDD ] --------------------------------------------------------------

// liftInstPHADDD lifts the given x86 PHADDD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPHADDD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHADDD: not yet implemented")
}

// --- [ PHADDSW ] -------------------------------------------------------------

// liftInstPHADDSW lifts the given x86 PHADDSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPHADDSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHADDSW: not yet implemented")
}

// --- [ PHADDW ] --------------------------------------------------------------

// liftInstPHADDW lifts the given x86 PHADDW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPHADDW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHADDW: not yet implemented")
}

// --- [ PHMINPOSUW ] ----------------------------------------------------------

// liftInstPHMINPOSUW lifts the given x86 PHMINPOSUW instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPHMINPOSUW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHMINPOSUW: not yet implemented")
}

// --- [ PHSUBD ] --------------------------------------------------------------

// liftInstPHSUBD lifts the given x86 PHSUBD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPHSUBD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHSUBD: not yet implemented")
}

// --- [ PHSUBSW ] -------------------------------------------------------------

// liftInstPHSUBSW lifts the given x86 PHSUBSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPHSUBSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHSUBSW: not yet implemented")
}

// --- [ PHSUBW ] --------------------------------------------------------------

// liftInstPHSUBW lifts the given x86 PHSUBW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPHSUBW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPHSUBW: not yet implemented")
}

// --- [ PINSRB ] --------------------------------------------------------------

// liftInstPINSRB lifts the given x86 PINSRB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPINSRB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPINSRB: not yet implemented")
}

// --- [ PINSRD ] --------------------------------------------------------------

// liftInstPINSRD lifts the given x86 PINSRD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPINSRD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPINSRD: not yet implemented")
}

// --- [ PINSRQ ] --------------------------------------------------------------

// liftInstPINSRQ lifts the given x86 PINSRQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPINSRQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPINSRQ: not yet implemented")
}

// --- [ PINSRW ] --------------------------------------------------------------

// liftInstPINSRW lifts the given x86 PINSRW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPINSRW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPINSRW: not yet implemented")
}

// --- [ PMADDUBSW ] -----------------------------------------------------------

// liftInstPMADDUBSW lifts the given x86 PMADDUBSW instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMADDUBSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMADDUBSW: not yet implemented")
}

// --- [ PMADDWD ] -------------------------------------------------------------

// liftInstPMADDWD lifts the given x86 PMADDWD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMADDWD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMADDWD: not yet implemented")
}

// --- [ PMAXSB ] --------------------------------------------------------------

// liftInstPMAXSB lifts the given x86 PMAXSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMAXSB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXSB: not yet implemented")
}

// --- [ PMAXSD ] --------------------------------------------------------------

// liftInstPMAXSD lifts the given x86 PMAXSD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMAXSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXSD: not yet implemented")
}

// --- [ PMAXSW ] --------------------------------------------------------------

// liftInstPMAXSW lifts the given x86 PMAXSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMAXSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXSW: not yet implemented")
}

// --- [ PMAXUB ] --------------------------------------------------------------

// liftInstPMAXUB lifts the given x86 PMAXUB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMAXUB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXUB: not yet implemented")
}

// --- [ PMAXUD ] --------------------------------------------------------------

// liftInstPMAXUD lifts the given x86 PMAXUD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMAXUD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXUD: not yet implemented")
}

// --- [ PMAXUW ] --------------------------------------------------------------

// liftInstPMAXUW lifts the given x86 PMAXUW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMAXUW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMAXUW: not yet implemented")
}

// --- [ PMINSB ] --------------------------------------------------------------

// liftInstPMINSB lifts the given x86 PMINSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMINSB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINSB: not yet implemented")
}

// --- [ PMINSD ] --------------------------------------------------------------

// liftInstPMINSD lifts the given x86 PMINSD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMINSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINSD: not yet implemented")
}

// --- [ PMINSW ] --------------------------------------------------------------

// liftInstPMINSW lifts the given x86 PMINSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMINSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINSW: not yet implemented")
}

// --- [ PMINUB ] --------------------------------------------------------------

// liftInstPMINUB lifts the given x86 PMINUB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMINUB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINUB: not yet implemented")
}

// --- [ PMINUD ] --------------------------------------------------------------

// liftInstPMINUD lifts the given x86 PMINUD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMINUD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINUD: not yet implemented")
}

// --- [ PMINUW ] --------------------------------------------------------------

// liftInstPMINUW lifts the given x86 PMINUW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMINUW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMINUW: not yet implemented")
}

// --- [ PMOVMSKB ] ------------------------------------------------------------

// liftInstPMOVMSKB lifts the given x86 PMOVMSKB instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVMSKB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVMSKB: not yet implemented")
}

// --- [ PMOVSXBD ] ------------------------------------------------------------

// liftInstPMOVSXBD lifts the given x86 PMOVSXBD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVSXBD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXBD: not yet implemented")
}

// --- [ PMOVSXBQ ] ------------------------------------------------------------

// liftInstPMOVSXBQ lifts the given x86 PMOVSXBQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVSXBQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXBQ: not yet implemented")
}

// --- [ PMOVSXBW ] ------------------------------------------------------------

// liftInstPMOVSXBW lifts the given x86 PMOVSXBW instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVSXBW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXBW: not yet implemented")
}

// --- [ PMOVSXDQ ] ------------------------------------------------------------

// liftInstPMOVSXDQ lifts the given x86 PMOVSXDQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVSXDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXDQ: not yet implemented")
}

// --- [ PMOVSXWD ] ------------------------------------------------------------

// liftInstPMOVSXWD lifts the given x86 PMOVSXWD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVSXWD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXWD: not yet implemented")
}

// --- [ PMOVSXWQ ] ------------------------------------------------------------

// liftInstPMOVSXWQ lifts the given x86 PMOVSXWQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVSXWQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVSXWQ: not yet implemented")
}

// --- [ PMOVZXBD ] ------------------------------------------------------------

// liftInstPMOVZXBD lifts the given x86 PMOVZXBD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVZXBD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXBD: not yet implemented")
}

// --- [ PMOVZXBQ ] ------------------------------------------------------------

// liftInstPMOVZXBQ lifts the given x86 PMOVZXBQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVZXBQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXBQ: not yet implemented")
}

// --- [ PMOVZXBW ] ------------------------------------------------------------

// liftInstPMOVZXBW lifts the given x86 PMOVZXBW instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVZXBW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXBW: not yet implemented")
}

// --- [ PMOVZXDQ ] ------------------------------------------------------------

// liftInstPMOVZXDQ lifts the given x86 PMOVZXDQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVZXDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXDQ: not yet implemented")
}

// --- [ PMOVZXWD ] ------------------------------------------------------------

// liftInstPMOVZXWD lifts the given x86 PMOVZXWD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVZXWD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXWD: not yet implemented")
}

// --- [ PMOVZXWQ ] ------------------------------------------------------------

// liftInstPMOVZXWQ lifts the given x86 PMOVZXWQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMOVZXWQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMOVZXWQ: not yet implemented")
}

// --- [ PMULDQ ] --------------------------------------------------------------

// liftInstPMULDQ lifts the given x86 PMULDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMULDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULDQ: not yet implemented")
}

// --- [ PMULHRSW ] ------------------------------------------------------------

// liftInstPMULHRSW lifts the given x86 PMULHRSW instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPMULHRSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULHRSW: not yet implemented")
}

// --- [ PMULHUW ] -------------------------------------------------------------

// liftInstPMULHUW lifts the given x86 PMULHUW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMULHUW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULHUW: not yet implemented")
}

// --- [ PMULHW ] --------------------------------------------------------------

// liftInstPMULHW lifts the given x86 PMULHW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMULHW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULHW: not yet implemented")
}

// --- [ PMULLD ] --------------------------------------------------------------

// liftInstPMULLD lifts the given x86 PMULLD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMULLD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULLD: not yet implemented")
}

// --- [ PMULLW ] --------------------------------------------------------------

// liftInstPMULLW lifts the given x86 PMULLW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMULLW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULLW: not yet implemented")
}

// --- [ PMULUDQ ] -------------------------------------------------------------

// liftInstPMULUDQ lifts the given x86 PMULUDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPMULUDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPMULUDQ: not yet implemented")
}

// --- [ POP ] -----------------------------------------------------------------

// liftInstPOP lifts the given x86 POP instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstPOP(inst *x86.Inst) error {
	v := f.pop()
	f.defArg(inst.Arg(0), v)
	return nil
}

// pop pops a value from the top of the stack of the function, emitting code to
// f.
func (f *Func) pop() value.Named {
	m := x86asm.Mem{
		Base: x86asm.ESP,
	}
	mem := x86.NewMem(m, nil)
	v := f.useMem(mem)
	f.espDisp += 4
	return v
}

// --- [ POPA ] ----------------------------------------------------------------

// liftInstPOPA lifts the given x86 POPA instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPOPA(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPA: not yet implemented")
}

// --- [ POPAD ] ---------------------------------------------------------------

// liftInstPOPAD lifts the given x86 POPAD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPOPAD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPAD: not yet implemented")
}

// --- [ POPCNT ] --------------------------------------------------------------

// liftInstPOPCNT lifts the given x86 POPCNT instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPOPCNT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPCNT: not yet implemented")
}

// --- [ POPF ] ----------------------------------------------------------------

// liftInstPOPF lifts the given x86 POPF instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPOPF(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPF: not yet implemented")
}

// --- [ POPFD ] ---------------------------------------------------------------

// liftInstPOPFD lifts the given x86 POPFD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPOPFD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPFD: not yet implemented")
}

// --- [ POPFQ ] ---------------------------------------------------------------

// liftInstPOPFQ lifts the given x86 POPFQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPOPFQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOPFQ: not yet implemented")
}

// --- [ POR ] -----------------------------------------------------------------

// liftInstPOR lifts the given x86 POR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstPOR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPOR: not yet implemented")
}

// --- [ PREFETCHNTA ] ---------------------------------------------------------

// liftInstPREFETCHNTA lifts the given x86 PREFETCHNTA instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPREFETCHNTA(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPREFETCHNTA: not yet implemented")
}

// --- [ PREFETCHT0 ] ----------------------------------------------------------

// liftInstPREFETCHT0 lifts the given x86 PREFETCHT0 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPREFETCHT0(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPREFETCHT0: not yet implemented")
}

// --- [ PREFETCHT1 ] ----------------------------------------------------------

// liftInstPREFETCHT1 lifts the given x86 PREFETCHT1 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPREFETCHT1(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPREFETCHT1: not yet implemented")
}

// --- [ PREFETCHT2 ] ----------------------------------------------------------

// liftInstPREFETCHT2 lifts the given x86 PREFETCHT2 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPREFETCHT2(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPREFETCHT2: not yet implemented")
}

// --- [ PREFETCHW ] -----------------------------------------------------------

// liftInstPREFETCHW lifts the given x86 PREFETCHW instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPREFETCHW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPREFETCHW: not yet implemented")
}

// --- [ PSADBW ] --------------------------------------------------------------

// liftInstPSADBW lifts the given x86 PSADBW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSADBW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSADBW: not yet implemented")
}

// --- [ PSHUFB ] --------------------------------------------------------------

// liftInstPSHUFB lifts the given x86 PSHUFB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSHUFB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSHUFB: not yet implemented")
}

// --- [ PSHUFD ] --------------------------------------------------------------

// liftInstPSHUFD lifts the given x86 PSHUFD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSHUFD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSHUFD: not yet implemented")
}

// --- [ PSHUFHW ] -------------------------------------------------------------

// liftInstPSHUFHW lifts the given x86 PSHUFHW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSHUFHW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSHUFHW: not yet implemented")
}

// --- [ PSHUFLW ] -------------------------------------------------------------

// liftInstPSHUFLW lifts the given x86 PSHUFLW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSHUFLW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSHUFLW: not yet implemented")
}

// --- [ PSHUFW ] --------------------------------------------------------------

// liftInstPSHUFW lifts the given x86 PSHUFW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSHUFW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSHUFW: not yet implemented")
}

// --- [ PSIGNB ] --------------------------------------------------------------

// liftInstPSIGNB lifts the given x86 PSIGNB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSIGNB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSIGNB: not yet implemented")
}

// --- [ PSIGND ] --------------------------------------------------------------

// liftInstPSIGND lifts the given x86 PSIGND instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSIGND(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSIGND: not yet implemented")
}

// --- [ PSIGNW ] --------------------------------------------------------------

// liftInstPSIGNW lifts the given x86 PSIGNW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSIGNW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSIGNW: not yet implemented")
}

// --- [ PSLLD ] ---------------------------------------------------------------

// liftInstPSLLD lifts the given x86 PSLLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSLLD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSLLD: not yet implemented")
}

// --- [ PSLLDQ ] --------------------------------------------------------------

// liftInstPSLLDQ lifts the given x86 PSLLDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSLLDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSLLDQ: not yet implemented")
}

// --- [ PSLLQ ] ---------------------------------------------------------------

// liftInstPSLLQ lifts the given x86 PSLLQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSLLQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSLLQ: not yet implemented")
}

// --- [ PSLLW ] ---------------------------------------------------------------

// liftInstPSLLW lifts the given x86 PSLLW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSLLW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSLLW: not yet implemented")
}

// --- [ PSRAD ] ---------------------------------------------------------------

// liftInstPSRAD lifts the given x86 PSRAD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSRAD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRAD: not yet implemented")
}

// --- [ PSRAW ] ---------------------------------------------------------------

// liftInstPSRAW lifts the given x86 PSRAW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSRAW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRAW: not yet implemented")
}

// --- [ PSRLD ] ---------------------------------------------------------------

// liftInstPSRLD lifts the given x86 PSRLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSRLD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRLD: not yet implemented")
}

// --- [ PSRLDQ ] --------------------------------------------------------------

// liftInstPSRLDQ lifts the given x86 PSRLDQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSRLDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRLDQ: not yet implemented")
}

// --- [ PSRLQ ] ---------------------------------------------------------------

// liftInstPSRLQ lifts the given x86 PSRLQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSRLQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRLQ: not yet implemented")
}

// --- [ PSRLW ] ---------------------------------------------------------------

// liftInstPSRLW lifts the given x86 PSRLW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSRLW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSRLW: not yet implemented")
}

// --- [ PSUBB ] ---------------------------------------------------------------

// liftInstPSUBB lifts the given x86 PSUBB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSUBB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBB: not yet implemented")
}

// --- [ PSUBD ] ---------------------------------------------------------------

// liftInstPSUBD lifts the given x86 PSUBD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSUBD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBD: not yet implemented")
}

// --- [ PSUBQ ] ---------------------------------------------------------------

// liftInstPSUBQ lifts the given x86 PSUBQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSUBQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBQ: not yet implemented")
}

// --- [ PSUBSB ] --------------------------------------------------------------

// liftInstPSUBSB lifts the given x86 PSUBSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSUBSB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBSB: not yet implemented")
}

// --- [ PSUBSW ] --------------------------------------------------------------

// liftInstPSUBSW lifts the given x86 PSUBSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSUBSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBSW: not yet implemented")
}

// --- [ PSUBUSB ] -------------------------------------------------------------

// liftInstPSUBUSB lifts the given x86 PSUBUSB instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSUBUSB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBUSB: not yet implemented")
}

// --- [ PSUBUSW ] -------------------------------------------------------------

// liftInstPSUBUSW lifts the given x86 PSUBUSW instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPSUBUSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBUSW: not yet implemented")
}

// --- [ PSUBW ] ---------------------------------------------------------------

// liftInstPSUBW lifts the given x86 PSUBW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPSUBW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPSUBW: not yet implemented")
}

// --- [ PTEST ] ---------------------------------------------------------------

// liftInstPTEST lifts the given x86 PTEST instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPTEST(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPTEST: not yet implemented")
}

// --- [ PUNPCKHBW ] -----------------------------------------------------------

// liftInstPUNPCKHBW lifts the given x86 PUNPCKHBW instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPUNPCKHBW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKHBW: not yet implemented")
}

// --- [ PUNPCKHDQ ] -----------------------------------------------------------

// liftInstPUNPCKHDQ lifts the given x86 PUNPCKHDQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPUNPCKHDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKHDQ: not yet implemented")
}

// --- [ PUNPCKHQDQ ] ----------------------------------------------------------

// liftInstPUNPCKHQDQ lifts the given x86 PUNPCKHQDQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPUNPCKHQDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKHQDQ: not yet implemented")
}

// --- [ PUNPCKHWD ] -----------------------------------------------------------

// liftInstPUNPCKHWD lifts the given x86 PUNPCKHWD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPUNPCKHWD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKHWD: not yet implemented")
}

// --- [ PUNPCKLBW ] -----------------------------------------------------------

// liftInstPUNPCKLBW lifts the given x86 PUNPCKLBW instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPUNPCKLBW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKLBW: not yet implemented")
}

// --- [ PUNPCKLDQ ] -----------------------------------------------------------

// liftInstPUNPCKLDQ lifts the given x86 PUNPCKLDQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPUNPCKLDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKLDQ: not yet implemented")
}

// --- [ PUNPCKLQDQ ] ----------------------------------------------------------

// liftInstPUNPCKLQDQ lifts the given x86 PUNPCKLQDQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPUNPCKLQDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKLQDQ: not yet implemented")
}

// --- [ PUNPCKLWD ] -----------------------------------------------------------

// liftInstPUNPCKLWD lifts the given x86 PUNPCKLWD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstPUNPCKLWD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUNPCKLWD: not yet implemented")
}

// --- [ PUSH ] ----------------------------------------------------------------

// liftInstPUSH lifts the given x86 PUSH instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPUSH(inst *x86.Inst) error {
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
	mem := x86.NewMem(m, nil)
	f.defMem(mem, v)
	f.espDisp -= 4
}

// --- [ PUSHA ] ---------------------------------------------------------------

// liftInstPUSHA lifts the given x86 PUSHA instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPUSHA(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUSHA: not yet implemented")
}

// --- [ PUSHAD ] --------------------------------------------------------------

// liftInstPUSHAD lifts the given x86 PUSHAD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPUSHAD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUSHAD: not yet implemented")
}

// --- [ PUSHF ] ---------------------------------------------------------------

// liftInstPUSHF lifts the given x86 PUSHF instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPUSHF(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUSHF: not yet implemented")
}

// --- [ PUSHFD ] --------------------------------------------------------------

// liftInstPUSHFD lifts the given x86 PUSHFD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPUSHFD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUSHFD: not yet implemented")
}

// --- [ PUSHFQ ] --------------------------------------------------------------

// liftInstPUSHFQ lifts the given x86 PUSHFQ instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstPUSHFQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPUSHFQ: not yet implemented")
}

// --- [ PXOR ] ----------------------------------------------------------------

// liftInstPXOR lifts the given x86 PXOR instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstPXOR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstPXOR: not yet implemented")
}

// --- [ RCL ] -----------------------------------------------------------------

// liftInstRCL lifts the given x86 RCL instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstRCL(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRCL: not yet implemented")
}

// --- [ RCPPS ] ---------------------------------------------------------------

// liftInstRCPPS lifts the given x86 RCPPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstRCPPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRCPPS: not yet implemented")
}

// --- [ RCPSS ] ---------------------------------------------------------------

// liftInstRCPSS lifts the given x86 RCPSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstRCPSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRCPSS: not yet implemented")
}

// --- [ RCR ] -----------------------------------------------------------------

// liftInstRCR lifts the given x86 RCR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstRCR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRCR: not yet implemented")
}

// --- [ RDFSBASE ] ------------------------------------------------------------

// liftInstRDFSBASE lifts the given x86 RDFSBASE instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstRDFSBASE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDFSBASE: not yet implemented")
}

// --- [ RDGSBASE ] ------------------------------------------------------------

// liftInstRDGSBASE lifts the given x86 RDGSBASE instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstRDGSBASE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDGSBASE: not yet implemented")
}

// --- [ RDMSR ] ---------------------------------------------------------------

// liftInstRDMSR lifts the given x86 RDMSR instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstRDMSR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDMSR: not yet implemented")
}

// --- [ RDPMC ] ---------------------------------------------------------------

// liftInstRDPMC lifts the given x86 RDPMC instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstRDPMC(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDPMC: not yet implemented")
}

// --- [ RDRAND ] --------------------------------------------------------------

// liftInstRDRAND lifts the given x86 RDRAND instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstRDRAND(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDRAND: not yet implemented")
}

// --- [ RDTSC ] ---------------------------------------------------------------

// liftInstRDTSC lifts the given x86 RDTSC instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstRDTSC(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDTSC: not yet implemented")
}

// --- [ RDTSCP ] --------------------------------------------------------------

// liftInstRDTSCP lifts the given x86 RDTSCP instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstRDTSCP(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRDTSCP: not yet implemented")
}

// --- [ ROL ] -----------------------------------------------------------------

// liftInstROL lifts the given x86 ROL instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstROL(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROL: not yet implemented")
}

// --- [ ROR ] -----------------------------------------------------------------

// liftInstROR lifts the given x86 ROR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstROR(inst *x86.Inst) error {
	// rotate right (ROR)
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	low := f.cur.NewLShr(x, y)
	typ, ok := y.Type().(*types.IntType)
	if !ok {
		panic(fmt.Errorf("invalid count operand type; expected *types.IntType, got %T", y.Type()))
	}
	bits := constant.NewInt(int64(typ.Size), typ)
	shift := f.cur.NewSub(bits, y)
	high := f.cur.NewShl(x, shift)
	result := f.cur.NewOr(low, high)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ ROUNDPD ] -------------------------------------------------------------

// liftInstROUNDPD lifts the given x86 ROUNDPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstROUNDPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROUNDPD: not yet implemented")
}

// --- [ ROUNDPS ] -------------------------------------------------------------

// liftInstROUNDPS lifts the given x86 ROUNDPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstROUNDPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROUNDPS: not yet implemented")
}

// --- [ ROUNDSD ] -------------------------------------------------------------

// liftInstROUNDSD lifts the given x86 ROUNDSD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstROUNDSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROUNDSD: not yet implemented")
}

// --- [ ROUNDSS ] -------------------------------------------------------------

// liftInstROUNDSS lifts the given x86 ROUNDSS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstROUNDSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstROUNDSS: not yet implemented")
}

// --- [ RSM ] -----------------------------------------------------------------

// liftInstRSM lifts the given x86 RSM instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstRSM(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRSM: not yet implemented")
}

// --- [ RSQRTPS ] -------------------------------------------------------------

// liftInstRSQRTPS lifts the given x86 RSQRTPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstRSQRTPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRSQRTPS: not yet implemented")
}

// --- [ RSQRTSS ] -------------------------------------------------------------

// liftInstRSQRTSS lifts the given x86 RSQRTSS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstRSQRTSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstRSQRTSS: not yet implemented")
}

// --- [ SAHF ] ----------------------------------------------------------------

// liftInstSAHF lifts the given x86 SAHF instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSAHF(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSAHF: not yet implemented")
}

// --- [ SAR ] -----------------------------------------------------------------

// liftInstSAR lifts the given x86 SAR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstSAR(inst *x86.Inst) error {
	// shift arithmetic right (SAR)
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewAShr(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ SBB ] -----------------------------------------------------------------

// liftInstSBB lifts the given x86 SBB instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstSBB(inst *x86.Inst) error {
	// SBB - Integer Subtraction with Borrow.
	//
	//    SBB AL, imm8        Subtract with borrow imm8 from AL.
	//    SBB AX, imm16       Subtract with borrow imm16 from AX.
	//    SBB EAX, imm32      Subtract with borrow imm32 from EAX.
	//    SBB RAX, imm32      Subtract with borrow sign-extended imm.32 to 64-bits from RAX.
	//    SBB r/m8, imm8      Subtract with borrow imm8 from r/m8.
	//    SBB r/m8*, imm8     Subtract with borrow imm8 from r/m8.
	//    SBB r/m16, imm16    Subtract with borrow imm16 from r/m16.
	//    SBB r/m32, imm32    Subtract with borrow imm32 from r/m32.
	//    SBB r/m64, imm32    Subtract with borrow sign-extended imm32 to 64-bits from r/m64.
	//    SBB r/m16, imm8     Subtract with borrow sign-extended imm8 from r/m16.
	//    SBB r/m32, imm8     Subtract with borrow sign-extended imm8 from r/m32.
	//    SBB r/m64, imm8     Subtract with borrow sign-extended imm8 from r/m64.
	//    SBB r/m8, r8        Subtract with borrow r8 from r/m8.
	//    SBB r/m8*, r8       Subtract with borrow r8 from r/m8.
	//    SBB r/m16, r16      Subtract with borrow r16 from r/m16.
	//    SBB r/m32, r32      Subtract with borrow r32 from r/m32.
	//    SBB r/m64, r64      Subtract with borrow r64 from r/m64.
	//    SBB r8, r/m8        Subtract with borrow r/m8 from r8.
	//    SBB r8*, r/m8*      Subtract with borrow r/m8 from r8.
	//    SBB r16, r/m16      Subtract with borrow r/m16 from r16.
	//    SBB r32, r/m32      Subtract with borrow r/m32 from r32.
	//    SBB r64, r/m64      Subtract with borrow r/m64 from r64.
	//
	// Adds the source operand (second operand) and the carry (CF) flag, and
	// subtracts the result from the destination operand (first operand). The
	// result of the subtraction is stored in the destination operand.
	dst := f.useArg(inst.Arg(0))
	src := f.useArg(inst.Arg(1))
	cf := f.useStatus(CF)
	v := f.cur.NewAdd(src, cf)
	result := f.cur.NewSub(dst, v)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ SCASB ] ---------------------------------------------------------------

// liftInstSCASB lifts the given x86 SCASB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSCASB(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSCASB: not yet implemented")
}

// --- [ SCASD ] ---------------------------------------------------------------

// liftInstSCASD lifts the given x86 SCASD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSCASD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSCASD: not yet implemented")
}

// --- [ SCASQ ] ---------------------------------------------------------------

// liftInstSCASQ lifts the given x86 SCASQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSCASQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSCASQ: not yet implemented")
}

// --- [ SCASW ] ---------------------------------------------------------------

// liftInstSCASW lifts the given x86 SCASW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSCASW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSCASW: not yet implemented")
}

// --- [ SFENCE ] --------------------------------------------------------------

// liftInstSFENCE lifts the given x86 SFENCE instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSFENCE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSFENCE: not yet implemented")
}

// --- [ SGDT ] ----------------------------------------------------------------

// liftInstSGDT lifts the given x86 SGDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSGDT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSGDT: not yet implemented")
}

// --- [ SHL ] -----------------------------------------------------------------

// liftInstSHL lifts the given x86 SHL instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstSHL(inst *x86.Inst) error {
	// shift logical left (SHL)
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewShl(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ SHLD ] ----------------------------------------------------------------

// liftInstSHLD lifts the given x86 SHLD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSHLD(inst *x86.Inst) error {
	// SHLD - Double Precision Shift Left
	//
	//    SHLD a1, a2, a3
	//
	// Shift a1 to left a3 places while shifting bits from a2 in from the right.
	a1, a2, a3 := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1)), f.useArg(inst.Arg(2))
	tmp1 := f.cur.NewZExt(a1, types.I64)
	n32 := constant.NewInt(32, types.I64)
	high := f.cur.NewShl(tmp1, n32)
	low := f.cur.NewZExt(a2, types.I64)
	tmp3 := f.cur.NewOr(high, low)
	tmp4 := f.cur.NewShl(tmp3, a3)
	tmp5 := f.cur.NewLShr(tmp4, n32)
	result := f.cur.NewTrunc(tmp5, types.I32)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ SHR ] -----------------------------------------------------------------

// liftInstSHR lifts the given x86 SHR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstSHR(inst *x86.Inst) error {
	// shift logical right (SHR)
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewLShr(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ SHRD ] ----------------------------------------------------------------

// liftInstSHRD lifts the given x86 SHRD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSHRD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSHRD: not yet implemented")
}

// --- [ SHUFPD ] --------------------------------------------------------------

// liftInstSHUFPD lifts the given x86 SHUFPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSHUFPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSHUFPD: not yet implemented")
}

// --- [ SHUFPS ] --------------------------------------------------------------

// liftInstSHUFPS lifts the given x86 SHUFPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSHUFPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSHUFPS: not yet implemented")
}

// --- [ SIDT ] ----------------------------------------------------------------

// liftInstSIDT lifts the given x86 SIDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSIDT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSIDT: not yet implemented")
}

// --- [ SLDT ] ----------------------------------------------------------------

// liftInstSLDT lifts the given x86 SLDT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSLDT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSLDT: not yet implemented")
}

// --- [ SMSW ] ----------------------------------------------------------------

// liftInstSMSW lifts the given x86 SMSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSMSW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSMSW: not yet implemented")
}

// --- [ SQRTPD ] --------------------------------------------------------------

// liftInstSQRTPD lifts the given x86 SQRTPD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSQRTPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSQRTPD: not yet implemented")
}

// --- [ SQRTPS ] --------------------------------------------------------------

// liftInstSQRTPS lifts the given x86 SQRTPS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSQRTPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSQRTPS: not yet implemented")
}

// --- [ SQRTSD ] --------------------------------------------------------------

// liftInstSQRTSD lifts the given x86 SQRTSD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSQRTSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSQRTSD: not yet implemented")
}

// --- [ SQRTSS ] --------------------------------------------------------------

// liftInstSQRTSS lifts the given x86 SQRTSS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSQRTSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSQRTSS: not yet implemented")
}

// --- [ STC ] -----------------------------------------------------------------

// liftInstSTC lifts the given x86 STC instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstSTC(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTC: not yet implemented")
}

// --- [ STD ] -----------------------------------------------------------------

// liftInstSTD lifts the given x86 STD instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstSTD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTD: not yet implemented")
}

// --- [ STI ] -----------------------------------------------------------------

// liftInstSTI lifts the given x86 STI instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstSTI(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTI: not yet implemented")
}

// --- [ STMXCSR ] -------------------------------------------------------------

// liftInstSTMXCSR lifts the given x86 STMXCSR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSTMXCSR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTMXCSR: not yet implemented")
}

// --- [ STOSB ] ---------------------------------------------------------------

// liftInstSTOSB lifts the given x86 STOSB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSTOSB(inst *x86.Inst) error {
	src := f.useArg(inst.Arg(1))
	f.defArgElem(inst.Arg(0), src, types.I8)
	return nil
}

// --- [ STOSD ] ---------------------------------------------------------------

// liftInstSTOSD lifts the given x86 STOSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSTOSD(inst *x86.Inst) error {
	src := f.useArg(inst.Arg(1))
	f.defArgElem(inst.Arg(0), src, types.I32)
	return nil
}

// --- [ STOSQ ] ---------------------------------------------------------------

// liftInstSTOSQ lifts the given x86 STOSQ instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSTOSQ(inst *x86.Inst) error {
	src := f.useArg(inst.Arg(1))
	f.defArgElem(inst.Arg(0), src, types.I64)
	return nil
}

// --- [ STOSW ] ---------------------------------------------------------------

// liftInstSTOSW lifts the given x86 STOSW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSTOSW(inst *x86.Inst) error {
	src := f.useArg(inst.Arg(1))
	f.defArgElem(inst.Arg(0), src, types.I16)
	return nil
}

// --- [ STR ] -----------------------------------------------------------------

// liftInstSTR lifts the given x86 STR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstSTR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSTR: not yet implemented")
}

// --- [ SUB ] -----------------------------------------------------------------

// liftInstSUB lifts the given x86 SUB instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstSUB(inst *x86.Inst) error {
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewSub(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ SUBPD ] ---------------------------------------------------------------

// liftInstSUBPD lifts the given x86 SUBPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSUBPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSUBPD: not yet implemented")
}

// --- [ SUBPS ] ---------------------------------------------------------------

// liftInstSUBPS lifts the given x86 SUBPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSUBPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSUBPS: not yet implemented")
}

// --- [ SUBSD ] ---------------------------------------------------------------

// liftInstSUBSD lifts the given x86 SUBSD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSUBSD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSUBSD: not yet implemented")
}

// --- [ SUBSS ] ---------------------------------------------------------------

// liftInstSUBSS lifts the given x86 SUBSS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstSUBSS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSUBSS: not yet implemented")
}

// --- [ SWAPGS ] --------------------------------------------------------------

// liftInstSWAPGS lifts the given x86 SWAPGS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSWAPGS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSWAPGS: not yet implemented")
}

// --- [ SYSCALL ] -------------------------------------------------------------

// liftInstSYSCALL lifts the given x86 SYSCALL instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSYSCALL(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSYSCALL: not yet implemented")
}

// --- [ SYSENTER ] ------------------------------------------------------------

// liftInstSYSENTER lifts the given x86 SYSENTER instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstSYSENTER(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSYSENTER: not yet implemented")
}

// --- [ SYSEXIT ] -------------------------------------------------------------

// liftInstSYSEXIT lifts the given x86 SYSEXIT instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSYSEXIT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSYSEXIT: not yet implemented")
}

// --- [ SYSRET ] --------------------------------------------------------------

// liftInstSYSRET lifts the given x86 SYSRET instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstSYSRET(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstSYSRET: not yet implemented")
}

// --- [ TEST ] ----------------------------------------------------------------

// liftInstTEST lifts the given x86 TEST instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstTEST(inst *x86.Inst) error {
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

// liftInstTZCNT lifts the given x86 TZCNT instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstTZCNT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstTZCNT: not yet implemented")
}

// --- [ UCOMISD ] -------------------------------------------------------------

// liftInstUCOMISD lifts the given x86 UCOMISD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstUCOMISD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUCOMISD: not yet implemented")
}

// --- [ UCOMISS ] -------------------------------------------------------------

// liftInstUCOMISS lifts the given x86 UCOMISS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstUCOMISS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUCOMISS: not yet implemented")
}

// --- [ UD1 ] -----------------------------------------------------------------

// liftInstUD1 lifts the given x86 UD1 instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstUD1(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUD1: not yet implemented")
}

// --- [ UD2 ] -----------------------------------------------------------------

// liftInstUD2 lifts the given x86 UD2 instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstUD2(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUD2: not yet implemented")
}

// --- [ UNPCKHPD ] ------------------------------------------------------------

// liftInstUNPCKHPD lifts the given x86 UNPCKHPD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstUNPCKHPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUNPCKHPD: not yet implemented")
}

// --- [ UNPCKHPS ] ------------------------------------------------------------

// liftInstUNPCKHPS lifts the given x86 UNPCKHPS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstUNPCKHPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUNPCKHPS: not yet implemented")
}

// --- [ UNPCKLPD ] ------------------------------------------------------------

// liftInstUNPCKLPD lifts the given x86 UNPCKLPD instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstUNPCKLPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUNPCKLPD: not yet implemented")
}

// --- [ UNPCKLPS ] ------------------------------------------------------------

// liftInstUNPCKLPS lifts the given x86 UNPCKLPS instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstUNPCKLPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstUNPCKLPS: not yet implemented")
}

// --- [ VERR ] ----------------------------------------------------------------

// liftInstVERR lifts the given x86 VERR instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstVERR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVERR: not yet implemented")
}

// --- [ VERW ] ----------------------------------------------------------------

// liftInstVERW lifts the given x86 VERW instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstVERW(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVERW: not yet implemented")
}

// --- [ VMOVDQA ] -------------------------------------------------------------

// liftInstVMOVDQA lifts the given x86 VMOVDQA instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstVMOVDQA(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVMOVDQA: not yet implemented")
}

// --- [ VMOVDQU ] -------------------------------------------------------------

// liftInstVMOVDQU lifts the given x86 VMOVDQU instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstVMOVDQU(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVMOVDQU: not yet implemented")
}

// --- [ VMOVNTDQ ] ------------------------------------------------------------

// liftInstVMOVNTDQ lifts the given x86 VMOVNTDQ instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstVMOVNTDQ(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVMOVNTDQ: not yet implemented")
}

// --- [ VMOVNTDQA ] -----------------------------------------------------------

// liftInstVMOVNTDQA lifts the given x86 VMOVNTDQA instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstVMOVNTDQA(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVMOVNTDQA: not yet implemented")
}

// --- [ VZEROUPPER ] ----------------------------------------------------------

// liftInstVZEROUPPER lifts the given x86 VZEROUPPER instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstVZEROUPPER(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstVZEROUPPER: not yet implemented")
}

// --- [ WBINVD ] --------------------------------------------------------------

// liftInstWBINVD lifts the given x86 WBINVD instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstWBINVD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstWBINVD: not yet implemented")
}

// --- [ WRFSBASE ] ------------------------------------------------------------

// liftInstWRFSBASE lifts the given x86 WRFSBASE instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstWRFSBASE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstWRFSBASE: not yet implemented")
}

// --- [ WRGSBASE ] ------------------------------------------------------------

// liftInstWRGSBASE lifts the given x86 WRGSBASE instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstWRGSBASE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstWRGSBASE: not yet implemented")
}

// --- [ WRMSR ] ---------------------------------------------------------------

// liftInstWRMSR lifts the given x86 WRMSR instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstWRMSR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstWRMSR: not yet implemented")
}

// --- [ XABORT ] --------------------------------------------------------------

// liftInstXABORT lifts the given x86 XABORT instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstXABORT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXABORT: not yet implemented")
}

// --- [ XADD ] ----------------------------------------------------------------

// liftInstXADD lifts the given x86 XADD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstXADD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXADD: not yet implemented")
}

// --- [ XBEGIN ] --------------------------------------------------------------

// liftInstXBEGIN lifts the given x86 XBEGIN instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstXBEGIN(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXBEGIN: not yet implemented")
}

// --- [ XCHG ] ----------------------------------------------------------------

// liftInstXCHG lifts the given x86 XCHG instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstXCHG(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXCHG: not yet implemented")
}

// --- [ XEND ] ----------------------------------------------------------------

// liftInstXEND lifts the given x86 XEND instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstXEND(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXEND: not yet implemented")
}

// --- [ XGETBV ] --------------------------------------------------------------

// liftInstXGETBV lifts the given x86 XGETBV instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstXGETBV(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXGETBV: not yet implemented")
}

// --- [ XLATB ] ---------------------------------------------------------------

// liftInstXLATB lifts the given x86 XLATB instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstXLATB(inst *x86.Inst) error {
	// Set AL to memory byte DS:[(E)BX + unsigned AL].
	mem := inst.Mem(0)
	if mem.Mem.Index != 0 {
		panic(fmt.Errorf("invalid index of XLAT memory reference; expected 0, got %v", mem.Mem.Index))
	}
	mem.Mem.Scale = 1
	mem.Mem.Index = x86asm.AL
	v := f.useMemElem(mem, types.I8)
	f.defReg(x86.AL, v)
	return nil
}

// --- [ XOR ] -----------------------------------------------------------------

// liftInstXOR lifts the given x86 XOR instruction to LLVM IR, emitting code to
// f.
func (f *Func) liftInstXOR(inst *x86.Inst) error {
	x, y := f.useArg(inst.Arg(0)), f.useArg(inst.Arg(1))
	result := f.cur.NewXor(x, y)
	f.defArg(inst.Arg(0), result)
	return nil
}

// --- [ XORPD ] ---------------------------------------------------------------

// liftInstXORPD lifts the given x86 XORPD instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstXORPD(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXORPD: not yet implemented")
}

// --- [ XORPS ] ---------------------------------------------------------------

// liftInstXORPS lifts the given x86 XORPS instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstXORPS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXORPS: not yet implemented")
}

// --- [ XRSTOR ] --------------------------------------------------------------

// liftInstXRSTOR lifts the given x86 XRSTOR instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstXRSTOR(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXRSTOR: not yet implemented")
}

// --- [ XRSTOR64 ] ------------------------------------------------------------

// liftInstXRSTOR64 lifts the given x86 XRSTOR64 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstXRSTOR64(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXRSTOR64: not yet implemented")
}

// --- [ XRSTORS ] -------------------------------------------------------------

// liftInstXRSTORS lifts the given x86 XRSTORS instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstXRSTORS(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXRSTORS: not yet implemented")
}

// --- [ XRSTORS64 ] -----------------------------------------------------------

// liftInstXRSTORS64 lifts the given x86 XRSTORS64 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstXRSTORS64(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXRSTORS64: not yet implemented")
}

// --- [ XSAVE ] ---------------------------------------------------------------

// liftInstXSAVE lifts the given x86 XSAVE instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstXSAVE(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVE: not yet implemented")
}

// --- [ XSAVE64 ] -------------------------------------------------------------

// liftInstXSAVE64 lifts the given x86 XSAVE64 instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstXSAVE64(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVE64: not yet implemented")
}

// --- [ XSAVEC ] --------------------------------------------------------------

// liftInstXSAVEC lifts the given x86 XSAVEC instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstXSAVEC(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVEC: not yet implemented")
}

// --- [ XSAVEC64 ] ------------------------------------------------------------

// liftInstXSAVEC64 lifts the given x86 XSAVEC64 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstXSAVEC64(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVEC64: not yet implemented")
}

// --- [ XSAVEOPT ] ------------------------------------------------------------

// liftInstXSAVEOPT lifts the given x86 XSAVEOPT instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstXSAVEOPT(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVEOPT: not yet implemented")
}

// --- [ XSAVEOPT64 ] ----------------------------------------------------------

// liftInstXSAVEOPT64 lifts the given x86 XSAVEOPT64 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstXSAVEOPT64(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVEOPT64: not yet implemented")
}

// --- [ XSAVES ] --------------------------------------------------------------

// liftInstXSAVES lifts the given x86 XSAVES instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstXSAVES(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVES: not yet implemented")
}

// --- [ XSAVES64 ] ------------------------------------------------------------

// liftInstXSAVES64 lifts the given x86 XSAVES64 instruction to LLVM IR,
// emitting code to f.
func (f *Func) liftInstXSAVES64(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSAVES64: not yet implemented")
}

// --- [ XSETBV ] --------------------------------------------------------------

// liftInstXSETBV lifts the given x86 XSETBV instruction to LLVM IR, emitting
// code to f.
func (f *Func) liftInstXSETBV(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXSETBV: not yet implemented")
}

// --- [ XTEST ] ---------------------------------------------------------------

// liftInstXTEST lifts the given x86 XTEST instruction to LLVM IR, emitting code
// to f.
func (f *Func) liftInstXTEST(inst *x86.Inst) error {
	pretty.Println("inst:", inst)
	panic("emitInstXTEST: not yet implemented")
}
