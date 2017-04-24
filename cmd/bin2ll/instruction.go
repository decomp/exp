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

// emitInst translates the given x86 AAA instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstAAA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AAD ] -----------------------------------------------------------------

// emitInst translates the given x86 AAD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstAAD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AAM ] -----------------------------------------------------------------

// emitInst translates the given x86 AAM instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstAAM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AAS ] -----------------------------------------------------------------

// emitInst translates the given x86 AAS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstAAS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADC ] -----------------------------------------------------------------

// emitInst translates the given x86 ADC instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstADC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADD ] -----------------------------------------------------------------

// emitInst translates the given x86 ADD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstADD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDPD ] ---------------------------------------------------------------

// emitInst translates the given x86 ADDPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstADDPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDPS ] ---------------------------------------------------------------

// emitInst translates the given x86 ADDPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstADDPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDSD ] ---------------------------------------------------------------

// emitInst translates the given x86 ADDSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstADDSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDSS ] ---------------------------------------------------------------

// emitInst translates the given x86 ADDSS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstADDSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDSUBPD ] ------------------------------------------------------------

// emitInst translates the given x86 ADDSUBPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstADDSUBPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ADDSUBPS ] ------------------------------------------------------------

// emitInst translates the given x86 ADDSUBPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstADDSUBPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESDEC ] --------------------------------------------------------------

// emitInst translates the given x86 AESDEC instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstAESDEC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESDECLAST ] ----------------------------------------------------------

// emitInst translates the given x86 AESDECLAST instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstAESDECLAST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESENC ] --------------------------------------------------------------

// emitInst translates the given x86 AESENC instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstAESENC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESENCLAST ] ----------------------------------------------------------

// emitInst translates the given x86 AESENCLAST instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstAESENCLAST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESIMC ] --------------------------------------------------------------

// emitInst translates the given x86 AESIMC instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstAESIMC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AESKEYGENASSIST ] -----------------------------------------------------

// emitInst translates the given x86 AESKEYGENASSIST instruction to LLVM IR,
// emitting code to f.
func (d *disassembler) emitInstAESKEYGENASSIST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ AND ] -----------------------------------------------------------------

// emitInst translates the given x86 AND instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstAND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ANDNPD ] --------------------------------------------------------------

// emitInst translates the given x86 ANDNPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstANDNPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ANDNPS ] --------------------------------------------------------------

// emitInst translates the given x86 ANDNPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstANDNPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ANDPD ] ---------------------------------------------------------------

// emitInst translates the given x86 ANDPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstANDPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ANDPS ] ---------------------------------------------------------------

// emitInst translates the given x86 ANDPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstANDPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ARPL ] ----------------------------------------------------------------

// emitInst translates the given x86 ARPL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstARPL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BLENDPD ] -------------------------------------------------------------

// emitInst translates the given x86 BLENDPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstBLENDPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BLENDPS ] -------------------------------------------------------------

// emitInst translates the given x86 BLENDPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstBLENDPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BLENDVPD ] ------------------------------------------------------------

// emitInst translates the given x86 BLENDVPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstBLENDVPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BLENDVPS ] ------------------------------------------------------------

// emitInst translates the given x86 BLENDVPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstBLENDVPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BOUND ] ---------------------------------------------------------------

// emitInst translates the given x86 BOUND instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstBOUND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BSF ] -----------------------------------------------------------------

// emitInst translates the given x86 BSF instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstBSF(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BSR ] -----------------------------------------------------------------

// emitInst translates the given x86 BSR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstBSR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BSWAP ] ---------------------------------------------------------------

// emitInst translates the given x86 BSWAP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstBSWAP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BT ] ------------------------------------------------------------------

// emitInst translates the given x86 BT instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstBT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BTC ] -----------------------------------------------------------------

// emitInst translates the given x86 BTC instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstBTC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BTR ] -----------------------------------------------------------------

// emitInst translates the given x86 BTR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstBTR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ BTS ] -----------------------------------------------------------------

// emitInst translates the given x86 BTS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstBTS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CALL ] ----------------------------------------------------------------

// emitInst translates the given x86 CALL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCALL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CBW ] -----------------------------------------------------------------

// emitInst translates the given x86 CBW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CDQ ] -----------------------------------------------------------------

// emitInst translates the given x86 CDQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CDQE ] ----------------------------------------------------------------

// emitInst translates the given x86 CDQE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCDQE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CLC ] -----------------------------------------------------------------

// emitInst translates the given x86 CLC instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCLC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CLD ] -----------------------------------------------------------------

// emitInst translates the given x86 CLD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CLFLUSH ] -------------------------------------------------------------

// emitInst translates the given x86 CLFLUSH instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCLFLUSH(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CLI ] -----------------------------------------------------------------

// emitInst translates the given x86 CLI instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCLI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CLTS ] ----------------------------------------------------------------

// emitInst translates the given x86 CLTS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCLTS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMC ] -----------------------------------------------------------------

// emitInst translates the given x86 CMC instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVA ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVA instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMOVA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVAE ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVAE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMOVAE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVB ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMOVB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVBE ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVBE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMOVBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVE ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMOVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVG ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVG instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMOVG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVGE ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVGE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMOVGE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVL ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMOVL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVLE ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVLE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMOVLE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVNE ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVNE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMOVNE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVNO ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVNO instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMOVNO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVNP ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVNP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMOVNP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVNS ] --------------------------------------------------------------

// emitInst translates the given x86 CMOVNS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMOVNS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVO ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVO instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMOVO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVP ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMOVP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMOVS ] ---------------------------------------------------------------

// emitInst translates the given x86 CMOVS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMOVS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMP ] -----------------------------------------------------------------

// emitInst translates the given x86 CMP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPPD ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMPPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPPS ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMPPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSB ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPSB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMPSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSD ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMPSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSD_XMM ] -----------------------------------------------------------

// emitInst translates the given x86 CMPSD_XMM instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMPSD_XMM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSQ ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPSQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMPSQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSS ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPSS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMPSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPSW ] ---------------------------------------------------------------

// emitInst translates the given x86 CMPSW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCMPSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPXCHG ] -------------------------------------------------------------

// emitInst translates the given x86 CMPXCHG instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMPXCHG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPXCHG16B ] ----------------------------------------------------------

// emitInst translates the given x86 CMPXCHG16B instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMPXCHG16B(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CMPXCHG8B ] -----------------------------------------------------------

// emitInst translates the given x86 CMPXCHG8B instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCMPXCHG8B(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ COMISD ] --------------------------------------------------------------

// emitInst translates the given x86 COMISD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCOMISD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ COMISS ] --------------------------------------------------------------

// emitInst translates the given x86 COMISS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCOMISS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CPUID ] ---------------------------------------------------------------

// emitInst translates the given x86 CPUID instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCPUID(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CQO ] -----------------------------------------------------------------

// emitInst translates the given x86 CQO instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCQO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CRC32 ] ---------------------------------------------------------------

// emitInst translates the given x86 CRC32 instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCRC32(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTDQ2PD ] ------------------------------------------------------------

// emitInst translates the given x86 CVTDQ2PD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTDQ2PD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTDQ2PS ] ------------------------------------------------------------

// emitInst translates the given x86 CVTDQ2PS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTDQ2PS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPD2DQ ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPD2DQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTPD2DQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPD2PI ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPD2PI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTPD2PI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPD2PS ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPD2PS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTPD2PS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPI2PD ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPI2PD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTPI2PD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPI2PS ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPI2PS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTPI2PS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPS2DQ ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPS2DQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTPS2DQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPS2PD ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPS2PD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTPS2PD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTPS2PI ] ------------------------------------------------------------

