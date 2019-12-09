package main

import (
	"strings"
)

func main() {

	lines := fileToLines("input.txt")
	strs := strings.Split(lines[0], ",")


	vm := VM1202{
		Name: "Test",
		mem:  StringToIntArray(strs),
	}
	vm.Input = make(chan int, 2)
	vm.Output = make(chan int)
	vm.OutputMode = "Channel"
	vm.InputMode = "Channel"
	vm.LogLevel = 99
	//vm.inputs <- phase

}
