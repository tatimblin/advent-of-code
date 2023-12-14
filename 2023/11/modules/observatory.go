package module

type Universe struct {
	Matrix  [][]string
	ExpandX []int
	ExpandY []int
}

type Observatory interface {
	CreateGalaxies() []Galaxy
	Expand()
}

func (u *Universe) CreateGalaxies() []Galaxy {
	galaxies := []Galaxy{}
	var count int

	for y := range (*u).Matrix {
		for x, cell := range (*u).Matrix[y] {
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

func (u *Universe) Expand() {
	u.expandY()
	u.expandX()
}

func (u *Universe) expandX() {
	indicies := []int{}
	for i, row := range (*u).Matrix {
		if isEvery(row, ".") {
			indicies = append(indicies, i)
		}
	}
	u.ExpandX = indicies
}

func (u *Universe) expandY() {
	indicies := []int{} // slice of columns to insert at
	columnCount := len((*u).Matrix[0])
	for i := 0; i < columnCount; i++ {
		arr := make([]string, len((*u).Matrix))
		for j, row := range (*u).Matrix {
			arr[j] = row[i]
		}
		if isEvery(arr, ".") {
			indicies = append(indicies, i)
		}
	}
	u.ExpandY = indicies
}

func isEvery(arr []string, match string) bool {
	for _, char := range arr {
		if char != match {
			return false
		}
	}
	return true
}
