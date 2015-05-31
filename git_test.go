package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestNewGit(t *testing.T) {
	func() {
		t.Log("when repository has not changes.")
		fn, err := takeGitDir()
		if err != nil {
			t.Fatal(err)
		}
		defer fn()

		git := NewGit()
		if git.method != methodCheckout {
			t.Error("method expected: checkout, but not")
		}
	}()

	func() {
		t.Log("when repository has some changes.")
		fn, err := takeGitDir()
		if err != nil {
			t.Fatal(err)
		}
		defer fn()

		exec.Command("touch", "hoge").Run()

		git := NewGit()
		if git.method != methodStash {
			t.Error("method expected: stash, but not")
		}
	}()
}

func TestStash(t *testing.T) {
	fn, err := takeGitDir()
	if err != nil {
		t.Fatal(err)
	}
	defer fn()

	exec.Command("touch", "hoge").Run()

	git := NewGit()
	err = git.stash()
	if err != nil {
		t.Fatal(err)
	}
	if git.hasChange() {
		t.Error("Expected: no changes, but got some changes")
	}

	err = git.recoverStash()
	if err != nil {
		t.Fatal(err)
	}
	if !git.hasChange() {
		t.Error("Expected: some changes, but got no changes")
	}
}

// test helper
func takeGitDir() (func(), error) {
	dir, err := ioutil.TempDir("", "benchcmp-vcs-test")
	if err != nil {
		return nil, err
	}
	prevDir, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}

	os.Chdir(dir)
	exec.Command("git", "init").Run()
	exec.Command("git", "commit", "-m", "commit 1", "--allow-empty").Run()
	exec.Command("git", "commit", "-m", "commit 2", "--allow-empty").Run()
	exec.Command("git", "commit", "-m", "commit 3", "--allow-empty").Run()

	fn := func() {
		os.Chdir(prevDir)
		os.RemoveAll(dir)
	}
	return fn, nil
}
