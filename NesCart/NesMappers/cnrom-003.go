package mappers

type CnromMapper struct {
	prgROM uint
	chrROM uint

	chrPage uint
}

func (m CnromMapper) MapCpu(addr uint16, cycle uint64) uint {
	return (uint(addr) - 0x8000) % m.prgROM
}

func (m CnromMapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr) + m.chrPage * 8192
}

func (m CnromMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {
	if addr > 0x8000 {
		m.chrPage = uint(val)
	}
}

func (m *CnromMapper) New(prg, chr uint) {
	m.prgROM, m.chrROM = prg, chr
    pages := chr / 8192
	m.chrPage = pages - 1
}

func (m *CnromMapper) GetMirror() MirrorType {
	return HARDWIRED
}
