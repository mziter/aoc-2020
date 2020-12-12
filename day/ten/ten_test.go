package ten

import "testing"

func TestOneAndThreeJoltProductExampleOne(t *testing.T) {
	nums := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	want := 35
	got := oneAndThreeJoltProduct(nums)
	if want != got {
		t.Errorf("expected one and three jolt product to be %d, but was %d", want, got)
	}
}

func TestOneAndThreeJoltProductExampleTwo(t *testing.T) {
	nums := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49,
		45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
	want := 220
	got := oneAndThreeJoltProduct(nums)
	if want != got {
		t.Errorf("expected one and three jolt product to be %d, but was %d", want, got)
	}
}

func TestPossibilitiesExampleOne(t *testing.T) {
	nums := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	want := 8
	got := possibilities(nums)
	if want != got {
		t.Errorf("expected possibilities be %d, but was %d", want, got)
	}
}
