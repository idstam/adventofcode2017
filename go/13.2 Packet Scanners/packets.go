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
	delay := 0
	scanners, maxScanner := initScanners(lines)

	for {
		step := 0
		delay++
		found := false

		for i := 0; i <= maxScanner; i++ {

			if scanners[i] == nil || (delay+step)%((scanners[i].Depth*2)-1) != 0 {
				step++
				continue
			} else {
				found = true
				break
			}

		}

		if !found {
			break
		}

	}
	fmt.Println(delay)

}
func initScanners(lines []string) (map[int]*Scanner, int) {
	scanners := make(map[int]*Scanner)
	max := 0
	for _, line := range lines {
		tokens := strings.Split(line, ":")
		srange, _ := strconv.Atoi(strings.TrimSpace(tokens[0]))
		depth, _ := strconv.Atoi(strings.TrimSpace(tokens[1]))

		scanners[srange] = &Scanner{Depth: depth, Direction: 1}
		if srange > max {
			max = srange
		}
	}

	return scanners, max

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
