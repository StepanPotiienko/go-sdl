package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth      = 1920
	screenHeight     = 1200
	applicationTitle = "Hello World!"

	initVideo         = sdl.INIT_VIDEO
	positionUndefined = sdl.WINDOWPOS_UNDEFINED
)

var (
	isRunning bool = true

	window        *sdl.Window
	screenSurface *sdl.Surface
)

func QuitApp(exitCode int) {
	os.Exit(exitCode)
}

func CreateWindow() (window *sdl.Window, err error) {
	return sdl.CreateWindow(
		applicationTitle,
		positionUndefined,
		positionUndefined,
		screenWidth,
		screenHeight,
		sdl.WINDOW_SHOWN,
	)
}

func FillRect(r, g, b uint8) {
	screenSurface.FillRect(nil, sdl.MapRGB(screenSurface.Format, r, g, b))
	window.UpdateSurface()
}

func RunApplication() (err error) {
	if err = sdl.Init(initVideo); err != nil {
		return err
	}

	window, err = CreateWindow()
	if err != nil {
		panic(err)
	}

	// This line of code does not execute until createWindow() does not return.
	defer sdl.Quit()

	if screenSurface, err = window.GetSurface(); err != nil {
		return err
	}

	FillRect(0x00, 0xFF, 0x00)

	for isRunning {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				isRunning = false
			}
		}
	}

	err = window.Destroy()
	return err
}

func main() {
	if err := RunApplication(); err != nil {
		QuitApp(1)
	}
}
