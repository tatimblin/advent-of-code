package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Operation string
	Argument  int
}

func main() {
	lines := readLines("input.txt")
	log.Println(fmt.Sprintf("Part 1: %d", breakLoop(lines).Accumulator))
	log.Println(fmt.Sprintf("Part 2: %d", findError(lines).Accumulator))
}

func readLines(path string) []Instruction {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		op := parts[0]
		arg, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		lines = append(lines, Instruction{
			Operation: op,
			Argument:  arg,
		})
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return lines
}
