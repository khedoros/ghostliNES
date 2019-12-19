package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/khedoros/ghostliNES/NesCpu"
	"github.com/khedoros/ghostliNES/NesMem"
	"github.com/khedoros/ghostliNES/NesPpu"
	"github.com/khedoros/ghostliNES/NesApu"
	//"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"

)

var winTitle string = "ghostliNES emulator"
var winWidth, winHeight int32 = 320, 240

func init() { runtime.LockOSThread() }

func run() int {

	var debug = flag.Bool("debug", false, "print debug output while running")
	var res = flag.Int("res", 1, "integer output scaling")
	var mapper = flag.Int("mapper", -1, "override detected iNES mapper number")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprint(os.Stderr, "Not enough arguments. You need to at least specify a filename.\n")
		flag.PrintDefaults()
	}
	var filename = flag.Arg(0)
	fmt.Printf("Options\n--------\nDebug: %#v\nResolution: %#v\nMapper: %#v\nFile: %v\n", *debug, *res, *mapper, filename)

	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize SDL: %s\n", err)
		return 1
	}
	defer sdl.Quit()

	if window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 2
	}
	defer window.Destroy()

	if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 3 // don't use os.Exit(3); otherwise, previous deferred calls will never run
	}
	defer renderer.Destroy()
	renderer.Clear()

	var mem NesMem.NesMem
	mem.New(&filename)

	var cpu NesCpu.Cpu6502
	cpu.New(&mem)

	var ppu NesPpu.NesPpu
	ppu.New(&mem, res)

	var apu NesApu.NesApu
	apu.New()

	return 0
}

func main() {
	os.Exit(run())
}
