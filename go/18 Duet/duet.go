package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var registers = make(map[int]map[string]int)
var waiting = make(map[int]bool)
var queues = make(map[int][]int)
var count int = 0

func main() {
	lines := fileToLines("input.txt")
	registers[0] = make(map[string]int)
	registers[1] = make(map[string]int)

	registers[0]["p"] = 0
	registers[1]["p"] = 1
	for {
		for pid := 0; pid < 2; pid++ {

			if registers[pid]["done"] == 1 {
				continue
			}

			cmd, a, b := parseLine(pid, lines[registers[pid]["pntr"]])
			execCmd(pid, cmd, a, b)
			registers[pid]["pntr"] = registers[pid]["pntr"] + 1
			if registers[pid]["pntr"] < 0 || registers[pid]["pntr"] >= len(lines) {
				registers[pid]["done"] = 1
			}
		}
		if waiting[0] && waiting[1] {
			break
		}
		if registers[0]["done"] == 1 && registers[1]["done"] == 1 {
			break
		}
	}

	fmt.Println(count)
}

func execCmd(pid int, cmd string, a string, b int) {
	switch cmd {
	case "snd":
		snd(pid, a)
		break
	case "set":
		set(pid, a, b)
		break
	case "add":
		add(pid, a, b)
		break
	case "mul":
		mul(pid, a, b)
		break
	case "mod":
		mod(pid, a, b)
		break
	case "rcv":
		rcv(pid, a)
		break
	case "jgz":
		jgz(pid, a, b)
		break

	}
}
func parseLine(pid int, line string) (cmd string, a string, b int) {
	tokens := strings.Split(line, " ")

	x := 0
	if len(tokens) == 3 {
		ib, err := strconv.Atoi(tokens[2])
		if err != nil {
			ib = registers[pid][tokens[2]]
		}
		x = ib
	}
	return tokens[0], tokens[1], x
}

func snd(pid int, reg string) {
	if pid == 0 {
		queues[1] = append(queues[1], registers[0][reg])
	} else {
		count++
		queues[1] = append(queues[0], registers[1][reg])
	}
}
func rcv(pid int, reg string) {
	if registers[pid][reg] != 0 {
		if len(queues[pid]) > 0 {
			registers[pid][reg] = queues[pid][len(queues[pid])-1]
			queues[pid] = queues[pid][1:len(queues[pid])]
		} else {
			waiting[pid] = true
		}
	}
}
func set(pid int, reg string, val int) {
	registers[pid][reg] = val
}
func add(pid int, reg string, val int) {
	registers[pid][reg] = registers[pid][reg] + val
}
func mul(pid int, reg string, val int) {
	registers[pid][reg] = registers[pid][reg] * val
}
func mod(pid int, reg string, val int) {
	registers[pid][reg] = registers[pid][reg] % val
}

func jgz(pid int, reg string, val int) {
	if registers[pid][reg] > 0 {
		registers[pid]["pntr"] = registers[pid]["pntr"] + (val - 1)
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
