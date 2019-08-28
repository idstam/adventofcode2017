package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := fileToLines("input02.txt")
	tot := 0
	for _, line := range lines {
		edges := strings.Split(line, "x")
		l, _ := strconv.Atoi(edges[0])
		w, _ := strconv.Atoi(edges[1])
		h, _ := strconv.Atoi(edges[2])

		a1 := 2 * l * w
		a2 := 2 * w * h
		a3 := 2 * h * l
		m := minInt(a1, a2, a3) / 2

		tot = tot + a1 + a2 + a3 + m

	}
	fmt.Println(tot)
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
