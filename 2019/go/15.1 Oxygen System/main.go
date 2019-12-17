package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
)

var world [][]string
var x, y, dx, dy, direction int
var backtracking bool

func main() {
	lines := FileToLines("input.txt")
	numbers := strings.Split(lines[0], ",")

	vm := VM1202{
		InputMode:  "Function",
		OutputMode: "Function",
		mem:        StringArrayToInt64Map(numbers),
		LogLevel:   99,
	}

	vm.Input = make(chan int64)
	vm.Output = make(chan int64)
	vm.InputFunction = VmInput
	vm.OutputFunction = VmOutput

	//vm.Input <- 0
	world = makeSquareStringMatrix(100, " ")
	x = 50
	y = 50
	dy = -1
	direction = 1
	world[y][x] = "."
	vm.Run()

}
func VmInput(vm *VM1202) int64 {
	return int64(direction)
}
func VmOutput(vm *VM1202, val int64) {
	switch val {
	case 0:
		world[y+dy][x+dx] = "#"
		ChangeDirection()
	case 1:
		world[y+dy][x+dx] = "."
		y += dy
		x += dx
	case 2:
		world[y+dy][x+dx] = "O"
		y += dy
		x += dx
		world[y][x] = "D"
		dumpStringMatrix(world, "Found OXYGEN")
		world[y][x] = "."
	}
}

func ChangeDirection() {

	for i := 0; i < 4; i++ {
		direction++
		if direction == 5 {
			direction = 1
		}
		SetDeltaDirections()
		if world[y+dy][x+dx] == " " {
			backtracking = false
			return
		}
	}
	if !backtracking {
		world[y][x] = "D"
		world[50][50] = "S"
		dumpStringMatrix(world, "Starting backtrack")
		world[y][x] = "."
		world[50][50] = "."
	}
	for i := 0; i < 4; i++ {
		direction++
		if direction == 5 {
			direction = 1
		}
		SetDeltaDirections()
		if world[y+dy][x+dx] == "." {
			backtracking = true
			return
		}
	}

	world[y][x] = "D"
	world[50][50] = "S"
	dumpStringMatrix(world, "Should not happen")
	world[y][x] = "."
	world[50][50] = "."

}
func SetDeltaDirections() {
	switch direction {
	case 1:
		dx = 0
		dy = -1
	case 2:
		dx = 0
		dy = 1
	case 3:
		dx = -1
		dy = 0
	case 4:
		dx = 1
		dy = 0
	}
}
func DrawText(x, y int, text string, style tcell.Style, s tcell.Screen) {
	for i, r := range text {
		s.SetContent(x+i, y, r, []rune(""), style)
	}
}
func DrawPoint(p Int64Point, sc tcell.Screen) {
	s := tcell.StyleDefault.Foreground(tcell.ColorBlack)
	switch p.Z {
	case 0:
		s = s.Foreground(tcell.ColorBlack).Background(tcell.ColorDarkGray)
	case 1:
		s = s.Foreground(tcell.ColorDarkBlue).Background(tcell.ColorBlue)
	case 2:
		s = s.Foreground(tcell.ColorDarkRed).Background(tcell.ColorRed)
	case 3:
		s = s.Foreground(tcell.ColorOrange).Background(tcell.ColorYellow)
	case 4:
		s = s.Foreground(tcell.ColorGray).Background(tcell.ColorWhite)
	default:
		log.Fatal("Invalid tile type ", p.Z)
	}

	r := []rune(strconv.FormatInt(p.Z, 10))
	sc.SetContent(int(p.X), int(p.Y), r[0], []rune(""), s)

}
