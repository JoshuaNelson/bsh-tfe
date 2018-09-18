package main

import (
	"github.com/nsf/termbox-go"
//	"time"
//	"strconv"
)

var ConsolePrompt string = ">"
var frontendSquareSize = 40

func color256(c int) termbox.Attribute {
	if c < 0 {
		c = 0
	} else if c > 255 {
		c = 255
	}

	return termbox.Attribute(c+1)
}

func Cell(x, y int, r rune) {
	termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
}

func Console(x, y int, consoleBuf string) {
	if Control.inputMode != Control.cli {
		return
	}
	drawText(x, y, ConsolePrompt)
	drawText(x + len(ConsolePrompt) + 1, y, consoleBuf)
}

func drawFrontend() {
	x, y := 0, 0
	boxWidth, boxHeight:= frontendSquareSize, frontendSquareSize
	Console(x, y, Control.cli.Buffer.String())
	World(x, y, boxWidth, boxHeight, Control.gameMap.focGridDes)
	drawText(1, 1+frontendSquareSize, "Cursor: " + Control.gameMap.curGridDes.ToString())
}

func drawText(x, y int, text string) {
	for i, ch := range text {
		Cell(x + i, y, ch)
	}
}

func World(x, y, width, height int, g GridDesignation) {
	//startTime := time.Now()
	g = g.adjustEasting(0 - (frontendSquareSize/2))
	g = g.adjustNorthing(0 - (frontendSquareSize/2))
	curGridE := g
	curGridEN := g

	mapX := x+1
	mapY := y+height
	// draw map left
	var eastingGrid *Grid
	var northingGrid *Grid
	var adjustEasting bool = false
	for e := 0; e < width; e++ {
		if eastingGrid == nil {
			// Just starting. Do lookup, set eastingGrid.
			adjustEasting = true
			curGridE = g.adjustEasting(e)
			eastingGrid = Control.getGrid(curGridE)
		} else {
			if eastingGrid.East != nil {
				// East grid shortcut is available, avoid lookup
				eastingGrid = eastingGrid.East
			} else {
				// Do lookup, save shortcut for future
				adjustEasting = true
				curGridE = g.adjustEasting(e)
				eastingGrid.East = Control.getGrid(curGridE)
				eastingGrid = eastingGrid.East
			}
		}
		// draw map up
		for n := 0; n < height; n++ {
			if northingGrid == nil {
				northingGrid = eastingGrid
			} else {
				if northingGrid.North != nil {
					// North shortcut is available
					northingGrid = northingGrid.North
				} else {
					// Do lookup, save shortcut
					if !adjustEasting {
						adjustEasting = true
						curGridE = g.adjustEasting(e)
					}
					curGridEN = curGridE.adjustNorthing(n)
					northingGrid.North = Control.getGrid(curGridEN)
					northingGrid = northingGrid.North
				}
			}
			if Control.gameMap.curGrid == northingGrid {
				defer DrawCursor(mapX+e*2, mapY-n, northingGrid)
			} else {
				DrawGridSquare(mapX+e*2, mapY-n, northingGrid)
			}
		}
		northingGrid = nil
		adjustEasting = false
	}
	//durationSec := time.Since(startTime).Seconds()
	//durationString := strconv.FormatFloat(durationSec, 'f', -1, 64)
	//durationString += " sec"
	//drawText(mapX, mapY+2, durationString)
}

func DrawGridSquare(x, y int, g *Grid) {
	if g == nil { return }

	biomeGraphic, biomeBackground, biomeForeground := g.Biome.style()
	unitBackground, unitForeground := biomeBackground, biomeForeground
	var unitGraphic rune = 0x0000

	if g.Unit != nil {
		unitGraphic, unitBackground, unitForeground = g.Unit.style()
	}

	unitBackground = unitBackground
	termbox.SetCell(x,   y, unitGraphic,  unitForeground,  biomeBackground)
	termbox.SetCell(x+1, y, biomeGraphic, biomeForeground, biomeBackground)
}

func DrawCursor(x, y int, g *Grid) {
	if g == nil { return }

	biomeGraphic, biomeBackground, biomeForeground := g.Biome.style()
	unitBackground, unitForeground := biomeBackground, biomeForeground
	var unitGraphic rune = 0x0000

	if g.Unit != nil {
		unitGraphic, unitBackground, unitForeground = g.Unit.style()
	}

	unitBackground = unitBackground
	termbox.SetCell(x,   y, unitGraphic,  unitForeground,  biomeBackground | termbox.AttrBold)
	termbox.SetCell(x+1, y, biomeGraphic, biomeForeground, biomeBackground)
}
