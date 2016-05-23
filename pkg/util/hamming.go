package util

// find the hamming distance between two byte strings
func Hamming(b1, b2 []byte) int {
	count := 0
	for i := range b1 {
		count += hamming_byte(b1[i], b2[i])
	}
	return count
}

// find the hamming distance between two strings
func HammingString(s1, s2 string) int {
	return Hamming([]byte(s1), []byte(s2))
}

// the bytes that differ will be 1 in the xor
// count the number of differing bits
func hamming_byte(b1, b2 byte) (count int) {
	diff := b1 ^ b2
	mask := byte(0x1)
	count = 0
	for i := uint(0); i < 8; i++ {
		count += int((((mask << i) & diff) >> i))
	}
	return
}