// emitInst translates the given x86 CVTPS2PI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTPS2PI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSD2SI ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSD2SI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTSD2SI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSD2SS ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSD2SS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTSD2SS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSI2SD ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSI2SD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTSI2SD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSI2SS ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSI2SS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTSI2SS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSS2SD ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSS2SD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTSS2SD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTSS2SI ] ------------------------------------------------------------

// emitInst translates the given x86 CVTSS2SI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTSS2SI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTPD2DQ ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTPD2DQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTTPD2DQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTPD2PI ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTPD2PI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTTPD2PI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTPS2DQ ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTPS2DQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTTPS2DQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTPS2PI ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTPS2PI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTTPS2PI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTSD2SI ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTSD2SI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTTSD2SI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CVTTSS2SI ] -----------------------------------------------------------

// emitInst translates the given x86 CVTTSS2SI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstCVTTSS2SI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CWD ] -----------------------------------------------------------------

// emitInst translates the given x86 CWD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ CWDE ] ----------------------------------------------------------------

// emitInst translates the given x86 CWDE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstCWDE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DAA ] -----------------------------------------------------------------

// emitInst translates the given x86 DAA instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstDAA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DAS ] -----------------------------------------------------------------

// emitInst translates the given x86 DAS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstDAS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DEC ] -----------------------------------------------------------------

// emitInst translates the given x86 DEC instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstDEC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DIV ] -----------------------------------------------------------------

// emitInst translates the given x86 DIV instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstDIV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DIVPD ] ---------------------------------------------------------------

// emitInst translates the given x86 DIVPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstDIVPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DIVPS ] ---------------------------------------------------------------

// emitInst translates the given x86 DIVPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstDIVPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DIVSD ] ---------------------------------------------------------------

// emitInst translates the given x86 DIVSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstDIVSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DIVSS ] ---------------------------------------------------------------

// emitInst translates the given x86 DIVSS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstDIVSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DPPD ] ----------------------------------------------------------------

// emitInst translates the given x86 DPPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstDPPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ DPPS ] ----------------------------------------------------------------

// emitInst translates the given x86 DPPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstDPPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ EMMS ] ----------------------------------------------------------------

// emitInst translates the given x86 EMMS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstEMMS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ENTER ] ---------------------------------------------------------------

// emitInst translates the given x86 ENTER instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstENTER(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ EXTRACTPS ] -----------------------------------------------------------

// emitInst translates the given x86 EXTRACTPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstEXTRACTPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ F2XM1 ] ---------------------------------------------------------------

// emitInst translates the given x86 F2XM1 instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstF2XM1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FABS ] ----------------------------------------------------------------

// emitInst translates the given x86 FABS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFABS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FADD ] ----------------------------------------------------------------

// emitInst translates the given x86 FADD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFADD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FADDP ] ---------------------------------------------------------------

// emitInst translates the given x86 FADDP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFADDP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FBLD ] ----------------------------------------------------------------

// emitInst translates the given x86 FBLD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFBLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FBSTP ] ---------------------------------------------------------------

// emitInst translates the given x86 FBSTP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFBSTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCHS ] ----------------------------------------------------------------

// emitInst translates the given x86 FCHS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFCHS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVB ] --------------------------------------------------------------

// emitInst translates the given x86 FCMOVB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFCMOVB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVBE ] -------------------------------------------------------------

// emitInst translates the given x86 FCMOVBE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFCMOVBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVE ] --------------------------------------------------------------

// emitInst translates the given x86 FCMOVE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFCMOVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVNB ] -------------------------------------------------------------

// emitInst translates the given x86 FCMOVNB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFCMOVNB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVNBE ] ------------------------------------------------------------

// emitInst translates the given x86 FCMOVNBE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFCMOVNBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVNE ] -------------------------------------------------------------

// emitInst translates the given x86 FCMOVNE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFCMOVNE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVNU ] -------------------------------------------------------------

// emitInst translates the given x86 FCMOVNU instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFCMOVNU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCMOVU ] --------------------------------------------------------------

// emitInst translates the given x86 FCMOVU instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFCMOVU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOM ] ----------------------------------------------------------------

// emitInst translates the given x86 FCOM instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFCOM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOMI ] ---------------------------------------------------------------

// emitInst translates the given x86 FCOMI instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFCOMI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOMIP ] --------------------------------------------------------------

// emitInst translates the given x86 FCOMIP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFCOMIP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOMP ] ---------------------------------------------------------------

// emitInst translates the given x86 FCOMP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFCOMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOMPP ] --------------------------------------------------------------

// emitInst translates the given x86 FCOMPP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFCOMPP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FCOS ] ----------------------------------------------------------------

// emitInst translates the given x86 FCOS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFCOS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FDECSTP ] -------------------------------------------------------------

// emitInst translates the given x86 FDECSTP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFDECSTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FDIV ] ----------------------------------------------------------------

// emitInst translates the given x86 FDIV instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFDIV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FDIVP ] ---------------------------------------------------------------

// emitInst translates the given x86 FDIVP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFDIVP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FDIVR ] ---------------------------------------------------------------

// emitInst translates the given x86 FDIVR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFDIVR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FDIVRP ] --------------------------------------------------------------

// emitInst translates the given x86 FDIVRP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFDIVRP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FFREE ] ---------------------------------------------------------------

// emitInst translates the given x86 FFREE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFFREE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FFREEP ] --------------------------------------------------------------

// emitInst translates the given x86 FFREEP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFFREEP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FIADD ] ---------------------------------------------------------------

// emitInst translates the given x86 FIADD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFIADD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FICOM ] ---------------------------------------------------------------

// emitInst translates the given x86 FICOM instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFICOM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FICOMP ] --------------------------------------------------------------

// emitInst translates the given x86 FICOMP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFICOMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FIDIV ] ---------------------------------------------------------------

// emitInst translates the given x86 FIDIV instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFIDIV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FIDIVR ] --------------------------------------------------------------

// emitInst translates the given x86 FIDIVR instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFIDIVR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FILD ] ----------------------------------------------------------------

// emitInst translates the given x86 FILD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFILD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FIMUL ] ---------------------------------------------------------------

// emitInst translates the given x86 FIMUL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFIMUL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FINCSTP ] -------------------------------------------------------------

// emitInst translates the given x86 FINCSTP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFINCSTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FIST ] ----------------------------------------------------------------

// emitInst translates the given x86 FIST instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFIST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FISTP ] ---------------------------------------------------------------

// emitInst translates the given x86 FISTP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFISTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FISTTP ] --------------------------------------------------------------

// emitInst translates the given x86 FISTTP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFISTTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FISUB ] ---------------------------------------------------------------

// emitInst translates the given x86 FISUB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFISUB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FISUBR ] --------------------------------------------------------------

// emitInst translates the given x86 FISUBR instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFISUBR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLD ] -----------------------------------------------------------------

// emitInst translates the given x86 FLD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLD1 ] ----------------------------------------------------------------

// emitInst translates the given x86 FLD1 instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFLD1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDCW ] ---------------------------------------------------------------

// emitInst translates the given x86 FLDCW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFLDCW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDENV ] --------------------------------------------------------------

// emitInst translates the given x86 FLDENV instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFLDENV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDL2E ] --------------------------------------------------------------

// emitInst translates the given x86 FLDL2E instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFLDL2E(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDL2T ] --------------------------------------------------------------

// emitInst translates the given x86 FLDL2T instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFLDL2T(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDLG2 ] --------------------------------------------------------------

// emitInst translates the given x86 FLDLG2 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFLDLG2(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDLN2 ] --------------------------------------------------------------

