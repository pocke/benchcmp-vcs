package main

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const (
	GO = "go"
)

func main() {
	before, beforeClose, err := Tempfile()
	if err != nil {
		panic(err)
	}
	defer beforeClose()

	after, afterClose, err := Tempfile()
	if err != nil {
		panic(err)
	}
	defer afterClose()

	err = ExecBench(after)
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

	err = ExecBench(before)
	if err != nil {
		panic(err)
	}

	err = ExecCmp(os.Stdout, before.Name(), after.Name())
	if err != nil {
		panic(err)
	}
}

func Tempfile() (*os.File, func(), error) {
	file, err := ioutil.TempFile("", "benchcmp-git")
	if err != nil {
		return nil, nil, err
	}
	fn := func() {
		file.Close()
		os.Remove(file.Name())
	}

	return file, fn, nil
}

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
