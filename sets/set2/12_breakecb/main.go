package main

import (
	"fmt"

	"../../../pkg/breakcrypto"
	"../../../pkg/oracle"
)

var charset = []byte{' ', 'e', 't', 'a', 'o', 'i', 'n', 's', 'h', 'r', 'd', 'l', 'c', 'u', 'm', 'w', 'f', 'g', 'y', 'p', 'b', 'v', 'k', 'j', 'x', 'q', 'z', 'E', 'T', 'A', 'O', 'I', 'N', 'S', 'H', 'R', 'D', 'L', 'C', 'U', 'M', 'W', 'F', 'G', 'Y', 'P', 'B', 'V', 'K', 'J', 'X', 'Q', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', ':', '?', '!', '(', ')', '[', ']'}

func init() {
	set := make(map[byte]bool)
	for _, c := range charset {
		set[c] = true
	}
	for i := 0; i < 0x100; i++ {
		if _, ok := set[byte(i)]; !ok {
			charset = append(charset, byte(i))
		}
	}
}

func main() {
	message := breakcrypto.BreakPrependECBOracle(oracle.EncryptAesEcbPrepend, charset)
	fmt.Printf("The message is: \n%s\n", message)
}
