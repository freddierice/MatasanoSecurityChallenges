package parse

import (
	"errors"
	"strings"
)

var ErrInvalidFormat = errors.New("error: invalid format")

func KVDecode(s string) (map[string]string, error) {
	m := make(map[string]string)
	kvstrings := strings.Split(s, "&")
	for _, kvstring := range kvstrings {
		kvpair := strings.Split(kvstring, "=")
		if len(kvpair) != 2 {
			return nil, ErrInvalidFormat
		}
		m[kvpair[0]] = kvpair[1]
	}
	return m, nil
}

func KVEncode(m map[string]string) string {
	kvpairs := make([]string, len(m))
	keys := []string{"email", "uid", "role"}
	for i, k := range keys {
		kvpairs[i] = k + "=" + m[k]
	}
	return strings.Join(kvpairs, "&")
}
