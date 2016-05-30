package oracle

import (
	"crypto/aes"
	"strings"

	"../block"
	"../parse"
	"../util"
)

var blockSize = 16

func EncryptAesCbcInsert(plaintext []byte) []byte {

	plainstr := string(plaintext)
	plainstr = strings.Replace(plainstr, ",", "%2c", -1)
	plainstr = strings.Replace(plainstr, "&", "%26", -1)

	plaintext = []byte(plaintext)
	plaintext = append([]byte("comment1=cooking%20MCs;userdata="), plaintext...)
	plaintext = append(plaintext, []byte(";comment2=%20like%20a%20pound%20of%20bacon")...)
	plaintext = block.Pad(plaintext, blockSize)

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		panic("aes did something weird")
	}

	ciphertext := make([]byte, len(plaintext)+blockSize)
	iv := util.GenerateRandomBytes(blockSize)
	block.CBCEncrypt(ciphertext, plaintext, iv, aesCipher)

	return ciphertext
}

func CBCAuthenticate(ciphertext []byte) bool {
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		panic("aes did something weird")
	}

	plaintext := make([]byte, len(ciphertext)-blockSize)
	block.CBCDecrypt(plaintext, ciphertext, aesCipher)

	plaintext, err = block.UnPad(plaintext, blockSize)
	if err != nil {
		return false
	}

	m, err := parse.KVDecodeSemicolon(string(plaintext))
	if v, ok := m["admin"]; ok {
		if v == "true" {
			return true
		}
	}

	return false
}
