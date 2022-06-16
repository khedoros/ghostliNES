package mappers

type ColorDreamsMapper struct {
	prgROM uint
	chrROM uint
}

func (m ColorDreamsMapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m ColorDreamsMapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m *ColorDreamsMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *ColorDreamsMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}