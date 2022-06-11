package mappers

//A Mapper handles mapping pages of data in a ROM to their physical addresses, visible to the CPU and PPU
type Mapper interface {
	MapCpu(addr uint16, cycle uint64) uint
	MapPpu(addr uint16, cycle uint64) uint
	WriteCpu(addr uint16, val uint8, cycle uint64)
	New()
}
