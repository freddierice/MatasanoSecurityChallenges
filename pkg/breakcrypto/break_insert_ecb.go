package breakcrypto

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"../block"
)

var blockSize = 16

type InsertECBOracle func([]byte) []byte

var paddingKit []byte
var paddingBytes []byte
var encryptedPaddingBytes []byte
var offsetBytes []byte

func init() {
	paddingKit = make([]byte, 2*16*16+16)
	paddingBytes = []byte("0123456789abcdef")
	offsetBytes = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ12345")
	curr := 0
	for i := 0; i < 16; i++ {
		paddingKit[curr] = byte('g')
		curr++
		copy(paddingKit[curr:], paddingBytes)
		curr += blockSize
		copy(paddingKit[curr:], paddingBytes)
		curr += blockSize
	}
}

func BreakInsertECBOracle(o InsertECBOracle, charset []byte) string {
	LoadEncryptedPaddingBytes(o)
	offsetCiphers := GetOffsets(o)
	checkBlock := make([]byte, blockSize)
	nblocks := len(offsetCiphers[0])/blockSize - 1

	copy(checkBlock, offsetBytes[16:])
	message := make([]byte, 0)
	for block := 0; block < nblocks; block++ {
		start0 := block * blockSize
		start1 := start0 + blockSize
		for i := 0; i < blockSize; i++ {
			for j := 0; j < len(charset); j++ {
				checkBlock[blockSize-1] = charset[j]
				if bytes.Equal(GetEncryptedForm(o, checkBlock), offsetCiphers[i][start0:start1]) {
					message = append(message, charset[j])
					break
				}
			}
			checkBlock = append(checkBlock[1:], byte(0))
		}
	}

	unpadded, err := block.UnPad(message, blockSize)
	if err != nil {
		return string(message)
	}

	return string(unpadded)
}

func LoadEncryptedPaddingBytes(o InsertECBOracle) {
	b := o([]byte(paddingKit))
	nblocks := len(b) / blockSize
	for i := 1; i < nblocks; i++ {
		start0 := (i - 1) * blockSize
		start1 := start0 + blockSize
		start2 := start1 + blockSize
		if bytes.Equal(b[start0:start1], b[start1:start2]) {
			encryptedPaddingBytes = make([]byte, blockSize)
			copy(encryptedPaddingBytes, b[start0:start1])
			return
		}
	}
	return
}

func LoadPaddingKit(b []byte) {
	if len(b) != blockSize {
		panic("len(b) must equal blockSize")
	}

	curr := 0
	for i := 0; i < 16; i++ {
		curr++
		curr += blockSize
		copy(paddingKit[curr:], b)
		curr += blockSize
	}
}

func GetEncryptedForm(o InsertECBOracle, b []byte) []byte {
	if len(b) != blockSize {
		panic("len(b) must equal blockSize")
	}
	LoadPaddingKit(b)
	encB := o(paddingKit)
	nblocks := len(encB) / blockSize
	for i := 1; i < nblocks; i++ {
		start0 := (i - 1) * blockSize
		start1 := start0 + blockSize
		if bytes.Equal(encryptedPaddingBytes, encB[start0:start1]) {
			return encB[start1 : start1+blockSize]
		}
	}
	return nil
}

func GetOffsets(o InsertECBOracle) [][]byte {

	offsetCiphers := make([][]byte, blockSize)
	offsets := make(map[string]int)
	done := make(map[int]bool)

	for i := 0; i < blockSize; i++ {
		encryptedForm := GetEncryptedForm(o, offsetBytes[i:i+blockSize])
		offsets[string(encryptedForm)] = i
		//fmt.Println(hex.EncodeToString(encryptedForm))
	}

	for len(done) != blockSize {
		b := o(offsetBytes)
		for i := 0; i < len(b)/blockSize; i++ {
			if v, ok := offsets[string(b[i*blockSize:(i+1)*blockSize])]; ok {
				if _, ok := done[v]; !ok {
					offsetCiphers[v] = b[(i+1)*blockSize:]
					done[v] = true
				}
			}
		}
	}

	return offsetCiphers
}

func PrintBlocks(b []byte) {
	nblocks := len(b) / blockSize
	for i := 0; i < nblocks; i++ {
		fmt.Println(hex.EncodeToString(b[i*blockSize : (i+1)*blockSize]))
	}
}
