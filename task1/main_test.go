package main

import "testing"

func TestCalculateAverage(t *testing.T) {
	dict := map[string]int{
		"Math":    85,
		"Science": 90,
	}
	expected := 87.5
	result := calculateAverage(dict)
	if result != expected {
		t.Errorf("calculateAverage(%v) = %v; expected %v", dict, result, expected)
	}
}