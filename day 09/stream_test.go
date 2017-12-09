package main

import "testing"

func TestScore(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}
	for _, c := range cases {
		root, _ := createGroup(c.input, 1)
		if score := root.Score(0); score != c.expected {
			t.Errorf("Expected score %v, instead got %v for input %s", c.expected, score, c.input)
		}
	}
}

func TestGarbage(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"<>", 0},
		{"<random characters>", 17},
		{"<<<<>", 3},
		{"<{!>}>", 2},
		{"<!!>", 0},
		{"<!!!>>", 0},
		{`<{o"i!a,<{i<a>`, 10},
	}
	for _, c := range cases {
		if _, score := createGarbage(c.input, 1); score != c.expected {
			t.Errorf("Expected garbage score %v, instead got %v for input %s", c.expected, score, c.input)
		}
	}
}
