package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strconv"
	"strings"
)

func main() {

	lines := hlp.FileToLines("input.txt", false)

	var maxID int64
	taken := map[int64]bool{}

	for _, s := range lines {
		sid := seatID(s)
		if sid > maxID {
			maxID = sid
		}
		taken[sid] = true
	}
	fmt.Println("Max seat id", maxID)

	var i int64
	for i = 0; i < maxID; i++ {
		if taken[i] && taken[i+2] && !taken[i+1] {
			fmt.Println("Your seat is", i+1)
			break
		}
	}

}

func seatID(seat string) int64 {

	s := strings.Replace(seat, "B", "1", -1)
	s = strings.Replace(s, "F", "0", -1)
	s = strings.Replace(s, "L", "0", -1)
	s = strings.Replace(s, "R", "1", -1)

	ret, _ := strconv.ParseInt(s, 2, 64)
	return ret
}
