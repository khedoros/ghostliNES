package NesCart



type Mapper interface {
	func (m *Mapper) read(addr uint16) uint8
	func (m *Mapper) write(cycle uint32, addr uint16, data uint8)
	func (m *Mapper) ppu_read(cycle uint32, addr uint16)
	func (m *Mapper) ppu_write(cycle uint32, addr uint16, data uint8)
	func (m *Mapper) new()
}

type NesCart struct {
	prg_rom []uint8
	chr_rom []uint8
	sram []uint8
	mapper *Mapper
}

func (cart *NesCart) load(filename string) bool {

}
