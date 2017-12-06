package main

import "testing"

func TestPart2(t *testing.T) {
	input := []int{0,3,0,1,-3}

	result := part2(input)
	if result != 10 {
		t.Errorf("Expected 10, received %v", result)
	}
}
