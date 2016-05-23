package xorbyte

import (
	"encoding/hex"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\n" +
		"I go crazy when I hear a cymbal")
	ciphertext := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a262" +
		"26324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165" +
		"286326302e27282f"
	key := []byte("ICE")
	result := make([]byte, len(plaintext))

	// test the encryption
	xorbyteCipher, err := NewCipher(key)
	if err != nil {
		t.Error("could not create NewCipher: ", err)
		t.FailNow()
	}

	xorbyteCipher.Encrypt(result, plaintext)

	if hexResult := hex.EncodeToString(result); ciphertext != hexResult {
		t.Error("expected hexstring\n", ciphertext, "\ngot\n", hexResult)
	}
}

func TestBreakSingle(t *testing.T) {
	ciphertext := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	plaintext := "Cooking MC's like a pound of bacon"
	key := byte('X')

	ciphertextBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		t.Error("invalid ciphertext")
		t.FailNow()
	}

	brokenPlaintext, brokenKey := BreakSingle(ciphertextBytes)
	if brokenKey != key {
		t.Error("keys do not match")
	}
	if string(brokenPlaintext) != plaintext {
		t.Error("plaintexts do not match")
	}
}
