package world

import (
	"bsh-tfe/mgrs"
)

var BIOME_ARID int = 0
var BIOME_FOREST int = 1
var BIOME_GRASSLAND int = 2
var BIOME_ROCKY int = 4

func Init() {
	_, err := mgrs.StringToGridDesignation("1C AB 005 009")
	if err != nil {
		panic(err)
	}
}
