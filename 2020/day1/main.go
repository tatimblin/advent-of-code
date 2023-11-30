package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Numbers struct {
	First  int
	Second int
}

func main() {
	var numbers []int

	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("string convert: %s", err)
		}
		numbers = append(numbers, val)
	}

	match := checkSum(numbers, 2020)
	if err != nil {
		log.Fatalf("Sum match not found")
	}
	log.Println(match)
}

func checkSum(numbers []int, sum int) Numbers {
	for i, number := range numbers {
		for _, otherNumber := range numbers[i:] {
			if number+otherNumber == sum {
				return Numbers{
					First:  number,
					Second: otherNumber,
				}
			}
		}
	}
	return Numbers{}
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
