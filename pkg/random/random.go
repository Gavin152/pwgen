package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func RandomInt() int {
	c := 4
	rawSlice := make([]byte, c)
	
	_, err := rand.Read(rawSlice)
	if err != nil {
		fmt.Println("error generating random byte slice:", err)
	}

	generatedNumber := int(big.NewInt(0).SetBytes(rawSlice).Uint64())

	return generatedNumber
}
