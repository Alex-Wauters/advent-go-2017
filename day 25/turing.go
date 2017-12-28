package main

import "fmt"

const (
	a = iota
	b
	c
	d
	e
	f
)

func main() {
	tape := make(map[int]bool)
	cursor := 0
	state := a
	for i := 0; i < 12173597; i ++ {
		switch state {
		case a:
			if tape[cursor] {
				tape[cursor] = false
				state = c
				cursor--
			} else {
				tape[cursor] = true
				state = b
				cursor++
			}
		case b:
			if tape[cursor] {
				cursor++
				state = d
			} else {
				tape[cursor] = true
				cursor--
				state = a
			}
		case c:
			if tape[cursor] {
				tape[cursor] = false
				cursor--
				state = e
			} else {
				tape[cursor] = true
				cursor++
				state = a
			}
		case d:
			if tape[cursor] {
				state = b
			} else {
				state = a
			}
			tape[cursor] = !tape[cursor]
			cursor++
		case e:
			if tape[cursor] {
				state = c
			} else {
				state = f
			}
			tape[cursor] = true
			cursor--
		case f:
			if tape[cursor] {
				state = a
			} else {
				state = d
			}
			tape[cursor] = true
			cursor++
		}
	}
	count := 0
	for _, k := range tape {
		if k {
			count++
		}
	}
	fmt.Printf("Diagnostic: %v \n", count)
}
