package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type IntPoint struct {
	X int
	Y int
	Z int
}

type Int64Point struct {
	X int64
	Y int64
	Z int64
}

func (p Int64Point) Name() string {
	return strconv.FormatInt(p.X, 10) + "," + strconv.FormatInt(p.Y, 10) + "," + strconv.FormatInt(p.Z, 10)
}

type Line struct {
	A    IntPoint
	B    IntPoint
	Meta string
}

func (p IntPoint) Name() string {
	return strconv.Itoa(p.X) + ":" + strconv.Itoa(p.Y) + "," + strconv.Itoa(p.Z)
}

func (l Line) IsHorizontal() bool {
	return l.A.Y == l.B.Y
}

func (l Line) IsVertical() bool {
	return l.A.X == l.B.X
}
func (l1 Line) IntersectionStraight(l2 Line) (bool, IntPoint) {
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

	return result, IntPoint{X: x, Y: y}
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
func StringToSlice(in string) []string {
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

func StringArrayToInt64Map(in []string) map[int64]int64 {
	ret := map[int64]int64{}
	for i, s := range in {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			v = 0
		}
		ret[int64(i)] = v
	}

	return ret
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
func IntMax(in ...int) int {
	m := in[0]
	for _, i := range in {
		if i > m {
			m = i
		}
	}
	return m
}
func Int64Min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
func Int64Max(in ...int64) int64 {
	m := in[0]
	for _, i := range in {
		if i > m {
			m = i
		}
	}
	return m
}
func IntBetween(p, a, b int) bool {
	return p >= IntMin(a, b) && p <= IntMax(a, b)
}
func (a IntPoint) Manhattan(b IntPoint) int {
	return Manhattan(a.X, a.Y, b.X, b.Y)
}
func Manhattan(aX, aY, bX, bY int) int {
	x := aX - bX
	y := aY - bY
	return IntAbs(x) + IntAbs(y)
}

func IntPermutateArray(in []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(in); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(in, 0)

	return permuts
}
func IntRemoveFromSlice(i int, s []int) []int {
	a := append(s[:i], s[i+1:]...)

	return a
}

func IntPushArray(i int, a []int) []int {
	return append([]int{i}, a...)
}
func IntPopArray(a []int) (int, []int) {
	if len(a) == 0 {
		log.Fatal("Trying to POP empty array")
	}
	val := a[0]
	return val, IntRemoveFromSlice(0, a)
}

//IntPointBetween Is the point p between a and b
func IntPointBetween(p, a, b IntPoint) bool {
	dxa := float64(a.X) - float64(b.X)
	dya := float64(a.Y) - float64(b.Y)
	slopeB := dxa / dya

	dxc := float64(a.X) - float64(p.X)
	dyc := float64(a.Y) - float64(p.Y)
	slopeC := dxc / dyc

	if slopeB != slopeC {
		return false
	}
	if !FloatSameSign(dxa, dxc) {
		return false
	}
	if !FloatSameSign(dya, dyc) {
		return false
	}

	if math.Abs(dxa) > math.Abs(dxc) {
		return true
	}
	return false
}

func FloatSameSign(in ...float64) bool {
	x := 2

	for _, f := range in {
		if x == 2 {
			if f < 0 {
				x = -1
			}
			if f > 0 {
				x = 1
			}
			x = 0
		}

		if f < 0 && x == 1 {
			return false
		}
		if f > 0 && x == -1 {
			return false
		}
		if f == 0 && x != 0 {
			return false
		}

	}

	return true
}
func IntDirection(a, b int) int {
	if b > a {
		return 1
	}
	if b < a {
		return -1
	}
	return 0
}
func (p IntPoint) SameAs(p2 IntPoint) bool {
	return p.X == p2.X && p.Y == p2.Y && p.Z == p2.Z
}
