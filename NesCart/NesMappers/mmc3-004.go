package mappers

import "fmt"

type Mmc3Mapper struct {
	prgROM uint
	chrROM uint

	chrAddrXor    uint  // either 0 or 0x1000, set by 8000 b 7
	prgLowIsFixed bool  // 1=0x8000-0x9fff is fixed, 0=0xc000-dfff is fixed, 8000 b6
	ctrlCmd       uint8 // value is 0-7, to specify 1 of 8 commands. 8000 b0-2

	mirror       MirrorType // a000 b0, 0=vert, 1=horiz
	sRamEnable   bool       // a001 b7
	sRamReadOnly bool       // a001 b6

	lowPrgPage uint // 0x8000-0x9fff
	midPrgPage uint // 0xa000-0xbfff
	hiPrgPage  uint // 0xc000-0xdfff

	chrPages [8]uint // 8 pages, 0x400 each

}

func (m Mmc3Mapper) MapCpu(addr uint16, cycle uint64) uint {
	if addr >= 0x8000 && addr <= 0x9fff {
		return (uint(addr) - 0x8000 + m.lowPrgPage*8192) % m.prgROM
	} else if addr >= 0xa000 && addr <= 0xbfff {
		return (uint(addr) - 0xa000 + m.midPrgPage*8192) % m.prgROM
	} else if addr >= 0xc000 && addr <= 0xdfff {
		return (uint(addr) - 0xc000 + m.hiPrgPage*8192) % m.prgROM
	} else if addr >= 0xe000 && addr <= 0xffff {
		return (uint(addr) - 0xe000 + m.prgROM - 8192) % m.prgROM
	}
	return 0
}

func (m Mmc3Mapper) MapPpu(addr uint16, cycle uint64) uint {
	//fmt.Printf("MMC3: addr %04x goes to page %x, currently set to %02x\n", addr, addr>>10, m.chrPages[addr>>10])
	return m.chrPages[addr>>10]*0x400 + uint(addr&0x3ff)

}

func (m *Mmc3Mapper) WriteCpu(addr uint16, val uint8, cycle uint64) {
	//fmt.Printf("MMC3: %04x = %02x\n", addr, val)
	addr &= 0xe001
	switch addr {
	case 0x8000: // ctrl1 register
		m.chrAddrXor = uint((val & 0x80) / 0x20) // 0 or 4
		m.prgLowIsFixed = val&0x40 == 0x40
		m.ctrlCmd = val & 0x7

		/*
			000b 0 - Select 2 1K CHR ROM pages at 0000h in PPU space
			001b 1 - Select 2 1K CHR ROM pages at 0800h in PPU space
			010b 2 - Select 1K CHR ROM page at 1000h in PPU space
			011b 3 - Select 1K CHR ROM page at 1400h in PPU space
			100b 4 - Select 1K CHR ROM page at 1800h in PPU space
			101b 5 - Select 1K CHR ROM page at 1C00h in PPU space
			110b 6 - Select 8K PRG ROM page at 8000h or C000h
			111b 7 - Select 8K PRG ROM page at A000h
		*/
	case 0x8001: // bank for ctrl1
		switch m.ctrlCmd {
		case 0:
			m.chrPages[0^m.chrAddrXor] = uint(val)
			m.chrPages[1^m.chrAddrXor] = uint(val + 1)
		case 1:
			m.chrPages[2^m.chrAddrXor] = uint(val)
			m.chrPages[3^m.chrAddrXor] = uint(val + 1)
		case 2:
			m.chrPages[4^m.chrAddrXor] = uint(val)
		case 3:
			m.chrPages[5^m.chrAddrXor] = uint(val)
		case 4:
			m.chrPages[6^m.chrAddrXor] = uint(val)
		case 5:
			m.chrPages[7^m.chrAddrXor] = uint(val)
		case 6:
			if m.prgLowIsFixed {
				m.hiPrgPage = uint(val)
			} else {
				m.lowPrgPage = uint(val)
			}
		case 7:
			m.midPrgPage = uint(val)
		}
	case 0xa000:
		m.mirror = MirrorType(val & 0x1)
	case 0xa001:
		m.sRamEnable = val&0x80 == 0x80
		m.sRamReadOnly = val&0x40 == 0x40
	case 0xc000:
		fmt.Println("MMC3 0xc000 IRQ not implemented")
	case 0xc001:
		fmt.Println("MMC3 0xc001 IRQ not implemented")
	case 0xe000:
		fmt.Println("MMC3 0xe000 IRQ not implemented")
	case 0xe001:
		fmt.Println("MMC3 0xe001 IRQ not implemented")
	}
}

func (m *Mmc3Mapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
	m.lowPrgPage, m.midPrgPage, m.hiPrgPage = 0, 1, (m.prgROM/8192)-2
	m.prgLowIsFixed = false
	m.sRamEnable = false
	m.sRamReadOnly = true
	m.chrPages = [8]uint{1, 1, 1, 1, 1, 1, 1, 1}
}

func (m *Mmc3Mapper) GetMirror() MirrorType {
	return m.mirror
}
