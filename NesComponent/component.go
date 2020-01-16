package nescomponent

//NesComponent is a hardware component connected to the address and data bus, capable of being written to and read from.
type NesComponent interface {
	Read(addr uint16, cycle uint64) uint8
	Write(addr uint16, val uint8, cycle uint64)
}
