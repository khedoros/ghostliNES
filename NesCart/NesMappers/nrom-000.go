package mappers

type NromMapper struct {
	prgROM uint
	chrROM uint
}

func (m NromMapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m NromMapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m NromMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *NromMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}
