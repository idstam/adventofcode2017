package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	lines := fileToLines("input.txt")
	strs := strings.Split(lines[0], ",")

	vm := VM1202{
		Name: "Test",
		mem:  StringToInt64Map(strs),
	}
	vm.Input = make(chan int64, 2)
	vm.Output = make(chan int64)
	vm.OutputMode = "Channel"
	vm.InputMode = "Channel"
	vm.LogLevel = 99
	//vm.Input <- 1
	go vm.Run()

	x := 0
	y := 0
	dx := 0
	dy := -1

	visited := map[string]int64{}

	matrix := makeSquareStringMatrix(1000, " ")
	visited["0:0"] = 0

	for vm.State != "Done" {
		vm.Input <- visited[strconv.Itoa(x)+":"+strconv.Itoa(y)]

		color := <-vm.Output

		visited[strconv.Itoa(x)+":"+strconv.Itoa(y)] = color

		fmt.Println(len(visited))
		if color == 1 {
			matrix[y+500][x+500] = "#"
		}

		if vm.State == "Done" {
			break
		}
		turn := <-vm.Output
		if turn == 0 {
			dx, dy = TurnLeft(dx, dy)
		} else {
			dx, dy = TurnRight(dx, dy)
		}
		y += dy
		x += dx
	}

	fmt.Println(len(visited))

	dumpStringMatrix(matrix, "asf")

}

func TurnLeft(dx, dy int) (int, int) {
	if dx == 0 {
		return dy, 0
	}
	if dy == 0 {
		return 0, -dx
	}

	return 0, 0
}
func TurnRight(dx, dy int) (int, int) {
	if dx == 0 {
		return -dy, 0
	}
	if dy == 0 {
		return 0, dx
	}

	return 0, 0
}
