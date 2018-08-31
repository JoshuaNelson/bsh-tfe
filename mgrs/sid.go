package mgrs // Military Grid Reference System

import (
	"bsh-tfe/util"
	"fmt"
	"strings"
)

var currentSID squareIdentifier
var sidSequence = []rune{'A','B','C','D','E','F','G','H','J','K','L','M','N','P','Q','R','S','T','U','V'}

type squareIdentifier struct {
	easting rune
	northing rune
}

func (sid squareIdentifier) validEasting(g gridZoneDesignation) (bool) {
	var seq[]rune
	startIdx := ((g.easting-1)*gzdRatio[0]) % len(sidSequence)
	endIdx := startIdx + gzdRatio[0]
	seq = util.WrapRuneSlice(sidSequence, startIdx, endIdx)

	for _, ch := range seq {
		if sid.easting == ch {
			return true
		}
	}
	return false
}

func (sid squareIdentifier) validNorthing(g gridZoneDesignation) (bool) {
	var seq[]rune
	var startIdx int = 0
	for idx, ch := range gzdNorthingSequence {
		if g.northing == ch {
			startIdx = (idx*gzdRatio[1]) % len(sidSequence)
			break
		}
	}

	endIdx := startIdx+gzdRatio[1]
	seq = util.WrapRuneSlice(sidSequence, startIdx, endIdx)

	for _, ch := range seq {
		if sid.northing == ch {
			return true
		}
	}
	return false
}

func stringToSID(gridWord string) (squareIdentifier, error) {
	var easting rune
	var northing rune
	e := fmt.Errorf("Invalid Square ID.")

	// Only 2 character coordinate is acceptable
	if len(gridWord) != 2 {
		return squareIdentifier{}, e
	}

	easting = rune(gridWord[0])
	northing = rune(gridWord[1])

	return squareIdentifier{easting, northing}, nil
}

func (sid squareIdentifier) toString() (string) {
	var sidBldr strings.Builder
	sidBldr.WriteRune(sid.easting)
	sidBldr.WriteRune(sid.northing)
	return sidBldr.String()
}
