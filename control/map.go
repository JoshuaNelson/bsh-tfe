package control

import (
	"github.com/nsf/termbox-go"
	"bsh-tfe/mgrs"
	"bsh-tfe/world"
)

var TerrainMap Map

var SelectedGrid *world.Grid
var SelectedGridDesig mgrs.GridDesignation

var CursorGrid *world.Grid
var CursorGridDesig mgrs.GridDesignation

var ViewGridDesig mgrs.GridDesignation

type Map struct {
	Planet *world.Planet
}

func (m Map) EventHandler(event termbox.Event) {
	switch event.Key {
	case termbox.KeyEsc:
		inputMode = CommandLine
		return

	case termbox.KeyEnter:
		SelectedGrid = CursorGrid
		SelectedGridDesig = CursorGridDesig
		return

	case termbox.KeyArrowUp:
		CursorGridDesig = CursorGridDesig.AdjustNorthing(1)
		CursorGrid = world.SelectedPlanet.GetGrid(CursorGridDesig)
		return

	case termbox.KeyArrowDown:
		CursorGridDesig = CursorGridDesig.AdjustNorthing(-1)
		CursorGrid = world.SelectedPlanet.GetGrid(CursorGridDesig)
		return

	case termbox.KeyArrowLeft:
		CursorGridDesig = CursorGridDesig.AdjustEasting(-1)
		CursorGrid = world.SelectedPlanet.GetGrid(CursorGridDesig)
		return

	case termbox.KeyArrowRight:
		CursorGridDesig = CursorGridDesig.AdjustEasting(1)
		CursorGrid = world.SelectedPlanet.GetGrid(CursorGridDesig)
		return

	case termbox.KeyBackspace, termbox.KeyBackspace2:
		ViewGridDesig = SelectedGridDesig
		CursorGridDesig = SelectedGridDesig
		CursorGrid = SelectedGrid

	case termbox.KeyPgup:
		CursorGridDesig = CursorGridDesig.AdjustNorthing(10)
		CursorGrid = world.SelectedPlanet.GetGrid(CursorGridDesig)
		return

	case termbox.KeyPgdn:
		CursorGridDesig = CursorGridDesig.AdjustNorthing(-10)
		CursorGrid = world.SelectedPlanet.GetGrid(CursorGridDesig)
		return

	}

	// Focus map
	switch event.Ch {
	case 't':
		inputMode = CommandLine
		return

	case 'f':
		ViewGridDesig = CursorGridDesig
		return

	case 'k': // vim up
		CursorGridDesig = CursorGridDesig.AdjustNorthing(1)
		CursorGrid = world.SelectedPlanet.GetGrid(CursorGridDesig)
		return

	case 'j': // vim down
		CursorGridDesig = CursorGridDesig.AdjustNorthing(-1)
		CursorGrid = world.SelectedPlanet.GetGrid(CursorGridDesig)
		return

	case 'h': // vim left
		CursorGridDesig = CursorGridDesig.AdjustEasting(-1)
		CursorGrid = world.SelectedPlanet.GetGrid(CursorGridDesig)
		return

	case 'l': // vim right
		CursorGridDesig = CursorGridDesig.AdjustEasting(1)
		CursorGrid = world.SelectedPlanet.GetGrid(CursorGridDesig)
		return
	}

}
