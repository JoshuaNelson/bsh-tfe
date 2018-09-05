package control

import (
	"github.com/nsf/termbox-go"
)

type Map struct {
}

func (m Map) EventHandler(event termbox.Event) {
	switch event.Key {
	case termbox.KeyEsc:
		inputMode = CommandLine
		return

	case termbox.KeyEnter:
	case termbox.KeyArrowUp:
	case termbox.KeyArrowDown:
	case termbox.KeyArrowLeft:
	case termbox.KeyArrowRight:
	}
}
