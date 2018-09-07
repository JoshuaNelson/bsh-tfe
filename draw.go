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

func Box(x, y, height, width int) {
	// Draw corners
	Cell(x, y, 0x250C)
	Cell(x+1+width*2, y, 0x2510)
	Cell(x, y+1+height, 0x2514)
	Cell(x+1+width*2, y+1+height, 0x2518)

	// Draw top and bottom
	for i := 1; i < width*2; i+=2 {
		Cell(x+i, y, 0x2500)
		Cell(x+i+1, y, 0x2500)
		Cell(x+i, y+1+height, 0x2500)
		Cell(x+i+1, y+1+height, 0x2500)
	}

	// Draw left and right
	for i := 1; i < height+1; i++ {
		Cell(x, y+i, 0x2502)
		Cell(x+1+width*2, y+i, 0x2502)
	}
}

func DBox(x, y, height, width int) {
	// Draw corners
	Cell(x, y, 0x2554)
	Cell(x+1+width*2, y, 0x2557)
	Cell(x, y+1+height, 0x255A)
	Cell(x+1+width*2, y+1+height, 0x255D)

	// Draw top and bottom
	for i := 1; i < width*2; i+=2 {
		Cell(x+i, y, 0x2550)
		Cell(x+i+1, y, 0x2550)
		Cell(x+i, y+1+height, 0x2550)
		Cell(x+i+1, y+1+height, 0x2550)
	}

	// Draw left and right
	for i := 1; i < height+1; i++ {
		Cell(x, y+i, 0x2551)
		Cell(x+1+width*2, y+i, 0x2551)
	}
}

func Cell(x, y int, r rune) {
	termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
}

func CellSelected(x, y int, r rune) {
	termbox.SetCell(x, y, r, termbox.ColorBlack, termbox.ColorWhite)
}

func Console(x, y int, consoleBuf string) {
	drawText(x, y, ConsolePrompt)
	drawText(x + len(ConsolePrompt) + 1, y, consoleBuf)
}

func drawFrontend() {
	x, y := 1, 1
	boxWidth, boxHeight:= frontendSquareSize, frontendSquareSize
	Console(x, y, Control.cli.Buffer.String())
	World(x, y+1, boxWidth, boxHeight, Control.gameMap.mapGridDes)
	//drawText(x+1, y+3+boxHeight, msgBuf)
}

func drawText(x, y int, text string) {
	for i, ch := range text {
		Cell(x + i, y, ch)
	}
}

func World(x, y, width, height int, g GridDesignation) {
	//startTime := time.Now()
	//Box(x, y, width, height)
	DBox(x, y, width, height)
	// This is kind of hacky clean up later when we separate cursor and view
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
			Terrain(mapX+e*2, mapY-n, northingGrid)
		}
		northingGrid = nil
		adjustEasting = false
	}
	//durationSec := time.Since(startTime).Seconds()
	//durationString := strconv.FormatFloat(durationSec, 'f', -1, 64)
	//durationString += " sec"
	//drawText(mapX, mapY+2, durationString)
}

func Terrain(x, y int, g *Grid) {
	if g == nil {
		return
	}

	//fgSelColor  := color256(1)

	fgColor2, bgColor2, ch1, ch2 := StyleBiome(g.Biome)
	bgSelColor  := color256(15)
	cursorColor := color256(1)
	fgColor1    := fgColor2
	bgColor1    := bgColor2

	if g.Unit != nil {
		fgColor1 = color256(teamColor)|termbox.AttrBold
		bgColor1 = color256(236)
		ch1 = g.Unit.style()
	}

	if g == Control.gameMap.selGrid {
		//fgColor = fgSelColor
		bgColor1 = bgSelColor
		bgColor2 = bgSelColor
	}

	if g == Control.gameMap.curGrid {
		ch1 = 0x27EA
		ch2 = 0x27EB
		fgColor1 = cursorColor
		fgColor2 = cursorColor
	}

	termbox.SetCell(x,   y, ch1, fgColor1, bgColor1)
	termbox.SetCell(x+1, y, ch2, fgColor2, bgColor2)
}
