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
func initControl() Controller {
	var control Controller

	control.cli  = &CommandLine{nil, &strings.Builder{}}
	control.cli.initCommands()

	control.gameMap = &Map{}
	control.gameMap.initMap()

	control.setInputMode(control.gameMap)

	return control
}

type Controller struct {
	inputMode Input
	cli *CommandLine
	gameMap *Map
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
		}
		Control.setInputMode(Control.gameMap)
		return

	case termbox.KeyBackspace, termbox.KeyBackspace2:
		tmpInput := cli.Buffer.String()
		if len(tmpInput) > 0 {
			cli.Buffer.Reset()
			cli.Buffer.WriteString(tmpInput[0:len(tmpInput)-1])
		}
		return

	case termbox.KeySpace:
		cli.Buffer.WriteRune(0x0020) // Space
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

	cmdList := cli.CmdTree.addSubCmd("list")
	cmdList.function = cli.CmdTree.listCommands

	cmdGrid := cli.CmdTree.addSubCmd("grid")
	cmdGridList := cmdGrid.addSubCmd("list")
	cmdGridList.function = cmdGrid.listCommands
	cmdGridSelect := cmdGrid.addSubCmd("select")
	cmdGridSelect.function = gridSelect
	cmdGridInfo := cmdGrid.addSubCmd("info")
	cmdGridInfo.function = gridInfo
	cmdGridGoto := cmdGrid.addSubCmd("goto")
	cmdGridGoto.function = gridGoto
	cmdGrid.addSubCmd("bookmark")

	cmdUnit := cli.CmdTree.addSubCmd("unit")
	cmdUnitList := cmdUnit.addSubCmd("list")
	cmdUnitList.function = cmdUnit.listCommands
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
	cmdGroupList := cmdGroup.addSubCmd("list")
	cmdGroupList.function = cmdGroup.listCommands
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
