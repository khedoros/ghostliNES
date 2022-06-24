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
	cart      *nescart.NesCart
	ppu       *nesppu.NesPpu
	apu       *nesapu.NesApu
	ram       [0x800]uint8
	joy       [8]bool
	joyStrobe bool
	joyIndex  int
}

const (
	KEYA int = iota
	KEYB
	SELECT
	START
	UP
	DOWN
	LEFT
	RIGHT
)

func (this *NesMem) InputEvent(event *sdl.Event) {
	switch t := (*event).(type) {
	case *sdl.KeyboardEvent:
		pressed := t.State == sdl.PRESSED
		switch t.Keysym.Sym {
		case sdl.K_a:
			this.joy[LEFT] = pressed
		case sdl.K_s:
			this.joy[DOWN] = pressed
		case sdl.K_d:
			this.joy[RIGHT] = pressed
		case sdl.K_w:
			this.joy[UP] = pressed
		case sdl.K_g:
			this.joy[SELECT] = pressed
		case sdl.K_h:
			this.joy[START] = pressed
		case sdl.K_k:
			this.joy[KEYB] = pressed
		case sdl.K_l:
			this.joy[KEYA] = pressed
		}
	}
}

func (this *NesMem) New(filename *string, mapper int, ppu *nesppu.NesPpu, apu *nesapu.NesApu) {
	this.cart = &nescart.NesCart{}
	this.ppu = ppu
	this.apu = apu
	fmt.Println("Loading file ", *filename)
	valid := this.cart.Load(filename, mapper)
	if !valid {
		fmt.Println("File failed to load")
	} else {
		fmt.Println("Loaded ROM.")
	}
}

func (this *NesMem) IsPpuNmi(cycle uint64) bool {
	return this.ppu.IsNmi(cycle)
}

func (this *NesMem) Read(addr uint16, cycle uint64) uint8 {
	if addr < 0x2000 {
		//fmt.Printf("Read %02x from %04x\n", this.ram[addr&0x7ff], addr)
		return this.ram[addr&0x7ff]
	} else if addr < 0x4000 {
		return this.ppu.Read(addr, cycle)
	} else if addr < 0x4020 {
		if addr == 0x4016 { // Joy1
			if this.joyIndex >= 8 {
				return 0
			} else {
				val := uint8(0)
				if this.joy[this.joyIndex] {
					val = 1
				}
				this.joyIndex++
				return val
			}
		}
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
		this.ppu.Write((addr%8 + 0x2000), val, cycle)
	} else if addr == 0x4014 { // Sprite DMA
		base := uint16(val) * 0x100
		for i := uint16(0); i < 256; i++ {
			this.ppu.Write(0x4014, this.Read(base+i, cycle+uint64(i)*2), cycle+uint64(i)*2)
		}
	} else if addr == 0x4016 {
		if val&0x01 == 1 {
			this.joyStrobe = true
		} else if val&0x01 == 0 && this.joyStrobe {
			this.joyStrobe = false
			this.joyIndex = 0
		}
	} else if addr < 0x4020 {
		fmt.Printf("write addr: %x val: %x: I/O and APU space\n", addr, val)
	} else {
		this.cart.Write(addr, val, cycle)
	}
}

func (this *NesMem) GetCart() *nescart.NesCart {
	return this.cart
}
