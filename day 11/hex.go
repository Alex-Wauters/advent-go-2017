package main

import (
	"strings"
	"io/ioutil"
	"fmt"
)

func main() { //based on https://www.redblobgames.com/grids/hexagons/#coordinates
	input, _ := ioutil.ReadFile("input.txt")
	dirs := strings.Split(string(input), ",")
	c := cube{}
	max := 0
	for _, d := range dirs {
		c = c.Move(d)
		if c.Distance() > max {
			max = c.Distance()
		}
	}
	fmt.Printf("Part 1: %v, Part 2: %v ", c.Distance(), max)
}

type cube struct {
	x, y, z int
}

func (c cube) Move(dir string) cube {
	switch dir {
	case "n":
		return cube{c.x + 1, c.y, c.z - 1}
	case "ne":
		return cube{c.x + 1, c.y - 1, c.z}
	case "se":
		return cube{c.x, c.y - 1, c.z + 1}
	case "s":
		return cube{c.x - 1, c.y, c.z + 1}
	case "sw":
		return cube{c.x - 1, c.y + 1, c.z}
	case "nw":
		return cube{c.x, c.y + 1, c.z - 1}
	}
	panic("Don't recognize dir " + dir)
}

func (c cube) Distance() int {
	if Abs(c.x) >= Abs(c.y) && Abs(c.x) >= Abs(c.z) {
		return Abs(c.y) + Abs(c.z)
	} else if Abs(c.y) >= Abs(c.x) && Abs(c.y) >= Abs(c.z) {
		return Abs(c.x) + Abs(c.z)
	}
	return Abs(c.x) + Abs(c.y)
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
