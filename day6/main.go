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

	var grades []int
	var groupAnswers string
	summed := 0

	for i, line := range lines {
		if line != "" {
			groupAnswers = groupAnswers + line
		} else {
			// End group on newline
			grades = append(grades, countUniqueChars(groupAnswers))
			summed = summed + countUniqueChars(groupAnswers)
			groupAnswers = ""
		}

		// End group on eof
		if i+1 == len(lines) {
			grades = append(grades, countUniqueChars(groupAnswers))
			summed = summed + countUniqueChars(groupAnswers)
			groupAnswers = ""
		}
	}

	log.Println(len(grades), summed)
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
