package main

import (
	"fmt"
	"strconv"
)

const MemSize = 256

var mem [][]int

func main() {
	input := "vbqugkhl"
	//input = "flqrgnkx"

	mem = make([][]int, 128, 128)

	//input := "AoC 2017"
	for line := 0; line < 128; line++ {
		key := fmt.Sprintf("%s-%d", input, line)
		knot := knotHash(key)
		bin := strToBin(knot)
		mem[line] = make([]int, 128, 128)
		for x, r := range bin {
			i, _ := strconv.Atoi(string(r))
			mem[line][x] = i * -1
		}
	}

	cnt := 0
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			fmt.Printf("%d %d %d\n", x, y, cnt)
			val := safeGet(x, y)
			if val == -1 {
				cnt++
				floodFill(x, y, cnt)
			}
		}
	}

	fmt.Println(cnt)
}

func floodFill(x, y, group int) {
	val := safeGet(x, y)
	if val == -1 {
		mem[x][y] = group
		floodFill(x+1, y, group)
		floodFill(x-1, y, group)
		floodFill(x, y+1, group)
		floodFill(x, y-1, group)
	}
}
func safeGet(x, y int) int {
	fmt.Printf("SafeGet %d %d \n", x, y)
	if x == 128 || y == 128 || x == -1 || y == -1 {
		return 0
	}
	return mem[x][y]
}

func knotHash(key string) string {
	ints := make([]int, MemSize, 256)
	lengths := stringToCharCodeList(key)
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
	hex := ""
	for i := range denseHash {
		hex += fmt.Sprintf(fmt.Sprintf("%02x", denseHash[i]))

	}
	return hex
}
func strToBin(s string) string {
	ret := ""

	for _, c := range s {
		n, _ := strconv.ParseUint(string(c), 16, 32)

		ret += fmt.Sprintf("%04b", n)
	}
	return ret
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
	ret := pos % 256
	return ret
}
