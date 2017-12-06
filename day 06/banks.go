package main

import (
	"fmt"
	"strconv"
)

type banks []int

func main() {
	input := banks{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}
	fmt.Printf("Part 1: There are %v cycles \n", part1(input))
	fmt.Printf("Part 2: There are %v cycles \n", part2(input))

}

func part1(b banks) (c int) {
	seen := make(map[string]bool)
	seen[b.String()] = true
	for {
		b.Redistribute()
		c++
		if seen[b.String()] == true {
			return c
		}
		seen[b.String()] = true
	}
}

func part2(b banks) (c int) {
	seen := make(map[string]bool)
	seen[b.String()] = true
	var firstSeen string
	for {
		b.Redistribute()
		if firstSeen == b.String() {
			return c + 1
		}
		if firstSeen != "" {
			c++
		} else if seen[b.String()] == true {
			firstSeen = b.String()
		} else {
			seen[b.String()] = true
		}
	}
}

func (b banks) String() (r string) {
	for _, n := range b {
		r += strconv.Itoa(n) + ","
	}
	return r
}

func (b banks) Redistribute() {
	index := 0
	value := b[0]
	for i, v := range b {
		if v > value {
			value = v
			index = i
		}
	}
	b[index] = 0
	for i := 1; i <= value; i++ {
		b[(index+i)%len(b)]++
	}
}
