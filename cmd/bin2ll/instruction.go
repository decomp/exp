package main

import (
	"fmt"

	"golang.org/x/arch/x86/x86asm"
)

// emitInst translates the given x86 instruction to LLVM IR, emitting code to f.
func (d *disassembler) emitInst(f *function, inst *instruction) error {
	switch inst.Op {
	case x86asm.AAA:
		return d.emitInstAAA(f, inst)
	case x86asm.AAD:
		return d.emitInstAAD(f, inst)
	case x86asm.AAM:
		return d.emitInstAAM(f, inst)
	case x86asm.AAS:
		return d.emitInstAAS(f, inst)
	case x86asm.ADC:
		return d.emitInstADC(f, inst)
	case x86asm.ADD:
		return d.emitInstADD(f, inst)
	case x86asm.ADDPD:
		return d.emitInstADDPD(f, inst)
	case x86asm.ADDPS:
		return d.emitInstADDPS(f, inst)
	case x86asm.ADDSD:
		return d.emitInstADDSD(f, inst)
	case x86asm.ADDSS:
		return d.emitInstADDSS(f, inst)
	case x86asm.ADDSUBPD:
		return d.emitInstADDSUBPD(f, inst)
	case x86asm.ADDSUBPS:
		return d.emitInstADDSUBPS(f, inst)
	case x86asm.AESDEC:
		return d.emitInstAESDEC(f, inst)
	case x86asm.AESDECLAST:
		return d.emitInstAESDECLAST(f, inst)
	case x86asm.AESENC:
		return d.emitInstAESENC(f, inst)
	case x86asm.AESENCLAST:
		return d.emitInstAESENCLAST(f, inst)
	case x86asm.AESIMC:
		return d.emitInstAESIMC(f, inst)
	case x86asm.AESKEYGENASSIST:
		return d.emitInstAESKEYGENASSIST(f, inst)
	case x86asm.AND:
		return d.emitInstAND(f, inst)
	case x86asm.ANDNPD:
		return d.emitInstANDNPD(f, inst)
	case x86asm.ANDNPS:
		return d.emitInstANDNPS(f, inst)
	case x86asm.ANDPD:
		return d.emitInstANDPD(f, inst)
	case x86asm.ANDPS:
		return d.emitInstANDPS(f, inst)
	case x86asm.ARPL:
		return d.emitInstARPL(f, inst)
	case x86asm.BLENDPD:
		return d.emitInstBLENDPD(f, inst)
	case x86asm.BLENDPS:
		return d.emitInstBLENDPS(f, inst)
	case x86asm.BLENDVPD:
		return d.emitInstBLENDVPD(f, inst)
	case x86asm.BLENDVPS:
		return d.emitInstBLENDVPS(f, inst)
	case x86asm.BOUND:
		return d.emitInstBOUND(f, inst)
	case x86asm.BSF:
		return d.emitInstBSF(f, inst)
	case x86asm.BSR:
		return d.emitInstBSR(f, inst)
	case x86asm.BSWAP:
		return d.emitInstBSWAP(f, inst)
	case x86asm.BT:
		return d.emitInstBT(f, inst)
	case x86asm.BTC:
		return d.emitInstBTC(f, inst)
	case x86asm.BTR:
		return d.emitInstBTR(f, inst)
	case x86asm.BTS:
		return d.emitInstBTS(f, inst)
	case x86asm.CALL:
		return d.emitInstCALL(f, inst)
	case x86asm.CBW:
		return d.emitInstCBW(f, inst)
	case x86asm.CDQ:
		return d.emitInstCDQ(f, inst)
	case x86asm.CDQE:
		return d.emitInstCDQE(f, inst)
	case x86asm.CLC:
		return d.emitInstCLC(f, inst)
	case x86asm.CLD:
		return d.emitInstCLD(f, inst)
	case x86asm.CLFLUSH:
		return d.emitInstCLFLUSH(f, inst)
	case x86asm.CLI:
		return d.emitInstCLI(f, inst)
	case x86asm.CLTS:
		return d.emitInstCLTS(f, inst)
	case x86asm.CMC:
		return d.emitInstCMC(f, inst)
	case x86asm.CMOVA:
		return d.emitInstCMOVA(f, inst)
	case x86asm.CMOVAE:
		return d.emitInstCMOVAE(f, inst)
	case x86asm.CMOVB:
		return d.emitInstCMOVB(f, inst)
	case x86asm.CMOVBE:
		return d.emitInstCMOVBE(f, inst)
	case x86asm.CMOVE:
		return d.emitInstCMOVE(f, inst)
	case x86asm.CMOVG:
		return d.emitInstCMOVG(f, inst)
	case x86asm.CMOVGE:
		return d.emitInstCMOVGE(f, inst)
	case x86asm.CMOVL:
		return d.emitInstCMOVL(f, inst)
	case x86asm.CMOVLE:
		return d.emitInstCMOVLE(f, inst)
	case x86asm.CMOVNE:
		return d.emitInstCMOVNE(f, inst)
	case x86asm.CMOVNO:
		return d.emitInstCMOVNO(f, inst)
	case x86asm.CMOVNP:
		return d.emitInstCMOVNP(f, inst)
	case x86asm.CMOVNS:
		return d.emitInstCMOVNS(f, inst)
	case x86asm.CMOVO:
		return d.emitInstCMOVO(f, inst)
	case x86asm.CMOVP:
		return d.emitInstCMOVP(f, inst)
	case x86asm.CMOVS:
		return d.emitInstCMOVS(f, inst)
	case x86asm.CMP:
		return d.emitInstCMP(f, inst)
	case x86asm.CMPPD:
		return d.emitInstCMPPD(f, inst)
	case x86asm.CMPPS:
		return d.emitInstCMPPS(f, inst)
	case x86asm.CMPSB:
		return d.emitInstCMPSB(f, inst)
	case x86asm.CMPSD:
		return d.emitInstCMPSD(f, inst)
	case x86asm.CMPSD_XMM:
		return d.emitInstCMPSD_XMM(f, inst)
	case x86asm.CMPSQ:
		return d.emitInstCMPSQ(f, inst)
	case x86asm.CMPSS:
		return d.emitInstCMPSS(f, inst)
	case x86asm.CMPSW:
		return d.emitInstCMPSW(f, inst)
	case x86asm.CMPXCHG:
		return d.emitInstCMPXCHG(f, inst)
	case x86asm.CMPXCHG16B:
		return d.emitInstCMPXCHG16B(f, inst)
	case x86asm.CMPXCHG8B:
		return d.emitInstCMPXCHG8B(f, inst)
	case x86asm.COMISD:
		return d.emitInstCOMISD(f, inst)
	case x86asm.COMISS:
		return d.emitInstCOMISS(f, inst)
	case x86asm.CPUID:
		return d.emitInstCPUID(f, inst)
	case x86asm.CQO:
		return d.emitInstCQO(f, inst)
	case x86asm.CRC32:
		return d.emitInstCRC32(f, inst)
	case x86asm.CVTDQ2PD:
		return d.emitInstCVTDQ2PD(f, inst)
	case x86asm.CVTDQ2PS:
		return d.emitInstCVTDQ2PS(f, inst)
	case x86asm.CVTPD2DQ:
		return d.emitInstCVTPD2DQ(f, inst)
	case x86asm.CVTPD2PI:
		return d.emitInstCVTPD2PI(f, inst)
	case x86asm.CVTPD2PS:
		return d.emitInstCVTPD2PS(f, inst)
	case x86asm.CVTPI2PD:
		return d.emitInstCVTPI2PD(f, inst)
	case x86asm.CVTPI2PS:
		return d.emitInstCVTPI2PS(f, inst)
	case x86asm.CVTPS2DQ:
		return d.emitInstCVTPS2DQ(f, inst)
	case x86asm.CVTPS2PD:
		return d.emitInstCVTPS2PD(f, inst)
	case x86asm.CVTPS2PI:
		return d.emitInstCVTPS2PI(f, inst)
	case x86asm.CVTSD2SI:
		return d.emitInstCVTSD2SI(f, inst)
	case x86asm.CVTSD2SS:
		return d.emitInstCVTSD2SS(f, inst)
	case x86asm.CVTSI2SD:
		return d.emitInstCVTSI2SD(f, inst)
	case x86asm.CVTSI2SS:
		return d.emitInstCVTSI2SS(f, inst)
	case x86asm.CVTSS2SD:
		return d.emitInstCVTSS2SD(f, inst)
	case x86asm.CVTSS2SI:
		return d.emitInstCVTSS2SI(f, inst)
	case x86asm.CVTTPD2DQ:
		return d.emitInstCVTTPD2DQ(f, inst)
	case x86asm.CVTTPD2PI:
		return d.emitInstCVTTPD2PI(f, inst)
	case x86asm.CVTTPS2DQ:
		return d.emitInstCVTTPS2DQ(f, inst)
	case x86asm.CVTTPS2PI:
		return d.emitInstCVTTPS2PI(f, inst)
	case x86asm.CVTTSD2SI:
		return d.emitInstCVTTSD2SI(f, inst)
	case x86asm.CVTTSS2SI:
		return d.emitInstCVTTSS2SI(f, inst)
	case x86asm.CWD:
		return d.emitInstCWD(f, inst)
	case x86asm.CWDE:
		return d.emitInstCWDE(f, inst)
	case x86asm.DAA:
		return d.emitInstDAA(f, inst)
	case x86asm.DAS:
		return d.emitInstDAS(f, inst)
	case x86asm.DEC:
		return d.emitInstDEC(f, inst)
	case x86asm.DIV:
		return d.emitInstDIV(f, inst)
	case x86asm.DIVPD:
		return d.emitInstDIVPD(f, inst)
	case x86asm.DIVPS:
		return d.emitInstDIVPS(f, inst)
	case x86asm.DIVSD:
		return d.emitInstDIVSD(f, inst)
	case x86asm.DIVSS:
		return d.emitInstDIVSS(f, inst)
	case x86asm.DPPD:
		return d.emitInstDPPD(f, inst)
	case x86asm.DPPS:
		return d.emitInstDPPS(f, inst)
	case x86asm.EMMS:
		return d.emitInstEMMS(f, inst)
	case x86asm.ENTER:
		return d.emitInstENTER(f, inst)
	case x86asm.EXTRACTPS:
		return d.emitInstEXTRACTPS(f, inst)
	case x86asm.F2XM1:
		return d.emitInstF2XM1(f, inst)
	case x86asm.FABS:
		return d.emitInstFABS(f, inst)
	case x86asm.FADD:
		return d.emitInstFADD(f, inst)
	case x86asm.FADDP:
		return d.emitInstFADDP(f, inst)
	case x86asm.FBLD:
		return d.emitInstFBLD(f, inst)
	case x86asm.FBSTP:
		return d.emitInstFBSTP(f, inst)
	case x86asm.FCHS:
		return d.emitInstFCHS(f, inst)
	case x86asm.FCMOVB:
		return d.emitInstFCMOVB(f, inst)
	case x86asm.FCMOVBE:
		return d.emitInstFCMOVBE(f, inst)
	case x86asm.FCMOVE:
		return d.emitInstFCMOVE(f, inst)
	case x86asm.FCMOVNB:
		return d.emitInstFCMOVNB(f, inst)
	case x86asm.FCMOVNBE:
		return d.emitInstFCMOVNBE(f, inst)
	case x86asm.FCMOVNE:
		return d.emitInstFCMOVNE(f, inst)
	case x86asm.FCMOVNU:
		return d.emitInstFCMOVNU(f, inst)
	case x86asm.FCMOVU:
		return d.emitInstFCMOVU(f, inst)
	case x86asm.FCOM:
		return d.emitInstFCOM(f, inst)
	case x86asm.FCOMI:
		return d.emitInstFCOMI(f, inst)
	case x86asm.FCOMIP:
		return d.emitInstFCOMIP(f, inst)
	case x86asm.FCOMP:
		return d.emitInstFCOMP(f, inst)
	case x86asm.FCOMPP:
		return d.emitInstFCOMPP(f, inst)
	case x86asm.FCOS:
		return d.emitInstFCOS(f, inst)
	case x86asm.FDECSTP:
		return d.emitInstFDECSTP(f, inst)
	case x86asm.FDIV:
		return d.emitInstFDIV(f, inst)
	case x86asm.FDIVP:
		return d.emitInstFDIVP(f, inst)
	case x86asm.FDIVR:
		return d.emitInstFDIVR(f, inst)
	case x86asm.FDIVRP:
		return d.emitInstFDIVRP(f, inst)
	case x86asm.FFREE:
		return d.emitInstFFREE(f, inst)
	case x86asm.FFREEP:
		return d.emitInstFFREEP(f, inst)
	case x86asm.FIADD:
		return d.emitInstFIADD(f, inst)
	case x86asm.FICOM:
		return d.emitInstFICOM(f, inst)
	case x86asm.FICOMP:
		return d.emitInstFICOMP(f, inst)
	case x86asm.FIDIV:
		return d.emitInstFIDIV(f, inst)
	case x86asm.FIDIVR:
		return d.emitInstFIDIVR(f, inst)
	case x86asm.FILD:
		return d.emitInstFILD(f, inst)
	case x86asm.FIMUL:
		return d.emitInstFIMUL(f, inst)
	case x86asm.FINCSTP:
		return d.emitInstFINCSTP(f, inst)
	case x86asm.FIST:
		return d.emitInstFIST(f, inst)
	case x86asm.FISTP:
		return d.emitInstFISTP(f, inst)
	case x86asm.FISTTP:
		return d.emitInstFISTTP(f, inst)
	case x86asm.FISUB:
		return d.emitInstFISUB(f, inst)
	case x86asm.FISUBR:
		return d.emitInstFISUBR(f, inst)
	case x86asm.FLD:
		return d.emitInstFLD(f, inst)
	case x86asm.FLD1:
		return d.emitInstFLD1(f, inst)
	case x86asm.FLDCW:
		return d.emitInstFLDCW(f, inst)
	case x86asm.FLDENV:
		return d.emitInstFLDENV(f, inst)
	case x86asm.FLDL2E:
		return d.emitInstFLDL2E(f, inst)
	case x86asm.FLDL2T:
		return d.emitInstFLDL2T(f, inst)
	case x86asm.FLDLG2:
		return d.emitInstFLDLG2(f, inst)
	case x86asm.FLDLN2:
		return d.emitInstFLDLN2(f, inst)
	case x86asm.FLDPI:
		return d.emitInstFLDPI(f, inst)
	case x86asm.FLDZ:
		return d.emitInstFLDZ(f, inst)
	case x86asm.FMUL:
		return d.emitInstFMUL(f, inst)
	case x86asm.FMULP:
		return d.emitInstFMULP(f, inst)
	case x86asm.FNCLEX:
		return d.emitInstFNCLEX(f, inst)
	case x86asm.FNINIT:
		return d.emitInstFNINIT(f, inst)
	case x86asm.FNOP:
		return d.emitInstFNOP(f, inst)
	case x86asm.FNSAVE:
		return d.emitInstFNSAVE(f, inst)
	case x86asm.FNSTCW:
		return d.emitInstFNSTCW(f, inst)
	case x86asm.FNSTENV:
		return d.emitInstFNSTENV(f, inst)
	case x86asm.FNSTSW:
		return d.emitInstFNSTSW(f, inst)
	case x86asm.FPATAN:
		return d.emitInstFPATAN(f, inst)
	case x86asm.FPREM:
		return d.emitInstFPREM(f, inst)
	case x86asm.FPREM1:
		return d.emitInstFPREM1(f, inst)
	case x86asm.FPTAN:
		return d.emitInstFPTAN(f, inst)
	case x86asm.FRNDINT:
		return d.emitInstFRNDINT(f, inst)
	case x86asm.FRSTOR:
		return d.emitInstFRSTOR(f, inst)
	case x86asm.FSCALE:
		return d.emitInstFSCALE(f, inst)
	case x86asm.FSIN:
		return d.emitInstFSIN(f, inst)
	case x86asm.FSINCOS:
		return d.emitInstFSINCOS(f, inst)
	case x86asm.FSQRT:
		return d.emitInstFSQRT(f, inst)
	case x86asm.FST:
		return d.emitInstFST(f, inst)
	case x86asm.FSTP:
		return d.emitInstFSTP(f, inst)
	case x86asm.FSUB:
		return d.emitInstFSUB(f, inst)
	case x86asm.FSUBP:
		return d.emitInstFSUBP(f, inst)
	case x86asm.FSUBR:
		return d.emitInstFSUBR(f, inst)
	case x86asm.FSUBRP:
		return d.emitInstFSUBRP(f, inst)
	case x86asm.FTST:
		return d.emitInstFTST(f, inst)
	case x86asm.FUCOM:
		return d.emitInstFUCOM(f, inst)
	case x86asm.FUCOMI:
		return d.emitInstFUCOMI(f, inst)
	case x86asm.FUCOMIP:
		return d.emitInstFUCOMIP(f, inst)
	case x86asm.FUCOMP:
		return d.emitInstFUCOMP(f, inst)
	case x86asm.FUCOMPP:
		return d.emitInstFUCOMPP(f, inst)
	case x86asm.FWAIT:
		return d.emitInstFWAIT(f, inst)
	case x86asm.FXAM:
		return d.emitInstFXAM(f, inst)
	case x86asm.FXCH:
		return d.emitInstFXCH(f, inst)
	case x86asm.FXRSTOR:
		return d.emitInstFXRSTOR(f, inst)
	case x86asm.FXRSTOR64:
		return d.emitInstFXRSTOR64(f, inst)
	case x86asm.FXSAVE:
		return d.emitInstFXSAVE(f, inst)
	case x86asm.FXSAVE64:
		return d.emitInstFXSAVE64(f, inst)
	case x86asm.FXTRACT:
		return d.emitInstFXTRACT(f, inst)
	case x86asm.FYL2X:
		return d.emitInstFYL2X(f, inst)
	case x86asm.FYL2XP1:
		return d.emitInstFYL2XP1(f, inst)
	case x86asm.HADDPD:
		return d.emitInstHADDPD(f, inst)
	case x86asm.HADDPS:
		return d.emitInstHADDPS(f, inst)
	case x86asm.HLT:
		return d.emitInstHLT(f, inst)
	case x86asm.HSUBPD:
		return d.emitInstHSUBPD(f, inst)
	case x86asm.HSUBPS:
		return d.emitInstHSUBPS(f, inst)
	case x86asm.ICEBP:
		return d.emitInstICEBP(f, inst)
	case x86asm.IDIV:
		return d.emitInstIDIV(f, inst)
	case x86asm.IMUL:
		return d.emitInstIMUL(f, inst)
	case x86asm.IN:
		return d.emitInstIN(f, inst)
	case x86asm.INC:
		return d.emitInstINC(f, inst)
	case x86asm.INSB:
		return d.emitInstINSB(f, inst)
	case x86asm.INSD:
		return d.emitInstINSD(f, inst)
	case x86asm.INSERTPS:
		return d.emitInstINSERTPS(f, inst)
	case x86asm.INSW:
		return d.emitInstINSW(f, inst)
	case x86asm.INT:
		return d.emitInstINT(f, inst)
	case x86asm.INTO:
		return d.emitInstINTO(f, inst)
	case x86asm.INVD:
		return d.emitInstINVD(f, inst)
	case x86asm.INVLPG:
		return d.emitInstINVLPG(f, inst)
	case x86asm.INVPCID:
		return d.emitInstINVPCID(f, inst)
	case x86asm.IRET:
		return d.emitInstIRET(f, inst)
	case x86asm.IRETD:
		return d.emitInstIRETD(f, inst)
	case x86asm.IRETQ:
		return d.emitInstIRETQ(f, inst)
	case x86asm.JA:
		return d.emitInstJA(f, inst)
	case x86asm.JAE:
		return d.emitInstJAE(f, inst)
	case x86asm.JB:
		return d.emitInstJB(f, inst)
	case x86asm.JBE:
		return d.emitInstJBE(f, inst)
	case x86asm.JCXZ:
		return d.emitInstJCXZ(f, inst)
	case x86asm.JE:
		return d.emitInstJE(f, inst)
	case x86asm.JECXZ:
		return d.emitInstJECXZ(f, inst)
	case x86asm.JG:
		return d.emitInstJG(f, inst)
	case x86asm.JGE:
		return d.emitInstJGE(f, inst)
	case x86asm.JL:
		return d.emitInstJL(f, inst)
	case x86asm.JLE:
		return d.emitInstJLE(f, inst)
	case x86asm.JMP:
		return d.emitInstJMP(f, inst)
	case x86asm.JNE:
		return d.emitInstJNE(f, inst)
	case x86asm.JNO:
		return d.emitInstJNO(f, inst)
	case x86asm.JNP:
		return d.emitInstJNP(f, inst)
	case x86asm.JNS:
		return d.emitInstJNS(f, inst)
	case x86asm.JO:
		return d.emitInstJO(f, inst)
	case x86asm.JP:
		return d.emitInstJP(f, inst)
	case x86asm.JRCXZ:
		return d.emitInstJRCXZ(f, inst)
	case x86asm.JS:
		return d.emitInstJS(f, inst)
	case x86asm.LAHF:
		return d.emitInstLAHF(f, inst)
	case x86asm.LAR:
		return d.emitInstLAR(f, inst)
	case x86asm.LCALL:
		return d.emitInstLCALL(f, inst)
	case x86asm.LDDQU:
		return d.emitInstLDDQU(f, inst)
	case x86asm.LDMXCSR:
		return d.emitInstLDMXCSR(f, inst)
	case x86asm.LDS:
		return d.emitInstLDS(f, inst)
	case x86asm.LEA:
		return d.emitInstLEA(f, inst)
	case x86asm.LEAVE:
		return d.emitInstLEAVE(f, inst)
	case x86asm.LES:
		return d.emitInstLES(f, inst)
	case x86asm.LFENCE:
		return d.emitInstLFENCE(f, inst)
	case x86asm.LFS:
		return d.emitInstLFS(f, inst)
	case x86asm.LGDT:
		return d.emitInstLGDT(f, inst)
	case x86asm.LGS:
		return d.emitInstLGS(f, inst)
	case x86asm.LIDT:
		return d.emitInstLIDT(f, inst)
	case x86asm.LJMP:
		return d.emitInstLJMP(f, inst)
	case x86asm.LLDT:
		return d.emitInstLLDT(f, inst)
	case x86asm.LMSW:
		return d.emitInstLMSW(f, inst)
	case x86asm.LODSB:
		return d.emitInstLODSB(f, inst)
	case x86asm.LODSD:
		return d.emitInstLODSD(f, inst)
	case x86asm.LODSQ:
		return d.emitInstLODSQ(f, inst)
	case x86asm.LODSW:
		return d.emitInstLODSW(f, inst)
	case x86asm.LOOP:
		return d.emitInstLOOP(f, inst)
	case x86asm.LOOPE:
		return d.emitInstLOOPE(f, inst)
	case x86asm.LOOPNE:
		return d.emitInstLOOPNE(f, inst)
	case x86asm.LRET:
		return d.emitInstLRET(f, inst)
	case x86asm.LSL:
		return d.emitInstLSL(f, inst)
	case x86asm.LSS:
		return d.emitInstLSS(f, inst)
	case x86asm.LTR:
		return d.emitInstLTR(f, inst)
	case x86asm.LZCNT:
		return d.emitInstLZCNT(f, inst)
	case x86asm.MASKMOVDQU:
		return d.emitInstMASKMOVDQU(f, inst)
	case x86asm.MASKMOVQ:
		return d.emitInstMASKMOVQ(f, inst)
	case x86asm.MAXPD:
		return d.emitInstMAXPD(f, inst)
	case x86asm.MAXPS:
		return d.emitInstMAXPS(f, inst)
	case x86asm.MAXSD:
		return d.emitInstMAXSD(f, inst)
	case x86asm.MAXSS:
		return d.emitInstMAXSS(f, inst)
	case x86asm.MFENCE:
		return d.emitInstMFENCE(f, inst)
	case x86asm.MINPD:
		return d.emitInstMINPD(f, inst)
	case x86asm.MINPS:
		return d.emitInstMINPS(f, inst)
	case x86asm.MINSD:
		return d.emitInstMINSD(f, inst)
	case x86asm.MINSS:
		return d.emitInstMINSS(f, inst)
	case x86asm.MONITOR:
		return d.emitInstMONITOR(f, inst)
	case x86asm.MOV:
		return d.emitInstMOV(f, inst)
	case x86asm.MOVAPD:
		return d.emitInstMOVAPD(f, inst)
	case x86asm.MOVAPS:
		return d.emitInstMOVAPS(f, inst)
	case x86asm.MOVBE:
		return d.emitInstMOVBE(f, inst)
	case x86asm.MOVD:
		return d.emitInstMOVD(f, inst)
	case x86asm.MOVDDUP:
		return d.emitInstMOVDDUP(f, inst)
	case x86asm.MOVDQ2Q:
		return d.emitInstMOVDQ2Q(f, inst)
	case x86asm.MOVDQA:
		return d.emitInstMOVDQA(f, inst)
	case x86asm.MOVDQU:
		return d.emitInstMOVDQU(f, inst)
	case x86asm.MOVHLPS:
		return d.emitInstMOVHLPS(f, inst)
	case x86asm.MOVHPD:
		return d.emitInstMOVHPD(f, inst)
	case x86asm.MOVHPS:
		return d.emitInstMOVHPS(f, inst)
	case x86asm.MOVLHPS:
		return d.emitInstMOVLHPS(f, inst)
	case x86asm.MOVLPD:
		return d.emitInstMOVLPD(f, inst)
	case x86asm.MOVLPS:
		return d.emitInstMOVLPS(f, inst)
	case x86asm.MOVMSKPD:
		return d.emitInstMOVMSKPD(f, inst)
	case x86asm.MOVMSKPS:
		return d.emitInstMOVMSKPS(f, inst)
	case x86asm.MOVNTDQ:
		return d.emitInstMOVNTDQ(f, inst)
	case x86asm.MOVNTDQA:
		return d.emitInstMOVNTDQA(f, inst)
	case x86asm.MOVNTI:
		return d.emitInstMOVNTI(f, inst)
	case x86asm.MOVNTPD:
		return d.emitInstMOVNTPD(f, inst)
	case x86asm.MOVNTPS:
		return d.emitInstMOVNTPS(f, inst)
	case x86asm.MOVNTQ:
		return d.emitInstMOVNTQ(f, inst)
	case x86asm.MOVNTSD:
		return d.emitInstMOVNTSD(f, inst)
	case x86asm.MOVNTSS:
		return d.emitInstMOVNTSS(f, inst)
	case x86asm.MOVQ:
		return d.emitInstMOVQ(f, inst)
	case x86asm.MOVQ2DQ:
		return d.emitInstMOVQ2DQ(f, inst)
	case x86asm.MOVSB:
		return d.emitInstMOVSB(f, inst)
	case x86asm.MOVSD:
		return d.emitInstMOVSD(f, inst)
	case x86asm.MOVSD_XMM:
		return d.emitInstMOVSD_XMM(f, inst)
	case x86asm.MOVSHDUP:
		return d.emitInstMOVSHDUP(f, inst)
	case x86asm.MOVSLDUP:
		return d.emitInstMOVSLDUP(f, inst)
	case x86asm.MOVSQ:
		return d.emitInstMOVSQ(f, inst)
	case x86asm.MOVSS:
		return d.emitInstMOVSS(f, inst)
	case x86asm.MOVSW:
		return d.emitInstMOVSW(f, inst)
	case x86asm.MOVSX:
		return d.emitInstMOVSX(f, inst)
	case x86asm.MOVSXD:
		return d.emitInstMOVSXD(f, inst)
	case x86asm.MOVUPD:
		return d.emitInstMOVUPD(f, inst)
	case x86asm.MOVUPS:
		return d.emitInstMOVUPS(f, inst)
	case x86asm.MOVZX:
		return d.emitInstMOVZX(f, inst)
	case x86asm.MPSADBW:
		return d.emitInstMPSADBW(f, inst)
	case x86asm.MUL:
		return d.emitInstMUL(f, inst)
	case x86asm.MULPD:
		return d.emitInstMULPD(f, inst)
	case x86asm.MULPS:
		return d.emitInstMULPS(f, inst)
	case x86asm.MULSD:
		return d.emitInstMULSD(f, inst)
	case x86asm.MULSS:
		return d.emitInstMULSS(f, inst)
	case x86asm.MWAIT:
		return d.emitInstMWAIT(f, inst)
	case x86asm.NEG:
		return d.emitInstNEG(f, inst)
	case x86asm.NOP:
		return d.emitInstNOP(f, inst)
	case x86asm.NOT:
		return d.emitInstNOT(f, inst)
	case x86asm.OR:
		return d.emitInstOR(f, inst)
	case x86asm.ORPD:
		return d.emitInstORPD(f, inst)
	case x86asm.ORPS:
		return d.emitInstORPS(f, inst)
	case x86asm.OUT:
		return d.emitInstOUT(f, inst)
	case x86asm.OUTSB:
		return d.emitInstOUTSB(f, inst)
	case x86asm.OUTSD:
		return d.emitInstOUTSD(f, inst)
	case x86asm.OUTSW:
		return d.emitInstOUTSW(f, inst)
	case x86asm.PABSB:
		return d.emitInstPABSB(f, inst)
	case x86asm.PABSD:
		return d.emitInstPABSD(f, inst)
	case x86asm.PABSW:
		return d.emitInstPABSW(f, inst)
	case x86asm.PACKSSDW:
		return d.emitInstPACKSSDW(f, inst)
	case x86asm.PACKSSWB:
		return d.emitInstPACKSSWB(f, inst)
	case x86asm.PACKUSDW:
		return d.emitInstPACKUSDW(f, inst)
	case x86asm.PACKUSWB:
		return d.emitInstPACKUSWB(f, inst)
	case x86asm.PADDB:
		return d.emitInstPADDB(f, inst)
	case x86asm.PADDD:
		return d.emitInstPADDD(f, inst)
	case x86asm.PADDQ:
		return d.emitInstPADDQ(f, inst)
	case x86asm.PADDSB:
		return d.emitInstPADDSB(f, inst)
	case x86asm.PADDSW:
		return d.emitInstPADDSW(f, inst)
	case x86asm.PADDUSB:
		return d.emitInstPADDUSB(f, inst)
	case x86asm.PADDUSW:
		return d.emitInstPADDUSW(f, inst)
	case x86asm.PADDW:
		return d.emitInstPADDW(f, inst)
	case x86asm.PALIGNR:
		return d.emitInstPALIGNR(f, inst)
	case x86asm.PAND:
		return d.emitInstPAND(f, inst)
	case x86asm.PANDN:
		return d.emitInstPANDN(f, inst)
	case x86asm.PAUSE:
		return d.emitInstPAUSE(f, inst)
	case x86asm.PAVGB:
		return d.emitInstPAVGB(f, inst)
	case x86asm.PAVGW:
		return d.emitInstPAVGW(f, inst)
	case x86asm.PBLENDVB:
		return d.emitInstPBLENDVB(f, inst)
	case x86asm.PBLENDW:
		return d.emitInstPBLENDW(f, inst)
	case x86asm.PCLMULQDQ:
		return d.emitInstPCLMULQDQ(f, inst)
	case x86asm.PCMPEQB:
		return d.emitInstPCMPEQB(f, inst)
	case x86asm.PCMPEQD:
		return d.emitInstPCMPEQD(f, inst)
	case x86asm.PCMPEQQ:
		return d.emitInstPCMPEQQ(f, inst)
	case x86asm.PCMPEQW:
		return d.emitInstPCMPEQW(f, inst)
	case x86asm.PCMPESTRI:
		return d.emitInstPCMPESTRI(f, inst)
	case x86asm.PCMPESTRM:
		return d.emitInstPCMPESTRM(f, inst)
	case x86asm.PCMPGTB:
		return d.emitInstPCMPGTB(f, inst)
	case x86asm.PCMPGTD:
		return d.emitInstPCMPGTD(f, inst)
	case x86asm.PCMPGTQ:
		return d.emitInstPCMPGTQ(f, inst)
	case x86asm.PCMPGTW:
		return d.emitInstPCMPGTW(f, inst)
	case x86asm.PCMPISTRI:
		return d.emitInstPCMPISTRI(f, inst)
	case x86asm.PCMPISTRM:
		return d.emitInstPCMPISTRM(f, inst)
	case x86asm.PEXTRB:
		return d.emitInstPEXTRB(f, inst)
	case x86asm.PEXTRD:
		return d.emitInstPEXTRD(f, inst)
	case x86asm.PEXTRQ:
		return d.emitInstPEXTRQ(f, inst)
	case x86asm.PEXTRW:
		return d.emitInstPEXTRW(f, inst)
	case x86asm.PHADDD:
		return d.emitInstPHADDD(f, inst)
	case x86asm.PHADDSW:
		return d.emitInstPHADDSW(f, inst)
	case x86asm.PHADDW:
		return d.emitInstPHADDW(f, inst)
	case x86asm.PHMINPOSUW:
		return d.emitInstPHMINPOSUW(f, inst)
	case x86asm.PHSUBD:
		return d.emitInstPHSUBD(f, inst)
	case x86asm.PHSUBSW:
		return d.emitInstPHSUBSW(f, inst)
	case x86asm.PHSUBW:
		return d.emitInstPHSUBW(f, inst)
	case x86asm.PINSRB:
		return d.emitInstPINSRB(f, inst)
	case x86asm.PINSRD:
		return d.emitInstPINSRD(f, inst)
	case x86asm.PINSRQ:
		return d.emitInstPINSRQ(f, inst)
	case x86asm.PINSRW:
		return d.emitInstPINSRW(f, inst)
	case x86asm.PMADDUBSW:
		return d.emitInstPMADDUBSW(f, inst)
	case x86asm.PMADDWD:
		return d.emitInstPMADDWD(f, inst)
	case x86asm.PMAXSB:
		return d.emitInstPMAXSB(f, inst)
	case x86asm.PMAXSD:
		return d.emitInstPMAXSD(f, inst)
	case x86asm.PMAXSW:
		return d.emitInstPMAXSW(f, inst)
	case x86asm.PMAXUB:
		return d.emitInstPMAXUB(f, inst)
	case x86asm.PMAXUD:
		return d.emitInstPMAXUD(f, inst)
	case x86asm.PMAXUW:
		return d.emitInstPMAXUW(f, inst)
	case x86asm.PMINSB:
		return d.emitInstPMINSB(f, inst)
	case x86asm.PMINSD:
		return d.emitInstPMINSD(f, inst)
	case x86asm.PMINSW:
		return d.emitInstPMINSW(f, inst)
	case x86asm.PMINUB:
		return d.emitInstPMINUB(f, inst)
	case x86asm.PMINUD:
		return d.emitInstPMINUD(f, inst)
	case x86asm.PMINUW:
		return d.emitInstPMINUW(f, inst)
	case x86asm.PMOVMSKB:
		return d.emitInstPMOVMSKB(f, inst)
	case x86asm.PMOVSXBD:
		return d.emitInstPMOVSXBD(f, inst)
	case x86asm.PMOVSXBQ:
		return d.emitInstPMOVSXBQ(f, inst)
	case x86asm.PMOVSXBW:
		return d.emitInstPMOVSXBW(f, inst)
	case x86asm.PMOVSXDQ:
		return d.emitInstPMOVSXDQ(f, inst)
	case x86asm.PMOVSXWD:
		return d.emitInstPMOVSXWD(f, inst)
	case x86asm.PMOVSXWQ:
		return d.emitInstPMOVSXWQ(f, inst)
	case x86asm.PMOVZXBD:
		return d.emitInstPMOVZXBD(f, inst)
	case x86asm.PMOVZXBQ:
		return d.emitInstPMOVZXBQ(f, inst)
	case x86asm.PMOVZXBW:
		return d.emitInstPMOVZXBW(f, inst)
	case x86asm.PMOVZXDQ:
		return d.emitInstPMOVZXDQ(f, inst)
	case x86asm.PMOVZXWD:
		return d.emitInstPMOVZXWD(f, inst)
	case x86asm.PMOVZXWQ:
		return d.emitInstPMOVZXWQ(f, inst)
	case x86asm.PMULDQ:
		return d.emitInstPMULDQ(f, inst)
	case x86asm.PMULHRSW:
		return d.emitInstPMULHRSW(f, inst)
	case x86asm.PMULHUW:
		return d.emitInstPMULHUW(f, inst)
	case x86asm.PMULHW:
		return d.emitInstPMULHW(f, inst)
	case x86asm.PMULLD:
		return d.emitInstPMULLD(f, inst)
	case x86asm.PMULLW:
		return d.emitInstPMULLW(f, inst)
	case x86asm.PMULUDQ:
		return d.emitInstPMULUDQ(f, inst)
	case x86asm.POP:
		return d.emitInstPOP(f, inst)
	case x86asm.POPA:
		return d.emitInstPOPA(f, inst)
	case x86asm.POPAD:
		return d.emitInstPOPAD(f, inst)
	case x86asm.POPCNT:
		return d.emitInstPOPCNT(f, inst)
	case x86asm.POPF:
		return d.emitInstPOPF(f, inst)
	case x86asm.POPFD:
		return d.emitInstPOPFD(f, inst)
	case x86asm.POPFQ:
		return d.emitInstPOPFQ(f, inst)
	case x86asm.POR:
		return d.emitInstPOR(f, inst)
	case x86asm.PREFETCHNTA:
		return d.emitInstPREFETCHNTA(f, inst)
	case x86asm.PREFETCHT0:
		return d.emitInstPREFETCHT0(f, inst)
	case x86asm.PREFETCHT1:
		return d.emitInstPREFETCHT1(f, inst)
	case x86asm.PREFETCHT2:
		return d.emitInstPREFETCHT2(f, inst)
	case x86asm.PREFETCHW:
		return d.emitInstPREFETCHW(f, inst)
	case x86asm.PSADBW:
		return d.emitInstPSADBW(f, inst)
	case x86asm.PSHUFB:
		return d.emitInstPSHUFB(f, inst)
	case x86asm.PSHUFD:
		return d.emitInstPSHUFD(f, inst)
	case x86asm.PSHUFHW:
		return d.emitInstPSHUFHW(f, inst)
	case x86asm.PSHUFLW:
		return d.emitInstPSHUFLW(f, inst)
	case x86asm.PSHUFW:
		return d.emitInstPSHUFW(f, inst)
	case x86asm.PSIGNB:
		return d.emitInstPSIGNB(f, inst)
	case x86asm.PSIGND:
		return d.emitInstPSIGND(f, inst)
	case x86asm.PSIGNW:
		return d.emitInstPSIGNW(f, inst)
	case x86asm.PSLLD:
		return d.emitInstPSLLD(f, inst)
	case x86asm.PSLLDQ:
		return d.emitInstPSLLDQ(f, inst)
	case x86asm.PSLLQ:
		return d.emitInstPSLLQ(f, inst)
	case x86asm.PSLLW:
		return d.emitInstPSLLW(f, inst)
	case x86asm.PSRAD:
		return d.emitInstPSRAD(f, inst)
	case x86asm.PSRAW:
		return d.emitInstPSRAW(f, inst)
	case x86asm.PSRLD:
		return d.emitInstPSRLD(f, inst)
	case x86asm.PSRLDQ:
		return d.emitInstPSRLDQ(f, inst)
	case x86asm.PSRLQ:
		return d.emitInstPSRLQ(f, inst)
	case x86asm.PSRLW:
		return d.emitInstPSRLW(f, inst)
	case x86asm.PSUBB:
		return d.emitInstPSUBB(f, inst)
	case x86asm.PSUBD:
		return d.emitInstPSUBD(f, inst)
	case x86asm.PSUBQ:
		return d.emitInstPSUBQ(f, inst)
	case x86asm.PSUBSB:
		return d.emitInstPSUBSB(f, inst)
	case x86asm.PSUBSW:
		return d.emitInstPSUBSW(f, inst)
	case x86asm.PSUBUSB:
		return d.emitInstPSUBUSB(f, inst)
	case x86asm.PSUBUSW:
		return d.emitInstPSUBUSW(f, inst)
	case x86asm.PSUBW:
		return d.emitInstPSUBW(f, inst)
	case x86asm.PTEST:
		return d.emitInstPTEST(f, inst)
	case x86asm.PUNPCKHBW:
		return d.emitInstPUNPCKHBW(f, inst)
	case x86asm.PUNPCKHDQ:
		return d.emitInstPUNPCKHDQ(f, inst)
	case x86asm.PUNPCKHQDQ:
		return d.emitInstPUNPCKHQDQ(f, inst)
	case x86asm.PUNPCKHWD:
		return d.emitInstPUNPCKHWD(f, inst)
	case x86asm.PUNPCKLBW:
		return d.emitInstPUNPCKLBW(f, inst)
	case x86asm.PUNPCKLDQ:
		return d.emitInstPUNPCKLDQ(f, inst)
	case x86asm.PUNPCKLQDQ:
		return d.emitInstPUNPCKLQDQ(f, inst)
	case x86asm.PUNPCKLWD:
		return d.emitInstPUNPCKLWD(f, inst)
	case x86asm.PUSH:
		return d.emitInstPUSH(f, inst)
	case x86asm.PUSHA:
		return d.emitInstPUSHA(f, inst)
	case x86asm.PUSHAD:
		return d.emitInstPUSHAD(f, inst)
	case x86asm.PUSHF:
		return d.emitInstPUSHF(f, inst)
	case x86asm.PUSHFD:
		return d.emitInstPUSHFD(f, inst)
	case x86asm.PUSHFQ:
		return d.emitInstPUSHFQ(f, inst)
	case x86asm.PXOR:
		return d.emitInstPXOR(f, inst)
	case x86asm.RCL:
		return d.emitInstRCL(f, inst)
	case x86asm.RCPPS:
		return d.emitInstRCPPS(f, inst)
	case x86asm.RCPSS:
		return d.emitInstRCPSS(f, inst)
	case x86asm.RCR:
		return d.emitInstRCR(f, inst)
	case x86asm.RDFSBASE:
		return d.emitInstRDFSBASE(f, inst)
	case x86asm.RDGSBASE:
		return d.emitInstRDGSBASE(f, inst)
	case x86asm.RDMSR:
		return d.emitInstRDMSR(f, inst)
	case x86asm.RDPMC:
		return d.emitInstRDPMC(f, inst)
	case x86asm.RDRAND:
		return d.emitInstRDRAND(f, inst)
	case x86asm.RDTSC:
		return d.emitInstRDTSC(f, inst)
	case x86asm.RDTSCP:
		return d.emitInstRDTSCP(f, inst)
	case x86asm.RET:
		return d.emitInstRET(f, inst)
	case x86asm.ROL:
		return d.emitInstROL(f, inst)
	case x86asm.ROR:
		return d.emitInstROR(f, inst)
	case x86asm.ROUNDPD:
		return d.emitInstROUNDPD(f, inst)
	case x86asm.ROUNDPS:
		return d.emitInstROUNDPS(f, inst)
	case x86asm.ROUNDSD:
		return d.emitInstROUNDSD(f, inst)
	case x86asm.ROUNDSS:
		return d.emitInstROUNDSS(f, inst)
	case x86asm.RSM:
		return d.emitInstRSM(f, inst)
	case x86asm.RSQRTPS:
		return d.emitInstRSQRTPS(f, inst)
	case x86asm.RSQRTSS:
		return d.emitInstRSQRTSS(f, inst)
	case x86asm.SAHF:
		return d.emitInstSAHF(f, inst)
	case x86asm.SAR:
		return d.emitInstSAR(f, inst)
	case x86asm.SBB:
		return d.emitInstSBB(f, inst)
	case x86asm.SCASB:
		return d.emitInstSCASB(f, inst)
	case x86asm.SCASD:
		return d.emitInstSCASD(f, inst)
	case x86asm.SCASQ:
		return d.emitInstSCASQ(f, inst)
	case x86asm.SCASW:
		return d.emitInstSCASW(f, inst)
	case x86asm.SETA:
		return d.emitInstSETA(f, inst)
	case x86asm.SETAE:
		return d.emitInstSETAE(f, inst)
	case x86asm.SETB:
		return d.emitInstSETB(f, inst)
	case x86asm.SETBE:
		return d.emitInstSETBE(f, inst)
	case x86asm.SETE:
		return d.emitInstSETE(f, inst)
	case x86asm.SETG:
		return d.emitInstSETG(f, inst)
	case x86asm.SETGE:
		return d.emitInstSETGE(f, inst)
	case x86asm.SETL:
		return d.emitInstSETL(f, inst)
	case x86asm.SETLE:
		return d.emitInstSETLE(f, inst)
	case x86asm.SETNE:
		return d.emitInstSETNE(f, inst)
	case x86asm.SETNO:
		return d.emitInstSETNO(f, inst)
	case x86asm.SETNP:
		return d.emitInstSETNP(f, inst)
	case x86asm.SETNS:
		return d.emitInstSETNS(f, inst)
	case x86asm.SETO:
		return d.emitInstSETO(f, inst)
	case x86asm.SETP:
		return d.emitInstSETP(f, inst)
	case x86asm.SETS:
		return d.emitInstSETS(f, inst)
	case x86asm.SFENCE:
		return d.emitInstSFENCE(f, inst)
	case x86asm.SGDT:
		return d.emitInstSGDT(f, inst)
	case x86asm.SHL:
		return d.emitInstSHL(f, inst)
	case x86asm.SHLD:
		return d.emitInstSHLD(f, inst)
	case x86asm.SHR:
		return d.emitInstSHR(f, inst)
	case x86asm.SHRD:
		return d.emitInstSHRD(f, inst)
	case x86asm.SHUFPD:
		return d.emitInstSHUFPD(f, inst)
	case x86asm.SHUFPS:
		return d.emitInstSHUFPS(f, inst)
	case x86asm.SIDT:
		return d.emitInstSIDT(f, inst)
	case x86asm.SLDT:
		return d.emitInstSLDT(f, inst)
	case x86asm.SMSW:
		return d.emitInstSMSW(f, inst)
	case x86asm.SQRTPD:
		return d.emitInstSQRTPD(f, inst)
	case x86asm.SQRTPS:
		return d.emitInstSQRTPS(f, inst)
	case x86asm.SQRTSD:
		return d.emitInstSQRTSD(f, inst)
	case x86asm.SQRTSS:
		return d.emitInstSQRTSS(f, inst)
	case x86asm.STC:
		return d.emitInstSTC(f, inst)
	case x86asm.STD:
		return d.emitInstSTD(f, inst)
	case x86asm.STI:
		return d.emitInstSTI(f, inst)
	case x86asm.STMXCSR:
		return d.emitInstSTMXCSR(f, inst)
	case x86asm.STOSB:
		return d.emitInstSTOSB(f, inst)
	case x86asm.STOSD:
		return d.emitInstSTOSD(f, inst)
	case x86asm.STOSQ:
		return d.emitInstSTOSQ(f, inst)
	case x86asm.STOSW:
		return d.emitInstSTOSW(f, inst)
	case x86asm.STR:
		return d.emitInstSTR(f, inst)
	case x86asm.SUB:
		return d.emitInstSUB(f, inst)
	case x86asm.SUBPD:
		return d.emitInstSUBPD(f, inst)
	case x86asm.SUBPS:
		return d.emitInstSUBPS(f, inst)
	case x86asm.SUBSD:
		return d.emitInstSUBSD(f, inst)
	case x86asm.SUBSS:
		return d.emitInstSUBSS(f, inst)
	case x86asm.SWAPGS:
		return d.emitInstSWAPGS(f, inst)
	case x86asm.SYSCALL:
		return d.emitInstSYSCALL(f, inst)
	case x86asm.SYSENTER:
		return d.emitInstSYSENTER(f, inst)
	case x86asm.SYSEXIT:
		return d.emitInstSYSEXIT(f, inst)
	case x86asm.SYSRET:
		return d.emitInstSYSRET(f, inst)
	case x86asm.TEST:
		return d.emitInstTEST(f, inst)
	case x86asm.TZCNT:
		return d.emitInstTZCNT(f, inst)
	case x86asm.UCOMISD:
		return d.emitInstUCOMISD(f, inst)
	case x86asm.UCOMISS:
		return d.emitInstUCOMISS(f, inst)
	case x86asm.UD1:
		return d.emitInstUD1(f, inst)
	case x86asm.UD2:
		return d.emitInstUD2(f, inst)
	case x86asm.UNPCKHPD:
		return d.emitInstUNPCKHPD(f, inst)
	case x86asm.UNPCKHPS:
		return d.emitInstUNPCKHPS(f, inst)
	case x86asm.UNPCKLPD:
		return d.emitInstUNPCKLPD(f, inst)
	case x86asm.UNPCKLPS:
		return d.emitInstUNPCKLPS(f, inst)
	case x86asm.VERR:
		return d.emitInstVERR(f, inst)
	case x86asm.VERW:
		return d.emitInstVERW(f, inst)
	case x86asm.VMOVDQA:
		return d.emitInstVMOVDQA(f, inst)
	case x86asm.VMOVDQU:
		return d.emitInstVMOVDQU(f, inst)
	case x86asm.VMOVNTDQ:
		return d.emitInstVMOVNTDQ(f, inst)
	case x86asm.VMOVNTDQA:
		return d.emitInstVMOVNTDQA(f, inst)
	case x86asm.VZEROUPPER:
		return d.emitInstVZEROUPPER(f, inst)
	case x86asm.WBINVD:
		return d.emitInstWBINVD(f, inst)
	case x86asm.WRFSBASE:
		return d.emitInstWRFSBASE(f, inst)
	case x86asm.WRGSBASE:
		return d.emitInstWRGSBASE(f, inst)
	case x86asm.WRMSR:
		return d.emitInstWRMSR(f, inst)
	case x86asm.XABORT:
		return d.emitInstXABORT(f, inst)
	case x86asm.XADD:
		return d.emitInstXADD(f, inst)
	case x86asm.XBEGIN:
		return d.emitInstXBEGIN(f, inst)
	case x86asm.XCHG:
		return d.emitInstXCHG(f, inst)
	case x86asm.XEND:
		return d.emitInstXEND(f, inst)
	case x86asm.XGETBV:
		return d.emitInstXGETBV(f, inst)
	case x86asm.XLATB:
		return d.emitInstXLATB(f, inst)
	case x86asm.XOR:
		return d.emitInstXOR(f, inst)
	case x86asm.XORPD:
		return d.emitInstXORPD(f, inst)
	case x86asm.XORPS:
		return d.emitInstXORPS(f, inst)
	case x86asm.XRSTOR:
		return d.emitInstXRSTOR(f, inst)
	case x86asm.XRSTOR64:
		return d.emitInstXRSTOR64(f, inst)
	case x86asm.XRSTORS:
		return d.emitInstXRSTORS(f, inst)
	case x86asm.XRSTORS64:
		return d.emitInstXRSTORS64(f, inst)
	case x86asm.XSAVE:
		return d.emitInstXSAVE(f, inst)
	case x86asm.XSAVE64:
		return d.emitInstXSAVE64(f, inst)
	case x86asm.XSAVEC:
		return d.emitInstXSAVEC(f, inst)
	case x86asm.XSAVEC64:
		return d.emitInstXSAVEC64(f, inst)
	case x86asm.XSAVEOPT:
		return d.emitInstXSAVEOPT(f, inst)
	case x86asm.XSAVEOPT64:
		return d.emitInstXSAVEOPT64(f, inst)
	case x86asm.XSAVES:
		return d.emitInstXSAVES(f, inst)
	case x86asm.XSAVES64:
		return d.emitInstXSAVES64(f, inst)
	case x86asm.XSETBV:
		return d.emitInstXSETBV(f, inst)
	case x86asm.XTEST:
		return d.emitInstXTEST(f, inst)
	default:
		panic(fmt.Errorf("support for x86 instruction opcode %v not yet implemented", inst.Op))
	}
}

