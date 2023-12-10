package main

import (
	"container/heap"
	"fmt"

	camelcard "github.com/tatimblin/advent-of-code/2023/07/camelcard"
	"github.com/tatimblin/advent-of-code/go-util"
)

type Item struct {
	Value    camelcard.Hand
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		return
	}

	pq := make(PriorityQueue, 0)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		hand := camelcard.Hand{}
		camelcard.ParseHand(line, &hand)
		score := hand.Score()

		heap.Push(&pq, &Item{
			Priority: score,
			Value:    hand,
		})
	}

	var part1 int
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		multiplier := len(lines) - pq.Len()
		part1 += item.Value.Bid * multiplier
		// fmt.Println(item.Value)
	}

	fmt.Println(part1)
}
