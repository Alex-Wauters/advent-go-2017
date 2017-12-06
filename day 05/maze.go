package main

import (
	"os"
	"bufio"
	"strconv"
	"log"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: We took %v steps \n", part1(readInput()))
	fmt.Printf("Part 2: We took %v steps", part2(readInput()))
}

func part1(i []int) int {
	steps := 0
	index := 0
	var jump int
	for {
		if index >= len(i) || index < 0 {
			return steps
		}
		jump = i[index]
		i[index]++
		index = index + jump
		steps++
	}
}

func part2(i []int) int {
	steps := 0
	index := 0
	var jump int
	for {
		if index >= len(i) || index < 0 {
			return steps
		}
		jump = i[index]

		if i[index] >= 3 {
			i[index]--
		} else {
			i[index]++
		}
		index = index + jump
		steps++
	}
}

func readInput() (i []int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())

		i = append(i, val)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return i
}