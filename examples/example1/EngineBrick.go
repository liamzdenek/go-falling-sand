package main

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/liamzdenek/go-falling-sand/engine"
)

type EngineBrick struct{ Physics *FSGPhysics }

func (eb *EngineBrick) Step(g *grid.Grid, x, y uint32) []*grid.GridUpdate { return nil }
func (eb *EngineBrick) Color() sdl.Color                                  { return sdl.Color{0x8B, 0, 0, 0} }
