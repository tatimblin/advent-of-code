package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tatimblin/advent-of-code/go-util"
)

type Round struct {
	ID    int
	Red   int
	Blue  int
	Green int
}

type Game struct {
	ID     int
	Red    int
	Blue   int
	Green  int
	Rounds *[]Round
}

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("Error: Could not read lines")
		return
	}

	games, _ := ParseGameLogs(lines)

	var part1 int
	var part2 int
	for _, game := range games {
		id := game.ID
		possible := true
		part2 += game.Red * game.Blue * game.Green
		for _, round := range *game.Rounds {
			if !isPossible(round) {
				possible = false
			}
		}

		if possible {
			part1 += id
		}
	}

	fmt.Println(part1, part2)
}

func ParseGameLogs(logs []string) ([]Game, []Round) {
	games := make([]Game, len(logs))
	var rounds []Round

	for _, gameLog := range logs {

		split := strings.Split(gameLog, ":")
		if len(split) != 2 {
			log.Printf("Had trouble parsing log: %s", gameLog)
			return games, rounds
		}

		id := getID(split[0])
		if id == -1 {
			log.Printf("Invalid Game.ID: %d", id)
			return games, rounds
		}

		game := Game{
			ID:     id,
			Rounds: &[]Round{},
		}

		gameRounds := getRounds(split[1])
		var red, blue, green int
		for _, round := range gameRounds {
			round.ID = id
			*game.Rounds = append(*game.Rounds, round)

			if round.Red > red {
				red = round.Red
			}
			if round.Blue > blue {
				blue = round.Blue
			}
			if round.Green > green {
				green = round.Green
			}

			rounds = append(rounds, round)
		}

		game.Red = red
		game.Blue = blue
		game.Green = green

		games[id-1] = game
	}

	return games, rounds
}

func getID(gameID string) int {
	idString := strings.Replace(gameID, "Game ", "", 1)
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Invalid Game.ID: %s", idString)
		return -1
	}

	return id
}

func getRounds(roundsString string) []Round {
	var rounds []Round
	roundStrings := strings.Split(roundsString, ";")
	for _, roundString := range roundStrings {
		var round Round
		scores := strings.Split(roundString, ",")
		for _, score := range scores {
			score = strings.Trim(score, " ")
			stat := strings.Split(score, " ")
			setCount(&round, stat)
		}
		rounds = append(rounds, round)
	}

	return rounds
}

func setCount(round *Round, stat []string) {
	count, err := strconv.Atoi(stat[0])
	if err != nil {
		return
	}

	switch stat[1] {
	case "red":
		round.Red = count
	case "blue":
		round.Blue = count
	case "green":
		round.Green = count
	}
}

func isPossible(round Round) bool {
	if round.Red <= 12 && round.Green <= 13 && round.Blue <= 14 {
		return true
	}
	return false
}
