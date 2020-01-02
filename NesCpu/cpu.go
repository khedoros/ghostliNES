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
}

func (this *Cpu6502) New(m *NesMem.NesMem) {
	this.status = StatReg{false, false, false, false, false, true, false, false}
	this.pc = 0x0000
	this.areg, this.xreg, this.yreg, this.spreg = 0, 0, 0, 0
	fmt.Println("init'd CPU")
	mem := m
	apu := NesApu.NesApu{}
	fmt.Println(mem.Blah, apu.Blah)
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


type op_func func(uint16)
type addr_func func() uint16

//func (this *Cpu6502) op_unimpl(arg ...byte) {
func (this *Cpu6502) op_unimpl(arg uint16) {
	fmt.Println("Joke's on you. That function doesn't exist.")
}
func (this *Cpu6502) addr_unimpl() uint16 {
	fmt.Println("Addressing mode unimplemented.")
	return 0
}

var addr = [256]addr_func {addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,addr_unimpl,
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

var op = [256]op_func {op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,op_unimpl,
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


