package main

import (
	"os"
	"bufio"
	"fmt"
	"unicode"
)

const (
	up    = iota
	down
	left
	right
)

var grid map[string]rune

func main() {
	readInput()
	x, y := 191, 0
	dir := down
	k := 0
	for {
		k++
		x, y = walk(x, y, dir)
		r := grid[coords(x, y)]
		switch {
		case unicode.IsLetter(grid[coords(x, y)]):
			fmt.Print(string(r))
		case r == '|':
		case r == '-':
			// No op
		case r == ' ':
			fmt.Printf("\n We jumped out after %v steps \n", k)
			return
		case r == '+':
			if dir == up || dir == down {
				n, ok := grid[coords(walk(x, y, left))]
				if !unicode.IsSpace(n) && ok {
					dir = left
				} else {
					dir = right
				}
			} else {
				n, ok := grid[coords(walk(x, y, up))]
				if !unicode.IsSpace(n) && ok {
					dir = up
				} else {
					dir = down
				}
			}
		default:
			fmt.Errorf("don't recognize %v for x: %v and y: %v", grid[coords(x, y)], x, y)
		}
	}
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
	panic("Durp")
}

func coords(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func readInput() {
	grid = make(map[string]rune)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	y := -1
	for scanner.Scan() {
		y++
		line := scanner.Text()
		for i, r := range line {
			grid[coords(i, y)] = r
		}
	}
}
