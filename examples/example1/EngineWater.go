package main

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/liamzdenek/go-falling-sand/engine"
	"math/rand"
)

type EngineWater struct{ Physics *FSGPhysics }

func (ew *EngineWater) String() string { return "Water" }

func (ew *EngineWater) Color() sdl.Color { return sdl.Color{0x16, 0xAC, 0xB0, 0} }

func (ew *EngineWater) Step(g *fsg.Grid, x, y uint32) []*fsg.GridUpdate {
	if y != 0 {
		if g.Grid[x][y-1].Engine == ew.Physics.Empty {
			return []*fsg.GridUpdate{
				&fsg.GridUpdate{X: x, Y: y - 1, Engine: ew},
				&fsg.GridUpdate{X: x, Y: y, Engine: g.Grid[x][y-1].Engine},
			}
		} else {
			if rand.Intn(2) == 1 {
				if x > 0 && g.Grid[x-1][y-1].Engine == ew.Physics.Empty {
					return []*fsg.GridUpdate{
						&fsg.GridUpdate{X: x - 1, Y: y - 1, Engine: ew},
						&fsg.GridUpdate{X: x, Y: y, Engine: ew.Physics.Empty},
					}
				}
			} else {
				if x < g.Width && g.Grid[x+1][y-1].Engine == ew.Physics.Empty {
					return []*fsg.GridUpdate{
						&fsg.GridUpdate{X: x + 1, Y: y - 1, Engine: ew},
						&fsg.GridUpdate{X: x, Y: y, Engine: ew.Physics.Empty},
					}
				}
			}
		}
	}
		if rand.Intn(2) == 1 {
			if x > 1 && g.Grid[x-1][y].Engine == ew.Physics.Empty &&
				g.Grid[x-2][y].Engine != ew &&
				(
					y >= g.Height-1 ||
					g.Grid[x-1][y+1].Engine == ew.Physics.Empty) {
				return []*fsg.GridUpdate{
					&fsg.GridUpdate{X: x - 1, Y: y, Engine: ew},
					&fsg.GridUpdate{X: x, Y: y, Engine: ew.Physics.Empty},
				}
			}
		} else {
			if x < g.Width-1 && g.Grid[x+1][y].Engine == ew.Physics.Empty &&
				g.Grid[x+2][y].Engine != ew &&
				(
					y >= g.Height-1 ||
					g.Grid[x+1][y+1].Engine ==ew.Physics.Empty) {
				return []*fsg.GridUpdate{
					&fsg.GridUpdate{X: x + 1, Y: y, Engine: ew},
					&fsg.GridUpdate{X: x, Y: y, Engine: ew.Physics.Empty},
				}
			}
		}
	return nil
}
