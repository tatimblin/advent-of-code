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

	min, max := 3, 10
	start, end := 0, len(matrix)*len(matrix[0])-1

	graph := modules.CreateGraph(matrix, min, max)
	result = modules.Dijkstra(graph, max, start, end)

	fmt.Println(result)
}
