package main

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/liamzdenek/go-falling-sand/engine"
	"math/rand"
)

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
