package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Color struct {
	Amount int
	Name   string
}

/*
[
	['poshcoral']: [
		[
			Amount: 0,
			Color: 'green'
		],
		[
			Amount: 0,
			Color: 'green'
		]
	]
]
*/

func main() {
	rules := readLines("input.txt")

	// Part 1
	contains := make(map[string]struct{})
	check("shiny gold", rules, contains)
	log.Println(len(contains))

	// Part 2
	log.Println(getAmount("shiny gold", rules))
}

func readLines(path string) map[string][]Color {
	file, _ := os.Open(path)
	defer file.Close()

	rules := make(map[string][]Color)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")

		mainColor := strings.Join(words[:2], " ")
		contains := make([]Color, 0)

		i := 5
		for i < len(words) {
			amt, err := strconv.Atoi(words[i-1])
			if err != nil {
				amt = 0
			}

			rule := Color{
				Amount: amt,
				Name:   strings.Join(words[i:i+2], " "),
			}

			contains = append(contains, rule)

			i = i + 4
		}

		rules[mainColor] = contains
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rules
}

// Save values that contain a queryed value as a map
func getContains(query string, data map[string][]Color) map[string]struct{} {
	contains := make(map[string]struct{})
	for key, values := range data {
		for _, value := range values {
			if query == value.Name {
				contains[key] = struct{}{}
			}
		}
	}
	return contains
}

func check(color string, rules map[string][]Color, contains map[string]struct{}) {
	for c := range getContains(color, rules) {
		contains[c] = struct{}{}
		check(c, rules, contains)
	}
}

func getAmount(color string, rules map[string][]Color) int {
	var amount int
	for _, col := range rules[color] {
		amount += col.Amount
		amount += col.Amount * getAmount(col.Name, rules)
	}

	return amount
}
