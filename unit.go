package main

func spawn(s string) string {
	g := Control.gameMap.selGrid
	g.Unit = Tank{5}
	return "Spawned tank at " + Control.gameMap.selGridDes.ToString()
}

type Unit interface {
	move(*Grid)
	style() rune
}

type Tank struct {
	speed int
}

func (t Tank) move(g *Grid) {
	return
}

func (t Tank) style() rune {
	return 'T'
	//return Perspective
	//return QuadDelta
}
