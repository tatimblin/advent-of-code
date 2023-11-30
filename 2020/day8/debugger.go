package main

type BreakLoopResponse struct {
	Accumulator int
	Success     bool
}

func breakLoop(instructions []Instruction) BreakLoopResponse {
	history := make(map[int]bool)
	var accumulator int
	var index int
	for {
		op := instructions[index].Operation
		arg := instructions[index].Argument

		if op == "acc" {
			accumulator += arg
		}

		if op == "jmp" {
			index += arg
		} else {
			index++
		}

		// Loop detected
		if history[index] {
			return BreakLoopResponse{
				Accumulator: accumulator,
				Success:     false,
			}
		}

		// Program successfully reached the end
		if index >= len(instructions) {
			return BreakLoopResponse{
				Accumulator: accumulator,
				Success:     true,
			}
		}

		history[index] = true
	}
}

func findError(instructions []Instruction) BreakLoopResponse {
	for i := range instructions {
		var (
			test = instructions
			op   = test[i].Operation
		)
		if op == "jmp" {
			test[i].Operation = "nop"
		} else if op == "nop" {
			test[i].Operation = "jmp"
		}

		breakLoop := breakLoop(test)
		if breakLoop.Success {
			return breakLoop
		}

		// Unsure why I need to reset the value
		if op == "jmp" {
			test[i].Operation = "jmp"
		} else if op == "nop" {
			test[i].Operation = "nop"
		}
	}
	return BreakLoopResponse{
		Accumulator: 0,
		Success:     false,
	}
}
