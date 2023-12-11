package instabilitysensor

import (
	"fmt"
)

type Log struct {
	History [][]int
}

type InstabilitySensor interface {
	CalculateDifferences([]int) []int
	AddEstimate()
}

func (log *Log) CalculateDifferences() {
	fmt.Println(log)
}

func (log *Log) AddEstimate() {
	fmt.Println(log)
}
