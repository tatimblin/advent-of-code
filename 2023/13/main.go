package main

import (
	"fmt"
	"strings"

	"github.com/tatimblin/advent-of-code/go-util"
)

var SMUDGE_COUNT int
var smudges int

func main() {
	var result int

	SMUDGE_COUNT = 1

	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("could not parse input file")
		return
	}

	var matrix Matrix
	for _, line := range lines {
		if line == "" {
			vertical := matrix.scanLeft()
			horizontal := matrix.scanDown()
			result += vertical + (100 * horizontal)
			matrix = Matrix{}
		} else {
			matrix = append(matrix, strings.Split(line, ""))
		}
	}

	fmt.Println(result)
}

type Matrix [][]string

func (m *Matrix) getRow(index int) []string {
	return (*m)[index]
}

func (m *Matrix) getCol(index int) []string {
	size := len(*m)
	col := make([]string, size)
	for i := 0; i < size; i++ {
		col[i] = (*m)[i][index]
	}
	return col
}

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if smudges < SMUDGE_COUNT {
				smudges += 1
			} else {
				return false
			}
		}
	}

	return true
}

func (m *Matrix) scanDown() int {
	size := len(*m) - 1

	for i := 0; i < size; i++ {
		top := m.getRow(i)
		bottom := m.getRow(i + 1)
		smudges = 0
		if isEqual(top, bottom) {
			t := i
			b := i + 1
			for t > 0 && b < size {
				t -= 1
				b += 1

				if !isEqual(m.getRow(t), m.getRow(b)) {
					t += 1
					b -= 1
					break
				}
			}

			if (t == 0 || b == size) && smudges == SMUDGE_COUNT {
				return i + 1
			}
		}
	}

	return 0
}

func (m *Matrix) scanLeft() int {
	size := len((*m)[0]) - 1

	for i := 0; i < size; i++ {
		left := m.getCol(i)
		right := m.getCol(i + 1)
		smudges = 0
		if isEqual(left, right) {
			l := i
			r := i + 1
			for l > 0 && r < size {
				l -= 1
				r += 1

				if !isEqual(m.getCol(l), m.getCol(r)) {
					l += 1
					r -= 1
					break
				}
			}

			if (l == 0 || r == size) && smudges == SMUDGE_COUNT {
				return i + 1
			}
		}
	}

	return 0
}
