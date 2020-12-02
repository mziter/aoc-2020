package one

import "testing"

func TestGetTwoNumSum(t *testing.T) {
	const target = 2020
	result := getTwoNumSum(exampleNums(), target)
	if result.num1+result.num2 != target {
		t.Errorf("Result of adding %d and %d should have been %d, but isn't", result.num1, result.num2, target)
	}
}

func TestGetThreeNumSum(t *testing.T) {
	const target = 2020
	result := getThreeNumSum(exampleNums(), target)
	if result.num1+result.num2+result.num3 != target {
		t.Errorf("Result of adding %d and %d should have been %d, but isn't", result.num1, result.num2, target)
	}
}

func exampleNums() []int {
	return []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
}
