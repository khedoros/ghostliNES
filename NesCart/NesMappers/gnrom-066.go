package mappers

type GnromMapper struct {
	prgROM uint
	chrROM uint
}

func (m GnromMapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m GnromMapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m *GnromMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *GnromMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}
