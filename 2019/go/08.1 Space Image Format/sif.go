package main

func main() {
	inputLines := fileToLines("input.txt")
	data := stringToSlice(inputLines[0])

	width := 25
	height := 6
	layers := len(data) / (width * height)
	pixels := MakeSquareStringMatrix(25, "2")

	z := 9999999
	for l := 0; l < layers; l++ {
		counts := map[string]int{}

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				c := data[(l*(height*width))+(y*width)+x]
				counts[c] = counts[c] + 1
				//fmt.Print(c)
				if pixels[y][x] == "2" {
					pixels[y][x] = c
				}
			}
			//fmt.Println("")
		}
		if counts["0"] < z {
			println(l, counts["0"], counts["1"], counts["2"], counts["1"]*counts["2"])
			z = counts["0"]
		}
	}
	//fmt.Println(minRowOnes, minRowTwos, (minRowOnes * minRowTwos))
	DumpStringMatrix(pixels, "")
}
