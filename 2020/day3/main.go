package main

import (
	"bufio"
	"log"
	"os"
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

	treeCount := 0

	for y, line := range lines {
		if y%2 == 0 {
			level := strings.Split(line, "")
			x := y * 7
			if x >= len(level) {
				x = x % len(level)
			}

			// Check for tree
			if level[x] == "#" {
				treeCount = treeCount + 1
			}
		}
	}

	log.Println(treeCount)
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
