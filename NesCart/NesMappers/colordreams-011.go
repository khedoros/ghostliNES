package mappers

type ColorDreamsMapper struct {
	prgROM uint
	chrROM uint

	prgPage uint
	chrPage uint
}

func (m ColorDreamsMapper) MapCpu(addr uint16, cycle uint64) uint {
	return ((uint(addr) - 0x8000) + 32768*m.prgPage) % m.prgROM
}

func (m ColorDreamsMapper) MapPpu(addr uint16, cycle uint64) uint {
	return (uint(addr) + 8192*m.chrPage) % m.chrROM
}

func (m *ColorDreamsMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {
	m.prgPage = uint(val & 0x0f)
	m.chrPage = uint(val&0xf0) >> 4
}

func (m *ColorDreamsMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
	m.prgPage, m.chrPage = 0, 0
}

func (m *ColorDreamsMapper) GetMirror() MirrorType {
	return HARDWIRED
}
