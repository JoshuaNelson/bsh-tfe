package main

import (
	"github.com/nsf/termbox-go"
	"logger"
)

var teamColor int = 33

func check(e error) { if e != nil { panic(e) } }

func main() {
	logger.Init()
	defer logger.Close()
	logger.Debug("Initializing frontend.")

	check(termbox.Init())
	defer termbox.Close()

	Control = initControl()

	logger.Debug("Initializing display.")
	termbox.SetInputMode(termbox.InputEsc) // | termbox.InputMouse)
	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	drawFrontend()
	termbox.Flush()

	logger.Debug("Polling for events.")
loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyCtrlX {
				break loop
			}
			Control.inputMode.EventHandler(event)

		case termbox.EventError:
			panic(event.Err)
		}

		termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
		drawFrontend()
		drawText(1, 1+40, "Cursor: " + Control.gameMap.curGridDes.ToString())
		drawText(1, 2+40, "Select: " + Control.gameMap.selGridDes.ToString())
		termbox.Flush()
	}
}
