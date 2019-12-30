package NesCart

import (
	"fmt"
	"io/ioutil"
)

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

type Mapper interface {
	Read(addr uint16) uint8
	Write(cycle uint32, addr uint16, data uint8)
	Ppu_read(cycle uint32, addr uint16)
	Ppu_write(cycle uint32, addr uint16, data uint8)
	New()
}

type NesCart struct {
	prg_rom []uint8
	chr_rom []uint8
	sram []uint8
	mapper *Mapper
}

func (cart *NesCart) Load(filename *string) bool {
	contents, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("File length: ",len(contents))
	header := contents[:16]
	if !Equal(header[:4], []byte{'N','E','S',0x1a}) {
		fmt.Println("Header doesn't look correct")
		return false
	}
	for _,c := range(header) {
		fmt.Printf("%02x ", c)
	}
	return true
}
