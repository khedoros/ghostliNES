package nescpu

var cpu_ops = [256]CPU6502instr{
	// 0x0_
	{OpSize: 1, OpTime: 7, OpFunc: opBrk, AddrFunc: addrImplied}, // BRK
	{OpSize: 2, OpTime: 6, OpFunc: opOra, AddrFunc: addrIndX},    // ORA IND_X
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 3, OpFunc: opDop, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opOra, AddrFunc: addrZp},  // ORA ZP
	{OpSize: 2, OpTime: 5, OpFunc: opAslm, AddrFunc: addrZp}, // ASL ZP
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 3, OpFunc: opPhp, AddrFunc: addrImplied},   // PHP
	{OpSize: 2, OpTime: 2, OpFunc: opOra, AddrFunc: addrImmediate}, // ORA IMM
	{OpSize: 1, OpTime: 2, OpFunc: opAsla, AddrFunc: addrAccum},    // ASLA
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opTop, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 4, OpFunc: opOra, AddrFunc: addrAbs},  // ORA ABS
	{OpSize: 3, OpTime: 6, OpFunc: opAslm, AddrFunc: addrAbs}, //ASL ABS
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0x1_
	{OpSize: 2, OpTime: 2, OpFunc: opBpl, AddrFunc: addrRelative}, //BPL
	{OpSize: 2, OpTime: 5, OpFunc: opOra, AddrFunc: addrIndY},     // ORA IND_Y
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 4, OpFunc: opDop, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 4, OpFunc: opOra, AddrFunc: addrZpX},  // ORA ZP_X
	{OpSize: 2, OpTime: 6, OpFunc: opAslm, AddrFunc: addrZpX}, // ASL ZP_X
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 2, OpFunc: opClc, AddrFunc: addrImplied}, // CLC
	{OpSize: 3, OpTime: 4, OpFunc: opOra, AddrFunc: addrAbsY},    // ORA ABS_Y
	{OpSize: 1, OpTime: 2, OpFunc: opNop, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opTop, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 4, OpFunc: opOra, AddrFunc: addrAbsX},  // ORA ABS_X
	{OpSize: 3, OpTime: 7, OpFunc: opAslm, AddrFunc: addrAbsX}, // ASL ABS_X
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0x2_
	{OpSize: 3, OpTime: 6, OpFunc: opJsr, AddrFunc: addrAbs},
	{OpSize: 2, OpTime: 6, OpFunc: opAnd, AddrFunc: addrIndX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 3, OpFunc: opBit, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opAnd, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 5, OpFunc: opRolm, AddrFunc: addrZp},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 4, OpFunc: opPlp, AddrFunc: addrImplied},
	{OpSize: 2, OpTime: 2, OpFunc: opAnd, AddrFunc: addrImmediate},
	{OpSize: 1, OpTime: 2, OpFunc: opRola, AddrFunc: addrAccum},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opBit, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 4, OpFunc: opAnd, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 6, OpFunc: opRolm, AddrFunc: addrAbs},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0x3_
	{OpSize: 2, OpTime: 2, OpFunc: opBmi, AddrFunc: addrRelative},
	{OpSize: 2, OpTime: 5, OpFunc: opAnd, AddrFunc: addrIndY},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 4, OpFunc: opDop, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 4, OpFunc: opAnd, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 6, OpFunc: opRolm, AddrFunc: addrZpX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 2, OpFunc: opSec, AddrFunc: addrImplied},
	{OpSize: 3, OpTime: 4, OpFunc: opAnd, AddrFunc: addrAbsY},
	{OpSize: 1, OpTime: 2, OpFunc: opNop, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opTop, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 4, OpFunc: opAnd, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 7, OpFunc: opRolm, AddrFunc: addrAbsX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0x4_
	{OpSize: 1, OpTime: 4, OpFunc: opRti, AddrFunc: addrImplied},
	{OpSize: 2, OpTime: 6, OpFunc: opEor, AddrFunc: addrIndX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 3, OpFunc: opDop, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opEor, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 5, OpFunc: opLsrm, AddrFunc: addrZp},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 3, OpFunc: opPha, AddrFunc: addrImplied},
	{OpSize: 2, OpTime: 2, OpFunc: opEor, AddrFunc: addrImmediate},
	{OpSize: 1, OpTime: 2, OpFunc: opLsra, AddrFunc: addrAccum},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 3, OpFunc: opJmp, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 6, OpFunc: opEor, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 6, OpFunc: opLsrm, AddrFunc: addrAbs},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0x5_
	{OpSize: 2, OpTime: 2, OpFunc: opBvc, AddrFunc: addrRelative},
	{OpSize: 2, OpTime: 5, OpFunc: opEor, AddrFunc: addrIndY},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 4, OpFunc: opDop, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 4, OpFunc: opEor, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 6, OpFunc: opLsrm, AddrFunc: addrZpX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 2, OpFunc: opCli, AddrFunc: addrImplied},
	{OpSize: 3, OpTime: 4, OpFunc: opEor, AddrFunc: addrAbsY},
	{OpSize: 1, OpTime: 2, OpFunc: opNop, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opTop, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 4, OpFunc: opEor, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 7, OpFunc: opLsrm, AddrFunc: addrAbsX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0x6_
	{OpSize: 1, OpTime: 6, OpFunc: opRts, AddrFunc: addrImplied},
	{OpSize: 2, OpTime: 6, OpFunc: opAdc, AddrFunc: addrIndX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 3, OpFunc: opDop, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opAdc, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 5, OpFunc: opRorm, AddrFunc: addrZp},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 4, OpFunc: opPla, AddrFunc: addrImplied},
	{OpSize: 2, OpTime: 2, OpFunc: opAdc, AddrFunc: addrImmediate},
	{OpSize: 1, OpTime: 2, OpFunc: opRora, AddrFunc: addrAccum},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 5, OpFunc: opJmp, AddrFunc: addrIndirect},
	{OpSize: 3, OpTime: 4, OpFunc: opAdc, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 6, OpFunc: opRorm, AddrFunc: addrAbs},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0x7_
	{OpSize: 2, OpTime: 2, OpFunc: opBvs, AddrFunc: addrRelative},
	{OpSize: 2, OpTime: 5, OpFunc: opAdc, AddrFunc: addrIndY},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 4, OpFunc: opDop, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 4, OpFunc: opAdc, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 6, OpFunc: opRorm, AddrFunc: addrZpX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 2, OpFunc: opSei, AddrFunc: addrImplied},
	{OpSize: 3, OpTime: 4, OpFunc: opAdc, AddrFunc: addrAbsY},
	{OpSize: 1, OpTime: 2, OpFunc: opNop, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opTop, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 4, OpFunc: opAdc, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 7, OpFunc: opRorm, AddrFunc: addrAbsX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0x8_
	{OpSize: 2, OpTime: 2, OpFunc: opDop, AddrFunc: addrImmediate},
	{OpSize: 2, OpTime: 6, OpFunc: opSta, AddrFunc: addrIndX},
	{OpSize: 2, OpTime: 2, OpFunc: opDop, AddrFunc: addrImmediate},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 3, OpFunc: opSty, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opSta, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opStx, AddrFunc: addrZp},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 2, OpFunc: opDey, AddrFunc: addrImplied},
	{OpSize: 2, OpTime: 2, OpFunc: opDop, AddrFunc: addrImmediate},
	{OpSize: 1, OpTime: 2, OpFunc: opTxa, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opSty, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 4, OpFunc: opSta, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 4, OpFunc: opStx, AddrFunc: addrAbs},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0x9_
	{OpSize: 2, OpTime: 2, OpFunc: opBcc, AddrFunc: addrRelative},
	{OpSize: 2, OpTime: 6, OpFunc: opSta, AddrFunc: addrIndY},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 4, OpFunc: opSty, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 4, OpFunc: opSta, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 4, OpFunc: opStx, AddrFunc: addrZpY},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 2, OpFunc: opTya, AddrFunc: addrImplied},
	{OpSize: 3, OpTime: 5, OpFunc: opSta, AddrFunc: addrAbsY},
	{OpSize: 1, OpTime: 2, OpFunc: opTxs, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 5, OpFunc: opSta, AddrFunc: addrAbsX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0xA_
	{OpSize: 2, OpTime: 2, OpFunc: opLdy, AddrFunc: addrImmediate},
	{OpSize: 2, OpTime: 6, OpFunc: opLda, AddrFunc: addrIndX},
	{OpSize: 2, OpTime: 2, OpFunc: opLdx, AddrFunc: addrImmediate},
	{OpSize: 2, OpTime: 6, OpFunc: opLax, AddrFunc: addrIndX},
	{OpSize: 2, OpTime: 3, OpFunc: opLdy, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opLda, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opLdx, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opLax, AddrFunc: addrZp},
	{OpSize: 1, OpTime: 2, OpFunc: opTay, AddrFunc: addrImplied},
	{OpSize: 2, OpTime: 2, OpFunc: opLda, AddrFunc: addrImmediate},
	{OpSize: 1, OpTime: 2, OpFunc: opTax, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opLdy, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 4, OpFunc: opLda, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 4, OpFunc: opLdx, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 4, OpFunc: opLax, AddrFunc: addrAbs},
	// 0xB_
	{OpSize: 2, OpTime: 2, OpFunc: opBcs, AddrFunc: addrRelative},
	{OpSize: 2, OpTime: 5, OpFunc: opLda, AddrFunc: addrIndY},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 5, OpFunc: opLax, AddrFunc: addrIndY},
	{OpSize: 2, OpTime: 4, OpFunc: opLdy, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 4, OpFunc: opLda, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 4, OpFunc: opLdx, AddrFunc: addrZpY},
	{OpSize: 2, OpTime: 4, OpFunc: opLax, AddrFunc: addrZpY},
	{OpSize: 1, OpTime: 2, OpFunc: opClv, AddrFunc: addrImplied},
	{OpSize: 3, OpTime: 4, OpFunc: opLda, AddrFunc: addrAbsY},
	{OpSize: 1, OpTime: 2, OpFunc: opTsx, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opLdy, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 4, OpFunc: opLda, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 4, OpFunc: opLdx, AddrFunc: addrAbsY},
	{OpSize: 3, OpTime: 4, OpFunc: opLax, AddrFunc: addrAbsY},
	// 0xC_
	{OpSize: 2, OpTime: 2, OpFunc: opCpy, AddrFunc: addrImmediate},
	{OpSize: 2, OpTime: 6, OpFunc: opCmp, AddrFunc: addrIndX},
	{OpSize: 2, OpTime: 2, OpFunc: opDop, AddrFunc: addrImmediate},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 3, OpFunc: opCpy, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opCmp, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 5, OpFunc: opDec, AddrFunc: addrZp},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 2, OpFunc: opIny, AddrFunc: addrImplied},
	{OpSize: 2, OpTime: 2, OpFunc: opCmp, AddrFunc: addrImmediate},
	{OpSize: 1, OpTime: 2, OpFunc: opDex, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opCpy, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 4, OpFunc: opCmp, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 6, OpFunc: opDec, AddrFunc: addrAbs},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0xD_
	{OpSize: 2, OpTime: 2, OpFunc: opBne, AddrFunc: addrRelative},
	{OpSize: 2, OpTime: 5, OpFunc: opCmp, AddrFunc: addrIndY},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 4, OpFunc: opDop, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 4, OpFunc: opCmp, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 6, OpFunc: opDec, AddrFunc: addrZpX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 2, OpFunc: opCld, AddrFunc: addrImplied},
	{OpSize: 3, OpTime: 4, OpFunc: opCmp, AddrFunc: addrAbsY},
	{OpSize: 1, OpTime: 2, OpFunc: opNop, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opTop, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 4, OpFunc: opCmp, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 7, OpFunc: opDec, AddrFunc: addrAbsX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0xE_
	{OpSize: 2, OpTime: 2, OpFunc: opCpx, AddrFunc: addrImmediate},
	{OpSize: 2, OpTime: 6, OpFunc: opSbc, AddrFunc: addrIndX},
	{OpSize: 2, OpTime: 2, OpFunc: opDop, AddrFunc: addrImmediate},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 3, OpFunc: opCpx, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 3, OpFunc: opSbc, AddrFunc: addrZp},
	{OpSize: 2, OpTime: 5, OpFunc: opInc, AddrFunc: addrZp},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 2, OpFunc: opInx, AddrFunc: addrImplied},
	{OpSize: 2, OpTime: 2, OpFunc: opSbc, AddrFunc: addrImmediate},
	{OpSize: 1, OpTime: 2, OpFunc: opNop, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opCpx, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 4, OpFunc: opSbc, AddrFunc: addrAbs},
	{OpSize: 3, OpTime: 6, OpFunc: opInc, AddrFunc: addrAbs},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	// 0xF_
	{OpSize: 2, OpTime: 2, OpFunc: opBeq, AddrFunc: addrRelative},
	{OpSize: 2, OpTime: 5, OpFunc: opSbc, AddrFunc: addrIndY},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 2, OpTime: 4, OpFunc: opDop, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 4, OpFunc: opSbc, AddrFunc: addrZpX},
	{OpSize: 2, OpTime: 6, OpFunc: opInc, AddrFunc: addrZpX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 1, OpTime: 2, OpFunc: opSed, AddrFunc: addrImplied},
	{OpSize: 3, OpTime: 4, OpFunc: opSbc, AddrFunc: addrAbsY},
	{OpSize: 1, OpTime: 2, OpFunc: opNop, AddrFunc: addrImplied},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
	{OpSize: 3, OpTime: 4, OpFunc: opTop, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 4, OpFunc: opSbc, AddrFunc: addrAbsX},
	{OpSize: 3, OpTime: 7, OpFunc: opInc, AddrFunc: addrAbsX},
	{OpSize: 0, OpTime: 0, OpFunc: opUnimpl, AddrFunc: addrUnimpl},
}
