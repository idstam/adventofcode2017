package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var recepies map[string]Recepie
var restChems map[string]int64

type RecepieLine struct {
	Amount   int64
	Chemical string
	IsOre    bool
	Original string
}

type Recepie struct {
	Result         string
	ResultAmount   int64
	ResultChemical string
	Lines          []RecepieLine
}

func main() {

	lines := FileToLines("example.txt")

	recepies = map[string]Recepie{}
	restChems = map[string]int64{}
	for _, line := range lines {
		line = strings.ReplaceAll(line, ">", "")
		recepieTokens := strings.Split(line, "=")

		result := strings.TrimSpace(recepieTokens[1])
		recepie := strings.TrimSpace(recepieTokens[0])

		r := MakeRecepie(result, recepie)
		recepies[result] = r
	}

	count := int64(0)
	fuelCount := 0
	for consumed := int64(0); consumed <= 1000000000000; consumed += count {
		count = 0
		_, oreCount := GetDemand2(MakeRecepieLine("1 FUEL"), 0)
		//fmt.Println(demand)

		// ores := strings.Split(demand, ",")
		// for _, ore := range ores {
		// 	tokens := strings.Split(ore, " ")
		// 	// if len(tokens) != 3 {
		// 	// 	continue
		// 	// }
		// 	a, _ := strconv.Atoi(tokens[0])
		// 	//b, _ := strconv.Atoi(tokens[1])
		// 	count += int64(a)
		// }
		fuelCount++
		consumed = oreCount
	}
	//	fmt.Println("Ore demand", count)
	fmt.Println(fuelCount)
}

func GetDemand2(result RecepieLine, oreDemand int64) ([]RecepieLine, int64) {
	amount, recepie := GetRecepie(result)

	ret := []RecepieLine{}
	for i := int64(0); i < amount; i++ {
		for _, demand := range recepie.Lines {

			if demand.IsOre {
				oreDemand += demand.Amount
				ret = append(ret, demand)
			} else {
				ret, oreDemand = GetDemand2(demand, oreDemand)
			}
		}
	}

	return ret, oreDemand
}

// func GetDemand(result string) string {
// 	amount, recepie := GetRecepie(result)

// 	ret := ""
// 	for _, demand := range recepie.Lines {
// 		if demand.IsOre {
// 			ret += strconv.Itoa(amount) + " " + demand + ","
// 		} else {
// 			ret += GetDemand(demand)
// 		}
// 	}

// 	return ret
// }
func GetRecepie(result RecepieLine) (int64, Recepie) {
	recepie := recepies[result.Original]
	if recepie.Result != "" {
		return 1, recepie
	}

	chemical := result.Chemical
	amount := result.Amount
	for _, v := range recepies {

		if v.ResultChemical == chemical {

			//If there are enough chems left from an old run, return from them
			if amount <= restChems[chemical] {
				restChems[chemical] = restChems[chemical] - amount
				return 0, v // Zero since we didn't need to create any new chems
			}

			//If there are some chems left from an old run, use them
			amount = amount - restChems[chemical]
			foo := (amount / v.ResultAmount)
			if amount%v.ResultAmount != 0 {
				foo++
			}
			totMade := foo * v.ResultAmount
			restChems[chemical] = totMade - amount

			return foo, v

		}
	}
	log.Fatal("Found no recepie")
	return 0, Recepie{}
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

func MakeRecepie(result, recepie string) Recepie {
	ret := Recepie{
		Result: result,
		Lines:  []RecepieLine{},
	}
	resultTokens := strings.Split(result, " ")
	ret.ResultAmount, _ = strconv.ParseInt(resultTokens[0], 10, 64)
	ret.ResultChemical = resultTokens[1]

	recepieLines := strings.Split(recepie, ",")
	for _, rl := range recepieLines {
		l := MakeRecepieLine(rl)
		ret.Lines = append(ret.Lines, l)

	}
	return ret
}

func MakeRecepieLine(lineData string) RecepieLine {
	lineTokens := strings.Split(lineData, " ")
	l := RecepieLine{
		Chemical: lineTokens[1],
		Original: lineData,
	}
	l.Amount, _ = strconv.ParseInt(lineTokens[0], 10, 64)
	l.IsOre = (l.Chemical == "ORE")

	return l
}
