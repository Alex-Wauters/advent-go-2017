package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	lines := readInput()
	part1(lines)
	part2(lines)
}

func part1(lines [][]string) {
	r := make(map[string]int)
	for i := 0; i < len(lines) && i >= 0; i++ {
		instr := lines[i]
		switch instr[0] {
		case "jgz":
			if r[instr[1]] > 0 {
				i += parseIntOrRegister(instr[2], r) - 1
			}
		case "rcv":
			if r[instr[1]] > 0 {
				fmt.Printf("The last received value is %v \n", r["snd"])
				return
			}
		case "snd":
			r["snd"] = r[instr[1]]
		case "set":
			r[instr[1]] = parseIntOrRegister(instr[2], r)
		case "add":
			r[instr[1]] += parseIntOrRegister(instr[2], r)
		case "mul":
			r[instr[1]] *= parseIntOrRegister(instr[2], r)
		case "mod":
			r[instr[1]] %= parseIntOrRegister(instr[2], r)
		}
	}
}

func part2(lines [][]string) {
	a, b, final := make(chan int, 10000), make(chan int, 10000), make(chan bool)
	go program(0, lines, a, b)
	go program(1, lines, b, a)
	<-final // wait until deadlock
}

func program(id int, lines [][]string, in <-chan int, out chan<- int) {
	r := map[string]int{"p": id}
	count := 0
	for i := 0; i < len(lines) && i >= 0; i++ {
		instr := lines[i]
		switch instr[0] {
		case "jgz":
			if parseIntOrRegister(instr[1], r) > 0 {
				i += parseIntOrRegister(instr[2], r) - 1
			}
		case "rcv":
			r[instr[1]] = <-in
		case "snd":
			out <- r[instr[1]]
			if id == 1 {
				count++
				fmt.Printf("Program 1 has sent %v values. \n", count)
			}
		case "set":
			r[instr[1]] = parseIntOrRegister(instr[2], r)
		case "add":
			r[instr[1]] += parseIntOrRegister(instr[2], r)
		case "mul":
			r[instr[1]] *= parseIntOrRegister(instr[2], r)
		case "mod":
			r[instr[1]] %= parseIntOrRegister(instr[2], r)
		}
	}
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
