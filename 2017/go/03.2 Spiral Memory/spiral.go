package main

import (
	"fmt"
	"math"
)

func main() {

	memory := [20][20]int{}

	maxX := 0
	maxY := 0
	minX := 0
	minY := 0
	x := 0
	y := 0
	xdir := 1
	ydir := 0
	memory[10][10] = 1
	for square := 1; square < 400; square++ {
		x = x + xdir
		y = y + ydir

		cell := 0
		cell += memory[x+10+1][y+10+1]
		cell += memory[x+10+1][y+10]
		cell += memory[x+10+1][y+10-1]
		cell += memory[x+10][y+10+1]
		cell += memory[x+10][y+10-1]
		cell += memory[x+10-1][y+10+1]
		cell += memory[x+10-1][y+10]
		cell += memory[x+10-1][y+10-1]
		memory[x+10][y+10] = cell

		if cell > 347991 {
			fmt.Println(cell)
			break
		}

		if x > maxX {
			maxX = x
			xdir = 0
			ydir = 1
			continue
		}

		if x < minX {
			minX = x
			xdir = 0
			ydir = -1
			continue
		}

		if y < minY {
			minY = y
			xdir = 1
			ydir = 0
			continue
		}
		if y > maxY {
			maxY = y
			xdir = -1
			ydir = 0
			continue
		}
	}
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))

}
