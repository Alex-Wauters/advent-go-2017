package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

var yellowPages map[string]*house

type house struct {
	visited           bool
	connections       []*house
	connectionsString []string
}

func main() {
	fmt.Printf("Part 1: There are %v houses connected to root \n", readInput().Unvisited())
	groups := 1
	for _, h := range yellowPages {
		if !h.visited {
			groups++
			h.Unvisited()
		}
	}
	fmt.Printf("Part 2: There are %v different groups \n", groups)
}

func (h *house) Unvisited() (sum int) {
	if h.visited {
		return 0
	}
	sum = 1
	h.visited = true
	for _, c := range h.connections {
		sum += c.Unvisited()
	}
	return sum
}

func readInput() (root *house) {
	yellowPages = make(map[string]*house)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		house := house{}
		fields := strings.Split(line, " <-> ")
		id, connections := fields[0], fields[1]
		house.connectionsString = strings.Split(connections, ", ")
		yellowPages[id] = &house
	}
	for _, h := range yellowPages {
		for _, s := range h.connectionsString {
			h.connections = append(h.connections, yellowPages[s])
		}
	}
	return yellowPages["0"]
}
