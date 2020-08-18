package main

import (
	"fmt"
	"strings"
)

var world [][]string
var line []string

var x, y, dx, dy, direction int
var probing bool
var steps int
var oxygenPoint IntPoint
var inputs []int

func main() {
	lines := FileToLines("input.txt")
	numbers := strings.Split(lines[0], ",")

	//For part 2
	numbers[0] = "2"

	vm := VM1202{
		InputMode:  "Function",
		OutputMode: "Function",
		mem:        StringArrayToInt64Map(numbers),
		LogLevel:   99,
	}

	vm.Input = make(chan int64)
	vm.Output = make(chan int64)
	vm.InputFunction = VmInput
	vm.OutputFunction = VmOutput2

	inputs = []int{
		65, 44, 66, 44, 66, 44, 67, 44, 67, 44, 65, 44, 65, 44, 66, 44, 66, 44, 67, 10,
		76, 44, 49, 50, 44, 82, 44, 52, 44, 82, 44, 52, 10,
		82, 44, 49, 50, 44, 82, 44, 52, 44, 76, 44, 49, 50, 10,
		82, 44, 49, 50, 44, 82, 44, 52, 44, 76, 44, 54, 44, 76, 44, 56, 44, 76, 44, 56, 10,
		121, 10,
	}

	//vm.Input <- 0
	world = [][]string{}
	line = []string{}

	vm.Run()

	dumpStringMatrix(world, "")
}

func VmInput(vm *VM1202) int64 {
	i := 0
	i, inputs = IntPopArray(inputs)
	return int64(i)
}

func VmOutput2(vm *VM1202, val int64) {
	//fmt.Println(val)
	if val < 255 {
		fmt.Print(string(rune(val)))
	} else {
		fmt.Println(val)
	}
	//os.Exit(0)
}

func VmOutput(vm *VM1202, val int64) {
	c := string(rune(val))

	if val == 10 {
		world = append(world, line)
		line = []string{}
		return
	}

	line = append(line, c)

}
