package random

import (
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
