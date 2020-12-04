package main

import (
	"fmt"
	hlp "idstam/aoc_idstam/2020/go/helpers"
	"strconv"
	"strings"
)

func main() {

	lines := hlp.FileToLines("input.txt", false)

	passport := map[string]string{}
	validCount1 := 0
	validCount2 := 0
	firstCheck := false
	for _, l := range lines {
		if l == "" {
			if len(passport) == 8 || (len(passport) == 7 && passport["cid"] == "") {
				validCount1++
				firstCheck = true
			}
			if firstCheck && checkValid2(passport) {
				validCount2++
			}

			firstCheck = false
			passport = map[string]string{}
			continue
		}

		lineTokens := strings.Split(l, " ")
		for _, lt := range lineTokens {
			tokens := strings.Split(lt, ":")
			passport[tokens[0]] = tokens[1]
		}
	}

	if len(passport) == 8 || (len(passport) == 7 && passport["cid"] == "") {
		validCount1++

		if checkValid2(passport) {
			validCount2++
		}

	}

	fmt.Println("ValidCount1", validCount1) //239
	fmt.Println("ValidCount2", validCount2) //188
}

func checkValid2(passport map[string]string) bool {

	byr, err := strconv.Atoi(passport["byr"])
	if err != nil || (byr < 1920 || byr > 2002) {
		return false
	}

	iyr, err := strconv.Atoi(passport["iyr"])
	if err != nil || (iyr < 2010 || iyr > 2020) {
		return false
	}

	eyr, err := strconv.Atoi(passport["eyr"])
	if err != nil || (eyr < 2020 || eyr > 2030) {
		return false
	}

	if !(strings.HasSuffix(passport["hgt"], "cm") || strings.HasSuffix(passport["hgt"], "in")) {
		return false
	}

	if strings.HasSuffix(passport["hgt"], "cm") {
		hgt, err := strconv.Atoi(strings.Replace(passport["hgt"], "cm", "", 1))
		if err != nil || (hgt < 150 || hgt > 193) {
			return false
		}
	}

	if strings.HasSuffix(passport["hgt"], "in") {
		hgt, err := strconv.Atoi(strings.Replace(passport["hgt"], "in", "", 1))
		if err != nil || (hgt < 59 || hgt > 76) {
			return false
		}
	}

	if len(passport["hcl"]) != 7 {
		return false
	}
	if !strings.HasPrefix(passport["hcl"], "#") {
		return false
	}

	_, err = strconv.ParseUint(strings.Replace(passport["hcl"], "#", "", 1), 16, 64)
	if err != nil {
		return false
	}

	if !strings.Contains("amb blu brn gry grn hzl oth", passport["ecl"]) {
		return false
	}

	if len(passport["pid"]) != 9 {
		return false
	}
	_, err = strconv.ParseUint(passport["pid"], 10, 64)
	if err != nil {
		return false
	}

	return true
}
