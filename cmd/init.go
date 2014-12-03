// +build go1.2

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/cpjolicoeur/git-presenter/models"
)

const PRESENTATION_FILE = ".git-presentation"

var CmdInit = cli.Command{
	Name:  "init",
	Usage: "Initialize a git-presenter config file",
	Description: `Init creates a .git-presenter configuration file listing the commits that will be used during the git presentation.
You can edit use this file to customize which commits will be included in the presentation and also the order the commits are presented.`,
	Action: runInit,
	Flags: []cli.Flag{
		cli.BoolFlag{"verbose, v", "show process details", ""},
		cli.StringFlag{"repo, r", ".", "present specified repository", ""},
	},
}

func runInit(ctx *cli.Context) {
	verbose := ctx.Bool("verbose")
	if verbose {
		fmt.Printf("Initializing %s file\n", PRESENTATION_FILE)
	}
	repoPath := ctx.String("repo")

	repository := &models.Repository{Path: repoPath}
	err := repository.Open()
	if err != nil {
		fmt.Println("Problem opening repository:", err)
		return
	}
	defer repository.Cleanup()

	err = repository.Process(verbose)
	if err != nil {
		fmt.Printf("Error processing repository: %s", err)
		return
	}

	err = createPresentationFile(repository, verbose)
	if err != nil {
		fmt.Printf("Error creating %s file: %q\n", PRESENTATION_FILE, err)
		return
	}
	return
}

func createPresentationFile(r *models.Repository, verbose bool) (err error) {
	f, err := os.Create(PRESENTATION_FILE)
	if err != nil {
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()

	config := new(models.PresentationConfig)
	config.Repo = r.Path

	for _, c := range r.Commits {
		config.Commits = append(config.Commits, models.ConfigMeta{c.Id().String(), c.Message()})
	}

	configJSON, err := json.Marshal(config)
	if err != nil {
		return
	}

	_, err = f.Write(configJSON)
	return
}
