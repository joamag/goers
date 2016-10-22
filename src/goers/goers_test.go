package goers

import "testing"

func TestSimple(t *testing.T) {
    result := 2
    if result != 2 {
        t.Error("Expected 2, got", result)
    }
}