// emitInst translates the given x86 FLDLN2 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFLDLN2(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDPI ] ---------------------------------------------------------------

// emitInst translates the given x86 FLDPI instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFLDPI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FLDZ ] ----------------------------------------------------------------

// emitInst translates the given x86 FLDZ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFLDZ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FMUL ] ----------------------------------------------------------------

// emitInst translates the given x86 FMUL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFMUL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FMULP ] ---------------------------------------------------------------

// emitInst translates the given x86 FMULP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFMULP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNCLEX ] --------------------------------------------------------------

// emitInst translates the given x86 FNCLEX instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFNCLEX(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNINIT ] --------------------------------------------------------------

// emitInst translates the given x86 FNINIT instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFNINIT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNOP ] ----------------------------------------------------------------

// emitInst translates the given x86 FNOP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFNOP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNSAVE ] --------------------------------------------------------------

// emitInst translates the given x86 FNSAVE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFNSAVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNSTCW ] --------------------------------------------------------------

// emitInst translates the given x86 FNSTCW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFNSTCW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNSTENV ] -------------------------------------------------------------

// emitInst translates the given x86 FNSTENV instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFNSTENV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FNSTSW ] --------------------------------------------------------------

// emitInst translates the given x86 FNSTSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFNSTSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FPATAN ] --------------------------------------------------------------

// emitInst translates the given x86 FPATAN instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFPATAN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FPREM ] ---------------------------------------------------------------

// emitInst translates the given x86 FPREM instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFPREM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FPREM1 ] --------------------------------------------------------------

// emitInst translates the given x86 FPREM1 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFPREM1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FPTAN ] ---------------------------------------------------------------

// emitInst translates the given x86 FPTAN instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFPTAN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FRNDINT ] -------------------------------------------------------------

// emitInst translates the given x86 FRNDINT instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFRNDINT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FRSTOR ] --------------------------------------------------------------

// emitInst translates the given x86 FRSTOR instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFRSTOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSCALE ] --------------------------------------------------------------

// emitInst translates the given x86 FSCALE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFSCALE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSIN ] ----------------------------------------------------------------

// emitInst translates the given x86 FSIN instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFSIN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSINCOS ] -------------------------------------------------------------

// emitInst translates the given x86 FSINCOS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFSINCOS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSQRT ] ---------------------------------------------------------------

// emitInst translates the given x86 FSQRT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFSQRT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FST ] -----------------------------------------------------------------

// emitInst translates the given x86 FST instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSTP ] ----------------------------------------------------------------

// emitInst translates the given x86 FSTP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFSTP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSUB ] ----------------------------------------------------------------

// emitInst translates the given x86 FSUB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFSUB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSUBP ] ---------------------------------------------------------------

// emitInst translates the given x86 FSUBP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFSUBP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSUBR ] ---------------------------------------------------------------

// emitInst translates the given x86 FSUBR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFSUBR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FSUBRP ] --------------------------------------------------------------

// emitInst translates the given x86 FSUBRP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFSUBRP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FTST ] ----------------------------------------------------------------

// emitInst translates the given x86 FTST instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFTST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FUCOM ] ---------------------------------------------------------------

// emitInst translates the given x86 FUCOM instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFUCOM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FUCOMI ] --------------------------------------------------------------

// emitInst translates the given x86 FUCOMI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFUCOMI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FUCOMIP ] -------------------------------------------------------------

// emitInst translates the given x86 FUCOMIP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFUCOMIP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FUCOMP ] --------------------------------------------------------------

// emitInst translates the given x86 FUCOMP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFUCOMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FUCOMPP ] -------------------------------------------------------------

// emitInst translates the given x86 FUCOMPP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFUCOMPP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FWAIT ] ---------------------------------------------------------------

// emitInst translates the given x86 FWAIT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFWAIT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXAM ] ----------------------------------------------------------------

// emitInst translates the given x86 FXAM instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFXAM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXCH ] ----------------------------------------------------------------

// emitInst translates the given x86 FXCH instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFXCH(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXRSTOR ] -------------------------------------------------------------

// emitInst translates the given x86 FXRSTOR instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFXRSTOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXRSTOR64 ] -----------------------------------------------------------

// emitInst translates the given x86 FXRSTOR64 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFXRSTOR64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXSAVE ] --------------------------------------------------------------

// emitInst translates the given x86 FXSAVE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFXSAVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXSAVE64 ] ------------------------------------------------------------

// emitInst translates the given x86 FXSAVE64 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFXSAVE64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FXTRACT ] -------------------------------------------------------------

// emitInst translates the given x86 FXTRACT instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFXTRACT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FYL2X ] ---------------------------------------------------------------

// emitInst translates the given x86 FYL2X instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstFYL2X(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ FYL2XP1 ] -------------------------------------------------------------

// emitInst translates the given x86 FYL2XP1 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstFYL2XP1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ HADDPD ] --------------------------------------------------------------

// emitInst translates the given x86 HADDPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstHADDPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ HADDPS ] --------------------------------------------------------------

// emitInst translates the given x86 HADDPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstHADDPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ HLT ] -----------------------------------------------------------------

// emitInst translates the given x86 HLT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstHLT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ HSUBPD ] --------------------------------------------------------------

// emitInst translates the given x86 HSUBPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstHSUBPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ HSUBPS ] --------------------------------------------------------------

// emitInst translates the given x86 HSUBPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstHSUBPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ICEBP ] ---------------------------------------------------------------

// emitInst translates the given x86 ICEBP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstICEBP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IDIV ] ----------------------------------------------------------------

// emitInst translates the given x86 IDIV instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstIDIV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IMUL ] ----------------------------------------------------------------

// emitInst translates the given x86 IMUL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstIMUL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IN ] ------------------------------------------------------------------

// emitInst translates the given x86 IN instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstIN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INC ] -----------------------------------------------------------------

// emitInst translates the given x86 INC instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstINC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INSB ] ----------------------------------------------------------------

// emitInst translates the given x86 INSB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstINSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INSD ] ----------------------------------------------------------------

// emitInst translates the given x86 INSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstINSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INSERTPS ] ------------------------------------------------------------

// emitInst translates the given x86 INSERTPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstINSERTPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INSW ] ----------------------------------------------------------------

// emitInst translates the given x86 INSW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstINSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INT ] -----------------------------------------------------------------

// emitInst translates the given x86 INT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstINT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INTO ] ----------------------------------------------------------------

// emitInst translates the given x86 INTO instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstINTO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INVD ] ----------------------------------------------------------------

// emitInst translates the given x86 INVD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstINVD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INVLPG ] --------------------------------------------------------------

// emitInst translates the given x86 INVLPG instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstINVLPG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ INVPCID ] -------------------------------------------------------------

// emitInst translates the given x86 INVPCID instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstINVPCID(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IRET ] ----------------------------------------------------------------

// emitInst translates the given x86 IRET instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstIRET(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IRETD ] ---------------------------------------------------------------

// emitInst translates the given x86 IRETD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstIRETD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ IRETQ ] ---------------------------------------------------------------

// emitInst translates the given x86 IRETQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstIRETQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JA ] ------------------------------------------------------------------

// emitInst translates the given x86 JA instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstJA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JAE ] -----------------------------------------------------------------

// emitInst translates the given x86 JAE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJAE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JB ] ------------------------------------------------------------------

// emitInst translates the given x86 JB instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstJB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JBE ] -----------------------------------------------------------------

// emitInst translates the given x86 JBE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JCXZ ] ----------------------------------------------------------------

// emitInst translates the given x86 JCXZ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJCXZ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JE ] ------------------------------------------------------------------

// emitInst translates the given x86 JE instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstJE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JECXZ ] ---------------------------------------------------------------

