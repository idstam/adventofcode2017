package main

import "fmt"

func main() {
	//inA := 699
	//inB := 124

	inA := 65
	inB := 8921
	eqCount := 0

	for i := 0; i < 5; i++ {
		inA = generatorA(inA)
		inB = generatorB(inB)

		fmt.Printf("%d %d \n", inA, inB)
		eq := true
		mask := 1

		for j := 0; j < 16; j++ {

			a := inA & mask
			b := inB & mask
			if a != b {
				eq = false
				break
			}
			mask = mask << 1
		}
		//fmt.Printf("%32b\n", inA)
		//fmt.Printf("%32b\n", inB)

		if eq {
			//fmt.Println(i)
			eqCount++
		}

	}
	fmt.Println(eqCount)
}

func generatorA(input int) int {
	for {
		ret := input * 16807
		foo := ret % 2147483647
		foo2 := foo % 4
		if foo2 == 0 {
			return ret
		}
		input = ret
	}
}

func generatorB(input int) int {
	for {
		ret := input * 48271
		foo := ret % 2147483647
		foo2 := foo % 8
		if foo2 == 0 {
			return ret
		}
		input = ret

	}
}
