package main

import (
	"github.com/nsf/termbox-go"
	"logger"
)

/*
 * MAP MAP MAP
 */
type Map struct {
	planet *planet
	selGrid *Grid
	curGrid *Grid
	selGridDes GridDesignation
	curGridDes GridDesignation
	mapGridDes GridDesignation
}

func (m *Map) Init() {
	m.initPlanet("Terra", 65) // Hardcoded for now
	grid, err := StringToGridDesignation("1C FC 803 205")
	check(err)
	m.selGrid = m.planet.getGrid(grid)
	m.curGrid = m.selGrid
	m.selGridDes = grid
	m.curGridDes = grid
	m.mapGridDes = grid

}

func (m *Map) initPlanet(name string, seed int) {
	m.planet = &planet{}
	m.planet.name = name
	m.planet.gridZone = make(map[GridZoneDesignation]*GridZone)
	m.planet.seed = seed

	logger.Debug("Generating new planet, %s.", name)
}

func (m *Map) EventHandler(event termbox.Event) {
	switch event.Key {
	case termbox.KeyEsc:
		Control.setInputMode(Control.cli)
		Control.Draw()
		return

	case termbox.KeyEnter:
		m.selGrid = m.curGrid
		m.selGridDes = m.curGridDes
		Control.Draw()
		return

	case termbox.KeyArrowUp:
		m.curGridDes = m.curGridDes.adjustNorthing(1)
		m.curGrid = Control.getGrid(m.curGridDes)
		Control.Draw()
		return

	case termbox.KeyArrowDown:
		m.curGridDes = m.curGridDes.adjustNorthing(-1)
		m.curGrid = Control.getGrid(m.curGridDes)
		Control.Draw()
		return

	case termbox.KeyArrowLeft:
		m.curGridDes = m.curGridDes.adjustEasting(-1)
		m.curGrid = Control.getGrid(m.curGridDes)
		Control.Draw()
		return

	case termbox.KeyArrowRight:
		m.curGridDes = m.curGridDes.adjustEasting(1)
		m.curGrid = Control.getGrid(m.curGridDes)
		Control.Draw()
		return

	case termbox.KeyBackspace, termbox.KeyBackspace2:
		m.mapGridDes = m.selGridDes
		m.curGridDes = m.selGridDes
		m.curGrid = m.selGrid
		Control.Draw()
		return

	case termbox.KeyPgup:
		m.curGridDes = m.curGridDes.adjustNorthing(10)
		m.curGrid = Control.getGrid(m.curGridDes)
		Control.Draw()
		return

	case termbox.KeyPgdn:
		m.curGridDes = m.curGridDes.adjustNorthing(-10)
		m.curGrid = Control.getGrid(m.curGridDes)
		Control.Draw()
		return
	}

	switch event.Ch {
	case 's':
		m.selGrid = m.curGrid
		m.selGridDes = m.curGridDes
		Control.Draw()
		return

	case 't':
		Control.setInputMode(Control.cli)
		Control.Draw()
		return

	case 'f':
		m.mapGridDes = m.curGridDes
		Control.Draw()
		return

	case 'k': // vim up
		m.curGridDes = m.curGridDes.adjustNorthing(1)
		m.curGrid = Control.getGrid(m.curGridDes)
		Control.Draw()
		return

	case 'j': // vim down
		m.curGridDes = m.curGridDes.adjustNorthing(-1)
		m.curGrid = Control.getGrid(m.curGridDes)
		Control.Draw()
		return

	case 'h': // vim left
		m.curGridDes = m.curGridDes.adjustEasting(-1)
		m.curGrid = Control.getGrid(m.curGridDes)
		Control.Draw()
		return

	case 'l': // vim right
		m.curGridDes = m.curGridDes.adjustEasting(1)
		m.curGrid = Control.getGrid(m.curGridDes)
		Control.Draw()
		return

	case 'm': // Move unit
		moveUnit(m.selGrid, m.curGrid)
		Control.Draw()
		return
	}
}
