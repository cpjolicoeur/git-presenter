// +build go1.2

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/libgit2/git2go"
)

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

const PRESENTATION_FILE = ".git-presentation"

type PresentationConfig struct {
	Repo    string       `json:"repository"`
	Commits []ConfigMeta `json:"commits"`
}

type ConfigMeta struct {
	Sha     string `json:"sha"`
	Message string `json:"message"`
}

func runInit(ctx *cli.Context) {
	if ctx.Bool("verbose") {
		fmt.Printf("Initializing %s file\n", PRESENTATION_FILE)
	}

	repoPath := ctx.String("repo")

	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		fmt.Println("Problem opening repository:", err)
		return
	}
	defer repo.Free()

	revWalker, err := repo.Walk()
	if err != nil {
		fmt.Println("Problem walking repository commit history:", err)
		return
	}
	defer revWalker.Free()

	err = revWalker.PushHead()
	if err != nil {
		fmt.Println("Problem pushing HEAD onto walker:", err)
		return
	}

	revWalker.Sorting(git.SortReverse)

	err = createPresentationFile(*revWalker, repo.Path(), ctx.Bool("verbose"))
	if err != nil {
		fmt.Printf("Error creating %s file: %q\n", PRESENTATION_FILE, err)
		return
	}
	return
}

func createPresentationFile(walker git.RevWalk, repoPath string, verbose bool) (err error) {
	f, err := os.Create(PRESENTATION_FILE)
	if err != nil {
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()

	config := new(PresentationConfig)
	config.Repo = repoPath

	iterator := func(commit *git.Commit) bool {
		if verbose {
			fmt.Printf("Adding Commit: %s => %s", commit.Id(), commit.Message())
		}
		config.Commits = append(config.Commits, ConfigMeta{commit.Id().String(), commit.Message()})
		return true
	}

	err = walker.Iterate(iterator)
	if err != nil {
		return
	}

	configJSON, err := json.Marshal(config)
	if err != nil {
		return
	}

	_, err = f.Write(configJSON)
	return
}
