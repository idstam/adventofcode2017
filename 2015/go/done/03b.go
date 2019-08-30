package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := fileToLines("input03.txt")
	karta := map[string]int{}
	sPosX := 0
	sPosY := 0
	rPosX := 0
	rPosY := 0
	posX := 0
	posY := 0
	key := ""
	karta["0,0"] = 2
	isSanta := true
	direction := 1
	for _, c := range lines[0] {
		if isSanta {
			direction = 1
			posY = sPosY
			posX = sPosX
		} else {
			direction = 1
			posY = rPosY
			posX = rPosX
		}
		switch c {
		case '<':
			posX = posX - direction
			break
		case '>':
			posX = posX + direction
			break
		case '^':
			posY = posY - direction
			break
		case 'v':
			posY = posY + direction
			break
		}

		key = strconv.Itoa(posX) + "," + strconv.Itoa(posY)
		karta[key] = karta[key] + 1

		if isSanta {
			sPosY = posY
			sPosX = posX
		} else {
			rPosY = posY
			rPosX = posX
		}
		isSanta = !isSanta

	}
	ret := 0
	for _, i := range karta {
		if i > 0 {
			ret++
		}
	}

	fmt.Println(ret)
}

func minInt(numbers ...int) int {
	m := numbers[0]
	for _, n := range numbers {
		if n < m {
			m = n
		}
	}

	return m
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
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			ret = append(ret, line)
		}
	}

	return ret
}
