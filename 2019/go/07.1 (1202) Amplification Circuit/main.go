package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var mem []int
var output int
var inputs []int

func main() {

	lines := fileToLines("input.txt")
	strs := strings.Split(lines[0], ",")
	resultCache := map[string]int{}

	allPhases := IntPermutateArray([]int{0, 1, 2, 3, 4})
	inputs = []int{}
	lastOutput := 0

	maxThrust := 0
	for _, phases := range allPhases {
		lastOutput = 0
		for _, phase := range phases {
			mem = StringToIntArray(strs)
			inputs = IntPushArray(lastOutput, inputs)
			inputs = IntPushArray(phase, inputs)

			r, exists := resultCache[strconv.Itoa(phase)+":"+strconv.Itoa(lastOutput)]
			if exists {
				output = r
			} else {
				ptr := 0
				for ptr >= 0 {
					ptr = Exec(ptr)
				}
				resultCache[strconv.Itoa(phase)+":"+strconv.Itoa(lastOutput)] = output
			}

			lastOutput = output

		}
		if lastOutput > maxThrust {
			fmt.Println("More thrust ", lastOutput, phases)
			maxThrust = lastOutput
		}
	}

}

func Exec(ptr int) int {
	fullOp := mem[ptr]
	_, mb, mc, op := ParseOpCode(fullOp)
	if op == 99 {
		//fmt.Println("Exit")
		return -1
	}
	v1, v2, v3 := GetValues(ptr)
	//fmt.Println(OpName(op), ": ", v1,mc, v2,mb, v3,ma)
	//fmt.Println("Before 225:", mem[225])

	switch op {
	case 1:
		dest := v3[1]
		ExpandTape(dest + 1)
		mem[dest] = v1[mc] + v2[mb]
		ptr += 4
	case 2:
		dest := v3[1]
		ExpandTape(dest + 1)
		mem[dest] = v1[mc] * v2[mb]
		ptr += 4
	case 3:
		dest := v1[1]
		ExpandTape(dest + 1)
		mem[dest] = GetInput()
		ptr += 2
	case 4:
		val := v1[mc]
		output = val
		//fmt.Printf("Output: %d \n", val)
		ptr += 2
	case 5:
		if v1[mc] != 0 {
			ptr = v2[mb]
		} else {
			ptr += 3
		}
	case 6:
		if v1[mc] == 0 {
			ptr = v2[mb]
		} else {
			ptr += 3
		}
	case 7:
		dest := v3[1]
		ExpandTape(dest + 1)

		if v1[mc] < v2[mb] {
			mem[dest] = 1
		} else {
			mem[dest] = 0
		}
		ptr += 4
	case 8:
		dest := v3[1]
		ExpandTape(dest + 1)

		if v1[mc] == v2[mb] {
			mem[dest] = 1
		} else {
			mem[dest] = 0
		}
		ptr += 4

	case 99:
		fmt.Println("Exit")
		return -1
	default:
		log.Fatalf("Unknown OP Code %d \n", op)

	}
	//fmt.Println("After 225:", mem[225])
	return ptr
}
func OpName(op int) string {
	switch op {
	case 1:
		return "Add a+b -> c"
	case 2:
		return "Add a*b -> c"
	case 3:
		return "Input -> a"
	case 4:
		return "Output a"
	case 5:
		return "JIT a=!0 => b"
	case 6:
		return "JIF a==0 => b"
	case 7:
		return "LT a < b -> c"
	case 8:
		return "EQ a == b -> c"
	case 99:
		return "Exit"
	default:
		return fmt.Sprintf("UNKNOWN %d", op)

	}
}
func ExpandTape(newSize int) {
	if len(mem) >= newSize {
		return
	}

	for i := len(mem); i <= newSize; i++ {
		mem = append(mem, 0)
	}
}
func GetValues(adress int) ([]int, []int, []int) {
	v1 := []int{0, 0}
	v2 := []int{0, 0}
	v3 := []int{0, 0}
	v1 = GetValue(adress + 1)
	v2 = GetValue(adress + 2)
	v3 = GetValue(adress + 3)
	return v1, v2, v3
}
func GetValue(adress int) []int {
	ExpandTape(IntMax(adress, mem[adress]) + 1)
	ret := -9999999
	from := mem[adress]
	if from >= 0 {
		ret = mem[from]
	}
	return []int{ret, mem[adress]}

}
func GetInput() int {
	ret := 0
	ret, inputs = IntPopArray(inputs)
	return ret
}

func ParseOpCode(in int) (int, int, int, int) {
	op := in % 100
	in -= op
	c := (in % 1000) / 100
	in -= c * 100
	b := (in % 10000) / 1000
	in -= b * 1000
	a := (in % 100000) / 10000
	return a, b, c, op
}
