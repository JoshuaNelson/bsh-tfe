package world

import (
	"bsh-tfe/mgrs"
	"logger"
)

var TERRAIN_ARID int = 0
var TERRAIN_FOREST int = 1

var Terra Planet
var SelectedGrid mgrs.GridDesignation

var testSelectGrid *Grid

type Terrain struct {
	x int
	y int
	Biome int
	//units []Unit // TODO add unit per terrain
}

func (t *Terrain) SetBiome(b int) {
	t.Biome = b
}

func Init() {
	g, err := mgrs.StringToGridDesignation("1C AB 100 200")
	if err != nil {
		panic(err)
	}
	logger.Debug("Initializing Planet.")
	logger.Debug("Starting with coordinate '%s'", g.ToString())
	Terra = initPlanet()
	testSelectGrid = Terra.GetGrid(g)

	// Select the grid
	SelectedGrid = g
}
