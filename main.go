package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/khedoros/ghostliNES/NesCpu"
	"github.com/veandco/go-sdl2/gfx"
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
	var vx, vy = make([]int16, 3), make([]int16, 3)
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
	renderer.Clear()
	defer renderer.Destroy()

	vx[0] = int16(winWidth / 3)
	vy[0] = int16(winHeight / 3)
	vx[1] = int16(winWidth * 2 / 3)
	vy[1] = int16(winHeight / 3)
	vx[2] = int16(winWidth / 2)
	vy[2] = int16(winHeight * 2 / 3)
	gfx.FilledPolygonColor(renderer, vx, vy, sdl.Color{0, 0, 255, 255})

	gfx.CharacterColor(renderer, winWidth-16, 16, 'X', sdl.Color{255, 0, 0, 255})
	gfx.StringColor(renderer, 16, 16, "GFX Demo", sdl.Color{0, 255, 0, 255})

	renderer.Present()
	sdl.Delay(3000)

	var cpu NesCpu.Cpu6502
	cpu.New()

	return 0
}

func main() {
	os.Exit(run())
}
