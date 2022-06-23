package nesppu

import (
	"fmt"

	nescart "github.com/khedoros/ghostliNES/NesCart"
)

type ctrl1Reg uint8
type ctrl2Reg uint8
type statReg uint8

//An NesPpu represents an NES's Picture Processing Unit
type NesPpu struct {
	frameCycle        uint
	cyclesPerFrame    uint
	linesPerFrame     uint
	cyclesBeforeVSync uint
	cpuPpuClockFactor float32
	handledNmi        bool
	clearedStatus     bool
	cart              *nescart.NesCart
	vram              [2048]uint8
	sprRam            [256]uint8
	palRam            [32]uint8
	palette           [64]Color
	vramPtr           uint16
	vramPtrShadow     uint16
	vramLatch         bool
	fineX             uint8
	control1          ctrl1Reg // $2000 (W)
	control2          ctrl2Reg // $2001 (W)
	status            statReg  // $2002 (R)
	sprAddr           uint8    // $2003 (W)
	sprData           uint8    // $2004 (RW)
	vramAddr1         uint8    // $2005 (W)
	vramAddr2         uint8    // $2006 (W)
	vramData          uint8    // $2007 (RW)
	sprDMA            uint8    // $4014 (W)
}

const (
	// Common values
	cyclesPerLine   = 341
	preRenderLines  = 1
	postRenderLines = 1

	// PAL values
	masterClockPal       = 26601712
	cpuClockDividerPal   = 16
	ppuClockDividerPal   = 5
	visibleLinesPal      = 240
	vblankLinesPal       = 70
	linesPerFramePal     = 312
	cyclesBeforeVSyncPal = (preRenderLines + visibleLinesPal + postRenderLines) * cyclesPerLine
	cyclesPerFramePal    = (preRenderLines + visibleLinesPal + postRenderLines + vblankLinesPal) * cyclesPerLine

	// NTSC values
	masterClockNtsc       = 21477272
	cpuClockDividerNtsc   = 12
	ppuClockDividerNtsc   = 4
	visibleLinesNtsc      = 240
	vblankLinesNtsc       = 20
	linesPerFrameNtsc     = 262
	cyclesBeforeVSyncNtsc = (preRenderLines + visibleLinesNtsc + postRenderLines) * cyclesPerLine
	cyclesPerFrameNtsc    = (preRenderLines + visibleLinesNtsc + postRenderLines + vblankLinesNtsc) * cyclesPerLine

	// Registers
	ppuControl1  = 0x2000
	ppuControl2  = 0x2001
	ppuStatus    = 0x2002
	ppuSprAddr   = 0x2003
	ppuSprData   = 0x2004
	ppuSprDma    = 0x4014
	ppuVramAddr1 = 0x2005
	ppuVramAddr2 = 0x2006
	ppuVramData  = 0x2007

	// Memory map
	cartRomBase = 0x0
	cartRomSize = 0x2000
	ppuVramBase = 0x2000
	ppuVramSize = 0x1000
	ppuPalBase  = 0x3f00
	ppuPalSize  = 0x20

	bitFlip16 = uint16(65535)
)

