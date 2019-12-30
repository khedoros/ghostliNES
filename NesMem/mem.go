package NesMem

import (
	"fmt"
	"github.com/khedoros/ghostliNES/NesCart"
)

//An NesMem struct holds the state of the NES's memory mapping circuitry
type NesMem struct {
	Blah int8
	cart *NesCart.NesCart
}

func (this *NesMem) New(filename *string) {
	cart := NesCart.NesCart{}
	fmt.Println("Loading file ", filename)
	valid := cart.Load(filename)
	if !valid {
		fmt.Println("File failed to load")
	}
}
