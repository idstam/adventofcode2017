package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strings"
)

type Rule struct {
	Class string
	AMin  int
	AMax  int
	BMin  int
	BMax  int
}

func main() {

	lines := hlp.FileToLines("input.txt", true)

	rules := map[string]Rule{}
	tokens := []string{}
	lineType := 0
	errorRate := 0
	myTicket := []int{}
	nearbyTickets := [][]int{}
	ruleLines := []string{}
	for _, l := range lines {
		if l == "your ticket:" || l == "nearby tickets:" {
			lineType++
			continue
		}
		if lineType == 0 {
			r := mapLineToRule(l)
			rules[r.Class] = r
			ruleLines = append(ruleLines, l)
			continue
		}

		if lineType == 1 {
			tokens = strings.Split(l, ",")
			myTicket = hlp.StringToIntArray(tokens)
			continue
		}

		if lineType == 2 {
			tokens = strings.Split(l, ",")
			ints := hlp.StringToIntArray(tokens)

			ticketHasError := false
			for _, field := range ints {
				matched := false
				for _, re := range rules {
					if hlp.IntBetween(field, re.AMin, re.AMax) || hlp.IntBetween(field, re.BMin, re.BMax) {
						matched = true
					}
				}
				if !matched {
					errorRate += field
					ticketHasError = true
					break
				}
			}
			if !ticketHasError {
				nearbyTickets = append(nearbyTickets, ints)
			}
		}
	}
	fmt.Println("Error Scanning Rate", errorRate) //Should be 23954
	fmt.Println("My Ticket", myTicket)

	matchedRules := map[string]int{}
	matchedFields := map[int]string{}

	for i := 0; i < 1000000; i++ { //Avoid endless loop
		matchCount := map[string]int{}
		ruleField := map[string]int{}
		for fieldNumber := 0; fieldNumber < len(myTicket); fieldNumber++ {
			for _, ruleLine := range ruleLines {

				rule := mapLineToRule(ruleLine)
				if matchedRules[rule.Class] != 0 || matchedFields[fieldNumber] != "" {
					continue
				}

				matched := true

				for _, ticket := range nearbyTickets {
					field := ticket[fieldNumber]
					if !(hlp.IntBetween(field, rule.AMin, rule.AMax) || hlp.IntBetween(field, rule.BMin, rule.BMax)) {
						matched = false
						break
					}
				}
				if matched {
					matchCount[rule.Class] = matchCount[rule.Class] + 1
					ruleField[rule.Class] = fieldNumber
				}
			}
		}
		if len(matchCount) == 0 {
			break
		}
		matchingRule := findSingleMatch(matchCount)
		matchedRules[matchingRule] = ruleField[matchingRule]
		matchedFields[ruleField[matchingRule]] = matchingRule

	}
	//fmt.Printf("Field Numbers %s \n", hlp.PrettyPrint(matchedRules))
	//fmt.Printf("Field Numbers %s \n", hlp.PrettyPrint(fieldNumbers))
	sum := 1
	for rule, field := range matchedRules {
		if strings.HasPrefix(rule, "departure") {
			v := myTicket[field]
			sum *= v
		}
	}

	fmt.Println("Product of departure fields", sum) //Answer 453459307723
}

func findSingleMatch(in map[string]int) string {
	for k, v := range in {
		if v == 1 {
			return k
		}
	}
	return ""
}

//1382560228363 high
//1169409308827
//1169409308827
//  46084864327
//    147778349 low

func mapLineToRule(line string) Rule {
	fields := strings.Split(line, ":")

	tokens := strings.Split(strings.TrimSpace(fields[1]), " ")
	rule := Rule{Class: fields[0]}
	minMax := strings.Split(tokens[0], "-")
	rule.AMin = hlp.Atoi(minMax[0])
	rule.AMax = hlp.Atoi(minMax[1])

	minMax = strings.Split(tokens[2], "-")
	rule.BMin = hlp.Atoi(minMax[0])
	rule.BMax = hlp.Atoi(minMax[1])

	return rule
}