// --- [ AAA ] -----------------------------------------------------------------

func (d *disassembler) emitInstAAA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AAD ] -----------------------------------------------------------------

func (d *disassembler) emitInstAAD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AAM ] -----------------------------------------------------------------

func (d *disassembler) emitInstAAM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AAS ] -----------------------------------------------------------------

func (d *disassembler) emitInstAAS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADC ] -----------------------------------------------------------------

func (d *disassembler) emitInstADC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADD ] -----------------------------------------------------------------

func (d *disassembler) emitInstADD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDPD ] ---------------------------------------------------------------

func (d *disassembler) emitInstADDPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDPS ] ---------------------------------------------------------------

func (d *disassembler) emitInstADDPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstADDSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDSS ] ---------------------------------------------------------------

func (d *disassembler) emitInstADDSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDSUBPD ] ------------------------------------------------------------

func (d *disassembler) emitInstADDSUBPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDSUBPS ] ------------------------------------------------------------

func (d *disassembler) emitInstADDSUBPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESDEC ] --------------------------------------------------------------

func (d *disassembler) emitInstAESDEC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESDECLAST ] ----------------------------------------------------------

func (d *disassembler) emitInstAESDECLAST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESENC ] --------------------------------------------------------------

func (d *disassembler) emitInstAESENC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESENCLAST ] ----------------------------------------------------------

