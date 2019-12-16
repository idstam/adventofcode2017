package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var recepies map[string]string
var restChems map[string]int

func main() {
	lines := FileToLines("example.txt")

	recepies = map[string]string{}
	restChems = map[string]int{}
	for _, line := range lines {
		line = strings.ReplaceAll(line, ">", "")
		recepieTokens := strings.Split(line, "=")
		result := strings.TrimSpace(recepieTokens[1])
		recepie := strings.TrimSpace(recepieTokens[0])
		recepies[result] = recepie
	}

	demand := GetDemand("2 AB")
	fmt.Println(demand)
	ores := strings.Split(demand, ",")

	count := int64(0)
	for _, ore := range ores {
		tokens := strings.Split(ore, " ")
		if len(tokens) != 3 {
			continue
		}
		a, _ := strconv.Atoi(tokens[0])
		b, _ := strconv.Atoi(tokens[1])
		count += (int64(a) * int64(b))
	}
	fmt.Println("Ore demand", count)
}

func GetDemand(result string) string {
	amount, recepie := GetRecepie(result)
	demands := strings.Split(recepie, ",")
	ret := ""
	for _, demand := range demands {
		demand = strings.TrimSpace(demand)
		if strings.Contains(demand, "ORE") {
			ret += strconv.Itoa(amount) + " " + demand + ","
		} else {
			ret += GetDemand(demand)
		}
	}

	return ret
}
func GetRecepie(result string) (int, string) {
	recepie := recepies[result]
	if recepie != "" {
		return 1, recepie
	}

	resultTokens := strings.Split(result, " ")
	chemical := resultTokens[1]
	amount, _ := strconv.Atoi(resultTokens[0])
	for k, v := range recepies {
		resultTokens = strings.Split(k, " ")
		if resultTokens[1] == chemical {
			recepieAmount, _ := strconv.Atoi(resultTokens[0])
			//If there are enough chems left from an old run, return from them
			if amount <= restChems[chemical] {
				restChems[chemical] = restChems[chemical] - amount
				return 0, v // Zero since we didn't need to create any new chems
			}

			//If there are some chems left from an old run, use them
			amount = amount - restChems[chemical]
			foo := (amount / recepieAmount)
			if amount%recepieAmount != 0 {
				foo++
			}
			totMade := foo * recepieAmount
			restChems[chemical] = totMade - amount
			consumed := CalcConsumption(amount, v)
			return foo, consumed

		}
	}
	log.Fatal("Found no recepie")
	return 0, ""
}
func CalcConsumption(amount int, recepie string) string {
	chemicals := strings.Split(recepie, ",")
	ret := ""
	for _, chemical := range chemicals {
		chemical = strings.TrimSpace(chemical)
		cTokens := strings.Split(chemical, " ")
		a, _ := strconv.Atoi(cTokens[0])
		c := cTokens[1]
		a *= amount
		ret += strconv.Itoa(a) + " " + c + ","
	}
	return strings.TrimSuffix(ret, ",")
}
