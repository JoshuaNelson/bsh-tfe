package world

import (
	"bsh-tfe/mgrs"
	"logger"
	"math/rand"
	"time"
)

type GridSquare struct {
	Grid map[mgrs.SixDigitCoordinate]*Grid
}

func initGridSquare(gsd mgrs.GridSquareDesignation) *GridSquare {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var gs GridSquare
	gs.Grid = make(map[mgrs.SixDigitCoordinate]*Grid)

	logger.Debug("Generating new Grid Square, %s.", gsd.ToString())
	for x := 0; x <= mgrs.GridSquareSize; x++ {
		for y := 0; y <= mgrs.GridSquareSize; y++ {
			sdc := mgrs.SixDigitCoordinate{x, y}
			//TODO generate random biomes
			var g Grid
			g.setBiome(r.Intn(5))
			gs.Grid[sdc] = &g
		}
	}

	return &gs
}

func (gs *GridSquare) getGrid(g mgrs.GridDesignation) *Grid {
	return gs.Grid[g.SDC]
}
