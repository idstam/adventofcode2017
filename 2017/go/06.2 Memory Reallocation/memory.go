package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "11 11 13 7 0 15 5 5 4 4 1 1 7 1 15 11"
	//input := "2 4 1 2"

	stringNumbers := strings.Split(input, " ")
	numbers := make([]int, len(stringNumbers), len(stringNumbers))

	for index, str := range stringNumbers {
		numbers[index], _ = strconv.Atoi(str)
	}

	previous := make(map[string]int)
	ctr := 1
	for {
		hash := makeHash(numbers)
		//fmt.Println(hash)
		if previous[hash] == 0 {
			previous[hash] = ctr
		} else {
			fmt.Println(ctr - previous[hash])
			break
		}

		maxIndex := findMaxIndex(numbers)
		numbers = reallocate(maxIndex, numbers)
		ctr++
	}
}

func reallocate(maxIndex int, numbers []int) []int {
	val := numbers[maxIndex]
	numbers[maxIndex] = 0
	i := maxIndex + 1
	for {
		if i >= len(numbers) {
			i = i - len(numbers)
		}

		numbers[i] = numbers[i] + 1
		val--

		if val == 0 {
			return numbers
		}
		i++
	}

}

func makeHash(numbers []int) string {
	ret := ""
	for _, v := range numbers {
		ret += fmt.Sprintf("%x", v)
	}
	return ret
}

func findMaxIndex(numbers []int) int {
	max := -1
	maxIndex := -1

	for i, v := range numbers {
		if v > max && i > maxIndex {
			max = v
			maxIndex = i
		}
	}
	return maxIndex
}
