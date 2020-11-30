package helpers

import (
	"strconv"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	A    Point
	B    Point
	Meta string
}

func (p Point) Name() string {
	return strconv.Itoa(p.X) + ":" + strconv.Itoa(p.Y)
}

func (l Line) IsHorizontal() bool {
	return l.A.Y == l.B.Y
}

func (l Line) IsVertical() bool {
	return l.A.X == l.B.X
}
func (l1 Line) IntersectionStraight(l2 Line) (bool, Point) {
	result := false
	x := 0
	y := 0
	if l1.IsVertical() && l2.IsHorizontal() {
		if IntBetween(l1.A.X, l2.A.X, l2.B.X) &&
			IntBetween(l2.A.Y, l1.A.Y, l1.B.Y) {
			x = l1.A.X
			y = l2.A.Y
			result = true
		}

	}

	if l2.IsVertical() && l1.IsHorizontal() {
		if IntBetween(l2.A.X, l1.A.X, l1.B.X) &&
			IntBetween(l1.A.Y, l2.A.Y, l2.B.Y) {
			x = l2.A.X
			y = l1.A.Y
			result = true
		}

	}

	return result, Point{X: x, Y: y}
}
func (a Point) Manhattan(b Point) int {
	return Manhattan(a.X, a.Y, b.X, b.Y)
}
func Manhattan(aX, aY, bX, bY int) int {
	x := aX - bX
	y := aY - bY
	return IntAbs(x) + IntAbs(y)
}
