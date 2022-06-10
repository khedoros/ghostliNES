package nescpu

import (
	"fmt"

	nesmem "github.com/khedoros/ghostliNES/NesMem"
)

type statReg struct {
	Carry, Zero, Interrupt, Dec, Break, True, Verflow, Sign bool
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
	cpu.status = statReg{false, false, false, false, false, true, false, false}
	cpu.pc = 0x0000
	cpu.areg, cpu.xreg, cpu.yreg, cpu.spreg, cpu.frameCycle = 0, 0, 0, 0, 0
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
		switch cpu_ops[op].OpSize {
		case 1:
			fmt.Printf("%04x: %02x\n", cpu.pc, op)
		case 2:
			fmt.Printf("%04x: %02x %02x\n", cpu.pc, op, cpu.mem.Read(cpu.pc+1, cpu.cycle))
		case 3:
			fmt.Printf("%04x: %02x %04x\n", cpu.pc, op, cpu.mem.Read16(cpu.pc+1, cpu.cycle))
		default:
			panic(fmt.Sprintf("%04x: %02x is an invalid operation\n", cpu.pc, op))
		}
		cpu.pc += cpu_ops[op].OpSize
		opCycles := cpu.ops[op]()
		cpu.frameCycle += opCycles
		cpu.cycle += uint64(opCycles)
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

func (cpu *CPU6502) opc(code byte, a addrFunc, o opFunc) func() int64 {
	return func() int64 { return cpu_ops[code].OpTime + o(cpu, a(cpu)) }
}

//        void set_sign(unsigned char);                   set sign true if data >= 128 else set sign false
//        void set_zero(unsigned char);                   set zero true if data == 0   else set zero false
//        void set_carry(unsigned char);                  set carry true if data != 0 else set carry false
//        void set_verflow(unsigned char, unsigned char); umm...no-op, apparently? probably not correct for the system...

//        void push(unsigned char val); write value to 0x100 + sp, dec sp
//        void push2(unsigned int val); write high to 0x100+sp, dec sp, write low to 0x100+sp, dec sp
//        unsigned int pop(); inc sp, read value from 0x100+sp
//        unsigned int pop2(); inc sp, read low from 0x100+sp, inc sp, read hi from 0x100+sp
