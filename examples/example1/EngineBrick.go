package main

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/liamzdenek/go-falling-sand/engine"
)

type EngineBrick struct{ Physics *FSGPhysics }

func (eb *EngineBrick) String() string { return "Brick" }
func (eb *EngineBrick) Step(g *fsg.Grid, x, y uint32) []*fsg.GridUpdate { return nil }
func (eb *EngineBrick) Color() sdl.Color                                { return sdl.Color{0x8B, 0, 0, 0} }
