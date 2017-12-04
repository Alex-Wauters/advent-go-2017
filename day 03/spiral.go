package main

import "fmt"

const (
	right = iota
	up
	left
	down
)

var dict map[string]*point

func main() {
	part1()
	part2()
}

func part1() {
	n := 325489
	initLookup()
	last := createPoint(0, 0, 1)
	dir := down
	for i := 1; i < n; i++ {
		switch dir {
		case right:
			_, ok := lookup(last.x, last.y+1)
			if !ok {
				dir = up
			}
		case up:
			_, ok := lookup(last.x-1, last.y)
			if !ok {
				dir = left
			}
		case left:
			_, ok := lookup(last.x, last.y-1)
			if !ok {
				dir = down
			}
		case down:
			_, ok := lookup(last.x+1, last.y)
			if !ok {
				dir = right
			}
		}
		last = createPointDir(last, dir)
	}
	fmt.Printf("Part 1: Last point %s has distance %v \n", last.String(), last.Distance())
}

func initLookup() {
	dict = make(map[string]*point)
}

func createPointDir(previous *point, dir int) (p *point) {
	switch dir {
	case up:
		return createPoint(previous.x, previous.y+1, previous.n+1)
	case left:
		return createPoint(previous.x-1, previous.y, previous.n+1)
	case down:
		return createPoint(previous.x, previous.y-1, previous.n+1)
	case right:
		return createPoint(previous.x+1, previous.y, previous.n+1)
	}
	panic("Do not recognize direction " + string(dir))
}

func createPoint(x, y, n int) (p *point) {
	p = &point{x, y, n}
	dict[p.String()] = p
	return
}

type point struct {
	x, y, n int
}

func (p *point) String() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

func (p *point) Distance() int {
	return Abs(p.x) + Abs(p.y)
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func lookup(x, y int) (*point, bool) {
	p, ok := dict[fmt.Sprintf("%v,%v", x, y)]
	return p, ok
}
