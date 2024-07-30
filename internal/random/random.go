package random

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
)

type RandomOptions struct {
	Lower   bool
	Upper   bool
	Numbers bool
	Symbols bool
}

func RandomIntAsString(count int) string {
	rawSlice := make([]byte, count)
	numString := ""

	_, err := rand.Read(rawSlice)
	if err != nil {
		fmt.Println("error generating random byte slice:", err)
	}

	for _, b := range rawSlice {
		numString += strconv.Itoa(int(b % 10))
	}

	return numString
}

func RandomChars(count int, options *RandomOptions) string {
	outputString := ""
	characterPool := []rune(inputChars(options))

	for i := 0; i < count; i++ {
		locString := RandomIntAsString(2)
		loc, err := strconv.Atoi(locString)
		if err != nil {
			fmt.Printf("Encountered an error when parsing random number string into int: %s", locString)
		}
		mapped := mapRange(float64(loc), 0, 99, 0, float64(len(characterPool)-1))
		//fmt.Printf("Random number mapped: %d\n", mapped)
		outputString += string(characterPool[mapped])
	}

	return outputString
}

func inputChars(options *RandomOptions) string {
	characters := "abcdefghijklmnopqrstuvwxyz"
	outputChars := ""
	if options.Lower {
		outputChars += characters
	}
	if options.Upper {
		outputChars += strings.ToUpper(characters)
	}
	if options.Numbers {
		outputChars += "0123456789"
	}
	if options.Symbols {
		outputChars += "!#$%^*()_+-={}[]|;<>,.?"
	}

	return outputChars
}

func mapRange(value, fromMin, fromMax, toMin, toMax float64) int {
	fromRange := fromMax - fromMin
	toRange := toMax - toMin
	ratio := (value - fromMin) / fromRange
	//fmt.Printf("Ratio: %f\n", ratio)
	mapped := toMin + (ratio * toRange)
	return int(mapped)
}
