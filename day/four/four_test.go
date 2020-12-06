package four

import (
	"strings"
	"testing"

	"github.com/mziter/aoc-2020/common"
)

func TestValidPartOnePassport(t *testing.T) {
	testCases := []struct {
		passport []string
		isValid  bool
	}{
		{[]string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
			"byr:1937 iyr:2017 cid:147 hgt:183cm",
		}, true},
		{[]string{
			"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
			"hcl:#cfa07d byr:1929",
		}, false},
		{[]string{
			"hcl:#ae17e1 iyr:2013",
			"eyr:2024",
			"ecl:brn pid:760753108 byr:1931",
			"hgt:179cm",
		}, true},
		{[]string{
			"hcl:#cfa07d eyr:2025 pid:166559648",
			"iyr:2011 ecl:brn hgt:59in",
		}, false},
	}
	for _, testCase := range testCases {
		res := validPartOnePassport(testCase.passport)
		if res != testCase.isValid {
			t.Errorf("Expected testcase of %v to be %v, but it was %v", testCase.passport, testCase.isValid, res)
		}
	}
}

func TestGetPassports(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
	lines := strings.Split(input, "\n")
	passports := common.SplitLines(lines)
	if len(passports) != 4 {
		t.Errorf("Expected four passports but there were actually %d", len(passports))
	}
}

func TestValidBirthYear(t *testing.T) {
	const key string = "byr"
	cases := []struct {
		value string
		valid bool
	}{
		{"1920", true},
		{"1921", true},
		{"1919", false},

		{"2002", true},
		{"2001", true},
		{"2003", false},

		{"1978", true},
		{"1998", true},
		{"2000", true},
	}
	for _, c := range cases {
		res := isValidField(key, c.value)
		if res != c.valid {
			t.Errorf("Expected birth year of %s to be %v, but was %v", c.value, c.valid, res)
		}
	}
}
func TestValidIssueYear(t *testing.T) {
	const key string = "iyr"
	cases := []struct {
		value string
		valid bool
	}{
		{"2010", true},
		{"2011", true},
		{"2009", false},

		{"2020", true},
		{"2019", true},
		{"2021", false},
	}
	for _, c := range cases {
		res := isValidField(key, c.value)
		if res != c.valid {
			t.Errorf("Expected issue year of %s to be %v, but was %v", c.value, c.valid, res)
		}
	}
}
func TestValidExpirationYear(t *testing.T) {
	const key string = "eyr"
	cases := []struct {
		value string
		valid bool
	}{
		{"2020", true},
		{"2021", true},
		{"2019", false},

		{"2030", true},
		{"2029", true},
		{"2031", false},
	}
	for _, c := range cases {
		res := isValidField(key, c.value)
		if res != c.valid {
			t.Errorf("Expected expiration year of %s to be %v, but was %v", c.value, c.valid, res)
		}
	}
}

func TestValidHairColor(t *testing.T) {
	cases := []struct {
		input string
		valid bool
	}{
		{"#abc123", true},
		{"#afc028", true},
		{"#bff098", true},
		{"#bff0988", false},
		{"#abc12", false},
		{"#agz123", false},
		{"abc123", false},
	}
	for _, c := range cases {
		res := isValidHairColor(c.input)
		if res != c.valid {
			t.Errorf("Expected hair color of %s to be %v, but was %v", c.input, c.valid, res)
		}
	}
}

func TestValidPID(t *testing.T) {
	cases := []struct {
		input string
		valid bool
	}{
		{"123456789", true},
		{"012345678", true},
		{"000000000", true},
		{"1", false},
		{"12345", false},
		{"00012345", false},
		{"0123456789", false},
	}
	for _, c := range cases {
		res := isValidPID(c.input)
		if res != c.valid {
			t.Errorf("Expected PID of %s to be %v, but was %v", c.input, c.valid, res)
		}
	}
}

func TestValidField(t *testing.T) {
	cases := []struct {
		key   string
		value string
		valid bool
	}{
		{"byr", "2002", true},
		{"byr", "2003", false},

		{"hgt", "60in", true},
		{"hgt", "190cm", true},
		{"hgt", "190in", false},
		{"hgt", "190", false},

		{"hcl", "#123abc", true},
		{"hcl", "#123abz", false},
		{"hcl", "123abc", false},

		{"ecl", "brn", true},
		{"ecl", "wat", false},

		{"pid", "000000001", true},
		{"pid", "0123456789", false},
	}
	for _, c := range cases {
		res := isValidField(c.key, c.value)
		if res != c.valid {
			t.Errorf("Expected [key:%s, val:%s] to be %v, but was %v", c.key, c.value, c.valid, res)
		}
	}
}
