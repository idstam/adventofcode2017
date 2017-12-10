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

	//input2 := "3,4,1,5"
	input2 := "106,118,236,1,130,0,235,254,59,205,2,87,129,25,255,118"
	ints := make([]int, MemSize, 256)

	lengths := stringsToInts(strings.Split(input2, ","))

	for i := 0; i < MemSize; i++ {
		ints[i] = i
	}

	pos := 0
	for skipLength, length := range lengths {
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

	}

	fmt.Println(ints)

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

func main1() {

	//input2 := "3,4,1,5"
	input2 := "106,118,236,1,130,0,235,254,59,205,2,87,129,25,255,118"
	ints := make([]int, 256, 256)

	lengths := stringsToInts(strings.Split(input2, ","))

	for i := 0; i < len(ints); i++ {
		ints[i] = i
	}

	rng := intsToRing(ints)
	startPos := rng
	//curPos := 0
	//skipSize := 0
	//printRing(rng)

	for skipSize, length := range lengths {
		if length < 1 {
			continue
		}

		fmt.Println(rng.Value)
		subList := make([]interface{}, length, length)
		for i := range subList {
			subList[i] = rng.Value
			rng = rng.Next()
		}
		for _, v := range subList {
			rng = rng.Prev()
			rng.Value = v

		}
		//printRing(startPos)

		rng = rng.Move(skipSize + length)
	}
	fmt.Println(rng.Value)
	printRing(startPos)
	println("Klar")
}
