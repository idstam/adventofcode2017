package main

import "strconv"

import "fmt"

import "strings"

func main() {

	count := 0
	for i := 245182; i <= 790572; i++ {
		//for i := 111122; i <= 111122; i++ {
		numStr := strconv.Itoa(i)
		numChars := stringToSlice(numStr)
		// if HasAdjacent(numChars) {
		// 	if NoDecrease(numChars) {
		// 		count++
		// 	}
		// }
		if HasPair(numChars, numStr) {
			if NoDecrease(numChars) {
				count++
			}
		}
	}

	fmt.Println(count)
}
func HasPair(numChars []string, numString string) bool {
	for _, s := range numChars {
		s2 := s + s
		s3 := s2 + s
		if strings.Contains(numString, s2) &&
			!strings.Contains(numString, s3) {
			return true
		}
	}
	return false
}
func HasAdjacent(numChars []string) bool {
	lastNum := ""
	for _, s := range numChars {
		if s == lastNum {
			return true
		}
		lastNum = s
	}
	return false
}
func NoDecrease(numChars []string) bool {
	lastNum := 0
	for _, s := range numChars {
		i, _ := strconv.Atoi(s)
		if i < lastNum {
			return false
		}
		lastNum = i
	}
	return true
}
