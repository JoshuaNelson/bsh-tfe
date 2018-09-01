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
