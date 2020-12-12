package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strconv"
	"strings"
)

var ints []int
var memo map[int]int

func main() {

	lines := hlp.FileToLines("input.txt", true)

	f := 90
	x := 0
	y := 0
	for _, line := range lines {
		dx := 0
		dy := 0
		f, dx, dy = parseLine(line, f)
		x += dx
		y += dy
	}
	fmt.Println("Part 1 ", x, y, "Manhattan from start", hlp.IntAbs(x)+hlp.IntAbs(y))

	x = 0
	y = 0
	wpx := 10
	wpy := -1
	for _, line := range lines {
		dx := 0
		dy := 0
		dx, dy = parseLine2(line, wpx, wpy)
		if strings.HasPrefix(line, "F") {
			x += dx
			y += dy
		} else {
			wpx = dx
			wpy = dy
		}
	}
	fmt.Println("Part 2 ", x, y, "Manhattan from start", hlp.IntAbs(x)+hlp.IntAbs(y))

}
func parseLine2(line string, wpX, wpY int) (int, int) {
	c, d := tokens(line)

	switch c {
	case "N":
		return wpX, wpY - d
	case "S":
		return wpX, wpY + d
	case "E":
		return wpX + d, wpY
	case "W":
		return wpX - d, wpY
	case "L":
		return rotateWp(-d, wpX, wpY)
	case "R":
		return rotateWp(d, wpX, wpY)
	case "F":

		return wpX * d, wpY * d
	}

	return 0, 0
}

func rotateWp(degrees, x, y int) (int, int) {
	if degrees < 0 {
		degrees = 360 + degrees
	}

	p := hlp.Point{X: x, Y: y}

	for i := 0; i < degrees/90; i++ {
		p.Rotate90(hlp.Point{X: 0, Y: 0})
	}

	return p.X, p.Y
}
func parseLine(line string, f int) (int, int, int) {
	c, d := tokens(line)

	switch c {
	case "N":
		return f, 0, -d
	case "S":
		return f, 0, d
	case "E":
		return f, d, 0
	case "W":
		return f, -d, 0
	case "L":
		f2 := (f - d) % 360
		if f2 < 0 {
			f2 = 360 + f2
		}
		return f2, 0, 0
	case "R":
		return (f + d) % 360, 0, 0

	case "F":
		fx, fy := fxfy(f)
		return f, fx * d, fy * d
	}

	return 0, 0, 0
}

func fxfy(f int) (int, int) {
	switch f {
	case 0:
		return 0, -1
	case 90:
		return 1, 0
	case 180:
		return 0, 1
	case 270:
		return -1, 0
	}
	return 0, 0
}
func tokens(line string) (string, int) {
	a := hlp.SubString(line, 0, 1)
	bStr := strings.Replace(line, a, "", 1)
	b, _ := strconv.Atoi(bStr)

	return a, b
}
