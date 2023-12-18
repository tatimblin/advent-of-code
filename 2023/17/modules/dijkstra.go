package modules

import (
	"container/heap"
)

type Direction int

const (
	Unknown Direction = 0
	Left    Direction = 1
	Right   Direction = 2
	Up      Direction = 3
	Down    Direction = 4
)

func Dijkstra(graph Graph, maxStraight int) int {

	for graph.Queue.Len() > 0 {
		pointer := heap.Pop(&graph.Queue).(*Vertex)

		if pointer.ID == len(graph.Verticies)-1 {
			return pointer.Distance
		}

		for _, edge := range graph.Edges[pointer.ID] {
			distance := pointer.Distance + edge.Weight
			if distance < graph.Verticies[edge.To].Distance {
				direction := getDirection(pointer.ID, edge.To, 13)
				// if exceedsConsecutiveLimit(pointer.Directions, direction, maxStraight) {
				// 	continue
				// }

				graph.Queue.Update(
					graph.Verticies[edge.To],
					distance,
					append(pointer.Directions, direction),
				)
			}
		}
	}

	return -1
}

func getDirection(from, to, size int) Direction {
	if from+1 == to {
		return Right
	}
	if from-1 == to {
		return Left
	}
	if from+1 < to {
		return Down
	}
	if from-1 > to {
		return Up
	}
	return Unknown
}

func exceedsConsecutiveLimit(prev []Direction, new Direction, limit int) bool {
	if len(prev) < limit {
		return false
	}

	var consecutive int
	for i := len(prev) - 1; i >= 0; i-- {
		if prev[i] != new {
			break
		}
		consecutive += 1
	}

	return consecutive >= limit
}
