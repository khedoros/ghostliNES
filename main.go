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

	running := true
	var event sdl.Event
	var joysticks [16]*sdl.Joystick

	for running == true {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				mem.InputEvent(&event)
			case *sdl.JoyAxisEvent:
				fmt.Printf("[%d ms] JoyAxis\ttype:%d\twhich:%c\taxis:%d\tvalue:%d\n",
					t.Timestamp, t.Type, t.Which, t.Axis, t.Value)
				mem.InputEvent(&event)
			case *sdl.JoyBallEvent:
				fmt.Printf("[%d ms] JoyBall\ttype:%d\twhich:%d\tball:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp, t.Type, t.Which, t.Ball, t.XRel, t.YRel)
				mem.InputEvent(&event)
			case *sdl.JoyButtonEvent:
				fmt.Printf("[%d ms] JoyButton\ttype:%d\twhich:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.Button, t.State)
				mem.InputEvent(&event)
			case *sdl.JoyHatEvent:
				fmt.Printf("[%d ms] JoyHat\ttype:%d\twhich:%d\that:%d\tvalue:%d\n",
					t.Timestamp, t.Type, t.Which, t.Hat, t.Value)
				mem.InputEvent(&event)
			case *sdl.JoyDeviceAddedEvent:
				joysticks[int(t.Which)] = sdl.JoystickOpen(t.Which)
				if joysticks[int(t.Which)] != nil {
					fmt.Printf("Joystick %d connected\n", t.Which)
				}
			case *sdl.JoyDeviceRemovedEvent:
				if joystick := joysticks[int(t.Which)]; joystick != nil {
					joystick.Close()
				}
				fmt.Printf("Joystick %d disconnected\n", t.Which)
			default:
				fmt.Printf("Some event\n")
			}

		}
		sdl.Delay(16) //TODO: Better frame timing solution
	}

	return 0
}

func main() {
	os.Exit(run())
}
