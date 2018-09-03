package world

import (
	"bsh-tfe/mgrs"
	"logger"
)

type GridZone struct {
	GridSquare map[mgrs.GridSquareDesignation]GridSquare
}

func initGridZone(gzd mgrs.GridZoneDesignation) GridZone {
	var gz GridZone
	gz.GridSquare = make(map[mgrs.GridSquareDesignation]GridSquare)

	logger.Debug("Generating new Grid Zone: %s.", gzd.ToString())
	return gz
}

func (gz GridZone) getGrid(g mgrs.GridDesignation) *Grid {
	gs, gsInitialized := gz.GridSquare[g.GSD]
	if !gsInitialized {
		gs = initGridSquare(g.GSD)
	}
	return gs.getGrid(g)
}
