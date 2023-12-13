package module

import "fmt"

type Universe [][]string

type Observatory interface {
	Expand()
}

func (u *Universe) Expand() {
	u.expandY()
	u.expandX()
}

func (u *Universe) expandX() {
	for i, row := range *u {
		if isEvery(row, ".") {
			u.insertArrayAt(row, i)
		}
	}
}

func (u *Universe) expandY() {
	indicies := []int{} // slice of columns to insert at
	columnCount := len((*u)[0])
	for i := 0; i < columnCount; i++ {
		arr := make([]string, len(*u))
		for j, row := range *u {
			arr[j] = row[i]
		}
		if isEvery(arr, ".") {
			indicies = append(indicies, i)
		}
	}
	for j := range *u { // for each row add strings
		u.insertStringsAt(".", j, indicies)
	}
}

func (u *Universe) insertArrayAt(item []string, y int) {
	if y < 0 || y > len(*u) {
		fmt.Println("out of range")
		return
	}

	result := make([][]string, len(*u)+1)
	copy(result[:y], (*u)[:y])
	result[y] = item
	copy(result[y+1:], (*u)[y:])
	*u = result
}

func (u *Universe) insertStringsAt(item string, y int, x []int) {
	if y < 0 || y > len(*u) {
		fmt.Println("out of range")
		return
	}

	length := len((*u)[y]) + len(x)
	var xIndex int
	result := make([]string, length)
	for i := 0; i < length; i++ {
		if xIndex < len(x) && i == x[xIndex] {
			xIndex += 1
			result[i] = item
		} else {
			result[i] = (*u)[y][i-xIndex]
		}
	}
	(*u)[y] = result
}

func isEvery(arr []string, match string) bool {
	for _, char := range arr {
		if char != match {
			return false
		}
	}
	return true
}
