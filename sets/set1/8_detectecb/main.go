package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"../../../pkg/block"
	"../../../pkg/util"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ciphertexts := make([][]byte, 0)
	for {
		ciphertext, err := util.ReadLineHex(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not read in ciphertexts: %v\n", err)
			os.Exit(1)
		}

		ciphertexts = append(ciphertexts, ciphertext)
	}

	minScore := block.ECBCiphertextScore(ciphertexts[0])
	minText := 0
	for i := 1; i < len(ciphertexts); i++ {
		if score := block.ECBCiphertextScore(ciphertexts[i]); score < minScore {
			minScore = score
			minText = i
		}
	}

	ciphertext := hex.EncodeToString(ciphertexts[minText])

	fmt.Printf("line: %d\nciphertext: %s\n", minText, ciphertext)
}
