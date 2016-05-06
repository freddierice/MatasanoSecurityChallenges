package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// pulled from wikipedia Letter_frequency
var KEYS = map[byte]int{
	'e': 127,
	't': 90,
	'a': 81,
	'o': 75,
	'i': 69,
	'n': 67,
	's': 63,
	'h': 60,
	'r': 59,
	'd': 42,
	'l': 40,
	'u': 27,
}

const TRIES = 5

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

	ciphertexts := [][]byte{}
	for {
		// break if the last one was read
		ciphertext, err := readHex(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not decode string as hex : %v", err)
		}
		ciphertexts = append(ciphertexts, ciphertext)
	}

	scoredPlaintexts := make([]ScoredPlaintext, TRIES*len(ciphertexts))
	for i, c := range ciphertexts {
		getGuesses(c, scoredPlaintexts[i*TRIES:], TRIES)
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

// Calculate a decryption score for each plaintext, then return the
// top scoring plaintexts
func getGuesses(ciphertext []byte, topPlaintexts []ScoredPlaintext, n int) {
	scoredPlaintexts := make([]ScoredPlaintext, 0xff)
	for i := 0; i < 0xff; i++ {
		plaintext := make([]byte, len(ciphertext))
		decrypt(plaintext, ciphertext, byte(i+1))

		scoredPlaintexts[i].Plaintext = plaintext
		scoredPlaintexts[i].Key = byte(i + 1)
		scoredPlaintexts[i].Score = score(plaintext)
	}
	sort.Sort(sort.Reverse(ByScore(scoredPlaintexts)))
	copy(topPlaintexts, scoredPlaintexts[:n])
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
		return
	}
	str = strings.Trim(str, "\n")

	bytes, err = hex.DecodeString(str)
	return
}
