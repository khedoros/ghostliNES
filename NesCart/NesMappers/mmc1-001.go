package mappers

import "fmt"

type Mmc1Mapper struct {
	prgROM    uint
	chrROM    uint
	prgLoPage uint // Low bits 1-3 of reg3, if r0b3 is 0, low 4 bits of reg3 if r0b2 = 1 and r0b3 = 1
	prgHiPage uint // Low bits 1-3 of reg3, +1, if r0b3 is 0, low 4 bits of reg3 if r0b2 = 0 and r0b3 = 1
	chrLoPage uint // Low 4 bits of reg1
	chrHiPage uint // Low 4 bits of reg2, or reg1 if bit4 of reg0 is 0

	buffer   uint
	bitCount uint

	mirror       bool // Reg0 Bit0 0=H, 1=V
	normalMirror bool // Reg0 Bit1 0=one-screen, 1=normal
	swapPromLo   bool // Reg0 Bit2 0=swap 0xc000, 1=swap 0x8000
	smallPrgSwap bool // Reg0 Bit3 0=32KiB PRG-ROM swap, 1=16KiB swap
	smallChrSwap bool // Reg0 Bit4 0=8KiB CHR-ROM swap, 1=4KiB swap
	sramDisabled bool // Reg3 Bit4

}

func (m Mmc1Mapper) MapCpu(addr uint16, cycle uint64) uint {
	if addr >= 0xc000 {
		return uint(addr-0xc000) + m.prgHiPage*16384
	} else {
		return uint(addr) - 0x8000 + m.prgLoPage*16384
	}
}

func (m Mmc1Mapper) MapPpu(addr uint16, cycle uint64) uint {
	if addr <= 0x1000 {
		return m.chrLoPage*4096 + uint(addr)
	} else {
		return m.chrHiPage*4096 + uint(addr&0xfff)
	}
}

func (m *Mmc1Mapper) WriteCpu(addr uint16, val uint8, cycle uint64) {
	fmt.Printf("MMC1 %04x = %02x\n", addr, val)
	if (val & 0x80) == 0x80 {
		m.buffer = 0
		m.bitCount = 0
		return
	}

	m.buffer |= uint((val & 0x01) << uint8(m.bitCount))
	m.bitCount++
	if m.bitCount == 5 {
		m.apply(addr, cycle)
		m.bitCount = 0
		m.buffer = 0
	}
}

func (m *Mmc1Mapper) apply(addr uint16, cycle uint64) {
	if addr >= 0x8000 && addr <= 0x9fff { // Reg0: RxxCFHPM
		m.mirror = ((m.buffer & 1) == 1) // TODO: Get bits 0 and 1 available to the PPU somehow
		m.normalMirror = ((m.buffer & 2) == 2)
		m.swapPromLo = ((m.buffer & 4) == 4)
		m.smallPrgSwap = ((m.buffer & 8) == 8)
		m.smallChrSwap = ((m.buffer & 16) == 16)
		fmt.Printf("MMC1 Reg0: %02x\n", m.buffer)
	} else if addr >= 0xa000 && addr <= 0xbfff { // Reg1: RxxPCCCC
		if m.smallChrSwap {
			m.chrLoPage = uint(m.buffer & 0xf)
		} else {
			m.chrLoPage = uint((m.buffer & 0xe) >> 1) // TODO: Verify that this *is* actually supposed to ignore the low bit in this case
			m.chrHiPage = m.chrLoPage + 1
		}
		fmt.Printf("MMC1 Reg1: %02x\n", m.buffer)
	} else if addr >= 0xc000 && addr <= 0xdfff { // Reg2: RxxPCCCC
		if m.smallChrSwap {
			m.chrHiPage = uint(m.buffer & 0x0f)
		}
		fmt.Printf("MMC1 Reg2: %02x\n", m.buffer)
	} else if addr >= 0xe000 && addr <= 0xffff { // Reg3: RxxPCCCC
		if m.smallPrgSwap && m.swapPromLo { // Swap region at 0x8000
			m.prgLoPage = uint(m.buffer & 0xf)
		} else if m.smallPrgSwap { // Swap region at 0xc000
			m.prgHiPage = uint(m.buffer & 0xf)
		} else { // Swap both 0x8000 and 0xc000 regions
			m.prgLoPage = uint((m.buffer & 0xe) >> 1)
			m.prgHiPage = m.prgLoPage + 1
		}

		m.sramDisabled = (m.buffer & 0x10) == 0x10

		fmt.Printf("MMC1 Reg3: %02x\n", m.buffer)
	}
}

func (m *Mmc1Mapper) New(prg, chr uint) {
	m.normalMirror = true
	m.smallChrSwap = true
	m.smallPrgSwap = true
	m.swapPromLo = true
	m.sramDisabled = true
	m.buffer = 0
	m.bitCount = 0
	m.prgROM, m.chrROM = prg, chr
	prgPages := m.prgROM / 16384
	m.prgLoPage = 0
	m.prgHiPage = prgPages - 1
	m.chrLoPage = 0
	m.chrHiPage = 1
}

func (m *Mmc1Mapper) GetMirror() MirrorType {
	if !m.normalMirror {
		return SINGLESCREEN
	} else if m.mirror {
		return VMIRROR
	} else {
		return HMIRROR
	}
}
