package thirteen

import (
	"fmt"
	"testing"
)

func TestFindFirstDepartureExample(t *testing.T) {
	arrival := 939
	buses := []int{7, 13, 59, 31, 19}
	want := 295
	bus, waitTime := findFirstDeparture(buses, arrival)
	got := bus * waitTime
	if got != want {
		t.Errorf("expected the answer to be %d, but was %d", want, got)
	}
}

func TestFindFirstAlignedExample(t *testing.T) {
	tests := []struct {
		buses []busInfo
		want  int64
	}{
		{[]busInfo{
			{0, 7},
			{1, 13},
			{4, 59},
			{6, 31},
			{7, 19},
		}, int64(1068781)},
		{[]busInfo{
			{0, 17},
			{2, 13},
			{3, 19},
		}, int64(3417)},
		{[]busInfo{
			{0, 1789},
			{1, 37},
			{2, 47},
			{3, 1889},
		}, int64(1202161486)},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.buses), func(t *testing.T) {
			got := findFirstAligned(tt.buses)
			if got != tt.want {
				t.Errorf("expected result of %v, but got %v", tt.want, got)
			}
		})
	}
}
