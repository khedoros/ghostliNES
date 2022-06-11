package mappers

type Rambo1Mapper struct {
	prgROM uint
	chrROM uint
}

func (m Rambo1Mapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m Rambo1Mapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m *Rambo1Mapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *Rambo1Mapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}
