package eleven

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mziter/aoc-2020/common"
)

func TestIsOccupied(t *testing.T) {
	seats := getExampleRoundZero()
	seats[0][3] = byte('#')
	seats[3][6] = byte('#')
	seats[7][5] = byte('#')

	tests := []struct {
		r    int
		c    int
		want bool
	}{
		{0, 1, false},
		{-1, 1, false},
		{-1, -1, false},
		{-1, -1, false},
		{20, 20, false},
		{0, 3, true},
		{3, 6, true},
		{7, 5, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("[%d, %d]", tt.r, tt.c), func(t *testing.T) {
			got := isOccupied(seats, tt.r, tt.c)
			if got != tt.want {
				t.Errorf("expected value of %v, but got %v", tt.want, got)
			}
		})
	}
}

func TestNewRoundOneAdj(t *testing.T) {
	roundZero := getExampleRoundZero()
	want := getExampleRoundOne()
	got, changes := newRoundWithAdj(roundZero)
	if !reflect.DeepEqual(want, got) {
		t.Error("result of next round was not what was expected")
	}
	if changes != 71 {
		t.Errorf("expected number of changes to be 35, but was %d", changes)
	}
}

func TestNewRoundOneVision(t *testing.T) {
	roundZero := getExampleRoundZero()
	want := getExampleRoundOne()
	got, changes := newRoundWithVision(roundZero)
	if !reflect.DeepEqual(want, got) {
		t.Error("result of next round was not what was expected")
	}
	if changes != 71 {
		t.Errorf("expected number of changes to be 35, but was %d", changes)
	}
}

func TestNewRoundTwoAdj(t *testing.T) {
	roundOne := getExampleRoundOne()
	want := getExampleRoundTwoAdj()
	got, changes := newRoundWithAdj(roundOne)
	if !reflect.DeepEqual(want, got) {
		t.Error("result of next round was not what was expected")
	}
	if changes != 51 {
		t.Errorf("expected number of changes to be 51, but was %d", changes)
	}
}

func TestNewRoundTwoVision(t *testing.T) {
	roundOne := getExampleRoundOne()
	want := getExampleRoundTwoVision()
	got, changes := newRoundWithVision(roundOne)
	if !reflect.DeepEqual(want, got) {
		t.Error("result of next round was not what was expected")
	}
	if changes != 64 {
		t.Errorf("expected number of changes to be 64, but was %d", changes)
	}
}

func TestCountSeenOccupied(t *testing.T) {
	seats := getExampleVisionSeats()
	empty := getEmpySeats()
	channel := getChannelSeats()
	countSeats := countSeenOccupied(seats, 4, 3)
	countEmpty := countSeenOccupied(empty, 4, 3)
	countChannel := countSeenOccupied(channel, 3, 3)
	if countSeats != 8 {
		t.Errorf("expected to see 8 occupied seats, but was %d", countSeats)
	}
	if countEmpty != 0 {
		t.Errorf("should never find an occupied seat when they are all empty, but was %d", countEmpty)
	}
	if countChannel != 0 {
		t.Errorf("should never find an occupied seat when they are empty along channels, but was %d", countEmpty)
	}

}

func getExampleRoundZero() [][]byte {
	bs, err := common.GetBytes(strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`))

	if err != nil {
		panic("couldn't get seats from lines")
	}

	return bs
}

func getExampleRoundOne() [][]byte {
	bs, err := common.GetBytes(strings.NewReader(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`))

	if err != nil {
		panic("couldn't get seats from lines")
	}

	return bs
}

func getExampleRoundTwoAdj() [][]byte {
	bs, err := common.GetBytes(strings.NewReader(`#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`))

	if err != nil {
		panic("couldn't get seats from lines")
	}

	return bs
}

func getExampleRoundTwoVision() [][]byte {
	bs, err := common.GetBytes(strings.NewReader(`#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`))

	if err != nil {
		panic("couldn't get seats from lines")
	}

	return bs
}

// Vision is from [4][3]
func getExampleVisionSeats() [][]byte {
	bs, err := common.GetBytes(strings.NewReader(`.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`))
	if err != nil {
		panic("couldn't get seats from lines")
	}

	return bs
}

func getEmpySeats() [][]byte {
	bs, err := common.GetBytes(strings.NewReader(`.........
.........
.........
.........
.........
.........
.........
.........
.........`))
	if err != nil {
		panic("couldn't get seats from lines")
	}

	return bs
}

func getChannelSeats() [][]byte {
	bs, err := common.GetBytes(strings.NewReader(`.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`))
	if err != nil {
		panic("couldn't get seats from lines")
	}

	return bs
}
