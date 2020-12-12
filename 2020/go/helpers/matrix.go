package helpers

import (
	"fmt"
)

type StringMatrix struct {
	Matrix [][]string
}

func (m *StringMatrix) InitSquare(size int, def string) {
	m.InitEmpty(size, size, def)
}

func (m *StringMatrix) InitEmpty(height, width int, def string) {
	ret := [][]string{}
	for y := 0; y < width; y++ {
		ret = append(ret, []string{})
		for x := 0; x < height; x++ {
			ret[y] = append(ret[y], def)
		}
	}
	m.Matrix = ret
}

func (m *StringMatrix) InitFromLines(lines []string) {
	width := 0
	for _, line := range lines {
		width = IntMax(width, len(line))
	}

	m.InitEmpty(len(lines), width, "")

	for y, line := range lines {
		l := StringToSlice(line)
		for x, p := range l {
			m.SafeSet(x, y, p)
		}
	}

}

func (m *StringMatrix) Clone(reuseMem bool) StringMatrix {
	ret := StringMatrix{}
	ret.InitEmpty(m.Height(), m.Width(), "")

	for y := 0; y < m.Height(); y++ {
		for x := 0; x < m.Width(); x++ {
			ret.Matrix[x][y] = m.Matrix[x][y]
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

	for y := 0; y < m.Height(); y++ {
		fmt.Print("|>")
		for x := 0; x < m.Width(); x++ {
			fmt.Print(m.SafeGet(x, y, " "))
		}
		fmt.Println("<|")
	}
	fmt.Println("")

}

func (m *StringMatrix) Width() int {
	return len(m.Matrix)
}

func (m *StringMatrix) Height() int {

	return len(m.Matrix[0])
}

//Count returns how many occurances of neddle ther is in the matrix
func (m *StringMatrix) Count(needle string) int {
	occupiedCount := 0
	for x := 0; x < m.Width(); x++ {
		for y := 0; y < m.Height(); y++ {
			p := m.SafeGet(x, y, needle+"XX")
			if p == needle {
				occupiedCount++
			}
		}
	}

	return occupiedCount
}
func (m *StringMatrix) CountAdjacent(x, y int, needle string) int {
	ret := 0

	if m.SafeGet(x, y-1, "") == needle {
		ret++
	}
	if m.SafeGet(x+1, y-1, "") == needle {
		ret++
	}

	if m.SafeGet(x+1, y, "") == needle {
		ret++
	}
	if m.SafeGet(x+1, y+1, "") == needle {
		ret++
	}
	if m.SafeGet(x, y+1, "") == needle {
		ret++
	}
	if m.SafeGet(x-1, y+1, "") == needle {
		ret++
	}
	if m.SafeGet(x-1, y, "") == needle {
		ret++
	}

	if m.SafeGet(x-1, y-1, "") == needle {
		ret++
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

func (m *StringMatrix) Look(x, y, dx, dy int) []string {
	ret := []string{}
	if dx == 0 && dy == 0 {
		return ret
	}

	p := m.SafeGet(x+dx, y+dy, "")
	if p != "" {
		ret = append(ret, p)
		ret = append(ret, m.Look(x+dx, y+dy, dx, dy)...)
	}

	return ret

}

func (p *Point) Rotate90(origin Point) {

	tx := p.X - origin.X
	ty := p.Y - origin.Y

	tmpX := -ty
	tmpY := tx

	p.X = tmpX + origin.X
	p.Y = tmpY + origin.Y

}
