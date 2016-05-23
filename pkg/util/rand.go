package util

import (
	cryptorand "crypto/rand"
	"fmt"
	mathrand "math/rand"
	"os"
)

func GenerateRandomBytesRange(smallest, largest int) []byte {
	randlen := mathrand.Intn(largest - smallest)
	return GenerateRandomBytes(randlen)
}

func GenerateRandomBytes(randlen int) []byte {

	b := make([]byte, randlen)
	if n, err := cryptorand.Read(b); err != nil || n != len(b) {
		fmt.Fprintf(os.Stderr, "this kind of error was unexpected: %v, %v", n, err)
		os.Exit(1)
	}

	return b
}
