package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strings"
)

var ints []int
var memo map[int]int

func main() {

	lines := hlp.FileToLines("input.txt", true)

	now := hlp.Atoi(lines[0])

	busses := strings.Split(lines[1], ",")

	minDeparture := now * 2
	earliestBus := 0
	minWait := 0
	for _, bus := range busses {
		if bus != "x" {
			departure := 0
			wait := 0
			id := hlp.Atoi(bus)
			for wait = 0; wait < id; wait++ {
				if (now+wait)%id == 0 {
					departure = now + wait
					break
				}
			}

			if departure < minDeparture {
				minDeparture = departure
				earliestBus = id
				minWait = wait
			}
		}
	}

	fmt.Println("Earliest bus", earliestBus, minWait, earliestBus*minWait)

	timeSlots := map[int]int{}
	intBusses := []int{}
	for i, bus := range busses {
		if bus != "x" {
			id := hlp.Atoi(bus)
			timeSlots[i] = id
			intBusses = append(intBusses, id)
		} else {
			intBusses = append(intBusses, 0)
		}
	}

	t := 0
	next := 1
	incrementor := intBusses[0]
	for true {
		if intBusses[next] == 0 {
			next++
			continue
		}

		if (t+next)%intBusses[next] == 0 {
			incrementor *= intBusses[next]

			valid := true
			for ts, id := range timeSlots {
				if (t+ts)%id != 0 {
					valid = false
					break
				}
			}
			if valid {
				fmt.Println("Valid", t)
				break
			}
			next++
		}

		t += incrementor
	}

}
