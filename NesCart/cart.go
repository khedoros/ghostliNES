package nescart

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"strings"

	mappers "github.com/khedoros/ghostliNES/NesCart/NesMappers"
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

//NesHeader defines the 16-byte structure that makes up the iNES header in most NES ROMs
type NesHeader struct {
	PrgSize      int        //byte 4, 16KB units
	ChrSize      int        //byte 5, 8KB units
	ChrRAM       bool       //byte 5 == 0
	MapperNum    MapperType //high-4 of byte 6 (low)  and high-4 of byte 7 (high)
	Mirroring    MirrorType //byte 6 bit 0 (0 == H-Mirror, 1 == V-Mirror)
	Battery      bool       //byte 6 bit 1
	Trainer      bool       //byte 6 bit 2
	FourScreen   bool       //byte 6 bit 3 (0 == no extra RAM, 1 == extra RAM?)
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
	prgROM []uint8
	chrROM []uint8
	sRAM   []uint8
	mapper mappers.Mapper
	header NesHeader
}

//MirrorType is the type of graphics memory mirroring that the cartridge is currently configured to use
type MirrorType int

//These specify how the cartridge is wired to mirror addresses in the name table
const (
	HMIRROR MirrorType = iota
	VMIRROR
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
	RAMBO1 = iota + 32
	H3001
	GNROM
	UNSUPPORTED
)

func (cart *NesCart) getMapper() mappers.Mapper {
	var mapper mappers.Mapper = nil
	switch cart.header.MapperNum {
	case NROM:
		mapper = &mappers.NromMapper{}
		fmt.Println("NROM Mapper")
	case MMC1:
		mapper = &mappers.Mmc1Mapper{}
		fmt.Println("MMC1 Mapper")
	case UxROM:
		mapper = &mappers.UnromMapper{}
		fmt.Println("UNROM Mapper")
	case CNROM:
		mapper = &mappers.CnromMapper{}
		fmt.Println("CNROM Mapper")
	case MMC3:
		mapper = &mappers.Mmc3Mapper{}
		fmt.Println("MMC3 Mapper")
	case MMC5:
		mapper = &mappers.Mmc5Mapper{}
		fmt.Println("MMC5 Mapper")
	case AxROM:
		mapper = &mappers.AoromMapper{}
		fmt.Println("AOROM Mapper")
	case MMC2:
		mapper = &mappers.Mmc2Mapper{}
		fmt.Println("MMC2 Mapper")
	case COLORDREAMS:
		mapper = &mappers.ColorDreamsMapper{}
		fmt.Println("Color Dreams Mapper")
	case NAMCO:
		mapper = &mappers.NamcoMapper{}
		fmt.Println("Namco Mapper")
	case RAMBO1:
		mapper = &mappers.Rambo1Mapper{}
		fmt.Println("Rambo-1 Mapper")
	case GNROM:
		mapper = &mappers.GnromMapper{}
		fmt.Println("GNROM Mapper")
	}
	return mapper
}

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
	contents = contents[16:]
	if !Equal(header[:4], []byte{'N', 'E', 'S', 0x1a}) {
		fmt.Println("Header doesn't look correct")
		return false
	}
	for _, c := range header {
		fmt.Printf("%02x ", c)
	}
	fmt.Println("")
	cart.header.PrgSize = int(header[4])
	cart.prgROM = contents[:16384*cart.header.PrgSize]
	contents = contents[16384*cart.header.PrgSize:]
	cart.header.ChrSize = int(header[5])
	if cart.header.ChrSize == 0 {
		cart.header.ChrRAM = true
		cart.chrROM = make([]uint8, 8192)
	} else {
		if len(contents) != cart.header.ChrSize*8192 {
			panic(fmt.Sprintf("Expected CHR ROM size of %v, saw %v bytes remaining", cart.header.ChrSize*8192, len(contents)))
		}
		cart.chrROM = contents
	}
	cart.header.MapperNum = MapperType(header[6]>>4 | (header[7] & 0xf0))
	cart.header.Mirroring = MirrorType(header[6] & 1)
	cart.header.Battery = (header[6] & 2) == 2
	cart.header.Trainer = (header[6] & 4) == 4
	cart.header.FourScreen = (header[6] & 8) == 8

	// Info beyond this point is "iffy"
	cart.header.SysType = SysType(header[7] & 3)
	cart.header.PrgRAMSize = int(header[8])
	cart.header.PrgRAM = header[10]&16 == 16

	// Mapper number from bytes 6 and 7 needs to be accurate
	cart.mapper = cart.getMapper()
	if cart.mapper == nil {
		fmt.Printf("Unknown mapper number %d\n", cart.header.MapperNum)
		return false
	}
	cart.mapper.New(uint(len(cart.prgROM)), uint(len(cart.chrROM)))

	return true
}

//Read reads a byte from the given address in the cartridge address-space
func (cart *NesCart) Read(addr uint16, cycle uint64) uint8 {
	return cart.prgROM[cart.mapper.MapCpu(addr, cycle)]
}

//Write handles bytes written onto the cartridge bus by the CPU
func (cart *NesCart) Write(addr uint16, val uint8, cycle uint64) {
	cart.mapper.WriteCpu(addr, val, cycle)
}

//ReadPpu reads a byte from the given address in the cartridge address-space
func (cart *NesCart) ReadPpu(addr uint16, cycle uint64) uint8 {
	return cart.chrROM[cart.mapper.MapPpu(addr, cycle)]
}

//WritePpu handles bytes written onto the cartridge bus by the PPU
func (cart *NesCart) WritePpu(addr uint16, val uint8, cycle uint64) {
	if cart.header.ChrRAM {
		cart.chrROM[cart.mapper.MapPpu(addr, cycle)] = val
	}
}
