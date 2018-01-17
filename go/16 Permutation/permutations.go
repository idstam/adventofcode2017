package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var Programs = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}

//var Programs = []string{"a", "b", "c", "d", "e"}

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	data := string(buf)
	input := strings.Split(data, ",")

	commands := parseCommands(input)

	for i := 0; i < 1000000000; i++ {
		for _, command := range commands {
			command.Execute()
		}
	}

	for _, name := range Programs {
		fmt.Print(name)
	}
}

func parseCommands(input []string) []Command {
	ret := make([]Command, 0)

	for _, cmd := range input {
		if strings.HasPrefix(cmd, "s") {
			command := &SpinCommand{}
			command.Parse(strings.TrimPrefix(cmd, "s"))
			ret = append(ret, command)
		}
		if strings.HasPrefix(cmd, "x") {
			command := &ExchangeCommand{}
			command.Parse(strings.TrimPrefix(cmd, "x"))
			ret = append(ret, command)
		}
		if strings.HasPrefix(cmd, "p") {
			command := &PartnerCommand{}
			command.Parse(strings.TrimPrefix(cmd, "p"))
			ret = append(ret, command)
		}
	}
	return ret
}
func findProgram(nameA, nameB string) (int, int) {
	retA := -1
	retB := -1

	for k, v := range Programs {
		if v == nameA {
			retA = k
		}
		if v == nameB {
			retB = k
		}
		if retA > -1 && retB > -1 {
			break
		}
	}
	return retA, retB
}

type PartnerCommand struct {
	nA string
	nB string
}

func (command *PartnerCommand) Parse(arg string) {
	args := strings.Split(arg, "/")
	command.nA = args[0]
	command.nB = args[1]

}
func (command *PartnerCommand) Execute() {
	a, b := findProgram(command.nA, command.nB)

	swap := Programs[a]
	Programs[a] = Programs[b]
	Programs[b] = swap
}

type ExchangeCommand struct {
	a int
	b int
}

func (command *ExchangeCommand) Parse(arg string) {
	args := strings.Split(arg, "/")
	command.a, _ = strconv.Atoi(args[0])
	command.b, _ = strconv.Atoi(args[1])
}
func (command *ExchangeCommand) Execute() {
	swap := Programs[command.a]
	Programs[command.a] = Programs[command.b]
	Programs[command.b] = swap
}

type SpinCommand struct {
	num int
}

func (command *SpinCommand) Parse(arg string) {
	command.num, _ = strconv.Atoi(arg)
}
func (command *SpinCommand) Execute() {
	b := Programs[:len(Programs)-command.num]
	a := Programs[len(Programs)-command.num:]

	Programs = append(a, b...)

}

type Command interface {
	Parse(arg string)
	Execute()
}
