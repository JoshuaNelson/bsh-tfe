package world

import (
	"bsh-tfe/mgrs"
	"logger"
)

var gridSquareSize int = 999

type GridSquare struct {
	Grid map[mgrs.SixDigitCoordinate]Grid
}

func initGridSquare(gsd mgrs.GridSquareDesignation) GridSquare {
	var gs GridSquare
	gs.Grid = make(map[mgrs.SixDigitCoordinate]Grid)

	logger.Debug("Generating biome grids in Grid Square %s.", gsd.ToString())
	for x := 0; x < gridSquareSize; x++ {
		for y := 0; y < gridSquareSize; y++ {
			sdc := mgrs.SixDigitCoordinate{x, y}
			//TODO generate random biomes
			gs.Grid[sdc] = Grid{TERRAIN_FOREST}
		}
	}

	return gs
}

func (gs GridSquare) getGrid(g mgrs.GridDesignation) Grid {
	return gs.Grid[g.SDC]
}
