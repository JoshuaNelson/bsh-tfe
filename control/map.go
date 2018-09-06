package control

import (
	"github.com/nsf/termbox-go"
	"bsh-tfe/mgrs"
	"bsh-tfe/world"
	"logger"
)

var TerrainMap Map
var SelectedGrid mgrs.GridDesignation
var CursorGrid mgrs.GridDesignation

type Map struct {
	Planet *world.Planet
}

func (m Map) EventHandler(event termbox.Event) {
	switch event.Key {
	case termbox.KeyEsc:
		logger.Debug("Returning to command line.")
		inputMode = CommandLine
		return

	case termbox.KeyEnter:
		return
	case termbox.KeyArrowUp:
		logger.Debug("Arrow Up")
		SelectedGrid = SelectedGrid.AdjustNorthing(1)
		return
	case termbox.KeyArrowDown:
		logger.Debug("Arrow Down")
		SelectedGrid = SelectedGrid.AdjustNorthing(-1)
		return
	case termbox.KeyArrowLeft:
		logger.Debug("Arrow Left")
		SelectedGrid = SelectedGrid.AdjustEasting(-1)
		return
	case termbox.KeyArrowRight:
		logger.Debug("Arrow Right")
		SelectedGrid = SelectedGrid.AdjustEasting(1)
		return
	}

	return
}
