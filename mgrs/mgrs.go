package mgrs // Military Grid Reference System

import (
	"logger"

	"fmt"
	"strconv"
	"strings"
)

var currentGZD gridZoneDesignation
var currentSID squareID

var gzdEastingSequence = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,
    20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,
    45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60}
var gzdNorthingSequence = []rune{'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L',
    'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X'}

type gridZoneDesignation struct {
	easting int
	northing rune
}

func stringToGZD(gridWord string) (gridZoneDesignation, error) {
	var easting int
	var northing rune
	e := fmt.Errorf("Invalid Grid Zone Designation.")

	// Should either be 3 or 2 characters
	if len(gridWord) > 4 || len(gridWord) < 2 {
		return gridZoneDesignation{}, e
	}

	// Get the easting value, should be 1 or 2 digit integer
	easting, err := strconv.Atoi(gridWord[:len(gridWord)-1])
	if err != nil {
		logger.Debug("Invalid GZD easting size")
		return gridZoneDesignation{}, e
	}

	gridWord = gridWord[1:]

	// Bounds checking for easting value (1 <= easting <= 60)
	if easting < gzdEastingSequence[0] ||
	    easting > gzdEastingSequence[len(gzdEastingSequence)-1] {
		logger.Debug("Invalid GZD easting value")
		return gridZoneDesignation{}, e
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
		return gridZoneDesignation{}, e
	}
	// We have valid easting and northing values, success
	return gridZoneDesignation{easting, northing}, nil
}

func (gzd gridZoneDesignation) toString() (string) {
	var gzdBldr strings.Builder
	gzdBldr.WriteString(strconv.Itoa(gzd.easting))
	gzdBldr.WriteRune(gzd.northing)
	return gzdBldr.String()
}

type squareID struct {
	easting rune
	northing rune
}

func stringToSID(gridWord string) (squareID, error) {
	var easting rune
	var northing rune
	e := fmt.Errorf("Invalid Square ID.")

	// Only 2 character coordinate is acceptable
	if len(gridWord) != 2 {
		return squareID{}, e
	}

	easting = rune(gridWord[0])
	northing = rune(gridWord[1])

	return squareID{easting, northing}, nil
}

func (sid squareID) toString() (string) {
	var sidBldr strings.Builder
	sidBldr.WriteRune(sid.easting)
	sidBldr.WriteRune(sid.northing)
	return sidBldr.String()
}

type sixDigitCoordinate struct {
	easting int
	northing int
}

func stringToSDC(s_easting string, s_northing string) (sixDigitCoordinate, error) {
	e := fmt.Errorf("Invalid 6 digit coordinate.")
	easting, err := strconv.Atoi(s_easting)
	if err != nil {
		return sixDigitCoordinate{}, e
	}
	northing, err := strconv.Atoi(s_northing)
	if err != nil {
		return sixDigitCoordinate{}, e
	}
	return sixDigitCoordinate{easting, northing}, nil
}

func (sdc sixDigitCoordinate) toString() (string) {
	var sdcBldr strings.Builder
	sdcBldr.WriteString(strconv.Itoa(sdc.easting))
	sdcBldr.WriteString(" ")
	sdcBldr.WriteString(strconv.Itoa(sdc.northing))
	return sdcBldr.String()
}

type Grid struct {
	GZD gridZoneDesignation
	SID squareID
	coord sixDigitCoordinate
	selected bool
}

func StringToGrid(s string) (Grid, error) {
	e := fmt.Errorf("Invalid MGRS.")
	gridWord := strings.Split(s, " ")

	// Build the gridZoneDesignation
	var gzd gridZoneDesignation
	if len(gridWord) >= 4 {
		var err error = nil
		gzd, err = stringToGZD(gridWord[0])
		if err != nil {
			logger.Debug("Invalid GZD")
			return Grid{}, err
		}
		gridWord = gridWord[1:]
	} else if currentGZD != (gridZoneDesignation{}) {
		gzd = currentGZD
	} else {
		return Grid{}, e
	}

	// Build the squareID
	var sid squareID
	if len(gridWord) >= 3 {
		var err error = nil
		sid, err = stringToSID(gridWord[0])
		if err != nil {
			logger.Debug("Invalid squareID")
			return Grid{}, err
		}
		gridWord = gridWord[1:]
	} else if currentSID != (squareID{}) {
		sid = currentSID
	} else {
		return Grid{}, e
	}

	// Build the sixDigitCoordinate
	var sdc sixDigitCoordinate
	sdc, err := stringToSDC(gridWord[0], gridWord[1])
	if err != nil {
		return Grid{}, err
	}

	return Grid{gzd, sid, sdc, false}, nil
}

func (g Grid) ToString() (string) {
	var gridBldr strings.Builder
	gridBldr.WriteString(g.GZD.toString())
	gridBldr.WriteString(" ")
	gridBldr.WriteString(g.SID.toString())
	gridBldr.WriteString(" ")
	gridBldr.WriteString(g.coord.toString())
	return gridBldr.String()
}
