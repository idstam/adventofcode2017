package main

import (
	"strings"
)

func main() {

	lines := fileToLines("input.txt")
	strs := strings.Split(lines[0], ",")

	vm := VM1202{
		Name: "Test",
		mem:  StringToInt64Array(strs),
	}
	vm.Input = make(chan int64, 2)
	vm.Output = make(chan int64)
	vm.OutputMode = "Console"
	vm.InputMode = "Channel"
	vm.LogLevel = 99
	vm.Input <- 1
	vm.Run()
	//vm.inputs <- phase

}
