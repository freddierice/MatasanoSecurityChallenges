package main

import (
	"bytes"
	"fmt"

	"../../../pkg/oracle"
)

// Proof that I can detect which one is chosen. the basic idea is
// if we encrypt two blocks of 16 A's, they will be the same only
// in ecb mode. So if I add enough padding, I can detect ECB in two
// of the blocks. If there is no repetition, then it is CBC mode.
func main() {

	correct := 0
	total := 100
	for i := 0; i < total; i++ {
		ciphertext, mode := oracle.EncryptAesEcbCbc([]byte("AAAAAAAAAAA" + "AAAAAAAAAAAAAAAA" + "AAAAAAAAAAAAAAAA" + "AAAAAAAAAAA"))
		if bytes.Equal(ciphertext[16:32], ciphertext[32:48]) && mode == oracle.ModeECB {
			correct++
		} else if mode == oracle.ModeCBC {
			correct++
		}
	}

	fmt.Printf("Got %d out of %d right: %v%%.\n", correct, total, float64(100)*float64(correct)/float64(total))
}
