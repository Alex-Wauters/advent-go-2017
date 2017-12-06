package main

import "testing"

func TestPart1(t *testing.T) {
	input := banks{0,2,7,0}
	if result := part1(input); result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestPart2(t *testing.T) {
	input := banks{0,2,7,0}
	if result := part2(input); result != 4 {
		t.Errorf("Expected 4, got %v", result)
	}
}


