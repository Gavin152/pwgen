package random

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

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
