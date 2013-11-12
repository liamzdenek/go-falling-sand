package fsg

const UpdateHistoryLength = 5

type SimplePhysics struct {
	UpdateHistoryModulo int
	UpdateHistory       []int
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

	/*
		if update_size > 200 {
			fmt.Printf("RAM Leak detected. aborting\n");
			os.Exit(0);
		}
	*/

	//fmt.Printf("Prepping of %v size...\n", update_size);
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

	//fmt.Printf("Len history: %v\n", p.UpdateHistory);

	{
		p.UpdateHistory[p.UpdateHistoryModulo] = real_count
		if p.UpdateHistoryModulo++; p.UpdateHistoryModulo > len(p.UpdateHistory)-1 {
			p.UpdateHistoryModulo = 0
		}
	}

	g.ApplyUpdates(updates)
}
