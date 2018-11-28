package main

import (
	"bufio"
	"container/ring"
	_ "container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MemSize = 256

func main() {
	main2()
}
func main2() {

	input := "106,118,236,1,130,0,235,254,59,205,2,87,129,25,255,118"
	//input := "AoC 2017"
	ints := make([]int, MemSize, 256)
	lengths := stringToCharCodeList(input)
	lengths = append(lengths, 17, 31, 73, 47, 23)

	for i := 0; i < MemSize; i++ {
		ints[i] = i
	}

	pos := 0
	//pos2 := 0
	skipLength := 0
	for round := 0; round < 64; round++ {
		for _, length := range lengths {
			//pos2 = pos
			sub := make([]int, length, length)
			for i := range sub {
				address := getAddress(pos + i)
				sub[i] = ints[address]
			}
			//fmt.Println(sub)
			for i := range sub {
				address := getAddress(pos + i)
				ints[address] = sub[length-1-i]
			}
			//fmt.Println(ints)
			pos = pos + skipLength + length
			skipLength++
		}
		//skipLength--
		//pos = pos2
	}

	denseHash := make([]int, 16, 16)
	for j := 0; j < 16; j++ {
		i := j * 16
		denseHash[j] = ints[0+i] ^ ints[1+i] ^ ints[2+i] ^ ints[3+i] ^
			ints[4+i] ^ ints[5+i] ^ ints[6+i] ^ ints[7+i] ^
			ints[8+i] ^ ints[9+i] ^ ints[10+i] ^ ints[11+i] ^
			ints[12+i] ^ ints[13+i] ^ ints[14+i] ^ ints[15+i]
	}

	for i := range denseHash {
		fmt.Printf(fmt.Sprintf("%02x", denseHash[i]))
	}

}

func stringToCharCodeList(s string) []int {
	ret := make([]int, len(s), len(s))
	for i := range s {
		ret[i] = int(charCodeAt(s, i))
	}
	return ret
}
func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}
func getAddress(pos int) int {
	ret := pos % MemSize
	return ret
}
func printRing(rng *ring.Ring) {

	fmt.Println("-----------------------")
	for i := 0; i < 5; i++ {
		fmt.Println(rng.Value)
		rng = rng.Next()
	}
}
func fileToLines(fileName string) []string {
	ret := make([]string, 1, 100)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		ret = append(ret, line)
	}

	return ret
}
func stringsToInts(numbers []string) []int {
	ret := make([]int, len(numbers), len(numbers))
	for i, v := range numbers {
		ret[i], _ = strconv.Atoi(v)
	}
	return ret
}

func intsToRing(ints []int) *ring.Ring {
	ret := ring.New(len(ints))

	for _, v := range ints {
		ret.Value = v
		ret = ret.Next()
	}

	return ret
}
