package util

import (
	"bufio"
	"encoding/hex"
	"strings"
)

func ReadLineHex(reader *bufio.Reader) ([]byte, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	str = strings.Trim(str, "\n")

	return hex.DecodeString(str)
}

func XORBytes(dst, src []byte) {
	for i := range dst {
		dst[i] ^= src[i]
	}
}
