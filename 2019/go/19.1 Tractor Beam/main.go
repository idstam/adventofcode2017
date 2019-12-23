package main

import (
	"fmt"
	"os"
	"strings"
)

var count int
var found bool
var program map[int64]int64
var targetSize int

func main() {
	lines := FileToLines("input.txt")
	numbers := strings.Split(lines[0], ",")
	program = StringArrayToInt64Map(numbers)
	targetSize = 100
	for y := 0; y < 1000; y++ {
		foundCount := 0
		fmt.Println("Line", y)

		for x := 0; x < 1000; x++ {

			found := HasPull(x, y)

			if found {
				foundCount++
			} else {
				foundCount = 0
			}

			if foundCount == targetSize {
				found := CheckSquareFrom(x, y)
				if found {
					fmt.Println("Result", x, y)
					os.Exit(0)
				}
			}
		}
	}
	//158 too high

	fmt.Println(count)
}

func CheckSquareFrom(x1, y1 int) bool {

	for y := y1; y < y1+targetSize; y++ {
		foundCount := 0
		for x := x1; x < x1+targetSize; x++ {

			found := HasPull(x, y)

			if !found {
				return false
			}

		}
	}

	return true
}
func HasPull(x, y int) bool {
	vm := VM1202{
		InputMode:  "Channel",
		OutputMode: "Function",
		mem:        program,
		LogLevel:   99,
	}

	vm.Input = make(chan int64, 10000)
	vm.Output = make(chan int64)
	//vm.InputFunction = VmInput
	vm.OutputFunction = VmOutput

	vm.Input <- int64(x)
	vm.Input <- int64(y)

	vm.Run()
	return found
}
func VmOutput(vm *VM1202, val int64) {
	if val == 1 {
		count++
		found = true
	}
}
