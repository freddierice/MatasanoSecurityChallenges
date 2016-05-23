package block

import (
	"errors"
)

var ErrPaddingInvalid = errors.New("invalid padding")

// implement PKCS7 padding

func Pad(src []byte, blockSize int) []byte {

	nblocks := len(src)/blockSize + 1
	blank := nblocks*blockSize - len(src)
	dst := make([]byte, nblocks*blockSize)

	copy(dst, src)

	for i := len(src); i < len(dst); i++ {
		dst[i] = byte(blank)
	}

	return dst
}

func UnPad(src []byte, blockSize int) ([]byte, error) {

	blank := int(src[len(src)-1])
	if len(src) < blank || blank > blockSize || blank == 0 {
		return nil, ErrPaddingInvalid
	}

	for i := len(src) - blank; i < len(src); i++ {
		if src[i] != byte(blank) {
			return nil, ErrPaddingInvalid
		}
	}

	return src[:len(src)-blank], nil
}
