package main

import (
	"bytes"
	"os"
)

import "testing"

func TestExecBench(t *testing.T) {
	b := make([]byte, 0)
	w := bytes.NewBuffer(b)
	err := ExecBench(w)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTempfile(t *testing.T) {
	f, fn, err := Tempfile()
	if err != nil {
		t.Fatal(err)
	}
	if !fileExists(f.Name()) {
		t.Fatalf("File %s doesn't exist.", f.Name())
	}

	fn()
	if fileExists(f.Name()) {
		t.Fatalf("File %s should be removed. But exist.", f.Name())
	}
}

func fileExists(fname string) bool {
	_, err := os.Stat(fname)
	return err == nil
}
