package mgrs // Military Grid Reference System

import (
	"fmt"
	"strconv"
	"strings"
)

type SixDigitCoordinate struct {
	Easting int
	Northing int
}

func (sdc SixDigitCoordinate) adjustEasting(x int) (SixDigitCoordinate, int) {
	carry := 0
	sdc.Easting += x
	for {
		if sdc.Easting > GridSquareSize {
			sdc.Easting -= (GridSquareSize+1)
			carry++
		} else if sdc.Easting < 0 {
			sdc.Easting += (GridSquareSize+1)
			carry--
		} else {
			break
		}
	}
	return sdc, carry
}

func (sdc SixDigitCoordinate) adjustNorthing(y int) (SixDigitCoordinate, int) {
	carry := 0
	sdc.Northing += y
	for {
		if sdc.Northing > GridSquareSize {
			sdc.Northing -= (GridSquareSize+1)
			carry++
		} else if sdc.Northing < 0 {
			sdc.Northing += (GridSquareSize+1)
			carry--
		} else {
			break
		}
	}
	return sdc, carry
}

func stringToSDC(eastingStr string, northingStr string) (SixDigitCoordinate, error) {
	e := fmt.Errorf("Invalid 6 digit coordinate.")
	easting, err := strconv.Atoi(eastingStr)
	if err != nil {
		return SixDigitCoordinate{}, e
	}
	northing, err := strconv.Atoi(northingStr)
	if err != nil {
		return SixDigitCoordinate{}, e
	}
	return SixDigitCoordinate{easting, northing}, nil
}

func (sdc SixDigitCoordinate) ToString() (string) {
	var sdcBldr strings.Builder
	sdcBldr.WriteString(fmt.Sprintf("%03d", sdc.Easting))
	sdcBldr.WriteString(" ")
	sdcBldr.WriteString(fmt.Sprintf("%03d", sdc.Northing))
	return sdcBldr.String()
}
