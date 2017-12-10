package main

import (
	"fmt"
)

func main() {
	hash := knotHash(generateInput(), []int{206, 63, 255, 131, 65, 80, 238, 157, 254, 24, 133, 2, 16, 0, 1, 3}, 1)
	fmt.Printf("Part 1: %v \n", hash[0]*hash[1])
	fmt.Println(hashToString(denseHash(knotHash(generateInput(), convertLengths("206,63,255,131,65,80,238,157,254,24,133,2,16,0,1,3"), 64))))
}

func knotHash(input, lengths []int, rounds int) []int {
	skip := 0
	index := 0
	for round := 1; round <= rounds; round++ {
		for _, l := range lengths {
			els := reverseElements(input, index, l)
			for i, e := range els {
				input[(index+i)%len(input)] = e
			}
			index = (index + skip + l) % len(input)
			skip++
		}
	}
	return input
}

func reverseElements(list []int, index, size int) (r []int) {
	for i := 0; i < size; i++ {
		r = append(r, list[(index+i)%len(list)])
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func denseHash(r []int) (d []int) {
	for i := 0; i < len(r); i += 16 {
		d = append(d, r[i]^r[i+1]^r[i+2]^r[i+3]^r[i+4]^r[i+5]^r[i+6]^r[i+7]^r[i+8]^r[i+9]^r[i+10]^r[i+11]^r[i+12]^r[i+13]^r[i+14]^r[i+15])
	}
	return d
}

func hashToString(hash []int) (result string) {
	for _, h := range hash {
		hex := fmt.Sprintf("%x", h)
		if len(hex) == 1 {
			hex = "0" + hex
		}
		result += hex
	}
	return result
}

func generateInput() (r []int) {
	r = make([]int, 256, 256)
	for i := 0; i < 256; i++ {
		r[i] = i
	}
	return r
}

func convertLengths(i string) (result []int) {
	for _, r := range i {
		result = append(result, int(r))
	}
	result = append(result, []int{17, 31, 73, 47, 23}...)
	return result
}
