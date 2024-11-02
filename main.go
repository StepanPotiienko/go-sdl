package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth      = 800
	screenHeight     = 600
	applicationTitle = "GO SDL Triangle"

	initVideo         = sdl.INIT_VIDEO
	positionUndefined = sdl.WINDOWPOS_UNDEFINED
)

var (
	buf          bytes.Buffer
	logger       = log.New(&buf, "logger:", log.Lshortfile)
	isRunning    bool = true
	window       *sdl.Window
	renderer     *sdl.Renderer
	points       []sdl.Point
)

// LogOutput takes multiple parameters for logging
// The first one is buffer, the second one is output string, and the last one is errorCode. Other ones are ignored.
func LogOutput(buffer *bytes.Buffer, params ...any) {
	if len(params) > 1 {
		logger.Println(params[0], params[1])
	} else {
		logger.Println(params[0])
	}
}

func QuitAppWithCode(exitCode int) {
	os.Exit(exitCode)
}

// DrawTriangle draws a triangle using the defined points
func DrawTriangle(points []sdl.Point) {
	if len(points) != 3 {
		LogOutput(&buf, "DrawTriangle requires exactly 3 points")
		return
	}

	// Set the draw color for the triangle (white)
	if err := renderer.SetDrawColor(255, 255, 255, 255); err != nil {
		LogOutput(&buf, "Failed to set draw color:", err)
		return
	}

	// Draw lines between the points
	for i := 0; i < 3; i++ {
		nextIndex := (i + 1) % 3
		if err := renderer.DrawLine(points[i].X, points[i].Y, points[nextIndex].X, points[nextIndex].Y); err != nil {
			LogOutput(&buf, "Failed to draw line:", err)
		}
	}
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

func RunApplication() (err error) {
	if err = sdl.Init(initVideo); err != nil {
		LogOutput(&buf, "Failed to initialize SDL:", err)
		return err
	}
	defer sdl.Quit()

	window, err = CreateWindow()
	if err != nil {
		LogOutput(&buf, "Failed to create window:", err)
		return err
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		LogOutput(&buf, "Failed to create renderer:", err)
		return err
	}

	points = []sdl.Point{
		{X: 400, Y: 100}, // Top point
		{X: 300, Y: 400}, // Bottom left point
		{X: 500, Y: 400}, // Bottom right point
	}

	DrawTriangle(points)

	// Present the renderer to show the changes
	renderer.Present()

	for isRunning {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				isRunning = false
			}
		}
	}

	err = window.Destroy()
	if err != nil {
		LogOutput(&buf, "Failed to destroy window:", err)
		return err
	}

	fmt.Print(&buf)
	return nil
}

func main() {
	if err := RunApplication(); err != nil {
		LogOutput(&buf, "Application error:", err)
		fmt.Print(&buf)
		fmt.Println("SDL Error:", sdl.GetError())
		QuitAppWithCode(1)
	}
}
