package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strconv"
)

func main() {

	lines := hlp.FileToLines("input.txt", true)

	mapA, maxX, maxY := initMap(lines)
	mapB := map[string]bool{}

	aIsCurrent := true

	minX := 0
	minY := 0
	minZ := 0
	maxZ := 0

	var current map[string]bool
	var next map[string]bool

	for i := 0; i < 6; i++ {
		if aIsCurrent {
			current = mapA
			mapB = make(map[string]bool)
			next = mapB

		} else {
			current = mapB
			mapA = make(map[string]bool)
			next = mapA
		}
		aIsCurrent = !aIsCurrent

		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for z := minZ - 1; z <= maxZ+1; z++ {
					coord := coordIndex(x, y, z, 0)
					count := countAdjacent(x, y, z, current)
					if count == 3 {
						next[coord] = true
						maxX, maxY, maxZ, _ = maxDims(x, y, z, 0, maxX, maxY, maxZ, 0)
						minX, minY, minZ, _ = minDims(x, y, z, 0, minX, minY, minZ, 0)
						continue
					}
					if current[coord] && count == 2 {
						next[coord] = true
						maxX, maxY, maxZ, _ = maxDims(x, y, z, 0, maxX, maxY, maxZ, 0)
						minX, minY, minZ, _ = minDims(x, y, z, 0, minX, minY, minZ, 0)
						continue
					}
				}
			}
		}
		//fmt.Println("Iteration", i, "Number of cubes", len(next))
	}
	fmt.Println("Number of cubes part1", len(next))

	mapA, maxX, maxY = initMap(lines)

	aIsCurrent = true

	minX = 0
	minY = 0
	minZ = 0
	maxZ = 0
	maxW := 0
	minW := 0
	for i := 0; i < 6; i++ {
		if aIsCurrent {
			current = mapA
			mapB = make(map[string]bool)
			next = mapB

		} else {
			current = mapB
			mapA = make(map[string]bool)
			next = mapA
		}
		aIsCurrent = !aIsCurrent

		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for z := minZ - 1; z <= maxZ+1; z++ {
					for w := minW - 1; w <= maxW+1; w++ {
						coord := coordIndex(x, y, z, w)
						count := countAdjacent2(x, y, z, w, current)
						if count == 3 {
							next[coord] = true
							maxX, maxY, maxZ, maxW = maxDims(x, y, z, w, maxX, maxY, maxZ, maxW)
							minX, minY, minZ, minW = minDims(x, y, z, w, minX, minY, minZ, minW)
							continue
						}
						if current[coord] && count == 2 {
							next[coord] = true
							maxX, maxY, maxZ, maxW = maxDims(x, y, z, w, maxX, maxY, maxZ, maxW)
							minX, minY, minZ, minW = minDims(x, y, z, w, minX, minY, minZ, minW)
							continue
						}
					}
				}
			}
		}
		//fmt.Println("Iteration", i, "Number of cubes", len(next))
	}
	fmt.Println("Number of cubes part2", len(next))

}

func maxDims(x, y, z, w, mx, my, mz, mw int) (int, int, int, int) {
	return hlp.IntMax(x, mx), hlp.IntMax(y, my), hlp.IntMax(z, mz), hlp.IntMax(w, mw)
}
func minDims(x, y, z, w, mx, my, mz, mw int) (int, int, int, int) {
	return hlp.IntMin(x, mx), hlp.IntMin(y, my), hlp.IntMin(z, mz), hlp.IntMin(w, mw)
}

func countAdjacent(x, y, z int, current map[string]bool) int {
	count := 0
	for ox := -1; ox <= 1; ox++ {
		for oy := -1; oy <= 1; oy++ {
			for oz := -1; oz <= 1; oz++ {
				if ox == 0 && oy == 0 && oz == 0 {
					continue
				}
				coord := coordIndex(x+ox, y+oy, z+oz, 0)
				if current[coord] {
					count++
				}
			}
		}
	}
	return count
}

func countAdjacent2(x, y, z, w int, current map[string]bool) int {
	count := 0
	for ox := -1; ox <= 1; ox++ {
		for oy := -1; oy <= 1; oy++ {
			for oz := -1; oz <= 1; oz++ {
				for ow := -1; ow <= 1; ow++ {
					if ox == 0 && oy == 0 && oz == 0 && ow == 0 {
						continue
					}
					coord := coordIndex(x+ox, y+oy, z+oz, w+ow)
					if current[coord] {
						count++
					}
				}
			}
		}
	}
	return count
}

func coordIndex(x, y, z, w int) string {
	return strconv.Itoa(x) + ":" + strconv.Itoa(y) + ":" + strconv.Itoa(z) + ":" + strconv.Itoa(w)
}
func initMap(lines []string) (map[string]bool, int, int) {
	ret := map[string]bool{}
	maxX, maxY := 0, 0
	for y, l := range lines {
		line := hlp.StringToSlice(l)
		for x, p := range line {
			if p == "#" {
				ret[coordIndex(x, y, 0, 0)] = true
			}
			maxX = hlp.IntMax(x, maxX)
		}
		maxY = hlp.IntMax(y, maxY)
	}

	return ret, maxX, maxY
}
