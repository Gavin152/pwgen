package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDigitCount(t *testing.T) {
	count := 13
	rNum := RandomIntAsString(count)
	assert.Equalf(t, count, len(rNum), "random number %d should contain %d digits", rNum, count)
}

func TestRandomness(t *testing.T) {
	count := 8
	rnum01 := RandomIntAsString(count)
	rnum02 := RandomIntAsString(count)
	assert.NotEqualf(t, rnum01, rnum02, "Both numbers should not be equal")
}

func TestRandomChars(t *testing.T) {
	count := 25
	options := RandomOptions{
		Lower:   true,
		Upper:   true,
		Numbers: false,
		Symbols: false,
	}
	randString := RandomChars(count, &options)
	assert.Equalf(t, count, len(randString), "Random string should be %d characters long", count)
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
		assert.Regexpf(t, scenario.matchString, charSet, "character set of %s should not match regex %s. CharSet: '%s'", key, scenario.matchString, charSet)
	}
}

func TestRangeMapper(t *testing.T) {
	mapped := mapRange(5, 0, 10, 0, 100)
	assert.Equalf(t, 50, mapped, "Mapping low range to high range failed")

	mapped = mapRange(50, 0, 100, 0, 10)
	assert.Equalf(t, 5, mapped, "Mapping high range to low range failed")
}
