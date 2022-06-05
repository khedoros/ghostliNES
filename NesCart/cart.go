package nescart

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"strings"
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

//A Mapper handles mapping pages of data in a ROM to their physical addresses, visible to the CPU and PPU
type Mapper interface {
	Read(addr uint16) uint8
	Write(cycle uint32, addr uint16, data uint8)
	Ppu_read(cycle uint32, addr uint16)
	Ppu_write(cycle uint32, addr uint16, data uint8)
	New()
}

//NesHeader defines the 16-byte structure that makes up the iNES header in most NES ROMs
type NesHeader struct {
	PrgSize      int        //byte 4, 16KB units
	ChrSize      int        //byte 5, 8KB units
	ChrRAM       bool       //byte 5 == 0
	MapperNum    MapperType //high-4 of byte 6 (low)  and high-4 of byte 7 (high)
	Mirroring    MirrorType //byte 6 bit 0 and bit 3
	Battery      bool       //byte 6 bit 1
	Trainer      bool       //byte 6 bit 2
	SysType      SysType    //byte 7 bits 0,1
	PrgRAM       bool       //byte 10 bit 4
	PrgRAMSize   int        //byte 8, 8KB units
	TvType       TvType     //byte 9 bit 0, byte 10 bits 0+1
	Nes2_0       bool       //byte 7 bits 2,3 == 0x02
	BusConflicts bool       //byte 10 bit 5
}

//if byte7 AND 0x0C == 0x08, and ROM size, taking into account byte 9 doesn't exceed actual filesize, then NES2.0
//if byte7 AND 0x0C == 0x00, and bytes 12-15 == 0 then full iNES
//otherwise, archaic iNES

//NesCart holds the actual data for an NES ROM file and save data
type NesCart struct {
	PrgROM []uint8
	ChrROM []uint8
	SRAM   []uint8
	Mapper *Mapper
	Header NesHeader
}

//MirrorType is the type of graphics memory mirroring that the cartridge is currently configured to use
type MirrorType int

//These specify how the cartridge is wired to mirror addresses in the name table
const (
	HMIRROR MirrorType = iota
	VMIRROR
	FOURSCREEN
)

//MapperType is an enum of constants giving the mapper numbers symbolic names.
type MapperType int

//These are values for the individual mapper numbers, to make them easier to identify in a symbolic way
const (
	NROM        MapperType = iota //247 games
	MMC1                          //680 games
	UxROM                         //269 games
	CNROM                         //155 games
	MMC3                          //599 games
	MMC5                          //24 games
	GAMEDOC                       //SKIP
	AxROM                         //75 games
	FFEDISK                       //See GAME_DOC docs for more details SKIP
	MMC2                          //Punch-Out
	MMC4                          //3 games SKIP
	COLORDREAMS                   //31 games
	MMC3A                         //MMC3 variant (probably implement with same code as MMC3)
	CPROM
	GOUDERHUANG //See MMC3 and VRC2 mappers; apparently does both?
	K1029
	BANDAIFCG
	FFSMDISK
	JALECOSS88006
	NAMCO //20 games
	FDS
	VRC4A
	VRC2A
	VRC2B
	VRC6A //1 game
	VRC2C
	VRC6B //2 games
	VRC4PIRATE
	ACTION53
	HOMEBREW
	UNROM512
	NSF
)

//SysType describes the type of NES-related machine that the ROM is for
type SysType int

//These are the NES-related systems
const (
	VSUnisystem SysType = iota
	Playchoice10
	NES
	Extended
)

//TvType specifies which video system the game was written for
type TvType int

//NTSC, PAL, and Dual system are the three categories of TV system support that NES ROMs can have
const (
	NTSC TvType = iota
	PAL
	DUAL
)

//Load reads the provided ROM file, and puts it into an NesCart struct for access by the emulator
func (cart *NesCart) Load(filename *string) bool {
	var contents []byte
	var err error
	if strings.HasSuffix(*filename, ".zip") {
		zipFile, err := zip.OpenReader(*filename)
		if err != nil {
			panic(err)
		}
		defer zipFile.Close()
		firstFile, err := zipFile.File[0].Open()
		contents, err = ioutil.ReadAll(firstFile)
		if err != nil {
			panic(err)
		}
	} else {
		contents, err = ioutil.ReadFile(*filename)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("File length: ", len(contents))
	header := contents[:16]
	if !Equal(header[:4], []byte{'N', 'E', 'S', 0x1a}) {
		fmt.Println("Header doesn't look correct")
		return false
	}
	for _, c := range header {
		fmt.Printf("%02x ", c)
	}
	return true
}

//Read reads a byte from the given address in the cartridge address-space
func (cart *NesCart) Read(addr uint16, cycle uint64) uint8 {
	return 0
}

//Write handles bytes written onto the cartridge bus
func (cart *NesCart) Write(addr uint16, val uint8, cycle uint64) {
}
