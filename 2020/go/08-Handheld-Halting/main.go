package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strings"
)

func main() {

	lines := hlp.FileToLines("input.txt", true)

	c := hlp.Computer{}
	c.Mode = "08"
	c.Boot(lines)
	fmt.Println("Break here if only running part one")

	swapped := -1
	for true {
		c.Reset()
		c.Run()
		if c.PC == len(c.Instructions) {
			fmt.Println("Part 2 Accumulator is ", c.Accumulator)
			break
		}

		if swapped > -1 {
			swapInstruction(c, swapped)
		}
		swapped = findNextSwapapbleInstruction(c, swapped+1)
		swapInstruction(c, swapped)

	}

	fmt.Println("ProgramCounter", c.PC)
}

func swapInstruction(c hlp.Computer, pc int) {
	if c.Instructions[pc].Name == "nop" {
		c.Instructions[pc].Name = "jmp"
	} else {
		c.Instructions[pc].Name = "nop"
	}

}
func findNextSwapapbleInstruction(c hlp.Computer, start int) int {
	for i := start; i < len(c.Instructions); i++ {
		if strings.Contains("nop jmp", c.Instructions[i].Name) {
			return i
		}
	}

	return -1
}
