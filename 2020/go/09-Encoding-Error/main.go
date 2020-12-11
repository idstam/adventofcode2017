package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
)

func main() {

	lines := hlp.FileToLines("input.txt", true)

	preLen := 25

	numbers := []*Number{}
	for _, l := range lines {
		newNumber := &Number{Value: hlp.Atoi(l), Sums: map[int]bool{}}

		if len(numbers) == preLen {

			found := false
			for _, n := range numbers {
				if n.Sums[newNumber.Value] {
					found = true
				}
			}
			if !found {
				fmt.Println("Invalid number", newNumber.Value)
				findBlock(lines, newNumber.Value)
				break
			}

			numbers = numbers[1:]
		}

		for _, n := range numbers {
			n.Sums[n.Value+newNumber.Value] = true
		}

		numbers = append(numbers, newNumber)
	}

	fmt.Println("Done")
}

func findBlock(lines []string, needle int) {
	start := 0
	end := 0
	numbers := hlp.StringToIntArray(lines)

	for true {
		sum := 0
		min := needle
		max := 0
		for i := start; i <= end; i++ {
			sum += numbers[i]
			min = hlp.IntMin(min, numbers[i])
			max = hlp.IntMax(max, numbers[i])
		}
		if sum == needle {
			fmt.Println(start, end, min, max, min+max)
			break
		}
		if sum > needle {
			start++
		}
		if sum < needle {
			end++
		}
	}
}

type Number struct {
	Value int
	Sums  map[int]bool
}
