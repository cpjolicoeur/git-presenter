// +build go1.2

package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var CmdInit = cli.Command{
	Name:  "init",
	Usage: "Initialize a git-presenter config file",
	Description: `Init creates a .git-presenter configuration file listing the commits that will be used during the git presentation.
You can edit use this file to customize which commits will be included in the presentation and also the order the commits are presented.`,
	Action: runInit,
	Flags: []cli.Flag{
		cli.BoolFlag{"verbose, v", "show process details", ""},
	},
}

func runInit(ctx *cli.Context) {
	if ctx.Bool("verbose") {
		fmt.Println("Running in verbose mode")
	}
	fmt.Println("TODO: run init process here...")
	return
}
