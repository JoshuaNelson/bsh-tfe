package draw

import (
	"github.com/nsf/termbox-go"
	"bsh-tfe/world"
)

var ConsolePrompt string = ">"
var frontendSquareSize = 10

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
	World(x, y+1, boxWidth, boxHeight, 0, 0)
	Text(x+1, y+3+boxHeight, msgBuf)
}

func Text(x, y int, text string) {
	for i, ch := range text {
		Cell(x + i, y, ch)
	}
}

func World(x, y, width, height, xCoord, yCoord int) {
	Box(x, y, width, height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			Terrain(x+1+i*2, y+j+1, i+xCoord, j+yCoord)
		}
	}
}

func Terrain(x, y, xWrldCoord, yWrldCoord int) {
	if xWrldCoord > world.Size || yWrldCoord > world.Size {
		return
	}

	var ch rune = 0x0000
	t := world.GameMap.Grid(xWrldCoord, yWrldCoord)

	switch t.Biome {
	case world.TERRAIN_ARID:
		ch = '_'
	case world.TERRAIN_FOREST:
		ch = 'f'
	}

	if t == world.Selected {
		CellSelected(x, y, ch)
		CellSelected(x+1, y, 0x0000)
	} else {
		Cell(x, y, ch)
		Cell(x+1, y, 0x0000)
	}
}
