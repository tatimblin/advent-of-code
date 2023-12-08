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
	races := []Race{
		{Time: 7, Distance: 9},
		{Time: 15, Distance: 40},
		{Time: 30, Distance: 200},
	}

	for _, race := range races {
		optimalHold := race.Time / 2
		// optimalDistance := CalcDistance(race.Time, optimalHold)

		nearestHold := FindHold(race, optimalHold/2, optimalHold)
		// possibilities := nearestHold * 2
		fmt.Println(nearestHold, nearestHold+optimalHold)
	}

	fmt.Println(races, races[0].Time%2)
}

func CalcDistance(total, hold int) int {
	return (total - hold) * hold
}

func FindHold(race Race, pivot int, size int) int {
	distance := CalcDistance(race.Time, pivot)

	if size == 0 {
		return pivot
	}

	if race.Distance > distance {
		return FindHold(race, pivot+(size/2), size/2)
	} else {
		return FindHold(race, pivot-(size/2), size/2)
	}
}
