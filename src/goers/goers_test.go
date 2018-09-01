package goers

import "testing"

func TestSimple(t *testing.T) {
    result := 42
    if result != 42 {
        t.Error("Expected 42, got", result)
    }
}

func TestMap(t *testing.T) {
    var result map[string]int
    result = make(map[string]int)
    result["test"] = 42
    if result["test"] != 42 {
        t.Error("Expected 42, got", result["test"])
    }
}
