package world

import (
	"bsh-tfe/mgrs"
	"github.com/aquilax/go-perlin"
	"logger"
)

type GridSquare struct {
	Grid map[mgrs.SixDigitCoordinate]*Grid
}

func initGridSquare(gsd mgrs.GridSquareDesignation) *GridSquare {
	var gs GridSquare
	gs.Grid = make(map[mgrs.SixDigitCoordinate]*Grid)

	logger.Debug("Spooling up Perlin noise generator. Cover your ears.")
	var seed int64 = 65

	var scale float64 = 25
	p := perlin.NewPerlin(2.1, 2.2, 3, seed)
	//p := perlin.NewPerlin(2, 2, 3, seed)

	// TODO Use scaled ints across world, not just each square

	logger.Debug("Generating new Grid Square, %s.", gsd.ToString())
	for x := 0; x <= mgrs.GridSquareSize; x++ {
		for y := 0; y <= mgrs.GridSquareSize; y++ {
			sdc := mgrs.SixDigitCoordinate{x, y}
			var g Grid
			g.newGrid(p.Noise2D(float64(x)/scale, float64(y)/scale))
			gs.Grid[sdc] = &g
		}
	}

	return &gs
}

func (gs *GridSquare) getGrid(g mgrs.GridDesignation) *Grid {
	return gs.Grid[g.SDC]
}
