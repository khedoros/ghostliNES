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
	frame      *sdl.Surface
	mem        nesmem.NesMem
	cpu        nescpu.CPU6502
	ppu        nesppu.NesPpu
	apu        nesapu.NesApu
}

func (emu *NesEmu) New() error {
	emu.debug = *flag.Bool("debug", false, "print debug output while running")
	emu.resolution = *flag.Int("res", 1, "integer output scaling")
	emu.mapper = *flag.Int("mapper", -1, "override detected iNES mapper number")
	emu.region = *flag.String("region", "ntsc", "override detected ROM region (ntsc/pal)")
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

func (emu *NesEmu) RunFrame() {
	//run a frame of CPU
	opChunk := int64(10000)
	emu.cpu.Run(opChunk)
	//finish PPU render
	frameDone := emu.ppu.Run(opChunk)
	if frameDone {
		emu.GetFrame()
	}
	//finish APU render

}

func (emu *NesEmu) GetFrame() *sdl.Surface {
	if emu.frame == nil {
		emu.frame, _ = sdl.CreateRGBSurface(0, renderWidth, renderHeight, 32, 0, 0, 0, 0)
	}
	frame := emu.ppu.Render()
	s := emu.frame.Pixels()
	for i, v := range *frame {
		s[i*4] = v.R
		s[i*4+1] = v.G
		s[i*4+2] = v.B
		s[i*4+3] = 255
	}
	for y := 0; y < renderHeight; y++ {
		for x := 0; x < renderWidth; x++ {
			//s.Set(x,y,color.RGBA{uint8(x),uint8(y),255,255})
		}
	}
	return emu.frame
}
