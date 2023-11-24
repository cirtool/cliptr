package main

import (
	"testing"
)

func TestTrim(t *testing.T) {
	want := "This is an awesome test"
	trimmed := Trim("    This is an awesome test     ")
	if trimmed != want {
		t.Fatalf(`Invalid, %q != %q`, trimmed, want)
	}
}
