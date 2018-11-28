package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := make(map[string]int)

		valid := true

		words := strings.Split(scanner.Text(), " ")
		for _, word := range words {
			trimmedWord := strings.TrimSpace(word)
			if m[trimmedWord] == 0 {
				m[trimmedWord] = 1
				continue
			} else {
				valid = false
				break
			}

		}
		if valid {
			count++
		}

	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
