package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"math"
)

func main() {
	lines := readInput()
	part1(lines)
	part2()
}

func part1(lines [][]string) {
	r := make(map[string]int)
	mulCount := 0
	for i := 0; i < len(lines) && i >= 0; i++ {
		instr := lines[i]
		switch instr[0] {
		case "jnz":
			if parseIntOrRegister(instr[1], r) != 0 {
				i += parseIntOrRegister(instr[2], r) - 1
			}
		case "set":
			r[instr[1]] = parseIntOrRegister(instr[2], r)
		case "sub":
			r[instr[1]] -= parseIntOrRegister(instr[2], r)
		case "mul":
			r[instr[1]] *= parseIntOrRegister(instr[2], r)
			mulCount++
		default:
			panic(instr[0])
		}
	}
	fmt.Printf("Mulcount: %v \n", mulCount)
}

func part2() {
	count := 0
	for b,c := 108100, 108100+17000 ; b <= c ; b+=17 {
		if !IsPrime(b) {
			count++
		}
	}
	fmt.Printf("H count: %v \n", count)
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func parseIntOrRegister(s string, registers map[string]int) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return registers[s]
	}
	return val
}

func readInput() (r [][]string) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		r = append(r, line)
	}
	return r
}
