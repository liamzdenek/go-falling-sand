package main

import (
	"github.com/liamzdenek/go-falling-sand/engine"
	"math"
)

type FSGPhysics struct {
	fsg.SimplePhysics

	Empty *EngineEmpty
	Sand  *EngineSand
	Brick *EngineBrick
	Fire  *EngineFire
}

func NewFSGPhysics() *FSGPhysics {
	p := &FSGPhysics{}
	p.Empty = &EngineEmpty{p}
	p.Sand = &EngineSand{p}
	p.Brick = &EngineBrick{p}
	p.Fire = &EngineFire{p}
	return p
}

func (p *FSGPhysics) Init(g *fsg.Grid) {
	sand_center_x := float64(int(g.Width / 2))
	sand_center_y := float64(int(g.Height / 2))

	brick_center_x := float64(int(g.Width / 2))
	brick_center_y := float64(int(g.Height / 4))

	fire_center_x := float64(int(g.Width / 4 * 3))
	fire_center_y := float64(int(g.Height / 2))
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
