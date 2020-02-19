package nescpu

import (
	"fmt"
	nesapu "github.com/khedoros/ghostliNES/NesApu"
	nesmem "github.com/khedoros/ghostliNES/NesMem"
)

type statReg struct {
	Carry, Zero, Interrupt, Dec, Break, True, Verflow, Sign bool
}

//CPU6502 defines the structs necessary to hold the NES CPU's registers and table of opcodes.
type CPU6502 struct {
	status                  statReg
	pc                      uint16
	areg, xreg, yreg, spreg byte
	mem                     *nesmem.NesMem
	ops                     [256]func() int
}

//New initializes a CPU6502 struct with its initial values
func (cpu *CPU6502) New(m *nesmem.NesMem) {
	cpu.status = statReg{false, false, false, false, false, true, false, false}
	cpu.pc = 0x0000
	cpu.areg, cpu.xreg, cpu.yreg, cpu.spreg = 0, 0, 0, 0
	fmt.Println("init'd CPU")
	mem := m
	apu := nesapu.NesApu{}
	fmt.Println(mem.Read(0, 0), apu.Read(0, 0))
	for i := 0; i < 256; i++ {
		cpu.ops[i] = cpu.opc(byte(i), addrMap[i], opMap[i])
	}
}

type CPU6502instr struct {
	OpSize      int
	OpTime      int
	OpExtraTime int
	OpFunc      func(*CPU6502, uint16) int
	AddrFunc    func(*CPU6502) uint16
}

var runtime = [256]int{7, 6, 0, 0, 0, 3, 5, 0, 3, 2, 2, 0, 0, 4, 6, 0,
	2, 5, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,
	6, 6, 0, 0, 3, 3, 5, 0, 4, 2, 2, 0, 4, 4, 6, 0,
	2, 5, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,

	4, 6, 0, 0, 0, 3, 5, 0, 3, 2, 2, 0, 3, 6, 6, 0,
	2, 5, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,
	6, 6, 0, 0, 0, 3, 5, 0, 4, 2, 2, 0, 5, 4, 6, 0,
	2, 5, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,

	0, 6, 0, 0, 3, 3, 3, 0, 2, 0, 2, 0, 4, 4, 4, 0,
	2, 6, 0, 0, 4, 4, 4, 0, 2, 5, 2, 0, 0, 5, 0, 0,
	2, 6, 2, 0, 3, 3, 3, 0, 2, 2, 2, 0, 4, 4, 4, 0,
	2, 5, 0, 0, 4, 4, 4, 0, 2, 4, 2, 0, 4, 4, 4, 0,

	2, 6, 0, 0, 3, 3, 5, 0, 2, 2, 2, 0, 4, 4, 6, 0,
	2, 5, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0,
	2, 6, 0, 0, 3, 3, 5, 0, 2, 2, 2, 0, 4, 4, 6, 0,
	2, 5, 0, 0, 0, 4, 6, 0, 2, 4, 0, 0, 0, 4, 7, 0}

type opFunc func(*CPU6502, uint16) int
type addrFunc func(*CPU6502) uint16

var opUnimpl opFunc = func(cpu *CPU6502, arg uint16) int {
	fmt.Println("Joke's on you. That function doesn't exist.")
	return 0
}

var addrUnimpl addrFunc = func(cpu *CPU6502) uint16 {
	fmt.Println("Addressing mode unimplemented.")
	return 0
}

func (cpu *CPU6502) opc(code byte, a addrFunc, o opFunc) func() int {
	return func() int { return runtime[code] + o(cpu, a(cpu)) }
}

var addrMap = [256]addrFunc{
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl,
	addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl, addrUnimpl}

var opMap = [256]opFunc{
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl,
	opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl, opUnimpl}

//        void set_sign(unsigned char);                   set sign true if data >= 128 else set sign false
//        void set_zero(unsigned char);                   set zero true if data == 0   else set zero false
//        void set_carry(unsigned char);                  set carry true if data != 0 else set carry false
//        void set_verflow(unsigned char, unsigned char); umm...no-op, apparently? probably not correct for the system...

//        void push(unsigned char val); write value to 0x100 + sp, dec sp
//        void push2(unsigned int val); write high to 0x100+sp, dec sp, write low to 0x100+sp, dec sp
//        unsigned int pop(); inc sp, read value from 0x100+sp
//        unsigned int pop2(); inc sp, read low from 0x100+sp, inc sp, read hi from 0x100+sp
