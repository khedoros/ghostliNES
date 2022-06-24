package mappers

//MirrorType is the type of graphics memory mirroring that the cartridge is currently configured to use
type MirrorType int

//These specify how the cartridge is wired to mirror addresses in the name table
const (
	HMIRROR MirrorType = iota
	VMIRROR
	SINGLESCREEN
	FOURSCREEN
	HARDWIRED
)

//A Mapper handles mapping pages of data in a ROM to their physical addresses, visible to the CPU and PPU
type Mapper interface {
	MapCpu(addr uint16, cycle uint64) uint
	MapPpu(addr uint16, cycle uint64) uint
	WriteCpu(addr uint16, val uint8, cycle uint64)
	GetMirror() MirrorType
	New(prgROM, chrROM uint)
}
