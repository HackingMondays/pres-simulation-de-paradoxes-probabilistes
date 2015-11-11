package main

import (
	"testing"
)

func TestGroupSizeSolver(t *testing.T) {

	expected := 23
	actual := findSize()

	if actual != expected {
		t.Errorf("Expected: %d, but was: %d", expected, actual)
	}
}
