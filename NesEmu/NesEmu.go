package NesEmu

import (
	"flag"
	"fmt"
	"os"

        "github.com/khedoros/ghostliNES/NesCpu"
        "github.com/khedoros/ghostliNES/NesMem"
        "github.com/khedoros/ghostliNES/NesPpu"
        "github.com/khedoros/ghostliNES/NesApu"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle = "ghostliNES Emulator"
        winWidth, winHeight = 320, 240
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
	debug bool
	resolution int
	mapper int
	filename string
	mem NesMem.NesMem
	cpu NesCpu.Cpu6502
	ppu NesPpu.NesPpu
	apu NesApu.NesApu
}

func (emu *NesEmu) New() error {
        emu.debug = *flag.Bool("debug", false, "print debug output while running")
        emu.resolution = *flag.Int("res", 1, "integer output scaling")
        emu.mapper = *flag.Int("mapper", -1, "override detected iNES mapper number")
        flag.Parse()
        if flag.NArg() < 1 {
                fmt.Fprint(os.Stderr, "Not enough arguments. You need to at least specify a filename.\n")
                flag.PrintDefaults()
        }
        emu.filename = flag.Arg(0)
        fmt.Printf("Options\n--------\nDebug: %v\nResolution: %v\nMapper: %v\nFile: %v\n", emu.debug, emu.resolution, emu.mapper, emu.filename)

        emu.mem.New(&emu.filename)
        emu.cpu.New(&emu.mem)
        emu.ppu.New(&emu.mem, emu.resolution)
        emu.apu.New()

	return nil
}

func (emu *NesEmu) Destroy() {
	//TODO: Use this to shut down the emulator, etc
}

func (emu *NesEmu) InputEvent(event *sdl.Event) {
	emu.mem.InputEvent(event)
}