func (d *disassembler) emitInstAESENCLAST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESIMC ] --------------------------------------------------------------

func (d *disassembler) emitInstAESIMC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESKEYGENASSIST ] -----------------------------------------------------

func (d *disassembler) emitInstAESKEYGENASSIST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AND ] -----------------------------------------------------------------

func (d *disassembler) emitInstAND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ANDNPD ] --------------------------------------------------------------

func (d *disassembler) emitInstANDNPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ANDNPS ] --------------------------------------------------------------

func (d *disassembler) emitInstANDNPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ANDPD ] ---------------------------------------------------------------

func (d *disassembler) emitInstANDPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ANDPS ] ---------------------------------------------------------------

func (d *disassembler) emitInstANDPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ARPL ] ----------------------------------------------------------------

func (d *disassembler) emitInstARPL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BLENDPD ] -------------------------------------------------------------

func (d *disassembler) emitInstBLENDPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BLENDPS ] -------------------------------------------------------------

func (d *disassembler) emitInstBLENDPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BLENDVPD ] ------------------------------------------------------------

func (d *disassembler) emitInstBLENDVPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BLENDVPS ] ------------------------------------------------------------

func (d *disassembler) emitInstBLENDVPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BOUND ] ---------------------------------------------------------------

func (d *disassembler) emitInstBOUND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BSF ] -----------------------------------------------------------------