func (this *NesPpu) New(mem *nescart.NesCart, region string) {
	this.cart = mem
	if region == "ntsc" {
		this.cyclesPerFrame = cyclesPerFrameNtsc
		this.cyclesBeforeVSync = cyclesBeforeVSyncNtsc
		this.linesPerFrame = linesPerFrameNtsc
		this.cpuPpuClockFactor = cpuClockDividerNtsc / ppuClockDividerNtsc
	} else {
		this.cyclesPerFrame = cyclesPerFramePal
		this.cyclesBeforeVSync = cyclesBeforeVSyncPal
		this.linesPerFrame = linesPerFramePal
		this.cpuPpuClockFactor = cpuClockDividerPal / ppuClockDividerPal
	}
	this.palette = [64]Color{
		{0x6a, 0x6d, 0x6a}, {0x00, 0x13, 0x80}, {0x1e, 0x00, 0x8a}, {0x39, 0x00, 0x7a},
		{0x55, 0x00, 0x56}, {0x5a, 0x00, 0x18}, {0x4f, 0x10, 0x00}, {0x3d, 0x1c, 0x00},
		{0x25, 0x32, 0x00}, {0x00, 0x3d, 0x00}, {0x00, 0x40, 0x00}, {0x00, 0x39, 0x24},
		{0x00, 0x2e, 0x55}, {0x00, 0x00, 0x00}, {0x00, 0x00, 0x00}, {0x00, 0x00, 0x00},
		{0xb9, 0xbc, 0xb9}, {0x18, 0x50, 0xc7}, {0x4b, 0x30, 0xe3}, {0x73, 0x22, 0xd6},
		{0x95, 0x1f, 0xa9}, {0x9d, 0x28, 0x5c}, {0x98, 0x37, 0x00}, {0x7f, 0x4c, 0x00},
		{0x5e, 0x64, 0x00}, {0x22, 0x77, 0x00}, {0x02, 0x7e, 0x02}, {0x00, 0x76, 0x45},
		{0x00, 0x6e, 0x8a}, {0x00, 0x00, 0x00}, {0x00, 0x00, 0x00}, {0x00, 0x00, 0x00},
		{0xff, 0xff, 0xff}, {0x68, 0xa6, 0xff}, {0x8c, 0x9c, 0xff}, {0xb5, 0x86, 0xff},
		{0xd9, 0x75, 0xfd}, {0xe3, 0x77, 0xb9}, {0xe5, 0x8d, 0x68}, {0xd4, 0x9d, 0x29},
		{0xb3, 0xaf, 0x0c}, {0x7b, 0xc2, 0x11}, {0x55, 0xca, 0x47}, {0x46, 0xcb, 0x81},
		{0x47, 0xc1, 0xc5}, {0x4a, 0x4d, 0x4a}, {0x00, 0x00, 0x00}, {0x00, 0x00, 0x00},
		{0xff, 0xff, 0xff}, {0xcc, 0xea, 0xff}, {0xdd, 0xde, 0xff}, {0xec, 0xda, 0xff},
		{0xf8, 0xd7, 0xfe}, {0xfc, 0xd6, 0xf5}, {0xfd, 0xdb, 0xcf}, {0xf9, 0xe7, 0xb5},
		{0xf1, 0xf0, 0xaa}, {0xda, 0xfa, 0xa9}, {0xc9, 0xff, 0xbc}, {0xc3, 0xfb, 0xd7},
		{0xc4, 0xf6, 0xf6}, {0xbe, 0xc1, 0xbe}, {0x00, 0x00, 0x00}, {0x00, 0x00, 0x00},
	}
}

func (this *NesPpu) IsNmi(cycles uint64) bool {
	if (this.control1)&0x80 > 0 && !this.handledNmi { // If NMI is enabled and haven't handled NMI yet
		cycles = uint64(float32(cycles) * this.cpuPpuClockFactor)
		cycles %= uint64(this.cyclesPerFrame)
		if cycles > uint64(this.cyclesBeforeVSync) {
			this.handledNmi = true
			return true
		}
	}
	return false
}

func (this *NesPpu) Run(cycles int64) bool {
	this.frameCycle += uint(this.cpuPpuClockFactor * float32(cycles))
	if this.frameCycle >= this.cyclesPerFrame {
		this.frameCycle -= this.cyclesPerFrame
		this.handledNmi = false
		this.clearedStatus = false
		return true
	}
	return false
}

// 0-1FFF: CRAM/CROM in cartridge
// 2000-2FFF: VRAM in PPU
// 3000-3F00: Mirror of VRAM
// 3F00-3F1F: palette memory
// 3F00-3FFF: Mirror of palette memory
// 4000-FFFF: Mirror of 0-3FFF

