package main

import (
	"os"
	"bufio"
	"strings"
	"log"
)

type passphrase []string

func main() {
	input := readInput()
	part1(input)
	part2(input)
}

func part1(i []passphrase) {
	sum := 0
	for _, p := range i {
		if p.noDuplicates() {
			sum++
		}
	}
	log.Printf("Part 1: There are %v valid passphrases \n", sum)
}

func part2(i []passphrase) {
	sum := 0
	for _, p := range i {
		if p.noAnagrams() {
			sum++
		}
	}
	log.Printf("Part 2: There are %v valid passphrases \n", sum)
}

func (p passphrase) noDuplicates() bool {
	for i1, n1 := range p {
		for i2, n2 := range p {
			if i1 != i2 && n1 == n2 {
				return false
			}
		}
	}
	return true
}

func (p passphrase) noAnagrams() bool {
	for i1, n1 := range p {
		for i2, n2 := range p {
			if i1 != i2 && len(n1) == len(n2) && isAnagram(n1, n2) {
				return false
			}
		}
	}
	return true
}

func isAnagram(a, b string) bool {
	remaining := strings.Split(b, "")
Step:
	for _, l := range strings.Split(a, "") {
		for i, r := range remaining {
			if l == r {
				remaining = append(remaining[:i], remaining[i+1:]...)
				continue Step
			}
		}
		return false
	}
	return len(remaining) == 0
}

func readInput() (r []passphrase) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, strings.Fields(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}
