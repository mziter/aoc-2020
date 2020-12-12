package eleven

import (
	"fmt"
	"strconv"

	"github.com/mziter/aoc-2020/common"
)

type (
	// PartOneSolver implements solver interface for part one
	PartOneSolver struct{}
	// PartTwoSolver implements solver interface for part one
	PartTwoSolver struct{}
)

// Solve implements solver interface for part one
func (d PartOneSolver) Solve() string {
	seats, err := common.GetBytesFromFile("day/eleven/input.txt")
	if err != nil {
		panic("couldn't read input file for day ten")
	}
	nextSeats := seats
	for {
		nextRound, numChanges := newRoundWithAdj(nextSeats)
		if numChanges == 0 {
			return strconv.Itoa(countOccupied(nextRound))
		} 
		nextSeats = nextRound
	}
}

// Solve implements solver interface for part one
func (d PartTwoSolver) Solve() string {
	seats, err := common.GetBytesFromFile("day/eleven/input.txt")
	if err != nil {
		panic("couldn't read input file for day ten")
	}
	nextSeats := seats
	for {
		nextRound, numChanges := newRoundWithVision(nextSeats)
		if numChanges == 0 {
			return strconv.Itoa(countOccupied(nextRound))
		}
		nextSeats = nextRound
	}
}

func newRoundWithAdj(seats [][]byte) ([][]byte, int) {
	changes := 0
	newSeats := make([][]byte, len(seats))
	for r := 0; r < len(seats); r++ {
		seatRow := make([]byte, len(seats[0]))
		for c := 0; c < len(seats[0]); c++ {
			if seats[r][c] == byte('L') && countAdjOccupied(seats, r, c) == 0 {
				seatRow[c] = byte('#')
				changes++
			} else if isOccupied(seats, r, c) && countAdjOccupied(seats, r, c) >= 4 {
				seatRow[c] = byte('L')
				changes++
			} else {
				seatRow[c] = seats[r][c]
			}
		}
		newSeats[r] = seatRow
	}
	return newSeats, changes
}

func newRoundWithVision(seats [][]byte) ([][]byte, int) {
	changes := 0
	newSeats := make([][]byte, len(seats))
	for r := 0; r < len(seats); r++ {
		seatRow := make([]byte, len(seats[0]))
		for c := 0; c < len(seats[0]); c++ {
			if seats[r][c] == byte('L') && countSeenOccupied(seats, r, c) == 0 {
				seatRow[c] = byte('#')
				changes++
			} else if isOccupied(seats, r, c) && countSeenOccupied(seats, r, c) >= 5 {
				seatRow[c] = byte('L')
				changes++
			} else {
				seatRow[c] = seats[r][c]
			}
		}
		newSeats[r] = seatRow
	}
	return newSeats, changes
}

func countAdjOccupied(seats [][]byte, r, c int) int {
	count := 0
	for offsetR := -1; offsetR <= 1; offsetR++ {
		for offsetC := -1; offsetC <= 1; offsetC++ {
			adjR := r + offsetR
			adjC := c + offsetC
			if !(r == adjR && c == adjC) && isOccupied(seats, adjR, adjC) {
				count++
			}
		}
	}
	return count
}

func countSeenOccupied(seats [][]byte, row, col int) int {
	maxCol := len(seats[0]) - 1
	maxRow := len(seats) - 1
	count := 0
	// looking left
	for c := col - 1; c >= 0; c-- {
		if seats[row][c] == byte('#') {
			count++
			break
		} else if seats[row][c] == byte('L') {
			break
		}
	}

	// looking diagonal up and to left
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if seats[r][c] == byte('#') {
			count++
			break
		} else if seats[r][c] == byte('L') {
			break
		}
	}

	// looking up
	for r := row - 1; r >= 0; r-- {
		if seats[r][col] == byte('#') {
			count++
			break
		} else if seats[r][col] == byte('L') {
			break
		}
	}

	// looking diagonal up and to right
	for r, c := row-1, col+1; r >= 0 && c <= maxCol; r, c = r-1, c+1 {
		if seats[r][c] == byte('#') {
			count++
			break
		} else if seats[r][c] == byte('L') {
			break
		}
	}

	// looking right
	for c := col + 1; c <= maxCol; c++ {
		if seats[row][c] == byte('#') {
			count++
			break
		} else if seats[row][c] == byte('L') {
			break
		}
	}

	// looking diagonal down and to right
	for r, c := row+1, col+1; r <= maxRow && c <= maxCol; r, c = r+1, c+1 {
		if seats[r][c] == byte('#') {
			count++
			break
		} else if seats[r][c] == byte('L') {
			break
		}
	}

	// looking down
	for r := row + 1; r <= maxRow; r++ {
		if seats[r][col] == byte('#') {
			count++
			break
		} else if seats[r][col] == byte('L') {
			break
		}
	}

	// looking diagonal down and to left
	for r, c := row+1, col-1; r <= maxRow && c >= 0; r, c = r+1, c-1 {
		if seats[r][c] == byte('#') {
			count++
			break
		} else if seats[r][c] == byte('L') {
			break
		}
	}

	return count
}

func countOccupied(seats [][]byte) int {
	count := 0
	for _, r := range seats {
		for _, c := range r {
			if c == byte('#') {
				count++
			}
		}
	}
	return count
}

func isOccupied(seats [][]byte, r, c int) bool {
	if r < 0 || r >= len(seats) {
		return false
	}
	if c < 0 || c >= len(seats[0]) {
		return false
	}
	return seats[r][c] == byte('#')
}

func print(seats [][]byte) {
	for _, r := range seats {
		for _, c := range r {
			fmt.Printf("%c", rune(c))
		}
		fmt.Println()
	}
}
