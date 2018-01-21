package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(1)

	r.Value = 0
	zero := r
	input := 376

	rings := make([]*ring.Ring, 50000000)
	for i := 1; i <= 50000000; i++ {
		rings[i-1] = ring.New(1)
	}
	for i := 1; i <= 50000000; i++ {
		r = r.Move(input)
		n := rings[i-1]
		n.Value = i
		r.Link(n)
		r = n
	}
	end := zero.Next()

	fmt.Println(end.Value)

}
