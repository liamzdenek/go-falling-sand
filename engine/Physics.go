package fsg

import (
	"github.com/banthar/Go-SDL/sdl"
)

type Physics interface {
	Init(g *Grid)
	Step(g *Grid)
	HandleEvent(g *Grid, e *sdl.Event)
}