// emitInst translates the given x86 JECXZ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJECXZ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JG ] ------------------------------------------------------------------

// emitInst translates the given x86 JG instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstJG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JGE ] -----------------------------------------------------------------

// emitInst translates the given x86 JGE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJGE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JL ] ------------------------------------------------------------------

// emitInst translates the given x86 JL instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstJL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JLE ] -----------------------------------------------------------------

// emitInst translates the given x86 JLE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJLE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JMP ] -----------------------------------------------------------------

// emitInst translates the given x86 JMP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JNE ] -----------------------------------------------------------------

// emitInst translates the given x86 JNE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJNE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JNO ] -----------------------------------------------------------------

// emitInst translates the given x86 JNO instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJNO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JNP ] -----------------------------------------------------------------

// emitInst translates the given x86 JNP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJNP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JNS ] -----------------------------------------------------------------

// emitInst translates the given x86 JNS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJNS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JO ] ------------------------------------------------------------------

// emitInst translates the given x86 JO instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstJO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JP ] ------------------------------------------------------------------

// emitInst translates the given x86 JP instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstJP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JRCXZ ] ---------------------------------------------------------------

// emitInst translates the given x86 JRCXZ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstJRCXZ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ JS ] ------------------------------------------------------------------

// emitInst translates the given x86 JS instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstJS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LAHF ] ----------------------------------------------------------------

// emitInst translates the given x86 LAHF instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLAHF(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LAR ] -----------------------------------------------------------------

// emitInst translates the given x86 LAR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLAR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LCALL ] ---------------------------------------------------------------

// emitInst translates the given x86 LCALL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLCALL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LDDQU ] ---------------------------------------------------------------

// emitInst translates the given x86 LDDQU instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLDDQU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LDMXCSR ] -------------------------------------------------------------

// emitInst translates the given x86 LDMXCSR instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstLDMXCSR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LDS ] -----------------------------------------------------------------

// emitInst translates the given x86 LDS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLDS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LEA ] -----------------------------------------------------------------

// emitInst translates the given x86 LEA instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLEA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LEAVE ] ---------------------------------------------------------------

// emitInst translates the given x86 LEAVE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLEAVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LES ] -----------------------------------------------------------------

// emitInst translates the given x86 LES instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLES(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LFENCE ] --------------------------------------------------------------

// emitInst translates the given x86 LFENCE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstLFENCE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LFS ] -----------------------------------------------------------------

// emitInst translates the given x86 LFS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLFS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LGDT ] ----------------------------------------------------------------

// emitInst translates the given x86 LGDT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLGDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LGS ] -----------------------------------------------------------------

// emitInst translates the given x86 LGS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLGS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LIDT ] ----------------------------------------------------------------

// emitInst translates the given x86 LIDT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLIDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LJMP ] ----------------------------------------------------------------

// emitInst translates the given x86 LJMP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLJMP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LLDT ] ----------------------------------------------------------------

// emitInst translates the given x86 LLDT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLLDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LMSW ] ----------------------------------------------------------------

// emitInst translates the given x86 LMSW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLMSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LODSB ] ---------------------------------------------------------------

// emitInst translates the given x86 LODSB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLODSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LODSD ] ---------------------------------------------------------------

// emitInst translates the given x86 LODSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLODSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LODSQ ] ---------------------------------------------------------------

// emitInst translates the given x86 LODSQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLODSQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LODSW ] ---------------------------------------------------------------

// emitInst translates the given x86 LODSW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLODSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LOOP ] ----------------------------------------------------------------

// emitInst translates the given x86 LOOP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLOOP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LOOPE ] ---------------------------------------------------------------

// emitInst translates the given x86 LOOPE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLOOPE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LOOPNE ] --------------------------------------------------------------

// emitInst translates the given x86 LOOPNE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstLOOPNE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LRET ] ----------------------------------------------------------------

// emitInst translates the given x86 LRET instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLRET(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LSL ] -----------------------------------------------------------------

// emitInst translates the given x86 LSL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLSL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LSS ] -----------------------------------------------------------------

// emitInst translates the given x86 LSS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LTR ] -----------------------------------------------------------------

// emitInst translates the given x86 LTR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLTR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ LZCNT ] ---------------------------------------------------------------

// emitInst translates the given x86 LZCNT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstLZCNT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MASKMOVDQU ] ----------------------------------------------------------

// emitInst translates the given x86 MASKMOVDQU instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMASKMOVDQU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MASKMOVQ ] ------------------------------------------------------------

// emitInst translates the given x86 MASKMOVQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMASKMOVQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MAXPD ] ---------------------------------------------------------------

// emitInst translates the given x86 MAXPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMAXPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MAXPS ] ---------------------------------------------------------------

// emitInst translates the given x86 MAXPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMAXPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MAXSD ] ---------------------------------------------------------------

// emitInst translates the given x86 MAXSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMAXSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MAXSS ] ---------------------------------------------------------------

// emitInst translates the given x86 MAXSS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMAXSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MFENCE ] --------------------------------------------------------------

// emitInst translates the given x86 MFENCE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMFENCE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MINPD ] ---------------------------------------------------------------

// emitInst translates the given x86 MINPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMINPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MINPS ] ---------------------------------------------------------------

// emitInst translates the given x86 MINPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMINPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MINSD ] ---------------------------------------------------------------

// emitInst translates the given x86 MINSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMINSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MINSS ] ---------------------------------------------------------------

// emitInst translates the given x86 MINSS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMINSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MONITOR ] -------------------------------------------------------------

// emitInst translates the given x86 MONITOR instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMONITOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOV ] -----------------------------------------------------------------

// emitInst translates the given x86 MOV instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVAPD ] --------------------------------------------------------------

// emitInst translates the given x86 MOVAPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVAPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVAPS ] --------------------------------------------------------------

// emitInst translates the given x86 MOVAPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVAPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVBE ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVBE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOVBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVD ] ----------------------------------------------------------------

// emitInst translates the given x86 MOVD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOVD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVDDUP ] -------------------------------------------------------------

// emitInst translates the given x86 MOVDDUP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVDDUP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVDQ2Q ] -------------------------------------------------------------

// emitInst translates the given x86 MOVDQ2Q instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVDQ2Q(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVDQA ] --------------------------------------------------------------

// emitInst translates the given x86 MOVDQA instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVDQA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVDQU ] --------------------------------------------------------------

// emitInst translates the given x86 MOVDQU instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVDQU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVHLPS ] -------------------------------------------------------------

// emitInst translates the given x86 MOVHLPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVHLPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVHPD ] --------------------------------------------------------------

// emitInst translates the given x86 MOVHPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVHPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVHPS ] --------------------------------------------------------------

// emitInst translates the given x86 MOVHPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVHPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVLHPS ] -------------------------------------------------------------

// emitInst translates the given x86 MOVLHPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVLHPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVLPD ] --------------------------------------------------------------

// emitInst translates the given x86 MOVLPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVLPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVLPS ] --------------------------------------------------------------

// emitInst translates the given x86 MOVLPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVLPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVMSKPD ] ------------------------------------------------------------

// emitInst translates the given x86 MOVMSKPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVMSKPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVMSKPS ] ------------------------------------------------------------

// emitInst translates the given x86 MOVMSKPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVMSKPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTDQ ] -------------------------------------------------------------

// emitInst translates the given x86 MOVNTDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVNTDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTDQA ] ------------------------------------------------------------

// emitInst translates the given x86 MOVNTDQA instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVNTDQA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTI ] --------------------------------------------------------------

// emitInst translates the given x86 MOVNTI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVNTI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTPD ] -------------------------------------------------------------

