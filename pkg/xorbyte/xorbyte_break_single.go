package xorbyte

import "../plaintext"

func BreakSingle(ciphertext []byte) ([]byte, byte) {

	plaintexts := make([][]byte, 0)
	keys := make([]interface{}, 0)
	for i := 0; i < 0x100; i++ {
		xorbyteCipher, err := NewCipher([]byte{byte(i)})
		if err != nil {
			panic("could not create cipher")
		}

		plaintext := make([]byte, len(ciphertext))
		xorbyteCipher.Decrypt(plaintext, ciphertext)

		plaintexts = append(plaintexts, plaintext)
		keys = append(keys, byte(i))
	}

	ps := plaintext.GetGuessesEnglish(plaintexts, keys, 1)

	key, _ := ps[0].Identifier.(byte)
	return ps[0].Plaintext, key
}
