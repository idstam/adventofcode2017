package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strings"
)

func main() {

	input := strings.Split("6,4,12,1,20,0,16", ",") //Actual input
	//input := strings.Split("6,4,12,1,20,0,16", ",")

	numbersPrev := map[int]int{}
	numbersLast := map[int]int{}
	numbersCount := map[int]int{}

	last := 0
	for i, s := range input {
		n := hlp.Atoi(s)
		numbersPrev[n] = i + 1
		numbersLast[n] = i + 1
		numbersCount[n] = 1
		last = n
	}

	for turn := len(numbersCount) + 1; turn <= 2020; turn++ {

		lastCount := numbersCount[last]

		if lastCount == 1 {
			last = 0
		} else {
			last = numbersLast[last] - numbersPrev[last]
		}

		numbersPrev[last] = numbersLast[last]
		numbersLast[last] = turn
		numbersCount[last] = numbersCount[last] + 1

	}

	fmt.Println("@2020", last)

	numbersPrev = map[int]int{}
	numbersLast = map[int]int{}
	numbersCount = map[int]int{}

	last = 0
	for i, s := range input {
		n := hlp.Atoi(s)
		numbersPrev[n] = i + 1
		numbersLast[n] = i + 1
		numbersCount[n] = 1
		last = n
	}

	for turn := len(numbersCount) + 1; turn <= 30000000; turn++ {

		lastCount := numbersCount[last]

		if lastCount == 1 {
			last = 0
		} else {
			last = numbersLast[last] - numbersPrev[last]
		}

		numbersPrev[last] = numbersLast[last]
		numbersLast[last] = turn
		numbersCount[last] = numbersCount[last] + 1

	}

	fmt.Println("@30000000", last)
}
