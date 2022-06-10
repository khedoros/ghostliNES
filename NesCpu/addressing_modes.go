package nescpu

import "fmt"

var addrUnimpl addrFunc = func(cpu *CPU6502) uint16 {
	fmt.Print("Addressing mode unimplemented.\t")
	return 0
}

// Addressing mode implementations
func addrIndX(cpu *CPU6502) uint16 {
	addr := uint16((cpu.mem.Read(cpu.pc-1, cpu.cycle) + cpu.xreg) & 0xff)
	retAddr := uint16(cpu.mem.Read(addr, cpu.cycle))
	retAddr += uint16(cpu.mem.Read((addr+1)&0xff, cpu.cycle)) << 8
	return retAddr
}

func addrIndY(cpu *CPU6502) uint16 {
	addr := uint16(cpu.mem.Read(cpu.pc-1, cpu.cycle))
	retAddr := uint16(cpu.mem.Read(addr, cpu.cycle))
	retAddr += uint16(cpu.mem.Read((addr+1)&0xff, cpu.cycle))<<8 + uint16(cpu.yreg)
	//return cpu.mem.Read16(addr, cpu.cycle) + uint16(cpu.yreg)
	return retAddr
}

func addrIndirect(cpu *CPU6502) uint16 {
	addr := cpu.mem.Read16(cpu.pc-2, cpu.cycle)
	page := addr & 0xff00
	addr2 := (addr+1)&0xff | page
	retAddr := uint16(cpu.mem.Read(addr, cpu.cycle))
	retAddr += uint16(cpu.mem.Read(addr2, cpu.cycle)) << 8
	return retAddr
}

func addrZp(cpu *CPU6502) uint16        { return uint16(cpu.mem.Read(cpu.pc-1, cpu.cycle)) }
func addrZpX(cpu *CPU6502) uint16       { return uint16(cpu.mem.Read(cpu.pc-1, cpu.cycle) + cpu.xreg) }
func addrZpY(cpu *CPU6502) uint16       { return uint16(cpu.mem.Read(cpu.pc-1, cpu.cycle) + cpu.yreg) }
func addrAbs(cpu *CPU6502) uint16       { return cpu.mem.Read16(cpu.pc-2, cpu.cycle) }
func addrAbsY(cpu *CPU6502) uint16      { return cpu.mem.Read16(cpu.pc-2, cpu.cycle) + uint16(cpu.yreg) }
func addrAbsX(cpu *CPU6502) uint16      { return cpu.mem.Read16(cpu.pc-2, cpu.cycle) + uint16(cpu.xreg) }
func addrImmediate(cpu *CPU6502) uint16 { return cpu.pc - 1 }
func addrRelative(cpu *CPU6502) uint16  { return cpu.pc - 1 }
func addrImplied(cpu *CPU6502) uint16   { return 0 }
func addrAccum(cpu *CPU6502) uint16     { return 0 }
