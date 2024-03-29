package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/tatimblin/advent-of-code/go-util"
)

type Conversion struct {
	Ranges [][]int // TODO: would be better sorted
	Offset []int
}

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("Error: Could not read lines")
		return
	}

	var almanac []Conversion
	var current Conversion

	for i := 3; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			almanac = append(almanac, current)
			current = Conversion{}
			i += 1
			continue
		}

		var instructions []int
		for _, _instruction := range strings.Split(line, " ") {
			instruction, err := strconv.Atoi(_instruction)
			if err != nil {
				fmt.Println("could not convert instruction:", _instruction)
				break
			}
			instructions = append(instructions, instruction)
		}

		current.Ranges = append(current.Ranges, []int{instructions[1], instructions[1] + (instructions[2] - 1)})
		current.Offset = append(current.Offset, instructions[0]-instructions[1])
	}
	almanac = append(almanac, current)

	seeds := GetSeeds(lines[0])

	part1 := math.MaxInt
	// for _, seed := range seeds {
	// 	fmt.Println(seed)
	// 	location := CalculateSoil(seed, almanac)
	// 	if location < int(part1) {
	// 		part1 = location
	// 	}
	// }

	// lowestSoil := math.MaxInt
	// for _, seed := range GetSeeds(lines[0]) {
	// 	fmt.Println(seed)
	// 	soil := CalculateSoil(seed, almanac)
	// 	fmt.Println(soil)
	// 	if soil < int(lowestSoil) {
	// 		lowestSoil = soil
	// 	}
	// }

	part2 := math.MaxInt
	for i := 0; i < len(seeds)-1; i += 2 {
		pointer := seeds[i]
		end := seeds[i+1]
		for pointer < seeds[i]+end {
			location := CalculateSoil(pointer, almanac)
			if location < int(part2) {
				part2 = location
			}
			pointer += 1
		}
	}

	fmt.Println(part1, part2)
}

func GetSeeds(input string) []int {
	input = strings.Replace(input, "seeds: ", "", 1)
	var seeds []int
	for _, s := range strings.Split(input, " ") {
		seed, err := strconv.Atoi(s)
		if err != nil {
			seeds = append(seeds, 0)
		} else {
			seeds = append(seeds, seed)
		}
	}
	return seeds
}

func CalculateSoil(seed int, almanac []Conversion) int {
	for _, conversion := range almanac {
		for i, r := range conversion.Ranges {
			if seed >= r[0] && seed <= r[1] {
				seed += conversion.Offset[i]
				break
			}
		}
	}

	return seed
}
