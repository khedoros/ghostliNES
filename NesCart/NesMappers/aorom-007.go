package mappers

type AoromMapper struct {
	prgROM uint
	chrROM uint
}

func (m AoromMapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m AoromMapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m *AoromMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *AoromMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}

func (m *AoromMapper) GetMirror() MirrorType {
	return HARDWIRED
}
