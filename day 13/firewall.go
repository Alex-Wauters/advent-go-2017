package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"time"
)

var fw []firewall

func main() {
	fw = readInput()
	part1()
	part2()
}

func part1() {
	severity := 0
	for _, f := range fw {
		severity += f.Severity()
	}
	fmt.Printf("Severity is %v \n", severity)
}

func part2() {
	defer track(time.Now(),"Part 2")
	c, quit := make(chan int, 1), make(chan bool)
	go dispatch(c, quit)
	fmt.Printf("The minimum delay is %v",  <- c )
	close(quit)
}

func dispatch(c chan int, quit <-chan bool) {
	i := 0
	routines := 20
	for ; i < routines; i++ {
		go worker(i, routines, c, quit)
	}
}

func worker(i, increment int, c chan<- int, done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			if canPass(i) {
				c <- i
			}
			i += increment
		}
	}
}

func canPass(delay int) bool {
	for _, f := range fw {
		if !f.CanPassWithDelay(delay) {
			return false
		}
	}
	return true
}

type firewall struct {
	depth, length int
}

func (f firewall) Severity() int {
	if f.length == 2 {
		if f.depth % 2 == 0 {
			return f.depth * f.length
		}
		return 0
	}
	if f.depth % (2*(f.length - 1)) == 0 { return f.depth * f.length}
	return 0
}

func (f firewall) CanPassWithDelay(delay int) bool {
	if f.length == 2 {
		return (f.depth + delay) % 2 != 0
	}
	return (f.depth + delay) % (2*(f.length - 1)) != 0
}

func readInput() (f []firewall) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ": ")
		f = append(f, firewall{toInt(fields[0]), toInt(fields[1])})
	}
	return f
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func track(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf(" %s took %s \n", name, elapsed)
}