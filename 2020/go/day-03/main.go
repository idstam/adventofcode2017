package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
)

func main() {

	lines := hlp.FileToLines("input.txt")

	treeCount := countTrees(lines, 3)
	fmt.Println("First treeCount", treeCount) //First = 220

	treeCount = countTrees(lines, 1)
	totTreeCount := treeCount

	treeCount = countTrees(lines, 3)
	totTreeCount *= treeCount

	treeCount = countTrees(lines, 5)
	totTreeCount *= treeCount

	treeCount = countTrees(lines, 7)
	totTreeCount *= treeCount

	treeCount = 0
	x := 0
	lookAtThisRow := true
	for _, l := range lines {
		if lookAtThisRow {
			mapLine := hlp.StringToSlice(l)
			if mapLine[x] == "#" {
				treeCount++
			}
			x++
			x = x % len(mapLine)
		}

		lookAtThisRow = !lookAtThisRow
	}
	totTreeCount *= treeCount
	fmt.Println("Tot treeCount", totTreeCount)

}

func countTrees(lines []string, step int) int {
	treeCount := 0
	x := 0
	for _, l := range lines {
		mapLine := hlp.StringToSlice(l)
		if mapLine[x] == "#" {
			treeCount++
		}
		x += step
		x = x % len(mapLine)
	}
	return treeCount
}
