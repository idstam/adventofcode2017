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
	b := a.Clone(false)

	madeChange := true
	iterations := 0
	for madeChange {
		iterations++
		madeChange = false
		for x := 0; x < a.Width(); x++ {
			for y := 0; y < a.Height(); y++ {
				p := a.Matrix[x][y]
				c := a.CountAdjacent(x, y, "#")
				if p == "L" && c == 0 {
					madeChange = true
					b.Matrix[x][y] = "#"
				}
				if p == "#" && c >= 4 {
					madeChange = true
					b.Matrix[x][y] = "L"
				}
			}
		}
		//b.Dump("")
		a = b.Clone(true)
	}
	fmt.Println("Part 1 Iterations", iterations)
	occupiedCount := a.Count("#")

	fmt.Println("Part 1 Occupied", occupiedCount)

	a = hlp.StringMatrix{}
	a.InitFromLines(lines)
	b = a.Clone(false)

	madeChange = true
	for madeChange {
		madeChange = false
		//a.Dump("")
		for x := 0; x < a.Width(); x++ {
			for y := 0; y < a.Height(); y++ {
				p := a.Matrix[x][y]
				c := 0
				for dx := -1; dx < 2; dx++ {
					for dy := -1; dy < 2; dy++ {
						ray := a.Look(x, y, dx, dy)
						for _, rp := range ray {
							if rp == "L" {
								break
							}
							if rp == "#" {
								c++
								break
							}
						}
					}
				}

				if p == "L" && c == 0 {
					madeChange = true
					b.Matrix[x][y] = "#"
				}
				if p == "#" && c >= 5 {
					madeChange = true
					b.Matrix[x][y] = "L"
				}
			}
		}
		a = b.Clone(true)
	}

	occupiedCount = a.Count("#")

	fmt.Println("Part 2 Occupied", occupiedCount)

}
