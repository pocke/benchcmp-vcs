package main

import (
	"io"
	"io/ioutil"
	"strings"
)
import (
	"os"
	"os/exec"
)

func main() {
	before, err := ioutil.TempFile("", "benchcmp-git")
	if err != nil {
		panic(err)
	}
	defer before.Close()
	defer os.Remove(before.Name())

	err = ExecBench(before)
	if err != nil {
		panic(err)
	}

	current, err := exec.Command("git", "name-rev", "--name-only", "HEAD").Output()
	if err != nil {
		panic(err)
	}

	exec.Command("git", "checkout", "HEAD~").Run()
	defer func(branch []byte) {
		b := strings.Trim(string(branch), "\n")
		c := exec.Command("git", "checkout", b)
		c.Run()
	}(current)

	after, err := ioutil.TempFile("", "benchcmp-git")
	if err != nil {
		panic(err)
	}
	defer after.Close()
	defer os.Remove(after.Name())

	err = ExecBench(after)
	if err != nil {
		panic(err)
	}

	err = ExecCmp(os.Stdout, before.Name(), after.Name())
	if err != nil {
		panic(err)
	}
}

const (
	GO = "go"
)

func ExecBench(w io.Writer, opts ...string) error {
	o := append([]string{"test", "-run=NONE", "-bench=.", "-benchmem"}, opts...)
	out, err := exec.Command(GO, o...).Output()
	if err != nil {
		return err
	}

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func ExecCmp(w io.Writer, before, after string, opts ...string) error {
	CMD := "benchcmp"
	o := append([]string{before, after}, opts...)
	out, err := exec.Command(CMD, o...).Output()
	if err != nil {
		return err
	}

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}
