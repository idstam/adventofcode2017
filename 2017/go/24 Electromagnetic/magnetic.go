package main

import (
	"fmt"
	"strconv"
	"strings"
)

var allParts []Part

type Part struct {
	Name     string
	Path string
	EndA     int
	EndB     int
	Parent   *Part
	Children map[string]Part
}

func main() {
	lines := fileToLines("example.txt")

	for _, line := range lines {
		allParts = append(allParts, lineToPart(line))

	}

	root := Part{Name: "0/0"}
	root.Children = getChildTree(root)

	root.PrintAllBridges()
	fmt.Println("asdf")
}

func (p Part)PrintAllBridges(){
	if len(p.Children) == 0{
		fmt.Println	(p.Path)
		return
	}
	for _, c := range p.Children{
		c.PrintAllBridges()
	}
}


func (parent Part) HasUsed(name string) bool {
	_, exists := parent.Children[name]
	if parent.Name == name {
		return true
	}
	if exists {
		return true
	}
	if parent.Name == "0/0" {
		return false
	}
	return parent.Parent.HasUsed(name)
}
func getChildTree(this Part) map[string]Part {
	ret := map[string]Part{}
	for _, p := range allParts {
		if !this.HasUsed(p.Name) {
			if this.EndA == p.EndA ||
				this.EndB == p.EndB ||
				this.EndA == p.EndB ||
				this.EndB == p.EndA {
				p.Parent = &this
				p.Path = this.Path + " - " + p.Name
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
