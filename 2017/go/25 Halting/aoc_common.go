package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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
		i, err := strconv.Atoi(s)
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
