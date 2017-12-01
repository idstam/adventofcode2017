package main

import (
	"fmt"
)

func main() {

	var captcha = "1122"

	var foo = captcha + captcha[0:1]

	fmt.Println(foo)
}