func (d *disassembler) emitInstBSF(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BSR ] -----------------------------------------------------------------

func (d *disassembler) emitInstBSR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BSWAP ] ---------------------------------------------------------------

func (d *disassembler) emitInstBSWAP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BT ] ------------------------------------------------------------------

func (d *disassembler) emitInstBT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BTC ] -----------------------------------------------------------------

func (d *disassembler) emitInstBTC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BTR ] -----------------------------------------------------------------

func (d *disassembler) emitInstBTR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BTS ] -----------------------------------------------------------------

func (d *disassembler) emitInstBTS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CALL ] ----------------------------------------------------------------

func (d *disassembler) emitInstCALL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CBW ] -----------------------------------------------------------------

func (d *disassembler) emitInstCBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CDQ ] -----------------------------------------------------------------

func (d *disassembler) emitInstCDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CDQE ] ----------------------------------------------------------------

func (d *disassembler) emitInstCDQE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CLC ] -----------------------------------------------------------------

func (d *disassembler) emitInstCLC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CLD ] -----------------------------------------------------------------

func (d *disassembler) emitInstCLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CLFLUSH ] -------------------------------------------------------------

func (d *disassembler) emitInstCLFLUSH(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CLI ] -----------------------------------------------------------------

func (d *disassembler) emitInstCLI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CLTS ] ----------------------------------------------------------------

