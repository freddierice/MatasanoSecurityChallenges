package main

import (
	"bufio"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"../../../pkg/block"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read all from the buffer : %v", err)
	}
	base64str := strings.Replace(string(input), "\n", "", -1)
	input, err = base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not base64 decode this string : %v", err)
	}

	output := make([]byte, len(input))

	aesCipher, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create AES cipher")
		os.Exit(1)
	}
	block.ECBDecrypt(output, input, aesCipher)
	fmt.Printf("%s\n", string(output))
}