// emitInst translates the given x86 MOVNTPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVNTPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTPS ] -------------------------------------------------------------

// emitInst translates the given x86 MOVNTPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVNTPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTQ ] --------------------------------------------------------------

// emitInst translates the given x86 MOVNTQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVNTQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTSD ] -------------------------------------------------------------

// emitInst translates the given x86 MOVNTSD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVNTSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVNTSS ] -------------------------------------------------------------

// emitInst translates the given x86 MOVNTSS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVNTSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVQ ] ----------------------------------------------------------------

// emitInst translates the given x86 MOVQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOVQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVQ2DQ ] -------------------------------------------------------------

// emitInst translates the given x86 MOVQ2DQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVQ2DQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSB ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOVSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSD ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOVSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSD_XMM ] -----------------------------------------------------------

// emitInst translates the given x86 MOVSD_XMM instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVSD_XMM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSHDUP ] ------------------------------------------------------------

// emitInst translates the given x86 MOVSHDUP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVSHDUP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSLDUP ] ------------------------------------------------------------

// emitInst translates the given x86 MOVSLDUP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVSLDUP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSQ ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOVSQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSS ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOVSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSW ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOVSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSX ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVSX instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOVSX(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVSXD ] --------------------------------------------------------------

// emitInst translates the given x86 MOVSXD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVSXD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVUPD ] --------------------------------------------------------------

// emitInst translates the given x86 MOVUPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVUPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVUPS ] --------------------------------------------------------------

// emitInst translates the given x86 MOVUPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMOVUPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MOVZX ] ---------------------------------------------------------------

// emitInst translates the given x86 MOVZX instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMOVZX(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MPSADBW ] -------------------------------------------------------------

// emitInst translates the given x86 MPSADBW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstMPSADBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MUL ] -----------------------------------------------------------------

// emitInst translates the given x86 MUL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMUL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MULPD ] ---------------------------------------------------------------

// emitInst translates the given x86 MULPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMULPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MULPS ] ---------------------------------------------------------------

// emitInst translates the given x86 MULPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMULPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MULSD ] ---------------------------------------------------------------

// emitInst translates the given x86 MULSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMULSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MULSS ] ---------------------------------------------------------------

// emitInst translates the given x86 MULSS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMULSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ MWAIT ] ---------------------------------------------------------------

// emitInst translates the given x86 MWAIT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstMWAIT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ NEG ] -----------------------------------------------------------------

// emitInst translates the given x86 NEG instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstNEG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ NOP ] -----------------------------------------------------------------

// emitInst translates the given x86 NOP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstNOP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ NOT ] -----------------------------------------------------------------

// emitInst translates the given x86 NOT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstNOT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ OR ] ------------------------------------------------------------------

// emitInst translates the given x86 OR instruction to LLVM IR, emitting code to
// f.
func (d *disassembler) emitInstOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ORPD ] ----------------------------------------------------------------

// emitInst translates the given x86 ORPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstORPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ORPS ] ----------------------------------------------------------------

// emitInst translates the given x86 ORPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstORPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ OUT ] -----------------------------------------------------------------

// emitInst translates the given x86 OUT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstOUT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ OUTSB ] ---------------------------------------------------------------

// emitInst translates the given x86 OUTSB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstOUTSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ OUTSD ] ---------------------------------------------------------------

// emitInst translates the given x86 OUTSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstOUTSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ OUTSW ] ---------------------------------------------------------------

// emitInst translates the given x86 OUTSW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstOUTSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PABSB ] ---------------------------------------------------------------

// emitInst translates the given x86 PABSB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPABSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PABSD ] ---------------------------------------------------------------

// emitInst translates the given x86 PABSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPABSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PABSW ] ---------------------------------------------------------------

// emitInst translates the given x86 PABSW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPABSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PACKSSDW ] ------------------------------------------------------------

// emitInst translates the given x86 PACKSSDW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPACKSSDW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PACKSSWB ] ------------------------------------------------------------

// emitInst translates the given x86 PACKSSWB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPACKSSWB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PACKUSDW ] ------------------------------------------------------------

// emitInst translates the given x86 PACKUSDW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPACKUSDW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PACKUSWB ] ------------------------------------------------------------

// emitInst translates the given x86 PACKUSWB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPACKUSWB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDB ] ---------------------------------------------------------------

// emitInst translates the given x86 PADDB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPADDB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDD ] ---------------------------------------------------------------

// emitInst translates the given x86 PADDD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPADDD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDQ ] ---------------------------------------------------------------

// emitInst translates the given x86 PADDQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPADDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDSB ] --------------------------------------------------------------

// emitInst translates the given x86 PADDSB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPADDSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDSW ] --------------------------------------------------------------

// emitInst translates the given x86 PADDSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPADDSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDUSB ] -------------------------------------------------------------

// emitInst translates the given x86 PADDUSB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPADDUSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDUSW ] -------------------------------------------------------------

// emitInst translates the given x86 PADDUSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPADDUSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PADDW ] ---------------------------------------------------------------

// emitInst translates the given x86 PADDW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPADDW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PALIGNR ] -------------------------------------------------------------

// emitInst translates the given x86 PALIGNR instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPALIGNR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PAND ] ----------------------------------------------------------------

// emitInst translates the given x86 PAND instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPAND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PANDN ] ---------------------------------------------------------------

// emitInst translates the given x86 PANDN instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPANDN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PAUSE ] ---------------------------------------------------------------

// emitInst translates the given x86 PAUSE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPAUSE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PAVGB ] ---------------------------------------------------------------

// emitInst translates the given x86 PAVGB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPAVGB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PAVGW ] ---------------------------------------------------------------

// emitInst translates the given x86 PAVGW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPAVGW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PBLENDVB ] ------------------------------------------------------------

// emitInst translates the given x86 PBLENDVB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPBLENDVB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PBLENDW ] -------------------------------------------------------------

// emitInst translates the given x86 PBLENDW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPBLENDW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCLMULQDQ ] -----------------------------------------------------------

// emitInst translates the given x86 PCLMULQDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCLMULQDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPEQB ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPEQB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPEQB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPEQD ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPEQD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPEQD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPEQQ ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPEQQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPEQQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPEQW ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPEQW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPEQW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPESTRI ] -----------------------------------------------------------

// emitInst translates the given x86 PCMPESTRI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPESTRI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPESTRM ] -----------------------------------------------------------

// emitInst translates the given x86 PCMPESTRM instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPESTRM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPGTB ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPGTB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPGTB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPGTD ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPGTD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPGTD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPGTQ ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPGTQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPGTQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPGTW ] -------------------------------------------------------------

// emitInst translates the given x86 PCMPGTW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPGTW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPISTRI ] -----------------------------------------------------------

// emitInst translates the given x86 PCMPISTRI instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPISTRI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PCMPISTRM ] -----------------------------------------------------------

// emitInst translates the given x86 PCMPISTRM instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPCMPISTRM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PEXTRB ] --------------------------------------------------------------

// emitInst translates the given x86 PEXTRB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPEXTRB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PEXTRD ] --------------------------------------------------------------

// emitInst translates the given x86 PEXTRD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPEXTRD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PEXTRQ ] --------------------------------------------------------------

// emitInst translates the given x86 PEXTRQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPEXTRQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PEXTRW ] --------------------------------------------------------------

// emitInst translates the given x86 PEXTRW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPEXTRW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHADDD ] --------------------------------------------------------------

// emitInst translates the given x86 PHADDD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPHADDD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHADDSW ] -------------------------------------------------------------

// emitInst translates the given x86 PHADDSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPHADDSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHADDW ] --------------------------------------------------------------

