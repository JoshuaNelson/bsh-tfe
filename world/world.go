package world

import (
	"logger"
)

var GameMap Map
var Size int = 64

type Terrain struct {
	Biome int
	//units []Unit // TODO add unit per terrain
}

type Map struct {
	grid [][]Terrain
}

func (m *Map) Grid(x, y int) *Terrain {
	return &m.grid[x][y]
}

func Init() {
	logger.Debug("Initializing game map.")
	if GameMap.grid != nil {
		return
	}
	for i := 0; i < Size; i++ {
		var t []Terrain
		for j := 0; j < Size; j++ {
			t = append(t, Terrain{0})
		}
		GameMap.grid = append(GameMap.grid, t)
	}
}
