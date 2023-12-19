package modules

import (
	"container/heap"
	"fmt"
)

type Direction int

const (
	Unknown Direction = 0
	Left    Direction = 1
	Right   Direction = 2
	Up      Direction = 3
	Down    Direction = 4
)

type Vertex struct {
	ID        int
	Distance  int
	Direction Direction
	SinceTurn int
	index     int
}

type MinHeap []*Vertex

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *MinHeap) Push(x interface{}) {
	item := x.(*Vertex)
	item.index = len(*h)
	*h = append(*h, item)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	x.index = -1
	*h = old[0 : n-1]
	return x
}

func Dijkstra(graph Graph, source *Vertex, end *Vertex) int {
	distances := make(map[int]int)

	for vertex := range graph {
		distances[vertex] = int(^uint(0) >> 1)
	}
	distances[source.ID] = source.Distance

	pq := &MinHeap{source}
	heap.Init(pq)

	for pq.Len() > 0 {
		currentVertex := heap.Pop(pq).(*Vertex)

		for _, edge := range graph[currentVertex.ID] {
			neighbor := edge.To
			newDistance := currentVertex.Distance + edge.Weight

			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance

				sinceTurn := currentVertex.SinceTurn + edge.Depth

				if (currentVertex.Direction == Up && edge.Direction == Down) ||
					(currentVertex.Direction == Down && edge.Direction == Up) ||
					(currentVertex.Direction == Left && edge.Direction == Right) ||
					(currentVertex.Direction == Right && edge.Direction == Left) {
					continue
				}

				if currentVertex.Direction != edge.Direction && currentVertex.Direction != Unknown {
					sinceTurn = 0
				}

				if sinceTurn > 3 {
					continue
				}

				heap.Push(pq, &Vertex{
					ID:        neighbor,
					Distance:  newDistance,
					Direction: edge.Direction,
					SinceTurn: sinceTurn,
				})
			}
		}
	}

	fmt.Println(distances)

	return distances[end.ID]
}
