package cmd

import (
	"bsh-tfe/mgrs"
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
	cmdGrid := root.addSubCmd("grid")
	cmdGridSelect := cmdGrid.addSubCmd("select")
	cmdGridSelect.function = gridSelect
	cmdGridInfo := cmdGrid.addSubCmd("info")
	cmdGridInfo.function = gridInfo
	cmdGrid.addSubCmd("goto")
	cmdGrid.addSubCmd("bookmark")
	cmdGridSet := cmdGrid.addSubCmd("set")
	cmdGridSetBiome := cmdGridSet.addSubCmd("biome")
	cmdGridSetBiome.function = gridSetBiome

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

func gridSelect(s string) string {
	grid, err := mgrs.StringToGridDesignation(s)
	if err != nil {
		return "Invalid grid designation."
	}

	world.SelectedGrid = grid
	return "Selected grid " + grid.ToString() + "."
}

func gridSetBiome(s string) string {
	msg := strings.Split(s, " ")
	t, err := strconv.Atoi(msg[0])
	if err != nil {
		return "Usage: map set biome <biome>"
	}

	if world.SelectedGrid != (mgrs.GridDesignation{}) {
		grid := world.Terra.GetGrid(world.SelectedGrid)
		grid.Biome = t
		return "Set biome successfully."
	} else {
		return "No grid selected. Use: map select <grid>"
	}
}

func gridInfo(s string) string {
	b := world.Terra.GetGrid(world.SelectedGrid).Biome
	return "Grid " + world.SelectedGrid.ToString() + ": Biome " + strconv.Itoa(b)
}
