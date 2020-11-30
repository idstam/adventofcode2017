package main

import "fmt"

func main() {
	lines := fileToLines("example.txt")
	m := [][]string{}
	for _, l := range lines {
		m = append(m, StringToSlice(l))
	}

	points := []IntPoint{}

	for y, l := range m {
		for x, c := range l {
			if c == "#" {
				points = append(points, IntPoint{X: x, Y: y})
			}
		}
	}

	maxPoint := IntPoint{}
	maxSeen := 0
	for _, pa := range points {
		seen := 0
		for _, pb := range points {
			blocked := false
			if pa.SameAs(pb) {
				continue
			}

			dx := IntDirection(pa.X, pb.X)
			dy := IntDirection(pa.Y, pb.Y)

			if dx == 0 {
				for y := pa.Y; y != pb.Y; y += dy {
					if m[y][pa.X] == "." {
						continue
					}
					pc := IntPoint{pa.X, y}
					if pc.SameAs(pa) || pc.SameAs(pb) {
						continue
					}
					if IntBetween(y, pa.Y, pb.Y) {
						blocked = true
					}
				}
				if !blocked {
					seen++
					continue
				}
			}

			if dy == 0 {
				for x := pa.X; x != pb.X; x += dx {
					if m[pa.Y][x] == "." {
						continue
					}
					pc := IntPoint{x, pa.Y}
					if pc.SameAs(pa) || pc.SameAs(pb) {
						continue
					}
					if IntBetween(x, pa.X, pb.X) {
						blocked = true
					}
				}
				if !blocked {
					seen++
					continue
				}
			}
			for x := pa.X; x != pb.X; x += dx {
				for y := pa.Y; y != pb.Y; y += dy {
					if m[y][x] == "." {
						continue
					}
					pc := IntPoint{x, y}
					if pc.SameAs(pa) || pc.SameAs(pb) {
						continue
					}
					if IntPointBetween(pc, pa, pb) {
						blocked = true
						break
					}
				}
				if blocked {
					break
				}
			}
			if !blocked {
				seen++
			}

		}
		if seen > maxSeen {
			maxSeen = seen
			maxPoint = pa
		}
	}

	fmt.Println(maxPoint, maxSeen)
}
