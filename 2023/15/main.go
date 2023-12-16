package main

import (
	"fmt"
	"strings"

	"github.com/tatimblin/advent-of-code/2023/15/modules"
	"github.com/tatimblin/advent-of-code/go-util"
)

func main() {
	var result int

	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Println("could not parse input file")
		return
	}

	hashmap := modules.HashMap{}

	for _, step := range strings.Split(lines[0], ",") {
		hash := modules.Encode(step)
		if hash.Slot != 0 {
			hashmap.Set(hash)
		} else {
			hashmap.Remove(hash)
		}
	}

	for boxi := 0; boxi < len(hashmap.Map); boxi++ {
		var pos int
		for sloti := 0; sloti < len(hashmap.Map[0]); sloti++ {
			hash := hashmap.Get(boxi, sloti)
			if hash != (modules.Hash{}) {
				pos += 1
				result += (hash.Box + 1) * pos * (hash.Slot)
			}
		}
	}

	fmt.Println(result)
}
