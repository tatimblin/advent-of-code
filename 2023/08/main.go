package main

import (
	"fmt"
	"strings"

	"github.com/tatimblin/advent-of-code/go-util"
)

var START, END string = "AAA", "ZZZ"

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		return
	}

	connections := map[string][]string{}
	pointer := START
	for i := 2; i < len(lines); i++ {
		lines[i] = strings.Replace(lines[i], " ", "", -1)
		lines[i] = strings.Replace(lines[i], "(", "", -1)
		lines[i] = strings.Replace(lines[i], ")", "", -1)
		parts := strings.Split(lines[i], "=")
		value := strings.Split(parts[1], ",")
		connections[parts[0]] = value
	}

	instructions := strings.Split(lines[0], "")
	i := 0
	for pointer != END {
		instruction := instructions[i%len(instructions)]
		i += 1

		if instruction == "L" {
			pointer = connections[pointer][0]
		} else {
			pointer = connections[pointer][1]
		}
	}

	fmt.Println("part one", i)
}
