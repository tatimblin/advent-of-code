package modules

import (
	"container/heap"
)

type MinHeap []*Vertex

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h MinHeap) Swap(i, j int) {
	if i >= 0 && i < len(h) && j >= 0 && j < len(h) {
		h[i], h[j] = h[j], h[i]
		h[i].index = i
		h[j].index = j
	}
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Vertex))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	x.index = -1
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Update(item *Vertex, distance int, directions []Direction) {
	item.Directions = directions
	item.Distance = distance
	if item.index == -1 {
		return
	}
	heap.Fix(h, item.index)
}
