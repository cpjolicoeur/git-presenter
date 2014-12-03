// +build go1.2

package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var CmdStart = cli.Command{
	Name:        "start",
	Usage:       "Start a git presentation",
	Description: `Start begins the git presentation process.  If a .git-presentation file is found it will be used, otherwise all commits will be used in reverse-chronological order.`,
	Action:      runStart,
	Flags: []cli.Flag{
		cli.BoolFlag{"verbose, v", "show process details", ""},
	},
}

func runStart(ctx *cli.Context) {
	fmt.Println("TODO: start presentation here...")
	return
}
