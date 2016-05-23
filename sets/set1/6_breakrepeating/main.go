package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"../../../pkg/xorbyte"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read all from the buffer : %v", err)
	}
	base64str := strings.Replace(string(bytes), "\n", "", -1)
	bytes, err = base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not base64 decode this string : %v", err)
	}

	plaintext, key := xorbyte.Break(bytes)
	fmt.Printf("Message: \n%s\n\nkey: \n%s\n", string(plaintext), string(key))
}
