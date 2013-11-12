package fsg

type Physics interface {
	Init(g *Grid)
	Step(g *Grid)
}

