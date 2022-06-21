package nescpu

import (
	"fmt"

	nesmem "github.com/khedoros/ghostliNES/NesMem"
)

type statReg struct {
	Carry, Zero, Interrupt, Dec, True, Verflow, Sign bool
}

const (
	NMIVector = uint16(0xfffa)
	RSTVector = uint16(0xfffc)
	IRQVector = uint16(0xfffe)
)

//CPU6502 defines the structs necessary to hold the NES CPU's registers and table of opcodes.
type CPU6502 struct {
	status                  statReg
	pc                      uint16
	areg, xreg, yreg, spreg byte
	mem                     *nesmem.NesMem
	ops                     [256]func() int64
	frameCycle              int64
	cycle                   uint64
	zeroFlagNote            byte
	negFlagNote             byte
}

//New initializes a CPU6502 struct with its initial values
func (cpu *CPU6502) New(m *nesmem.NesMem) {
	cpu.status = statReg{Carry: false, Zero: false, Interrupt: true, Dec: false, True: true, Verflow: false, Sign: false}
	cpu.negFlagNote, cpu.zeroFlagNote = 0, 1
	cpu.pc = 0x0000
	cpu.areg, cpu.xreg, cpu.yreg, cpu.spreg, cpu.frameCycle = 0, 0, 0, 0xfd, 0
	fmt.Println("init'd CPU")
	cpu.mem = m
	for i := 0; i < 256; i++ {
		cpu.ops[i] = cpu.opc(byte(i), cpu_ops[i].AddrFunc, cpu_ops[i].OpFunc)
	}
	cpu.pc = cpu.mem.Read16(RSTVector, 0)
	fmt.Printf("Read address %04x and got vector %04x\n", RSTVector, cpu.pc)
}

func (cpu *CPU6502) Run(cycles int64) {
	for cpu.frameCycle < cycles {
		op := cpu.mem.Read(cpu.pc, cpu.cycle+uint64(cpu.cycle))
		/*
			switch cpu_ops[op].OpSize {
			case 1:
				fmt.Printf("%04X  %02X      ", cpu.pc, op)
			case 2:
				fmt.Printf("%04X  %02X %02X   ", cpu.pc, op, cpu.mem.Read(cpu.pc+1, cpu.cycle))
			case 3:
				tmp := cpu.mem.Read16(cpu.pc+1, cpu.cycle)
				fmt.Printf("%04X  %02X %02X %02X", cpu.pc, op, tmp&0xff, tmp>>8)
			default:
				panic(fmt.Sprintf("%04x  %02x is an invalid operation\n", cpu.pc, op))
			}
			fmt.Printf("    A:%02X X:%02X Y:%02X P:%02X SP:%02X\n", cpu.areg, cpu.xreg, cpu.yreg, cpu.getStatus(), cpu.spreg)
		*/
		cpu.pc += cpu_ops[op].OpSize
		opCycles := cpu.ops[op]()
		if opCycles == 0 {
			panic(fmt.Sprintf("Zero-time operation found at pc %04x: %02x", cpu.pc, op))
		}
		cpu.frameCycle += opCycles
		cpu.cycle += uint64(opCycles)

		if cpu.mem.IsPpuNmi(cpu.cycle) {
			cpu.nmi()
		}
	}
	cpu.frameCycle -= cycles
}

type CPU6502instr struct {
	OpSize      uint16
	OpTime      int64
	OpExtraTime int64
	OpFunc      func(*CPU6502, uint16) int64
	AddrFunc    func(*CPU6502) uint16
}

type opFunc func(*CPU6502, uint16) int64
type addrFunc func(*CPU6502) uint16

func (cpu *CPU6502) nmi() {
	cpu.push2(cpu.pc)
	cpu.push(cpu.getStatus())
	cpu.status.Interrupt = true
	//fmt.Printf("INTERRUPT: NMI from PC: %04x\n", cpu.pc)
	cpu.pc = cpu.mem.Read16(NMIVector, cpu.cycle)
}

func (cpu *CPU6502) irq() {
	cpu.push2(cpu.pc)
	cpu.push(cpu.getStatus())
	cpu.status.Interrupt = true
	cpu.pc = cpu.mem.Read16(IRQVector, cpu.cycle)
	//fmt.Printf("INTERRUPT: IRQ\n")
}

func (cpu *CPU6502) opc(code byte, a addrFunc, o opFunc) func() int64 {
	return func() int64 { return cpu_ops[code].OpTime + o(cpu, a(cpu)) }
}

func (cpu *CPU6502) push2(val uint16) {
	low := byte(val & 0xff)
	high := byte(val >> 8)
	//fmt.Printf("Push2 %02x %02x to %04x and %04x\n", high, low, 0x100+uint16(cpu.spreg), 0x100+uint16(cpu.spreg-1))
	cpu.mem.Write(uint16(0x100)+uint16(cpu.spreg), high, cpu.cycle)
	cpu.spreg--
	cpu.mem.Write(uint16(0x100)+uint16(cpu.spreg), low, cpu.cycle)
	cpu.spreg--
}

func (cpu *CPU6502) push(val byte) {
	//fmt.Printf("Push1 %02x to %04x\n", val, 0x100+uint16(cpu.spreg))
	cpu.mem.Write(uint16(0x100)+uint16(cpu.spreg), val, cpu.cycle)
	cpu.spreg--
}

func (cpu *CPU6502) pop2() uint16 {
	cpu.spreg++
	low := cpu.mem.Read(uint16(0x100)+uint16(cpu.spreg), cpu.cycle)
	cpu.spreg++
	high := cpu.mem.Read(uint16(0x100)+uint16(cpu.spreg), cpu.cycle)
	//fmt.Printf("Pop2 %02x %02x from %04x and %04x\n", high, low, 0x100+uint16(cpu.spreg), 0x100+uint16(cpu.spreg-1))
	return (uint16(low) + (uint16(high) << 8))
}

func (cpu *CPU6502) pop() byte {
	cpu.spreg++
	val := cpu.mem.Read(uint16(0x100)+uint16(cpu.spreg), cpu.cycle)
	//fmt.Printf("Pop1 %02x from %04x\n", val, 0x100+uint16(cpu.spreg))
	return val
}

func (cpu *CPU6502) getStatus() byte {
	status := cpu.negFlagNote & 0x80
	if cpu.status.Carry {
		status |= 1
	}
	if cpu.zeroFlagNote == 0 {
		status |= 2
	}
	if cpu.status.Interrupt {
		status |= 4
	}
	if cpu.status.Dec {
		status |= 8
	}
	status |= 0x20
	if cpu.status.Verflow {
		status |= 0x40
	}
	return status
}

func (cpu *CPU6502) setStatus(status byte) {
	cpu.negFlagNote = status
	cpu.status.Verflow = status&0x40 == 0x40
	cpu.status.Dec = status&0x08 == 0x08
	cpu.status.Interrupt = status&0x04 == 0x04
	if status&2 == 2 {
		cpu.zeroFlagNote = 0
	} else {
		cpu.zeroFlagNote = 1
	}
	cpu.status.Carry = status&0x01 == 0x01
}