func (d *disassembler) emitInstCLTS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMC ] -----------------------------------------------------------------

func (d *disassembler) emitInstCMC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVA ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMOVA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVAE ] --------------------------------------------------------------

func (d *disassembler) emitInstCMOVAE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVB ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMOVB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVBE ] --------------------------------------------------------------

func (d *disassembler) emitInstCMOVBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVE ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMOVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVG ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMOVG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVGE ] --------------------------------------------------------------

func (d *disassembler) emitInstCMOVGE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVL ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMOVL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVLE ] --------------------------------------------------------------

func (d *disassembler) emitInstCMOVLE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVNE ] --------------------------------------------------------------

func (d *disassembler) emitInstCMOVNE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVNO ] --------------------------------------------------------------

func (d *disassembler) emitInstCMOVNO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVNP ] --------------------------------------------------------------

func (d *disassembler) emitInstCMOVNP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVNS ] --------------------------------------------------------------

func (d *disassembler) emitInstCMOVNS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVO ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMOVO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVP ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMOVP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVS ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMOVS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMP ] -----------------------------------------------------------------

func (d *disassembler) emitInstCMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPPD ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMPPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPPS ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMPPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSB ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMPSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMPSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSD_XMM ] -----------------------------------------------------------

func (d *disassembler) emitInstCMPSD_XMM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMPSQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSS ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMPSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSW ] ---------------------------------------------------------------

func (d *disassembler) emitInstCMPSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPXCHG ] -------------------------------------------------------------

func (d *disassembler) emitInstCMPXCHG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPXCHG16B ] ----------------------------------------------------------

func (d *disassembler) emitInstCMPXCHG16B(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPXCHG8B ] -----------------------------------------------------------

func (d *disassembler) emitInstCMPXCHG8B(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ COMISD ] --------------------------------------------------------------

func (d *disassembler) emitInstCOMISD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ COMISS ] --------------------------------------------------------------

func (d *disassembler) emitInstCOMISS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CPUID ] ---------------------------------------------------------------

func (d *disassembler) emitInstCPUID(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CQO ] -----------------------------------------------------------------

func (d *disassembler) emitInstCQO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CRC32 ] ---------------------------------------------------------------

func (d *disassembler) emitInstCRC32(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTDQ2PD ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTDQ2PD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTDQ2PS ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTDQ2PS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPD2DQ ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTPD2DQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPD2PI ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTPD2PI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPD2PS ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTPD2PS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPI2PD ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTPI2PD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPI2PS ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTPI2PS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPS2DQ ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTPS2DQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPS2PD ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTPS2PD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPS2PI ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTPS2PI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSD2SI ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTSD2SI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSD2SS ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTSD2SS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSI2SD ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTSI2SD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSI2SS ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTSI2SS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSS2SD ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTSS2SD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSS2SI ] ------------------------------------------------------------

func (d *disassembler) emitInstCVTSS2SI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTPD2DQ ] -----------------------------------------------------------

func (d *disassembler) emitInstCVTTPD2DQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTPD2PI ] -----------------------------------------------------------

func (d *disassembler) emitInstCVTTPD2PI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTPS2DQ ] -----------------------------------------------------------

func (d *disassembler) emitInstCVTTPS2DQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTPS2PI ] -----------------------------------------------------------

func (d *disassembler) emitInstCVTTPS2PI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTSD2SI ] -----------------------------------------------------------

func (d *disassembler) emitInstCVTTSD2SI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTSS2SI ] -----------------------------------------------------------

func (d *disassembler) emitInstCVTTSS2SI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CWD ] -----------------------------------------------------------------

func (d *disassembler) emitInstCWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CWDE ] ----------------------------------------------------------------

func (d *disassembler) emitInstCWDE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DAA ] -----------------------------------------------------------------

func (d *disassembler) emitInstDAA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DAS ] -----------------------------------------------------------------

func (d *disassembler) emitInstDAS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DEC ] -----------------------------------------------------------------

func (d *disassembler) emitInstDEC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DIV ] -----------------------------------------------------------------

func (d *disassembler) emitInstDIV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DIVPD ] ---------------------------------------------------------------

func (d *disassembler) emitInstDIVPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DIVPS ] ---------------------------------------------------------------

func (d *disassembler) emitInstDIVPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DIVSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstDIVSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DIVSS ] ---------------------------------------------------------------

func (d *disassembler) emitInstDIVSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DPPD ] ----------------------------------------------------------------

func (d *disassembler) emitInstDPPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DPPS ] ----------------------------------------------------------------

func (d *disassembler) emitInstDPPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ EMMS ] ----------------------------------------------------------------

func (d *disassembler) emitInstEMMS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ENTER ] ---------------------------------------------------------------

func (d *disassembler) emitInstENTER(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ EXTRACTPS ] -----------------------------------------------------------

func (d *disassembler) emitInstEXTRACTPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ F2XM1 ] ---------------------------------------------------------------

func (d *disassembler) emitInstF2XM1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FABS ] ----------------------------------------------------------------

func (d *disassembler) emitInstFABS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FADD ] ----------------------------------------------------------------

func (d *disassembler) emitInstFADD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FADDP ] ---------------------------------------------------------------

func (d *disassembler) emitInstFADDP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FBLD ] ----------------------------------------------------------------

func (d *disassembler) emitInstFBLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FBSTP ] ---------------------------------------------------------------

func (d *disassembler) emitInstFBSTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCHS ] ----------------------------------------------------------------

func (d *disassembler) emitInstFCHS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVB ] --------------------------------------------------------------

func (d *disassembler) emitInstFCMOVB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVBE ] -------------------------------------------------------------

func (d *disassembler) emitInstFCMOVBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVE ] --------------------------------------------------------------

func (d *disassembler) emitInstFCMOVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVNB ] -------------------------------------------------------------

func (d *disassembler) emitInstFCMOVNB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVNBE ] ------------------------------------------------------------

func (d *disassembler) emitInstFCMOVNBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVNE ] -------------------------------------------------------------

func (d *disassembler) emitInstFCMOVNE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVNU ] -------------------------------------------------------------

func (d *disassembler) emitInstFCMOVNU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVU ] --------------------------------------------------------------

func (d *disassembler) emitInstFCMOVU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOM ] ----------------------------------------------------------------

func (d *disassembler) emitInstFCOM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOMI ] ---------------------------------------------------------------

func (d *disassembler) emitInstFCOMI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOMIP ] --------------------------------------------------------------

func (d *disassembler) emitInstFCOMIP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOMP ] ---------------------------------------------------------------

func (d *disassembler) emitInstFCOMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOMPP ] --------------------------------------------------------------

func (d *disassembler) emitInstFCOMPP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOS ] ----------------------------------------------------------------

func (d *disassembler) emitInstFCOS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FDECSTP ] -------------------------------------------------------------

func (d *disassembler) emitInstFDECSTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FDIV ] ----------------------------------------------------------------

func (d *disassembler) emitInstFDIV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FDIVP ] ---------------------------------------------------------------

func (d *disassembler) emitInstFDIVP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FDIVR ] ---------------------------------------------------------------

func (d *disassembler) emitInstFDIVR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FDIVRP ] --------------------------------------------------------------

func (d *disassembler) emitInstFDIVRP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FFREE ] ---------------------------------------------------------------

func (d *disassembler) emitInstFFREE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FFREEP ] --------------------------------------------------------------

func (d *disassembler) emitInstFFREEP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FIADD ] ---------------------------------------------------------------

func (d *disassembler) emitInstFIADD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FICOM ] ---------------------------------------------------------------

func (d *disassembler) emitInstFICOM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FICOMP ] --------------------------------------------------------------

func (d *disassembler) emitInstFICOMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FIDIV ] ---------------------------------------------------------------

func (d *disassembler) emitInstFIDIV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FIDIVR ] --------------------------------------------------------------

func (d *disassembler) emitInstFIDIVR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FILD ] ----------------------------------------------------------------

func (d *disassembler) emitInstFILD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FIMUL ] ---------------------------------------------------------------

func (d *disassembler) emitInstFIMUL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FINCSTP ] -------------------------------------------------------------

func (d *disassembler) emitInstFINCSTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FIST ] ----------------------------------------------------------------

func (d *disassembler) emitInstFIST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FISTP ] ---------------------------------------------------------------

func (d *disassembler) emitInstFISTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FISTTP ] --------------------------------------------------------------

func (d *disassembler) emitInstFISTTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FISUB ] ---------------------------------------------------------------

func (d *disassembler) emitInstFISUB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FISUBR ] --------------------------------------------------------------

func (d *disassembler) emitInstFISUBR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLD ] -----------------------------------------------------------------

