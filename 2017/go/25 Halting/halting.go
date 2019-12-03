package main

import (
	"fmt"
)

var mem []int

func main() {

	maxStep := 12586542
	mem = make([]int, maxStep*2)
	state := "A"
	ptr := maxStep

	for i := 0; i < maxStep; i++ {
		ptr, state = Exec(ptr, state)
	}

	checkSum := SumIntArray(mem)
	fmt.Println(checkSum)

}

func Exec(ptr int, state string) (int, string) {
	switch state {
	case "A":
		if mem[ptr] == 0 {
			mem[ptr] = 1
			ptr++
			state = "B"
		} else {
			mem[ptr] = 0
			ptr--
			state = "B"
		}
	case "B":
		if mem[ptr] == 0 {
			mem[ptr] = 0
			ptr++
			state = "C"
		} else {
			mem[ptr] = 1
			ptr--
			state = "B"
		}
	case "C":
		if mem[ptr] == 0 {
			mem[ptr] = 1
			ptr++
			state = "D"
		} else {
			mem[ptr] = 0
			ptr--
			state = "A"
		}
	case "D":
		if mem[ptr] == 0 {
			mem[ptr] = 1
			ptr--
			state = "E"
		} else {
			mem[ptr] = 1
			ptr--
			state = "F"
		}
	case "E":
		if mem[ptr] == 0 {
			mem[ptr] = 1
			ptr--
			state = "A"
		} else {
			mem[ptr] = 0
			ptr--
			state = "D"
		}
	case "F":
		if mem[ptr] == 0 {
			mem[ptr] = 1
			ptr++
			state = "A"
		} else {
			mem[ptr] = 1
			ptr--
			state = "E"
		}
	default:
		panic("Unknown state")
	}

	return ptr, state
}
