package goers-run

import "testing"

func TestSimple(t *testing.T) {
	result := 42
	if result != 42 {
		t.Error("Expected 42, got", result)
	}
}