func (d *disassembler) emitInstFLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLD1 ] ----------------------------------------------------------------

func (d *disassembler) emitInstFLD1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDCW ] ---------------------------------------------------------------

func (d *disassembler) emitInstFLDCW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDENV ] --------------------------------------------------------------

func (d *disassembler) emitInstFLDENV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDL2E ] --------------------------------------------------------------

func (d *disassembler) emitInstFLDL2E(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDL2T ] --------------------------------------------------------------

func (d *disassembler) emitInstFLDL2T(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDLG2 ] --------------------------------------------------------------

func (d *disassembler) emitInstFLDLG2(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDLN2 ] --------------------------------------------------------------

func (d *disassembler) emitInstFLDLN2(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDPI ] ---------------------------------------------------------------

func (d *disassembler) emitInstFLDPI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDZ ] ----------------------------------------------------------------

func (d *disassembler) emitInstFLDZ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FMUL ] ----------------------------------------------------------------

func (d *disassembler) emitInstFMUL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FMULP ] ---------------------------------------------------------------

func (d *disassembler) emitInstFMULP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNCLEX ] --------------------------------------------------------------

func (d *disassembler) emitInstFNCLEX(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNINIT ] --------------------------------------------------------------

func (d *disassembler) emitInstFNINIT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNOP ] ----------------------------------------------------------------

func (d *disassembler) emitInstFNOP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNSAVE ] --------------------------------------------------------------

func (d *disassembler) emitInstFNSAVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNSTCW ] --------------------------------------------------------------

func (d *disassembler) emitInstFNSTCW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNSTENV ] -------------------------------------------------------------

func (d *disassembler) emitInstFNSTENV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNSTSW ] --------------------------------------------------------------

func (d *disassembler) emitInstFNSTSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FPATAN ] --------------------------------------------------------------

func (d *disassembler) emitInstFPATAN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FPREM ] ---------------------------------------------------------------

func (d *disassembler) emitInstFPREM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FPREM1 ] --------------------------------------------------------------

func (d *disassembler) emitInstFPREM1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FPTAN ] ---------------------------------------------------------------

func (d *disassembler) emitInstFPTAN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FRNDINT ] -------------------------------------------------------------

func (d *disassembler) emitInstFRNDINT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FRSTOR ] --------------------------------------------------------------

func (d *disassembler) emitInstFRSTOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSCALE ] --------------------------------------------------------------

func (d *disassembler) emitInstFSCALE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSIN ] ----------------------------------------------------------------

func (d *disassembler) emitInstFSIN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSINCOS ] -------------------------------------------------------------

func (d *disassembler) emitInstFSINCOS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSQRT ] ---------------------------------------------------------------

func (d *disassembler) emitInstFSQRT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FST ] -----------------------------------------------------------------

func (d *disassembler) emitInstFST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSTP ] ----------------------------------------------------------------

func (d *disassembler) emitInstFSTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSUB ] ----------------------------------------------------------------

func (d *disassembler) emitInstFSUB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSUBP ] ---------------------------------------------------------------

func (d *disassembler) emitInstFSUBP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSUBR ] ---------------------------------------------------------------

func (d *disassembler) emitInstFSUBR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSUBRP ] --------------------------------------------------------------

func (d *disassembler) emitInstFSUBRP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FTST ] ----------------------------------------------------------------

func (d *disassembler) emitInstFTST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FUCOM ] ---------------------------------------------------------------

func (d *disassembler) emitInstFUCOM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FUCOMI ] --------------------------------------------------------------

func (d *disassembler) emitInstFUCOMI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FUCOMIP ] -------------------------------------------------------------

func (d *disassembler) emitInstFUCOMIP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FUCOMP ] --------------------------------------------------------------

func (d *disassembler) emitInstFUCOMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FUCOMPP ] -------------------------------------------------------------

func (d *disassembler) emitInstFUCOMPP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FWAIT ] ---------------------------------------------------------------

func (d *disassembler) emitInstFWAIT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXAM ] ----------------------------------------------------------------

func (d *disassembler) emitInstFXAM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXCH ] ----------------------------------------------------------------

func (d *disassembler) emitInstFXCH(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXRSTOR ] -------------------------------------------------------------

func (d *disassembler) emitInstFXRSTOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXRSTOR64 ] -----------------------------------------------------------

func (d *disassembler) emitInstFXRSTOR64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXSAVE ] --------------------------------------------------------------

func (d *disassembler) emitInstFXSAVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXSAVE64 ] ------------------------------------------------------------

func (d *disassembler) emitInstFXSAVE64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXTRACT ] -------------------------------------------------------------

func (d *disassembler) emitInstFXTRACT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FYL2X ] ---------------------------------------------------------------

func (d *disassembler) emitInstFYL2X(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FYL2XP1 ] -------------------------------------------------------------

func (d *disassembler) emitInstFYL2XP1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ HADDPD ] --------------------------------------------------------------

func (d *disassembler) emitInstHADDPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ HADDPS ] --------------------------------------------------------------

func (d *disassembler) emitInstHADDPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ HLT ] -----------------------------------------------------------------

func (d *disassembler) emitInstHLT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ HSUBPD ] --------------------------------------------------------------

func (d *disassembler) emitInstHSUBPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ HSUBPS ] --------------------------------------------------------------

func (d *disassembler) emitInstHSUBPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ICEBP ] ---------------------------------------------------------------

func (d *disassembler) emitInstICEBP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IDIV ] ----------------------------------------------------------------

func (d *disassembler) emitInstIDIV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IMUL ] ----------------------------------------------------------------

func (d *disassembler) emitInstIMUL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IN ] ------------------------------------------------------------------

func (d *disassembler) emitInstIN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INC ] -----------------------------------------------------------------

func (d *disassembler) emitInstINC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INSB ] ----------------------------------------------------------------

func (d *disassembler) emitInstINSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INSD ] ----------------------------------------------------------------

func (d *disassembler) emitInstINSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INSERTPS ] ------------------------------------------------------------

func (d *disassembler) emitInstINSERTPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INSW ] ----------------------------------------------------------------

func (d *disassembler) emitInstINSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INT ] -----------------------------------------------------------------

func (d *disassembler) emitInstINT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INTO ] ----------------------------------------------------------------

func (d *disassembler) emitInstINTO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INVD ] ----------------------------------------------------------------

func (d *disassembler) emitInstINVD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INVLPG ] --------------------------------------------------------------

func (d *disassembler) emitInstINVLPG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INVPCID ] -------------------------------------------------------------

func (d *disassembler) emitInstINVPCID(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IRET ] ----------------------------------------------------------------

func (d *disassembler) emitInstIRET(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IRETD ] ---------------------------------------------------------------

func (d *disassembler) emitInstIRETD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IRETQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstIRETQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JA ] ------------------------------------------------------------------

func (d *disassembler) emitInstJA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JAE ] -----------------------------------------------------------------

func (d *disassembler) emitInstJAE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JB ] ------------------------------------------------------------------

func (d *disassembler) emitInstJB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JBE ] -----------------------------------------------------------------

func (d *disassembler) emitInstJBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JCXZ ] ----------------------------------------------------------------

func (d *disassembler) emitInstJCXZ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JE ] ------------------------------------------------------------------

func (d *disassembler) emitInstJE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JECXZ ] ---------------------------------------------------------------

func (d *disassembler) emitInstJECXZ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JG ] ------------------------------------------------------------------

func (d *disassembler) emitInstJG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JGE ] -----------------------------------------------------------------

func (d *disassembler) emitInstJGE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JL ] ------------------------------------------------------------------

func (d *disassembler) emitInstJL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JLE ] -----------------------------------------------------------------

func (d *disassembler) emitInstJLE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JMP ] -----------------------------------------------------------------

func (d *disassembler) emitInstJMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JNE ] -----------------------------------------------------------------

func (d *disassembler) emitInstJNE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JNO ] -----------------------------------------------------------------

func (d *disassembler) emitInstJNO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JNP ] -----------------------------------------------------------------

func (d *disassembler) emitInstJNP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JNS ] -----------------------------------------------------------------

func (d *disassembler) emitInstJNS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JO ] ------------------------------------------------------------------

func (d *disassembler) emitInstJO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JP ] ------------------------------------------------------------------

func (d *disassembler) emitInstJP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JRCXZ ] ---------------------------------------------------------------

func (d *disassembler) emitInstJRCXZ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JS ] ------------------------------------------------------------------

func (d *disassembler) emitInstJS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LAHF ] ----------------------------------------------------------------

func (d *disassembler) emitInstLAHF(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LAR ] -----------------------------------------------------------------

func (d *disassembler) emitInstLAR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LCALL ] ---------------------------------------------------------------

func (d *disassembler) emitInstLCALL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LDDQU ] ---------------------------------------------------------------

func (d *disassembler) emitInstLDDQU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LDMXCSR ] -------------------------------------------------------------

func (d *disassembler) emitInstLDMXCSR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LDS ] -----------------------------------------------------------------

func (d *disassembler) emitInstLDS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LEA ] -----------------------------------------------------------------

func (d *disassembler) emitInstLEA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LEAVE ] ---------------------------------------------------------------

func (d *disassembler) emitInstLEAVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LES ] -----------------------------------------------------------------

func (d *disassembler) emitInstLES(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LFENCE ] --------------------------------------------------------------

func (d *disassembler) emitInstLFENCE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LFS ] -----------------------------------------------------------------

func (d *disassembler) emitInstLFS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LGDT ] ----------------------------------------------------------------

func (d *disassembler) emitInstLGDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LGS ] -----------------------------------------------------------------

func (d *disassembler) emitInstLGS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LIDT ] ----------------------------------------------------------------

func (d *disassembler) emitInstLIDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LJMP ] ----------------------------------------------------------------

func (d *disassembler) emitInstLJMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LLDT ] ----------------------------------------------------------------

func (d *disassembler) emitInstLLDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LMSW ] ----------------------------------------------------------------

func (d *disassembler) emitInstLMSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LODSB ] ---------------------------------------------------------------

func (d *disassembler) emitInstLODSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LODSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstLODSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LODSQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstLODSQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LODSW ] ---------------------------------------------------------------

func (d *disassembler) emitInstLODSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LOOP ] ----------------------------------------------------------------

func (d *disassembler) emitInstLOOP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LOOPE ] ---------------------------------------------------------------

func (d *disassembler) emitInstLOOPE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LOOPNE ] --------------------------------------------------------------

func (d *disassembler) emitInstLOOPNE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LRET ] ----------------------------------------------------------------

func (d *disassembler) emitInstLRET(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LSL ] -----------------------------------------------------------------

func (d *disassembler) emitInstLSL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LSS ] -----------------------------------------------------------------

func (d *disassembler) emitInstLSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LTR ] -----------------------------------------------------------------

func (d *disassembler) emitInstLTR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LZCNT ] ---------------------------------------------------------------

func (d *disassembler) emitInstLZCNT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MASKMOVDQU ] ----------------------------------------------------------

func (d *disassembler) emitInstMASKMOVDQU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MASKMOVQ ] ------------------------------------------------------------

func (d *disassembler) emitInstMASKMOVQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MAXPD ] ---------------------------------------------------------------

func (d *disassembler) emitInstMAXPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MAXPS ] ---------------------------------------------------------------

func (d *disassembler) emitInstMAXPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MAXSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstMAXSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MAXSS ] ---------------------------------------------------------------

func (d *disassembler) emitInstMAXSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MFENCE ] --------------------------------------------------------------

func (d *disassembler) emitInstMFENCE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MINPD ] ---------------------------------------------------------------

func (d *disassembler) emitInstMINPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MINPS ] ---------------------------------------------------------------

func (d *disassembler) emitInstMINPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MINSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstMINSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MINSS ] ---------------------------------------------------------------

func (d *disassembler) emitInstMINSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MONITOR ] -------------------------------------------------------------

func (d *disassembler) emitInstMONITOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOV ] -----------------------------------------------------------------

func (d *disassembler) emitInstMOV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVAPD ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVAPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVAPS ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVAPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVBE ] ---------------------------------------------------------------

func (d *disassembler) emitInstMOVBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVD ] ----------------------------------------------------------------

func (d *disassembler) emitInstMOVD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVDDUP ] -------------------------------------------------------------

func (d *disassembler) emitInstMOVDDUP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVDQ2Q ] -------------------------------------------------------------

func (d *disassembler) emitInstMOVDQ2Q(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVDQA ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVDQA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVDQU ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVDQU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVHLPS ] -------------------------------------------------------------

func (d *disassembler) emitInstMOVHLPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVHPD ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVHPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVHPS ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVHPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVLHPS ] -------------------------------------------------------------

func (d *disassembler) emitInstMOVLHPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVLPD ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVLPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVLPS ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVLPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVMSKPD ] ------------------------------------------------------------

func (d *disassembler) emitInstMOVMSKPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVMSKPS ] ------------------------------------------------------------

func (d *disassembler) emitInstMOVMSKPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTDQ ] -------------------------------------------------------------

func (d *disassembler) emitInstMOVNTDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTDQA ] ------------------------------------------------------------

