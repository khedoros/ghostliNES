package nescpu

import "fmt"

var opUnimpl opFunc = func(cpu *CPU6502, arg uint16) int64 {
	fmt.Print("Operation unimplemented.\t")
	return 0
}

// Opcode implementations
func opAdc(cpu *CPU6502, arg uint16 /*int*/) int64 { return 0 }

func opAnd(cpu *CPU6502, arg uint16 /*int*/) int64 {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.areg &= val
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opAslm(cpu *CPU6502, arg uint16 /*int*/) int64 { return 0 }
func opAsla(cpu *CPU6502, arg uint16) int64         { return 0 }

func opBcc(cpu *CPU6502, arg uint16 /*signed char*/) int64 {
	if !cpu.status.Carry {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := int64(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBcs(cpu *CPU6502, arg uint16 /*signed char*/) int64 {
	if cpu.status.Carry {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := int64(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBeq(cpu *CPU6502, arg uint16 /*signed char*/) int64 {
	if cpu.zeroFlagNote == 0 {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := int64(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBit(cpu *CPU6502, arg uint16 /*int*/) int64 { return 0 }

func opBmi(cpu *CPU6502, arg uint16 /*signed char*/) int64 {
	if cpu.negFlagNote&0x80 == 0x80 {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := int64(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBne(cpu *CPU6502, arg uint16 /*signed char*/) int64 {
	if cpu.zeroFlagNote != 0 {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := int64(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBpl(cpu *CPU6502, arg uint16 /*signed char*/) int64 {
	if cpu.negFlagNote&0x80 == 0 {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := int64(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBrk(cpu *CPU6502, arg uint16) int64 { return 0 }

func opBvc(cpu *CPU6502, arg uint16 /*signed char*/) int64 {
	if !cpu.status.Verflow {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := int64(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBvs(cpu *CPU6502, arg uint16 /*signed char*/) int64 {
	if cpu.status.Verflow {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := int64(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opClc(cpu *CPU6502, arg uint16) int64 {
	cpu.status.Carry = false
	return 0
}

func opCld(cpu *CPU6502, arg uint16) int64 {
	cpu.status.Dec = false
	return 0
}

func opCli(cpu *CPU6502, arg uint16) int64 {
	cpu.status.Interrupt = false
	return 0
}

func opClv(cpu *CPU6502, arg uint16) int64 {
	cpu.status.Verflow = false
	return 0
}

func opCmp(cpu *CPU6502, arg uint16 /*int*/) int64 { return 0 }
func opCpx(cpu *CPU6502, arg uint16 /*int*/) int64 { return 0 }
func opCpy(cpu *CPU6502, arg uint16 /*int*/) int64 { return 0 }

func opDec(cpu *CPU6502, arg uint16) int64 {
	val := cpu.mem.Read(arg, cpu.cycle)
	val--
	cpu.mem.Write(arg, val, cpu.cycle)
	cpu.negFlagNote = val
	cpu.zeroFlagNote = val
	return 0
}

func opDex(cpu *CPU6502, arg uint16) int64 {
	cpu.xreg--
	cpu.negFlagNote = cpu.xreg
	cpu.zeroFlagNote = cpu.xreg
	return 0
}

func opDey(cpu *CPU6502, arg uint16) int64 {
	cpu.yreg++
	cpu.negFlagNote = cpu.yreg
	cpu.zeroFlagNote = cpu.yreg
	return 0
}

func opEor(cpu *CPU6502, arg uint16 /*int*/) int64 {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.areg ^= val
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opInc(cpu *CPU6502, arg uint16 /*int*/) int64 {
	val := cpu.mem.Read(arg, cpu.cycle)
	val++
	cpu.mem.Write(arg, val, cpu.cycle)
	cpu.negFlagNote = val
	cpu.zeroFlagNote = val
	return 0
}

func opInx(cpu *CPU6502, arg uint16) int64 {
	cpu.xreg++
	cpu.negFlagNote = cpu.xreg
	cpu.zeroFlagNote = cpu.xreg
	return 0
}

func opIny(cpu *CPU6502, arg uint16) int64 {
	cpu.yreg++
	cpu.negFlagNote = cpu.yreg
	cpu.zeroFlagNote = cpu.yreg
	return 0
}

func opJmp(cpu *CPU6502, arg uint16 /*int*/) int64 {
	cpu.pc = arg
	return 0
}

func opJsr(cpu *CPU6502, arg uint16 /*int*/) int64 { return 0 }

func opLda(cpu *CPU6502, arg uint16 /*int*/) int64 {
	cpu.areg = cpu.mem.Read(arg, cpu.cycle)
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opLdx(cpu *CPU6502, arg uint16 /*int*/) int64 {
	cpu.xreg = cpu.mem.Read(arg, cpu.cycle)
	cpu.negFlagNote = cpu.xreg
	cpu.zeroFlagNote = cpu.xreg
	return 0
}

func opLdy(cpu *CPU6502, arg uint16 /*int*/) int64 {
	cpu.yreg = cpu.mem.Read(arg, cpu.cycle)
	cpu.negFlagNote = cpu.yreg
	cpu.zeroFlagNote = cpu.yreg
	return 0
}

func opLsra(cpu *CPU6502, arg uint16) int64         { return 0 }
func opLsrm(cpu *CPU6502, arg uint16 /*int*/) int64 { return 0 }

func opNop(cpu *CPU6502, arg uint16) int64 {
	return 0
}

func opOra(cpu *CPU6502, arg uint16 /*int*/) int64 {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.areg |= val
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opPha(cpu *CPU6502, arg uint16) int64          { return 0 }
func opPhp(cpu *CPU6502, arg uint16) int64          { return 0 }
func opPla(cpu *CPU6502, arg uint16) int64          { return 0 }
func opPlp(cpu *CPU6502, arg uint16) int64          { return 0 }
func opRola(cpu *CPU6502, arg uint16) int64         { return 0 }
func opRolm(cpu *CPU6502, arg uint16 /*int*/) int64 { return 0 }
func opRora(cpu *CPU6502, arg uint16) int64         { return 0 }
func opRorm(cpu *CPU6502, arg uint16 /*int*/) int64 { return 0 }
func opRti(cpu *CPU6502, arg uint16) int64          { return 0 }
func opRts(cpu *CPU6502, arg uint16) int64          { return 0 }
func opSbc(cpu *CPU6502, arg uint16 /*int*/) int64  { return 0 }

func opSec(cpu *CPU6502, arg uint16) int64 {
	cpu.status.Carry = true
	return 0
}

func opSed(cpu *CPU6502, arg uint16) int64 {
	cpu.status.Dec = true
	return 0
}
func opSei(cpu *CPU6502, arg uint16) int64 {
	cpu.status.Interrupt = true
	return 0
}

func opSta(cpu *CPU6502, arg uint16 /*int*/) int64 {
	cpu.mem.Write(arg, cpu.areg, cpu.cycle)
	return 0
}

func opStx(cpu *CPU6502, arg uint16 /*int*/) int64 {
	cpu.mem.Write(arg, cpu.xreg, cpu.cycle)
	return 0
}

func opSty(cpu *CPU6502, arg uint16 /*int*/) int64 {
	cpu.mem.Write(arg, cpu.yreg, cpu.cycle)
	return 0
}

func opTax(cpu *CPU6502, arg uint16) int64 {
	cpu.xreg = cpu.areg
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opTay(cpu *CPU6502, arg uint16) int64 {
	cpu.yreg = cpu.areg
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opTsx(cpu *CPU6502, arg uint16) int64 {
	cpu.xreg = cpu.spreg
	cpu.negFlagNote = cpu.spreg
	cpu.zeroFlagNote = cpu.spreg
	return 0
}

func opTxa(cpu *CPU6502, arg uint16) int64 {
	cpu.areg = cpu.xreg
	cpu.negFlagNote = cpu.xreg
	cpu.zeroFlagNote = cpu.xreg
	return 0
}

func opTxs(cpu *CPU6502, arg uint16) int64 {
	cpu.spreg = cpu.xreg
	cpu.negFlagNote = cpu.xreg
	cpu.zeroFlagNote = cpu.xreg
	return 0
}

func opTya(cpu *CPU6502, arg uint16) int64 {
	cpu.areg = cpu.yreg
	cpu.negFlagNote = cpu.yreg
	cpu.zeroFlagNote = cpu.yreg
	return 0
}
