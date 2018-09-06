package draw

import (
	"github.com/nsf/termbox-go"
	"bsh-tfe/control"
	"bsh-tfe/mgrs"
	"bsh-tfe/world"
//	"time"
//	"strconv"
)

var ConsolePrompt string = ">"
var frontendSquareSize = 40

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
	Text(x, y, ConsolePrompt)
	Text(x + len(ConsolePrompt) + 1, y, consoleBuf)
}

func Frontend() {
	x, y := 1, 1
	boxWidth, boxHeight:= frontendSquareSize, frontendSquareSize
	Console(x, y, control.CommandLine.Buffer.String())
	World(x, y+1, boxWidth, boxHeight, control.ViewGridDesig)
	//Text(x+1, y+3+boxHeight, msgBuf)
}

func Text(x, y int, text string) {
	for i, ch := range text {
		Cell(x + i, y, ch)
	}
}

func World(x, y, width, height int, g mgrs.GridDesignation) {
	//startTime := time.Now()
	//Box(x, y, width, height)
	DBox(x, y, width, height)
	// This is kind of hacky clean up later when we separate cursor and view
	g = g.AdjustEasting(0 - (frontendSquareSize/2))
	g = g.AdjustNorthing(0 - (frontendSquareSize/2))
	curGridE := g
	curGridEN := g

	mapX := x+1
	mapY := y+height
	// draw map left
	var eastingGrid *world.Grid
	var northingGrid *world.Grid
	var adjustEasting bool = false
	for e := 0; e < width; e++ {
		if eastingGrid == nil {
			// Just starting. Do lookup, set eastingGrid.
			adjustEasting = true
			curGridE = g.AdjustEasting(e)
			eastingGrid = world.SelectedPlanet.GetGrid(curGridE)
		} else {
			if eastingGrid.East != nil {
				// East grid shortcut is available, avoid lookup
				eastingGrid = eastingGrid.East
			} else {
				// Do lookup, save shortcut for future
				adjustEasting = true
				curGridE = g.AdjustEasting(e)
				eastingGrid.East = world.SelectedPlanet.GetGrid(curGridE)
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
						curGridE = g.AdjustEasting(e)
					}
					curGridEN = curGridE.AdjustNorthing(n)
					northingGrid.North = world.SelectedPlanet.GetGrid(curGridEN)
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
	//Text(mapX, mapY+2, durationString)
}

func Terrain(x, y int, g *world.Grid) {
	if g == nil {
		return
	}

	var ch1 rune = 0x0000
	var ch2 rune = 0x0000

	fgColor     := Color256(0)
	bgColor     := Color256(0)
	//fgSelColor  := Color256(1)
	bgSelColor  := Color256(15)
	cursorColor := Color256(1)

	switch g.Biome {
	case world.BIOME_ARID:
		ch2 = UpArrow
		fgColor = termbox.ColorBlack
		bgColor = Color256(137)
	case world.BIOME_FOREST:
		ch2 = Delta
		fgColor = Color256(22) // Dark Green Trees
		bgColor = Color256(34) // Green Tile
	case world.BIOME_GRASSLAND:
		ch2 = DownArrow
		fgColor = Color256(22) // Dark Green Trees
		bgColor = Color256(40) // Light Green
	case world.BIOME_ROCKY:
		ch2 = Gravel
		fgColor = Color256(249)
		bgColor = Color256(245)
	case world.BIOME_WATER:
		ch2 = '^'
		fgColor = Color256(81)
		bgColor = Color256(27)
	}

	if g == control.SelectedGrid {
		//fgColor = fgSelColor
		bgColor = bgSelColor
	}

	if g == control.CursorGrid {
		ch1 = 0x27EA
		ch2 = 0x27EB
		fgColor = cursorColor
	}

	termbox.SetCell(x,   y, ch1, fgColor, bgColor)
	termbox.SetCell(x+1, y, ch2, fgColor, bgColor)
}
