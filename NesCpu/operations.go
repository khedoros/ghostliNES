package nescpu

import "fmt"

var opUnimpl opFunc = func(cpu *CPU6502, arg uint16) uint {
	fmt.Print("Operation unimplemented.\t")
	return 0
}

// Opcode implementations
func opAdc(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	carry := uint8(0)
	if cpu.status.Carry {
		carry = 1
	}
	//val2 := int16(int8(val)) + int16(int8(cpu.areg)) + int16(int8(carry))
	val2 := val + cpu.areg + carry

	cpu.status.Verflow = !(((cpu.areg ^ val) & 0x80) == 0x80) && (((cpu.areg ^ uint8(val2)) & 0x80) == 0x80)
	cpu.zeroFlagNote = uint8(val2)
	cpu.negFlagNote = uint8(val2)
	cpu.status.Carry = (uint16(val) + uint16(cpu.areg) + uint16(carry)) > 0xff
	cpu.areg = uint8(val2)

	return 0
}

func opAnd(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.areg &= val
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opAslm(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.status.Carry = val&0x80 > 0
	val <<= 1
	cpu.zeroFlagNote = val
	cpu.negFlagNote = val
	cpu.mem.Write(arg, val, cpu.cycle)
	return 0
}

func opAsla(cpu *CPU6502, arg uint16) uint {
	cpu.status.Carry = cpu.areg&0x80 > 0
	cpu.areg <<= 1
	cpu.zeroFlagNote = cpu.areg
	cpu.negFlagNote = cpu.areg
	return 0
}

func opBcc(cpu *CPU6502, arg uint16 /*signed char*/) uint {
	if !cpu.status.Carry {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := uint(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBcs(cpu *CPU6502, arg uint16 /*signed char*/) uint {
	if cpu.status.Carry {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := uint(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBeq(cpu *CPU6502, arg uint16 /*signed char*/) uint {
	if cpu.zeroFlagNote == 0 {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := uint(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBit(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.zeroFlagNote = val & cpu.areg
	cpu.negFlagNote = val
	cpu.status.Verflow = (val&0x40 > 0)
	return 0
}

func opBmi(cpu *CPU6502, arg uint16 /*signed char*/) uint {
	if cpu.negFlagNote&0x80 == 0x80 {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := uint(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBne(cpu *CPU6502, arg uint16 /*signed char*/) uint {
	if cpu.zeroFlagNote != 0 {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := uint(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBpl(cpu *CPU6502, arg uint16 /*signed char*/) uint {
	if cpu.negFlagNote&0x80 == 0 {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := uint(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBrk(cpu *CPU6502, arg uint16) uint {
	cpu.pc++
	cpu.push2(cpu.pc)
	cpu.push(cpu.getStatus() | 0x10)
	cpu.status.Interrupt = true
	cpu.pc = cpu.mem.Read16(IRQVector, cpu.cycle)
	//fmt.Printf("INTERRUPT: BRK\n")
	return 0
}

func opBvc(cpu *CPU6502, arg uint16 /*signed char*/) uint {
	if !cpu.status.Verflow {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := uint(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opBvs(cpu *CPU6502, arg uint16 /*signed char*/) uint {
	if cpu.status.Verflow {
		newAddr := cpu.pc + uint16(int8(cpu.mem.Read(arg, cpu.cycle)))
		extra := uint(1)
		if (newAddr & 0xff00) != (cpu.pc & 0xff00) {
			extra++
		}
		cpu.pc = newAddr
		return extra
	}
	return 0
}

func opClc(cpu *CPU6502, arg uint16) uint {
	cpu.status.Carry = false
	return 0
}

func opCld(cpu *CPU6502, arg uint16) uint {
	cpu.status.Dec = false
	return 0
}

func opCli(cpu *CPU6502, arg uint16) uint {
	cpu.status.Interrupt = false
	return 0
}

func opClv(cpu *CPU6502, arg uint16) uint {
	cpu.status.Verflow = false
	return 0
}

func opCmp(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.status.Carry = (cpu.areg >= val)
	val = cpu.areg - val
	cpu.negFlagNote = val
	cpu.zeroFlagNote = val
	//fmt.Printf("CMP Result: %02x\n", val)
	return 0
}

func opCpx(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.status.Carry = (cpu.xreg >= val)
	val = cpu.xreg - val
	cpu.negFlagNote = val
	cpu.zeroFlagNote = val
	return 0
}

func opCpy(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.status.Carry = (cpu.yreg >= val)
	val = cpu.yreg - val
	cpu.negFlagNote = val
	cpu.zeroFlagNote = val
	return 0
}

func opDec(cpu *CPU6502, arg uint16) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	val--
	cpu.mem.Write(arg, val, cpu.cycle)
	cpu.negFlagNote = val
	cpu.zeroFlagNote = val
	return 0
}

func opDex(cpu *CPU6502, arg uint16) uint {
	cpu.xreg--
	cpu.negFlagNote = cpu.xreg
	cpu.zeroFlagNote = cpu.xreg
	return 0
}

func opDey(cpu *CPU6502, arg uint16) uint {
	cpu.yreg--
	cpu.negFlagNote = cpu.yreg
	cpu.zeroFlagNote = cpu.yreg
	return 0
}

// Illegal 2-byte nop
func opDop(cpu *CPU6502, arg uint16) uint {
	_ = cpu.mem.Read(arg, cpu.cycle)
	return 0
}

func opEor(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.areg ^= val
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opInc(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	val++
	cpu.mem.Write(arg, val, cpu.cycle)
	cpu.negFlagNote = val
	cpu.zeroFlagNote = val
	return 0
}

func opInx(cpu *CPU6502, arg uint16) uint {
	cpu.xreg++
	cpu.negFlagNote = cpu.xreg
	cpu.zeroFlagNote = cpu.xreg
	return 0
}

func opIny(cpu *CPU6502, arg uint16) uint {
	cpu.yreg++
	cpu.negFlagNote = cpu.yreg
	cpu.zeroFlagNote = cpu.yreg
	return 0
}

func opJmp(cpu *CPU6502, arg uint16 /*int*/) uint {
	cpu.pc = arg
	return 0
}

func opJsr(cpu *CPU6502, arg uint16 /*int*/) uint {
	cpu.pc--
	cpu.push2(cpu.pc)
	cpu.pc = arg
	return 0
}

func opLax(cpu *CPU6502, arg uint16) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.areg = val
	cpu.xreg = val
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opLda(cpu *CPU6502, arg uint16 /*int*/) uint {
	cpu.areg = cpu.mem.Read(arg, cpu.cycle)
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opLdx(cpu *CPU6502, arg uint16 /*int*/) uint {
	cpu.xreg = cpu.mem.Read(arg, cpu.cycle)
	cpu.negFlagNote = cpu.xreg
	cpu.zeroFlagNote = cpu.xreg
	return 0
}

func opLdy(cpu *CPU6502, arg uint16 /*int*/) uint {
	cpu.yreg = cpu.mem.Read(arg, cpu.cycle)
	cpu.negFlagNote = cpu.yreg
	cpu.zeroFlagNote = cpu.yreg
	return 0
}

func opLsra(cpu *CPU6502, arg uint16) uint {
	cpu.negFlagNote = 0
	cpu.status.Carry = cpu.areg&0x01 == 0x01
	cpu.areg >>= 1
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opLsrm(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.negFlagNote = 0
	cpu.status.Carry = val&0x01 == 0x01
	val >>= 1
	cpu.zeroFlagNote = val
	cpu.mem.Write(arg, val, cpu.cycle)
	return 0
}

func opNop(cpu *CPU6502, arg uint16) uint {
	return 0
}

func opOra(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	cpu.areg |= val
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opPha(cpu *CPU6502, arg uint16) uint {
	cpu.push(cpu.areg)
	return 0
}

func opPhp(cpu *CPU6502, arg uint16) uint {
	cpu.push(cpu.getStatus() | 0x10)
	return 0
}

func opPla(cpu *CPU6502, arg uint16) uint {
	val := cpu.pop()
	cpu.negFlagNote = val
	cpu.zeroFlagNote = val
	cpu.areg = val
	return 0
}

func opPlp(cpu *CPU6502, arg uint16) uint {
	status := cpu.pop()
	cpu.setStatus(status)
	return 0
}

func opRola(cpu *CPU6502, arg uint16) uint {
	carry := uint8(0)
	if cpu.status.Carry {
		carry = 0x01
	}
	cpu.status.Carry = cpu.areg&0x80 == 0x80
	cpu.areg <<= 1
	cpu.areg += carry
	cpu.zeroFlagNote = cpu.areg
	cpu.negFlagNote = cpu.areg
	return 0
}

func opRolm(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	carry := uint8(0)
	if cpu.status.Carry {
		carry = 0x01
	}
	cpu.status.Carry = val&0x80 == 0x80
	val <<= 1
	val += carry
	cpu.zeroFlagNote = val
	cpu.negFlagNote = val
	cpu.mem.Write(arg, val, cpu.cycle)
	return 0
}

func opRora(cpu *CPU6502, arg uint16) uint {
	carry := uint8(0)
	if cpu.status.Carry {
		carry = 0x80
	}
	cpu.status.Carry = cpu.areg&0x01 == 0x01
	cpu.areg >>= 1
	cpu.areg += carry
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opRorm(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	carry := uint8(0)
	if cpu.status.Carry {
		carry = 0x80
	}
	cpu.status.Carry = val&0x01 == 0x01
	val >>= 1
	val += carry
	cpu.negFlagNote = val
	cpu.zeroFlagNote = val
	cpu.mem.Write(arg, val, cpu.cycle)
	return 0
}

func opRti(cpu *CPU6502, arg uint16) uint {
	cpu.setStatus(cpu.pop())
	cpu.pc = cpu.pop2()
	return 0
}

func opRts(cpu *CPU6502, arg uint16) uint {
	cpu.pc = cpu.pop2()
	cpu.pc++
	return 0
}

func opSbc(cpu *CPU6502, arg uint16 /*int*/) uint {
	val := cpu.mem.Read(arg, cpu.cycle)
	val2 := val
	carry := uint8(0)
	if !cpu.status.Carry {
		carry = 1
	}
	val += carry
	cpu.zeroFlagNote = cpu.areg - val
	cpu.negFlagNote = cpu.areg - val
	cpu.status.Verflow = (((cpu.areg ^ val) & 0x80) > 0) && ((cpu.areg^val2)&0x80) > 0
	cpu.status.Carry = cpu.areg >= val
	cpu.areg -= val
	return 0
}

func opSec(cpu *CPU6502, arg uint16) uint {
	cpu.status.Carry = true
	return 0
}

func opSed(cpu *CPU6502, arg uint16) uint {
	cpu.status.Dec = true
	return 0
}
func opSei(cpu *CPU6502, arg uint16) uint {
	cpu.status.Interrupt = true
	return 0
}

func opSta(cpu *CPU6502, arg uint16 /*int*/) uint {
	cpu.mem.Write(arg, cpu.areg, cpu.cycle)
	return 0
}

func opStx(cpu *CPU6502, arg uint16 /*int*/) uint {
	cpu.mem.Write(arg, cpu.xreg, cpu.cycle)
	return 0
}

func opSty(cpu *CPU6502, arg uint16 /*int*/) uint {
	cpu.mem.Write(arg, cpu.yreg, cpu.cycle)
	return 0
}

func opTax(cpu *CPU6502, arg uint16) uint {
	cpu.xreg = cpu.areg
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

func opTay(cpu *CPU6502, arg uint16) uint {
	cpu.yreg = cpu.areg
	cpu.negFlagNote = cpu.areg
	cpu.zeroFlagNote = cpu.areg
	return 0
}

// Illegal 3-byte nop
func opTop(cpu *CPU6502, arg uint16) uint {
	_ = cpu.mem.Read16(arg, cpu.cycle)
	return 0
}

func opTsx(cpu *CPU6502, arg uint16) uint {
	cpu.xreg = cpu.spreg
	cpu.negFlagNote = cpu.spreg
	cpu.zeroFlagNote = cpu.spreg
	return 0
}

func opTxa(cpu *CPU6502, arg uint16) uint {
	cpu.areg = cpu.xreg
	cpu.negFlagNote = cpu.xreg
	cpu.zeroFlagNote = cpu.xreg
	return 0
}

func opTxs(cpu *CPU6502, arg uint16) uint {
	cpu.spreg = cpu.xreg
	return 0
}

func opTya(cpu *CPU6502, arg uint16) uint {
	cpu.areg = cpu.yreg
	cpu.negFlagNote = cpu.yreg
	cpu.zeroFlagNote = cpu.yreg
	return 0
}
