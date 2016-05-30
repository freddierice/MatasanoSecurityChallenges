package parse

import (
	"errors"
	"strings"
)

var ErrInvalidFormat = errors.New("error: invalid format")

func KVEncodeSemicolon(m map[string]string) string {
	return kvEncode(m, ";", "=")
}

func KVDecodeSemicolon(s string) (map[string]string, error) {
	return kvDecode(s, ";", "=")
}

func KVEncode(m map[string]string) string {
	return kvEncode(m, "&", "=")
}

func KVDecode(s string) (map[string]string, error) {
	return kvDecode(s, "&", "=")
}

func kvEncode(m map[string]string, amp, eq string) string {
	kvpairs := make([]string, len(m))
	keys := []string{"email", "uid", "role"}
	for i, k := range keys {
		kvpairs[i] = k + eq + m[k]
	}
	return strings.Join(kvpairs, amp)
}

func kvDecode(s, amp, eq string) (map[string]string, error) {
	m := make(map[string]string)
	kvstrings := strings.Split(s, amp)
	for _, kvstring := range kvstrings {
		kvpair := strings.Split(kvstring, eq)
		if len(kvpair) != 2 {
			return nil, ErrInvalidFormat
		}
		m[kvpair[0]] = kvpair[1]
	}

	return m, nil
}
