package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//bgvyzdsv

	i := 1
	for {
		data := []byte("bgvyzdsv" + strconv.Itoa(i))
		check := md5.Sum(data)
		//if check[0] == 0 && check[1] == 0 && check[2] < 16 {
		if check[0] == 0 && check[1] == 0 && check[2] == 0 {
			fmt.Println(i)
			return
		}
		i++

	}
}

func minInt(numbers ...int) int {
	m := numbers[0]
	for _, n := range numbers {
		if n < m {
			m = n
		}
	}

	return m
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
		if line != "" {
			ret = append(ret, line)
		}
	}

	return ret
}
