package helpers

import (
	"fmt"
)

type StringMatrix struct {
	Matrix [][]string
}

func (m *StringMatrix) InitSquare(size int, def string) {
	ret := [][]string{}
	for y := 0; y < size; y++ {
		ret = append(ret, []string{})
		for x := 0; x < size; x++ {
			ret[y] = append(ret[y], def)
		}
	}
	m.Matrix = ret
}

func (m *StringMatrix) InitFromLines(lines []string) {

	for y, line := range lines {
		l := StringToSlice(line)
		m.Matrix = append(m.Matrix, []string{})
		for _, p := range l {
			m.Matrix[y] = append(m.Matrix[y], p)
		}
	}

}

func (m *StringMatrix) Clone() StringMatrix {
	ret := StringMatrix{}
	ret.Matrix = [][]string{}
	for x := 0; x < m.Width(); x++ {
		ret.Matrix = append(ret.Matrix, []string{})

		for y := 0; y < m.Height(); y++ {
			ret.Matrix[x] = append(ret.Matrix[x], m.SafeGet(x, y, ""))
		}
	}
	return ret
}
func (m *StringMatrix) GetSubStringMatrix(x1, y1, blockSize int) [][]string {
	a := make([][]string, blockSize)

	for y := y1; y < x1+blockSize; y++ {
		for x := x1; x < x1+blockSize; x++ {
			a[y] = append(a[y], m.Matrix[y][x])
		}
	}
	return a
}
func (m *StringMatrix) FlipY() {
	ret := StringMatrix{}
	ret.Matrix = m.Matrix
	for readY := len(m.Matrix) - 1; readY >= 0; readY-- {
		ret.Matrix = append(ret.Matrix, m.Matrix[readY])
	}
	m.Matrix = ret.Matrix

}

func (m *StringMatrix) FlipX() {
	ret := StringMatrix{}
	ret.InitSquare(len(m.Matrix[0]), "")

	maxX := len(m.Matrix[0]) - 1
	for x := 0; x < len(m.Matrix[0]); x++ {
		for y := 0; y < len(m.Matrix); y++ {
			ret.Matrix[y][maxX-x] = m.Matrix[y][x]
		}
	}

	m.Matrix = ret.Matrix
}

func (m *StringMatrix) RotateSquareStringMatrix90() [][]string {
	a := make([][]string, len(m.Matrix[0]))
	for x := 0; x < len(m.Matrix[0]); x++ {
		for y := len(m.Matrix[0]) - 1; y >= 0; y-- {
			a[x] = append(a[x], m.Matrix[y][x])
		}
	}
	return a
}
func BlitStringMatrix(small, large [][]string, x, y int) [][]string {
	for dy := 0; dy < len(small[0]); dy++ {
		for dx := 0; dx < len(small[0]); dx++ {
			large[dy+y][dx+x] = small[dy][dx]
		}
	}

	return large
}

func (m *StringMatrix) Dump(caption string) {
	fmt.Println(caption + ":")
	for _, line := range m.Matrix {
		fmt.Print("|>")
		for _, s := range line {
			fmt.Print(s)
		}
		fmt.Println("<|")
	}
	fmt.Println("")
}

func (m *StringMatrix) Width() int {
	return len(m.Matrix)
}

func (m *StringMatrix) Height() int {

	if m.Width() > 0 {
		return len(m.Matrix[0])
	}
	return 0
}
func (m *StringMatrix) CountAdjacent(x, y int, needle string) int {
	ret := 0
	vals := []string{}
	vals = append(vals, m.SafeGet(x, y-1, ""))
	vals = append(vals, m.SafeGet(x+1, y-1, ""))
	vals = append(vals, m.SafeGet(x+1, y, ""))
	vals = append(vals, m.SafeGet(x+1, y+1, ""))
	vals = append(vals, m.SafeGet(x, y+1, ""))
	vals = append(vals, m.SafeGet(x-1, y+1, ""))
	vals = append(vals, m.SafeGet(x-1, y, ""))
	vals = append(vals, m.SafeGet(x-1, y-1, ""))

	for _, v := range vals {
		if v == needle {
			ret++
		}
	}
	return ret
}

func (m *StringMatrix) SafeGet(x, y int, def string) string {

	if x < 0 || x >= m.Width() {
		return def
	}
	if y < 0 || y >= m.Height() {
		return def
	}

	return m.Matrix[x][y]
}
func (m *StringMatrix) SafeSet(x, y int, val string) bool {

	if x < 0 || x >= m.Width() {
		return false
	}
	if y < 0 || y >= m.Height() {
		return false
	}

	m.Matrix[x][y] = val
	return true
}
