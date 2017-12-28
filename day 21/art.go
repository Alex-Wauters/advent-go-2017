package main

import (
	"regexp"
	"os"
	"bufio"
	"log"
	"fmt"
)

var rules []rule

type rule struct {
	size   int
	match  string
	output string
}

type block struct {
	input, output string
	x, y, size    int
}

func main() {
	rules = readInput()
	grid := [][]rune{[]rune(".#."), []rune("..#"), []rune("###")}
	for i := 0; i < 5; i ++ { // change to 18 for part 2
		blocks := make([]block, 0)
		newSize := 0
		if len(grid[0])%2 == 0 {
			for y := 0; y < len(grid[0]); y = y + 2 {
				for x := 0; x < len(grid[0]); x = x + 2 {
					b := block{input: string(grid[y][x:x+2]) + string(grid[y+1][x:x+2]), size: 2, x: x * 3 / 2, y: y * 3 / 2}
					b.output = match(b)
					blocks = append(blocks, b)
				}
			}
			newSize = (len(grid[0]) / 2) * 3
		} else {
			for y := 0; y < len(grid[0]); y = y + 3 {
				for x := 0; x < len(grid[0]); x = x + 3 {
					b := block{input: string(grid[y][x:x+3]) + string(grid[y+1][x:x+3]) + string(grid[y+2][x:x+3]), size: 3, x: x * 4 / 3, y: y * 4 / 3}
					b.output = match(b)
					blocks = append(blocks, b)
				}
			}
			newSize = (len(grid[0]) / 3) * 4
		}
		grid = createGrid(newSize, blocks)
	}
	count := 0
	for _, row := range grid {
		for _, c := range row {
			if c == '#' {
				count++
			}
		}
	}
	fmt.Printf("There are %v pixels on \n", count)
}

func createGrid(size int, blocks []block) [][]rune {
	g := make([][]rune, size)
	for i := range g {
		g[i] = make([]rune, size)
	}
	for _, b := range blocks {
		if b.size == 2 {
			g[b.y+0][b.x+0], g[b.y+0][b.x+1], g[b.y+0][b.x+2], g[b.y+1][b.x+0], g[b.y+1][b.x+1], g[b.y+1][b.x+2], g[b.y+2][b.x+0], g[b.y+2][b.x+1], g[b.y+2][b.x+2] = r(b.output, 0), r(b.output, 1), r(b.output, 2), r(b.output, 3), r(b.output, 4), r(b.output, 5), r(b.output, 6), r(b.output, 7), r(b.output, 8)
		} else {
			g[b.y+0][b.x+0], g[b.y+0][b.x+1], g[b.y+0][b.x+2], g[b.y+0][b.x+3], g[b.y+1][b.x+0], g[b.y+1][b.x+1], g[b.y+1][b.x+2], g[b.y+1][b.x+3], g[b.y+2][b.x+0], g[b.y+2][b.x+1], g[b.y+2][b.x+2], g[b.y+2][b.x+3], g[b.y+3][b.x+0], g[b.y+3][b.x+1], g[b.y+3][b.x+2], g[b.y+3][b.x+3] = r(b.output, 0), r(b.output, 1), r(b.output, 2), r(b.output, 3), r(b.output, 4), r(b.output, 5), r(b.output, 6), r(b.output, 7), r(b.output, 8), r(b.output, 9), r(b.output, 10), r(b.output, 11), r(b.output, 12), r(b.output, 13), r(b.output, 14), r(b.output, 15)
		}
	}
	return g
}

func r(s string, i int) rune {
	return rune(s[i])
}

func match(b block) string {
	for _, r := range rules {
		if r.size == b.size {
			for _, p := range permutations(b.input) {
				if r.match == p {
					return r.output
				}
			}
		}
	}
	panic("No match found for " + b.input)
}

func permutations(match string) (r []string) {
	last := match
	for i := 0; i < 4; i ++ {
		last = rotate(last)
		r = append(r, last)
	}
	last = flip(last)
	for i := 0; i < 4; i ++ {
		last = rotate(last)
		r = append(r, last)
	}
	return r
}

func flip(s string) string {
	r := []rune(s)
	if len(s)%2 == 0 {
		return string([]rune{r[1], r[0], r[3], r[2]})
	} else {
		return string([]rune{r[2], r[1], r[0], r[5], r[4], r[3], r[8], r[7], r[6]})
	}
}

func rotate(s string) string {
	r := []rune(s)
	if len(s)%2 == 0 {
		return string([]rune{r[2], r[0], r[3], r[1]})
	} else {
		return string([]rune{r[6], r[3], r[0], r[7], r[4], r[1], r[8], r[5], r[2]})
	}
}

func readInput() (i []rule) {
	twoRE := regexp.MustCompile(`^([.#]{2})/([.#]{2}) => ([.#]{3})/([.#]{3})/([.#]{3})$`)
	threeRE := regexp.MustCompile(`^([.#]{3})/([.#]{3})/([.#]{3}) => ([.#]{4})/([.#]{4})/([.#]{4})/([.#]{4})$`)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r := rule{size: 2}
		m := twoRE.FindAllStringSubmatch(line, -1)
		if m == nil {
			m = threeRE.FindAllStringSubmatch(line, -1)
			r.size = 3
			r.match = m[0][1] + m[0][2] + m[0][3]
			r.output = m[0][4] + m[0][5] + m[0][6] + m[0][7]
		} else {
			r.match = m[0][1] + m[0][2]
			r.output = m[0][3] + m[0][4] + m[0][5]
		}
		i = append(i, r)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return i
}
