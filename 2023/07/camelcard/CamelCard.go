package camelcard

import (
	"fmt"
	"strconv"
	"strings"
)

type Hand struct {
	Cards []string
	Bid   int
}

func ParseHand(str string, hand *Hand) error {
	parts := strings.Split(str, " ")

	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("could not parse bid for hand %s", parts[0])
	}
	hand.Bid = bid

	hand.Cards = strings.Split(parts[0], "")

	return nil
}

type CamelCard interface {
	Score() int
}

func (hand Hand) Score() int {
	cardType := getType(hand.Cards)
	cardValues := parseCards(hand.Cards)
	score, err := strconv.Atoi(fmt.Sprintf("%d%s", cardType, cardValues))
	if err != nil {
		fmt.Errorf("failed to convert score to int")
	}
	return score
}

func getType(cards []string) int {
	history := make(map[string]int)

	for _, card := range cards {
		history[card] += 1
	}

	var duplicates [5]int

	for card, count := range history {
		if card != "J" {
			duplicates[count-1] += 1
		}
	}

	var result int
	if duplicates[4] == 1 {
		// five of a kind
		result += 6
	} else if duplicates[3] == 1 {
		// four of a kind
		result += 5
	} else if duplicates[2] == 1 && duplicates[1] == 1 {
		// full house
		result += 4
	} else if duplicates[2] == 1 {
		// three of a kind
		result += 3
		if history["J"] > 0 {
			result += 1
		}
	} else if duplicates[1] == 2 {
		// two pair
		result += 2
		if history["J"] > 0 {
			result += 1
		}
	} else if duplicates[1] == 1 {
		// one pair
		result += 1
		if history["J"] > 1 {
			result += 2
		} else if history["J"] > 0 {
			result += 1
		}
	} else {
		if history["J"] < 5 && history["J"] > 2 {
			result += 2
		} else if history["J"] > 1 {
			result += 1
		}
	}

	result += history["J"]
	if history["J"] > 2 {
		fmt.Println(result, cards)
	}

	// high card
	return result
}

func parseCards(cards []string) string {
	suits := map[string]string{
		"T": "10",
		"J": "01", // part1: 11
		"Q": "12",
		"K": "13",
		"A": "14",
	}
	var result string

	for _, card := range cards {

		suitValue, ok := suits[card]
		if ok {
			card = suitValue
		} else {
			card = fmt.Sprintf("0%s", card)
		}

		result = fmt.Sprintf("%s%s", result, card)
	}

	return result
}
