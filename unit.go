package main

import (
	"github.com/nsf/termbox-go"
)

func spawn(s string) string {
	g := Control.gameMap.selGrid
	g.Unit = Tank{5}
	return "Spawned tank at " + Control.gameMap.selGridDes.ToString()
}

func moveUnit(sg *Grid, cg *Grid) {
	if sg.Unit != nil {
		cg.Unit = sg.Unit
		sg.Unit = nil
		// BUG move to same space deletes unit
	}
}

type Unit interface {
	move(*Grid)
	style() (rune, termbox.Attribute, termbox.Attribute)
}

type Tank struct {
	speed int
}

func (t Tank) move(g *Grid) {
	return
}

func (t Tank) style() (rune, termbox.Attribute, termbox.Attribute) {
	return 'T', color256(236), color256(teamColor)|termbox.AttrBold
}
