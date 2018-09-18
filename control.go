package main

import (
	"github.com/nsf/termbox-go"
	"logger"
	"strings"
)

var Control Controller

/*
 * CONTROLLER
 */
type Controller struct {
	inputMode  Input
	cli        *CommandLine
	gameMap    *Map
	viewUpdate chan int
	quit       chan int
}

func (ctl* Controller) Init() {
	ctl.cli  = &CommandLine{nil, &strings.Builder{}}
	ctl.cli.initCommands()

	ctl.gameMap = &Map{}
	ctl.gameMap.Init()

	ctl.viewUpdate = make(chan int)
	ctl.quit = make(chan int)

	ctl.setInputMode(ctl.gameMap)
	go eventListener()
}

func (ctl* Controller) Close() {
	close(ctl.viewUpdate)
	close(ctl.quit)
}

func (ctl *Controller) getGrid(g GridDesignation) *Grid {
	return ctl.planet().getGrid(g)
}

func (ctl *Controller) planet() *planet {
	return ctl.gameMap.planet
}

func (ctl *Controller) setInputMode(mode Input) {
	ctl.inputMode = mode
}

func (ctl *Controller) Draw() {
	ctl.viewUpdate <- 0
}

func (ctl *Controller) Quit() {
	logger.Debug("Calling controller quit")
	ctl.quit <- 0
}

/*
 * INPUT
 */
type Input interface {
	EventHandler(termbox.Event)
}

/*
 * CLI CLI CLI
 */
type CommandLine struct {
	CmdTree *Command
	Buffer  *strings.Builder
}

func (cli *CommandLine) EventHandler(event termbox.Event) {
	switch event.Key {
	case termbox.KeyEnter:
		defer cli.Buffer.Reset()
		cmd, msg := cli.find(cli.Buffer.String())
		logger.Debug("Executing buffer: %s", cli.Buffer.String())
		if cmd != nil {
			logger.Debug("Running with msg %s", msg)
			cmd.Run(msg)
			Control.Draw()
		}
		Control.setInputMode(Control.gameMap)
		return

	case termbox.KeyBackspace, termbox.KeyBackspace2:
		tmpInput := cli.Buffer.String()
		if len(tmpInput) > 0 {
			cli.Buffer.Reset()
			cli.Buffer.WriteString(tmpInput[0:len(tmpInput)-1])
			Control.Draw()
		}
		return

	case termbox.KeySpace:
		cli.Buffer.WriteRune(0x0020) // Space
		Control.Draw()
		return

	case termbox.KeyEsc:
		Control.setInputMode(Control.gameMap)
		return

	case termbox.KeyArrowUp, termbox.KeyArrowDown, termbox.KeyArrowRight,
	    termbox.KeyArrowLeft:
		Control.setInputMode(Control.gameMap)
		Control.inputMode.EventHandler(event)
		return
	}

	cli.Buffer.WriteRune(event.Ch)
	Control.Draw()
	return
}

func (cli *CommandLine) find(s string) (*Command, string) {
	var cmdNode *Command = cli.CmdTree
	commandSplit := strings.Split(s, " ")

	if cmdNode == nil {
		return nil, ""
	}

	cmdIdx := 0
	for _, word := range commandSplit {
		foundWord := false
		for _, cmd := range cmdNode.children {
			if 0 == strings.Compare(cmd.text, word) {
				foundWord = true
				cmdIdx++
				cmdNode = cmd
			}
		}
		if !foundWord {
			break
		}
	}

	return cmdNode, strings.Join(commandSplit[cmdIdx:], " ")
}

/*
 * CLI COMMANDS
 */
func (cli *CommandLine) initCommands() {
	logger.Debug("Initializing commands.")
	cli.CmdTree = &Command{"", nil, usageRoot}

	cli.CmdTree.addSubCmd("exit")
	cli.CmdTree.addSubCmd("help")

	cli.CmdTree.addSubCmd("list")

	cmdGrid := cli.CmdTree.addSubCmd("grid")
	cmdGrid.addSubCmd("list")
	cmdGrid.addSubCmd("info")
	cmdGrid.addSubCmd("goto")
	cmdGrid.addSubCmd("bookmark")

	cmdUnit := cli.CmdTree.addSubCmd("unit")
	cmdUnit.addSubCmd("list")
	cmdUnit.addSubCmd("select")
	cmdUnit.addSubCmd("deselect")
	cmdUnit.addSubCmd("alias")
	cmdUnit.addSubCmd("move")
	cmdUnit.addSubCmd("stop")
	cmdUnit.addSubCmd("scan")
	cmdUnitBuild := cmdUnit.addSubCmd("build")
	cmdUnitBuild.addSubCmd("building")
	cmdUnitSpawn := cmdUnit.addSubCmd("spawn")
	cmdUnitSpawn.function = spawn

	cmdGroup := cli.CmdTree.addSubCmd("group")
	cmdGroup.addSubCmd("list")
	cmdGroup.addSubCmd("select")
	cmdGroup.addSubCmd("deselect")
	cmdGroup.addSubCmd("set")
	cmdGroup.addSubCmd("unset")
	cmdGroup.addSubCmd("alias")
	cmdGroup.addSubCmd("move")
	cmdGroup.addSubCmd("stop")
}

/*
 * COMMAND COMMAND COMMAND
 */

type Command struct {
	text string
	children []*Command
	function Function
}

type Function func(string) (string)

func (cmd *Command) addSubCmd(t string) *Command {
	subCmd := &Command{t, nil, nil}
	cmd.children = append(cmd.children, subCmd)
	return subCmd
}

func (cmd *Command) Run(s string) string {
	var msg string = ""
	logger.Debug("Executing command '%s':'%s'", cmd.text, s)
	if cmd.function != nil {
		msg = cmd.function(s)
	}
	return msg
}

func (cmd *Command) listCommands(s string) string {
	var msg []string
	for _, cmd := range cmd.children {
		msg = append(msg, cmd.text)
	}
	return "Commands: " + strings.Join(msg, ", ")
}

/*
 * LISTENER LISTENER LISTENER
 */
func eventListener() {
	logger.Debug("Initializing event listener.")
loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyCtrlX {
				break loop
			}
			Control.inputMode.EventHandler(event)
		case termbox.EventError:
			Control.Quit()
		}
	}
	Control.Quit()
}
