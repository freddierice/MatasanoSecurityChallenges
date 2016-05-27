package block

import (
	"crypto/cipher"
	"errors"
	"sync"

	"../xorbyte"
)

var NWORKERS = 2

var ErrBlockSize = errors.New("dst/src have the wrong lengths")
var ErrIVLen = errors.New("len(iv) must be the block size")

func CBCEncrypt(dst, src, iv []byte, blockCipher cipher.Block) error {

	blockSize := blockCipher.BlockSize()
	if len(dst) != len(src)+blockSize || len(src)%blockSize != 0 {
		return ErrBlockSize
	}
	if len(iv) != blockSize {
		return ErrIVLen
	}

	// first block of CBC is iv
	copy(dst, iv)

	// next block is enc(prev_ciphertext xor current_plaintext)
	for i := 0; i < len(src)/blockSize; i++ {
		start0 := i * blockSize
		start1 := start0 + blockSize
		start2 := start1 + blockSize
		copy(dst[start1:start2], src[start0:start1])
		xorbyte.XOR(dst[start1:start2], dst[start0:start1])
		blockCipher.Encrypt(dst[start1:start2], dst[start1:start2])
	}

	return nil
}

func CBCDecrypt(dst, src []byte, blockCipher cipher.Block) error {
	blockSize := blockCipher.BlockSize()
	if len(dst)+blockSize != len(src) || len(src)%blockSize != 0 {
		return ErrBlockSize
	}

	blockchan := make(chan int, NWORKERS)
	wg := &sync.WaitGroup{}
	worker := func() {
		defer wg.Done()
		for {
			if block, ok := <-blockchan; ok {
				start0 := block * blockSize
				start1 := start0 + blockSize
				start2 := start1 + blockSize
				blockCipher.Decrypt(dst[start0:start1], src[start1:start2])
				xorbyte.XOR(dst[start0:start1], src[start0:start1])
			} else {
				break
			}
		}
	}

	for i := 0; i < NWORKERS; i++ {
		go worker()
		wg.Add(1)
	}
	for i := 0; i < len(dst)/blockSize; i++ {
		blockchan <- i
	}
	close(blockchan)
	wg.Wait()

	return nil
}
