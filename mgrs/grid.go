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

func (g GridDesignation) AdjustEasting(x int) GridDesignation {
	carry := 0
	g.SDC, carry = g.SDC.adjustEasting(x)

	if carry != 0 {
		g.GSD, carry = g.GSD.adjustEasting(carry)
	}

	if carry != 0 {
		g.GZD, carry = g.GZD.adjustEasting(carry)
	}

	return g
}

func (g GridDesignation) AdjustNorthing(y int) GridDesignation {
	carry := 0
	g.SDC, carry = g.SDC.adjustNorthing(y)

	if carry != 0 {
		g.GSD, carry = g.GSD.adjustNorthing(carry)
	}

	if carry != 0 {
		g.GZD, carry = g.GZD.adjustNorthing(carry)
	}

	return g
}

func StringToGridDesignation(s string) (GridDesignation, error) {
	var gzd GridZoneDesignation
	var gsd GridSquareDesignation
	var sdc SixDigitCoordinate
	var err error = nil
	e := fmt.Errorf("Invalid grid designation.")
	gridWord := strings.Split(s, " ")

	if len(gridWord) < 4 {
		logger.Debug("Invalid grid designation.")
		return GridDesignation{}, e
	}

	// Build the GridZoneDesignation
	gzd, err = stringToGZD(gridWord[0])
	if err != nil {
		logger.Debug("Invalid grid zone designation.")
		return GridDesignation{}, err
	}

	// Build the GridSquareDesignation
	gsd, err = stringToGSD(gridWord[1], gzd)
	if err != nil {
		logger.Debug("Invalid grid square designation.")
		return GridDesignation{}, err
	}

	// Build the SixDigitCoordinate
	sdc, err = stringToSDC(gridWord[2], gridWord[3])
	if err != nil {
		logger.Debug("Invalid six digit coordinate.")
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
