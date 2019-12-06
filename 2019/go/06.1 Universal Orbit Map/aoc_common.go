package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
func getSubStringMatrix(in [][]string, x1, y1, blockSize int) [][]string {
	a := make([][]string, blockSize)

	for y := y1; y < x1+blockSize; y++ {
		for x := x1; x < x1+blockSize; x++ {
			a[y] = append(a[y], in[y][x])
		}
	}
	return a
}
func flipSquareStringMatrixY(pattern [][]string) [][]string {
	ret := [][]string{}
	for readY := len(pattern) - 1; readY >= 0; readY-- {
		ret = append(ret, pattern[readY])
	}
	return ret

}

func flipSquareStringMatrixX(in [][]string) [][]string {
	ret := makeSquareStringMatrix(len(in[0]), "")
	maxX := len(in[0]) - 1
	for x := 0; x < len(in[0]); x++ {
		for y := 0; y < len(in); y++ {
			ret[y][maxX-x] = in[y][x]
		}
	}

	return ret
}

func rotateSquareStringMatrix90(in [][]string) [][]string {
	a := make([][]string, len(in[0]))
	for x := 0; x < len(in[0]); x++ {
		for y := len(in[0]) - 1; y >= 0; y-- {
			a[x] = append(a[x], in[y][x])
		}
	}
	return a
}
func makeSquareStringMatrix(size int, def string) [][]string {
	ret := [][]string{}
	for y := 0; y < size; y++ {
		ret = append(ret, []string{})
		for x := 0; x < size; x++ {
			ret[y] = append(ret[y], def)
		}
	}
	return ret
}
func blitStringMatrix(small, large [][]string, x, y int) [][]string {
	for dy := 0; dy < len(small[0]); dy++ {
		for dx := 0; dx < len(small[0]); dx++ {
			large[dy+y][dx+x] = small[dy][dx]
		}
	}

	return large
}
func stringToSlice(in string) []string {
	ret := []string{}
	for _, r := range in {
		ret = append(ret, string(r))
	}
	return ret
}

func fileToLines(fileName string) []string {
	ret := make([]string, 0, 100)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			ret = append(ret, line)
		}
	}

	return ret
}

func dumpStringMatrix(matrix [][]string, caption string) {
	fmt.Println(caption + ":")
	for _, line := range matrix {
		fmt.Print("|>")
		for _, s := range line {
			fmt.Print(s)
		}
		fmt.Println("<|")
	}
	fmt.Println("")
}

func StringToIntArray(in []string) []int {
	ret := []int{}
	for _, s := range in {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			i = 0
		}
		ret = append(ret, i)
	}

	return ret
}
func SumIntArray(in []int) int {
	ret := 0
	for _, i := range in {
		ret += i
	}
	return ret
}

func DivideByOnIntArray(in []int, div int) []int {
	ret := []int{}

	for _, i := range in {
		x := i / div
		ret = append(ret, x)
	}
	return ret
}
func SubtractOnIntArray(in []int, sub int) []int {
	ret := []int{}

	for _, i := range in {
		x := i - sub
		ret = append(ret, x)
	}
	return ret
}

func IntAbs(in int) int {
	if in < 0 {
		return in * -1
	}
	return in
}
func IntMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func IntMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func IntBetween(p, a, b int) bool {
	return p >= IntMin(a, b) && p <= IntMax(a, b)
}
func (a Point) Manhattan(b Point) int {
	return Manhattan(a.X, a.Y, b.X, b.Y)
}
func Manhattan(aX, aY, bX, bY int) int {
	x := aX - bX
	y := aY - bY
	return IntAbs(x) + IntAbs(y)
}
