package main

import (
	"fmt"
	"math"
)

func main() {

	maxX := 0
	maxY := 0
	minX := 0
	minY := 0
	x := 0
	y := 0
	xdir := 1
	ydir := 0
	for square := 1; square < 347991; square++ {
		x = x + xdir
		y = y + ydir

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
