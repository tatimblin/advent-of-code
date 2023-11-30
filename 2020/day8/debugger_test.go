package main

import (
	"fmt"
	"testing"
)

func TestBreakLoop(t *testing.T) {
	t.Run("Catch infinite loop", testBreakLoop([]Instruction{
		Instruction{
			Operation: "noc",
			Argument:  0,
		},
		Instruction{
			Operation: "acc",
			Argument:  1,
		},
		Instruction{
			Operation: "jmp",
			Argument:  4,
		},
		Instruction{
			Operation: "acc",
			Argument:  3,
		},
		Instruction{
			Operation: "jmp",
			Argument:  -3,
		},
		Instruction{
			Operation: "acc",
			Argument:  -99,
		},
		Instruction{
			Operation: "acc",
			Argument:  1,
		},
		Instruction{
			Operation: "jmp",
			Argument:  -4,
		},
		Instruction{
			Operation: "acc",
			Argument:  6,
		},
	}, BreakLoopResponse{
		Accumulator: 5,
		Success:     false,
	}))
	t.Run("Successful run", testBreakLoop([]Instruction{
		Instruction{
			Operation: "noc",
			Argument:  0,
		},
		Instruction{
			Operation: "acc",
			Argument:  1,
		},
		Instruction{
			Operation: "jmp",
			Argument:  4,
		},
		Instruction{
			Operation: "acc",
			Argument:  3,
		},
		Instruction{
			Operation: "jmp",
			Argument:  -3,
		},
		Instruction{
			Operation: "acc",
			Argument:  -99,
		},
		Instruction{
			Operation: "acc",
			Argument:  1,
		},
		Instruction{
			Operation: "nop",
			Argument:  -4,
		},
		Instruction{
			Operation: "acc",
			Argument:  6,
		},
	}, BreakLoopResponse{
		Accumulator: 8,
		Success:     true,
	}))
}

func testBreakLoop(body []Instruction, expected BreakLoopResponse) func(*testing.T) {
	return func(t *testing.T) {
		actual := breakLoop(body)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected: %v but got %v!", expected.Accumulator, actual.Accumulator))
		}
	}
}
