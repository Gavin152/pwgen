package random

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strings"
)

type RandomOptions struct {
	Lower   bool
	Upper   bool
	Numbers bool
	Symbols bool
}

func RandomInt(count int8) uint64 {
	c := 8
	rawSlice := make([]byte, c)

	_, err := rand.Read(rawSlice)
	if err != nil {
		fmt.Println("error generating random byte slice:", err)
	}

	generatedNumber := big.NewInt(0).SetBytes(rawSlice).Uint64()
	stripped := generatedNumber % uint64((math.Pow(float64(10), float64(count))))
	return stripped
}

func RandomChars(count int, options *RandomOptions) string {
	outputString := ""
	characterPool := []rune(inputChars(options))

	for i := 0; i < count; i++ {
		loc := int(RandomInt(int8(2)))
		//fmt.Printf("Random numder generated: %d\n", loc)
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