// emitInst translates the given x86 PHADDW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPHADDW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHMINPOSUW ] ----------------------------------------------------------

// emitInst translates the given x86 PHMINPOSUW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPHMINPOSUW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHSUBD ] --------------------------------------------------------------

// emitInst translates the given x86 PHSUBD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPHSUBD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHSUBSW ] -------------------------------------------------------------

// emitInst translates the given x86 PHSUBSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPHSUBSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PHSUBW ] --------------------------------------------------------------

// emitInst translates the given x86 PHSUBW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPHSUBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PINSRB ] --------------------------------------------------------------

// emitInst translates the given x86 PINSRB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPINSRB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PINSRD ] --------------------------------------------------------------

// emitInst translates the given x86 PINSRD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPINSRD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PINSRQ ] --------------------------------------------------------------

// emitInst translates the given x86 PINSRQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPINSRQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PINSRW ] --------------------------------------------------------------

// emitInst translates the given x86 PINSRW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPINSRW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMADDUBSW ] -----------------------------------------------------------

// emitInst translates the given x86 PMADDUBSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMADDUBSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMADDWD ] -------------------------------------------------------------

// emitInst translates the given x86 PMADDWD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMADDWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXSB ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXSB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMAXSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXSD ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXSD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMAXSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXSW ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMAXSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXUB ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXUB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMAXUB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXUD ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXUD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMAXUD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMAXUW ] --------------------------------------------------------------

// emitInst translates the given x86 PMAXUW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMAXUW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINSB ] --------------------------------------------------------------

// emitInst translates the given x86 PMINSB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMINSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINSD ] --------------------------------------------------------------

// emitInst translates the given x86 PMINSD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMINSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINSW ] --------------------------------------------------------------

// emitInst translates the given x86 PMINSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMINSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINUB ] --------------------------------------------------------------

// emitInst translates the given x86 PMINUB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMINUB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINUD ] --------------------------------------------------------------

// emitInst translates the given x86 PMINUD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMINUD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMINUW ] --------------------------------------------------------------

// emitInst translates the given x86 PMINUW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMINUW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVMSKB ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVMSKB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVMSKB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXBD ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXBD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVSXBD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXBQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXBQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVSXBQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXBW ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXBW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVSXBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXDQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVSXDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXWD ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXWD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVSXWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVSXWQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVSXWQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVSXWQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXBD ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXBD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVZXBD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXBQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXBQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVZXBQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXBW ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXBW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVZXBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXDQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVZXDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXWD ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXWD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVZXWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMOVZXWQ ] ------------------------------------------------------------

// emitInst translates the given x86 PMOVZXWQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMOVZXWQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULDQ ] --------------------------------------------------------------

// emitInst translates the given x86 PMULDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMULDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULHRSW ] ------------------------------------------------------------

// emitInst translates the given x86 PMULHRSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMULHRSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULHUW ] -------------------------------------------------------------

// emitInst translates the given x86 PMULHUW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMULHUW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULHW ] --------------------------------------------------------------

// emitInst translates the given x86 PMULHW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMULHW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULLD ] --------------------------------------------------------------

// emitInst translates the given x86 PMULLD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMULLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULLW ] --------------------------------------------------------------

// emitInst translates the given x86 PMULLW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMULLW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PMULUDQ ] -------------------------------------------------------------

// emitInst translates the given x86 PMULUDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPMULUDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POP ] -----------------------------------------------------------------

// emitInst translates the given x86 POP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPOP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPA ] ----------------------------------------------------------------

// emitInst translates the given x86 POPA instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPOPA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPAD ] ---------------------------------------------------------------

// emitInst translates the given x86 POPAD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPOPAD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPCNT ] --------------------------------------------------------------

// emitInst translates the given x86 POPCNT instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPOPCNT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPF ] ----------------------------------------------------------------

// emitInst translates the given x86 POPF instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPOPF(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPFD ] ---------------------------------------------------------------

// emitInst translates the given x86 POPFD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPOPFD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POPFQ ] ---------------------------------------------------------------

// emitInst translates the given x86 POPFQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPOPFQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ POR ] -----------------------------------------------------------------

// emitInst translates the given x86 POR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PREFETCHNTA ] ---------------------------------------------------------

// emitInst translates the given x86 PREFETCHNTA instruction to LLVM IR,
// emitting code to f.
func (d *disassembler) emitInstPREFETCHNTA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PREFETCHT0 ] ----------------------------------------------------------

// emitInst translates the given x86 PREFETCHT0 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPREFETCHT0(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PREFETCHT1 ] ----------------------------------------------------------

// emitInst translates the given x86 PREFETCHT1 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPREFETCHT1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PREFETCHT2 ] ----------------------------------------------------------

// emitInst translates the given x86 PREFETCHT2 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPREFETCHT2(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PREFETCHW ] -----------------------------------------------------------

// emitInst translates the given x86 PREFETCHW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPREFETCHW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSADBW ] --------------------------------------------------------------

// emitInst translates the given x86 PSADBW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSADBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSHUFB ] --------------------------------------------------------------

// emitInst translates the given x86 PSHUFB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSHUFB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSHUFD ] --------------------------------------------------------------

// emitInst translates the given x86 PSHUFD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSHUFD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSHUFHW ] -------------------------------------------------------------

// emitInst translates the given x86 PSHUFHW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSHUFHW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSHUFLW ] -------------------------------------------------------------

// emitInst translates the given x86 PSHUFLW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSHUFLW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSHUFW ] --------------------------------------------------------------

// emitInst translates the given x86 PSHUFW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSHUFW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSIGNB ] --------------------------------------------------------------

// emitInst translates the given x86 PSIGNB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSIGNB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSIGND ] --------------------------------------------------------------

// emitInst translates the given x86 PSIGND instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSIGND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSIGNW ] --------------------------------------------------------------

// emitInst translates the given x86 PSIGNW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSIGNW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSLLD ] ---------------------------------------------------------------

// emitInst translates the given x86 PSLLD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSLLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSLLDQ ] --------------------------------------------------------------

// emitInst translates the given x86 PSLLDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSLLDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSLLQ ] ---------------------------------------------------------------

// emitInst translates the given x86 PSLLQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSLLQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSLLW ] ---------------------------------------------------------------

// emitInst translates the given x86 PSLLW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSLLW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRAD ] ---------------------------------------------------------------

// emitInst translates the given x86 PSRAD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSRAD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRAW ] ---------------------------------------------------------------

// emitInst translates the given x86 PSRAW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSRAW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRLD ] ---------------------------------------------------------------

// emitInst translates the given x86 PSRLD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSRLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRLDQ ] --------------------------------------------------------------

// emitInst translates the given x86 PSRLDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSRLDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRLQ ] ---------------------------------------------------------------

// emitInst translates the given x86 PSRLQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSRLQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSRLW ] ---------------------------------------------------------------

// emitInst translates the given x86 PSRLW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSRLW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBB ] ---------------------------------------------------------------

// emitInst translates the given x86 PSUBB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSUBB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBD ] ---------------------------------------------------------------

// emitInst translates the given x86 PSUBD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSUBD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBQ ] ---------------------------------------------------------------

// emitInst translates the given x86 PSUBQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSUBQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBSB ] --------------------------------------------------------------

// emitInst translates the given x86 PSUBSB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSUBSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBSW ] --------------------------------------------------------------

// emitInst translates the given x86 PSUBSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSUBSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBUSB ] -------------------------------------------------------------

// emitInst translates the given x86 PSUBUSB instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSUBUSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBUSW ] -------------------------------------------------------------

// emitInst translates the given x86 PSUBUSW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPSUBUSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PSUBW ] ---------------------------------------------------------------

