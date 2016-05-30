package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"../../../pkg/block"
)

func main() {

	b, err := ioutil.ReadFile("encdata")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not read from encdata\n%v\n", err)
		os.Exit(1)
	}

	str, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not base64 decode\n%v\n", err)
		os.Exit(1)
	}

	key := []byte("YELLOW SUBMARINE")
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ciphertext := make([]byte, len(str)+len(iv))
	plaintext := make([]byte, len(str))

	copy(ciphertext, iv)
	copy(ciphertext[len(iv):], []byte(str))

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		panic("aes is acting up :/")
	}

	block.CBCDecrypt(plaintext, ciphertext, aesCipher)
	plaintext, err = block.UnPad(plaintext, len(iv))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in padding\n%v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(plaintext))
}
