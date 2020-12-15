package fourteen

import (
	"fmt"
	"strconv"
	"testing"
)

func TestParseMask(t *testing.T) {
	tests := []struct {
		input string
		want  [36]int
	}{
		{"1110X1110XXX101X0011010X110X10X0110X",
			[36]int{-1, 0, 1, 1, 0, -1, 0, 1, -1, 0, 1, 1,
				-1, 0, 1, 0, 1, 1, 0, 0, -1, 1, 0, 1,
				-1, -1, -1, 0, 1, 1, 1, -1, 0, 1, 1, 1}},
	}

	for _, tt := range tests {
		got := parseMask(tt.input)
		t.Run(tt.input, func(t *testing.T) {
			if got != tt.want {
				t.Errorf("expected arrays to be equal\nWANT: %v\n GOT: %v", tt.want, got)
			}
		})
	}
}

func TestApplyMask(t *testing.T) {
	mask := [36]int{-1, 0, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	tests := []struct {
		num  int
		want int
	}{
		{11, 73},
		{101, 101},
		{0, 64},
	}

	for _, tt := range tests {
		got := applyMask(mask, tt.num)
		t.Run(fmt.Sprintf("%d", tt.num), func(t *testing.T) {
			if got != tt.want {
				t.Errorf("unexpected result of mask apply to %d\nWANT: %s\n GOT: %s",
					tt.num,
					strconv.FormatInt(int64(tt.want), 2),
					strconv.FormatInt(int64(got), 2))
			}
		})
	}
}
