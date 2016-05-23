package block

import (
	"bytes"
	"compress/gzip"
	"crypto/cipher"
	"fmt"
	"os"
)

func ECBEncrypt(dst, src []byte, blockCipher cipher.Block) {

	blockSize := blockCipher.BlockSize()
	if len(dst) != len(src) || len(src)%blockSize != 0 {
		panic("len(dst) != len(src) || len(src) is not a multiple of block size")
	}

	for i := 0; i < len(src)/blockSize; i++ {
		blockCipher.Encrypt(dst[i*blockSize:(i+1)*blockSize],
			src[i*blockSize:(i+1)*blockSize])
	}
}

func ECBDecrypt(dst, src []byte, blockCipher cipher.Block) {
	blockSize := blockCipher.BlockSize()
	if len(dst) != len(src) || len(src)%blockSize != 0 {
		panic("len(dst) != len(src) || len(src) is not a multiple of block size")
	}

	for i := 0; i < len(src)/blockSize; i++ {
		blockCipher.Decrypt(dst[i*blockSize:(i+1)*blockSize],
			src[i*blockSize:(i+1)*blockSize])
	}
}

// just check the compressed length
func ECBCiphertextScore(ciphertext []byte) int {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(ciphertext); err != nil {
		fmt.Fprintf(os.Stderr, "could not gzip data")
		os.Exit(1)
	}
	gz.Flush()
	gz.Close()
	return b.Len()
}
