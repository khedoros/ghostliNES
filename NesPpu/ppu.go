package nesppu

import (
	"fmt"

	nescart "github.com/khedoros/ghostliNES/NesCart"
)

type ctrl1Reg uint8
type ctrl2Reg uint8
type statReg uint8

type ppuWrite struct {
	addr       uint16
	val        uint8
	frame      uint64
	frameCycle uint64
}

// An NesPpu represents an NES's Picture Processing Unit
type NesPpu struct {
	buffer            [256 * 240]Color
	frameCycle        uint
	cyclesPerFrame    uint
	linesPerFrame     uint
	cyclesBeforeVSync uint
	cpuPpuClockFactor float64
	handledNmi        bool
	clearedStatus     bool

	writeQueue     [3000]ppuWrite
	writeQueueCnt  uint
	processedUntil uint

	cart          *nescart.NesCart
	vram          [2048]uint8
	sprRam        [256]uint8
	palRam        [32]uint8
	palette       [64]Color
	vramPtr       uint16
	vramPtrShadow uint16
	vramLatch     bool
	fineX         uint8
	control1      ctrl1Reg // $2000 (W)
	control2      ctrl2Reg // $2001 (W)
	status        statReg  // $2002 (R)
	sprAddr       uint8    // $2003 (W)
	//sprData       uint8    // $2004 (RW)
	vramAddr1 uint8 // $2005 (W)
	vramAddr2 uint8 // $2006 (W)
	vramData  uint8 // $2007 (RW)
	sprDMA    uint8 // $4014 (W)

	tile1_byte1 uint8
	tile1_byte2 uint8
	tile1_attr  uint8
	tile2_byte1 uint8
	tile2_byte2 uint8
	tile2_attr  uint8
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
		_, cycle := this.cpuToPpuCycle(cycles)
		if uint(cycle) > this.cyclesBeforeVSync {
			this.handledNmi = true
			return true
		}
	}
	return false
}

func (this *NesPpu) Run(cpuCyclesToRunFor uint) bool {
	frames, ppuCycles := this.cpuToPpuCycle(uint64(cpuCyclesToRunFor))
	if frames > 0 {
		panic(fmt.Sprintf("Told to run for %d CPU cycles; more than the number of PPU cycles per frame (%d)", cpuCyclesToRunFor, this.cyclesPerFrame))
	}
	for ; ppuCycles > 0; ppuCycles-- {
		this.apply(this.frameCycle)
		this.frameCycle++
	}
	//this.frameCycle += ppuCycles
	//this.apply(this.frameCycle)
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
	_, frameCycle := this.cpuToPpuCycle(cycle)
	if addr != 0x2002 {
		fmt.Printf("Read PPU %04x at frame cycle %d", addr, frameCycle)
	}

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
		//fmt.Printf(", returning %02x\n", status)
		return status
	case ppuSprData:
		//fmt.Printf(", returning %02x\n", 0)
		return 0
	case ppuVramData:
		val := this.read(this.vramPtr, cycle)
		fmt.Printf(", returning %02x\n", val)
		this.vramPtr++
		return val
	default:
		fmt.Printf(", returning %02x\n", 0)
		return 0
	}
}

func (this *NesPpu) cpuToPpuCycle(cycle uint64) (frame, frameCycles uint64) {
	cycles := uint64(float64(cycle) * this.cpuPpuClockFactor)
	return uint64(cycles / uint64(this.cyclesPerFrame)), (cycles % uint64(this.cyclesPerFrame))
}

// CPU interface to write to externally-accessible registers
func (this *NesPpu) Write(addr uint16, val uint8, cycle uint64) {
	if addr != 0x4014 {
		fmt.Printf("Write PPU %04x = %02x (write queue %d)\n", addr, val, this.writeQueueCnt)
	}
	if this.writeQueueCnt < uint(len(this.writeQueue)) {
		frame, frameCycle := this.cpuToPpuCycle(cycle)
		this.writeQueue[this.writeQueueCnt] = ppuWrite{addr, val, frame, frameCycle}
		this.writeQueueCnt++
	} else {
		panic(fmt.Sprintf("Exceeded size of writeQueue (%d items)", len(this.writeQueue)))
	}
}

