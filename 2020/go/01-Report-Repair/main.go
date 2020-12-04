package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strconv"
)

func main() {

	lines := hlp.FileToLines("input.txt")

	numbers := []int{}

	for _, l := range lines {
		n, _ := strconv.Atoi(l)

		for _, n2 := range numbers {
			if n+n2 == 2020 {
				fmt.Println("First")
				fmt.Println(n * n2)
			}
			for _, n3 := range numbers {
				if n+n2+n3 == 2020 {
					fmt.Println("Second")
					fmt.Println(n * n2 * n3)
				}

			}

		}

		numbers = append(numbers, n)
	}
}
