package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strconv"
	"strings"
)

func main() {

	lines := hlp.FileToLines("input.txt", true)

	mem := map[int]int64{}
	mask := []string{}
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if tokens[0] == "mask" {
			mask = hlp.StringToSlice(tokens[2])
		} else {
			adress := valFromToken(tokens[0])
			value := hlp.Atoi(tokens[2])
			mem[adress] = applyMask(value, mask)
		}
	}
	result := int64(0)
	for _, v := range mem {
		result += v
	}
	fmt.Println("Part1", result)

	mem = map[int]int64{}
	mask = []string{}
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if tokens[0] == "mask" {
			mask = hlp.StringToSlice(tokens[2])
		} else {
			adress := valFromToken(tokens[0])
			value := hlp.Atoi(tokens[2])
			adresses := applyMask2(adress, mask)
			for _, appliedAdress := range adresses {
				mem[int(appliedAdress)] = int64(value)
			}
		}
	}

	result = int64(0)
	for _, v := range mem {
		result += v
	}
	fmt.Println("Part2", result)

}

func valFromToken(token string) int {
	token = strings.Replace(token, "mem[", "", 1)
	token = strings.Replace(token, "]", "", 1)
	return hlp.Atoi(token)
}
func applyMask(value int, mask []string) int64 {
	valStr := fmt.Sprintf("%b", value)
	valStr = fmt.Sprintf("%036s", valStr)
	values := hlp.StringToSlice(valStr)

	fmt.Println(valStr)

	result := ""

	for i, m := range mask {
		switch m {
		case "0":
			result += "0"
		case "1":
			result += "1"
		default:
			result += values[i]
		}
	}

	ret, _ := strconv.ParseInt(result, 2, 64)
	return ret
}

func applyMask2(value int, mask []string) []int64 {
	valStr := fmt.Sprintf("%b", value)
	valStr = fmt.Sprintf("%036s", valStr)
	values := hlp.StringToSlice(valStr)

	//fmt.Println(valStr)

	result := ""

	for i, m := range mask {
		switch m {
		case "0":
			result += values[i]
		case "1":
			result += "1"
		default:
			result += "x"
		}
	}
	foundX := true
	results := [][]string{}
	results = append(results, hlp.StringToSlice(result))

	for foundX {
		foundX = false

		for _, fr := range results {
			for i, m := range fr {
				if string(m) == "x" {
					tmp := make([]string, len(fr))
					fr[i] = "0"
					copy(tmp, fr)
					tmp[i] = "1"
					results = append(results, tmp)
					foundX = true
					break
				}
			}
		}
	}

	returns := []int64{}
	for _, r := range results {
		res := hlp.SliceToString(r)
		ret, _ := strconv.ParseInt(res, 2, 64)
		returns = append(returns, ret)
		//fmt.Println(r)
	}
	return returns
}
