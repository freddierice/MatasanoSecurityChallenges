package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"../../../pkg/util"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	bytes1, err := util.ReadLineHex(reader)
	if err != nil {
		log.Fatalf("could not decode string as hex : %v", err)
	}

	bytes2, err := util.ReadLineHex(reader)
	if err != nil {
		log.Fatalf("could not decode string as hex : %v", err)
	}

	if len(bytes1) != len(bytes2) {
		fmt.Println("len of byte strings differ")
		os.Exit(1)
	}

	util.XORBytes(bytes1, bytes2)

	str := hex.EncodeToString(bytes1)
	fmt.Println(str)
}
