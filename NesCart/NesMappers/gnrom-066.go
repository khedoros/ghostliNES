package mappers

type GnromMapper struct {
	prgROM uint
	chrROM uint

	prgPage uint
	chrPage uint
}

func (m GnromMapper) MapCpu(addr uint16, cycle uint64) uint {
	return ((uint(addr) - 0x8000) + m.prgPage*32768) % m.prgROM
}

func (m GnromMapper) MapPpu(addr uint16, cycle uint64) uint {
	return (uint(addr) + m.chrPage*8192) % m.chrROM
}

func (m *GnromMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {
	m.chrPage = uint(val & 0x3)
	m.prgPage = uint(val&0x30) >> 4
}

func (m *GnromMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
	m.prgPage, m.chrPage = 0, 0
}

func (m *GnromMapper) GetMirror() MirrorType {
	return HARDWIRED
}
