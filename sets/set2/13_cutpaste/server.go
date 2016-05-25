package main

import (
	"crypto/aes"
	"crypto/cipher"
	"strings"

	"../../../pkg/block"
	"../../../pkg/parse"
	"../../../pkg/util"
)

var aesCipher cipher.Block
var key []byte
var blockSize int

func init() {
	var err error

	blockSize = 16
	key = util.GenerateRandomBytes(blockSize)
	aesCipher, err = aes.NewCipher(key)
	if err != nil {
		panic("aes made a mistake")
	}
}

func profile_for(email string) string {
	m := make(map[string]string)

	email = strings.Replace(email, "&", "", -1)
	email = strings.Replace(email, "=", "", -1)

	m["email"] = email
	m["uid"] = "10"
	m["role"] = "user"

	return parse.KVEncode(m)
}

func EncryptProfile(email string) []byte {
	profileBytes := block.Pad([]byte(profile_for(email)), aesCipher.BlockSize())

	block.ECBEncrypt(profileBytes, profileBytes, aesCipher)

	return profileBytes
}

func ValidateAdmin(b []byte) bool {
	if len(b)%aesCipher.BlockSize() != 0 {
		return false
	}

	block.ECBDecrypt(b, b, aesCipher)

	b, err := block.UnPad(b, aesCipher.BlockSize())
	if err != nil {
		return false
	}

	m, err := parse.KVDecode(string(b))
	if err != nil {
		return false
	}
	if v, ok := m["role"]; ok {
		if v == "admin" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

	return false
}
