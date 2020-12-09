package nine

import "testing"

func TestFirstNonSum(t *testing.T) {
	nums := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	got := firstNonSum(nums, 5)
	want := 127
	if got != want {
		t.Errorf("answer from aoc example should have been %d, but was %d", want, got)
	}

}

func TestContiguousSumMinMax(t *testing.T) {
	nums := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	got := contiguousSumMinMax(nums, 127)
	want := 62
	if got != want {
		t.Errorf("answer from aoc example should have been %d, but was %d", want, got)
	}

}
