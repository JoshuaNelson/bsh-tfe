package draw

import (
	"github.com/nsf/termbox-go"
	"bsh-tfe/mgrs"
	"bsh-tfe/world"
)

var ConsolePrompt string = ">"
var frontendSquareSize = 20

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

func Frontend(consoleBuf, msgBuf string) {
	x, y := 1, 1
	boxWidth, boxHeight:= frontendSquareSize, frontendSquareSize
	Console(x, y, consoleBuf)
	World(x, y+1, boxWidth, boxHeight, world.SelectedGrid)
	Text(x+1, y+3+boxHeight, msgBuf)
}

func Text(x, y int, text string) {
	for i, ch := range text {
		Cell(x + i, y, ch)
	}
}

func World(x, y, width, height int, g mgrs.GridDesignation) {
	Box(x, y, width, height)
	origNorthing := g.SDC.Northing
	//origEasting :=  g.SDC.Easting

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			Terrain(x+1+i*2, y+j+1, g)
			g.SDC.Northing -= 1
		}
		g.SDC.Northing = origNorthing
		g.SDC.Easting -= 1
	}
}

func Terrain(x, y int, g mgrs.GridDesignation) {
	var ch1 rune = 0x0000
	var ch2 rune = 0x0000
	grid := world.Terra.GetGrid(g)

	if grid == nil {
		return
	}

	fgColor := termbox.ColorDefault
	bgColor := termbox.ColorDefault
	fgSelColor := termbox.ColorBlack
	bgSelColor := termbox.ColorWhite

	switch grid.Biome {
	case world.BIOME_ARID:
		ch2 = UpArrow
		fgColor = termbox.ColorBlack
		bgColor = termbox.Attribute(222)
	case world.BIOME_FOREST:
		ch2 = Delta
		fgColor = termbox.Attribute(23) // Dark Green Trees
		bgColor = termbox.Attribute(35) // Green Tile
	case world.BIOME_GRASSLAND:
		ch2 = DownArrow
		fgColor = termbox.Attribute(23) // Dark Green Trees
		bgColor = termbox.Attribute(41) // Light Green
	case world.BIOME_ROCKY:
		ch2 = Gravel
		fgColor = termbox.Attribute(250)
		bgColor = termbox.Attribute(246)
	}

	if g == world.SelectedGrid {
		ch1 = Machine
		termbox.SetCell(x,   y, ch1, fgSelColor, bgSelColor)
		termbox.SetCell(x+1, y, ch2, fgSelColor, bgSelColor)
	} else {
		termbox.SetCell(x,   y, ch1, fgColor, bgColor)
		termbox.SetCell(x+1, y, ch2, fgColor, bgColor)
	}
}
