package mappers

type CnromMapper struct {
	prgROM uint
	chrROM uint
}

func (m CnromMapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m CnromMapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m CnromMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *CnromMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}
