package mappers

type Mmc2Mapper struct {
	prgROM uint
	chrROM uint
}

func (m Mmc2Mapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m Mmc2Mapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m *Mmc2Mapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *Mmc2Mapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}
