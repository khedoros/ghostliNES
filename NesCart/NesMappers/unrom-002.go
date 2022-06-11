package mappers

type UnromMapper struct {
	prgROM uint
	chrROM uint
}

func (m UnromMapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m UnromMapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m *UnromMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *UnromMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}
