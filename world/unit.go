package world

import (
	"github.com/nsf/termbox-go"
)

var unitMap map[string]Unit

func InitUnitMap() {
	unitMap["tank"] = Tank{1}
}

type Unit interface {
	Move(*Grid)
	Style(StyleFunction) (termbox.Attribute, rune)
}

func UnitSpawn(s string) {
	var g *Grid
	return
	//g = control.SelectedGrid

	unit, mapped := unitMap[s]

	if mapped {
		g.Unit = unit
	}
	return
}

type Tank struct {
	speed int
}

func (t Tank) Move(g *Grid) {
	return
}

func (t Tank) Style(color StyleFunction) (termbox.Attribute, rune) {
	return color(0), 0x235E
}
