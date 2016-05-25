package main

import (
	"bytes"
	"fmt"
)

// email=&uid=10&role=user

// this is the payload we want:
// email=foob@rbar.
// com&uid=10&role=
// admin&uid=10&rol
// =user
func main() {

	blocks := make([][]byte, 4)

	enc1 := EncryptProfile("foob@rbar.com")
	blocks[0] = enc1[0:blockSize]
	blocks[1] = enc1[blockSize : 2*blockSize]

	enc2 := EncryptProfile("AAAAAAAAAAadmin")
	blocks[2] = enc2[blockSize : 2*blockSize]

	enc3 := EncryptProfile("AAAAAAAAAAAAAA")
	blocks[3] = enc3[2*blockSize : 3*blockSize]

	ciphertext := bytes.Join(blocks, []byte{})
	if ValidateAdmin(ciphertext) {
		fmt.Printf("Congrats, you are admin!\n")
	}
}
