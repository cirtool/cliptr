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

func TestCapitalizeUppercase(t *testing.T) {
	want := "An Awesome Test"
	capitalized := CapitalizeUppercase("AN AWESOME TEST")
	if capitalized != want {
		t.Fatalf(`Invalid, %q != %q`, capitalized, want)
	}
}

func TestCapitalizeUppercaseInEmptyString(t *testing.T) {
	want := "   "
	capitalized := CapitalizeUppercase("   ")
	if capitalized != want {
		t.Fatalf(`Invalid, %q != %q`, capitalized, want)
	}
}
