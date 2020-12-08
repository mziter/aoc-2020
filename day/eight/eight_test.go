package eight

import (
	"testing"
)

func TestMustParseInstruction(t *testing.T) {
	testCases := []struct {
		input string
		want  Instruction
	}{
		{"nop +0", Instruction{NOP, 0}},
		{"acc +1", Instruction{ACC, 1}},
		{"jmp +4", Instruction{JMP, 4}},
		{"acc +3", Instruction{ACC, 3}},
		{"jmp -3", Instruction{JMP, -3}},
		{"acc -99", Instruction{ACC, -99}},
		{"acc +1", Instruction{ACC, 1}},
		{"jmp -4", Instruction{JMP, -4}},
		{"acc +6", Instruction{ACC, 6}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := mustParseInstruction(tc.input)
			if got != tc.want {
				t.Errorf("expected input of %s to result in %v, but was instead %v", tc.input, tc.want, got)
			}
		})
	}
}

func TestGetAccAtLoop(t *testing.T) {
	testCases := []struct {
		name  string
		input BootCode
		want  int
	}{
		{"Example One", getExampleBootCode(), 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := getAccAtLoop(tc.input)
			if !ok {
				t.Error("expected a loop, but there wasn't one")
			}
			if got != tc.want {
				t.Errorf("expected accumulator value to be %d, but was %d", tc.want, got)
			}
		})
	}
}

func TestGetAccAtLoopTerminates(t *testing.T) {
	bc := getExampleBootCode()
	bc.instructions[7].Op = NOP
	got, isLoop := getAccAtLoop(bc)
	if isLoop {
		t.Error("result should not be a loop but a properly terminated program")
	}
	want := 8
	if got != want {
		t.Errorf("expected result to be %d, not %d", want, got)
	}
}

func getExampleBootCode() BootCode {
	return NewBootCodeFromStringInstructions([]string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	})
}
