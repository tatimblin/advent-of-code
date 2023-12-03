package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tatimblin/advent-of-code/go-util"
)

type Match struct {
	Matcher string
	Replace string
}

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("Error: Could not read lines")
		return
	}

	var sum int
	matches := []Match{
		{ Matcher: "zero", Replace: "0" },
		{ Matcher: "one", Replace: "1" },
		{ Matcher: "two", Replace: "2" },
		{ Matcher: "three", Replace: "3" },
		{ Matcher: "four", Replace: "4" },
		{ Matcher: "five", Replace: "5" },
		{ Matcher: "six", Replace: "6" },
		{ Matcher: "seven", Replace: "7" },
		{ Matcher: "eight", Replace: "8" },
		{ Matcher: "nine", Replace: "9" },
	}

	for _, line := range lines {
		ReplaceMatches(&line, matches)
		sum += ParseString(line)
	}

	fmt.Println(sum)
}

func ReplaceMatches(str *string, matches []Match) {
	for _, match := range matches {
		merge := fmt.Sprintf("%s%s%s", match.Matcher, match.Replace, match.Matcher)
		*str = strings.Replace(*str, match.Matcher, merge, -1)
	}
}

func ParseString(str string) int {
	var left byte = '0'
	var right byte = '0'

	for i := 0; i < len(str); i++ {
		if IsByteNumber(str[i]) {
			left = str[i]
			break
		}
	}

	for i := len(str) - 1; i >= 0; i-- {
		if IsByteNumber(str[i]) {
			right = str[i]
			break
		}
	}

	concat, err := strconv.Atoi(fmt.Sprintf("%c%c", left, right))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	return concat
}

func IsByteNumber(b byte) bool {
	_, err := strconv.Atoi(string(b))
	if err != nil {
		return false
	}
	return true
}