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

// co is git checkout
func (_ *Git) co(to string) error {
	return exec.Command("git", "checkout", to).Run()
}

func (_ *Git) hasChange() bool {
	out, _ := exec.Command("git", "status", "--short").Output()
	return len(out) != 0
}

func (_ *Git) stash() error        { return exec.Command("git", "stash", "save", "--include-untracked").Run() }
func (_ *Git) recoverStash() error { return exec.Command("git", "stash", "pop").Run() }
