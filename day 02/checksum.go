package main

import (
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	numbers := readInput()
	part1(numbers)
	part2(numbers)
}

func part1(numbers [][] int) {
	sum := 0
	for _, line := range numbers {
		sum = sum + lineChecksum1(line)
	}
	fmt.Printf("Part 1: The checksum is %v \n", sum)
}

func lineChecksum1(l []int) int {
	min, max := l[0], l[0]
	for _, n := range l {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return max - min
}

func part2(numbers [][] int) {
	sum := 0
	for _, line := range numbers {
		sum = sum + lineChecksum2(line)
	}
	fmt.Printf("Part 2: The checksum is %v \n", sum)
}

func lineChecksum2(l []int) int {
	for _, n := range l {
		for _, x := range l {
			if n != x && (n/x)*x == n {
				return n/x
			}
		}
	}
	panic(fmt.Sprintf("Could not find checksum for line %v", l))
}

func readInput() (r [][]int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := make([]int, 0)
		for _, s := range strings.Fields(scanner.Text()) {
			numbers = append(numbers, toInt(s))
		}
		r = append(r, numbers)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}