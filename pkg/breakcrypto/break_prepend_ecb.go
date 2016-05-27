package breakcrypto

import "bytes"

type PrependECBOracle func([]byte) []byte

func BreakPrependECBOracle(o PrependECBOracle, charset []byte) string {

	plaintextLen, blockSize := GetPlaintextLenAndBlockSize(o)
	nblocks := len(o([]byte(""))) / blockSize
	message := make([]byte, blockSize)
	blockPadding := make([]byte, blockSize)

	for i := 0; i < blockSize; i++ {
		message[i] = byte('A')
	}

	ciphertextOffsets := make([][]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		ciphertextOffsets[i] = o([]byte(message[:i]))
	}

	// get the message byte by byte
	for block := 1; block <= nblocks; block++ {
		start0 := (block - 1) * blockSize
		start1 := start0 + blockSize
		copy(blockPadding, []byte(message)[start0+1:start1])
		for i := 0; i < blockSize; i++ {
			var c byte
			offset := blockSize - i - 1
			if offset < 0 {
				offset += blockSize
			}
			b := ciphertextOffsets[offset][start0:start1]
			for j := 0; j < len(charset); j++ {
				blockPadding[blockSize-1] = charset[j]
				if bytes.Equal(o(blockPadding)[0:blockSize], b) {
					c = charset[j]
					break
				}
			}
			if block == nblocks && (plaintextLen%blockSize) == i {
				break
			}
			message = append(message, c)
			blockPadding = append(blockPadding[1:], c)
		}
	}

	// remove the AA..A section of the message
	return string(message[blockSize:])
}

func GetPlaintextLenAndBlockSize(o PrependECBOracle) (int, int) {
	input := ""
	initialLength := len(o([]byte(input)))
	for {
		input = input + "A"
		newLength := len(o([]byte(input)))
		if newLength != initialLength {
			blockSize := newLength - initialLength
			plaintextLen := initialLength - len(input) // +1 if no padding
			return plaintextLen, blockSize
		}
	}
}
