package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tatimblin/advent-of-code/go-util"
)

type Permutation struct {
	group  int
	amount int
	count  int
}

func main() {
	var part1, part2 int

	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("could not parse input file")
		return
	}

	for _, line := range lines {
		records := strings.Split(line, " ")
		row := strings.Split(records[0], "")
		pattern := []int{}
		for _, char := range strings.Split(records[1], ",") {
			pat, err := strconv.Atoi(char)
			if err != nil {
				fmt.Println("coud not parse pattern")
				return
			}
			pattern = append(pattern, pat)
		}

		// ugly part 2 stuff
		bigRow := []string{}
		bigPattern := []int{}
		i := 0
		for i < 5 {
			i++
			bigRow = append(bigRow, row...)
			if i < 5 {
				bigRow = append(bigRow, "?")
			}
			bigPattern = append(bigPattern, pattern...)
		}

		part1 += countPermutations(row, pattern)
		part2 += countPermutations(bigRow, bigPattern)
	}

	fmt.Println(part1, part2)
}

func countPermutations(row []string, pattern []int) int {
	permutations := map[string]int{
		"0,0": 1,
	}

	for _, char := range row {
		queue := []Permutation{}
		for key, count := range permutations {
			group, amount := parseKey(key)
			permutation := Permutation{
				group:  group,
				amount: amount,
				count:  count,
			}
			if char != "#" { // is . or ?
				if permutation.amount == 0 {
					queue = append(queue, permutation)
				} else if permutation.amount == pattern[permutation.group] {
					queue = append(queue, Permutation{
						group:  permutation.group + 1,
						amount: 0,
						count:  permutation.count,
					})
				}
			}
			if char != "." { // is # or ?
				if permutation.group < len(pattern) && permutation.amount < pattern[permutation.group] {
					queue = append(queue, Permutation{
						group:  permutation.group,
						amount: permutation.amount + 1,
						count:  permutation.count,
					})
				}
			}
		}
		permutations = map[string]int{}
		for _, permutation := range queue {
			permutations[fmt.Sprintf("%d,%d", permutation.group, permutation.amount)] += permutation.count
		}
	}

	var count int
	for key, permpermutation := range permutations {
		group, amount := parseKey(key)
		if group == len(pattern) || (group == len(pattern)-1 && amount == pattern[len(pattern)-1]) {
			count += permpermutation
		}
	}

	return count
}

func parseKey(key string) (int, int) {
	ctx := strings.Split(key, ",")
	group, _ := strconv.Atoi(ctx[0])
	amount, _ := strconv.Atoi(ctx[1])
	return group, amount
}
