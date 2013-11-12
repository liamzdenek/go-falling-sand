package fsg

import (
	"github.com/banthar/Go-SDL/sdl"
)

type PixelEngine interface {
	Step(*Grid, uint32, uint32) []*GridUpdate
	Color() sdl.Color
}
