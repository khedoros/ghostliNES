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
	cart              *nescart.NesCart
	vram              [2048]uint8
	sprRam            [256]uint8
	palRam            [32]uint8
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
	fmt.Printf("Read PPU %04x\n", addr)
	switch addr {
	case ppuStatus:
		return 0x80
	case ppuSprData:
		return 0
	case ppuVramData:
		return 0
	default:
		return 0
	}
}

// CPU interface to write to externally-accessible registers
func (this *NesPpu) Write(addr uint16, val uint8, cycle uint64) {
	fmt.Printf("Write PPU %04x = %02x\n", addr, val)
	switch addr {
	case ppuControl1:
		this.control1 = ctrl1Reg(val)
	case ppuControl2:
		this.control2 = ctrl2Reg(val)
	case ppuSprAddr:
		this.sprAddr = val
	case ppuSprData:
		this.sprData = val
		this.sprRam[this.sprAddr] = val
		this.sprAddr++
	case ppuVramAddr1: // Scrolling register
		if this.vramLatch {
		} else {
		}
	case ppuVramAddr2: // VRAM access register
		if this.vramLatch { // set lower 8 bits
			this.vramPtrShadow &= 0xff00
			this.vramPtrShadow |= uint16(val)
			this.vramPtr = this.vramPtrShadow
		} else {
			this.vramPtrShadow &= 0x00ff
			this.vramPtrShadow |= uint16(val&0b00111111) << 8
		}
		this.vramLatch = !this.vramLatch
	case ppuVramData:
		this.vram[this.vramPtr] = val
		this.vramPtr++
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
func (this *NesPpu) write(addr uint16, cycle uint64) {}
