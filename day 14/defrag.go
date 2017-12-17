package main

import (
	"github.com/AE-nv/aedvent-code-2017/dependencies/Alex/knot"
	"fmt"
	"encoding/hex"
	"log"
)

func main() {
	input := generateInput("vbqugkhl")
	part1(input)
	part2(input)
}

type grid [][]*square

type square struct {
	used   bool
	region int
}

func part1(i grid) {
	sum := 0
	for _, row := range i {
		for _, c := range row {
			if c.used {
				sum++
			}
		}
	}
	fmt.Printf("Part 1: %v \n", sum)
}

func part2(i grid) {
	regions := 0
	for x, row := range i {
		for y, sq := range row {
			if sq.used && sq.region == 0 {
				regions++
				i.Visit(x, y, regions)
			}
		}
	}
	fmt.Printf("Part 2: %v regions \n", regions)
}

func (g grid) Visit(x, y, region int) {
	if x < 0 || y < 0 || x >= 128 || y >= 128 || !g[x][y].used || g[x][y].region != 0 {
		return
	}
	g[x][y].region = region
	g.Visit(x-1, y, region)
	g.Visit(x, y-1, region)
	g.Visit(x+1, y, region)
	g.Visit(x, y+1, region)
}

func HexToBin(s string) (result string) {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range decoded {
		result += fmt.Sprintf("%08b", d)

	}
	return result
}

func generateInput(p string) (rows grid) {
	for i := 0; i < 128; i++ {
		line := HexToBin(knot.Hash(fmt.Sprintf("%v-%v", p, i)))
		row := make([]*square, 128)
		for i, c := range line {
			sq := &square{}
			if c == '1' {
				sq.used = true
			}
			row[i] = sq
		}
		rows = append(rows, row)
	}
	return rows
}
