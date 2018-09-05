package main

import (
	"bsh-tfe/control"
	"bsh-tfe/view"
	"bsh-tfe/world"
	"github.com/nsf/termbox-go"
	"logger"
)

func check(e error) { if e != nil { panic(e) } }

func main() {
	logger.Init()
	defer logger.Close()
	logger.Debug("Initializing frontend.")

	check(termbox.Init())
	defer termbox.Close()

	control.Init()
	world.Init()

	logger.Debug("Initializing display.")
	termbox.SetInputMode(termbox.InputEsc) // | termbox.InputMouse)
	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	draw.Frontend()
	termbox.Flush()

	logger.Debug("Polling for events.")
loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyCtrlX {
				break loop
			}
			control.InputMode().EventHandler(event)

		case termbox.EventError:
			panic(event.Err)
		}

		termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

		draw.Frontend()
		termbox.Flush()
	}
}
