package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	lines := fileToLines("input.txt")
	line := lines[1]
	opened := 0
	score := 0
	garbage := 0
	ignore := false
	inGarbage := false
	for _, c := range line {
		if ignore {
			ignore = false
			continue
		}
		if c == '!' {
			ignore = true
			continue
		}
		if c == '<' && !inGarbage {
			inGarbage = true
			continue
		}
		if inGarbage {
			if c == '>' {
				inGarbage = false
			} else {
				garbage++
			}
			continue
		}
		if c == '{' {
			opened++
			score += opened
		}
		if c == '}' {
			opened--
		}
	}

	fmt.Println(garbage)
}

func fileToLines(fileName string) []string {
	ret := make([]string, 1, 100)

	file, err := os.Open("input.txt")
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
