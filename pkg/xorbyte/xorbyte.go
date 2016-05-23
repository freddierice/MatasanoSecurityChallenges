package xorbyte

import "crypto/cipher"

type xorbyteCipher struct {
	blockSize int
	key       []byte
	idxEnc    int
	idxDec    int
}

func NewCipher(key []byte) (cipher.Block, error) {
	ks := len(key)
	c := xorbyteCipher{
		blockSize: ks,
		key:       make([]byte, ks),
		idxEnc:    0,
		idxDec:    0,
	}
	copy(c.key, key)
	return &c, nil
}

func (c *xorbyteCipher) BlockSize() int { return c.blockSize }
func (c *xorbyteCipher) Encrypt(dst, src []byte) {
	if len(dst) < len(src) {
		panic("cannot encrypt: dst too small.")
	}
	for i := range dst {
		dst[i] = src[i] ^ c.key[c.idxEnc]
		c.idxEnc = (c.idxEnc + 1) % c.blockSize
	}
}
func (c *xorbyteCipher) Decrypt(dst, src []byte) {
	if len(dst) < len(src) {
		panic("cannot encrypt: dst too small.")
	}
	for i := range dst {
		dst[i] = src[i] ^ c.key[c.idxDec]
		c.idxDec = (c.idxDec + 1) % c.blockSize
	}
}
