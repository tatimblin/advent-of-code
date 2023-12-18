package modules

import (
	"container/heap"
	"math"
)

type Matrix [][]int

type Coordinate [2]int

type Graph struct {
	Edges     map[int][]Edge
	Verticies map[int]*Vertex
	Queue     MinHeap
}

type Edge struct {
	To     int
	Weight int
}

type Vertex struct {
	ID         int
	Distance   int
	Directions []Direction
	index      int
}

func CreateGraph(matrix Matrix) Graph {
	var graph Graph
	graph.Edges = make(map[int][]Edge)
	graph.Verticies = map[int]*Vertex{}
	graph.Queue = make(MinHeap, len(matrix)*len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fromID := i*len(matrix[0]) + j

			v := &Vertex{
				ID:       fromID,
				Distance: initDistance(i, j, matrix[0][0]),
				index:    fromID,
			}

			graph.Verticies[fromID] = v
			graph.Queue[fromID] = v
			for _, adjacent := range [4]Coordinate{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				ii := i + adjacent[0]
				jj := j + adjacent[1]

				if matrix.isCoordinate(Coordinate{i, j}) && matrix.isCoordinate(Coordinate{ii, jj}) {
					toID := ii*len(matrix[ii]) + jj
					graph.addEdge(fromID, toID, matrix[ii][jj])
				}
			}
		}
	}

	heap.Init(&graph.Queue)

	return graph
}

func initDistance(i, j, distance int) int {
	if i == 0 && j == 0 {
		return distance
	}
	return math.MaxInt
}

func (matrix Matrix) isCoordinate(coord Coordinate) bool {
	maxY := len(matrix)
	maxX := len(matrix[0])
	if coord[0] < 0 || coord[0] >= maxY || coord[1] < 0 || coord[1] >= maxX {
		return false
	}
	return true
}

func (graph *Graph) addEdge(from int, to int, weight int) {
	(*graph).Edges[from] = append((*graph).Edges[from], Edge{
		To:     to,
		Weight: weight,
	})
}
