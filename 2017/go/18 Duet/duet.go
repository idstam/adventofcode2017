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
var lastSound = 0
var tick = 0

func main() {
	lines := fileToLines("input.txt")
	registers[0] = make(map[string]int)
	registers[1] = make(map[string]int)

	registers[0]["p"] = 0
	registers[1]["p"] = 1

	for {
		tick++
		for pid := 0; pid < 2; pid++ {
			// if waiting[pid] {
			// 	registers[pid]["waiting"] = 1
			// } else {
			// 	registers[pid]["waiting"] = 0
			// }
			// registers[pid]["queueLength"] = len(queues[pid])
			// fmt.Println(registers[pid])
			// fmt.Println(lines[registers[pid]["pntr"]])

			if registers[pid]["done"] == 1 {
				continue
			}
			if waiting[pid] {
				continue
			}
			cmd, a, b := parseLine(pid, lines[registers[pid]["pntr"]])
			execCmd(pid, cmd, a, b)

			registers[pid]["pntr"] = registers[pid]["pntr"] + 1
			if registers[pid]["pntr"] <= 0 || registers[pid]["pntr"] >= len(lines) {
				registers[pid]["done"] = 1
			}
		}
		//fmt.Println("------------------------------------------")
		if (waiting[0] || registers[0]["done"] == 1) && (waiting[1] || registers[1]["done"] == 1) {
			fmt.Println("Both terminated")
			break
		}

		if tick%1000000 == 0 {
			fmt.Println("Tick count too large")
			//break
		}
	}

	fmt.Printf("Program 1 send count %d \n", count)
}

func execCmd(pid int, cmd string, a string, b int) {

	switch cmd {
	case "snd":
		snd(pid, a)
		lastSound = registers[0][a]
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

	val := registers[pid][reg]

	if pid == 0 {
		queues[1] = append(queues[1], val)
		waiting[1] = false
	} else {
		count++
		queues[0] = append(queues[0], val)
		waiting[0] = false
	}
}
func rcv(pid int, reg string) {
	if len(queues[pid]) == 0 {
		waiting[pid] = true
		registers[pid]["pntr"] = registers[pid]["pntr"] - 1 //So that we'll try this instruction again when the wait is over
		return
	}

	// if registers[pid][reg] != 0 {
	// 	if lastSound != 0 && pid == 0 {
	// 		fmt.Printf("Last sound for part one: %d \n", lastSound)
	// 	}
	// }
	registers[pid][reg] = queues[pid][0]
	queues[pid] = queues[pid][1:]

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
	a, err := strconv.Atoi(reg)
	if err != nil {
		a = registers[pid][reg]
	}
	if a > 0 {
		registers[pid]["pntr"] = registers[pid]["pntr"] + (val - 1) //When I get out of here the pointer will ++, hence -1
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
