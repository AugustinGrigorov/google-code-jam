package main

import (
	"bytes"
	"os"
	"testing"
)

var want = `Case #1: CJC
Case #2: IMPOSSIBLE
Case #3: JCCJJ
Case #4: CC
`

func TestParentingPartneringReturns(t *testing.T) {
	in, _ = os.Open("./input.txt")
	out = new(bytes.Buffer)
	main()
	got := out.(*bytes.Buffer).String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
