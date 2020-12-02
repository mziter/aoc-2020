package two

import "testing"

func TestValidFreq(t *testing.T) {
	cases := []struct {
		pwd         string
		least       int
		most        int
		char        rune
		expectedOut bool
	}{
		{"abcde", 1, 3, 'a', true},
		{"cdefg", 1, 3, 'b', false},
		{"ccccccccc", 2, 9, 'c', true},
		{"abbcdd", 1, 2, 'b', true},
		{"abbbcdd", 1, 2, 'b', false},
		{"abcdd", 2, 3, 'b', false},
	}
	for _, c := range cases {
		result := validFreq(c.pwd, c.least, c.most, c.char)
		if result != c.expectedOut {
			t.Errorf("Expected case of %+v to result in %v, but was %v", c, c.expectedOut, result)
		}
	}
}

func TestValidLoc(t *testing.T) {
	cases := []struct {
		pwd         string
		least       int
		most        int
		char        rune
		expectedOut bool
	}{
		{"abcde", 1, 3, 'a', true},
		{"cdefg", 1, 3, 'b', false},
		{"ccccccccc", 2, 9, 'c', false},
	}
	for _, c := range cases {
		result := validLoc(c.pwd, c.least, c.most, c.char)
		if result != c.expectedOut {
			t.Errorf("Expected case of %+v to result in %v, but was %v", c, c.expectedOut, result)
		}
	}
}
