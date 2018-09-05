package control

import (
	"logger"
	"strings"
)

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

func (t *Terminal) Find(s string) (*Command, string) {
	var cmdNode *Command = t.CmdTree
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

func (c *Command) Run(s string) string {
	var msg string = ""
	logger.Debug("Executing command '%s':'%s'", c.text, s)
	if c.function != nil {
		msg = c.function(s)
	}
	return msg
}

func (c *Command) listCommands(s string) string {
	var msg []string
	for _, cmd := range c.children {
		msg = append(msg, cmd.text)
	}
	return "Commands: " + strings.Join(msg, ", ")
}
