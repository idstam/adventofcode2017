package helpers

import (
	"fmt"
)

func GetSubStringMatrix(in [][]string, x1, y1, blockSize int) [][]string {
	a := make([][]string, blockSize)

	for y := y1; y < x1+blockSize; y++ {
		for x := x1; x < x1+blockSize; x++ {
			a[y] = append(a[y], in[y][x])
		}
	}
	return a
}
func FlipSquareStringMatrixY(pattern [][]string) [][]string {
	ret := [][]string{}
	for readY := len(pattern) - 1; readY >= 0; readY-- {
		ret = append(ret, pattern[readY])
	}
	return ret

}

func FlipSquareStringMatrixX(in [][]string) [][]string {
	ret := MakeSquareStringMatrix(len(in[0]), "")
	maxX := len(in[0]) - 1
	for x := 0; x < len(in[0]); x++ {
		for y := 0; y < len(in); y++ {
			ret[y][maxX-x] = in[y][x]
		}
	}

	return ret
}

func RotateSquareStringMatrix90(in [][]string) [][]string {
	a := make([][]string, len(in[0]))
	for x := 0; x < len(in[0]); x++ {
		for y := len(in[0]) - 1; y >= 0; y-- {
			a[x] = append(a[x], in[y][x])
		}
	}
	return a
}
func MakeSquareStringMatrix(size int, def string) [][]string {
	ret := [][]string{}
	for y := 0; y < size; y++ {
		ret = append(ret, []string{})
		for x := 0; x < size; x++ {
			ret[y] = append(ret[y], def)
		}
	}
	return ret
}
func BlitStringMatrix(small, large [][]string, x, y int) [][]string {
	for dy := 0; dy < len(small[0]); dy++ {
		for dx := 0; dx < len(small[0]); dx++ {
			large[dy+y][dx+x] = small[dy][dx]
		}
	}

	return large
}
func DumpStringMatrix(matrix [][]string, caption string) {
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
