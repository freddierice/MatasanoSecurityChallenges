package xorbyte

func XOR(dst, src []byte) {
	for i := range src {
		dst[i] ^= src[i]
	}
}
