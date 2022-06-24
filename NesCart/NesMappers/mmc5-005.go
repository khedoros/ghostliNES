package mappers

type Mmc5Mapper struct {
	prgROM uint
	chrROM uint
}

func (m Mmc5Mapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m Mmc5Mapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m *Mmc5Mapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m *Mmc5Mapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
}

func (m *Mmc5Mapper) GetMirror() MirrorType {
	return HARDWIRED
}
