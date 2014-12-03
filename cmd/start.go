// +build go1.2

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var CmdStart = cli.Command{
	Name:        "start",
	Usage:       "Start a git presentation",
	Description: `Start begins the git presentation process.  If a .git-presentation file is found it will be used, otherwise all commits will be used in reverse-chronological order.`,
	Action:      runStart,
	Flags: []cli.Flag{
		cli.BoolFlag{"verbose, v", "show process details", ""},
		cli.StringFlag{"repo, r", ".", "present specified repository", ""},
	},
}

func runStart(ctx *cli.Context) {
	configFile, exists := configFileExists()
	if exists {
		defer configFile.Close()
		if ctx.Bool("verbose") {
			fmt.Printf("Found %s file\n", PRESENTATION_FILE)
		}

		jsonConfig, err := ioutil.ReadAll(configFile)
		if err != nil {
			log.Fatal(err)
		}

		var presentationConfig PresentationConfig
		err = json.Unmarshal(jsonConfig, &presentationConfig)
		if err != nil { // bad JSON config for some reason,
			fmt.Printf("Your %s appears to be invalid.  Please either run `git-presenter init` again to rebuild it, or remove the file completely and run `git-presenter start` again.\n\tError: %q\n", PRESENTATION_FILE, err)
			return
		}

		if ctx.Bool("verbose") {
			fmt.Println("Presenting repository from:", presentationConfig.Repo)
		}
		// presentFromConfig()
	} else {
		// present()
	}
	return
}

func configFileExists() (*os.File, bool) {
	file, err := os.Open(PRESENTATION_FILE)
	if err != nil {
		return nil, false
	}

	return file, true
}
