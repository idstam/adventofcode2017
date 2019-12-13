package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
)

var paddleX int64
var ballX int64

func main() {
	lines := fileToLines("input.txt")
	numbers := strings.Split(lines[0], ",")

	vm := VM1202{
		InputMode:  "Function",
		OutputMode: "Channel",
		mem:        StringArrayToInt64Map(numbers),
		LogLevel:   2,
	}

	vm.Input = make(chan int64)
	vm.Output = make(chan int64)
	vm.InputFunction = NextValue
	go vm.Run()
	//vm.Input <- 0
	world := map[string]Int64Point{}

	score := int64(0)
	paddleX = int64(-1)
	var screen tcell.Screen
	// if vm.LogLevel == 99 {
	//screen = initScreen()
	// } else {
	screen = tcell.NewSimulationScreen("UTF8")
	// }

	vm.mem[0] = 2
	count := 0
	for vm.State != "Done" {
		t := Int64Point{}
		select {
		case t.X = <-vm.Output:

			t.Y = <-vm.Output
			t.Z = <-vm.Output

			if t.X == -1 && t.Y == 0 {
				score = t.Z
			} else {
				vm.Log(1, "Draw", t)

				if t.Z == 3 {
					paddleX = t.X
					vm.Log(2, "Paddle", paddleX, t.Y)
				}
				if t.Z == 4 {
					ballX = t.X
					vm.Log(2, "Ball", t.X, t.Y)
					//time.Sleep(500 * time.Millisecond)
					count++

				}
				DrawPoint(t, screen)

			}
		default:
		}
		_, y := screen.Size()
		world[t.Name()] = t

		style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
		DrawText(3, y+1, strconv.FormatInt(score, 10), style, screen)
		screen.Show()

		//time.Sleep(time.Millisecond)
	}
	//screen.Fini()
	fmt.Println()
	vm.DumpLogLines()

	fmt.Println("Score", score)
	fmt.Println("Count", count)
	vm.DumpInfo()

}
func NextValue(vm *VM1202) int64 {

	ret := int64(0)
	if paddleX != -1 {
		if ballX < paddleX {
			ret = -1
		}
		if ballX > paddleX {
			ret = 1
		}
		if ballX == paddleX {
			ret = 0
		}
	}
	vm.Log(2, "NextValue", ret)
	fmt.Println("-------------NextValue", ret)
	return ret
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
