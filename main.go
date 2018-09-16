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

	Control.Init()
	defer Control.Close()

	logger.Debug("Initializing display.")
	termbox.SetInputMode(termbox.InputEsc) // | termbox.InputMouse)
	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	drawFrontend()
	termbox.Flush()

loop:
	for {
		select {
		case <-Control.viewUpdate:
			termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
			drawFrontend()
			termbox.Flush()

		case <-Control.quit:
			logger.Debug("Quitting main")
			break loop
		}
	}
}
