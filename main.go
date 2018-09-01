package main

import (
	"bsh-tfe/control"
	"bsh-tfe/view"
	"bsh-tfe/world"
	"github.com/nsf/termbox-go"
	"logger"

	"strings"
)

var wordExit string = "exit"

func inputEventHandler(e termbox.Event, input *strings.Builder) (
    *cmd.Command, string) {
	switch e.Key {
	// Exit program
	case termbox.KeyCtrlX:
		cmdExit, _ := cmd.Find(wordExit)
		return cmdExit, ""

	// Execute program
	case termbox.KeyEnter:
		defer input.Reset()
		return cmd.Find(strings.TrimSpace(input.String()))

	// Backspace
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		tmpInput := input.String()
		if len(tmpInput) > 0 {
			input.Reset()
			input.WriteString(tmpInput[0:len(tmpInput)-1])
		}
		return nil, ""

	// Space sets rune as NULL (0x0000) we want space (0x0020)
	case termbox.KeySpace:
		input.WriteRune(0x0020) // Space
		return nil, ""
	}

	// Other
	input.WriteRune(e.Ch)
	return nil, ""
}

func check(e error) { if e != nil { panic(e) } }

func main() {
	logger.Init()
	defer logger.Close()

	logger.Debug("Initializing frontend.")
	check(termbox.Init())
	defer termbox.Close()

	cmd.Init()
	world.Init()
	world.InitPlanet()

	logger.Debug("Initializing display.")
	var textIn strings.Builder
	var msg string
	termbox.SetInputMode(termbox.InputEsc) // | termbox.InputMouse)
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	draw.Frontend(textIn.String(), msg)
	termbox.Flush()

	logger.Debug("Polling for events.")
loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			command, s := inputEventHandler(event, &textIn)
			cmdExit, _ := cmd.Find(wordExit)
			if command == nil {
				break
			} else if command == cmdExit {
				logger.Debug("Exiting frontend.")
				break loop
			} else {
				msg = cmd.Run(command, s)
				logger.Debug("Running with msg %s", msg)
			}

		case termbox.EventError:
			panic(event.Err)
		}

		termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
		draw.Frontend(textIn.String(), msg)
		termbox.Flush()

		msg = ""
	}
}
