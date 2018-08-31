package world

import (
	"logger"
)

var GameMap Map
var Size int = 64
var Selected *Terrain

var TERRAIN_ARID int = 0
var TERRAIN_FOREST int = 1

type Terrain struct {
	x int
	y int
	Biome int
	//units []Unit // TODO add unit per terrain
}

type Map struct {
	grid [][]Terrain
}

func (m *Map) Grid(x, y int) *Terrain {
	return &m.grid[x][y]
}

func (t *Terrain) SetBiome(b int) {
	t.Biome = b
}

func Init() {
	logger.Debug("Initializing game map.")
	if GameMap.grid != nil {
		return
	}
	for i := 0; i < Size; i++ {
		var t []Terrain
		for j := 0; j < Size; j++ {
			t = append(t, Terrain{i, j, TERRAIN_ARID})
		}
		GameMap.grid = append(GameMap.grid, t)
	}
}
