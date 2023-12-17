package main

import (
	"fmt"

	"github.com/tatimblin/advent-of-code/go-util"
)

type Matrix struct {
	Instructions [][]string
	Left         [][]bool
	Right        [][]bool
	Up           [][]bool
	Down         [][]bool
}

var direction map[string][2]int
var result int

func main() {

	direction = map[string][2]int{
		"Left":  {0, -1},
		"Down":  {1, 0},
		"Right": {0, 1},
		"Up":    {-1, 0},
	}

	var matrix Matrix
	input, err := util.ReadMatrix("input.txt")
	if err != nil {
		fmt.Println("could not parse input file")
		return
	}

	matrix.Instructions = input
	matrix.Left = make([][]bool, len(input))
	matrix.Right = make([][]bool, len(input))
	matrix.Up = make([][]bool, len(input))
	matrix.Down = make([][]bool, len(input))
	for i := 0; i < len(input); i++ {
		matrix.Left[i] = make([]bool, len(input[0]))
		matrix.Right[i] = make([]bool, len(input[0]))
		matrix.Up[i] = make([]bool, len(input[0]))
		matrix.Down[i] = make([]bool, len(input[0]))
	}

	matrix.traverse([2]int{109, 0}, "Right")

	for i := 0; i < len(matrix.Left); i++ {
		for j := 0; j < len(matrix.Left); j++ {
			if matrix.Left[i][j] || matrix.Right[i][j] || matrix.Up[i][j] || matrix.Down[i][j] {
				result += 1
			}
		}
	}

	fmt.Println(result)
}

func (m Matrix) badRange(coord [2]int) bool {
	if coord[0] < 0 || coord[0] >= len(m.Instructions) || coord[1] < 0 || coord[1] >= len(m.Instructions[0]) {
		return true
	}
	return false
}

func (m Matrix) traverse(coord [2]int, dir string) {
	if m.badRange(coord) {
		return
	}
	if (dir == "Left" && m.Left[coord[0]][coord[1]]) ||
		(dir == "Right" && m.Right[coord[0]][coord[1]]) ||
		(dir == "Up" && m.Up[coord[0]][coord[1]]) ||
		(dir == "Down" && m.Down[coord[0]][coord[1]]) {
		return
	}

	if dir == "Left" && !m.Left[coord[0]][coord[1]] {
		m.Left[coord[0]][coord[1]] = true
	}
	if dir == "Right" && !m.Right[coord[0]][coord[1]] {
		m.Right[coord[0]][coord[1]] = true
	}
	if dir == "Up" && !m.Up[coord[0]][coord[1]] {
		m.Up[coord[0]][coord[1]] = true
	}
	if dir == "Down" && !m.Down[coord[0]][coord[1]] {
		m.Down[coord[0]][coord[1]] = true
	}
	cell := m.Instructions[coord[0]][coord[1]]

	switch cell {
	case "|":
		if dir != "Down" {
			m.traverse([2]int{
				coord[0] + direction["Up"][0],
				coord[1] + direction["Up"][1],
			}, "Up")
		}
		if dir != "Up" {
			m.traverse([2]int{
				coord[0] + direction["Down"][0],
				coord[1] + direction["Down"][1],
			}, "Down")
		}
		return
	case "-":
		if dir != "Right" {
			m.traverse([2]int{
				coord[0] + direction["Left"][0],
				coord[1] + direction["Left"][1],
			}, "Left")
		}
		if dir != "Left" {
			m.traverse([2]int{
				coord[0] + direction["Right"][0],
				coord[1] + direction["Right"][1],
			}, "Right")
		}
		return
	case "/":
		if dir == "Left" {
			dir = "Down"
		} else if dir == "Up" {
			dir = "Right"
		} else if dir == "Right" {
			dir = "Up"
		} else if dir == "Down" {
			dir = "Left"
		}
	case "\\":
		if dir == "Left" {
			dir = "Up"
		} else if dir == "Up" {
			dir = "Left"
		} else if dir == "Right" {
			dir = "Down"
		} else if dir == "Down" {
			dir = "Right"
		}
	}

	m.traverse([2]int{
		coord[0] + direction[dir][0],
		coord[1] + direction[dir][1],
	}, dir)
}
