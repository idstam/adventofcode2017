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
	key := ""
	for _, c := range lines[0] {
		switch c {
		case '<':
			sPosX--
			break
		case '>':
			sPosX++
			break
		case '^':
			sPosY--
			break
		case 'v':
			sPosY++
			break
		}

		key = strconv.Itoa(sPosX) + "," + strconv.Itoa(sPosY)
		karta[key] = karta[key] + 1
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
