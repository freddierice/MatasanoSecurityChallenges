package oracle

import (
	"crypto/aes"
	"encoding/base64"
	"math/rand"
	"time"

	"../block"
	"../util"
)

var key []byte
var message []byte

var messageString = "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
	"aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
	"dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
	"YnkK"
var keylengths = [...]int{16, 32}

func init() {

	var err error

	rand.Seed(time.Now().UnixNano())
	keylength := keylengths[rand.Intn(2)]
	key = util.GenerateRandomBytes(keylength)
	message, err = base64.StdEncoding.DecodeString(messageString)
	if err != nil {
		panic("message is not properly base64 encoded")
	}
}

func EncryptAesEcbPrepend(plaintext []byte) []byte {

	plaintext = append(plaintext, message...)
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		panic("aes did something weird")
	}

	plaintext = block.Pad(plaintext, len(key))
	ciphertext := make([]byte, len(plaintext))

	block.ECBEncrypt(ciphertext, plaintext, aesCipher)
	return ciphertext
}
