package main

import "fmt"

func main() {

	input := 314
	ring := Ring{}
	ring.Init()
	ring.Insert(0)

	//Part 1
	// for i := 1; i <= 2017; i++ {
	// 	ring.Move(input)
	// 	ring.Insert(i)
	// }
	// ring.Move(1)
	//fmt.Println(ring.Value())

	pointer := 0
	//atOne := 0
	for len := 1; len <= 50000000; len++ {
		//Move
		realSteps := input % len
		pointer += realSteps
		if pointer >= len {
			pointer -= len
		}
		if pointer == 0 {
			fmt.Println(len)
		}
		pointer++
	}

	//ring.SetPos(1)
	//fmt.Println(ring.Value())

}

type Ring struct {
	//https://github.com/golang/go/wiki/SliceTricks
	buffer  []int
	pointer int
}

//Init initialises the Ring
func (r *Ring) Init() {
	r.pointer = -1
}

//Pos is the current position in the ring buffer
func (r *Ring) Pos() int {
	return r.pointer
}
func (r *Ring) SetPos(pos int) {
	r.pointer = pos
}

//Value the value at the current position
func (r *Ring) Value() int {
	return r.buffer[r.pointer]
}

//Move move the pointer and return its position fter the move
func (r *Ring) Move(steps int) int {
	realSteps := steps % len(r.buffer)
	r.pointer += realSteps
	if r.pointer >= len(r.buffer) {
		r.pointer -= len(r.buffer)
	}

	return r.pointer
}

//Insert insert a value and return the size
func (r *Ring) Insert(value int) int {
	r.pointer++
	if r.pointer == len(r.buffer) {
		r.buffer = append(r.buffer, value)
	} else {

		r.buffer = append(r.buffer, 0)
		copy(r.buffer[r.pointer+1:], r.buffer[r.pointer:])
		r.buffer[r.pointer] = value
	}
	return len(r.buffer)
}

/*

	r := ring.New(1)

	r.Value = 0
	zero := r
	input := 3

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

*/
