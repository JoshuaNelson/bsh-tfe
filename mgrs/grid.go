package mgrs // Military Grid Reference System

import (
	"fmt"
	"logger"
	"strings"
)

type GridDesignation struct {
	GZD GridZoneDesignation
	GSD GridSquareDesignation
	SDC SixDigitCoordinate
}

func StringToGridDesignation(s string) (GridDesignation, error) {
	e := fmt.Errorf("Invalid MGRS.")
	gridWord := strings.Split(s, " ")

	// Build the GridZoneDesignation
	var gzd GridZoneDesignation
	if len(gridWord) >= 4 {
		var err error = nil
		gzd, err = stringToGZD(gridWord[0])
		if err != nil {
			logger.Debug("Invalid GZD")
			return GridDesignation{}, err
		}
		gridWord = gridWord[1:]
	} else if currentGZD != (GridZoneDesignation{}) {
		gzd = currentGZD
	} else {
		return GridDesignation{}, e
	}

	// Build the GridSquareDesignation
	var gsd GridSquareDesignation
	if len(gridWord) >= 3 {
		var err error = nil
		gsd, err = stringToSID(gridWord[0])
		if err != nil {
			logger.Debug("Invalid square identifier")
			return GridDesignation{}, err
		}
		gridWord = gridWord[1:]
	} else if currentGSD != (GridSquareDesignation{}) {
		gsd = currentGSD
	} else {
		return GridDesignation{}, e
	}

	if !gsd.validEasting(gzd) {
		logger.Debug("Invalid easting square identifier coordinate.")
		return GridDesignation{}, e
	} else if !gsd.validNorthing(gzd) {
		logger.Debug("Invalid northing square identifier coordinate.")
		return GridDesignation{}, e
	}

	// Build the SixDigitCoordinate
	var sdc SixDigitCoordinate
	sdc, err := stringToSDC(gridWord[0], gridWord[1])
	if err != nil {
		return GridDesignation{}, err
	}

	return GridDesignation{gzd, gsd, sdc}, nil
}

func (g GridDesignation) ToString() (string) {
	var gridBldr strings.Builder
	gridBldr.WriteString(g.GZD.ToString())
	gridBldr.WriteString(" ")
	gridBldr.WriteString(g.GSD.ToString())
	gridBldr.WriteString(" ")
	gridBldr.WriteString(g.SDC.ToString())
	return gridBldr.String()
}
