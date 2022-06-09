package nescpu

import "fmt"

var opUnimpl opFunc = func(cpu *CPU6502, arg uint16) int64 {
	fmt.Print("Operation unimplemented.\t")
	return 0
}

var addrUnimpl addrFunc = func(cpu *CPU6502) uint16 {
	fmt.Print("Addressing mode unimplemented.\t")
	return 0
}

func addrZpX(cpu *CPU6502) uint16       { return 0 }
func addrZpY(cpu *CPU6502) uint16       { return 0 }
func addrIndX(cpu *CPU6502) uint16      { return 0 }
func addrIndY(cpu *CPU6502) uint16      { return 0 }
func addrZp(cpu *CPU6502) uint16        { return 0 }
func addrImmediate(cpu *CPU6502) uint16 { return 0 }
func addrAbs(cpu *CPU6502) uint16       { return 0 }
func addrAbsY(cpu *CPU6502) uint16      { return 0 }
func addrAbsX(cpu *CPU6502) uint16      { return 0 }
func addrRelative(cpu *CPU6502) uint16  { return 0 }
func addrIndirect(cpu *CPU6502) uint16  { return 0 }
func addrImplied(cpu *CPU6502) uint16   { return 0 }
func addrAccum(cpu *CPU6502) uint16     { return 0 }

func opBpl(cpu *CPU6502, arg uint16 /*signed char*/) int64 { return 0 }
func opBrk(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opOra(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opAslm(cpu *CPU6502, arg uint16 /*int*/) int64        { return 0 }
func opAsla(cpu *CPU6502, arg uint16) int64                { return 0 }
func opPhp(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opClc(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opJsr(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opBit(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opAnd(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opRolm(cpu *CPU6502, arg uint16 /*int*/) int64        { return 0 }
func opPlp(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opRola(cpu *CPU6502, arg uint16) int64                { return 0 }
func opBmi(cpu *CPU6502, arg uint16 /*signed char*/) int64 { return 0 }
func opSec(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opRti(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opEor(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opLsrm(cpu *CPU6502, arg uint16 /*int*/) int64        { return 0 }
func opPha(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opLsra(cpu *CPU6502, arg uint16) int64                { return 0 }
func opJmp(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opBvc(cpu *CPU6502, arg uint16 /*signed char*/) int64 { return 0 }
func opCli(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opRts(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opAdc(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opRorm(cpu *CPU6502, arg uint16 /*int*/) int64        { return 0 }
func opPla(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opRora(cpu *CPU6502, arg uint16) int64                { return 0 }
func opBvs(cpu *CPU6502, arg uint16 /*signed char*/) int64 { return 0 }
func opSei(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opSty(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opSta(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opStx(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opDey(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opTxa(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opBcc(cpu *CPU6502, arg uint16 /*signed char*/) int64 { return 0 }
func opTya(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opTxs(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opLdy(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opLda(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opLdx(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opTay(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opTax(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opBcs(cpu *CPU6502, arg uint16 /*signed char*/) int64 { return 0 }
func opClv(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opTsx(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opCpy(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opCmp(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opCec(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opIny(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opDex(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opBne(cpu *CPU6502, arg uint16 /*signed char*/) int64 { return 0 }
func opCld(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opCpx(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opSbc(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opInc(cpu *CPU6502, arg uint16 /*int*/) int64         { return 0 }
func opInx(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opNop(cpu *CPU6502, arg uint16) int64                 { return 0 }
func opBeq(cpu *CPU6502, arg uint16 /*signed char*/) int64 { return 0 }
func opSed(cpu *CPU6502, arg uint16) int64                 { return 0 }
