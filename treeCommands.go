package main

import (
	"strconv"
	"strings"
)

func usageRoot(s string) string {
	msg := strings.Split(s, " ")
	return "Command '" + msg[0] + "' not recognized."
}

func gridSelect(s string) string {
	grid, err := StringToGridDesignation(s)
	if err != nil {
		return "Invalid grid designation."
	}

	Control.gameMap.selGridDes = grid
	Control.gameMap.selGrid = Control.getGrid(grid)
	return "Selected grid " + grid.ToString() + "."
}

func gridGoto(s string) string {
	grid, err := StringToGridDesignation(s)
	if err != nil {
		return "Invalid grid designation."
	}

	Control.gameMap.mapGridDes = grid
	Control.gameMap.curGridDes = grid
	Control.gameMap.curGrid = Control.getGrid(grid)
	return "Viewing grid " + grid.ToString() + "."
	}

func gridInfo(s string) string {
	b := Control.gameMap.selGrid.Biome
	return "Grid " + Control.gameMap.selGridDes.ToString() + ": Biome " + strconv.Itoa(b.code)
}
