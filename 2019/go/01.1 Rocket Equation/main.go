package main

import "fmt"

func main() {
	lines := fileToLines("input.txt")

	ints := StringToIntArray(lines)
	ints = DivideByOnIntArray(ints, 3)
	ints = SubtractOnIntArray(ints, 2)
	sum := SumIntArray(ints)
	fmt.Println(sum)
}
