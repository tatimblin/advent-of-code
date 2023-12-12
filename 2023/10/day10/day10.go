package day10

import (
	"slices"
)

type Matrix struct {
	Pipe        [][]string
	Distance    [][]int
	Queue       [][3]int
	MaxDistance int
}

type Traverse interface {
	Step()
}

func (matrix *Matrix) Step() {
	coords, queue := dequeue(matrix.Queue)
	matrix.Queue = queue
	pipe := matrix.Pipe[coords[0]][coords[1]]

	movement := map[string][][2]int{
		"|": {{1, 0}, {-1, 0}},
		"-": {{0, 1}, {0, -1}},
		"L": {{0, 1}, {-1, 0}},
		"J": {{0, -1}, {-1, 0}},
		"7": {{1, 0}, {0, -1}},
		"F": {{0, 1}, {1, 0}},
		".": {},
		"S": {{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
	}

	for _, offset := range movement[pipe] {
		y := coords[0] + offset[0]
		x := coords[1] + offset[1]

		if y < 0 || y > len(matrix.Pipe) || x < 0 || x > len(matrix.Pipe[0]) {
			continue
		}

		if matrix.Distance[y][x] != 0 {
			continue
		}

		toPipe := matrix.Pipe[y][x]
		pipeMatchIndex := slices.Index(movement[toPipe], invertMovement(offset))
		if pipeMatchIndex == -1 {
			continue
		}

		matrix.Distance[y][x] = coords[2]
		matrix.Queue = enqueue(matrix.Queue, [3]int{y, x, coords[2] + 1})

		if coords[2] > matrix.MaxDistance {
			matrix.MaxDistance = coords[2]
		}
	}
}

func enqueue(queue [][3]int, element [3]int) [][3]int {
	queue = append(queue, element)
	return queue
}

func dequeue(queue [][3]int) ([3]int, [][3]int) {
	element := queue[0]
	if len(queue) == 1 {
		temp := [][3]int{}
		return element, temp
	}
	return element, queue[1:]
}

func invertMovement(move [2]int) [2]int {
	if move[0] != 0 {
		move[0] = move[0] * -1
	}
	if move[1] != 0 {
		move[1] = move[1] * -1
	}

	return move
}
