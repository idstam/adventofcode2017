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

	lines := fileToLines("input.txt")

	registers := make(map[string]int)

	for _, line := range lines {
		if line == "" {
			continue
		}
		tokens := strings.Split(line, " ")

		register := tokens[0]
		command := tokens[1]
		value, _ := strconv.Atoi(tokens[2])
		leftOp := tokens[4]
		operator := tokens[5]
		rightOp, _ := strconv.Atoi(tokens[6])

		comparison := compare(registers[leftOp], operator, rightOp)

		if comparison {
			registers[register] = registers[register] + commandValue(command, value)
		}
	}
	fmt.Println(maxInMapValues(registers))
}

func maxInMapValues(m map[string]int) int {
	isSet := false
	max := 0

	for _, v := range m {
		if v > max || isSet == false {
			max = v
			isSet = true
		}
	}
	return max
}
func commandValue(command string, value int) int {
	if command == "inc" {
		return value
	} else {
		return value * -1
	}
}
func compare(register int, operator string, value int) bool {
	switch operator {
	case "==":
		return register == value
	case ">":
		return register > value
	case "<":
		return register < value
	case "!=":
		return register != value
	case ">=":
		return register >= value
	case "<=":
		return register <= value
	}

	return false
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
