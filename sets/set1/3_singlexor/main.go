package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"../../../pkg/util"
	"../../../pkg/xorbyte"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ciphertext, err := util.ReadLineHex(reader)
	if err != nil {
		log.Fatalf("could not decode string as hex : %v", err)
	}

	plaintext, key := xorbyte.BreakSingle(ciphertext)
	fmt.Printf("key: %c\nplaintext: %s\n", key, string(plaintext))
}
