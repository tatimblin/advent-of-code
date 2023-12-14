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
		universe.Matrix = append(universe.Matrix, chars)
	}

	universe.Expand()
	galaxies := universe.CreateGalaxies()

	for offset, fromGalaxy := range galaxies {
		for i := offset + 1; i < len(galaxies); i++ {
			toGalaxy := galaxies[i]
			distance := module.CalculateDistance(fromGalaxy, toGalaxy)
			expansions := module.CountRange(universe.ExpandY, fromGalaxy.X, toGalaxy.X) + module.CountRange(universe.ExpandX, fromGalaxy.Y, toGalaxy.Y)
			part1 += distance + expansions
			part2 += distance + (expansions * 999999)
		}
	}

	fmt.Println(part1, part2)
}
