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

	return exec.Command("git", "checkout", "HEAD~").Run()
}

func (g *Git) BackToTheFuture() error {
	return exec.Command("git", "checkout", g.branch).Run()
}
