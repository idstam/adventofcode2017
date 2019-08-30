package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	niceCount := 0
	for _, line := range fileToLines("input05.txt") {
		//line = "dvszwmarrgswjxmb"
		line += " "
		vowelCount := 0
		hasDouble := false

		if strings.Contains(line, "ab") {
			continue
		}
		if strings.Contains(line, "cd") {
			continue
		}
		if strings.Contains(line, "pq") {
			continue
		}
		if strings.Contains(line, "xy") {
			continue
		}

		for pos, c := range line {
			if strings.ContainsRune("aeiou", c) {
				vowelCount++
			}

			if pos < len(line)-1 && line[pos] == line[pos+1] {
				hasDouble = true
			}
			if hasDouble && vowelCount > 2 {
				niceCount++
				break
			}
		}
		//break
	}

	fmt.Println(niceCount)
}

func minInt(numbers ...int) int {
	m := numbers[0]
	for _, n := range numbers {
		if n < m {
			m = n
		}
	}

	return m
}

func fileToLines(fileName string) []string {
	ret := make([]string, 0, 100)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			ret = append(ret, line)
		}
	}

	return ret
}
