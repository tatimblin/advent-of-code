package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/tatimblin/advent-of-code/2023/10/day10"
	"github.com/tatimblin/advent-of-code/go-util"
)

func main() {
	var part1, part2 int

	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("could not parse input file")
		return
	}

	matrix := day10.Matrix{
		Pipe:     [][]string{},
		Distance: [][]int{},
	}

	for i, line := range lines {
		chars := strings.Split(line, "")
		matrix.Pipe = append(matrix.Pipe, chars)
		matrix.Distance = append(matrix.Distance, make([]int, len(chars)))
		sIndex := slices.Index(chars, "S")
		if sIndex >= 0 {
			matrix.Queue = [][3]int{{i, sIndex, 1}}
		}
	}

	for len(matrix.Queue) > 0 {
		matrix.Step()
	}

	part1 = matrix.MaxDistance

	fmt.Println(part1, part2)
}
