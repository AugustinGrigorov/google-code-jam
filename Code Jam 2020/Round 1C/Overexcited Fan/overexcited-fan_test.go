package main

import (
	"bytes"
	"os"
	"testing"
)

var want = `Case #1: 4
Case #2: IMPOSSIBLE
Case #3: IMPOSSIBLE
Case #4: 1
Case #5: 5
Case #6: 4
Case #7: 4
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
