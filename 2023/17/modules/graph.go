package modules

type Matrix [][]int

type Coordinate [2]int
type Step struct {
	Coordinate Coordinate
	Direction  Direction
}

type Edge struct {
	To     int
	Weight int
	Depth  int
	Direction
}

type Graph map[int][]Edge

func CreateGraph(matrix Matrix, min int, max int) Graph {
	edges := make(map[int][]Edge)
	steps := []Step{
		{Coordinate: [2]int{0, 1}, Direction: Right},
		{Coordinate: [2]int{1, 0}, Direction: Down},
		{Coordinate: [2]int{0, -1}, Direction: Left},
		{Coordinate: [2]int{-1, 0}, Direction: Up},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fromID := i*len(matrix[0]) + j

			for _, step := range steps {
				var weight int
				for depth := min; depth < max-1; depth++ {
					ii := i + step.Coordinate[0]*depth
					jj := j + step.Coordinate[1]*depth
					if matrix.isCoordinate(Coordinate{i, j}) && matrix.isCoordinate(Coordinate{ii, jj}) {
						weight += matrix[ii][jj]
						toID := ii*len(matrix[ii]) + jj
						edges[fromID] = append(edges[fromID], Edge{
							To:        toID,
							Weight:    weight,
							Depth:     depth,
							Direction: step.Direction,
						})
					}
				}
			}
		}
	}

	return edges
}

func (matrix Matrix) isCoordinate(coord Coordinate) bool {
	maxY := len(matrix)
	maxX := len(matrix[0])
	if coord[0] < 0 || coord[0] >= maxY || coord[1] < 0 || coord[1] >= maxX {
		return false
	}
	return true
}
