package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tatimblin/advent-of-code/go-util"
)

type Scratcher struct {
	ID       int
	Winning  []int
	Matching []int
	Points   int
}

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("Error: Could not read lines")
		return
	}

	var totalPoints int
	scratchers := make([]Scratcher, len(lines))
	for i, line := range lines {
		scratchers[i] = ParseScratcher(line)
		totalPoints += scratchers[i].Points
	}

	fmt.Println(totalPoints)
}

func ParseScratcher(input string) Scratcher {
	scratcher := Scratcher{}

	card := strings.Split(input, ":")

	if len(card) > 0 {
		idString := strings.Replace(card[0], "Card ", "", 1)
		id, err := strconv.Atoi(idString)
		if err == nil {
			scratcher.ID = id
		}
	}

	if len(card) > 1 {
		card[1] = strings.Replace(card[1], "  ", " ", -1)
		groups := strings.Split(card[1], "|")
		scratcher.Winning = setIntSlice(groups[0])
		scratcher.Matching = setIntSlice(groups[1])
		intersect := getIntersect(scratcher.Winning, scratcher.Matching)
		fmt.Println(scratcher, len(intersect))
		scratcher.Points = getPoints(len(intersect))
	}

	return scratcher
}

func setIntSlice(data string) []int {
	data = strings.Trim(data, " ")
	numbers := strings.Split(data, " ")
	result := make([]int, len(numbers))

	for i, num := range numbers {
		number, err := strconv.Atoi(num)
		if err == nil {
			result[i] = number
		}
	}

	return result
}

func getIntersect(a, b []int) []int {
	var intersect []int
	offset := 0

	// TODO
	for _, numA := range a {
		for j := 0; j < len(b)-offset; j++ {
			numB := b[j]
			if numA == numB {
				intersect = append(intersect, numA)
				break
			}
		}
	}

	return intersect
}

func getPoints(matches int) int {
	if matches == 0 {
		return 0
	}

	points := 1
	for i := 1; i < matches; i++ {
		points *= 2
	}

	return points
}
