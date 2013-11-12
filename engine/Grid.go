package grid

import (
	"github.com/banthar/Go-SDL/sdl"
	"fmt"
	"runtime"
)

const (
	ENGINE_EMPTY = iota
)

type GridUpdate struct {
	X, Y   uint32
	Engine PixelEngine
}

type PixelEngine interface {
	Step(*Grid, uint32, uint32) []*GridUpdate
	Color() sdl.Color
}

type Pixel struct {
	Engine PixelEngine
}

type Grid struct {
	// consts
	Width, Height uint32

	// grid
	Grid [][]Pixel

	// physics engine
	Physics Physics

	Surface *sdl.Surface
}

func NewGrid(w, h uint32, physics Physics, surface *sdl.Surface) (g *Grid) {
	g = &Grid{
		Width:   w,
		Height:  h,
		Grid:    make([][]Pixel, w),
		Physics: physics,
		Surface: surface,
	}
	for i := range g.Grid {
		g.Grid[i] = make([]Pixel, h)
	}
	physics.Init(g)

	g.Blit();
	return
}

func (g *Grid) Run() {
	fmt.Printf("Starting Physics Thread...\n")
	go func() {
		fmt.Printf("Beginning physics\n")
		for {
			g.Step()
		}
	}()
	for {
		runtime.LockOSThread()
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.MouseButtonEvent:
				fmt.Printf("Mouse click\n")
			case *sdl.MouseMotionEvent:
				fmt.Printf("Mouse Move\n")
			}
		}
	}
}

func (g *Grid) Step() {
	g.Physics.Step(g)
}

func (g *Grid) ApplyUpdates(updates []*GridUpdate) {
	for _, update := range updates {
		if update == nil {
			break;
		}
		//fmt.Printf("UPDATE: %v x %v = %v\n", update.X, update.Y, update.Engine.Color());
		g.Grid[update.X][update.Y].Engine = update.Engine
		g.Surface.Set(int(update.X), int((g.Height-1)-update.Y), g.Grid[update.X][update.Y].Engine.Color())
	}
	g.Surface.UpdateRect(0, 0, g.Width, g.Height)
}

func (g *Grid) Blit() {
	for i := range g.Grid {
		for j := range g.Grid[i] {
			g.Surface.Set(i, int(g.Height-1)-j, g.Grid[i][j].Engine.Color())
		}
	}
}
type Physics interface {
	Init(g *Grid)
	Step(g *Grid)
}

const UpdateHistoryLength = 5

type SimplePhysics struct{
	UpdateHistoryModulo int
	UpdateHistory []int
}

func (p *SimplePhysics) Step(g *Grid) {
	update_size := 0;
	{
		if(p.UpdateHistory == nil) {
			p.UpdateHistory = make([]int, UpdateHistoryLength);
		} else {
			for _, val := range p.UpdateHistory {
				update_size += val;
			}
			update_size /= len(p.UpdateHistory);
		}
	}

	/*
	if update_size > 200 {
		fmt.Printf("RAM Leak detected. aborting\n");
		os.Exit(0);
	}
	*/

	//fmt.Printf("Prepping of %v size...\n", update_size);
	updates := make([]*GridUpdate, 0, update_size);
	real_count := 0;

	for i := range g.Grid {
		for j := range g.Grid[i] {
			u_local := g.Grid[i][j].Engine.Step(g, uint32(i), uint32(j))
			if len(u_local) > 0 {
				updates = append(updates, u_local...);
			}
		}
	}

	//fmt.Printf("Len history: %v\n", p.UpdateHistory);

	{
		p.UpdateHistory[p.UpdateHistoryModulo] = real_count;
		if p.UpdateHistoryModulo++; p.UpdateHistoryModulo > len(p.UpdateHistory)-1 {
			p.UpdateHistoryModulo = 0;
		}
	}

	g.ApplyUpdates(updates)
}
