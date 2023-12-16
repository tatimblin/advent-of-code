package main

import (
	"fmt"

	"github.com/tatimblin/advent-of-code/2023/14/modules"
	"github.com/tatimblin/advent-of-code/go-util"
)

func main() {
	var result int

	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("could not parse input file")
		return
	}

	columns := make([][]string, len(lines[0]))

	config := modules.Properties{
		Left:      "O",
		Right:     ".",
		Collision: "#",
	}

	for i := 0; i < len(columns); i++ {
		column := make([]string, len(lines))
		for j := 0; j < len(lines); j++ {
			column[j] = string([]rune(lines[j])[i])
		}
		sorted := config.Sort(column)
		result += config.Count(sorted)
	}

	fmt.Println(result)
}
