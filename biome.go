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
)

var BiomeArid      Biome = Biome{0, 0x2303, color256(235), color256(137)}
var BiomeForest    Biome = Biome{1, 0x234B, color256(22),  color256(34)}
var BiomeGrass     Biome = Biome{2, 0x2304, color256(22),  color256(40)}
var BiomeRock      Biome = Biome{3, 0x2591, color256(249), color256(245)}
var BiomeWater     Biome = Biome{4, 0x2303, color256(255), color256(39)}
var BiomeSnow      Biome = Biome{5, 0x0000, color256(15),  color256(15)}
var BiomeSand      Biome = Biome{6, 0x0000, color256(238), color256(228)}
var BiomeDeepwater Biome = Biome{7, 0x2303, color256(255), color256(27)}

type Biome struct {
	code int
	graphic rune
	foreground termbox.Attribute
	background termbox.Attribute
}

func (b *Biome) style() (rune, termbox.Attribute, termbox.Attribute) {
	return b.graphic, b.background, b.foreground
}
