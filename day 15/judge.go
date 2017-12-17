package main

import "fmt"

func main() {
	a, b, solution := make(chan uint16, 1), make(chan uint16, 1), make(chan uint32, 1)
	part1(a, b, solution)
	part2(a, b, solution)

	fmt.Printf("There are %v samples which are the same ", <-solution)
}

func part1(a, b chan uint16, solution chan uint32) {
	go Generator(512, 16807, 2147483647, a, func(i int) bool { return true })
	go Generator(191, 48271, 2147483647, b, func(i int) bool { return true })
	go Judge(a, b, solution, 40000000)
}

func part2(a, b chan uint16, solution chan uint32) {
	go Generator(512, 16807, 2147483647, a, func(i int) bool { return i%4 == 0 })
	go Generator(191, 48271, 2147483647, b, func(i int) bool { return i%8 == 0 })
	go Judge(a, b, solution, 5000000)
}

func Judge(a, b chan uint16, solution chan uint32, rounds uint32) {
	count := uint32(0)
	for i := uint32(1); i <= rounds; i++ {
		if <-a == <-b {
			count++
		}
	}
	solution <- count
}

func Generator(start, multiplier, remainder int, output chan<- uint16, filter func(int) bool) {
	value := start
	for {
		value = (value * multiplier) % remainder
		if filter(value) {
			output <- uint16(value)
		}
	}
}
