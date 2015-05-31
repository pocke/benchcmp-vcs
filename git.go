package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type gitMethod int

const (
	methodCheckout gitMethod = iota + 1
	methodStash
)

type Git struct {
	// Branch before benchcmp-vcs command run.
	// For BackToTheFuture method when use checkout.
	now string

	// revision hash for show revision.
	old, new string

	method gitMethod
}

func NewGit() *Git {
	git := &Git{}
	if git.hasChange() {
		git.method = methodStash
	} else {
		git.method = methodCheckout
	}
	return git
}

func (g *Git) BackToThePast() error {
	g.new = g.getRevision()

	var err error
	switch g.method {
	case methodCheckout:
		err = g.checkoutPast()
	case methodStash:
		err = g.stash()
	default:
		err = fmt.Errorf("unexpected method")
	}

	g.old = g.getRevision()
	return err
}

func (g *Git) BackToTheFuture() error {
	switch g.method {
	case methodCheckout:
		return g.checkoutFuture()
	case methodStash:
		return g.recoverStash()
	default:
		return fmt.Errorf("unexpected method")
	}
}

func (g *Git) NewRevision() string {
	if g.method != methodCheckout {
		return ""
	}
	return g.new
}

func (g *Git) OldRevision() string {
	return g.old
}

// co is git checkout
func (_ *Git) co(to string) error { return exec.Command("git", "checkout", to).Run() }

func (_ *Git) hasChange() bool {
	out, _ := exec.Command("git", "status", "--short").Output()
	return len(out) != 0
}

func (_ *Git) getRevision() string {
	out, _ := exec.Command("git", "rev-parse", "HEAD").Output()
	return strings.Trim(string(out), "\n")
}

func (g *Git) checkoutPast() error {
	now, err := exec.Command("git", "name-rev", "--name-only", "HEAD").Output()
	if err != nil {
		return err
	}
	g.now = strings.Trim(string(now), "\n")

	return g.co("HEAD~")
}
func (g *Git) checkoutFuture() error { return g.co(g.now) }

func (_ *Git) stash() error        { return exec.Command("git", "stash", "save", "--include-untracked").Run() }
func (_ *Git) recoverStash() error { return exec.Command("git", "stash", "pop").Run() }