// emitInst translates the given x86 PSUBW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPSUBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PTEST ] ---------------------------------------------------------------

// emitInst translates the given x86 PTEST instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPTEST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKHBW ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKHBW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUNPCKHBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKHDQ ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKHDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUNPCKHDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKHQDQ ] ----------------------------------------------------------

// emitInst translates the given x86 PUNPCKHQDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUNPCKHQDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKHWD ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKHWD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUNPCKHWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKLBW ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKLBW instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUNPCKLBW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKLDQ ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKLDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUNPCKLDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKLQDQ ] ----------------------------------------------------------

// emitInst translates the given x86 PUNPCKLQDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUNPCKLQDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUNPCKLWD ] -----------------------------------------------------------

// emitInst translates the given x86 PUNPCKLWD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUNPCKLWD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSH ] ----------------------------------------------------------------

// emitInst translates the given x86 PUSH instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPUSH(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSHA ] ---------------------------------------------------------------

// emitInst translates the given x86 PUSHA instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPUSHA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSHAD ] --------------------------------------------------------------

// emitInst translates the given x86 PUSHAD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUSHAD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSHF ] ---------------------------------------------------------------

// emitInst translates the given x86 PUSHF instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPUSHF(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSHFD ] --------------------------------------------------------------

// emitInst translates the given x86 PUSHFD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUSHFD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PUSHFQ ] --------------------------------------------------------------

// emitInst translates the given x86 PUSHFQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstPUSHFQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ PXOR ] ----------------------------------------------------------------

// emitInst translates the given x86 PXOR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstPXOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RCL ] -----------------------------------------------------------------

// emitInst translates the given x86 RCL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstRCL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RCPPS ] ---------------------------------------------------------------

// emitInst translates the given x86 RCPPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstRCPPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RCPSS ] ---------------------------------------------------------------

// emitInst translates the given x86 RCPSS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstRCPSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RCR ] -----------------------------------------------------------------

// emitInst translates the given x86 RCR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstRCR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDFSBASE ] ------------------------------------------------------------

// emitInst translates the given x86 RDFSBASE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstRDFSBASE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDGSBASE ] ------------------------------------------------------------

// emitInst translates the given x86 RDGSBASE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstRDGSBASE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDMSR ] ---------------------------------------------------------------

// emitInst translates the given x86 RDMSR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstRDMSR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDPMC ] ---------------------------------------------------------------

// emitInst translates the given x86 RDPMC instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstRDPMC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDRAND ] --------------------------------------------------------------

// emitInst translates the given x86 RDRAND instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstRDRAND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDTSC ] ---------------------------------------------------------------

// emitInst translates the given x86 RDTSC instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstRDTSC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RDTSCP ] --------------------------------------------------------------

// emitInst translates the given x86 RDTSCP instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstRDTSCP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RET ] -----------------------------------------------------------------

// emitInst translates the given x86 RET instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstRET(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROL ] -----------------------------------------------------------------

// emitInst translates the given x86 ROL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstROL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROR ] -----------------------------------------------------------------

// emitInst translates the given x86 ROR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstROR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROUNDPD ] -------------------------------------------------------------

// emitInst translates the given x86 ROUNDPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstROUNDPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROUNDPS ] -------------------------------------------------------------

// emitInst translates the given x86 ROUNDPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstROUNDPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROUNDSD ] -------------------------------------------------------------

// emitInst translates the given x86 ROUNDSD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstROUNDSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ ROUNDSS ] -------------------------------------------------------------

// emitInst translates the given x86 ROUNDSS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstROUNDSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RSM ] -----------------------------------------------------------------

// emitInst translates the given x86 RSM instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstRSM(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RSQRTPS ] -------------------------------------------------------------

// emitInst translates the given x86 RSQRTPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstRSQRTPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ RSQRTSS ] -------------------------------------------------------------

// emitInst translates the given x86 RSQRTSS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstRSQRTSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SAHF ] ----------------------------------------------------------------

// emitInst translates the given x86 SAHF instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSAHF(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SAR ] -----------------------------------------------------------------

// emitInst translates the given x86 SAR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSAR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SBB ] -----------------------------------------------------------------

// emitInst translates the given x86 SBB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSBB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SCASB ] ---------------------------------------------------------------

// emitInst translates the given x86 SCASB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSCASB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SCASD ] ---------------------------------------------------------------

// emitInst translates the given x86 SCASD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSCASD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SCASQ ] ---------------------------------------------------------------

// emitInst translates the given x86 SCASQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSCASQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SCASW ] ---------------------------------------------------------------

// emitInst translates the given x86 SCASW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSCASW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETA ] ----------------------------------------------------------------

// emitInst translates the given x86 SETA instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETAE ] ---------------------------------------------------------------

// emitInst translates the given x86 SETAE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETAE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETB ] ----------------------------------------------------------------

// emitInst translates the given x86 SETB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETBE ] ---------------------------------------------------------------

// emitInst translates the given x86 SETBE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETBE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETE ] ----------------------------------------------------------------

// emitInst translates the given x86 SETE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETG ] ----------------------------------------------------------------

// emitInst translates the given x86 SETG instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETGE ] ---------------------------------------------------------------

// emitInst translates the given x86 SETGE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETGE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETL ] ----------------------------------------------------------------

// emitInst translates the given x86 SETL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETLE ] ---------------------------------------------------------------

// emitInst translates the given x86 SETLE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETLE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETNE ] ---------------------------------------------------------------

// emitInst translates the given x86 SETNE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETNE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETNO ] ---------------------------------------------------------------

// emitInst translates the given x86 SETNO instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETNO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETNP ] ---------------------------------------------------------------

// emitInst translates the given x86 SETNP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETNP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETNS ] ---------------------------------------------------------------

// emitInst translates the given x86 SETNS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETNS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETO ] ----------------------------------------------------------------

// emitInst translates the given x86 SETO instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETO(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETP ] ----------------------------------------------------------------

// emitInst translates the given x86 SETP instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETP(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SETS ] ----------------------------------------------------------------

// emitInst translates the given x86 SETS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSETS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SFENCE ] --------------------------------------------------------------

// emitInst translates the given x86 SFENCE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSFENCE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SGDT ] ----------------------------------------------------------------

// emitInst translates the given x86 SGDT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSGDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHL ] -----------------------------------------------------------------

// emitInst translates the given x86 SHL instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSHL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHLD ] ----------------------------------------------------------------

// emitInst translates the given x86 SHLD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSHLD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHR ] -----------------------------------------------------------------

// emitInst translates the given x86 SHR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSHR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHRD ] ----------------------------------------------------------------

// emitInst translates the given x86 SHRD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSHRD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHUFPD ] --------------------------------------------------------------

// emitInst translates the given x86 SHUFPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSHUFPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SHUFPS ] --------------------------------------------------------------

// emitInst translates the given x86 SHUFPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSHUFPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SIDT ] ----------------------------------------------------------------

// emitInst translates the given x86 SIDT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSIDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SLDT ] ----------------------------------------------------------------

// emitInst translates the given x86 SLDT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSLDT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SMSW ] ----------------------------------------------------------------

// emitInst translates the given x86 SMSW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSMSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SQRTPD ] --------------------------------------------------------------

// emitInst translates the given x86 SQRTPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSQRTPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SQRTPS ] --------------------------------------------------------------

// emitInst translates the given x86 SQRTPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSQRTPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SQRTSD ] --------------------------------------------------------------

// emitInst translates the given x86 SQRTSD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSQRTSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SQRTSS ] --------------------------------------------------------------

