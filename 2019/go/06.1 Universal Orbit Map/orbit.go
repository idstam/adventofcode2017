package main

import "strings"

import "fmt"

type Constellation struct {
	Center   string
	Children map[string]Constellation
	Parent   string
}

var constellations map[string]Constellation

func main() {
	lines := fileToLines("input.txt")
	constellations = map[string]Constellation{}

	for _, line := range lines {
		o, orbiter := LineToOrbit(line)
		old, exists := constellations[o.Center]
		if exists {
			old.Children[orbiter] = o.Children[orbiter]
			constellations[o.Center] = old
		} else {
			constellations[o.Center] = o
		}
		for _, child := range constellations[o.Center].Children {
			old, exists = constellations[child.Center]
			if !exists {
				constellations[child.Center] = child
			}
		}
	}
	dist := 0
	for _, c := range constellations {
		if c.Center != "COM" {
			dist += DistanceToCOM(c, 0)
		}
	}

	fmt.Println(dist)
}

func DistanceToCOM(con Constellation, count int) int {
	pName := FindParent(con)

	if pName == "COM" {
		return count + 1
	}
	count++
	return DistanceToCOM(constellations[pName], count)
}

func FindParent(con Constellation) string {
	if con.Parent != "" || con.Center == "COM" {
		return con.Center
	}
	for _, p := range constellations {
		for _, c := range p.Children {
			if c.Center == con.Center {
				return p.Center
			}
		}
	}
	return "NOT FOUND"
}
func LineToOrbit(line string) (Constellation, string) {
	planets := strings.Split(line, ")")
	center := Constellation{Center: planets[0], Children: map[string]Constellation{}}
	center.Children[planets[1]] = Constellation{Center: planets[1], Children: map[string]Constellation{}}
	return center, planets[1]
}
