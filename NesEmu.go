package NesEmu

import (
	"flag"
	"fmt"
	"os"

        "github.com/khedoros/ghostliNES/NesCpu"
        "github.com/khedoros/ghostliNES/NesMem"
        "github.com/khedoros/ghostliNES/NesPpu"
        "github.com/khedoros/ghostliNES/NesApu"
)

const (
	WindowTitle = "ghostliNES Emulator"
        WinWidth, WinHeight = 320, 240
	RenderWidth, RenderHeight = 256, 240
)

func GetWindowTitle() *String {
	return windowTitle
}

func GetWindowSize() int32, int32 {
	return WinWidth, WinHeight
}

func GetRenderSize() int32, int32 {
	return RenderWidth, RenderHeight
}

struct NesEmu {
	debug bool
	resolution int
	mapper int
	filename string
	mem NesMem.NesMem
	cpu NesCpu.Cpu6502
	ppu NesPpu.NesPpu
	apu NesApu.NesApu
}

func (emu *NesEmu) New() {
        emu.debug = flag.Bool("debug", false, "print debug output while running")
        emu.resolution = flag.Int("res", 1, "integer output scaling")
        emu.mapper = flag.Int("mapper", -1, "override detected iNES mapper number")
        flag.Parse()
        if flag.NArg() < 1 {
                fmt.Fprint(os.Stderr, "Not enough arguments. You need to at least specify a filename.\n")
                flag.PrintDefaults()
        }
        emu.filename = flag.Arg(0)
        fmt.Printf("Options\n--------\nDebug: %#v\nResolution: %#v\nMapper: %#v\nFile: %v\n", *debug, *res, *mapper, filename)

        emu.mem.New(&filename)
        emu.cpu.New(&mem)
        emu.ppu.New(&mem, res)
        emu.apu.New()
}

func (emu *NesEmu) InputEvent(event &sdl.Event) {
	emu.mem.InputEvent(event)
}
