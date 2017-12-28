package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
)

const (
	clean    = iota
	weakened
	infected
	flagged
)

var world2 map[string]int

func part2() {
	readInput2()
	x, y := 0, 0
	dir := up
	infectCount := 0
	for i := 0; i < 10000000; i++ {
		health := world2[coords(x, y)]
		switch health {
		case infected:
			dir = (dir + 1) % 4
		case weakened:
			infectCount++
		case clean:
			dir = (dir - 1) % 4
			if dir == -1 {
				dir = 3
			}
		case flagged:
			dir = (dir + 2) % 4
		}
		world2[coords(x, y)] = (health + 1) % 4
		x, y = walk(x, y, dir)
	}
	fmt.Printf("New infections: %v \n", infectCount)
}

func readInput2() {
	world2 = make(map[string]int)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	y := -12
	for scanner.Scan() {
		line := scanner.Text()
		for k, c := range line {
			state := clean
			if c == '#' {
				state = infected
			}
			world2[coords(-12+k, y)] = state
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