// CPU interface to read from externally-accessible registers
func (this *NesPpu) Read(addr uint16, cycle uint64) uint8 {
	frameCycle := uint64(float64(this.cpuPpuClockFactor)*float64(cycle)) % uint64(this.cyclesPerFrame)
	fmt.Printf("Read PPU %04x at frame cycle %d", addr, frameCycle)

	switch addr {
	case ppuStatus:
		status := uint8(0x0)
		if frameCycle > uint64(20*cyclesPerLine) { // TODO: actually calculate Sprite0 collision, rather than just setting it at the beginning of line 20
			status |= 0x40
		}
		if frameCycle > uint64(this.cyclesBeforeVSync) && !this.clearedStatus {
			status |= 0x80
			this.clearedStatus = true
		}
		this.vramLatch = false
		fmt.Printf(", returning %02x\n", status)
		return status
	case ppuSprData:
		fmt.Printf(", returning %02x\n", 0)
		return 0
	case ppuVramData:
		val := this.read(this.vramPtr, cycle)
		fmt.Printf(", returning %02x\n", val)
		if this.control1&4 == 4 {
			this.vramPtr += 32
		} else {
			this.vramPtr++
		}
		return val
	default:
		fmt.Printf(", returning %02x\n", 0)
		return 0
	}
}

// CPU interface to write to externally-accessible registers
func (this *NesPpu) Write(addr uint16, val uint8, cycle uint64) {
	fmt.Printf("Write PPU %04x = %02x\n", addr, val)
	switch addr {
	case ppuControl1:
		this.control1 = ctrl1Reg(val)

		// Put bits 0+1 into bits 10+11 of the vramPtrShadow
		clearBits := bitFlip16 ^ (uint16(3 << 10))
		setBits := (uint16(val & 3)) << 10
		this.vramPtrShadow &= clearBits
		this.vramPtrShadow |= setBits

	case ppuControl2:
		this.control2 = ctrl2Reg(val)
	case ppuSprAddr:
		this.sprAddr = val
	case ppuSprData:
		this.sprData = val
		this.sprRam[this.sprAddr] = val
		this.sprAddr++
	case ppuVramAddr1: // Scrolling register
		if this.vramLatch { // Set y scroll
			clearBits := bitFlip16 ^ (0b111001111100000)
			fineY := uint16(val&0b111) << 12
			coarseY := uint16(val&0b11111000) << 2
			this.vramPtrShadow &= clearBits
			this.vramPtrShadow |= (fineY | coarseY)
		} else { // set x scroll
			clearBits := bitFlip16 ^ (0b11111)
			coarseX := (uint16(val & 0b11111000)) >> 3
			this.vramPtrShadow &= clearBits
			this.vramPtrShadow |= coarseX
			this.fineX = val & 0b111
		}
		this.vramLatch = !this.vramLatch
	case ppuVramAddr2: // VRAM access register
		if this.vramLatch { // set lower 8 bits
			this.vramPtrShadow &= 0xff00
			this.vramPtrShadow |= uint16(0b00111111 & val)
			this.vramPtr = this.vramPtrShadow
		} else {
			this.vramPtrShadow &= 0x00ff
			this.vramPtrShadow |= uint16(val&0b00111111) << 8
		}
		this.vramLatch = !this.vramLatch
	case ppuVramData:
		this.write(this.vramPtr, val, cycle)
		if this.control1&4 == 4 {
			this.vramPtr += 32
		} else {
			this.vramPtr++
		}
		this.vramPtr &= 0x3fff
	case ppuSprDma: // Probably actually do this via 256 writes, so this access represents 1 write
		this.sprRam[this.sprAddr] = val
		this.sprAddr++
	}

}

// Internal PPU memory read
func (this *NesPpu) read(addr uint16, cycle uint64) uint8 {
	if addr < 0x2000 {
		return this.cart.ReadPpu(addr, cycle)
	}
	return 0
}

// Internal PPU memory write
func (this *NesPpu) write(addr uint16, val uint8, cycle uint64) {
	if addr < ppuVramBase { // Write to CRAM/CROM
		this.cart.WritePpu(addr, val, cycle)
	} else if addr >= 0x3f00 && addr < 0x4000 { // Write to palette RAM
		this.palRam[addr&0x1f] = val
	} else if addr < 0x4000 { // Write to VRAM
		// TODO: Mirroring. For now, just keep the writes within physical VRAM
		this.vram[addr&0x7ff] = val
	}
}

type Color struct {
	R, G, B uint8
}

func (this *NesPpu) Render() *[]Color {
	c := []Color{}
	for i := 0; i < 256*240; i++ {
		if i%256 < 32 && i/256 < 30 {
			val := this.vram[(i/256)*32+i%256]
			c = append(c, Color{val, val, val})
		} else {
			c = append(c, Color{0, 0, 0})
		}
	}
	return &c
}
