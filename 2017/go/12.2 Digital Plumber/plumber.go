package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	Name        string
	Connections map[string]int
	Group       int
}

func main() {
	lines := fileToLines("input2.txt")
	programs := make(map[string]*Program)

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		var program Program
		for i, token := range tokens {

			if i == 0 {
				program = Program{Name: token, Connections: make(map[string]int)}
			}
			if i > 1 {
				connection := strings.Replace(token, ",", "", -1)
				connection = strings.TrimSpace(connection)
				program.Connections[connection] = 1
			}
		}
		programs[program.Name] = &program
	}

	for _, program := range programs {
		for name := range program.Connections {
			for cn := range program.Connections {
				if name != cn {
					programs[name].Connections[cn] = 1
				}
			}
		}
	}

	groupCount := 0
	for _, program := range programs {
		group := 0
		for name := range program.Connections {
			if programs[name].Group != 0 {
				group = programs[name].Group
			}
		}
		if group == 0 {
			groupCount++
			program.Group = groupCount
		}

		for name := range program.Connections {
			if programs[name].Group == 0 {
				programs[name].Group = program.Group
			} else {

			}
		}
	}

	fmt.Println(groupCount)

}
func fileToLines(fileName string) []string {
	ret := make([]string, 0, 100)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		ret = append(ret, line)
	}

	return ret
}
func stringsToInts(numbers []string) []int {
	ret := make([]int, len(numbers), len(numbers))
	for i, v := range numbers {
		ret[i], _ = strconv.Atoi(v)
	}
	return ret
}
