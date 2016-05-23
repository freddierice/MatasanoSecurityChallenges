package util

import "testing"

func TestHamming(t *testing.T) {
	distance := HammingString("this is a test", "wokka wokka!!!")
	if distance != 37 {
		t.Error("expected hamming distance 37, got ", distance)
	}
}
