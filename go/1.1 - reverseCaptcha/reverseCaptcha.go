package main

import (
	"fmt"
	"strconv"
)

func main() {

	var captcha = "1122"

	var foo = captcha + captcha[0:1]

	i := 0
	sum := 0
	for i < len(foo)-1 {
		a := foo[i : i+1]
		b := foo[i+1 : i+2]
		if a == b {
			loopVal, _ := strconv.Atoi(foo[i : i+1])
			sum = sum + loopVal

		}
		i = i + 1
	}
	fmt.Println(sum)
}
