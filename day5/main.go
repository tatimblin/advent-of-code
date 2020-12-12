package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Range struct {
	Min int
	Max int
}

type Position struct {
	Row    int
	Column int
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	highSeatID := 0
	takenSeats := make(map[int]bool)
	openSeats := []int{}

	for _, line := range lines {
		intstructions := strings.Split(line, "")
		seatPos := readSeatInst(intstructions)
		seatID := getSeatID(seatPos)

		if seatID > highSeatID {
			highSeatID = seatID
		}

		takenSeats[seatID] = true
	}

	for i := 1; i <= highSeatID; i++ {
		var (
			above = i - 1
			below = i + 1
		)

		if !takenSeats[i] && takenSeats[above] && takenSeats[below] {
			openSeats = append(openSeats, i)
		}
	}

	log.Println(openSeats)
}

func readSeatInst(instructions []string) Position {
	var (
		rowInst  = instructions[:7]
		rowRange = Range{
			Min: 0,
			Max: 127,
		}
		colInst  = instructions[7:]
		colRange = Range{
			Min: 0,
			Max: 7,
		}
	)

	row := departition(rowInst, rowRange, "F", "B")
	col := departition(colInst, colRange, "L", "R")

	return Position{
		Row:    row,
		Column: col,
	}
}

func departition(instruction []string, spread Range, lowKey, highKey string) int {
	currentSpread := spread
	result := 0
	for _, step := range instruction {
		difference := currentSpread.Max - currentSpread.Min + 1

		if step == lowKey || step == highKey {
			goUp := true
			if step == lowKey {
				goUp = false
			}

			// Last is special
			if currentSpread.Max-currentSpread.Min == 1 {
				result = currentSpread.Min
				if goUp {
					result = currentSpread.Max
				}
			}

			if goUp {
				currentSpread.Min = currentSpread.Min + difference/2
			} else {
				currentSpread.Max = currentSpread.Max - difference/2
			}
		}
	}

	return result
}

func getSeatID(seat Position) int {
	return seat.Row*8 + seat.Column
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
