package breakcrypto

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"../oracle"
	"../util"
)

func TestGetEncryptedForm(t *testing.T) {
	testBytes := util.GenerateRandomBytes(blockSize)
	blockSize := 16
	o := oracle.EncryptAesEcbInsert

	LoadEncryptedPaddingBytes(o)

	actual := oracle.EncryptAesEcb(testBytes)[0:blockSize]
	guess := GetEncryptedForm(o, testBytes)
	if !bytes.Equal(guess, actual) {
		fmt.Printf("recieved: %v\n", hex.EncodeToString(guess))
		fmt.Printf("expected: %v\n", hex.EncodeToString(actual))
		t.FailNow()
	}
}

func TestGetOffsets(t *testing.T) {
	/*
		o := EncryptAesEcbInsert
		LoadEncryptedPaddingBytes(o)

		for _, b := range GetOffsets(o) {
			fmt.Println(len(hex.EncodeToString(b)))
		}
	*/
}

func TestBreakInsertECBOracle(t *testing.T) {
	var charset = []byte{' ', 'e', 't', 'a', 'o', 'i', 'n', 's', 'h', 'r', 'd', 'l', 'c', 'u', 'm', 'w', 'f', 'g', 'y', 'p', 'b', 'v', 'k', 'j', 'x', 'q', 'z', 'E', 'T', 'A', 'O', 'I', 'N', 'S', 'H', 'R', 'D', 'L', 'C', 'U', 'M', 'W', 'F', 'G', 'Y', 'P', 'B', 'V', 'K', 'J', 'X', 'Q', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', ':', '?', '!', '(', ')', '[', ']'}
	set := make(map[byte]bool)
	for _, c := range charset {
		set[c] = true
	}
	for i := 0; i < 0x100; i++ {
		if _, ok := set[byte(i)]; !ok {
			charset = append(charset, byte(i))
		}
	}

	m := BreakInsertECBOracle(oracle.EncryptAesEcbInsert, charset)
	fmt.Println(m)
}
