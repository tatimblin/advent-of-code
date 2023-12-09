package main

import "fmt"

/*
Time:      7  15   30
Distance:  9  40  200

Time:        41     96     88     94
Distance:   214   1789   1127   1055

7 9
0  1  2  3  4  5  6  7
0  6 10 12 12 10  6  0

hold * (time - hold) = distance

calc number of ways to beat distance
optimal distance is to hold button for half the time
1. get optimal distance
2. use quick search to find farthest winning hold
3. multiply distance by 2
4. handle odd number
*/

type Race struct {
	Time     int
	Distance int
}

func main() {
	// races := []Race{
	// 	{Time: 41, Distance: 214},
	// 	{Time: 96, Distance: 1789},
	// 	{Time: 88, Distance: 1127},
	// 	{Time: 94, Distance: 1055},
	// }
	// races := []Race{
	// 	{Time: 7, Distance: 9},
	// 	{Time: 15, Distance: 40},
	// 	{Time: 30, Distance: 200},
	// }
	races := []Race{
		{Time: 41968894, Distance: 214178911271055},
	}

	part1 := 1
	for _, race := range races {
		optimalHold := race.Time / 2
		// optimalDistance := CalcDistance(race.Time, optimalHold)

		min := FindHold(race, 0, optimalHold)
		max := race.Time - min
		part1 *= max - min + 1
	}

	fmt.Println(part1)
}

func CalcDistance(total, hold int) int {
	return (total - hold) * hold
}

func FindHold(race Race, left int, right int) int {
	pivot := left + (right-left)/2
	distance := CalcDistance(race.Time, pivot)

	if left >= right {
		return pivot
	}

	if race.Distance >= distance {
		return FindHold(race, pivot+1, right)
	} else {
		return FindHold(race, left, pivot)
	}
}
