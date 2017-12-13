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
	delay := 9
	for {
		delay++
		scanners := initScanners(lines)
		//printState(scanners, 0)
		for pico := 0; pico < delay; pico++ {
			for _, scanner := range scanners {
				scanner.Pos = scanner.Pos + scanner.Direction
				if scanner.Pos == scanner.Depth || scanner.Pos == -1 {
					scanner.Direction = scanner.Direction * -1
					scanner.Pos = scanner.Pos + scanner.Direction
					scanner.Pos = scanner.Pos + scanner.Direction
				}
			}
			//printState(scanners, pico)
		}

		//printState(scanners, delay)

		severity := 0
		for myPos := 0; myPos < 88; myPos++ {
			if scanners[myPos] != nil && scanners[myPos].Pos == 0 {
				severity = 1
				break
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
		if severity == 0 {
			break
		}

	}
	fmt.Println(delay)

}
func printState(scanners map[int]*Scanner, pos int) {
	fmt.Println(pos)
	for k, scanner := range scanners {

		fmt.Print(k)
		fmt.Print(":")
		for i := 0; i < scanner.Depth; i++ {
			if scanner.Pos == i {
				fmt.Print("[X]")
			} else {
				fmt.Print("[ ]")
			}
		}
		fmt.Println("")

	}
	fmt.Println("--------------------------------------------------------------")
}
func initScanners(lines []string) map[int]*Scanner {
	scanners := make(map[int]*Scanner)

	for _, line := range lines {
		tokens := strings.Split(line, ":")
		srange, _ := strconv.Atoi(strings.TrimSpace(tokens[0]))
		depth, _ := strconv.Atoi(strings.TrimSpace(tokens[1]))

		scanners[srange] = &Scanner{Depth: depth, Direction: 1}

	}

	return scanners

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
