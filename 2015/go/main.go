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
		//line = "ieodomkazucvgmuy"
		hasA := false
		hasB := false

		for pos := range line {

			if pos > 1 && pos < len(line)-3 {
				rest := line[pos:]
				last2 := line[pos-2 : pos]
				if strings.Contains(rest, last2) {
					hasA = true
				}
			}
			if pos < len(line)-2 && line[pos] == line[pos+2] {
				hasB = true
			}

			if hasA && hasB {
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
