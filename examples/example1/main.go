package main

import (
	"fmt"
	"runtime"
	//"flag"
	"github.com/banthar/Go-SDL/sdl"
	//"github.com/go-gl/gl"
	"github.com/liamzdenek/go-falling-sand/engine"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	sdl.Init(sdl.INIT_VIDEO)
	defer sdl.Quit()

	sdl.GL_SetAttribute(sdl.GL_SWAP_CONTROL, 1)
	surface := sdl.SetVideoMode(512, 512, 32, 0)

	if surface == nil {
		panic("sdl error")
	}

	sdl.WM_SetCaption("Falling Sand Game", "Falling Sand Game")

	fmt.Printf("Initializing physics...\n")
	p := NewFSGPhysics()
	fmt.Printf("Initializing the grid...\n")
	g := fsg.NewGrid(512, 512, p, surface)

	g.Run()
}

