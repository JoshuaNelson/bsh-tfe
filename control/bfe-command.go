package control

import (
	"bsh-tfe/mgrs"
	"bsh-tfe/world"
	"strconv"
	"strings"
)

func Init() {
	CommandLine = InitMainCommandLine()
	SetInputMode(CommandLine)
}

func usageRoot(s string) string {
	msg := strings.Split(s, " ")
	return "Command '" + msg[0] + "' not recognized."
}

func gridSelect(s string) string {
	grid, err := mgrs.StringToGridDesignation(s)
	if err != nil {
		return "Invalid grid designation."
	}

	world.SelectedGrid = grid
	return "Selected grid " + grid.ToString() + "."
}

func gridSetBiome(s string) string {
	msg := strings.Split(s, " ")
	t, err := strconv.Atoi(msg[0])
	if err != nil {
		return "Usage: map set biome <biome>"
	}

	if world.SelectedGrid != (mgrs.GridDesignation{}) {
		grid := world.Terra.GetGrid(world.SelectedGrid)
		grid.Biome = t
		return "Set biome successfully."
	} else {
		return "No grid selected. Use: map select <grid>"
	}
}

func gridInfo(s string) string {
	b := world.Terra.GetGrid(world.SelectedGrid).Biome
	return "Grid " + world.SelectedGrid.ToString() + ": Biome " + strconv.Itoa(b)
}
