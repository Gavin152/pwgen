package random

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestDigitCount(t *testing.T) {
	count := 13
	rNum := RandomInt(int8(count))
	rStr := strconv.Itoa(int(rNum))
	assert.Equalf(t, count, len(rStr), "random number %d should contain %d digits", rNum, count)
}

func TestRandomness(t *testing.T) {
	count := 8
	rNum_01 := RandomInt(int8(count))
	rNum_02 := RandomInt(int8(count))
	assert.NotEqualf(t, rNum_01, rNum_02, "Both numbers should not be equal")
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
