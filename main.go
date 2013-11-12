package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	//"flag"
	"github.com/banthar/Go-SDL/sdl"
	//"github.com/go-gl/gl"
	"./grid"
	"time"
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

	f, _ := os.Create("fsg-profile")
	pprof.StartCPUProfile(f)

	fmt.Printf("Starting Physics Thread...\n")
	go func() {
		fmt.Printf("Initializing physics...\n")
		p := NewFSGPhysics()
		fmt.Printf("Initializing the grid...\n")
		g := grid.NewGrid(512, 512, p, surface)
		fmt.Printf("Beginning physics\n")
		for /*i := 0; i < 100; i++*/ {
			start := time.Now()
			fmt.Printf("Stepping...\n")
			g.Step()
			fmt.Printf("TOOK X: %v\n", time.Since(start))
		}
		pprof.StopCPUProfile()
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

type EngineEmpty struct{ Physics *FSGPhysics }

func (ee *EngineEmpty) Step(g *grid.Grid, x, y uint32) []*grid.GridUpdate { return nil }
func (ee *EngineEmpty) Color() sdl.Color                                  { return sdl.Color{0, 0, 0, 0} }

type EngineSand struct{ Physics *FSGPhysics }

func (es *EngineSand) Step(g *grid.Grid, x, y uint32) []*grid.GridUpdate {
	if y == 0 {
		return nil
	}

	if g.Grid[x][y-1].Engine == es.Physics.Empty {
		return []*grid.GridUpdate{
			&grid.GridUpdate{X: x, Y: y - 1, Engine: es},
			&grid.GridUpdate{X: x, Y: y, Engine: es.Physics.Empty},
		}
	} else {
		if rand.Intn(2) == 1 {
			if x != 0 && g.Grid[x-1][y-1].Engine == es.Physics.Empty {
				return []*grid.GridUpdate{
					&grid.GridUpdate{X: x - 1, Y: y - 1, Engine: es},
					&grid.GridUpdate{X: x, Y: y, Engine: es.Physics.Empty},
				}
			}
		} else {
			if x != g.Width-1 && g.Grid[x+1][y-1].Engine == es.Physics.Empty {
				return []*grid.GridUpdate{
					&grid.GridUpdate{X: x + 1, Y: y - 1, Engine: es},
					&grid.GridUpdate{X: x, Y: y, Engine: es.Physics.Empty},
				}
			}
		}
	}
	return nil
}
func (es *EngineSand) Color() sdl.Color { return sdl.Color{0xFF, 0xFF, 0, 0} }

type EngineBrick struct{ Physics *FSGPhysics }

func (eb *EngineBrick) Step(g *grid.Grid, x, y uint32) []*grid.GridUpdate { return nil }
func (eb *EngineBrick) Color() sdl.Color                                  { return sdl.Color{0x8B, 0, 0, 0} }

type EngineFire struct { Physics *FSGPhysics }

func (ef *EngineFire) Step(g *grid.Grid, x, y uint32) []*grid.GridUpdate {
	if rand.Intn(5) == 1 && y != g.Height && g.Grid[x][y+1].Engine == ef.Physics.Empty {
		return []*grid.GridUpdate{
			&grid.GridUpdate{X: x, Y: y+1, Engine: ef},
			&grid.GridUpdate{X: x, Y: y, Engine: ef.Physics.Empty},
		}
	} else if rand.Intn(10) == 1 {
		return []*grid.GridUpdate{
			&grid.GridUpdate{X: x, Y: y, Engine: ef.Physics.Empty},
		}
	}
	return nil
}

func (ef *EngineFire) Color() sdl.Color {
	if rand.Intn(5) == 1 {
		return sdl.Color{0xFF, 0, 0, 0};
	} else {
		return sdl.Color{0xFF, 0xFF, 0, 0};
	}
}

type FSGPhysics struct {
	grid.SimplePhysics

	Empty *EngineEmpty
	Sand  *EngineSand
	Brick *EngineBrick
	Fire *EngineFire
}

func NewFSGPhysics() *FSGPhysics {
	p := &FSGPhysics{}
	p.Empty = &EngineEmpty{p}
	p.Sand = &EngineSand{p}
	p.Brick = &EngineBrick{p}
	p.Fire = &EngineFire{p}
	return p
}

func (p *FSGPhysics) Init(g *grid.Grid) {
	sand_center_x := float64(int(g.Width / 2))
	sand_center_y := float64(int(g.Height / 2))

	brick_center_x := float64(int(g.Width / 2))
	brick_center_y := float64(int(g.Height / 4))
	
	fire_center_x := float64(int(g.Width / 4 * 3))
	fire_center_y := float64(int(g.Height / 2));
	for i := range g.Grid {
		for j := range g.Grid[i] {
			if math.Sqrt(math.Pow(sand_center_x-float64(i), 2)+math.Pow(sand_center_y-float64(j), 2)) < 100 {
				g.Grid[i][j].Engine = p.Sand
			} else if math.Sqrt(math.Pow(brick_center_x-float64(i), 2)+math.Pow(brick_center_y-float64(j), 2)) < 30 {
				g.Grid[i][j].Engine = p.Brick
			} else if math.Sqrt(math.Pow(fire_center_x-float64(i), 2)+math.Pow(fire_center_y-float64(j), 2)) < 30 {
				g.Grid[i][j].Engine = p.Fire
			} else {
				g.Grid[i][j].Engine = p.Empty
			}
		}
	}
}
