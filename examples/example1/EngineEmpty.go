package main;

import ("github.com/liamzdenek/go-falling-sand/engine";"github.com/banthar/Go-SDL/sdl");

type EngineEmpty struct{ Physics *FSGPhysics }

func (ee *EngineEmpty) Step(g *grid.Grid, x, y uint32) []*grid.GridUpdate { return nil }
func (ee *EngineEmpty) Color() sdl.Color                                  { return sdl.Color{0, 0, 0, 0} }

