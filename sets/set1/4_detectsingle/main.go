package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"../../../pkg/util"
	"../../../pkg/xorbyte"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ciphertexts := [][]byte{}
	for {
		// break if the last one was read
		ciphertext, err := util.ReadLineHex(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not decode string as hex : %v", err)
		}
		ciphertexts = append(ciphertexts, ciphertext)
	}

	ptext, key := xorbyte.DetectSingle(ciphertexts)
	fmt.Printf("key: %c\nplaintext: %s\n", key, ptext)
}
