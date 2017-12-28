package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
)

const (
	left  = iota
	up
	right
	down
)

var world map[string]bool // isInfected

func main() {
	part1()
	part2()
}

func part1() {
	readInput()
	x, y := 0, 0
	dir := up
	infectCount := 0
	for i := 0; i < 10000; i++ {
		if world[coords(x, y)] {
			dir = (dir + 1) % 4
		} else {
			dir = (dir - 1) % 4
			if dir == -1 {
				dir = 3
			}
		}
		isInfected := world[coords(x, y)]
		world[coords(x, y)] = !isInfected
		if !isInfected {
			infectCount++
		}
		x, y = walk(x, y, dir)
	}
	fmt.Printf("New infections: %v \n", infectCount)
}

func walk(x, y, dir int) (int, int) {
	switch dir {
	case up:
		return x, y - 1
	case down:
		return x, y + 1
	case left:
		return x - 1, y
	case right:
		return x + 1, y
	}
	panic(fmt.Sprintf("Unrecognized direction %v", dir))
}

func readInput() {
	world = make(map[string]bool)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	y := -12
	for scanner.Scan() {
		line := scanner.Text()
		for k, c := range line {
			world[coords(-12+k, y)] = c == '#'
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func coords(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}
