package draw

import (
	"github.com/nsf/termbox-go"
	"bsh-tfe/world"
)

var ConsolePrompt string = ">"

func Box(x, y, height, width int) {
	// Draw corners
	Cell(x, y, 0x250C)
	Cell(x+width, y, 0x2510)
	Cell(x, y+height, 0x2514)
	Cell(x+width, y+height, 0x2518)

	// Draw top and bottom
	for i := 1; i < width; i++ {
		Cell(x+i, y, 0x2500)
		Cell(x+i, y+height, 0x2500)
	}

	// Draw left and right
	for i := 1; i < height; i++ {
		Cell(x, y+i, 0x2502)
		Cell(x+width, y+i, 0x2502)
	}
}

func Cell(x, y int, r rune) {
	termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
}

func Console(x, y int, consoleBuf string) {
	Text(x, y, ConsolePrompt)
	Text(x + len(ConsolePrompt) + 1, y, consoleBuf)
}

func Frontend(consoleBuf, msgBuf string) {
	x, y := 1, 1
	boxHeight, boxWidth := 30, 60
	Console(x, y, consoleBuf)
	World(x, y+1, boxHeight, boxWidth, 0, 0)
	Text(x+1, y+2, msgBuf)
}

func Text(x, y int, text string) {
	for i, ch := range text {
		Cell(x + i, y, ch)
	}
}

func World(x, y, height, width, xCoord, yCoord int) {
	Box(x, y, height, width)
	for i := 0; i < width-1; i++ {
		for j := 0; j < height-1; j++ {
			Terrain(x+i+1, y+j+1, i, j)
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
	case 0:
		ch = '.'
	}

	if ch != 0x0000 {
		Cell(x, y, ch)
	}
}
