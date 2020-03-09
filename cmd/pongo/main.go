package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("Could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Pongo", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("Could not create window: %v", err)
	}
	defer window.Destroy()

	render, err := sdl.CreateRenderer(window, 0, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return fmt.Errorf("Could not  create render: %v", err)
	}
	defer render.Destroy()

	if err := render.DrawLine(10, 10, 10, 10); err != nil {
		return fmt.Errorf("Could not draw a line: %v", err)
	}

	time.Sleep(10 * time.Second)

	return nil

}
