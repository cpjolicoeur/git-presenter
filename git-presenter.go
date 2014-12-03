// +build go1.2

package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/cpjolicoeur/git-presenter/cmd"
)

const APP_VER = "0.1.1"

func main() {
	app := cli.NewApp()
	app.Name = "git-presenter"
	app.Usage = "Git presentation tool"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdInit,
		cmd.CmdStart,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
