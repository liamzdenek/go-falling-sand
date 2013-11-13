package fsg

import (
	"fmt"
	"github.com/banthar/Go-SDL/sdl"
)

const UpdateHistoryLength = 5

type SimplePhysics struct {
	// UpdateHistory is a mechanism that is used for []GridUpdate preallocation
	UpdateHistoryModulo int
	UpdateHistory       []int

	// this is a list of updates given by the user
	ManualUpdates []*GridUpdate

	// this is a list of acceptable buildable types
	// SelectedType is the index of the currently building item
	BuildableTypes    []PixelEngine
	SelectedType      int

	// bool that will tell whether the mouse button is down. Used for pixel placement
	// events
	IsMouseButtonDown bool
}

func (p *SimplePhysics) Step(g *Grid) {
	update_size := 0
	{
		if p.UpdateHistory == nil {
			p.UpdateHistory = make([]int, UpdateHistoryLength)
		} else {
			for _, val := range p.UpdateHistory {
				update_size += val
			}
			update_size /= len(p.UpdateHistory)
		}
	}

	updates := make([]*GridUpdate, 0, update_size)
	real_count := 0

	for i := range g.Grid {
		for j := range g.Grid[i] {
			u_local := g.Grid[i][j].Engine.Step(g, uint32(i), uint32(j))
			if len(u_local) > 0 {
				updates = append(updates, u_local...)
			}
		}
	}

	{
		p.UpdateHistory[p.UpdateHistoryModulo] = real_count
		if p.UpdateHistoryModulo++; p.UpdateHistoryModulo > len(p.UpdateHistory)-1 {
			p.UpdateHistoryModulo = 0
		}
	}

	// pull in the updates from the user
	{
		updates = append(updates, p.ManualUpdates...);
		p.ManualUpdates = []*GridUpdate{}
	}

	g.ApplyUpdates(updates)
}

func (p *SimplePhysics) HandleEvent(g *Grid, e *sdl.Event) {
	switch e := (*e).(type) {
	case *sdl.KeyboardEvent:
		//fmt.Printf("GOT EVENT: \n", e);
		if e.Type == sdl.KEYDOWN && e.Keysym.Scancode == KEY_TAB {
			p.SelectedType++
			if p.SelectedType >= len(p.BuildableTypes) {
				p.SelectedType = 0
			}
			fmt.Printf("Got tab - now building: %v\n", p.BuildableTypes[p.SelectedType].String())
		}
	case *sdl.MouseButtonEvent:
		if e.Type == sdl.MOUSEBUTTONDOWN {
			p.IsMouseButtonDown = true
		} else if e.Type == sdl.MOUSEBUTTONUP {
			p.IsMouseButtonDown = false
		}
		//fmt.Printf("Mouse click\n", e)
	case *sdl.MouseMotionEvent:
		if p.IsMouseButtonDown {
			p.Set(g, uint32(e.X), (uint32(g.Height)-1)-uint32(e.Y), p.BuildableTypes[p.SelectedType])
		}
		//fmt.Printf("Mouse Move -- %v\n", p.IsMouseButtonDown, e)
	}
}

func (p *SimplePhysics) Set(g *Grid, x, y uint32, engine PixelEngine) {
	p.ManualUpdates = append(p.ManualUpdates, &GridUpdate{
		X: x,
		Y: y,
		Engine: engine,
	});
	//g.Grid[x][y].Engine = engine
}
