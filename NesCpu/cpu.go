package NesCpu

import (
	"fmt"
	"github.com/khedoros/ghostliNES/NesApu"
	"github.com/khedoros/ghostliNES/NesMem"
)

type StatReg struct {
	Carry, Zero, Interrupt, Dec, Break, True, Verflow, Sign bool
}

type Cpu6502 struct {
	status                  StatReg
	pc                      uint16
	areg, xreg, yreg, spreg byte
	mem                     *NesMem.NesMem
	ops [256]func() int
}

func (this *Cpu6502) New(m *NesMem.NesMem) {
	this.status = StatReg{false, false, false, false, false, true, false, false}
	this.pc = 0x0000
	this.areg, this.xreg, this.yreg, this.spreg = 0, 0, 0, 0
	fmt.Println("init'd CPU")
	mem := m
	apu := NesApu.NesApu{}
	fmt.Println(mem.Blah, apu.Blah)
	for i:=0; i<256; i++ {
		this.ops[i] = this.opc(byte(i),addr_map[i],op_map[i])
	}
}

var runtime = [256]int { 7,6,0,0, 0,3,5,0, 3,2,2,0, 0,4,6,0,
                         2,5,0,0, 0,4,6,0, 2,4,0,0, 0,4,7,0,
                         6,6,0,0, 3,3,5,0, 4,2,2,0, 4,4,6,0,
                         2,5,0,0, 0,4,6,0, 2,4,0,0, 0,4,7,0,

                         4,6,0,0, 0,3,5,0, 3,2,2,0, 3,6,6,0,
                         2,5,0,0, 0,4,6,0, 2,4,0,0, 0,4,7,0,
                         6,6,0,0, 0,3,5,0, 4,2,2,0, 5,4,6,0,
                         2,5,0,0, 0,4,6,0, 2,4,0,0, 0,4,7,0,

                         0,6,0,0, 3,3,3,0, 2,0,2,0, 4,4,4,0,
                         2,6,0,0, 4,4,4,0, 2,5,2,0, 0,5,0,0,
                         2,6,2,0, 3,3,3,0, 2,2,2,0, 4,4,4,0,
                         2,5,0,0, 4,4,4,0, 2,4,2,0, 4,4,4,0,

                         2,6,0,0, 3,3,5,0, 2,2,2,0, 4,4,6,0,
                         2,5,0,0, 0,4,6,0, 2,4,0,0, 0,4,7,0,
                         2,6,0,0, 3,3,5,0, 2,2,2,0, 4,4,6,0,
                         2,5,0,0, 0,4,6,0, 2,4,0,0, 0,4,7,0,}

type op_func func(*Cpu6502, uint16) int
type addr_func func(*Cpu6502) uint16

var op_unimpl op_func = func(this *Cpu6502, arg uint16) int {
	fmt.Println("Joke's on you. That function doesn't exist.")
	return 0
}

var addr_unimpl addr_func = func(this *Cpu6502) uint16 {
	fmt.Println("Addressing mode unimplemented.")
	return 0
}

func (cpu *Cpu6502) opc(code byte, a addr_func, o op_func) func() int {
	return func() int {return runtime[code] + o(cpu, a(cpu))}
}

var addr_map = [256]addr_func {addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
                               addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl, }

var op_map = [256]op_func {op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
                           op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl, }

//        void set_sign(unsigned char);                   set sign true if data >= 128 else set sign false
//        void set_zero(unsigned char);                   set zero true if data == 0   else set zero false
//        void set_carry(unsigned char);                  set carry true if data != 0 else set carry false
//        void set_verflow(unsigned char, unsigned char); umm...no-op, apparently? probably not correct for the system...

//        void push(unsigned char val); write value to 0x100 + sp, dec sp
//        void push2(unsigned int val); write high to 0x100+sp, dec sp, write low to 0x100+sp, dec sp
//        unsigned int pop(); inc sp, read value from 0x100+sp
//        unsigned int pop2(); inc sp, read low from 0x100+sp, inc sp, read hi from 0x100+sp

