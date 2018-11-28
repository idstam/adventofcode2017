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

	programs := make(map[string]int)

	for _, line := range lines {
		if line == "" {
			continue
		}
		program := strings.Split(line, " ")[0]
		if programs[program] == 0 {
			programs[program] = 0
		}
		subPrograms := getSubPrograms(line)

		for _, subProgram := range subPrograms {
			programs[subProgram] = programs[subProgram] + 1
		}
	}
	for program, count := range programs {
		if count == 0 {
			fmt.Println(program)
			return
		}

	}

}

func getSubPrograms(line string) []string {
	ret := make([]string, 1, 1)

	pos := strings.Index(line, "->")
	if pos == -1 {
		return ret
	}
	pos += 3
	subStrings := line[pos:len(line)]

	ret = strings.Split(subStrings, ", ")
	return ret

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
