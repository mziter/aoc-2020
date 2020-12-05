package five

import "testing"

func TestGetBinaryNumber(t *testing.T) {
	cases := []struct {
		input  string
		output float64
	}{
		{"FBFBBFF", 44.0},
		{"RLR", 5.0},
	}
	for _, c := range cases {
		res := getBoardingNumber(c.input)
		if res != c.output {
			t.Errorf("Expected result of %s to be %f, but was %f", c.input, c.output, res)
		}
	}
}

func TestGetBoardingPassID(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}
	for _, c := range cases {
		res := getBoardingPassID(c.input)
		if res != c.output {
			t.Errorf("Expected result of %s to be %d, but was %d", c.input, c.output, res)
		}
	}
}
