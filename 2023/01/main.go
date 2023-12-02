package main

import (
	"fmt"
	"strconv"

	"github.com/tatimblin/advent-of-code/go-util"
)

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("Error: Could not read lines")
		return
	}

	var sum int
	for _, line := range lines {
		sum += ParseString(line)
	}

	fmt.Println(sum)
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