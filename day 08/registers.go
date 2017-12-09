package main

import (
	"os"
	"regexp"
	"bufio"
	"fmt"
	"log"
	"strconv"
)

var r map[string]int

func main() {
	input := readInput()
	part1(input)
	part2(input)
}

func part1(i []command) {
	r = make(map[string]int)
	for _, c := range i {
		if c.condition.isTrue() {
			if c.instruction == "dec" {
				c.value *= -1
			}
			r[c.reg] += c.value
		}
	}
	maxRegister, maxValue := maxRegisterValue()
	fmt.Printf("Part 1: The max register is %s with value %v \n", maxRegister, maxValue)
}

func part2(i []command) {
	r = make(map[string]int)
	maxValue := 0
	maxRegister := ""
	for _, c := range i {
		if c.condition.isTrue() {
			if c.instruction == "dec" {
				c.value *= -1
			}
			r[c.reg] += c.value
		}
		cRegister, cValue := maxRegisterValue()
		if cValue > maxValue {
			maxValue = cValue
			maxRegister = cRegister
		}
	}
	fmt.Printf("Part 2: The max register is %s with value %v \n", maxRegister, maxValue)
}

func maxRegisterValue() (register string, value int) {
	for k, v := range r {
		if v > value {
			register = k
			value = v
		}
	}
	return
}

type command struct {
	reg         string
	instruction string
	value       int
	condition   condition
}

type condition struct {
	reg         string
	conditional string
	value       int
}

func (c condition) isTrue() bool {
	reg, _ := r[c.reg]
	switch c.conditional {
	case "==":
		return reg == c.value
	case ">=":
		return reg >= c.value
	case "!=":
		return reg != c.value
	case "<=":
		return reg <= c.value
	case "<":
		return reg < c.value
	case ">":
		return reg > c.value
	}
	panic("Do not recognize conditional " + c.conditional)
}

func readInput() (i []command) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	commandRE := regexp.MustCompile(`^([a-z]+) (inc|dec) (-?\d+) if ([a-z]+) (>|<|==|!=|<=|>=) (-?\d+)$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := commandRE.FindAllStringSubmatch(line, -1)[0]
		i = append(i, command{matches[1], matches[2], toInt(matches[3]), condition{matches[4], matches[5], toInt(matches[6])}})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return i
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