// emitInst translates the given x86 SQRTSS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSQRTSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STC ] -----------------------------------------------------------------

// emitInst translates the given x86 STC instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSTC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STD ] -----------------------------------------------------------------

// emitInst translates the given x86 STD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSTD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STI ] -----------------------------------------------------------------

// emitInst translates the given x86 STI instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSTI(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STMXCSR ] -------------------------------------------------------------

// emitInst translates the given x86 STMXCSR instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSTMXCSR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STOSB ] ---------------------------------------------------------------

// emitInst translates the given x86 STOSB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSTOSB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STOSD ] ---------------------------------------------------------------

// emitInst translates the given x86 STOSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSTOSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STOSQ ] ---------------------------------------------------------------

// emitInst translates the given x86 STOSQ instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSTOSQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STOSW ] ---------------------------------------------------------------

// emitInst translates the given x86 STOSW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSTOSW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ STR ] -----------------------------------------------------------------

// emitInst translates the given x86 STR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSTR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SUB ] -----------------------------------------------------------------

// emitInst translates the given x86 SUB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSUB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SUBPD ] ---------------------------------------------------------------

// emitInst translates the given x86 SUBPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSUBPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SUBPS ] ---------------------------------------------------------------

// emitInst translates the given x86 SUBPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSUBPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SUBSD ] ---------------------------------------------------------------

// emitInst translates the given x86 SUBSD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSUBSD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SUBSS ] ---------------------------------------------------------------

// emitInst translates the given x86 SUBSS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstSUBSS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SWAPGS ] --------------------------------------------------------------

// emitInst translates the given x86 SWAPGS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSWAPGS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SYSCALL ] -------------------------------------------------------------

// emitInst translates the given x86 SYSCALL instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSYSCALL(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SYSENTER ] ------------------------------------------------------------

// emitInst translates the given x86 SYSENTER instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSYSENTER(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SYSEXIT ] -------------------------------------------------------------

// emitInst translates the given x86 SYSEXIT instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSYSEXIT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ SYSRET ] --------------------------------------------------------------

// emitInst translates the given x86 SYSRET instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstSYSRET(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ TEST ] ----------------------------------------------------------------

// emitInst translates the given x86 TEST instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstTEST(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ TZCNT ] ---------------------------------------------------------------

// emitInst translates the given x86 TZCNT instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstTZCNT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UCOMISD ] -------------------------------------------------------------

// emitInst translates the given x86 UCOMISD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstUCOMISD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UCOMISS ] -------------------------------------------------------------

// emitInst translates the given x86 UCOMISS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstUCOMISS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UD1 ] -----------------------------------------------------------------

// emitInst translates the given x86 UD1 instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstUD1(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UD2 ] -----------------------------------------------------------------

// emitInst translates the given x86 UD2 instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstUD2(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UNPCKHPD ] ------------------------------------------------------------

// emitInst translates the given x86 UNPCKHPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstUNPCKHPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UNPCKHPS ] ------------------------------------------------------------

// emitInst translates the given x86 UNPCKHPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstUNPCKHPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UNPCKLPD ] ------------------------------------------------------------

// emitInst translates the given x86 UNPCKLPD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstUNPCKLPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ UNPCKLPS ] ------------------------------------------------------------

// emitInst translates the given x86 UNPCKLPS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstUNPCKLPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VERR ] ----------------------------------------------------------------

// emitInst translates the given x86 VERR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstVERR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VERW ] ----------------------------------------------------------------

// emitInst translates the given x86 VERW instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstVERW(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VMOVDQA ] -------------------------------------------------------------

// emitInst translates the given x86 VMOVDQA instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstVMOVDQA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VMOVDQU ] -------------------------------------------------------------

// emitInst translates the given x86 VMOVDQU instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstVMOVDQU(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VMOVNTDQ ] ------------------------------------------------------------

// emitInst translates the given x86 VMOVNTDQ instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstVMOVNTDQ(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VMOVNTDQA ] -----------------------------------------------------------

// emitInst translates the given x86 VMOVNTDQA instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstVMOVNTDQA(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ VZEROUPPER ] ----------------------------------------------------------

// emitInst translates the given x86 VZEROUPPER instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstVZEROUPPER(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ WBINVD ] --------------------------------------------------------------

// emitInst translates the given x86 WBINVD instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstWBINVD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ WRFSBASE ] ------------------------------------------------------------

// emitInst translates the given x86 WRFSBASE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstWRFSBASE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ WRGSBASE ] ------------------------------------------------------------

// emitInst translates the given x86 WRGSBASE instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstWRGSBASE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ WRMSR ] ---------------------------------------------------------------

// emitInst translates the given x86 WRMSR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstWRMSR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XABORT ] --------------------------------------------------------------

// emitInst translates the given x86 XABORT instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXABORT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XADD ] ----------------------------------------------------------------

// emitInst translates the given x86 XADD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstXADD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XBEGIN ] --------------------------------------------------------------

// emitInst translates the given x86 XBEGIN instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXBEGIN(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XCHG ] ----------------------------------------------------------------

// emitInst translates the given x86 XCHG instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstXCHG(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XEND ] ----------------------------------------------------------------

// emitInst translates the given x86 XEND instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstXEND(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XGETBV ] --------------------------------------------------------------

// emitInst translates the given x86 XGETBV instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXGETBV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XLATB ] ---------------------------------------------------------------

// emitInst translates the given x86 XLATB instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstXLATB(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XOR ] -----------------------------------------------------------------

// emitInst translates the given x86 XOR instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstXOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XORPD ] ---------------------------------------------------------------

// emitInst translates the given x86 XORPD instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstXORPD(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XORPS ] ---------------------------------------------------------------

// emitInst translates the given x86 XORPS instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstXORPS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XRSTOR ] --------------------------------------------------------------

// emitInst translates the given x86 XRSTOR instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXRSTOR(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XRSTOR64 ] ------------------------------------------------------------

// emitInst translates the given x86 XRSTOR64 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXRSTOR64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XRSTORS ] -------------------------------------------------------------

// emitInst translates the given x86 XRSTORS instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXRSTORS(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XRSTORS64 ] -----------------------------------------------------------

// emitInst translates the given x86 XRSTORS64 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXRSTORS64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVE ] ---------------------------------------------------------------

// emitInst translates the given x86 XSAVE instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstXSAVE(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVE64 ] -------------------------------------------------------------

// emitInst translates the given x86 XSAVE64 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXSAVE64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVEC ] --------------------------------------------------------------

// emitInst translates the given x86 XSAVEC instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXSAVEC(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVEC64 ] ------------------------------------------------------------

// emitInst translates the given x86 XSAVEC64 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXSAVEC64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVEOPT ] ------------------------------------------------------------

// emitInst translates the given x86 XSAVEOPT instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXSAVEOPT(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVEOPT64 ] ----------------------------------------------------------

// emitInst translates the given x86 XSAVEOPT64 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXSAVEOPT64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVES ] --------------------------------------------------------------

// emitInst translates the given x86 XSAVES instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXSAVES(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSAVES64 ] ------------------------------------------------------------

// emitInst translates the given x86 XSAVES64 instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXSAVES64(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XSETBV ] --------------------------------------------------------------

// emitInst translates the given x86 XSETBV instruction to LLVM IR, emitting
// code to f.
func (d *disassembler) emitInstXSETBV(f *function, inst *instruction) error {
	panic("not yet implemented")
}

// --- [ XTEST ] ---------------------------------------------------------------

// emitInst translates the given x86 XTEST instruction to LLVM IR, emitting code
// to f.
func (d *disassembler) emitInstXTEST(f *function, inst *instruction) error {
	panic("not yet implemented")
}
