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
	curGrid *Grid
	curGridDes GridDesignation
	focGridDes GridDesignation
}

func (m *Map) Init() {
	m.initPlanet("Terra", 65) // Hardcoded for now
	grid, err := StringToGridDesignation("1C FC 803 205")
	check(err)
	m.curGrid = m.planet.getGrid(grid)
	m.curGridDes = grid
	m.focGridDes = grid

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
		setInputCLI()

	case termbox.KeyEnter:
		//m.selectGrid()

	case termbox.KeyBackspace, termbox.KeyBackspace2:
		//m.centerOnSelectedGrid()

	case termbox.KeyArrowUp:
		m.adjustNorthing(1)

	case termbox.KeyArrowDown:
		m.adjustNorthing(-1)

	case termbox.KeyArrowLeft:
		m.adjustEasting(-1)

	case termbox.KeyArrowRight:
		m.adjustEasting(1)

	case termbox.KeyPgup:
		m.adjustNorthing(10)

	case termbox.KeyPgdn:
		m.adjustNorthing(-10)
	}

	switch event.Ch {
	case 's':
		//m.selectGrid()

	case 't':
		setInputCLI()

	case 'f':
		m.focusOnGrid()

	case 'k': // vim up
		m.adjustNorthing(1)

	case 'j': // vim down
		m.adjustNorthing(-1)

	case 'h': // vim left
		m.adjustEasting(-1)

	case 'l': // vim right
		m.adjustEasting(1)

	case 'm': // Move unit
		//moveUnit(m.selGrid, m.curGrid)
		Control.Draw()
		return
	}
}
