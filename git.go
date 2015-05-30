package main

import (
	"os/exec"
	"strings"
)

type Git struct {
	branch string
}

func NewGit() *Git {
	return &Git{}
}

func (g *Git) BackToThePast() error {
	current, err := exec.Command("git", "name-rev", "--name-only", "HEAD").Output()
	if err != nil {
		return err
	}
	g.branch = strings.Trim(string(current), "\n")

	return g.co("HEAD~")
}

func (g *Git) BackToTheFuture() error {
	return g.co(g.branch)
}

func (_ *Git) co(to string) error {
	return exec.Command("git", "checkout", to).Run()
}
