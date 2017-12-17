package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

const (
	spin     = iota
	exchange
	partner
)

var p programs

func main() {
	seen := make(map[string]bool)
	p = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}
	moves := readDances()
	for i := 1; i <= 1000000000; i++ {
		for _, move := range moves {
			dance(move)
		}
		if _, ok := seen[p.String()]; ok {
			for k := 0; k < (1000000000 % (i-1)) - 1; k ++ {
				for _, move := range moves {
					dance(move)
				}
			}
			fmt.Println(p.String())
			return
		}
		seen[p.String()] = true
	}
}

func (p programs) String() (s string) {
	for _, r := range p {
		s += string(r)
	}
	return s
}

type programs []rune
type move struct {
	instruction int
	n           []int
	r           []rune
}

func dance(m move) {
	switch m.instruction {
	case spin:
		p = append(p[16-m.n[0]:], p[:16-m.n[0]]...)
	case exchange:
		p[m.n[0]], p[m.n[1]] = p[m.n[1]], p[m.n[0]]
	case partner:
		x, y := p.Find(m.r[0]), p.Find(m.r[1])
		p[x], p[y] = p[y], p[x]
	}
}

func (p programs) Find(r rune) int {
	for i, program := range p {
		if program == r {
			return i
		}
	}
	panic("Could not find program " + string(r))
}

func readDances() (moves []move) {
	input, _ := ioutil.ReadFile("input.txt")
	instr := strings.Split(string(input), ",")
	for _, i := range instr {
		if i[0] == 's' {
			moves = append(moves, move{instruction: spin, n: []int{toInt(i[1:])}})
		} else if i[0] == 'p' {
			moves = append(moves, move{instruction: partner, r: []rune{rune(i[1]), rune(i[3])}})
		} else if i[0] == 'x' {
			numbers := strings.Split(i[1:], "/")
			moves = append(moves, move{instruction: exchange, n: []int{toInt(numbers[0]), toInt(numbers[1])}})
		}
	}
	return moves
}

func toInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
