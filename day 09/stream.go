package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
)

func main() {
	rootGroup := readInput()
	fmt.Printf("Part 1: %v, Part 2: %v \n", rootGroup.Score(0), rootGroup.TotalGarbage())
}

type group struct {
	children []group
	garbage  int
}

func (g *group) Score(p int) (s int) {
	s = p + 1
	for _, c := range g.children {
		s += c.Score(p + 1)
	}
	return s
}

func (g *group) TotalGarbage() (s int) {
	s = g.garbage
	for _, c := range g.children {
		s += c.TotalGarbage()
	}
	return s
}

func createGroup(line string, offset int) (g group, k int) {
	for {
		if offset >= len(line) {
			return g, offset
		}
		switch line[offset] {
		case '}':
			return g, offset
		case '{':
			child, k := createGroup(line, offset+1)
			offset = k
			g.children = append(g.children, child)
		case '<':
			k, garbageCount := createGarbage(line, offset+1)
			offset = k
			g.garbage += garbageCount
		}
		offset++
	}
}

func createGarbage(line string, i int) (offset int, score int) {
	ignoreNext := false
	for {
		if !ignoreNext {
			switch line[i] {
			case '>':
				return i, score
			case '!':
				ignoreNext = true
			default:
				score++
			}
		} else {
			ignoreNext = false
		}
		i++
	}
}

func readInput() (*group) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	root, _ := createGroup(line, 1)
	return &root
}
