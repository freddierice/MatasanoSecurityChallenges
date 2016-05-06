package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	str = strings.Trim(str, "\n")
	if err != nil {
		log.Fatalf("no hex string entered : %v", err)
	}

	bytes, err := hex.DecodeString(str)
	if err != nil {
		log.Fatalf("could not decode string as hex : %v", err)
	}

	str = base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(str)
}
