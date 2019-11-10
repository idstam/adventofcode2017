package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	pipeMap, x := initPipeMap("input.txt")
	result := ""
	dX := 0
	dY := 1
	y := -1
	stepCount := 0
	for {
		stepCount++
		x += dX
		y += dY

		if x < 0 || y < 0 || x >= 250 || y >= 250 {
			fmt.Println("Out of bounds")
			fmt.Println(result)
			break
		}
		posChar := pipeMap[x][y]
		fmt.Printf("%s, %d,%d \n", posChar, x, y)

		if strings.TrimSpace(posChar) == "" {
			fmt.Println("Done")
			fmt.Println(result)
			break
		}
		if !(strings.Contains("|-+", posChar)) {
			result += posChar
		}

		if posChar == "+" && dX != 0 {
			dX = 0
			if strings.TrimSpace(pipeMap[x][y-1]) != "" {
				dY = -1
			} else {
				dY = 1
			}
			continue
		}

		if posChar == "+" && dY != 0 {
			dY = 0
			if strings.TrimSpace(pipeMap[x-1][y]) != "" {
				dX = -1
			} else {
				dX = 1
			}
		}

	}

	fmt.Printf("StepCount %d", stepCount-1)
}

func initPipeMap(fileName string) ([250][250]string, int) {
	lines := fileToLines(fileName)
	ret := [250][250]string{}
	startPos := -1
	for y, line := range lines {
		for x, r := range line {
			if y == 0 && startPos == -1 && string(r) == "|" {
				startPos = x
			}
			ret[x][y] = string(r)
		}
	}

	return ret, startPos
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
		line := scanner.Text()
		if line != "" {
			ret = append(ret, line)
		}
	}

	return ret
}
