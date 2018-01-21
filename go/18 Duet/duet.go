package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var lastSound int
var registers = make(map[string]int)
var pointer = 0

func main() {
	lines := fileToLines("input.txt")

	for {
		cmd, a, b := parseLine(lines[pointer])
		execCmd(cmd, a, b)
		pointer++
		if pointer < 0 || pointer >= len(lines) {
			break
		}
	}

	fmt.Println(lastSound)
}

func execCmd(cmd string, a string, b int) {
	switch cmd {
	case "snd":
		snd(a)
		break
	case "set":
		set(a, b)
		break
	case "add":
		add(a, b)
		break
	case "mul":
		mul(a, b)
		break
	case "mod":
		mod(a, b)
		break
	case "rcv":
		rcv(a)
		break
	case "jgz":
		jgz(a, b)
		break

	}
}
func parseLine(line string) (cmd string, a string, b int) {
	tokens := strings.Split(line, " ")

	x := 0
	if len(tokens) == 3 {
		ib, err := strconv.Atoi(tokens[2])
		if err != nil {
			ib = registers[tokens[2]]
		}
		x = ib
	}
	return tokens[0], tokens[1], x
}

func snd(reg string) {
	lastSound = registers[reg]
}
func rcv(reg string) {
	if registers[reg] != 0 {
		registers[reg] = lastSound //Break here to get value for part one
	}
}
func set(reg string, val int) {
	registers[reg] = val
}
func add(reg string, val int) {
	registers[reg] = registers[reg] + val
}
func mul(reg string, val int) {
	registers[reg] = registers[reg] * val
}
func mod(reg string, val int) {
	registers[reg] = registers[reg] % val
}

func jgz(reg string, val int) {
	if registers[reg] > 0 {
		pointer += (val - 1)
	}
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
