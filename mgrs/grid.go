package mgrs // Military Grid Reference System

import (
	"fmt"
	"logger"
	"strings"
)

type Grid struct {
	GZD gridZoneDesignation
	SID squareIdentifier
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

	// Build the squareIdentifier
	var sid squareIdentifier
	if len(gridWord) >= 3 {
		var err error = nil
		sid, err = stringToSID(gridWord[0])
		if err != nil {
			logger.Debug("Invalid square identifier")
			return Grid{}, err
		}
		gridWord = gridWord[1:]
	} else if currentSID != (squareIdentifier{}) {
		sid = currentSID
	} else {
		return Grid{}, e
	}

	if !sid.validEasting(gzd) {
		logger.Debug("Invalid easting square identifier coordinate.")
		return Grid{}, e
	} else if !sid.validNorthing(gzd) {
		logger.Debug("Invalid northing square identifier coordinate.")
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