func (d *disassembler) emitInstMOVNTDQA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTI ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVNTI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTPD ] -------------------------------------------------------------

func (d *disassembler) emitInstMOVNTPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTPS ] -------------------------------------------------------------

func (d *disassembler) emitInstMOVNTPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTQ ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVNTQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTSD ] -------------------------------------------------------------

func (d *disassembler) emitInstMOVNTSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTSS ] -------------------------------------------------------------

func (d *disassembler) emitInstMOVNTSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVQ ] ----------------------------------------------------------------

func (d *disassembler) emitInstMOVQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVQ2DQ ] -------------------------------------------------------------

func (d *disassembler) emitInstMOVQ2DQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSB ] ---------------------------------------------------------------

func (d *disassembler) emitInstMOVSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstMOVSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSD_XMM ] -----------------------------------------------------------

func (d *disassembler) emitInstMOVSD_XMM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSHDUP ] ------------------------------------------------------------

func (d *disassembler) emitInstMOVSHDUP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSLDUP ] ------------------------------------------------------------

func (d *disassembler) emitInstMOVSLDUP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstMOVSQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSS ] ---------------------------------------------------------------

func (d *disassembler) emitInstMOVSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSW ] ---------------------------------------------------------------

func (d *disassembler) emitInstMOVSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSX ] ---------------------------------------------------------------

func (d *disassembler) emitInstMOVSX(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSXD ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVSXD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVUPD ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVUPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVUPS ] --------------------------------------------------------------

func (d *disassembler) emitInstMOVUPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVZX ] ---------------------------------------------------------------

func (d *disassembler) emitInstMOVZX(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MPSADBW ] -------------------------------------------------------------

func (d *disassembler) emitInstMPSADBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MUL ] -----------------------------------------------------------------

func (d *disassembler) emitInstMUL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MULPD ] ---------------------------------------------------------------

func (d *disassembler) emitInstMULPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MULPS ] ---------------------------------------------------------------

func (d *disassembler) emitInstMULPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MULSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstMULSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MULSS ] ---------------------------------------------------------------

func (d *disassembler) emitInstMULSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MWAIT ] ---------------------------------------------------------------

func (d *disassembler) emitInstMWAIT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ NEG ] -----------------------------------------------------------------

func (d *disassembler) emitInstNEG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ NOP ] -----------------------------------------------------------------

func (d *disassembler) emitInstNOP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ NOT ] -----------------------------------------------------------------

func (d *disassembler) emitInstNOT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ OR ] ------------------------------------------------------------------

func (d *disassembler) emitInstOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ORPD ] ----------------------------------------------------------------

func (d *disassembler) emitInstORPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ORPS ] ----------------------------------------------------------------

func (d *disassembler) emitInstORPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ OUT ] -----------------------------------------------------------------

func (d *disassembler) emitInstOUT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ OUTSB ] ---------------------------------------------------------------

func (d *disassembler) emitInstOUTSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ OUTSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstOUTSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ OUTSW ] ---------------------------------------------------------------

func (d *disassembler) emitInstOUTSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PABSB ] ---------------------------------------------------------------

func (d *disassembler) emitInstPABSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PABSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstPABSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PABSW ] ---------------------------------------------------------------

func (d *disassembler) emitInstPABSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PACKSSDW ] ------------------------------------------------------------

func (d *disassembler) emitInstPACKSSDW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PACKSSWB ] ------------------------------------------------------------

func (d *disassembler) emitInstPACKSSWB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PACKUSDW ] ------------------------------------------------------------

func (d *disassembler) emitInstPACKUSDW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PACKUSWB ] ------------------------------------------------------------

func (d *disassembler) emitInstPACKUSWB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDB ] ---------------------------------------------------------------

func (d *disassembler) emitInstPADDB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDD ] ---------------------------------------------------------------

func (d *disassembler) emitInstPADDD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstPADDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDSB ] --------------------------------------------------------------

func (d *disassembler) emitInstPADDSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDSW ] --------------------------------------------------------------

func (d *disassembler) emitInstPADDSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDUSB ] -------------------------------------------------------------

func (d *disassembler) emitInstPADDUSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDUSW ] -------------------------------------------------------------

func (d *disassembler) emitInstPADDUSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDW ] ---------------------------------------------------------------

func (d *disassembler) emitInstPADDW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PALIGNR ] -------------------------------------------------------------

func (d *disassembler) emitInstPALIGNR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PAND ] ----------------------------------------------------------------

func (d *disassembler) emitInstPAND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PANDN ] ---------------------------------------------------------------

func (d *disassembler) emitInstPANDN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PAUSE ] ---------------------------------------------------------------

func (d *disassembler) emitInstPAUSE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PAVGB ] ---------------------------------------------------------------

func (d *disassembler) emitInstPAVGB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PAVGW ] ---------------------------------------------------------------

func (d *disassembler) emitInstPAVGW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PBLENDVB ] ------------------------------------------------------------

func (d *disassembler) emitInstPBLENDVB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PBLENDW ] -------------------------------------------------------------

func (d *disassembler) emitInstPBLENDW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCLMULQDQ ] -----------------------------------------------------------

func (d *disassembler) emitInstPCLMULQDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPEQB ] -------------------------------------------------------------

func (d *disassembler) emitInstPCMPEQB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPEQD ] -------------------------------------------------------------

func (d *disassembler) emitInstPCMPEQD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPEQQ ] -------------------------------------------------------------

func (d *disassembler) emitInstPCMPEQQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPEQW ] -------------------------------------------------------------

func (d *disassembler) emitInstPCMPEQW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPESTRI ] -----------------------------------------------------------

func (d *disassembler) emitInstPCMPESTRI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPESTRM ] -----------------------------------------------------------

func (d *disassembler) emitInstPCMPESTRM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPGTB ] -------------------------------------------------------------

func (d *disassembler) emitInstPCMPGTB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPGTD ] -------------------------------------------------------------

func (d *disassembler) emitInstPCMPGTD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPGTQ ] -------------------------------------------------------------

func (d *disassembler) emitInstPCMPGTQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPGTW ] -------------------------------------------------------------

func (d *disassembler) emitInstPCMPGTW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPISTRI ] -----------------------------------------------------------

func (d *disassembler) emitInstPCMPISTRI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPISTRM ] -----------------------------------------------------------

func (d *disassembler) emitInstPCMPISTRM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PEXTRB ] --------------------------------------------------------------

func (d *disassembler) emitInstPEXTRB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PEXTRD ] --------------------------------------------------------------

func (d *disassembler) emitInstPEXTRD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PEXTRQ ] --------------------------------------------------------------

func (d *disassembler) emitInstPEXTRQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PEXTRW ] --------------------------------------------------------------

func (d *disassembler) emitInstPEXTRW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHADDD ] --------------------------------------------------------------

func (d *disassembler) emitInstPHADDD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHADDSW ] -------------------------------------------------------------

func (d *disassembler) emitInstPHADDSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHADDW ] --------------------------------------------------------------

func (d *disassembler) emitInstPHADDW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHMINPOSUW ] ----------------------------------------------------------

func (d *disassembler) emitInstPHMINPOSUW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHSUBD ] --------------------------------------------------------------

func (d *disassembler) emitInstPHSUBD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHSUBSW ] -------------------------------------------------------------

func (d *disassembler) emitInstPHSUBSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHSUBW ] --------------------------------------------------------------

func (d *disassembler) emitInstPHSUBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PINSRB ] --------------------------------------------------------------

func (d *disassembler) emitInstPINSRB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PINSRD ] --------------------------------------------------------------

func (d *disassembler) emitInstPINSRD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PINSRQ ] --------------------------------------------------------------

func (d *disassembler) emitInstPINSRQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PINSRW ] --------------------------------------------------------------

func (d *disassembler) emitInstPINSRW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMADDUBSW ] -----------------------------------------------------------

func (d *disassembler) emitInstPMADDUBSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMADDWD ] -------------------------------------------------------------

func (d *disassembler) emitInstPMADDWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXSB ] --------------------------------------------------------------

func (d *disassembler) emitInstPMAXSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXSD ] --------------------------------------------------------------

func (d *disassembler) emitInstPMAXSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXSW ] --------------------------------------------------------------

func (d *disassembler) emitInstPMAXSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXUB ] --------------------------------------------------------------

func (d *disassembler) emitInstPMAXUB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXUD ] --------------------------------------------------------------

func (d *disassembler) emitInstPMAXUD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXUW ] --------------------------------------------------------------

func (d *disassembler) emitInstPMAXUW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINSB ] --------------------------------------------------------------

func (d *disassembler) emitInstPMINSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINSD ] --------------------------------------------------------------

func (d *disassembler) emitInstPMINSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINSW ] --------------------------------------------------------------

func (d *disassembler) emitInstPMINSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINUB ] --------------------------------------------------------------

func (d *disassembler) emitInstPMINUB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINUD ] --------------------------------------------------------------

func (d *disassembler) emitInstPMINUD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINUW ] --------------------------------------------------------------

func (d *disassembler) emitInstPMINUW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVMSKB ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVMSKB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXBD ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVSXBD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXBQ ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVSXBQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXBW ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVSXBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXDQ ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVSXDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXWD ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVSXWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXWQ ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVSXWQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXBD ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVZXBD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXBQ ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVZXBQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXBW ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVZXBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXDQ ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVZXDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXWD ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVZXWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXWQ ] ------------------------------------------------------------

func (d *disassembler) emitInstPMOVZXWQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULDQ ] --------------------------------------------------------------

func (d *disassembler) emitInstPMULDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULHRSW ] ------------------------------------------------------------

func (d *disassembler) emitInstPMULHRSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULHUW ] -------------------------------------------------------------

func (d *disassembler) emitInstPMULHUW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULHW ] --------------------------------------------------------------

func (d *disassembler) emitInstPMULHW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULLD ] --------------------------------------------------------------

func (d *disassembler) emitInstPMULLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULLW ] --------------------------------------------------------------

func (d *disassembler) emitInstPMULLW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULUDQ ] -------------------------------------------------------------

func (d *disassembler) emitInstPMULUDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POP ] -----------------------------------------------------------------

func (d *disassembler) emitInstPOP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPA ] ----------------------------------------------------------------

func (d *disassembler) emitInstPOPA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPAD ] ---------------------------------------------------------------

func (d *disassembler) emitInstPOPAD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPCNT ] --------------------------------------------------------------

func (d *disassembler) emitInstPOPCNT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPF ] ----------------------------------------------------------------

func (d *disassembler) emitInstPOPF(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPFD ] ---------------------------------------------------------------

func (d *disassembler) emitInstPOPFD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPFQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstPOPFQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POR ] -----------------------------------------------------------------

func (d *disassembler) emitInstPOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PREFETCHNTA ] ---------------------------------------------------------

func (d *disassembler) emitInstPREFETCHNTA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PREFETCHT0 ] ----------------------------------------------------------

func (d *disassembler) emitInstPREFETCHT0(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PREFETCHT1 ] ----------------------------------------------------------

func (d *disassembler) emitInstPREFETCHT1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PREFETCHT2 ] ----------------------------------------------------------

func (d *disassembler) emitInstPREFETCHT2(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PREFETCHW ] -----------------------------------------------------------

func (d *disassembler) emitInstPREFETCHW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSADBW ] --------------------------------------------------------------

func (d *disassembler) emitInstPSADBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSHUFB ] --------------------------------------------------------------

func (d *disassembler) emitInstPSHUFB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSHUFD ] --------------------------------------------------------------

func (d *disassembler) emitInstPSHUFD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSHUFHW ] -------------------------------------------------------------

func (d *disassembler) emitInstPSHUFHW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSHUFLW ] -------------------------------------------------------------

func (d *disassembler) emitInstPSHUFLW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSHUFW ] --------------------------------------------------------------

func (d *disassembler) emitInstPSHUFW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSIGNB ] --------------------------------------------------------------

func (d *disassembler) emitInstPSIGNB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSIGND ] --------------------------------------------------------------

func (d *disassembler) emitInstPSIGND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSIGNW ] --------------------------------------------------------------

func (d *disassembler) emitInstPSIGNW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSLLD ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSLLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSLLDQ ] --------------------------------------------------------------

func (d *disassembler) emitInstPSLLDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSLLQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSLLQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSLLW ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSLLW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRAD ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSRAD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRAW ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSRAW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRLD ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSRLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRLDQ ] --------------------------------------------------------------

func (d *disassembler) emitInstPSRLDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRLQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSRLQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRLW ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSRLW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBB ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSUBB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBD ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSUBD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSUBQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBSB ] --------------------------------------------------------------

func (d *disassembler) emitInstPSUBSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBSW ] --------------------------------------------------------------

func (d *disassembler) emitInstPSUBSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBUSB ] -------------------------------------------------------------

func (d *disassembler) emitInstPSUBUSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBUSW ] -------------------------------------------------------------

func (d *disassembler) emitInstPSUBUSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBW ] ---------------------------------------------------------------

