package NesApu

//An NesApu represents an NES's Audio Processing Unit
type NesApu struct {
	Blah int8
}

func (this *NesApu) New() {
}

func (this *NesApu) Read(addr uint16, cycle uint64) uint8 {
        return 0
}

func (this *NesApu) Write(addr uint16, val uint8, cycle uint64) {
}

