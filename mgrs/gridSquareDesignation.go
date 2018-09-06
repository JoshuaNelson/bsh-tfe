package mgrs // Military Grid Reference System

import (
	"bsh-tfe/util"
	"fmt"
	"logger"
	"strings"
)

var currentGSD *GridSquareDesignation
var gsdSequence = []rune{'A','B','C','D','E','F','G','H','J','K','L','M','N','P','Q','R','S','T','U','V'}
var GridSquareSize = 999

type GridSquareDesignation struct {
	Easting rune
	Northing rune
	eastingSeqIdx int // Index in gsdSequence starting valid eastings
	northingSeqIdx int // Index in gsdSequence starting valid northings
}

func (gsd GridSquareDesignation) adjustEasting(x int) (GridSquareDesignation, int) {
	carry := 0

	// Get the index value of current Easting value
	gzdIdx := -1
	for idx, ch := range gsd.getEastingSequence() {
		if gsd.Easting == ch {
			gzdIdx = idx
		}
	}

	//Calculate what new easting is
	gsdIdx := 0
	for idx, ch := range gsdSequence {
		if gsd.Easting == ch {
			gsdIdx = idx
		}
	}
	newEastingIdx := (gsdIdx + x) % len(gsdSequence)
	if newEastingIdx < 0 {
		newEastingIdx += len(gsdSequence)
	}
	gsd.Easting = gsdSequence[newEastingIdx]


	// Carry value is how many times we cross the GSD border
	relativeX := x + gzdIdx
	for {
		if relativeX > (gzdRatio[0] - 1) {
			carry++
			relativeX -= gzdRatio[0]
			gsd.eastingSeqIdx += gzdRatio[0]
			if gsd.eastingSeqIdx > len(gsdSequence) {
				gsd.eastingSeqIdx -= len(gsdSequence)
			}
		} else if relativeX < 0 {
			carry--
			relativeX += gzdRatio[0]
			gsd.eastingSeqIdx -= gzdRatio[0]
			if gsd.eastingSeqIdx < 0 {
				gsd.eastingSeqIdx += len(gsdSequence)
			}
		} else {
			break
		}
	}

	return gsd, carry
}

func (gsd GridSquareDesignation) adjustNorthing(y int) (GridSquareDesignation, int) {
	carry := 0

	// Get the index value of current Northing value
	gzdIdx := 0
	for idx, ch := range gsd.getNorthingSequence() {
		if gsd.Northing == ch {
			gzdIdx = idx
		}
	}

	//Calculate what new easting is
	gsdIdx := 0
	for idx, ch := range gsdSequence {
		if gsd.Northing == ch {
			gsdIdx = idx
		}
	}
	newNorthingIdx := (gsdIdx + y) % len(gsdSequence)
	if newNorthingIdx < 0 {
		newNorthingIdx += len(gsdSequence)
	}
	gsd.Northing = gsdSequence[newNorthingIdx]


	// Carry value is how many times we cross the GSD border
	relativeX := y + gzdIdx
	for {
		if relativeX > (gzdRatio[1] - 1) {
			carry++
			relativeX -= gzdRatio[1]
			gsd.northingSeqIdx += gzdRatio[1]
			if gsd.northingSeqIdx > len(gsdSequence) {
				gsd.northingSeqIdx -= len(gsdSequence)
			}
		} else if relativeX < 0 {
			carry--
			relativeX += gzdRatio[1]
			gsd.northingSeqIdx -= gzdRatio[1]
			if gsd.northingSeqIdx < 0 {
				gsd.northingSeqIdx += len(gsdSequence)
			}
		} else {
			break
		}
	}

	return gsd, carry
}

func (gsd GridSquareDesignation) getEastingSequence() []rune {
	startIdx := gsd.eastingSeqIdx
	endIdx := gsd.eastingSeqIdx + gzdRatio[0]
	return util.WrapRuneSlice(gsdSequence, startIdx, endIdx)
}

func (gsd GridSquareDesignation) getNorthingSequence() []rune {
	startIdx := gsd.northingSeqIdx
	endIdx := gsd.northingSeqIdx + gzdRatio[1]
	return util.WrapRuneSlice(gsdSequence, startIdx, endIdx)
}

func (gsd GridSquareDesignation) validEasting() (bool) {
	for _, ch := range gsd.getEastingSequence() {
		if gsd.Easting == ch {
			return true
		}
	}
	return false
}

func (gsd GridSquareDesignation) validNorthing() (bool) {
	for _, ch := range gsd.getNorthingSequence() {
		if gsd.Northing == ch {
			return true
		}
	}
	return false
}

func initEastingIdx(g GridZoneDesignation) int {
	return ((g.Easting-1)*gzdRatio[0]) % len(gsdSequence)
}

func initNorthingIdx(g GridZoneDesignation) int {
	for idx, ch := range gzdNorthingSequence {
		if g.Northing == ch {
			return (idx*gzdRatio[1]) % len(gsdSequence)
		}
	}
	return 0
}

func stringToGSD(gridWord string, gzd GridZoneDesignation) (GridSquareDesignation, error) {
	var gsd GridSquareDesignation
	e := fmt.Errorf("Invalid grid square designation.")

	// Only 2 character coordinate is acceptable
	if len(gridWord) != 2 {
		return GridSquareDesignation{}, e
	}

	// Build GSD
	gsd.Easting = rune(gridWord[0])
	gsd.Northing = rune(gridWord[1])
	gsd.eastingSeqIdx = initEastingIdx(gzd)
	gsd.northingSeqIdx = initNorthingIdx(gzd)

	if !gsd.validEasting() {
		logger.Debug("Invalid easting grid square designation.")
		return GridSquareDesignation{}, e
	} else if !gsd.validNorthing() {
		logger.Debug("Invalid northing grid square designation.")
		return GridSquareDesignation{}, e
	}

	return gsd, nil
}

func (gsd GridSquareDesignation) ToString() (string) {
	var gsdBldr strings.Builder
	gsdBldr.WriteRune(gsd.Easting)
	gsdBldr.WriteRune(gsd.Northing)
	return gsdBldr.String()
}
