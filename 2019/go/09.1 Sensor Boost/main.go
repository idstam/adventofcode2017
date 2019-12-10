package main

import (
	"strings"
)

func main() {

	lines := fileToLines("example.txt")
	strs := strings.Split(lines[0], ",")

	vm := VM1202{
		Name: "Test",
		mem:  StringToInt64Array(strs),
	}
	vm.Input = make(chan int64, 2)
	vm.Output = make(chan int64)
	vm.OutputMode = "Console"
	vm.InputMode = "Console"
	vm.LogLevel = 99

	vm.Run()
	//vm.inputs <- phase

}
