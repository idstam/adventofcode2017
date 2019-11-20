package main

import "fmt"

func main() {
	size := 1000
	world := initWorld(size)
	dx := 0
	dy := -1
	x := size / 2
	y := size / 2
	infectCount := 0

	//dumpStringMatrix(world, "Initial")
	rounds := 10000000
	for t := 1; t < rounds; t++ {
		switch world[y][x] {
		case "#":
			dx, dy = turnRight(dx, dy)
			world[y][x] = "f"
		case ".":
			dx, dy = turnLeft(dx, dy)
			world[y][x] = "w"
		case "w":
			world[y][x] = "#"
			infectCount++
		case "f":
			dx, dy = reverse(dx, dy)
			world[y][x] = "."
		}
		//dumpStringMatrix(world, "Tick "+strconv.Itoa(t))
		x += dx
		y += dy
	}
	fmt.Printf("Infection count after %d %d", rounds, infectCount)
}

func main_a() {
	size := 1000
	world := initWorld(size)
	dx := 0
	dy := -1
	x := size / 2
	y := size / 2
	infectCount := 0

	//dumpStringMatrix(world, "Initial")
	for t := 1; t < 10000; t++ {
		if world[y][x] == "#" {
			dx, dy = turnRight(dx, dy)
			world[y][x] = "."

		} else {
			dx, dy = turnLeft(dx, dy)
			world[y][x] = "#"
			infectCount++
		}
		//dumpStringMatrix(world, "Tick "+strconv.Itoa(t))
		x += dx
		y += dy
	}
	fmt.Printf("Infection count after 10000 %d", infectCount)
}

func reverse(dx, dy int) (int, int) {
	return dx * -1, dy * -1
}
func turnRight(dx, dy int) (int, int) {
	if dx == 0 && dy == -1 {
		return 1, 0
	}
	if dx == 0 && dy == 1 {
		return -1, 0
	}
	if dx == -1 && dy == 0 {
		return 0, -1
	}
	if dx == 1 && dy == 0 {
		return 0, 1
	}

	return 100000, 100000

}
func turnLeft(dx, dy int) (int, int) {
	if dx == 0 && dy == -1 {
		return -1, 0
	}
	if dx == 0 && dy == 1 {
		return 1, 0
	}
	if dx == -1 && dy == 0 {
		return 0, 1
	}
	if dx == 1 && dy == 0 {
		return 0, -1
	}

	return 100000, 100000

}

func initWorld(size int) [][]string {
	ret := [][]string{}

	for y := 0; y < size; y++ {
		ret = append(ret, []string{})
		for x := 0; x < size; x++ {
			ret[y] = append(ret[y], ".")
		}
	}
	lines := fileToLines("input.txt")
	offset := (size / 2) - (len(lines) / 2)
	for y, line := range lines {
		stringSlice := stringToSlice(line)
		for x, s := range stringSlice {
			if s == "." {
				ret[y+offset][x+offset] = "."
			} else {
				ret[y+offset][x+offset] = "#"
			}
		}
	}
	return ret
}
