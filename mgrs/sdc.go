package mgrs // Military Grid Reference System

import (
	"fmt"
	"strconv"
	"strings"
)

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
	sdcBldr.WriteString(fmt.Sprintf("%03d", sdc.easting))
	sdcBldr.WriteString(" ")
	sdcBldr.WriteString(fmt.Sprintf("%03d", sdc.northing))
	return sdcBldr.String()
}
