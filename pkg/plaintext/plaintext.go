package plaintext

import "sort"

// pulled from wikipedia Letter_frequency
// (letter frequencies times 1000
var KEYS = map[byte]int{
	' ': 130,
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

type PlaintextScorer func([]byte) int

func GetGuessesEnglish(plaintexts [][]byte, identifier []interface{}, n int) []PlaintextScored {
	return GetGuesses(plaintexts, identifier, n, EnglishScorer)
}

// Calculate a decryption score for each plaintext, then return the
// top scoring plaintexts
func GetGuesses(plaintexts [][]byte, identifier []interface{}, n int, plaintextScorer PlaintextScorer) []PlaintextScored {
	ps := make([]PlaintextScored, len(plaintexts))
	for i := range plaintexts {
		ps[i].Plaintext = make([]byte, len(plaintexts[i]))
		copy(ps[i].Plaintext, plaintexts[i])
		ps[i].Identifier = identifier[i]
		ps[i].Score = plaintextScorer(plaintexts[i])
	}

	sort.Sort(sort.Reverse(ByScore(ps)))

	if len(plaintexts) < n {
		return ps
	}

	return ps[:n]
}

// Calculates the score of the plaintext. It is based on a quadratic,
// where having a space will double the score
func EnglishScorer(plaintext []byte) int {
	x := 0
	for _, p := range plaintext {
		if val, ok := KEYS[p]; ok {
			x += val
		}
	}

	return x
}
