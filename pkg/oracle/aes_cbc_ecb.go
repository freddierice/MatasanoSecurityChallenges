package oracle

import (
	"crypto/aes"
	"math/rand"

	"../block"
	"../util"
)

type EncryptionMode int

const (
	ModeCBC = iota
	ModeECB
)

func EncryptAesEcbCbc(plaintext []byte) ([]byte, EncryptionMode) {

	blockSize := 16
	nblocks := (len(plaintext)+blockSize-1)/blockSize + 1

	key := util.GenerateRandomBytes(blockSize)
	prepend := util.GenerateRandomBytesRange(5, 10)
	postpend := util.GenerateRandomBytesRange(5, 10)

	plaintext = append(prepend, plaintext...)
	plaintext = append(plaintext, postpend...)

	mode := EncryptionMode(rand.Intn(1))
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		panic("something is terribly wrong with aes")
	}

	b := make([]byte, nblocks*blockSize)

	switch mode {
	case ModeCBC:
		iv := util.GenerateRandomBytes(blockSize)
		block.CBCEncrypt(b, plaintext, iv, aesCipher)
	case ModeECB:
		// ecb has no iv, cut out first block
		block.ECBEncrypt(b[blockSize:], plaintext, aesCipher)
	default:
		panic("should not get here")
	}

	// do not send iv, cut out first block
	return b[blockSize:], mode
}
