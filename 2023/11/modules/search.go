package module

import (
	"math"
)

type Galaxy struct {
	ID    int
	X     int
	Y     int
	Total int
}

type Search interface {
	CalculateDistance(from Galaxy, to Galaxy) int
	CountRange(arr []int, from int, to int) int
}

func CalculateDistance(from Galaxy, to Galaxy) int {
	xDistance := math.Abs(float64(from.X) - float64(to.X))
	yDistance := math.Abs(float64(from.Y) - float64(to.Y))
	return int(xDistance + yDistance)
}

func CountRange(arr []int, from int, to int) int {
	left := 0
	right := len(arr)
	low := min(from, to)
	high := max(from, to)

	for low > arr[left] {
		left += 1
		if left == len(arr) {
			break
		}
	}

	for right > 0 && high < arr[right-1] {
		right -= 1
	}

	return right - left
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
