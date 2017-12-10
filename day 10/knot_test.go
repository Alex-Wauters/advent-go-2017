package main

import "testing"

func TestKnotHash(t *testing.T) {
	i := []int{0, 1, 2, 3, 4}
	l := []int{3, 4, 1, 5}
	v := knotHash(i, l, 1)
	//3 4 2 1 [0]
	if !(v[0] == 3 && v[1] == 4) {
		t.Errorf("Expected %v, instead got %v", []int{3, 4, 2, 1, 0}, v)
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
	}
	for _, c := range cases {
		result := hashToString(denseHash(knotHash(generateInput(), convertLengths(c.input), 64)))
		if result != c.expected {
			t.Errorf("Expected %s for input %s, instead got %s \n", c.expected, c.input, result)
		}
	}

}
