package control

import (
	"github.com/nsf/termbox-go"
)

var inputMode Input = CommandLine

type Input interface {
	EventHandler(termbox.Event)
}

func SetInputMode(mode Input) {
	inputMode = mode
}

func InputMode() Input {
	return inputMode
}
