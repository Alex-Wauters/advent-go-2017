package main

import "testing"

func TestDistance(t *testing.T) {
	cases := []struct {
		dirs []string
		expected int
	}{
		{[]string{"ne","ne","ne"}, 3},
		{[]string{"ne","ne","sw","sw"}, 0},
		{[]string{"ne","ne","s","s"}, 2},
	}
	for _, c := range cases {
		hex := cube{}
		for _, d := range c.dirs {
			hex = hex.Move(d)
		}
		if distance := hex.Distance(); distance != c.expected {
			t.Errorf("Expected for %v the distance value %v instead of %v", c.dirs, c.expected, distance)
		}
	}
}