func (this *NesPpu) apply(upToCycle uint) {
	itemIdx := this.processedUntil
	prevFrameCycle := this.writeQueue[itemIdx].frameCycle
	for ; itemIdx < this.writeQueueCnt && uint(this.writeQueue[itemIdx].frameCycle) <= upToCycle; itemIdx++ {
		curItem := this.writeQueue[itemIdx]
		addr := curItem.addr
		val := curItem.val
		frameCycle := curItem.frameCycle
		if frameCycle < prevFrameCycle {
			// panic(fmt.Sprintf("item %d marked as cycle %d with previous item being at cycle %d", itemIdx, frameCycle, prevFrameCycle))
		}
		prevFrameCycle = curItem.frameCycle

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
			this.sprRam[this.sprAddr] = val
			this.sprAddr++
		case ppuVramAddr1: // Scrolling register
			if this.vramLatch { // Set y scroll
				clearBits := bitFlip16 ^ (0b1111001111100000)
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
				this.vramPtrShadow |= uint16(val)
				this.vramPtr = this.vramPtrShadow
			} else {
				this.vramPtrShadow &= 0x00ff
				this.vramPtrShadow |= (uint16(val&0b00111111) << 8)
			}
			this.vramLatch = !this.vramLatch
		case ppuVramData:
			this.write(this.vramPtr, val, uint64(frameCycle))
			if this.control1&4 == 4 {
				this.vramPtr += 32
			} else {
				this.vramPtr++
			}
			this.vramPtr &= 0x3fff
		case ppuSprDma: // Probably actually do this via 256 writes, so this access represents 1 write
			//fmt.Printf("DMA[%02x] = %02x\n", this.sprAddr, val)
			this.sprRam[this.sprAddr] = val
			this.sprAddr++
		}
	}

	// Mark queue as empty when we get to the end
	// or save last-processed index to pick up later
	if itemIdx == this.writeQueueCnt {
		this.processedUntil = 0
		this.writeQueueCnt = 0
	} else {
		this.processedUntil = itemIdx
	}
}

// Internal PPU memory read
func (this *NesPpu) read(addr uint16, cycle uint64) uint8 {
	addr &= 0x3fff
	if addr < 0x2000 { // Read from pattern table
		return this.cart.ReadPpu(addr, cycle)
	} else if addr >= 0x3f00 && addr < 0x4000 { // Read from palette RAM
		return 0
	} else { // Read from name/attrib table
		return this.vram[addr&0x7ff]
	}
}

// Internal PPU memory write
func (this *NesPpu) write(addr uint16, val uint8, cycle uint64) {
	addr &= 0x3fff
	if addr < ppuVramBase { // Write to CRAM/CROM
		this.cart.WritePpu(addr, val, cycle)
	} else if addr >= 0x3f00 { // Write to palette RAM
		val &= 0x3f
		addr &= 0x1f
		element := addr % 4
		this.palRam[addr] = val
		if element == 0 {
			this.palRam[addr^0x10] = val
		}
	} else { // Write to name/attrib table
		// TODO: Mirroring. For now, just keep the writes within physical VRAM
		this.vram[addr&0x7ff] = val
	}
}

type Color struct {
	R, G, B uint8
}

