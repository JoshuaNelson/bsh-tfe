package main

import (
	"strings"
)

func usageRoot(s string) string {
	msg := strings.Split(s, " ")
	return "Command '" + msg[0] + "' not recognized."
}

func gridGoto(s string) string {
	grid, err := StringToGridDesignation(s)
	if err != nil {
		return "Invalid grid designation."
	}

	Control.gameMap.focGridDes = grid
	return "Viewing grid " + grid.ToString() + "."
}

func setInputCLI() {
	Control.setInputMode(Control.cli)
	Control.Draw()
}

func (m *Map) adjustNorthing(y int) {
	m.curGridDes = m.curGridDes.adjustNorthing(y)
	m.curGrid = Control.getGrid(m.curGridDes)
	Control.Draw()
}

func (m *Map) adjustEasting(x int) {
	m.curGridDes = m.curGridDes.adjustEasting(x)
	m.curGrid = Control.getGrid(m.curGridDes)
	Control.Draw()
}

func (m *Map) focusOnGrid() {
	m.focGridDes = m.curGridDes
	Control.Draw()
}
