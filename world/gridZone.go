package world

import (
	"bsh-tfe/mgrs"
//	"logger"
)

type GridZone struct {
	gs map[mgrs.GridSquareDesignation]GridSquare
}

func (gz GridZone) getGrid(g mgrs.GridDesignation) Grid {
	return gz.gs[g.GSD].getGrid(g)
}