func (this *NesPpu) getTileLine(base uint16, tileNum, fineY uint8) [8]uint8 {
	line := [8]uint8{}
	addr := base + uint16(tileNum)*16 + uint16(fineY)
	if fineY >= 8 {
		addr += 8
	}
	dat1 := this.cart.ReadPpu(addr, 0)
	dat2 := this.cart.ReadPpu(addr+8, 0)
	mask := uint8(128)
	for pix := 0; pix < 8; pix++ {
		bit1 := (dat1 & mask) >> (7 - pix)
		bit2 := (dat2 & mask) >> (7 - pix)
		idx := bit1 | (bit2 << 1)
		line[pix] = idx
		mask /= 2
	}
	return line
}

func (this *NesPpu) getAttrib(base uint16, coarseX, coarseY uint8) uint8 {
	addr := base + 0x3c0 + uint16(coarseX)/4 + ((uint16(coarseY) / 4) * 8)
	attribByte := this.read(addr, 0)
	shift := 2 * ((coarseX%4)/2 + 2*((coarseY%4)/2))
	return (attribByte & (3 << shift)) >> shift
}

func (this *NesPpu) Print() {
	for y := 0; y < 30; y++ {
		for x := 0; x < 32; x++ {
			fmt.Printf("%02x ", this.vram[y*32+x])
		}
		fmt.Printf("\n")
	}
}

func (this *NesPpu) Render() *[61440]Color {
	//this.apply(this.cyclesPerFrame)
	if this.control2&16 == 16 {
		this.drawSprites(true)
	}
	if this.control2&8 == 8 {
		this.drawBackground()
	}
	if this.control2&16 == 16 {
		this.drawSprites(false)
	}
	return &this.buffer
}

func (this *NesPpu) drawBackground() {
	pix := 0
	bgBase := 256 * uint16(this.control1&0x10)
	for coarseY := uint(0); coarseY < 30; coarseY++ {
		for fineY := uint8(0); fineY < 8; fineY++ {
			for coarseX := uint(0); coarseX < 32; coarseX++ {
				tileNum := this.vram[coarseY*32+coarseX+1024]
				tileLine := this.getTileLine(bgBase, tileNum, fineY)
				tileAttrib := this.getAttrib(0x2000, uint8(coarseX), uint8(coarseY)) << 2

				for fineX := 0; fineX < 8; fineX++ {
					this.buffer[pix] = this.palette[this.palRam[tileAttrib|tileLine[fineX]]]
					pix++
				}
			}
		}
	}
}

func (this *NesPpu) drawSprites(lowPriority bool) {
	for spr := 63; spr >= 0; spr-- {
		priority := this.sprRam[spr*4+2]&0x20 == 0x20
		if this.sprRam[spr*4] >= 0xef || priority != lowPriority { // Sprite isn't visible, or doesn't match desired priority
			continue
		} else {
			//fmt.Printf("Render Sprite %02x\n", spr)
		}
		y := this.sprRam[spr*4] + 1
		t := this.sprRam[spr*4+1]
		s := this.sprRam[spr*4+2]
		x := this.sprRam[spr*4+3]
		pal := (s & 3) << 2
		yf := (s & 0x80) == 0x80
		xf := (s & 0x40) == 0x40
		height := uint8(8)
		sprBase := uint16(this.control1&0x08) * 512
		if this.control1&0x20 == 0x20 { // tall sprites
			height = 16
			if t&1 == 1 {
				sprBase = 0x1000
			} else {
				sprBase = 0
			}
			t &= 0xfe
		}
		for line := uint8(0); line < height && line+y < 240; line++ {
			yPix := line
			if yf {
				yPix = height - line - 1
			}

			tileLine := this.getTileLine(sprBase, t, yPix)
			for xFine := uint8(0); xFine < 8 && uint(xFine)+uint(x) < 256; xFine++ {
				xPix := xFine
				if xf {
					xPix = 7 - xFine
				}
				bufferOffset := uint(x+xFine) + uint(y+line)*256
				if tileLine[xPix] != 0 {
					col := this.palette[this.palRam[(0x10+pal|tileLine[xPix])]]
					this.buffer[bufferOffset] = col
				}
			}
		}
	}
}
