package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	groupNumber := 0
	allYes := []int{}
	groupAnswer := []string{}

	for i, line := range lines {
		if line == "" || i+1 == len(lines) {
			groupNumber++
			allYes = append(allYes, countMatchingChars(groupAnswer))
			groupAnswer = []string{}
		} else {
			groupAnswer = append(groupAnswer, line)
		}
	}

	matches := 0
	for _, match := range allYes {
		matches += match
	}

	log.Println(matches)
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
