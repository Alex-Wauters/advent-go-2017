package main

import (
	"os"
	"bufio"
	"regexp"
	"log"
	"strconv"
	"fmt"
)

func main() {
	root := part1(readInput())
	part2(root)
}

func part1(towers []*tower) *tower {
	n := towers[0]
	for n.parent != nil {
		n = n.parent
	}
	fmt.Println("Part 1: The root node is " + n.name)
	return n
}

func part2(root *tower) {
	unbalanced := root
	unbalancedCandidate := root
	expectedWeight := 0
	weightCandidate := 0
	for { // Keep going down until all the children have correct weights
		unbalancedCandidate, weightCandidate = unbalanced.children.unbalancedNode()
		if unbalancedCandidate == nil {
			fmt.Printf("The unbalanced node's total weight should be %v instead of %v. Node's own weight: %v \n", expectedWeight, unbalanced.totalWeight(), unbalanced.weight)
			return
		}
		unbalanced, expectedWeight = unbalancedCandidate, weightCandidate
	}
}

type tower struct {
	name          string
	weight        int
	childrenNames []string
	children      towers
	parent        *tower
}

type towers []*tower

func (t *tower) totalWeight() int {
	return t.weight + t.children.totalWeight()
}

func (t towers) find(n string) *tower {
	for _, tower := range t {
		if tower.name == n {
			return tower
		}
	}
	return nil
}

func (t towers) totalWeight() (sum int) {
	for _, tower := range t {
		sum += tower.weight + tower.children.totalWeight()
	}
	return sum
}

func (t towers) unbalancedNode() (n *tower, normalWeight int) {
	if len(t) == 0 {
		return nil, 0
	}
	for i := 1; i < len(t); i++ {
		if t[i].totalWeight() != t[0].totalWeight() {
			thirdNode := i + 1
			if len(t) <= thirdNode {
				thirdNode = i - 1
			}
			if t[0].totalWeight() == t[thirdNode].totalWeight() {
				return t[i], t[0].totalWeight() // i is the unbalanced node
			} else {
				return t[0], t[i].totalWeight() // First child does not match 2nd and third
			}
		}
	}
	return nil, 0
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func readInput() (i towers) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	towerRE := regexp.MustCompile(`^(?P<name>[a-z]+) \((?P<weight>[0-9]+)\)`)
	childrenRE := regexp.MustCompile(`(?P<child>[a-z]+)(?:,|\z)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		t := towerRE.FindAllStringSubmatch(line, -1)
		c := childrenRE.FindAllStringSubmatch(line, -1)
		tower := &tower{name: t[0][1], weight: toInt(t[0][2])}
		for _, child := range c {
			tower.childrenNames = append(tower.childrenNames, child[1])
		}
		i = append(i, tower)
	}
	for _, t := range i {
		for _, cName := range t.childrenNames {
			child := i.find(cName)
			t.children = append(t.children, child)
			child.parent = t
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return i
}
