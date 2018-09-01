package mgrs // Military Grid Reference System

import (
	"bsh-tfe/util"
	"fmt"
	"strings"
)

var currentGSD GridSquareDesignation
var gsdSequence = []rune{'A','B','C','D','E','F','G','H','J','K','L','M','N','P','Q','R','S','T','U','V'}

type GridSquareDesignation struct {
	Easting rune
	Northing rune
}

func (sid GridSquareDesignation) validEasting(g GridZoneDesignation) (bool) {
	var seq[]rune
	startIdx := ((g.Easting-1)*gzdRatio[0]) % len(gsdSequence)
	endIdx := startIdx + gzdRatio[0]
	seq = util.WrapRuneSlice(gsdSequence, startIdx, endIdx)

	for _, ch := range seq {
		if sid.Easting == ch {
			return true
		}
	}
	return false
}

func (sid GridSquareDesignation) validNorthing(g GridZoneDesignation) (bool) {
	var seq[]rune
	var startIdx int = 0
	for idx, ch := range gzdNorthingSequence {
		if g.Northing == ch {
			startIdx = (idx*gzdRatio[1]) % len(gsdSequence)
			break
		}
	}

	endIdx := startIdx+gzdRatio[1]
	seq = util.WrapRuneSlice(gsdSequence, startIdx, endIdx)

	for _, ch := range seq {
		if sid.Northing == ch {
			return true
		}
	}
	return false
}

func stringToSID(gridWord string) (GridSquareDesignation, error) {
	var easting rune
	var northing rune
	e := fmt.Errorf("Invalid Square ID.")

	// Only 2 character coordinate is acceptable
	if len(gridWord) != 2 {
		return GridSquareDesignation{}, e
	}

	easting = rune(gridWord[0])
	northing = rune(gridWord[1])

	return GridSquareDesignation{easting, northing}, nil
}

func (sid GridSquareDesignation) ToString() (string) {
	var sidBldr strings.Builder
	sidBldr.WriteRune(sid.Easting)
	sidBldr.WriteRune(sid.Northing)
	return sidBldr.String()
}
