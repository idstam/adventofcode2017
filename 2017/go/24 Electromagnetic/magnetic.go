package main

import (
	"fmt"
	"strconv"
	"strings"
)

var allParts []Part

type Part struct {
	Name        string
	Path        string
	EndA        int
	EndB        int
	Parent      *Part
	Children    map[string]Part
	SumStrength int
}

func main() {

	// 1563 too low


	lines := fileToLines("input.txt")

	for _, line := range lines {
		allParts = append(allParts, lineToPart(line))

	}

	root := Part{Name: "0/0"}
	root.Children = getChildTree(root)

	maxStrength := root.PrintAllBridges(0)
	fmt.Println(maxStrength)
}

func (p Part) PrintAllBridges(lastMaxStrength int) int {
	if len(p.Children) == 0 {
		if p.SumStrength > lastMaxStrength{
			lastMaxStrength = p.SumStrength
		}
		fmt.Printf("Strength: %d Path: %s \n", p.SumStrength, p.Path)
		return lastMaxStrength
	}
	for _, c := range p.Children {
		lastMaxStrength = c.PrintAllBridges(lastMaxStrength)
	}
	return lastMaxStrength
}

func getChildTree(this Part) map[string]Part {
	ret := map[string]Part{}
	for _, p := range allParts {
		if !strings.Contains(this.Path, p.Name){
			if this.EndB == p.EndA {
				p.Parent = &this
				p.Path = this.Path + " - " + p.Name
				p.SumStrength = this.SumStrength + p.EndA + p.EndB
				p.Children = getChildTree(p)
				ret[p.Name] = p
			}
			if this.EndB == p.EndB {
				p.Parent = &this
				p.Path = this.Path + " - " + p.Name
				p.EndB = p.EndA
				p.EndA = this.EndB
				p.SumStrength = this.SumStrength + p.EndA + p.EndB
				p.Children = getChildTree(p)
				ret[p.Name] = p
			}
		}
	}
	return ret
}

func lineToPart(in string) Part {
	stringEnds := strings.Split(in, "/")
	ret := Part{
		Name: in,
	}
	ret.EndA, _ = strconv.Atoi(stringEnds[0])
	ret.EndB, _ = strconv.Atoi(stringEnds[1])

	return ret
}
