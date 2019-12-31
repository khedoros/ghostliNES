package NesCart

import (
	"fmt"
	"io/ioutil"
)

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

type Mapper interface {
	Read(addr uint16) uint8
	Write(cycle uint32, addr uint16, data uint8)
	Ppu_read(cycle uint32, addr uint16)
	Ppu_write(cycle uint32, addr uint16, data uint8)
	New()
}

type NesHeader struct {
	Prg_size int //byte 4, 16KB units
	Chr_size int //byte 5, 8KB units
	Chr_ram bool //byte 5 == 0
	Mapper_num Mapper_type //high-4 of byte 6 (low)  and high-4 of byte 7 (high)
	Mirroring Mirror_type //byte 6 bit 0 and bit 3
	Battery bool //byte 6 bit 1
	Trainer bool //byte 6 bit 2
	Sys_type Sys_type //byte 7 bits 0,1
	Prg_ram bool //byte 10 bit 4
	Prg_ram_size int //byte 8, 8KB units
	Tv_type Tv_type //byte 9 bit 0, byte 10 bits 0+1
	NES_2_0 bool //byte 7 bits 2,3 == 0x02
	Bus_conflicts bool //byte 10 bit 5
}
//if byte7 AND 0x0C == 0x08, and ROM size, taking into account byte 9 doesn't exceed actual filesize, then NES2.0
//if byte7 AND 0x0C == 0x00, and bytes 12-15 == 0 then full iNES
//otherwise, archaic iNES

type NesCart struct {
	Prg_rom []uint8
	Chr_rom []uint8
	Sram []uint8
	Mapper *Mapper
	Header NesHeader
}

type Mirror_type int
const (
	HMIRROR mirror_type = iota
	VMIRROR
	FOURSCREEN
)

type Mapper_type int
const (
	NROM mapper_type = iota //247 games
	MMC1 //680 games
	UxROM //269 games
	CNROM //155 games
	MMC3 //599 games
	MMC5 //24 games
	GAME_DOC //SKIP
	AxROM //75 games
	FFE_DISK //See GAME_DOC docs for more details SKIP
	MMC2 //Punch-Out
	MMC4 //3 games SKIP
	COLOR_DREAMS //31 games
	MMC3A //MMC3 variant
	CPROM
	GOUDER_HUANG //See MMC3 and VRC2 mappers
	K1029
	BANDAI_FCG
	FFSM_DISK
	JALECO_SS88006
	NAMCO //20 games
	FDS
	VRC4A
	VRC2A
	VRC2B
	VRC6A //1 game
	VRC2C
	VRC6B //2 games
	VRC4_PIRATE
	ACTION_53
	HOMEBREW
	UNROM_512
	NSF
)

type Sys_type int
const (
	VSUnisystem
	Playchoice10
	NES
	Extended
)

type Tv_type int
const (
	NTSC
	PAL
	DUAL
)

func (cart *NesCart) Load(filename *string) bool {
	contents, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("File length: ",len(contents))
	header := contents[:16]
	if !Equal(header[:4], []byte{'N','E','S',0x1a}) {
		fmt.Println("Header doesn't look correct")
		return false
	}
	for _,c := range(header) {
		fmt.Printf("%02x ", c)
	}
	return true
}
