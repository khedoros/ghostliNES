package mappers

type NromMapper struct{}

func (m NromMapper) MapCpu(addr uint16, cycle uint64) uint {
	return uint(addr) - 0x8000
}

func (m NromMapper) MapPpu(addr uint16, cycle uint64) uint {
	return uint(addr)
}

func (m NromMapper) WriteCpu(addr uint16, val uint8, cycle uint64) {}

func (m NromMapper) New() {}
