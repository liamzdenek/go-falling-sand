package main

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/liamzdenek/go-falling-sand/engine"
	"math/rand"
)

type EngineFire struct{ Physics *FSGPhysics }

func (ef *EngineFire) Step(g *grid.Grid, x, y uint32) []*grid.GridUpdate {
	if rand.Intn(5) == 1 && y != g.Height && g.Grid[x][y+1].Engine == ef.Physics.Empty {
		return []*grid.GridUpdate{
			&grid.GridUpdate{X: x, Y: y + 1, Engine: ef},
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
		return sdl.Color{0xFF, 0, 0, 0}
	} else {
		return sdl.Color{0xFF, 0xFF, 0, 0}
	}
}
