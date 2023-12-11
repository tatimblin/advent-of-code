package instabilitysensor

type Log struct {
	History [][]int
}

type InstabilitySensor interface {
	CalculateDifferences()
	AddEstimates() int
}

func (log *Log) CalculateDifferences() {
	var isZeros bool
	for !isZeros {
		isZeros = true
		size := len(log.History)
		differences := []int{}
		for i := 0; i < len(log.History[size-1])-1; i++ {
			difference := log.History[size-1][i+1] - log.History[size-1][i]
			differences = append(differences, difference)
			if difference != 0 {
				isZeros = false
			}
		}
		log.History = append(log.History, differences)
	}
}

func (log *Log) AddEstimates() int {
	var est int
	size := len(log.History)
	for i := size - 1; i >= 0; i-- {
		if i != size-1 {
			left := log.History[i][len(log.History[i])-1]
			down := log.History[i+1][len(log.History[i+1])-1]
			est = left + down
		}
		log.History[i] = append(log.History[i], est)
	}

	return est
}

func (log *Log) AddEstimatesLeft() int {
	var est int
	size := len(log.History)
	for i := size - 1; i >= 0; i-- {
		if i != size-1 {
			right := log.History[i][0]
			down := log.History[i+1][0]
			est = right - down
		}
		log.History[i] = append([]int{est}, log.History[i]...)
	}

	return est
}
