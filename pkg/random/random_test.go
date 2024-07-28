package random

import (
	"regexp"
	"strconv"
	"testing"
)

func TestDigitCount(t *testing.T) {
	count := 13
	rNum := RandomInt(int8(count))
	rStr := strconv.Itoa(int(rNum))
	rNumDigits := len(rStr)
	if rNumDigits != count {
		t.Fatalf(`%d is not of length %d`, rNumDigits, count)
	}
}

func TestRandomness(t *testing.T) {
	count := 8
	rNum_01 := RandomInt(int8(count))
	rNum_02 := RandomInt(int8(count))
	if rNum_01 == rNum_02 {
		t.Fatalf(`Random int 1 was equal to int 2. %d == %d`, rNum_01, rNum_02)
	}
}

func TestCharacterSet(t *testing.T) {
	type scenario struct {
		matchString string
		options     RandomOptions
	}
	scenarios := make(map[string]scenario)
	scenarios["LowerCaseOnly"] = scenario{
		matchString: `^[a-z]+$`,
		options: RandomOptions{
			Lower:   true,
			Upper:   false,
			Numbers: false,
			Symbols: false,
		},
	}
	scenarios["UpperCaseOnly"] = scenario{
		matchString: `^[A-Z]+$`,
		options: RandomOptions{
			Lower:   false,
			Upper:   true,
			Numbers: false,
			Symbols: false,
		},
	}
	scenarios["NumberOnly"] = scenario{
		matchString: `^[0-9]+$`,
		options: RandomOptions{
			Lower:   false,
			Upper:   false,
			Numbers: true,
			Symbols: false,
		},
	}
	scenarios["SymbolOnly"] = scenario{
		matchString: `^[!#$%^*()_+-={}[]|;<>,.?]+$`,
		options: RandomOptions{
			Lower:   false,
			Upper:   false,
			Numbers: false,
			Symbols: true,
		},
	}
	scenarios["UpperAndLower"] = scenario{
		matchString: `^[a-zA-Z]+$`,
		options: RandomOptions{
			Lower:   true,
			Upper:   true,
			Numbers: false,
			Symbols: false,
		},
	}
	scenarios["UpperAndNumbers"] = scenario{
		matchString: `^[A-Z0-9]+$`,
		options: RandomOptions{
			Lower:   false,
			Upper:   true,
			Numbers: true,
			Symbols: false,
		},
	}

	for key, scenario := range scenarios {
		charSet := inputChars(&scenario.options)
		found, _ := regexp.MatchString(scenario.matchString, charSet)
		if !found {
			t.Logf("Character set is invalid for %s: %s", key, charSet)
			t.Fail()
		}
	}
}

func TestRangeMapper(t *testing.T) {
	mapped := mapRange(5, 0, 10, 0, 100)
	if mapped != 50 {
		t.Logf("Mapped value is %d, want 50", mapped)
		t.Fail()
	}
	mapped = mapRange(50, 0, 100, 0, 10)
	if mapped != 5 {
		t.Logf("Mapped value is %d, want 5", mapped)
		t.Fail()
	}
}
