package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	lines := fileToLines("input.txt")
	numbers := strings.Split(lines[0], ",")

	vm := VM1202{
		InputMode:  "Channel",
		OutputMode: "Channel",
		mem:        StringArrayToInt64Map(numbers),
		LogLevel:   99,
	}

	vm.Input = make(chan int64, 2)
	vm.Output = make(chan int64)

	go vm.Run()
	vm.Input <- 0
	screen := map[string]Int64Point{}

	for vm.State != "Done" {
		t := Int64Point{}
		t.X = <-vm.Output
		t.Y = <-vm.Output
		t.Z = <-vm.Output

		screen[t.Name()] = t
		time.Sleep(time.Millisecond)
	}
	count := 0
	for _, t := range screen {
		if t.Z == 2 {
			count++
		}
	}

	fmt.Println(count)
}
