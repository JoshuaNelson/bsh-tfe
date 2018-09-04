package world

import (
	"bsh-tfe/mgrs"
)

var SelectedGrid mgrs.GridDesignation
var testSelectGrid *Grid

type Grid struct {
	Biome int
}

func (g *Grid) setBiome(b int) {
	g.Biome = b
}