func (d *disassembler) emitInstPSUBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PTEST ] ---------------------------------------------------------------

func (d *disassembler) emitInstPTEST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKHBW ] -----------------------------------------------------------

func (d *disassembler) emitInstPUNPCKHBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKHDQ ] -----------------------------------------------------------

func (d *disassembler) emitInstPUNPCKHDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKHQDQ ] ----------------------------------------------------------

func (d *disassembler) emitInstPUNPCKHQDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKHWD ] -----------------------------------------------------------

func (d *disassembler) emitInstPUNPCKHWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKLBW ] -----------------------------------------------------------

func (d *disassembler) emitInstPUNPCKLBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKLDQ ] -----------------------------------------------------------

func (d *disassembler) emitInstPUNPCKLDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKLQDQ ] ----------------------------------------------------------

func (d *disassembler) emitInstPUNPCKLQDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKLWD ] -----------------------------------------------------------

func (d *disassembler) emitInstPUNPCKLWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSH ] ----------------------------------------------------------------

func (d *disassembler) emitInstPUSH(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSHA ] ---------------------------------------------------------------

func (d *disassembler) emitInstPUSHA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSHAD ] --------------------------------------------------------------

func (d *disassembler) emitInstPUSHAD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSHF ] ---------------------------------------------------------------

func (d *disassembler) emitInstPUSHF(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSHFD ] --------------------------------------------------------------

func (d *disassembler) emitInstPUSHFD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSHFQ ] --------------------------------------------------------------

func (d *disassembler) emitInstPUSHFQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PXOR ] ----------------------------------------------------------------

func (d *disassembler) emitInstPXOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RCL ] -----------------------------------------------------------------

func (d *disassembler) emitInstRCL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RCPPS ] ---------------------------------------------------------------

func (d *disassembler) emitInstRCPPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RCPSS ] ---------------------------------------------------------------

func (d *disassembler) emitInstRCPSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RCR ] -----------------------------------------------------------------

func (d *disassembler) emitInstRCR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDFSBASE ] ------------------------------------------------------------

func (d *disassembler) emitInstRDFSBASE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDGSBASE ] ------------------------------------------------------------

func (d *disassembler) emitInstRDGSBASE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDMSR ] ---------------------------------------------------------------

func (d *disassembler) emitInstRDMSR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDPMC ] ---------------------------------------------------------------

func (d *disassembler) emitInstRDPMC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDRAND ] --------------------------------------------------------------

func (d *disassembler) emitInstRDRAND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDTSC ] ---------------------------------------------------------------

func (d *disassembler) emitInstRDTSC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDTSCP ] --------------------------------------------------------------

func (d *disassembler) emitInstRDTSCP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RET ] -----------------------------------------------------------------

func (d *disassembler) emitInstRET(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROL ] -----------------------------------------------------------------

func (d *disassembler) emitInstROL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROR ] -----------------------------------------------------------------

func (d *disassembler) emitInstROR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROUNDPD ] -------------------------------------------------------------

func (d *disassembler) emitInstROUNDPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROUNDPS ] -------------------------------------------------------------

func (d *disassembler) emitInstROUNDPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROUNDSD ] -------------------------------------------------------------

func (d *disassembler) emitInstROUNDSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROUNDSS ] -------------------------------------------------------------

func (d *disassembler) emitInstROUNDSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RSM ] -----------------------------------------------------------------

func (d *disassembler) emitInstRSM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RSQRTPS ] -------------------------------------------------------------

func (d *disassembler) emitInstRSQRTPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RSQRTSS ] -------------------------------------------------------------

func (d *disassembler) emitInstRSQRTSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SAHF ] ----------------------------------------------------------------

func (d *disassembler) emitInstSAHF(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SAR ] -----------------------------------------------------------------

func (d *disassembler) emitInstSAR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SBB ] -----------------------------------------------------------------

func (d *disassembler) emitInstSBB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SCASB ] ---------------------------------------------------------------

func (d *disassembler) emitInstSCASB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SCASD ] ---------------------------------------------------------------

func (d *disassembler) emitInstSCASD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SCASQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstSCASQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SCASW ] ---------------------------------------------------------------

func (d *disassembler) emitInstSCASW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETA ] ----------------------------------------------------------------

func (d *disassembler) emitInstSETA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETAE ] ---------------------------------------------------------------

func (d *disassembler) emitInstSETAE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETB ] ----------------------------------------------------------------

func (d *disassembler) emitInstSETB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETBE ] ---------------------------------------------------------------

func (d *disassembler) emitInstSETBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETE ] ----------------------------------------------------------------

func (d *disassembler) emitInstSETE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETG ] ----------------------------------------------------------------

func (d *disassembler) emitInstSETG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETGE ] ---------------------------------------------------------------

func (d *disassembler) emitInstSETGE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETL ] ----------------------------------------------------------------

func (d *disassembler) emitInstSETL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETLE ] ---------------------------------------------------------------

func (d *disassembler) emitInstSETLE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETNE ] ---------------------------------------------------------------

func (d *disassembler) emitInstSETNE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETNO ] ---------------------------------------------------------------

func (d *disassembler) emitInstSETNO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETNP ] ---------------------------------------------------------------

func (d *disassembler) emitInstSETNP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETNS ] ---------------------------------------------------------------

func (d *disassembler) emitInstSETNS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETO ] ----------------------------------------------------------------

func (d *disassembler) emitInstSETO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETP ] ----------------------------------------------------------------

func (d *disassembler) emitInstSETP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETS ] ----------------------------------------------------------------

func (d *disassembler) emitInstSETS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SFENCE ] --------------------------------------------------------------

func (d *disassembler) emitInstSFENCE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SGDT ] ----------------------------------------------------------------

func (d *disassembler) emitInstSGDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHL ] -----------------------------------------------------------------

func (d *disassembler) emitInstSHL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHLD ] ----------------------------------------------------------------

func (d *disassembler) emitInstSHLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHR ] -----------------------------------------------------------------

func (d *disassembler) emitInstSHR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHRD ] ----------------------------------------------------------------

func (d *disassembler) emitInstSHRD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHUFPD ] --------------------------------------------------------------

func (d *disassembler) emitInstSHUFPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHUFPS ] --------------------------------------------------------------

func (d *disassembler) emitInstSHUFPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SIDT ] ----------------------------------------------------------------

func (d *disassembler) emitInstSIDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SLDT ] ----------------------------------------------------------------

func (d *disassembler) emitInstSLDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SMSW ] ----------------------------------------------------------------

func (d *disassembler) emitInstSMSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SQRTPD ] --------------------------------------------------------------

func (d *disassembler) emitInstSQRTPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SQRTPS ] --------------------------------------------------------------

func (d *disassembler) emitInstSQRTPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SQRTSD ] --------------------------------------------------------------

func (d *disassembler) emitInstSQRTSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SQRTSS ] --------------------------------------------------------------

func (d *disassembler) emitInstSQRTSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STC ] -----------------------------------------------------------------

func (d *disassembler) emitInstSTC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STD ] -----------------------------------------------------------------

func (d *disassembler) emitInstSTD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STI ] -----------------------------------------------------------------

func (d *disassembler) emitInstSTI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STMXCSR ] -------------------------------------------------------------

func (d *disassembler) emitInstSTMXCSR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STOSB ] ---------------------------------------------------------------

func (d *disassembler) emitInstSTOSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STOSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstSTOSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STOSQ ] ---------------------------------------------------------------

func (d *disassembler) emitInstSTOSQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STOSW ] ---------------------------------------------------------------

func (d *disassembler) emitInstSTOSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STR ] -----------------------------------------------------------------

func (d *disassembler) emitInstSTR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SUB ] -----------------------------------------------------------------

func (d *disassembler) emitInstSUB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SUBPD ] ---------------------------------------------------------------

func (d *disassembler) emitInstSUBPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SUBPS ] ---------------------------------------------------------------

func (d *disassembler) emitInstSUBPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SUBSD ] ---------------------------------------------------------------

func (d *disassembler) emitInstSUBSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SUBSS ] ---------------------------------------------------------------

func (d *disassembler) emitInstSUBSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SWAPGS ] --------------------------------------------------------------

func (d *disassembler) emitInstSWAPGS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SYSCALL ] -------------------------------------------------------------

func (d *disassembler) emitInstSYSCALL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SYSENTER ] ------------------------------------------------------------

func (d *disassembler) emitInstSYSENTER(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SYSEXIT ] -------------------------------------------------------------

func (d *disassembler) emitInstSYSEXIT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SYSRET ] --------------------------------------------------------------

func (d *disassembler) emitInstSYSRET(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ TEST ] ----------------------------------------------------------------

func (d *disassembler) emitInstTEST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ TZCNT ] ---------------------------------------------------------------

func (d *disassembler) emitInstTZCNT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UCOMISD ] -------------------------------------------------------------

func (d *disassembler) emitInstUCOMISD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UCOMISS ] -------------------------------------------------------------

func (d *disassembler) emitInstUCOMISS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UD1 ] -----------------------------------------------------------------

func (d *disassembler) emitInstUD1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UD2 ] -----------------------------------------------------------------

func (d *disassembler) emitInstUD2(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UNPCKHPD ] ------------------------------------------------------------

func (d *disassembler) emitInstUNPCKHPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UNPCKHPS ] ------------------------------------------------------------

func (d *disassembler) emitInstUNPCKHPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UNPCKLPD ] ------------------------------------------------------------

func (d *disassembler) emitInstUNPCKLPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UNPCKLPS ] ------------------------------------------------------------

func (d *disassembler) emitInstUNPCKLPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VERR ] ----------------------------------------------------------------

func (d *disassembler) emitInstVERR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VERW ] ----------------------------------------------------------------

func (d *disassembler) emitInstVERW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VMOVDQA ] -------------------------------------------------------------

func (d *disassembler) emitInstVMOVDQA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VMOVDQU ] -------------------------------------------------------------

func (d *disassembler) emitInstVMOVDQU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VMOVNTDQ ] ------------------------------------------------------------

func (d *disassembler) emitInstVMOVNTDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VMOVNTDQA ] -----------------------------------------------------------

func (d *disassembler) emitInstVMOVNTDQA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VZEROUPPER ] ----------------------------------------------------------

func (d *disassembler) emitInstVZEROUPPER(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ WBINVD ] --------------------------------------------------------------

func (d *disassembler) emitInstWBINVD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ WRFSBASE ] ------------------------------------------------------------

func (d *disassembler) emitInstWRFSBASE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ WRGSBASE ] ------------------------------------------------------------

func (d *disassembler) emitInstWRGSBASE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ WRMSR ] ---------------------------------------------------------------

func (d *disassembler) emitInstWRMSR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XABORT ] --------------------------------------------------------------

func (d *disassembler) emitInstXABORT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XADD ] ----------------------------------------------------------------

func (d *disassembler) emitInstXADD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XBEGIN ] --------------------------------------------------------------

func (d *disassembler) emitInstXBEGIN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XCHG ] ----------------------------------------------------------------

func (d *disassembler) emitInstXCHG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XEND ] ----------------------------------------------------------------

func (d *disassembler) emitInstXEND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XGETBV ] --------------------------------------------------------------

func (d *disassembler) emitInstXGETBV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XLATB ] ---------------------------------------------------------------

func (d *disassembler) emitInstXLATB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XOR ] -----------------------------------------------------------------

func (d *disassembler) emitInstXOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XORPD ] ---------------------------------------------------------------

func (d *disassembler) emitInstXORPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XORPS ] ---------------------------------------------------------------

func (d *disassembler) emitInstXORPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XRSTOR ] --------------------------------------------------------------

func (d *disassembler) emitInstXRSTOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XRSTOR64 ] ------------------------------------------------------------

func (d *disassembler) emitInstXRSTOR64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XRSTORS ] -------------------------------------------------------------

func (d *disassembler) emitInstXRSTORS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XRSTORS64 ] -----------------------------------------------------------

func (d *disassembler) emitInstXRSTORS64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVE ] ---------------------------------------------------------------

func (d *disassembler) emitInstXSAVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVE64 ] -------------------------------------------------------------

func (d *disassembler) emitInstXSAVE64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVEC ] --------------------------------------------------------------

func (d *disassembler) emitInstXSAVEC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVEC64 ] ------------------------------------------------------------

func (d *disassembler) emitInstXSAVEC64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVEOPT ] ------------------------------------------------------------

func (d *disassembler) emitInstXSAVEOPT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVEOPT64 ] ----------------------------------------------------------

func (d *disassembler) emitInstXSAVEOPT64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVES ] --------------------------------------------------------------

func (d *disassembler) emitInstXSAVES(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVES64 ] ------------------------------------------------------------

func (d *disassembler) emitInstXSAVES64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSETBV ] --------------------------------------------------------------

func (d *disassembler) emitInstXSETBV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XTEST ] ---------------------------------------------------------------

func (d *disassembler) emitInstXTEST(f *function, inst *instruction) error {
	panic("not yet implemented")
}
