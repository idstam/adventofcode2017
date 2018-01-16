package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var programs = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}

//var programs = []string{"a", "b", "c", "d", "e"}

func main() {

	commands := make(chan string)
	go parseCommands(commands)
	receiveCommands(commands)

	for _, name := range programs {
		fmt.Print(name)
	}
}

func parseCommands(c chan string) {
	buf, _ := ioutil.ReadFile("input.txt")
	data := string(buf)
	commands := strings.Split(data, ",")
	for _, command := range commands {
		c <- command
	}
	close(c)

}

func receiveCommands(c chan string) {
	for {
		cmd, ok := <-c
		if !ok {
			break
		}
		fmt.Println(cmd)
		if strings.HasPrefix(cmd, "s") {
			spin(strings.TrimPrefix(cmd, "s"))
		}
		if strings.HasPrefix(cmd, "x") {
			exchange(strings.TrimPrefix(cmd, "x"))
		}
		if strings.HasPrefix(cmd, "p") {
			partner(strings.TrimPrefix(cmd, "p"))
		}

		fmt.Println(programs)

	}
}
func partner(arg string) {
	args := strings.Split(arg, "/")

	a := findProgram(args[0])
	b := findProgram(args[1])
	swap := programs[a]
	programs[a] = programs[b]
	programs[b] = swap
}
func findProgram(name string) int {
	for k, v := range programs {
		if v == name {
			return k
		}
	}
	return -1
}
func exchange(arg string) {
	args := strings.Split(arg, "/")
	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])
	swap := programs[a]
	programs[a] = programs[b]
	programs[b] = swap
}
func spin(arg string) {
	num, _ := strconv.Atoi(arg)
	b := programs[:len(programs)-num]
	a := programs[len(programs)-num:]

	programs = append(a, b...)

}
