package NesMem

type Mapper interface {
	CpuWrite(cycle uint32, addr uint16, data uint8)
	CpuRead(cycle uint32, addr uint16) uint8
	GetNextChangeCycle() uint32
}
