package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Disc struct {
	Name       string
	Weight     int
	Children   []*Disc
	TotWeight  int
	ParentName string
}

func main() {

	lines := fileToLines("input.txt")

	programs := make(map[string]*Disc)

	for _, line := range lines {
		if line == "" {
			continue
		}
		program := strings.Split(line, " ")[0]
		weightStr := strings.Split(line, " ")[1]
		weightStr = strings.Replace(weightStr, "(", "", 1)
		weightStr = strings.Replace(weightStr, ")", "", 1)
		weight, _ := strconv.Atoi(weightStr)

		programs[program] = &Disc{Name: program, Weight: weight, TotWeight: weight}

	}
	for _, line2 := range lines {
		if line2 == "" {
			continue
		}
		program := strings.Split(line2, " ")[0]
		subPrograms := getSubPrograms(line2)

		for _, subProgram := range subPrograms {
			if subProgram == "" {
				continue
			}
			p := programs[program]
			s := programs[subProgram]
			p.Children = append(p.Children, s)
			s.ParentName = program
			programs[program] = p
			programs[subProgram] = s
		}
	}

	for _, disc := range programs {
		if disc.ParentName == "" {
			calcTotWeight(disc)
			printTree(disc, "")
		}
	}

}
func calcTotWeight(disc *Disc) {
	for _, s := range disc.Children {
		calcTotWeight(s)
		disc.TotWeight += s.TotWeight
	}

}

func printTree(disc *Disc, level string) {
	fmt.Println(".", level, ".", disc.Name, disc.TotWeight, disc.Weight)
	for _, s := range disc.Children {
		printTree(s, level+" -")
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
