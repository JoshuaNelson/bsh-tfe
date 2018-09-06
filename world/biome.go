package world

import (
	"github.com/nsf/termbox-go"
)

var BIOME_ARID int = 0
var BIOME_FOREST int = 1
var BIOME_GRASSLAND int = 2
var BIOME_ROCK int = 3
var BIOME_WATER int = 4
var BIOME_SNOW int = 5
var BIOME_SAND = 6
var BIOME_DEEPWATER = 7

type StyleFunction func(int) (termbox.Attribute)

func StyleBiome(b int, color StyleFunction) (termbox.Attribute, termbox.Attribute, rune, rune) {
	var ch1 rune = 0x0000
	var ch2 rune = 0x0000

	fgColor := color(0)
	bgColor := color(0)

	switch b {
	case BIOME_ARID:
		ch2 = 0x2303 // UpArrow
		fgColor = color(235) // Dark Gray
		bgColor = color(137) // Red sand
	case BIOME_FOREST:
		ch2 = 0x234B // Delta
		fgColor = color(22) // Dark Green Trees
		bgColor = color(34) // Green Tile
	case BIOME_GRASSLAND:
		ch2 = 0x2304 // DownArrow
		fgColor = color(22) // Dark Green Trees
		bgColor = color(40) // Light Green
	case BIOME_ROCK:
		ch2 = 0x2591 // Gravel
		fgColor = color(249) // Light Gray
		bgColor = color(245) // Gray
	case BIOME_WATER:
		ch2 = 0x2303
		fgColor = color(255) // Lightest Gray
		bgColor = color(39) // Medium Blue
	case BIOME_DEEPWATER:
		ch2 = 0x2303
		fgColor = color(255) // Medium Blue
		bgColor = color(27) // Dark Blue
	case BIOME_SAND:
		fgColor = color(238) // Dark Gray
		bgColor = color(228) // Bright sand
	case BIOME_SNOW:
		bgColor = color(15) // White
	}

	return fgColor, bgColor, ch1, ch2
}
