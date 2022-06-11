package mappers

type Mmc1Mapper struct {
	prgROM uint
	chrROM uint
}

func (m Mmc1Mapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m Mmc1Mapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m *Mmc1Mapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *Mmc1Mapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}
