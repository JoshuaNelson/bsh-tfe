package main

import (
	"strconv"
)

func WrapRuneSlice(s []rune, start int, end int) []rune {
	// Check for edge start and end values
	if start < 0 {
		panic("slice start out of range")
	} else if start > end {
		panic("invalid slice index: " + strconv.Itoa(start) + " > " + strconv.Itoa(end))
	}

	// Reduce start index to within slice
	for {
		if start < len(s) {
			break
		}
		start -= len(s)
		end -= len(s)
	}

	var rv []rune
	if end > len(s) {
		tempSlice := s[start:]
		rv = append(rv, tempSlice...)
		numWraps := (end / len(s)) - 1
		remainder := end % len(s)
		for i := 0; i < numWraps; i++ {
			rv = append(rv, s...)
		}
		tempSlice = s[:remainder]
		rv = append(rv, tempSlice...)
		return rv
	}

	return s[start:end]
}
