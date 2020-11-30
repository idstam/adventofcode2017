package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
)

func main() {

	lines := hlp.FileToLines("adsf")

	for _, l := range lines {
		fmt.Println(l)
	}
}
