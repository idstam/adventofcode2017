package main

import (
	"strings"
)

func main() {

	lines := fileToLines("example.txt")
	ruleMap := map[string]string{}

	for _, line := range lines {
		tokens := strings.Split(line, " => ")
		ruleMap[tokens[0]] = tokens[1]
	}

	pattern := [][]string{}
	pattern = append(pattern, stringToSlice(".#."))
	pattern = append(pattern, stringToSlice("..#"))
	pattern = append(pattern, stringToSlice("###"))

	for t := 0; t < 5; t++ {
		blockSize := 2
		if len(pattern[0])%3 == 0 {
			blockSize = 3
		}
		for y := 0; y < len(pattern[0])/blockSize; y++ {
			for x := 0; x < len(pattern[0])/blockSize; x++ {
				block := getBlock(y, x, blockSize, pattern)
				rule := findRule(ruleMap, block)
				dumpStringMatrix(rule)
			}
		}
	}
}

func getBlock(y, x, blockSize int, pattern [][]string) [][]string {
	ret := [][]string{}
	for y2 := 0; y2 < blockSize; y2++ {
		ret = append(ret, pattern[(y*blockSize)+y2][(x*blockSize):blockSize])
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
	p = flipSquareStringMatrix(pattern)

	p = rotateSquareStringMatrix90(p)
	for i := 0; i < 4; i++ {
		p = rotateSquareStringMatrix90(pattern)
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
