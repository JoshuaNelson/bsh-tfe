package main

import (
	"fmt"
	"github.com/aquilax/go-perlin"
	"logger"
	"strconv"
	"strings"
)

/*
 * GRID GRID GRID
 */
type Grid struct {
	Biome int
	North *Grid
	East  *Grid
	Unit Unit
}

func (g *Grid) setBiome(b int) {
	g.Biome = b
}

func (g *Grid) newGrid(n float64) {
	switch {
	case n < BiomeLD:
		g.setBiome(BiomeDeepwater)
	case n >= BiomeLD && n < BiomeL0:
		g.setBiome(BiomeWater)
	case n >= BiomeL0 && n < BiomeL1:
		g.setBiome(BiomeSand)
	case n >= BiomeL1 && n < BiomeL2:
		g.setBiome(BiomeArid)
	case n >= BiomeL2 && n < BiomeL3:
		g.setBiome(BiomeGrass)
	case n >= BiomeL3 && n < BiomeL4:
		g.setBiome(BiomeForest)
	case n >= BiomeL4 && n < BiomeL5:
		g.setBiome(BiomeRock)
	case n >= BiomeL5:
		g.setBiome(BiomeSnow)
	}
}



/*
 * GRIDSQUARE GRIDSQUARE GRIDSQUARE
 */
func initGridSquare(gsd GridSquareDesignation) *GridSquare {
	var gs GridSquare
	gs.Grid = make(map[SixDigitCoordinate]*Grid)

	logger.Debug("Spooling up Perlin noise generator. Cover your ears.")
	var seed int64 = 65

	var scale float64 = 25
	p := perlin.NewPerlin(2.1, 2.2, 3, seed)
	//p := perlin.NewPerlin(2, 2, 3, seed)

	// TODO Use scaled ints across world, not just each square

	logger.Debug("Generating new Grid Square, %s.", gsd.ToString())
	for x := 0; x <= GridSquareSize; x++ {
		for y := 0; y <= GridSquareSize; y++ {
			sdc := SixDigitCoordinate{x, y}
			var g Grid
			g.newGrid(p.Noise2D(float64(x)/scale, float64(y)/scale))
			gs.Grid[sdc] = &g
		}
	}

	return &gs
}

type GridSquare struct {
	Grid map[SixDigitCoordinate]*Grid
}

func (gs *GridSquare) getGrid(g GridDesignation) *Grid {
	return gs.Grid[g.SDC]
}

/*
 * GRIDZONE GRIDZONE GRIDZONE
 */
func initGridZone(gzd GridZoneDesignation) *GridZone {
	var gz GridZone
	gz.GridSquare = make(map[GridSquareDesignation]*GridSquare)

	logger.Debug("Generating new Grid Zone: %s.", gzd.ToString())
	return &gz
}

type GridZone struct {
	GridSquare map[GridSquareDesignation]*GridSquare
}

func (gz *GridZone) getGrid(g GridDesignation) *Grid {
	_, gsInitialized := gz.GridSquare[g.GSD]
	if !gsInitialized {
		gz.GridSquare[g.GSD] = initGridSquare(g.GSD)
	}
	return gz.GridSquare[g.GSD].getGrid(g)
}

/*
 * PLANET PLANET PLANET
 */
type planet struct {
	name string
	gridZone map[GridZoneDesignation]*GridZone
}

func (p *planet) getGrid(g GridDesignation) *Grid {
	_, gzInitialized := p.gridZone[g.GZD]
	if !gzInitialized {
		p.gridZone[g.GZD] = initGridZone(g.GZD)
	}
	return p.gridZone[g.GZD].getGrid(g)
}

/*
 * SIXDIGITCOORDINATE SIXDIGITCOORDINATE SIXDIGITCOORDINATE
 */
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

func (sdc SixDigitCoordinate) ToString() (string) {
	var sdcBldr strings.Builder
	sdcBldr.WriteString(fmt.Sprintf("%03d", sdc.Easting))
	sdcBldr.WriteString(" ")
	sdcBldr.WriteString(fmt.Sprintf("%03d", sdc.Northing))
	return sdcBldr.String()
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

/*
 * GRIDSQUAREDESIGNATION GRIDSQUAREDESIGNATION GRIDSQUAREDESIGNATION
 */
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
	return WrapRuneSlice(gsdSequence, startIdx, endIdx)
}

func (gsd GridSquareDesignation) getNorthingSequence() []rune {
	startIdx := gsd.northingSeqIdx
	endIdx := gsd.northingSeqIdx + gzdRatio[1]
	return WrapRuneSlice(gsdSequence, startIdx, endIdx)
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

/*
 * GRIDZONEDESIGNATION GRIDZONEDESIGNATION GRIDZONEDESIGNATION
 */
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

/*
 * GRIDDESIGNATION GRIDDESIGNATION GRIDDESIGNATION
 */
type GridDesignation struct {
	GZD GridZoneDesignation
	GSD GridSquareDesignation
	SDC SixDigitCoordinate
}

func (g GridDesignation) adjustEasting(x int) GridDesignation {
	if x == 0 { return g }

	carry := 0
	g.SDC, carry = g.SDC.adjustEasting(x)
	if carry != 0 { g.GSD, carry = g.GSD.adjustEasting(carry) }
	if carry != 0 { g.GZD, carry = g.GZD.adjustEasting(carry) }
	return g
}

func (g GridDesignation) adjustNorthing(y int) GridDesignation {
	if y == 0 { return g }

	carry := 0
	g.SDC, carry = g.SDC.adjustNorthing(y)
	if carry != 0 { g.GSD, carry = g.GSD.adjustNorthing(carry) }
	if carry != 0 { g.GZD, carry = g.GZD.adjustNorthing(carry) }
	return g
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
