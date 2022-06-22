package nesemu

import (
	"flag"
	"fmt"
	"os"
	"strings"

	nesapu "github.com/khedoros/ghostliNES/NesApu"
	nescpu "github.com/khedoros/ghostliNES/NesCpu"
	nesmem "github.com/khedoros/ghostliNES/NesMem"
	nesppu "github.com/khedoros/ghostliNES/NesPpu"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle               = "ghostliNES Emulator"
	winWidth, winHeight       = 320, 240
	renderWidth, renderHeight = 256, 240
)

func GetWindowTitle() string {
	return windowTitle
}

func GetWindowSize() (int32, int32) {
	return winWidth, winHeight
}

func GetRenderSize() (int32, int32) {
	return renderWidth, renderHeight
}

type NesEmu struct {
	debug      bool
	resolution int
	mapper     int
	filename   string
	region     string
	frame      []byte
	mem        nesmem.NesMem
	cpu        nescpu.CPU6502
	ppu        nesppu.NesPpu
	apu        nesapu.NesApu
}

func (emu *NesEmu) New() error {
	flag.BoolVar(&emu.debug, "debug", false, "print debug output while running")
	flag.IntVar(&emu.resolution, "res", 1, "integer output scaling")
	flag.IntVar(&emu.mapper, "mapper", -1, "override detected iNES mapper number")
	flag.StringVar(&emu.region, "region", "ntsc", "override detected ROM region (ntsc/pal)")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprint(os.Stderr, "Not enough arguments. You need to at least specify a filename.\n")
		flag.PrintDefaults()
		return fmt.Errorf("Not enough arguments. You need to at least specify a filename.")
	}
	emu.region = strings.ToLower(emu.region)
	if emu.region != "ntsc" && emu.region != "pal" {
		return fmt.Errorf("Invalid region \"%s\" specified", emu.region)
	}

	emu.filename = flag.Arg(0)
	fmt.Printf("Options\n--------\nDebug: %v\nResolution: %v\nMapper: %v\nFile: %v\n", emu.debug, emu.resolution, emu.mapper, emu.filename)

	emu.frame = make([]byte, renderWidth*renderHeight*4)
	emu.mem.New(&emu.filename, emu.mapper, &emu.ppu, &emu.apu)
	emu.cpu.New(&emu.mem)
	emu.ppu.New(emu.mem.GetCart(), emu.region)
	emu.apu.New()

	return nil
}

func (emu *NesEmu) Destroy() {
	//TODO: Use this to shut down the emulator, etc
}

func (emu *NesEmu) InputEvent(event *sdl.Event) {
	emu.mem.InputEvent(event)
}

func (emu *NesEmu) RunFrame() *[]byte {
	frameDone := false
	for !frameDone {
		//run a chunk of CPU
		opChunk := int64(10000)
		emu.cpu.Run(opChunk)

		//finish PPU render
		frameDone = emu.ppu.Run(opChunk)
		//finish APU render
	}
	emu.GetFrame()
	return &emu.frame

}

func (emu *NesEmu) GetFrame() *[]byte {
	frame := emu.ppu.Render()
	s := emu.frame
	for i, v := range *frame {
		s[i*4] = v.R
		s[i*4+1] = v.G
		s[i*4+2] = v.B
		s[i*4+3] = 255
	}
	return &emu.frame
}
