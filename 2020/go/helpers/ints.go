package helpers

import (
	"strconv"
	"strings"
)

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

//XorInts returns true if exactly one of a and b equals t
func XorInts(t, a, b int) bool {

	return (t == a || t == b) && (a != b)
}
