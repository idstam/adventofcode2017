package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strings"
)

func main() {

	lines := hlp.FileToLines("input.txt", false)

	uniqueAnswers := map[string]int{}
	groupAnswerCount := 0
	answerCount1 := 0
	peopleCount := 0
	for _, l := range lines {
		if l == "" {
			answerCount1 += len(uniqueAnswers)

			for _, count := range uniqueAnswers {
				if count == peopleCount {
					groupAnswerCount++
				}
			}
			uniqueAnswers = map[string]int{}
			peopleCount = 0
		} else {
			peopleCount++
			answers := strings.Split(l, "")
			for _, a := range answers {
				uniqueAnswers[a] = uniqueAnswers[a] + 1
			}
		}

	}

	answerCount1 += len(uniqueAnswers)

	fmt.Println("First", answerCount1)

	for _, count := range uniqueAnswers {
		if count == peopleCount {
			groupAnswerCount++
		}
	}
	fmt.Println("Second", groupAnswerCount)

}
