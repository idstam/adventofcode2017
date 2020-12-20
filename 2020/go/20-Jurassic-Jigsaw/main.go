package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strings"
)

type Square struct {
	Name   string
	Number int
	Sides  []string
	Lines  []string
	Matrix hlp.StringMatrix
}

func main() {

	lines := hlp.FileToLines("input.txt", true)
	squares := map[string]Square{}

	currentSquare := Square{Lines: []string{}, Sides: []string{}}

	for _, line := range lines {
		if strings.HasPrefix(line, "Tile") {
			if currentSquare.Name == "" {
				currentSquare.Name = line
			} else {
				squares[currentSquare.Name] = calcSquareSides(currentSquare)
				currentSquare = Square{Name: line, Lines: []string{}, Sides: []string{}}

			}
			continue
		}
		currentSquare.Lines = append(currentSquare.Lines, line)
	}
	squares[currentSquare.Name] = calcSquareSides(currentSquare)

	fmt.Println("Number of squares", len(squares))

	corners, firstCorner := getCorners(squares)
	fmt.Println("Corner Count", len(corners))
	cornerSum := 1
	for _, sq := range corners {
		cornerSum *= sq.Number
	}
	fmt.Println("Corner Sum", cornerSum)

}

func getCorners(squares map[string]Square) (map[string]Square, Square) {
	corners := map[string]Square{}
	first := Square{}

	for nName, needle := range squares {

		matchCount := 0

		for hName, haystack := range squares {
			if nName == hName {
				continue
			}
			match := false
			for _, nSide := range needle.Sides {
				for _, hSide := range haystack.Sides {
					if nSide == hSide {
						match = true
						break
					}
				}
				if match {
					break
				}
			}
			if match {
				matchCount++
			}
		}
		if matchCount == 2 {
			corners[nName] = needle
			first = needle
		}
	}
	return corners, first
}
func calcSquareSides(in Square) Square {
	in.Sides = append(in.Sides, in.Lines[0])
	in.Sides = append(in.Sides, hlp.StringReverse(in.Lines[0]))
	in.Sides = append(in.Sides, in.Lines[9])
	in.Sides = append(in.Sides, hlp.StringReverse(in.Lines[9]))
	a := ""
	b := ""
	for _, line := range in.Lines {
		l := hlp.StringToSlice(line)
		a += l[0]
		b += l[9]
	}
	in.Sides = append(in.Sides, a)
	in.Sides = append(in.Sides, hlp.StringReverse(a))
	in.Sides = append(in.Sides, b)
	in.Sides = append(in.Sides, hlp.StringReverse(b))

	tmp := strings.Replace(in.Name, "Tile ", "", -1)
	tmp = strings.Replace(tmp, ":", "", -1)
	in.Number = hlp.Atoi(tmp)

	in.Matrix = hlp.StringMatrix{}
	in.Matrix.InitFromLines(in.Lines)
	return in

}
