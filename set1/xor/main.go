package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func readHex(reader *bufio.Reader) (bytes []byte, err error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		err = errors.New("no hex string entered")
		return
	}
	str = strings.Trim(str, "\n")

	bytes, err = hex.DecodeString(str)
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	bytes1, err := readHex(reader)
	if err != nil {
		log.Fatalf("could not decode string as hex : %v", err)
	}

	bytes2, err := readHex(reader)
	if err != nil {
		log.Fatalf("could not decode string as hex : %v", err)
	}

	if len(bytes1) != len(bytes2) {
		fmt.Println("len of byte strings differ")
		os.Exit(1)
	}

	for i := range bytes1 {
		bytes1[i] ^= bytes2[i]
	}

	str := hex.EncodeToString(bytes1)
	fmt.Println(str)
}
