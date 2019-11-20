package main

import (
	"fmt"
	"strings"
)

// 4 16 37 56 160
func main() {

	lines := fileToLines("input.txt")
	ruleMap := map[string]string{}

	for _, line := range lines {
		tokens := strings.Split(line, " => ")
		ruleMap[tokens[0]] = tokens[1]
	}

	pattern := [][]string{}
	blockSize := 0
	newBlockSize := 0
	numBlocks := 0
	pattern = append(pattern, stringToSlice(".#."))
	pattern = append(pattern, stringToSlice("..#"))
	pattern = append(pattern, stringToSlice("###"))
	result := makeSquareStringMatrix(4000, "*")

	for t := 1; t <= 18; t++ {
		fmt.Printf("Iteration %d \n", t)
		if len(pattern)%2 == 0 {
			blockSize = 2
		} else {
			blockSize = 3
		}
		numBlocks = (len(pattern) / blockSize)
		for y := 0; y < numBlocks; y++ {
			for x := 0; x < numBlocks; x++ {
				block := getBlock(y, x, blockSize, pattern)
				rule := findRule(ruleMap, block)
				newBlockSize = len(rule)
				result = blitStringMatrix(rule, result, x*newBlockSize, y*newBlockSize)
				//dumpStringMatrix(result, "")
			}
		}
		pattern = getBlock(0, 0, newBlockSize*numBlocks, result)
		countOn(pattern)
		fmt.Printf("Size %d\n", len(pattern))
	}
}

func countOn(in [][]string) {
	count := 0
	for _, line := range in {
		for _, cell := range line {
			if cell == "#" {
				count++
			}
		}
	}
	fmt.Printf("Count %d \n", count)
}
func getNextLen(curLen, newBlockSize int) int {
	for i := curLen + 1; i < curLen*2; i++ {
		if i%newBlockSize == 0 {
			return i
		}
	}

	panic("What")
}

func getBlock(y, x, blockSize int, pattern [][]string) [][]string {
	ret := [][]string{}
	for y2 := 0; y2 < blockSize; y2++ {
		ret = append(ret, []string{})
		for x2 := 0; x2 < blockSize; x2++ {
			ret[y2] = append(ret[y2], pattern[y2+(y*blockSize)][x2+(x*blockSize)])
		}
	}
	return ret
}
func findRule(ruleMap map[string]string, pattern [][]string) [][]string {
	ret := [][]string{}
	p := pattern
	for i := 0; i < 4; i++ {
		p = rotateSquareStringMatrix90(p)
		//dumpStringMatrix(p)
		k := flattenPattern(p)
		r, exists := ruleMap[k]
		if exists {
			lines := strings.Split(r, "/")
			for _, line := range lines {
				ret = append(ret, stringToSlice(line))
			}
			return ret
		}
	}
	p = flipSquareStringMatrixY(pattern)
	for i := 0; i < 4; i++ {
		p = rotateSquareStringMatrix90(p)
		k := flattenPattern(p)
		r, exists := ruleMap[k]
		if exists {
			lines := strings.Split(r, "/")
			for _, line := range lines {
				ret = append(ret, stringToSlice(line))
			}
			return ret
		}
	}

	p = flipSquareStringMatrixX(pattern)
	for i := 0; i < 4; i++ {
		p = rotateSquareStringMatrix90(p)
		k := flattenPattern(p)
		r, exists := ruleMap[k]
		if exists {
			lines := strings.Split(r, "/")
			for _, line := range lines {
				ret = append(ret, stringToSlice(line))
			}
			return ret
		}
	}
	return ret
}

func flattenPattern(pattern [][]string) string {
	ret := ""
	for _, line := range pattern {
		for _, s := range line {
			ret += s
		}
		ret += "/"
	}
	return strings.TrimSuffix(ret, "/")
}
