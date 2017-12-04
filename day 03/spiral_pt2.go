package main

import "fmt"

func part2() {
	n := 325489
	initLookup()
	last := createPoint(0, 0, 1)
	dir := down
	for last.n < n {
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
		last = createPointDir2(last, dir)
	}
	fmt.Printf("Part 2: The value is %v \n", last.n)
}

func createPointDir2(previous *point, dir int) (p *point) {
	switch dir {
	case up:
		return createPoint(previous.x, previous.y+1, surroundingValues(previous.x, previous.y+1))
	case left:
		return createPoint(previous.x-1, previous.y, surroundingValues(previous.x-1, previous.y))
	case down:
		return createPoint(previous.x, previous.y-1, surroundingValues(previous.x, previous.y-1))
	case right:
		return createPoint(previous.x+1, previous.y, surroundingValues(previous.x+1, previous.y))
	}
	panic("Do not recognize direction " + string(dir))
}

func surroundingValues(x, y int) (s int) {
	for _, nx := range []int{-1, 0, 1} {
		for _, ny := range []int{-1, 0, 1} {
			p, ok := lookup(x+nx, y+ny)
			if ok {
				s = s + p.n
			}
		}
	}
	return
}
