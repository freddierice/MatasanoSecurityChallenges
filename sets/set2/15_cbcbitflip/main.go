package main

import (
	"fmt"

	"../../../pkg/oracle"
	"../../../pkg/xorbyte"
)

var blockSize = 16

// EncryptAesCbcInsert
// CBCAuthenticate
// comment1=cooking
// %20MCs;userdata=
// ;comment2=%20lik
// e%20a%20pound%20
// of%20bacon
func main() {
	payload := make([]byte, blockSize*2)

	copy(payload, oracle.EncryptAesCbcInsert([]byte("AAAAAAAAAAAAAAAA"))[:2*blockSize])

	change := []byte("A=A;admin=true")
	change = append(change, byte(2))
	change = append(change, byte(2))

	xorbyte.XOR(payload, []byte("comment1=cooking"))
	xorbyte.XOR(payload, change)

	if oracle.CBCAuthenticate(payload) {
		fmt.Println("You are admin!")
	}
}
