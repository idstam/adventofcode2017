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
	}

	go vm.Run()

	screen := map[string]Int64Point{}
	for vm.State != "Done" {
		time.Sleep(time.Second)
		t := Int64Point{}
		t.X = <-vm.Output
		t.Y = <-vm.Output
		t.Z = <-vm.Output

		screen[t.Name()] = t
	}
	count := 0
	for _, t := range screen {
		if t.Z == 2 {
			count++
		}
	}

	fmt.Println(count)
}
