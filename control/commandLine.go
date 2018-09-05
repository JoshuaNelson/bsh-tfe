package control

import (
	"github.com/nsf/termbox-go"
	"logger"
	"strings"
)

var CommandLine Terminal

type Terminal struct {
	CmdTree *Command
	Buffer *strings.Builder
}

func (t Terminal) EventHandler(event termbox.Event) {
	switch event.Key {
	case termbox.KeyEnter:
		defer t.Buffer.Reset()
		cmd, msg := t.Find(t.Buffer.String())
		logger.Debug("Executing buffer: %s", t.Buffer.String())
		if cmd != nil {
			logger.Debug("Running with msg %s", msg)
			cmd.Run(msg)
		}
		return

	case termbox.KeyBackspace, termbox.KeyBackspace2:
		tmpInput := t.Buffer.String()
		if len(tmpInput) > 0 {
			t.Buffer.Reset()
			t.Buffer.WriteString(tmpInput[0:len(tmpInput)-1])
		}
		return

	case termbox.KeySpace:
		t.Buffer.WriteRune(0x0020) // Space
		return
	}

	t.Buffer.WriteRune(event.Ch)
	return
}

func InitMainCommandLine() Terminal {
	logger.Debug("Initializing MCL commands.")
	term := Terminal{}
	term.Buffer = &strings.Builder{}
	term.CmdTree = &Command{"", nil, usageRoot}

	term.CmdTree.addSubCmd("exit")
	term.CmdTree.addSubCmd("help")

	cmdList := term.CmdTree.addSubCmd("list")
	cmdList.function = term.CmdTree.listCommands

	cmdGrid := term.CmdTree.addSubCmd("grid")
	cmdGridList := cmdGrid.addSubCmd("list")
	cmdGridList.function = cmdGrid.listCommands
	cmdGridSelect := cmdGrid.addSubCmd("select")
	cmdGridSelect.function = gridSelect
	cmdGridInfo := cmdGrid.addSubCmd("info")
	cmdGridInfo.function = gridInfo
	cmdGrid.addSubCmd("goto")
	cmdGrid.addSubCmd("bookmark")
	cmdGridSet := cmdGrid.addSubCmd("set")
	cmdGridSetBiome := cmdGridSet.addSubCmd("biome")
	cmdGridSetBiome.function = gridSetBiome

	cmdUnit := term.CmdTree.addSubCmd("unit")
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

	cmdGroup := term.CmdTree.addSubCmd("group")
	cmdGroupList := cmdGroup.addSubCmd("list")
	cmdGroupList.function = cmdGroup.listCommands
	cmdGroup.addSubCmd("select")
	cmdGroup.addSubCmd("deselect")
	cmdGroup.addSubCmd("set")
	cmdGroup.addSubCmd("unset")
	cmdGroup.addSubCmd("alias")
	cmdGroup.addSubCmd("move")
	cmdGroup.addSubCmd("stop")

	return term
}
