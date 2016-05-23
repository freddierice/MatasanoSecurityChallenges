package block

import "crypto/cipher"

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
