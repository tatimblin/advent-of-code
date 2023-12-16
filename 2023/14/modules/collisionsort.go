package modules

type Properties struct {
	Left      string
	Right     string
	Collision string
}

type CollisionSort interface {
	Sort([]string) ([]string, int)
}

func (config *Properties) Sort(input []string) []string {
	for i := 0; i < len(input); i++ {
		var score int
		if input[i] == config.Left {
			toSwap := i
			j := i - 1
			for j >= 0 {
				if input[j] == config.Collision {
					break
				}
				if input[j] == config.Right {
					toSwap = j
				}
				j -= 1
			}
			input[i], input[toSwap] = input[toSwap], input[i]
			score += 1
		}
	}
	return input
}

func (config *Properties) Count(input []string) int {
	var result int
	for i, cell := range input {
		if cell == config.Left {
			result += len(input) - i
		}
	}
	return result
}
