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
	areg, xreg, yreg, spreg uint8
	mem                     *NesMem.NesMem
}

func (this *Cpu6502) New(m *NesMem.NesMem) {
	this.status = StatReg{false, false, false, false, false, true, false, false}
	this.pc = 0x0000
	this.areg, this.xreg, this.yreg, this.spreg = 0, 0, 0, 0
	fmt.Println("init'd CPU")
	mem := NesMem.NesMem{}
	apu := NesApu.NesApu{}
	fmt.Println(mem.Blah, apu.Blah)
}
