package main

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/liamzdenek/go-falling-sand/engine"
)

type EngineEmpty struct{ Physics *FSGPhysics }

func (ee *EngineEmpty) Step(g *fsg.Grid, x, y uint32) []*fsg.GridUpdate { return nil }
func (ee *EngineEmpty) Color() sdl.Color                                  { return sdl.Color{0, 0, 0, 0} }
