package world

import (
	"bsh-tfe/mgrs"
	"logger"
)

type Planet struct {
	GridZone map[mgrs.GridZoneDesignation]GridZone
}

func initPlanet() Planet {
	var p Planet
	p.GridZone = make(map[mgrs.GridZoneDesignation]GridZone)

	logger.Debug("Generating new planet.")
	return p
}

func (p Planet) GetGrid(g mgrs.GridDesignation) *Grid {
	gz, gzInitialized := p.GridZone[g.GZD]
	if !gzInitialized {
		gz = initGridZone(g.GZD)
	}
	return gz.getGrid(g)
}
