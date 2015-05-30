package main

import "bytes"

import "testing"

func TestExecBench(t *testing.T) {
	b := make([]byte, 0)
	w := bytes.NewBuffer(b)
	err := ExecBench(w)
	if err != nil {
		t.Fatal(err)
	}
}
