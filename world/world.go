package world

import (
	"bsh-tfe/mgrs"
	"logger"
)

var BIOME_ARID int = 0
var BIOME_FOREST int = 1
var BIOME_GRASSLAND int = 2
var BIOME_ROCKY int = 4

var Terra *Planet

func Init() {
	g, err := mgrs.StringToGridDesignation("1C AB 005 009")
	if err != nil {
		panic(err)
	}
	logger.Debug("Initializing Planet.")
	logger.Debug("Starting with coordinate '%s'", g.ToString())
	//Terra = initPlanet()
	Terra = initPlanet()
	testSelectGrid = Terra.GetGrid(g)
	testSelectGrid = Terra.GetGrid(g)

	// Select the grid
	SelectedGrid = g
}
