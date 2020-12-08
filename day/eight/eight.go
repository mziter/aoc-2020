package eight

import (
	"strconv"
	"strings"

	"github.com/mziter/aoc-2020/common"
)

type (
	// PartOneSolver implements solver interface for part one
	PartOneSolver struct{}
	// PartTwoSolver implements solver interface for part one
	PartTwoSolver struct{}
)

type (
	// OpType is a string alias
	OpType string

	// Instruction contains the operation to be performed and the value
	// associated with it
	Instruction struct {
		Op    OpType
		Value int
	}

	// BootCode contains all instrutions, current position,
	// and positions that have been executed
	BootCode struct {
		instructions []Instruction
		posExecuted  map[int]bool
		pos          int
		acc          int
	}
)

const (
	// NOP stands for No Operation
	NOP = "nop"
	// ACC stands for Accumulator
	ACC = "acc"
	// JMP stands for Jump
	JMP = "jmp"
)

var ops = map[string]OpType{
	"nop": NOP,
	"acc": ACC,
	"jmp": JMP,
}

func mustParseOp(s string) OpType {
	v, ok := ops[s]
	if !ok {
		panic("couldn't parse input:" + s)
	}
	return v
}

// Solve implements solver interface for part one
func (d PartOneSolver) Solve() string {
	lines, err := common.GetLines("day/eight/input.txt")
	if err != nil {
		panic("couldn't open input file for day eight")
	}

	bootCode := NewBootCodeFromStringInstructions(lines)
	res, ok := getAccAtLoop(bootCode)
	if !ok {
		panic("couldn't detect loop")
	}
	return strconv.Itoa(res)
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	lines, err := common.GetLines("day/eight/input.txt")
	if err != nil {
		panic("couldn't open input file for day eight")
	}

	orig := mustParseInstructions(lines)

	jmpPos := []int{}
	for pos, i := range orig {
		if i.Op == JMP {
			jmpPos = append(jmpPos, pos)
		}
	}
	for pos := range jmpPos {
		instruc := make([]Instruction, len(orig))
		copy(instruc, orig)
		instruc[pos].Op = NOP
		bc := NewBootCode(instruc)
		acc, isLoop := getAccAtLoop(bc)
		if !isLoop && acc != -1 {
			return strconv.Itoa(acc)
		}
	}

	nopPos := []int{}
	for pos, i := range orig {
		if i.Op == NOP {
			nopPos = append(nopPos, pos)
		}
	}
	for pos := range nopPos {
		instruc := make([]Instruction, len(orig))
		copy(instruc, orig)
		instruc[pos].Op = JMP
		bc := NewBootCode(instruc)
		acc, isLoop := getAccAtLoop(bc)
		if !isLoop && acc != -1 {
			return strconv.Itoa(acc)
		}
	}

	return "NO SOLUTION FOUND"
}

// NewBootCodeFromStringInstructions creates a new boot code struct from
// a slice of strings
func NewBootCodeFromStringInstructions(input []string) BootCode {
	return NewBootCode(mustParseInstructions(input))
}

// NewBootCode creates a new boot code struct
func NewBootCode(instructions []Instruction) BootCode {
	return BootCode{
		instructions: instructions,
		posExecuted:  map[int]bool{},
		pos:          0,
		acc:          0,
	}
}

func getAccAtLoop(bc BootCode) (int, bool) {
	if bc.posExecuted[bc.pos] {
		return bc.acc, true
	}
	if bc.pos < 0 || bc.pos >= len(bc.instructions) {
		return bc.acc, false
	}
	instruction := bc.instructions[bc.pos]
	newBc := executeInstruction(bc, instruction)
	return getAccAtLoop(newBc)
}

func executeInstruction(bc BootCode, i Instruction) BootCode {
	if bc.pos < 0 || bc.pos >= len(bc.instructions) {
		panic("position out of bounds")
	}
	bc.posExecuted[bc.pos] = true
	switch val := i.Value; i.Op {
	case NOP:
		return BootCode{
			instructions: bc.instructions,
			posExecuted:  bc.posExecuted,
			pos:          bc.pos + 1,
			acc:          bc.acc,
		}
	case JMP:
		return BootCode{
			instructions: bc.instructions,
			posExecuted:  bc.posExecuted,
			pos:          bc.pos + val,
			acc:          bc.acc,
		}
	case ACC:
		return BootCode{
			instructions: bc.instructions,
			posExecuted:  bc.posExecuted,
			pos:          bc.pos + 1,
			acc:          bc.acc + val,
		}
	default:
		panic("invalid operation encountered")
	}
}

func mustParseInstructions(ss []string) []Instruction {
	instructions := []Instruction{}
	for _, s := range ss {
		instructions = append(instructions, mustParseInstruction(s))
	}
	return instructions
}

func mustParseInstruction(s string) Instruction {
	tokens := strings.Split(s, " ")
	op := mustParseOp(tokens[0])
	val, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic("couldn't parse instruction: " + s)
	}
	return Instruction{
		Op:    op,
		Value: val,
	}
}
