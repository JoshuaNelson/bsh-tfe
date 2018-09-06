package mgrs // Military Grid Reference System

import (
	"fmt"
	"logger"
	"strconv"
	"strings"
)


var currentGZD *GridZoneDesignation
var gzdEastingSequence = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,
    20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,
    45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60}
var gzdNorthingSequence = []rune{'C','D','E','F','G','H','J','K','L','M','N','P','Q','R','S','T','U','V','W','X'}
var gzdRatio = [2]int{6,8}

type GridZoneDesignation struct {
	Easting int
	Northing rune
}

func (gzd GridZoneDesignation) adjustEasting(x int) (GridZoneDesignation, int) {
	eastingIdx := 0
	for idx, i := range gzdEastingSequence {
		if gzd.Easting == i {
			eastingIdx = idx
		}
	}
	eastingIdx += x
	if eastingIdx > len(gzdEastingSequence) || eastingIdx < 0 {
		return gzd, -1
	} else {
		gzd.Easting = gzdEastingSequence[eastingIdx]
	}
	return gzd, 0
}

func (gzd GridZoneDesignation) adjustNorthing(y int) (GridZoneDesignation, int) {
	northingIdx := 0
	for idx, ch := range gzdNorthingSequence {
		if gzd.Northing == ch {
			northingIdx = idx
		}
	}
	northingIdx += y
	if northingIdx > len(gzdNorthingSequence) || northingIdx < 0 {
		return gzd, -1
	} else {
		gzd.Northing = gzdNorthingSequence[northingIdx]
	}
	return gzd, 0
}

func stringToGZD(gridWord string) (GridZoneDesignation, error) {
	var easting int
	var northing rune
	e := fmt.Errorf("Invalid Grid Zone Designation.")

	// Should either be 3 or 2 characters
	if len(gridWord) > 4 || len(gridWord) < 2 {
		return GridZoneDesignation{}, e
	}

	// Get the easting value, should be 1 or 2 digit integer
	easting, err := strconv.Atoi(gridWord[:len(gridWord)-1])
	if err != nil {
		logger.Debug("Invalid GZD easting size")
		return GridZoneDesignation{}, e
	}

	gridWord = gridWord[1:]

	// Bounds checking for easting value (1 <= easting <= 60)
	if easting < gzdEastingSequence[0] ||
	    easting > gzdEastingSequence[len(gzdEastingSequence)-1] {
		logger.Debug("Invalid GZD easting value")
		return GridZoneDesignation{}, e
	}

	// Get the northing value, should be a single rune
	northing = rune(gridWord[0])
	northingValid := false
	// Check that northing value is in valid values
	for _, ch := range gzdNorthingSequence {
		if ch == northing {
			northingValid = true
		}
	}

	if !northingValid {
		logger.Debug("Invalid GZD northing value")
		return GridZoneDesignation{}, e
	}
	// We have valid easting and northing values, success
	return GridZoneDesignation{easting, northing}, nil
}

func (gzd GridZoneDesignation) ToString() (string) {
	var gzdBldr strings.Builder
	gzdBldr.WriteString(strconv.Itoa(gzd.Easting))
	gzdBldr.WriteRune(gzd.Northing)
	return gzdBldr.String()
}
