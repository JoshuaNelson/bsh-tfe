package world

import (
	"bsh-tfe/mgrs"
)

var SelectedGrid mgrs.GridDesignation

type Grid struct {
	Biome int
	North *Grid
	//South *Grid
	East  *Grid
	//West  *Grid
}

func (g *Grid) setBiome(b int) {
	g.Biome = b
}
