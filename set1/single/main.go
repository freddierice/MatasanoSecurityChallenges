package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var KEYS = map[byte]int{
	'e': 50,
	't': 40,
	'a': 30,
	'o': 20,
	'i': 10,
	'n': 1,
}

type ScoredPlaintext struct {
	Plaintext []byte
	Score     int
	Key       byte
}

type ByScore []ScoredPlaintext

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

func main() {
	reader := bufio.NewReader(os.Stdin)

	ciphertext, err := readHex(reader)
	if err != nil {
		log.Fatalf("could not decode string as hex : %v", err)
	}

	// Calculate a decryption score for each plaintext, and remember the key
	// with the best score.
	scoredPlaintexts := make([]ScoredPlaintext, 0xff)
	for i := 0; i < 0xff; i++ {
		plaintext := make([]byte, len(ciphertext))
		decrypt(plaintext, ciphertext, byte(i+1))

		scoredPlaintexts[i].Plaintext = plaintext
		scoredPlaintexts[i].Key = byte(i + 1)
		scoredPlaintexts[i].Score = score(plaintext)
	}

	// sort them by their score
	sort.Sort(sort.Reverse(ByScore(scoredPlaintexts)))

	//list the top few
	top := 5
	fmt.Printf("Here are the top %d\n", top)
	for i := 0; i < top; i++ {
		fmt.Printf("%c: %s\n", scoredPlaintexts[i].Key, string(scoredPlaintexts[i].Plaintext))
	}
}

func decrypt(plaintext, ciphertext []byte, key byte) {
	for i := range ciphertext {
		plaintext[i] = ciphertext[i] ^ key
	}
}

func score(plaintext []byte) int {
	x := 0
	for _, p := range plaintext {
		if val, ok := KEYS[p]; ok {
			x += val
		}
	}
	return x
}

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
