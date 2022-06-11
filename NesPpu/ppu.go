package nesppu

import (
	nescart "github.com/khedoros/ghostliNES/NesCart"
)

//An NesPpu represents an NES's Picture Processing Unit
type NesPpu struct {
	cart *nescart.NesCart
	vram []uint8
}

func (this *NesPpu) New(mem *nescart.NesCart) {
}

// 0-1FFF: CRAM/CROM in cartridge
// 2000-2FFF: VRAM in PPU
// 3000-3F00: Mirror of VRAM
// 3F00-3F1F: palette memory
// 3F00-3FFF: Mirror of palette memory
// 4000-FFFF: Mirror of 0-3FFF

func (this *NesPpu) Read(addr uint16, cycle uint64) uint8 {
	if addr < 0x2000 {
		return this.cart.ReadPpu(addr, cycle)
	}
	switch addr {
	case 0x2002:
		return 0x80
	default:
		return 0
	}
}

func (this *NesPpu) Write(addr uint16, val uint8, cycle uint64) {
	if addr < 0x2000 {
		this.cart.WritePpu(addr, val, cycle)
	}

}

func (this *NesPpu) Run(cycles int64) int64 {
	return 0
}
