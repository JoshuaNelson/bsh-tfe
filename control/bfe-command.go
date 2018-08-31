package cmd

import (
	"bsh-tfe/world"
	"logger"
	"strconv"
	"strings"
)

var root *Command = &Command{"", nil, usageRoot}
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
	// TODO add help func
	root.addSubCmd("help")

	cmdList := root.addSubCmd("list")
	cmdList.function = listCommands

	// TODO flesh out commands for game
	cmdMap := root.addSubCmd("map")
	cmdMapSelect := cmdMap.addSubCmd("select")
	cmdMapSelect.function = mapSelect
	cmdMapInfo := cmdMap.addSubCmd("info")
	cmdMapInfo.function = mapInfo
	cmdMap.addSubCmd("goto")
	cmdMap.addSubCmd("bookmark")
	cmdMapSet := cmdMap.addSubCmd("set")
	cmdMapSetBiome := cmdMapSet.addSubCmd("biome")
	cmdMapSetBiome.function = mapSetBiome

	cmdUnit := root.addSubCmd("unit")
	cmdUnit.addSubCmd("select")
	cmdUnit.addSubCmd("deselect")
	cmdUnit.addSubCmd("list")
	cmdUnit.addSubCmd("alias")
	cmdUnit.addSubCmd("move")
	cmdUnit.addSubCmd("stop")
	cmdUnit.addSubCmd("scan")
	cmdUnitBuild := cmdUnit.addSubCmd("build")
	cmdUnitBuild.addSubCmd("building")

	cmdGroup := root.addSubCmd("group")
	cmdGroup.addSubCmd("select")
	cmdGroup.addSubCmd("deselect")
	cmdGroup.addSubCmd("list")
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

func usageRoot(s string) string {
	msg := strings.Split(s, " ")
	return "Command '" + msg[0] + "' not recognized."
}

func listCommands(s string) string {
	var msg []string
	for _, cmd := range root.children {
		msg = append(msg, cmd.text)
	}
	return "Commands: " + strings.Join(msg, ", ")
}

func mapSelect(s string) string {
	msg := strings.Split(s, " ")
	x, err := strconv.Atoi(msg[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(msg[1])
	if err != nil {
		panic(err)
	}

	world.Selected = world.GameMap.Grid(x,y)
	return "Selected map at " + msg[0] + "x" + msg[1]
}

func mapSetBiome(s string) string {
	msg := strings.Split(s, " ")
	t, err := strconv.Atoi(msg[0])
	if err != nil {
		return "Usage: map set biome <biome>"
	}

	if world.Selected != nil {
		world.Selected.SetBiome(t)
		return "Set biome successfully."
	} else {
		return "No grid selected. Use: map select <grid>"
	}
}

func mapInfo(s string) string {
	msg := strings.Split(s, " ")
	x, err := strconv.Atoi(msg[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(msg[1])
	if err != nil {
		panic(err)
	}

	b := world.GameMap.Grid(x,y).Biome

	return "Selected map at " + msg[0] + "x" + msg[1] + ": Biome " + strconv.Itoa(b)
}
