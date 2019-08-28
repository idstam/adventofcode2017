package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := fileToLines("input_01.txt")

	floor := 0

	for _, c := range lines[0] {
		if c == '(' {
			floor++
		}
		if c == ')' {
			floor--
		}

	}

	fmt.Println(floor)
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
