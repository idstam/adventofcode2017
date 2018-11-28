package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	memory := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		word, _ := strconv.Atoi(scanner.Text())
		memory = append(memory, word)
	}

	address := 0
	ctr := 0
	for {
		ctr++
		word := memory[address]
		if word < 3 {
			memory[address] = memory[address] + 1
		} else {
			memory[address] = memory[address] - 1
		}
		address += word
		if address >= len(memory) {
			break
		}
	}

	fmt.Println(ctr)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
