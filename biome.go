package main

import (
	"github.com/nsf/termbox-go"
)
const (
	BiomeLD float64 = -0.4
	BiomeL0 float64 = -0.2
	BiomeL1 float64 = -0.12
	BiomeL2 float64 = 0
	BiomeL3 float64 = 0.2
	BiomeL4 float64 = 0.5
	BiomeL5 float64 = 0.7
	BiomeArid      int = 0
	BiomeForest    int = 1
	BiomeGrass     int = 2
	BiomeRock      int = 3
	BiomeWater     int = 4
	BiomeSnow      int = 5
	BiomeSand      int = 6
	BiomeDeepwater int = 7
)

func StyleBiome(b int) (termbox.Attribute, termbox.Attribute, rune, rune) {
	var ch1 rune = 0x0000
	var ch2 rune = 0x0000

	fgColor := color256(0)
	bgColor := color256(0)

	switch b {
	case BiomeArid:
		ch2 = 0x2303 // UpArrow
		fgColor = color256(235) // Dark Gray
		bgColor = color256(137) // Red sand
	case BiomeForest:
		ch2 = 0x234B // Delta
		fgColor = color256(22) // Dark Green Trees
		bgColor = color256(34) // Green Tile
	case BiomeGrass:
		ch2 = 0x2304 // DownArrow
		fgColor = color256(22) // Dark Green Trees
		bgColor = color256(40) // Light Green
	case BiomeRock:
		ch2 = 0x2591 // Gravel
		fgColor = color256(249) // Light Gray
		bgColor = color256(245) // Gray
	case BiomeWater:
		ch2 = 0x2303
		fgColor = color256(255) // Lightest Gray
		bgColor = color256(39) // Medium Blue
	case BiomeDeepwater:
		ch2 = 0x2303
		fgColor = color256(255) // Medium Blue
		bgColor = color256(27) // Dark Blue
	case BiomeSand:
		fgColor = color256(238) // Dark Gray
		bgColor = color256(228) // Bright sand
	case BiomeSnow:
		bgColor = color256(15) // White
	}

	return fgColor, bgColor, ch1, ch2
}
