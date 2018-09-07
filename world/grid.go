package world

var BiomeLD float64 = -0.4
var BiomeL0 float64 = -0.2
var BiomeL1 float64 = -0.12
var BiomeL2 float64 = 0
var BiomeL3 float64 = 0.2
var BiomeL4 float64 = 0.5
var BiomeL5 float64 = 0.7

type Grid struct {
	Biome int
	North *Grid
	//South *Grid
	East  *Grid
	//West  *Grid
	Unit Unit
}

func (g *Grid) setBiome(b int) {
	g.Biome = b
}

func (g *Grid) newGrid(n float64) {
	switch {
	case n < BiomeLD:
		g.setBiome(BIOME_DEEPWATER)
	case n >= BiomeLD && n < BiomeL0:
		g.setBiome(BIOME_WATER)
	case n >= BiomeL0 && n < BiomeL1:
		g.setBiome(BIOME_SAND)
	case n >= BiomeL1 && n < BiomeL2:
		g.setBiome(BIOME_ARID)
	case n >= BiomeL2 && n < BiomeL3:
		g.setBiome(BIOME_GRASSLAND)
	case n >= BiomeL3 && n < BiomeL4:
		g.setBiome(BIOME_FOREST)
	case n >= BiomeL4 && n < BiomeL5:
		g.setBiome(BIOME_ROCK)
	case n >= BiomeL5:
		g.setBiome(BIOME_SNOW)
	}
}
