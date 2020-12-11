package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"sort"
)

var ints []int
var memo map[int]int

func main() {

	lines := hlp.FileToLines("input.txt", true)
	ints = hlp.StringToIntArray(lines)
	ints = append(ints, 0)
	memo = map[int]int{}
	sort.Ints(ints)

	lastWattage := 0
	diffs := map[int]int{}
	for _, i := range ints {
		diffs[i-lastWattage] = diffs[i-lastWattage] + 1
		lastWattage = i
	}
	diffs[3] = diffs[3] + 1

	fmt.Printf("%v \n", diffs)
	fmt.Println("Part 1", diffs[1]*diffs[3])

	totalValidArrangements := next(0, 0)
	fmt.Println("Part 2", totalValidArrangements)
}
func next(from int, level int) int {
	numbers := ints
	if memo[from] > 0 {
		return memo[from]
	}

	if from == len(numbers)-1 {
		memo[from] = 1
		return 1
	}

	sum := 0
	for i := from + 1; i < from+4; i++ {

		if i == len(numbers) {
			break
		}
		if numbers[i] < numbers[from]+4 {
			sum += next(i, level+1)
		}
	}

	memo[from] = sum
	return sum
}
