package six

import "testing"

func TestCountDistinctChars(t *testing.T) {
	cases := []struct {
		input  []string
		output int
	}{
		{[]string{"abc"}, 3},
		{[]string{"a", "b", "c"}, 3},
		{[]string{"ab", "bc"}, 3},
		{[]string{"a", "a", "a", "a"}, 1},
		{[]string{"b"}, 1},
	}
	for _, c := range cases {
		res := countDistinctChars(c.input)
		if res != c.output {
			t.Errorf("Expected result of input %s, to be %d, but was %d", c.input, c.output, res)
		}
	}
}

func TestCountAllYes(t *testing.T) {
	cases := []struct {
		input  []string
		output int
	}{
		{[]string{"abc"}, 3},
		{[]string{"a", "b", "c"}, 0},
		{[]string{"ab", "bc"}, 1},
		{[]string{"a", "a", "a", "a"}, 1},
		{[]string{"b"}, 1},
	}
	for _, c := range cases {
		res := countAllYes(c.input)
		if res != c.output {
			t.Errorf("Expected result of input %s, to be %d, but was %d", c.input, c.output, res)
		}
	}
}
