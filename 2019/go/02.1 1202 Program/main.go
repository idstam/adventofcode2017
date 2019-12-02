package main

import "strings"

import "fmt"

var mem []int

func main() {
	lines := fileToLines("input.txt")
	strs := strings.Split(lines[0], ",")

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			mem = StringToIntArray(strs)

			mem[1] = noun
			mem[2] = verb
			ptr := 0
			for ptr >= 0 {
				ptr = Exec(ptr)

			}

			if mem[0] == 19690720 {
				fmt.Printf("Noun %d  Verb %d  Answer %d\n", noun, verb, (100*noun)+verb)
			}
		}
	}
	//Too low: 29891
	//fmt.Println(mem[0])
}

func Exec(ptr int) int {
	switch mem[ptr] {
	case 1:
		mem[mem[ptr+3]] = mem[mem[ptr+1]] + mem[mem[ptr+2]]
		ptr += 4
	case 2:
		mem[mem[ptr+3]] = mem[mem[ptr+1]] * mem[mem[ptr+2]]
		ptr += 4
	case 99:
		return -1
	default:
		panic("Unknown OP code")

	}

	return ptr
}
