package helpers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func StringToSlice(in string) []string {
	return strings.Split(in, "")
}
func SliceToString(in []string) string {
	ret := ""
	for _, s := range in {
		ret += s
	}
	return ret
}

func FileToLines(fileName string, skipEmpty bool) []string {
	ret := make([]string, 0, 100)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && skipEmpty {
			continue
		}
		ret = append(ret, line)
	}

	return ret
}

func SubString(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

//XorString returns true if exactly one of a and b equals t
func XorStrings(t, a, b string) bool {

	return (t == a || t == b) && (a != b)
}
