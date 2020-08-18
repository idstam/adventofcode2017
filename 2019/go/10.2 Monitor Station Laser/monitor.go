package main

import "sort"

import "fmt"

import "math"

type SlopePoint struct {
	Point     IntPoint
	Angle     float64
	Direction float64
	DirX      int
	DirY      int
	Dist      int
}

func main() {

	// pa := IntPoint{0, 1}
	// pb := IntPoint{0, 0}
	// a := Angle(pa, pb)

	// fmt.Println(a)

	lines := fileToLines("input.txt")
	m := [][]string{}
	for _, l := range lines {
		m = append(m, StringToSlice(l))
	}

	//Example {11 13} 210
	//Real {22 28} 326

	points := []SlopePoint{}
	laser := IntPoint{22, 28}

	direction := func(p1, p2 *SlopePoint) bool {
		if p1.Angle == p2.Angle {
			return p1.Dist < p2.Dist
		}
		return p1.Angle < p2.Angle
	}

	for y, l := range m {
		for x, c := range l {
			if c == "#" {
				sp := SlopePoint{}
				sp.Point = IntPoint{X: x, Y: y}
				sp.DirX = IntDirection(laser.X, x)
				sp.DirY = IntDirection(laser.Y, y)
				sp.Dist = laser.Manhattan(sp.Point)
				sp.Angle = Angle(laser, sp.Point)
				sp.Direction = Direction(sp.DirX, sp.DirY, sp.Angle)
				points = append(points, sp)
			}
		}
	}

	By(direction).Sort(points)
	//Burn them
	lastAngle := -100.0
	count := 0
	for _, p := range points {

		if p.Angle != lastAngle {
			count++
			fmt.Println(count, p)
			lastAngle = p.Angle
		}
	}

}

func IndexOfSlopePoint(points []SlopePoint, p IntPoint) int {
	for i, sp := range points {
		if sp.Point.SameAs(p) {
			return i
		}
	}
	return -1
}

func Angle(a, b IntPoint) float64 {

	dx := float64(b.X) - float64(a.X)
	dy := float64(b.Y) - float64(a.Y)

	theta_rad := math.Atan2(dy, dx)
	foo := 0.0
	theta_rad += math.Pi / 2
	if theta_rad < 0 {
		foo = 360
	}

	theta_deg := (theta_rad / math.Pi * 180) + foo

	//	degrees = (degrees + 360) % 360;  // +360 for implementations where mod returns negative numbers
	//if theta_deg = 360
	return theta_deg
}

func Direction(dx, dy int, slope float64) float64 {
	if dx == 0 && dy == -1 {
		return 1 + slope
	}
	if dx == 1 && dy == -1 {
		return 2 + slope
	}
	if dx == 1 && dy == 0 {
		return 3 + slope
	}
	if dx == 1 && dy == 1 {
		return 4 + slope
	}
	if dx == 0 && dy == 1 {
		return 5 + slope
	}
	if dx == -1 && dy == 1 {
		return 6 + slope
	}
	if dx == -1 && dy == 0 {
		return 7 + slope
	}
	if dx == -1 && dy == 1 {
		return 8 + slope
	}

	return 999 + slope
}

type planetSorter struct {
	planets []SlopePoint
	by      func(p1, p2 *SlopePoint) bool // Closure used in the Less method.
}

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *SlopePoint) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(planets []SlopePoint) {
	ps := &planetSorter{
		planets: planets,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// Len is part of sort.Interface.
func (s *planetSorter) Len() int {
	return len(s.planets)
}

// Swap is part of sort.Interface.
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}
