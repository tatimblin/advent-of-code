package main

import (
	"fmt"
	"strings"

	module "github.com/tatimblin/advent-of-code/2023/11/modules"
	"github.com/tatimblin/advent-of-code/go-util"
)

func main() {
	var part1, part2 int

	// get a list of galaxies
	// bfs off of all of them
	// update Galaxy.Pair whenever one is found
	// count total pairs and escape when reached

	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("could not parse input file")
		return
	}

	universe := module.Universe{}
	for _, line := range lines {
		chars := strings.Split(line, "")
		universe = append(universe, chars)
	}

	universe.Expand() // does not scale for part 2 :(
	galaxies := universe.CreateGalaxies()

	for i := 0; i < len(galaxies); i++ {
		universe.BFS(galaxies[i])
	}

	fmt.Println(part1, part2)
}
