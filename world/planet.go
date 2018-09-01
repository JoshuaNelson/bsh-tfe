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
	logger.Debug("Generating planet.")
	p.GridZone = make(map[mgrs.GridZoneDesignation]GridZone)
	return p
}

func (p Planet) GetGrid(g mgrs.GridDesignation) Grid {
	return p.GridZone[g.GZD].getGrid(g)
}
