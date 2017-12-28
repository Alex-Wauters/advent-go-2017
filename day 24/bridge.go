package main

import "fmt"

type connector struct {
	port1, port2 int
}
type bridge []connector
type state struct {
	bridge, rest bridge
}

var maxStrength int
var maxLength int

func main() {
	maxStrength = 0
	components := []connector{{31, 13}, {34, 4}, {49, 49}, {23, 37}, {47, 45}, {32, 4}, {12, 35}, {37, 30}, {41, 48}, {0, 47}, {32, 30}, {12, 5}, {37, 31}, {7, 41}, {10, 28}, {35, 4}, {28, 35}, {20, 29}, {32, 20}, {31, 43}, {48, 14}, {10, 11}, {27, 6}, {9, 24}, {8, 28}, {45, 48}, {8, 1}, {16, 19}, {45, 45}, {0, 4}, {29, 33}, {2, 5}, {33, 9}, {11, 7}, {32, 10}, {44, 1}, {40, 32}, {2, 45}, {16, 16}, {1, 18}, {38, 36}, {34, 24}, {39, 44}, {32, 37}, {26, 46}, {25, 33}, {9, 10}, {0, 29}, {38, 8}, {33, 33}, {49, 19}, {18, 20}, {49, 39}, {18, 39}, {26, 13}, {19, 32}}
	solve(state{[]connector{}, components}, 0)
	fmt.Printf("Max strength is %v \n", maxStrength)
}

func solve(s state, match int) {
	for i, c := range s.rest {
		if c.port1 == match || c.port2 == match {
			left := c.port1
			if c.port1 == match {
				left = c.port2
			}
			newBridge := make(bridge, len(s.bridge))
			newRest := make(bridge, len(s.rest))
			copy(newRest, s.rest)
			copy(newBridge, s.bridge)
			solve(state{append(newBridge, c), append(newRest[:i], newRest[i+1:]...)}, left)
		}
	}
	//part1(s)
	part2(s)
}

func part1(s state) {
	if strength := s.bridge.Strength(); strength > maxStrength {
		maxStrength = strength
	}
}

func part2(s state) {
	if len(s.bridge) > maxLength {
		maxLength = len(s.bridge)
		maxStrength = s.bridge.Strength()
	} else if len(s.bridge) == maxLength {
		if s.bridge.Strength() > maxStrength {
			maxStrength = s.bridge.Strength()
		}
	}
}

func (b bridge) Strength() (s int) {
	for _, c := range b {
		s += c.port1 + c.port2
	}
	return s
}
