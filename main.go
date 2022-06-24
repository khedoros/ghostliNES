package main

import (
	"fmt"
	"os"
	"runtime"

	nesemu "github.com/khedoros/ghostliNES/NesEmu"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = nesemu.GetWindowTitle()
var winWidth, winHeight int32 = nesemu.GetWindowSize()

func init() { runtime.LockOSThread() }

func run() int {

	var emulator nesemu.NesEmu
	var emuInitErr, sdlInitErr, createWinErr, createRendErr error
	emuInitErr = emulator.New()
	if emuInitErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize emulator: %s\n", emuInitErr)
		return 1
	}
	defer emulator.Destroy()

	var window *sdl.Window
	var renderer *sdl.Renderer

	if sdlInitErr = sdl.Init(sdl.INIT_EVERYTHING); sdlInitErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize SDL: %s\n", sdlInitErr)
		return 1
	}
	defer sdl.Quit()

	if window, createWinErr = sdl.CreateWindow(winTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE); createWinErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", createWinErr)
		return 2
	}
	defer window.Destroy()

	if renderer, createRendErr = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); createRendErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", createRendErr)
		return 3 // don't use os.Exit(3); otherwise, previous deferred calls will never run
	}
	defer renderer.Destroy()
	renderer.Clear()

	running := true
	var event sdl.Event
	var joysticks [16]*sdl.Joystick
	renderWidth, renderHeight := nesemu.GetRenderSize()
	frameBuffer, _ := sdl.CreateRGBSurface(0, renderWidth, renderHeight, 32, 0, 0, 0, 0)
	texture, _ := renderer.CreateTextureFromSurface(frameBuffer)

	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%d\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				emulator.InputEvent(&event)
				if t.Keysym.Sym == sdl.K_q {
					running = false
				}
			case *sdl.JoyAxisEvent:
				fmt.Printf("[%d ms] JoyAxis\ttype:%d\twhich:%c\taxis:%d\tvalue:%d\n",
					t.Timestamp, t.Type, t.Which, t.Axis, t.Value)
				emulator.InputEvent(&event)
			case *sdl.JoyBallEvent:
				fmt.Printf("[%d ms] JoyBall\ttype:%d\twhich:%d\tball:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp, t.Type, t.Which, t.Ball, t.XRel, t.YRel)
				emulator.InputEvent(&event)
			case *sdl.JoyButtonEvent:
				fmt.Printf("[%d ms] JoyButton\ttype:%d\twhich:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.Button, t.State)
				emulator.InputEvent(&event)
			case *sdl.JoyHatEvent:
				fmt.Printf("[%d ms] JoyHat\ttype:%d\twhich:%d\that:%d\tvalue:%d\n",
					t.Timestamp, t.Type, t.Which, t.Hat, t.Value)
				emulator.InputEvent(&event)
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
				//fmt.Printf("Some event\n")
			}

		}

		emulator.RunFrame()
		newFrame := emulator.GetFrame()
		renderer.Clear()
		texture.Update(nil, *newFrame, int(renderWidth)*4)
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		sdl.Delay(16) //TODO: Better frame timing solution
	}

	return 0
}

func main() {
	os.Exit(run())
}
