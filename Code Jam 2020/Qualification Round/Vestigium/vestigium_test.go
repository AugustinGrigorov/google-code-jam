package main

import (
	"bytes"
	"os"
	"testing"
)

var want = `Case #1: 4 0 0
Case #2: 9 4 4
Case #3: 8 0 2
`

func TestVestigium(t *testing.T) {
	in, _ = os.Open("./input.txt")
	out = new(bytes.Buffer)
	main()
	got := out.(*bytes.Buffer).String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
