package module

import "fmt"

type Galaxy struct {
	ID    int
	X     int
	Y     int
	Total int
}

type Search interface {
	Traverse()
	CreateGalaxies()
}

func (u *Universe) Traverse(galaxy Galaxy) {
	queue := [][3]int{{galaxy.Y, galaxy.X, 0}}

	func bfs() {
		coords, q := dequeue(queue)
		queue = q

		for _, offset := range [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			y := coords[0] + offset[0]
			x := coords[1] + offset[1]
			enqueue(queue, [3]int{y, x, coords[2]+1})
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

func (u *Universe) CreateGalaxies() []Galaxy {
	galaxies := []Galaxy{}
	var count int

	for y := range *u {
		for x, cell := range (*u)[y] {
			if cell == "#" {
				galaxies = append(galaxies, Galaxy{
					ID: count,
					X:  x,
					Y:  y,
				})
				count += 1
			}
		}
	}

	return galaxies
}
