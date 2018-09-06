package world

import (
	"bsh-tfe/mgrs"
	"logger"
)

var SelectedPlanet *Planet

type Planet struct {
	Name string
	GridZone map[mgrs.GridZoneDesignation]*GridZone
}

func InitPlanet(name string) *Planet {
	var p Planet
	p.Name = name
	p.GridZone = make(map[mgrs.GridZoneDesignation]*GridZone)

	logger.Debug("Generating new planet, %s.", name)
	return &p
}

func (p *Planet) GetGrid(g mgrs.GridDesignation) *Grid {
	_, gzInitialized := p.GridZone[g.GZD]
	if !gzInitialized {
		p.GridZone[g.GZD] = initGridZone(g.GZD)
	}
	return p.GridZone[g.GZD].getGrid(g)
}
