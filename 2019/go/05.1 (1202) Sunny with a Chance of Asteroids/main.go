package main

import (
	"fmt"
	"log"
	"strings"
)

var mem []int

func main() {
	lines := fileToLines("example.txt")
	strs := strings.Split(lines[0], ",")

	mem = StringToIntArray(strs)

	ptr := 0
	for ptr >= 0 {
		ptr = Exec(ptr)

	}

	//Too low: 29891
	//fmt.Println(mem[0])
}

func Exec(ptr int) int {
	fullOp := mem[ptr]
	ma, mb, mc, op := ParseOpCode(fullOp)
	a, b, c := GetValues(ma, mb, mc, ptr)
	val := 0
	switch op {
	case 1:
		mem[mem[ptr+3]] = a + b
		ptr += 4
	case 2:
		mem[mem[ptr+3]] = a * b
		ptr += 4
	case 3:
		mem[mem[ptr+1]] = GetInput()
		ptr += 2
	case 4:
		val = GetValue(c, ptr)
		fmt.Printf("Output: %d \n", val)
	case 99:
		return -1
	default:
		panic("Unknown OP code")

	}

	return ptr
}
func GetValues(ma, mb, mc, adress int) (int, int, int) {
	a := 0
	b := 0
	c := 0
	if adress+3 < len(mem) {
		a = GetValue(ma, adress+3)
	}
	if adress+2 < len(mem) {
		b = GetValue(ma, adress+2)
	}
	if adress+1 < len(mem) {
		b = GetValue(ma, adress+1)
	}
	return a, b, c
}
func GetValue(mode, adress int) int {
	if adress >= len(mem) {
		log.Fatalf("index out of range index:%d len:%d", adress, len(mem))
	}
	if mode == 1 {
		return mem[adress]
	}
	if mem[adress] >= len(mem) {
		log.Fatalf("index out of range index:%d(%d) len:%d", mem[adress], adress, len(mem))
	}

	return mem[mem[adress]]
}
func GetInput() int {
	return 1
}

func ParseOpCode(in int) (int, int, int, int) {
	op := in % 100
	in -= op
	c := in % 1000
	in -= c
	b := in % 10000
	in -= b
	a := in % 100000
	return a, b, c, op
}
