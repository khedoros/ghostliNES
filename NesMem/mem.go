package nesmem

import (
	"fmt"

	nesapu "github.com/khedoros/ghostliNES/NesApu"
	nescart "github.com/khedoros/ghostliNES/NesCart"
	nesppu "github.com/khedoros/ghostliNES/NesPpu"
	"github.com/veandco/go-sdl2/sdl"
)

//An NesMem struct holds the state of the NES's memory mapping circuitry
type NesMem struct {
	cart *nescart.NesCart
	ppu  *nesppu.NesPpu
	apu  *nesapu.NesApu
	ram  [0x800]uint8
}

func (this *NesMem) InputEvent(event *sdl.Event) {

}

func (this *NesMem) New(filename *string, mapper int, ppu *nesppu.NesPpu, apu *nesapu.NesApu) {
	this.cart = &nescart.NesCart{}
	fmt.Println("Loading file ", *filename)
	valid := this.cart.Load(filename)
	if !valid {
		fmt.Println("File failed to load")
	} else {
		fmt.Println("Loaded ROM.")
	}
}

func (this *NesMem) Read(addr uint16, cycle uint64) uint8 {
	if addr < 0x2000 {
		//fmt.Printf("Read %02x from %04x\n", this.ram[addr&0x7ff], addr)
		return this.ram[addr&0x7ff]
	} else if addr < 0x4000 {
		return this.ppu.Read(addr, cycle)
	} else if addr < 0x4020 {
		fmt.Printf("read addr: %x: I/O and APU space\n", addr)
		return 0
	} else {
		return this.cart.Read(addr, cycle)
	}
}

func (this *NesMem) Read16(addr uint16, cycle uint64) uint16 {
	mem := uint16(this.Read(addr+1, cycle)) << 8
	mem += uint16(this.Read(addr, cycle))
	return mem
}

func (this *NesMem) Write(addr uint16, val uint8, cycle uint64) {
	if addr < 0x2000 {
		//fmt.Printf("write %02x to %04x\n", val, addr)
		this.ram[addr&0x7ff] = val
	} else if addr < 0x4000 {
		this.ppu.Write(addr, val, cycle)
	} else if addr < 0x4020 {
		fmt.Printf("write addr: %x val: %x: I/O and APU space\n", addr, val)
	} else {
		this.cart.Write(addr, val, cycle)
	}
}

func (this *NesMem) GetCart() *nescart.NesCart {
	return this.cart
}
