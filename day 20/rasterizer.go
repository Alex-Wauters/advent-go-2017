package main

import (
	"os"
	"regexp"
	"bufio"
	"log"
	"strconv"
	"fmt"
)

type point struct {
	x, y, z int
}
type particle struct {
	p, v, a   *point
	IsRemoved bool
}

func (p point) Distance() int {
	return Abs(p.x) + Abs(p.y) + Abs(p.z)
}
func (p point) String() string {
	return fmt.Sprintf("%v,%v,%v", p.x, p.y, p.z)
}

func (p particle) Update() {
	p.v.x += p.a.x
	p.v.y += p.a.y
	p.v.z += p.a.z
	p.p.x += p.v.x
	p.p.y += p.v.y
	p.p.z += p.v.z
}

func main() {
	part1()
	part2()
}

func part1() {
	list := readInput()
	minNode := -1
	roundDistance := -1
	for i := 0; i < 1000; i++ {
		for k := range list {
			list[k].Update()
			if i == 999 && (roundDistance < 0 || list[k].p.Distance() < roundDistance) {
				minNode = k
				roundDistance = list[k].p.Distance()
			}
		}
	}
	fmt.Printf("The closest particle is %v \n", minNode)
}

func part2() {
	list := readInput()
	for i := 0; i < 1000; i++ {
		positions := make(map[string]int)
		for k := range list {
			if !list[k].IsRemoved {
				list[k].Update()
				pId, ok := positions[list[k].p.String()]
				if ok {
					list[k].IsRemoved = true
					list[pId].IsRemoved = true
				}
				positions[list[k].p.String()] = k
			}
		}
	}
	count := 0
	for _, p := range list {
		if !p.IsRemoved {
			count++
		}
	}
	fmt.Printf("There are %v particles which have not collided \n", count)
}

func readInput() (i []particle) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	particleRE := regexp.MustCompile(`^p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m := particleRE.FindAllStringSubmatch(line, -1)[0]
		i = append(i, particle{&point{toInt(m[1]), toInt(m[2]), toInt(m[3])}, &point{toInt(m[4]), toInt(m[5]), toInt(m[6])}, &point{toInt(m[7]), toInt(m[8]), toInt(m[9])}, false})
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

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
