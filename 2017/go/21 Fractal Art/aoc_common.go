package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
func flipSquareStringMatrix(pattern [][]string) [][]string {
	ret := [][]string{}
	for readY := len(pattern) - 1; readY >= 0; readY-- {
		ret = append(ret, pattern[readY])
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

func dumpStringMatrix(matrix [][]string) {
	fmt.Println("StringMatrix:")
	for _, line := range matrix {
		for _, s := range line {
			fmt.Print(s)
		}
		fmt.Println("")
	}
	fmt.Println("")
}
