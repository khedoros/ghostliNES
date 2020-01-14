package NesPpu

import (
	"github.com/khedoros/ghostliNES/NesCart"
)

//An NesPpu represents an NES's Picture Processing Unit
type NesPpu struct {
	Blah int8
}

func (this *NesPpu) New(mem *NesCart.NesCart, res int) {
}

func (this *NesPpu) Read(addr uint16, cycle uint64) uint8 {
	return 0
}

func (this *NesPpu) Write(addr uint16, val uint8, cycle uint64) {
}
