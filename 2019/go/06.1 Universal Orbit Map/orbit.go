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

	youPathToCOM := GetPathToAncestor("COM", constellations["YOU"], []string{"YOU"})
	santaPathToCOM := GetPathToAncestor("COM", constellations["SAN"], []string{"SAN"})
	//fmt.Println(youPathToCOM)
	//fmt.Println(santaPathToCOM)

	firstCommon := FindFirstCommon(youPathToCOM, santaPathToCOM)

	youDist :=DistanceToAncestor(firstCommon, constellations["YOU"], 0) 
	santaDist :=DistanceToAncestor(firstCommon, constellations["SAN"], 0) 
	
	fmt.Println(youDist + santaDist -2)

	//dist := 0
	// for _, c := range constellations {
	// 	if c.Center != "COM" {
	// 		dist += DistanceToAncestor("COM", c, 0)
	// 	}
	// }

	//dist = DistanceToAncestor("E", constellations["YOU"], 0)
	//fmt.Println(dist)
}

func FindFirstCommon(aa, bb []string)string{
	for _, a := range aa{
		for _, b := range bb{
			if a== b{
				return a
			}
		}
	}

	return "NO COMMON ORBIT FOUND"
}
func GetPathToAncestor(root string, con Constellation, path []string)[]string{
	pName := FindParent(con)

	if pName == root {
		return append(path, pName)
	}
	
	return GetPathToAncestor(root, constellations[pName], append(path, pName))

}


func DistanceToAncestor(root string, con Constellation, count int) int {
	pName := FindParent(con)

	if pName == root {
		return count + 1
	}
	count++
	return DistanceToAncestor(root, constellations[pName], count)
}

func FindParent(con Constellation) string {
	if con.Parent != "" || con.Center == "COM" {
		return con.Parent
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
	center.Children[planets[1]] = Constellation{
		Center: planets[1], 
		Children: map[string]Constellation{}, 
		Parent : planets[0],
	}
	return center, planets[1]
}
