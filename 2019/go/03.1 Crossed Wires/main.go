package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//var world [][]string

func main() {

	lines := fileToLines("input.txt")
	linesA := InputLineToLines(lines[0])
	linesB := InputLineToLines(lines[1])
	//world = makeSquareStringMatrix(16000, "")
	//DrawLines(linesA, "A")
	//DrawLines(linesB, "B")
	d := FindClosestIntersection(linesA, linesB)
	fmt.Println(d)

	d = FindShortestIntersection(linesA, linesB)
	fmt.Println(d)

}

func FindShortestIntersection(linesA, linesB []Line) int {
	minDist := math.MaxInt32
	distA := 0
	for _, l1 := range linesA {
		distB := 0

		for _, l2 := range linesB {
			intersects, p := l1.IntersectionStraight(l2)

			if intersects && Manhattan(0, 0, p.X, p.Y) != 0 {
				dist := distA + distB
				dist += p.Manhattan(l1.A)
				dist += p.Manhattan(l2.A)
				if dist != 0 && dist < minDist {
					minDist = dist
					fmt.Println(dist, l1.Meta, l2.Meta)
				}
			}
			distB += Manhattan(l2.A.X, l2.A.Y, l2.B.X, l2.B.Y)
		}
		distA += Manhattan(l1.A.X, l1.A.Y, l1.B.X, l1.B.Y)
	}
	return minDist
}

func FindClosestIntersection(linesA, linesB []Line) int {
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

func InputLineToLines(line string) []Line {
	fmt.Println("LineToPoints")
	ret := []Line{}
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
		l := Line{a, Point{x, y}, turn}
		ret = append(ret, l)
		//fmt.Printf("%s %s %d\n", turn, strconv.Itoa(x)+":"+strconv.Itoa(y), IntAbs(x)+IntAbs(y))
	}
	return ret
}
