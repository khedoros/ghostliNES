package mappers

type UnromMapper struct {
	prgROM  uint
	chrROM  uint
	prgPage uint
}

func (m UnromMapper) MapCpu(addr uint16, cycle uint64) uint {
	if addr >= 0xc000 {
		return uint(addr-0xc000) + m.prgROM - 16384
	}
	return uint(addr) - 0x8000 + m.prgPage*16384
}

func (m UnromMapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m *UnromMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {
	if addr > 0x8000 && addr < 0xffff {
		m.prgPage = uint(val) % (m.prgROM / 16384)
	}
}

func (m *UnromMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM, m.prgPage = prg, chr, 0
}

func (m *UnromMapper) GetMirror() MirrorType {
	return HARDWIRED
}
