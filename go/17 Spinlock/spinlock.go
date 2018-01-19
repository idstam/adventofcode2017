package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(1)

	r.Value = 0

	input := 376

	for i := 1; i < 2018; i++ {
		r = r.Move(input)
		n := ring.New(1)
		n.Value = i
		r.Link(n)
		r = n
	}
	end := r.Next()

	fmt.Println(end.Value)

}
