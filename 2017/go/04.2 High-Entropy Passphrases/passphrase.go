package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := make(map[string]int)

		valid := true

		words := strings.Split(scanner.Text(), " ")
		for _, word := range words {
			trimmedWord := strings.TrimSpace(word)
			sortedWord := SortString(trimmedWord)
			if m[sortedWord] == 0 {
				m[sortedWord] = 1
				continue
			} else {
				valid = false
				break
			}

		}
		if valid {
			count++
		}

	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
