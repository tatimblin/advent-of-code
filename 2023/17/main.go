package main

import (
	"fmt"

	"github.com/tatimblin/advent-of-code/2023/17/modules"
	"github.com/tatimblin/advent-of-code/go-util"
)

func main() {
	var result int

	// backtrack from answer calculating lowest heatloss from point
	// track where direction changes happen and increment them

	matrix, err := util.ReadMatrixInt("input.mock.txt", "")
	if err != nil {
		fmt.Println("could not parse input file")
		return
	}

	end := len(matrix)*len(matrix[0]) - 1
	graph := modules.CreateGraph(matrix, 1, 3)
	for _, g := range graph {
		fmt.Println(g)
	}
	result = modules.Dijkstra(
		graph,
		&modules.Vertex{
			ID:        0,
			Distance:  matrix[0][0],
			SinceTurn: 1,
			Direction: modules.Unknown,
		},
		&modules.Vertex{
			ID: end,
		},
	)

	fmt.Println(result)
}
