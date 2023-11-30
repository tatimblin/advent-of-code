package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type PasswordLogic struct {
	Min      int
	Max      int
	Char     string
	Password string
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	debugs := processLines(lines)
	validCount := 0

	for _, debug := range debugs {
		min := string(debug.Password[debug.Min-1]) == debug.Char
		max := string(debug.Password[debug.Max-1]) == debug.Char

		if (min || max) && !(min && max) {
			validCount = validCount + 1
		}
	}

	log.Println(validCount)
}

func processLines(lines []string) []PasswordLogic {
	passwordLogics := []PasswordLogic{}

	for _, line := range lines {
		parts := strings.Fields(line)
		threshold := strings.Split(parts[0], "-")
		threshold2 := []int{}

		for _, num := range threshold {
			j, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			threshold2 = append(threshold2, j)
		}

		passwordLogic := PasswordLogic{
			Min:      threshold2[0],
			Max:      threshold2[1],
			Char:     firstChar(parts[1]),
			Password: parts[2],
		}
		passwordLogics = append(passwordLogics, passwordLogic)
	}

	return passwordLogics
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

func firstChar(str string) string {
	return strings.Split(str, "")[0]
}
