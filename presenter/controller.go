// +build go1.2

package presenter

import (
	"fmt"

	"github.com/cpjolicoeur/git-presenter/models"
)

type Controller struct {
	repoPath string
	commits  []models.ConfigMeta
}

/*
 * Setup the presentation based on configuration file
 */
func (c *Controller) Load(config models.PresentationConfig, verbose bool) {
	if verbose {
		fmt.Println("Controller Load", config)
	}
	c.repoPath = config.Repo
	c.commits = config.Commits
}

/*
 * Setup the presentation based on direct repository parsing
 */
func (c *Controller) Initialize(repositoryPath string, verbose bool) {
	if verbose {
		fmt.Println("Controller Initialize", repositoryPath)
	}
	c.repoPath = repositoryPath
}

/*
 * Run the main presentation loop
 */
func (c *Controller) Start(verbose bool) {
	if verbose {
		fmt.Println("Starting presentation...")
	}
}
