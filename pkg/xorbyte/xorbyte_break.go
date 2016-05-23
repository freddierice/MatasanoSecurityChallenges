package xorbyte

import (
	"fmt"

	"../util"
)

func Break(ciphertext []byte) ([]byte, []byte) {

	// get the key size
	keySize := BreakKeySize(ciphertext)

	// get the key for each offset modulo the key size
	key := make([]byte, keySize)
	ciphertextSubset := make([]byte, len(ciphertext)/keySize)
	for offset := 0; offset < keySize; offset++ {
		for i := 0; i < len(ciphertext)/keySize; i++ {
			ciphertextSubset[i] = ciphertext[offset+i*keySize]
		}

		_, keyPart := BreakSingle(ciphertextSubset)
		key[offset] = keyPart
	}

	xorbyteCipher, err := NewCipher(key)
	if err != nil {
		panic("could not create the xorbyteCipiher")
	}

	xorbyteCipher.Encrypt(ciphertext, ciphertext)

	return ciphertext, key
}

func BreakKeySize(ciphertext []byte) int {
	maxKeySize := 128
	if len(ciphertext)/2 < maxKeySize {
		maxKeySize = len(ciphertext) / 2
	}

	minSize := 1
	minScore := KeySizeScore(ciphertext, minSize)
	for keySize := 2; keySize < maxKeySize; keySize++ {
		keyScore := KeySizeScore(ciphertext, keySize)
		if keyScore < minScore {
			minScore = keyScore
			minSize = keySize
		}
		fmt.Printf("(%d, %d)\n", keySize, keyScore)
	}

	return minSize
}

func KeySizeScore(ciphertext []byte, keySize int) int {
	blocks := len(ciphertext) / keySize
	score := float32(0)
	counts := 0
	for bIter := 2; bIter < blocks; bIter += 2 {
		counts++
		score += float32(
			util.Hamming(
				ciphertext[(bIter-2)*keySize:(bIter-1)*keySize],
				ciphertext[(bIter-1)*keySize:bIter*keySize]))
	}
	score = score / float32(counts*keySize)
	return int(score * 100)
}
