package main

import (
	"bytes"
	"os"
	"testing"
)

var want = `Case #1: 0000
Case #2: (1)0(1)
Case #3: (111)000
Case #4: (1)
`

func TestNestingDepth(t *testing.T) {
	in, _ = os.Open("./input.txt")
	out = new(bytes.Buffer)
	main()
	got := out.(*bytes.Buffer).String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
