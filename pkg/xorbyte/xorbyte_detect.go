package xorbyte

import (
	"../plaintext"
)

func DetectSingle(ciphertexts [][]byte) ([]byte, byte) {
	keys := make([]interface{}, 0)
	plaintexts := make([][]byte, 0)
	for _, ciphertext := range ciphertexts {
		for i := 0; i < 256; i++ {
			plaintext := make([]byte, len(ciphertext))

			xorbyteCipher, err := NewCipher([]byte{byte(i)})
			if err != nil {
				panic("could not create the xorbyte Cipher")
			}
			xorbyteCipher.Decrypt(plaintext, ciphertext)

			keys = append(keys, byte(i))
			plaintexts = append(plaintexts, plaintext)
		}
	}

	ps := plaintext.GetGuessesEnglish(plaintexts, keys, 1)[0]
	key, _ := ps.Identifier.(byte)
	return ps.Plaintext, key

}
