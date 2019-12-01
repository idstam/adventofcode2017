package main

import "fmt"

func main() {

	lines := fileToLines("input.txt")

	ints := StringToIntArray(lines)
	ints = DivideByOnIntArray(ints, 3)
	ints = SubtractOnIntArray(ints, 2)
	fuels := ExtraFuelArray(ints)

	sumFuel := SumIntArray(fuels)
	sum := SumIntArray(ints)
	fmt.Printf("Fule requirement for modules: %d\n", sum)
	fmt.Printf("Fule requirement for modules and it's fuel: %d\n", sum+sumFuel)
}

func ExtraFuelArray(in []int) []int {
	ret := []int{}
	for _, i := range in {
		ret = append(ret, ExtraFuel(i))
	}
	return ret
}
func ExtraFuel(mass int) int {
	f := (mass / 3) - 2
	ret := mass
	for f > 0 {
		ret += f
		f = (f / 3) - 2
	}

	return ret - mass
}
