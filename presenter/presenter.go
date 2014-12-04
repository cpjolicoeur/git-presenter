// +build go1.2

package presenter

import (
	"fmt"

	"github.com/cpjolicoeur/git-presenter/models"
)

type Presenter struct {
	Verbose    bool
	Repository *models.Repository
	Controller *Controller
}

/*
 * Setup the presentation based on configuration file
 */
func (p *Presenter) Load(config *models.PresentationConfig) {
	if p.Verbose {
		fmt.Println("Presenting repository from:", config.Repo)
	}
	p.Repository = &models.Repository{Path: config.Repo}
	err := p.Repository.Open()
	if err != nil {
		fmt.Println("Problem opening repository at:", config.Repo)
		return
	}
	defer p.Repository.Cleanup()

	// Load the Commits manually based on the config
	err = p.Repository.ProcessFromConfig(config.Commits, p.Verbose)
	if err != nil {
		fmt.Printf("Problem loading commits:\n\t%q", err)
		return
	}
}

/*
 * Setup the presentation based on direct repository parsing
 */
func (p *Presenter) Initialize(repositoryPath string) {
	if p.Verbose {
		fmt.Println("Initializing Presentation from repository:", repositoryPath)
	}

	p.Repository = &models.Repository{Path: repositoryPath}
	err := p.Repository.Open()
	if err != nil {
		fmt.Println("Problem opening repository at:", repositoryPath)
		return
	}
	defer p.Repository.Cleanup()

	err = p.Repository.Process(p.Verbose)
	if err != nil {
		fmt.Println("Problem parsing repository:", err)
		return
	}
}

/*
 * Run the main presentation loop
 */
func (p *Presenter) Start() (err error) {
	fmt.Println("Starting presentation...")
	p.Controller.Commits = p.Repository.Commits
	err = p.Controller.StartPresentation()
	return
}
