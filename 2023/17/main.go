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

	graph := modules.CreateGraph(matrix)
	result = modules.Dijkstra(graph, 3)

	// for _, v := range graph.Verticies {
	// 	fmt.Println(v)
	// }
	fmt.Println(graph.Verticies[168])
	// 2 2 4 2 2 2 3 2 2 2 4 4 2 2 4 4 2 4 4 4 2 4 4 4 1 4 4 2

	fmt.Println(result)
}
