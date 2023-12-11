package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tatimblin/advent-of-code/2023/09/instabilitysensor"
	"github.com/tatimblin/advent-of-code/go-util"
)

func main() {
	var part1, part2 int

	lines, err := util.ReadLines("input.txt")
	if err != nil {
		return
	}

	for _, line := range lines {
		parsed := parseLine(line)
		log := instabilitysensor.Log{
			History: [][]int{parsed},
		}
		log.CalculateDifferences()
	}

	fmt.Println(part1, part2)
}

func parseLine(line string) []int {
	strs := strings.Split(line, " ")
	result := make([]int, len(strs))
	for i, str := range strs {
		number, err := strconv.Atoi(str)
		if err != nil {
			fmt.Errorf("could not convert %s to a number", str)
		}
		result[i] = number
	}
	return result
}
