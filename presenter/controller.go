// +build go1.2

package presenter

import (
	"fmt"

	"github.com/libgit2/git2go"
)

type Controller struct {
	Commits []*git.Commit

	currentCommit *git.Commit
	currentIndex  int
	lastCommand   string
}

func (c *Controller) StartPresentation() (err error) {
	c.currentIndex = 0
	err = c.loadCommit(c.currentIndex)
	if err != nil {
		return
	}

	// c.WaitForCommand()
	return
}

func (c *Controller) NextSlide() (err error) {
	return
}

func (c *Controller) PreviousSlide() (err error) {
	return
}

func (c *Controller) loadCommit(index int) (err error) {
	c.currentCommit = c.Commits[index]
	// tree, err := c.currentCommit.Tree()
	// tree, err := c.Commits[index].Tree()
	// if err != nil {
	// 	fmt.Println("commit tree error:", err)
	// }
	// tree.Free()
	// fmt.Println("commit tree", tree)
	// fmt.Println("commit tree", tree.Owner().Path)

	fmt.Println("commit to load", c.Commits[index].Message())
	return
}
