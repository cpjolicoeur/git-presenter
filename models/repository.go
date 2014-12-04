// +build go1.2

package models

import (
	"fmt"

	"github.com/libgit2/git2go"
)

type Repository struct {
	Path       string
	Repository *git.Repository
	Commits    []*git.Commit
}

func (r *Repository) Open() (err error) {
	r.Repository, err = git.OpenRepository(r.Path)
	if err != nil {
		return
	}
	r.Path = r.Repository.Path()

	return
}

func (r *Repository) Cleanup() {
	r.Repository.Free()
}

func (r *Repository) Process(verbose bool) (err error) {
	revWalk, err := r.Repository.Walk()
	if err != nil {
		return
	}
	defer revWalk.Free()

	err = revWalk.PushHead()
	if err != nil {
		return
	}

	revWalk.Sorting(git.SortReverse)

	err = r.parseCommits(revWalk, verbose)
	return
}

func (r *Repository) ProcessFromConfig(configCommits []ConfigMeta, verbose bool) (err error) {
	for _, c := range configCommits {
		oid, err := git.NewOid(c.Sha)
		if err != nil {
			fmt.Printf("Error loading commit %s\n\t%q", c.Sha, err)
			return err
		}

		commit, err := r.Repository.LookupCommit(oid)
		if err != nil {
			fmt.Printf("Error loading commit %s\n\t%q", c.Sha, err)
			return err
		}
		if verbose {
			fmt.Println("Adding commit:", commit.Id().String())
		}
		r.Commits = append(r.Commits, commit)
	}
	return
}

func (r *Repository) parseCommits(revWalk *git.RevWalk, verbose bool) (err error) {
	iterator := func(commit *git.Commit) bool {
		if verbose {
			fmt.Printf("Adding Commit: %s => %s", commit.Id(), commit.Message())
		}
		r.Commits = append(r.Commits, commit)
		return true
	}

	err = revWalk.Iterate(iterator)
	if err != nil {
		return
	}

	return
}
