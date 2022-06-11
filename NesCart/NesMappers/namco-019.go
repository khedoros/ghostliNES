package mappers

type NamcoMapper struct {
	prgROM uint
	chrROM uint
}

func (m NamcoMapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m NamcoMapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m *NamcoMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *NamcoMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}
