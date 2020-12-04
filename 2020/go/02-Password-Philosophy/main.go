package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strconv"
	"strings"
)

func main() {

	lines := hlp.FileToLines("input.txt")

	validPasswords := map[string]bool{}
	validPasswords2 := map[string]bool{}
	validCount := 0
	validCount2 := 0
	for _, l := range lines {

		tokens := strings.FieldsFunc(l, func(c rune) bool { return strings.Contains("- :", string(c)) })
		password := tokens[3]

		mandatoryLetter := tokens[2]

		min, _ := strconv.Atoi(tokens[0])
		max, _ := strconv.Atoi(tokens[1])

		letters := hlp.StringToSlice(password)

		letterCount := 0

		for _, s := range letters {

			if s == mandatoryLetter {
				letterCount++
			}
		}
		if letterCount >= min && letterCount <= max {
			validPasswords[password] = true
			validCount++
		}

		a := min - 1
		b := max - 1

		if hlp.XorStrings(mandatoryLetter, letters[a], letters[b]) {
			validCount2++
			validPasswords2[password] = true
		}
	}

	fmt.Println(validCount) //Answer is 418
	fmt.Println(len(validPasswords))
	fmt.Println("---")

	fmt.Println(validCount2) //Answer 616
	fmt.Println(len(validPasswords2))

}
