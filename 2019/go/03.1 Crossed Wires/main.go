package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//var world [][]string

func main() {

	l1 := Line{Point{3, 3}, Point{3, 10}}
	l2 := Line{Point{1, 5}, Point{5, 5}}

	fmt.Println(l1.IntersectionStraight(l2))

	lines := fileToLines("input.txt")
	linesA := InputLineToLines(lines[0])
	linesB := InputLineToLines(lines[1])
	//world = makeSquareStringMatrix(16000, "")
	//DrawLines(linesA, "A")
	//DrawLines(linesB, "B")
	d := FindClosestIntersection(linesA, linesB)

	fmt.Println(d)

}
func FindClosestIntersection(linesA, linesB map[string]Line) int {
	minDist := math.MaxInt32
	for _, l1 := range linesA {
		for _, l2 := range linesB {
			intersects, p := l1.IntersectionStraight(l2)
			if intersects {
				dist := Manhattan(0, 0, p.X, p.Y)
				if dist != 0 && dist < minDist {
					minDist = dist
				}
			}
		}
	}
	return minDist
}

func InputLineToLines(line string) map[string]Line {
	fmt.Println("LineToPoints")
	ret := map[string]Line{}
	turns := strings.Split(line, ",")

	x := 0
	y := 0
	for _, turn := range turns {
		a := Point{x, y}
		dir := string(turn[0])
		distStr := strings.TrimPrefix(turn, dir)
		dist, _ := strconv.Atoi(distStr)
		switch dir {
		case "U":
			y += dist
		case "D":
			y -= dist
		case "L":
			x -= dist
		case "R":
			x += dist
		}
		l := Line{a, Point{x, y}}
		ret[strconv.Itoa(x)+":"+strconv.Itoa(y)] = l
		//fmt.Printf("%s %s %d\n", turn, strconv.Itoa(x)+":"+strconv.Itoa(y), IntAbs(x)+IntAbs(y))
	}
	return ret
}
