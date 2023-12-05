package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/tatimblin/advent-of-code/go-util"
)

type Piece struct {
	Read   string
	Number int
	End    bool
	Start  bool
	Valid  bool
}

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("Error: Could not read lines")
		return
	}

	fmt.Println(part1(lines))
}

func part1(lines []string) int {
	var result int
	var activePiece Piece

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			char := line[j]

			if strings.Contains("!@#$%^&*()/\\-+=.", string(char)) && activePiece.Start {
				closePiece(&activePiece)
				if activePiece.Valid {
					// pieces = append(pieces, activePiece)
					result += activePiece.Number
				}
				activePiece = Piece{}
			}

			if unicode.IsNumber(rune(char)) {
				updatePiece(&activePiece, string(char))
				instructions := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
				for _, instruction := range instructions {
					y := i + instruction[0]
					x := j + instruction[1]

					if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[y]) {
						continue
					}

					if strings.Contains("!@#$%^&*()/\\-+=", string(lines[y][x])) {
						activePiece.Valid = true
						break
					}
				}
			}
		}
	}

	return result
}

func closePiece(piece *Piece) {
	piece.End = true

	number, err := strconv.Atoi(piece.Read)
	if err == nil {
		piece.Number = number
	}
}

func updatePiece(piece *Piece, digit string) {
	piece.Read = piece.Read + digit
	piece.Start = true
}
