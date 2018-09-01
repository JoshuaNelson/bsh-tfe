package world

import (
	"bsh-tfe/mgrs"
	"logger"
)

type GridZone struct {
	zone map[mgrs.GridSquareDesignation]GridSquare
}

func initGridZone(gzd mgrs.GridZoneDesignation) GridZone {
	var gz GridZone
	gz.zone = make(map[mgrs.GridSquareDesignation]GridSquare)

	logger.Debug("Initializing new Grid Zone%s.", gzd.ToString())
	return gz
}

func (gz GridZone) getGrid(g mgrs.GridDesignation) Grid {
	gs, gsInitialized := gz.zone[g.GSD]
	if !gsInitialized {
		gs = initGridSquare(g.GSD)
	}
	return gs.getGrid(g)
}
