package cmd

import (
	"logger"
	"strings"
)

var root *Command = &Command{"", nil, nil}
var exit *Command = root.addSubCmd("exit")

func (c *Command) addSubCmd(t string) *Command {
	cmd := &Command{t, nil, nil}
	c.children = append(c.children, cmd)
	return cmd
}

type Command struct {
	text string
	children []*Command
	function Function
}

type Function func(string) (string)

func Find(t string) (*Command, string) {
	var cmdNode *Command = root
	commandSplit := strings.Split(t, " ")

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

func Init() {
	logger.Debug("Initializing commands.")
	// TODO remove these
	cmdTest := root.addSubCmd("test")
	cmdThis := cmdTest.addSubCmd("this")
	cmdThis.addSubCmd("command")

	// TODO add help func
	root.addSubCmd("help")
	cmdList := root.addSubCmd("list")
	cmdList.function = listCommands

	// TODO flesh out commands for game
	cmdWorld := root.addSubCmd("world")
	cmdWorld.addSubCmd("goto")
	cmdWorld.addSubCmd("bookmark")

	cmdUnit := root.addSubCmd("unit")
	cmdUnit.addSubCmd("select")
	cmdUnit.addSubCmd("deselect")
	cmdUnit.addSubCmd("alias")
	cmdUnit.addSubCmd("move")
	cmdUnit.addSubCmd("stop")
	cmdUnit.addSubCmd("scan")
	cmdUnitBuild := cmdUnit.addSubCmd("build")
	cmdUnitBuild.addSubCmd("building")

	cmdGroup := root.addSubCmd("group")
	cmdGroup.addSubCmd("select")
	cmdGroup.addSubCmd("deselect")
	cmdGroup.addSubCmd("set")
	cmdGroup.addSubCmd("unset")
	cmdGroup.addSubCmd("alias")
	cmdGroup.addSubCmd("move")
	cmdGroup.addSubCmd("stop")
}

func Run(cmd *Command, s string) string {
	var msg string = ""
	logger.Debug("Executing command '%s':'%s'", cmd.text, s)
	if cmd.function != nil {
		msg = cmd.function(s)
	}
	return msg
}

func listCommands(s string) string {
	var msg []string
	for _, cmd := range root.children {
		msg = append(msg, cmd.text)
	}
	return "Commands: " + strings.Join(msg, ", ")
}
