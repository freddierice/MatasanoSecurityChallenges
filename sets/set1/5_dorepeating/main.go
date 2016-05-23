package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"../../../pkg/xorbyte"
)

// reads the key from stdin then encrypts the bytes until eof
func main() {
	reader := bufio.NewReader(os.Stdin)

	key, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("could not get the key : %v", err)
	}
	key = strings.Trim(key, "\n")

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("could not read from stdin : %v", err)
	}

	xorbyteCipher, err := xorbyte.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("could not create a NewCipher")
	}

	xorbyteCipher.Encrypt(bytes, bytes)
	str := hex.EncodeToString(bytes)
	fmt.Printf("\n%s\n", str)
}

func encrypt(plaintext []byte, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	keylen := len(key)
	for i := range plaintext {
		ciphertext[i] = plaintext[i] ^ key[i%keylen]
	}
	return ciphertext
}
