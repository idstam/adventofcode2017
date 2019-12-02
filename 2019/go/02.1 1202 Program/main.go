package main

import "strings"

import "fmt"

var mem []int

func main() {
	lines := fileToLines("input.txt")
	strs := strings.Split(lines[0], ",")
	mem = StringToIntArray(strs)

	mem[1] = 12
	mem[2] = 2
	ptr := 0
	for ptr >= 0 {
		ptr = Exec(ptr)

	}

	//Too low: 29891
	fmt.Println(mem[0])
}

func Exec(ptr int) int {
	switch mem[ptr] {
	case 1:
		mem[mem[ptr+3]] = mem[mem[ptr+1]] + mem[mem[ptr+2]]
	case 2:
		mem[mem[ptr+3]] = mem[mem[ptr+1]] * mem[mem[ptr+2]]
	case 99:
		return -1
	default:
		panic("Unknown OP code")

	}

	return ptr + 4
}
