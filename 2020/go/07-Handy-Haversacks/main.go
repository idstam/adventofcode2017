package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strconv"
	"strings"
)

var isIn map[string]bagType
var rules map[string]bagType

func main() {

	lines := hlp.FileToLines("input.txt", false)

	rules = map[string]bagType{}
	isIn = map[string]bagType{}

	for _, l := range lines {
		rule := mapLineToRule(l)
		rules[rule.color] = rule
		for _, subBag := range rule.content {
			b, exists := isIn[subBag.color]
			if !exists {
				newBag := bagType{}
				newBag.color = subBag.color
				newBag.content = map[string]bagType{}
				isIn[subBag.color] = newBag
				b = newBag
			}
			b.content[rule.color] = rule
		}
	}

	outerBags := findOuterBags("shiny gold")
	uniqueOuterBags := map[string]bool{}
	for _, b := range outerBags {
		uniqueOuterBags[b.color] = true
	}
	fmt.Println("Number of outer bags", len(uniqueOuterBags)) // 302
	fmt.Println("Number of rules", len(rules))

	c := countSubBags("shiny gold")
	fmt.Println("Number of inner bags", c) //4165
}

type bagType struct {
	color   string
	count   int
	content map[string]bagType
}

func countSubBags(start string) int {
	ret := 0
	rule := rules[start]
	for _, bag := range rule.content {
		ret += bag.count
		ret += countSubBags(bag.color) * bag.count
	}
	return ret
}
func findOuterBags(find string) []bagType {
	ret := []bagType{}
	parents := isIn[find]
	for _, rule := range parents.content {
		ret = append(ret, rule)
		parent, exists := isIn[rule.color]
		if exists {
			ret = append(ret, findOuterBags(parent.color)...)
		}
	}

	return ret
}

func mapLineToRule(line string) bagType {
	//light red bags contain 1 bright white bag, 2 muted yellow bags.
	tokens := strings.Split(line, " ")

	bag := bagType{}
	bag.content = map[string]bagType{}

	bag.color = tokens[0] + " " + tokens[1]

	for i := 4; i < len(tokens); i += 4 {
		if tokens[i] == "no" {
			break
		}
		subBag := bagType{}
		subBag.content = map[string]bagType{}

		subBag.count, _ = strconv.Atoi(tokens[i])
		subBag.color = tokens[i+1] + " " + tokens[i+2]

		bag.content[subBag.color] = subBag
	}

	return bag
}
