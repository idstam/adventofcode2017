package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
)

var ints []int
var memo map[int]int

func main() {

	lines := hlp.FileToLines("input.txt", true)

	a := hlp.StringMatrix{}

	a.InitFromLines(lines)
	b := a.Clone()

	madeChange := true
	iterations := 0
	for madeChange {
		iterations++
		madeChange = false
		for x := 0; x < a.Width(); x++ {
			for y := 0; y < a.Height(); y++ {
				p := a.SafeGet(x, y, "")
				c := a.CountAdjacent(x, y, "#")
				if p == "L" && c == 0 {
					madeChange = true
					b.SafeSet(x, y, "#")
				}
				if p == "#" && c >= 4 {
					madeChange = true
					b.SafeSet(x, y, "L")
				}
			}
		}
		//b.Dump("")
		a = b.Clone()
	}
	occupiedCount := 0
	fmt.Println("Part 1 Iterations", iterations)
	for x := 0; x < a.Width(); x++ {
		for y := 0; y < a.Height(); y++ {
			p := a.SafeGet(x, y, "")
			if p == "#" {
				occupiedCount++
			}
		}
	}
	fmt.Println("Part 1 Occupied", occupiedCount)

}
