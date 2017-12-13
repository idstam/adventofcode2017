package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Scanner struct {
	Depth     int
	Pos       int
	Direction int
}

func main() {
	lines := fileToLines("input.txt")
	scanners := make(map[int]*Scanner)

	for _, line := range lines {
		tokens := strings.Split(line, ":")
		srange, _ := strconv.Atoi(strings.TrimSpace(tokens[0]))
		depth, _ := strconv.Atoi(strings.TrimSpace(tokens[1]))

		scanners[srange] = &Scanner{Depth: depth, Direction: 1}

	}

	severity := 0
	for myPos := 0; myPos < 89; myPos++ {
		if scanners[myPos] != nil && scanners[myPos].Pos == 0 {
			d := scanners[myPos].Depth
			severity += myPos * d
		}
		for _, scanner := range scanners {
			scanner.Pos = scanner.Pos + scanner.Direction
			if scanner.Pos == scanner.Depth || scanner.Pos == -1 {
				scanner.Direction = scanner.Direction * -1
				scanner.Pos = scanner.Pos + scanner.Direction
				scanner.Pos = scanner.Pos + scanner.Direction
			}
		}
	}
	fmt.Println(severity)

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
		ret = append(ret, line)
	}

	return ret
}
func stringsToInts(numbers []string) []int {
	ret := make([]int, len(numbers), len(numbers))
	for i, v := range numbers {
		ret[i], _ = strconv.Atoi(v)
	}
	return ret
}
