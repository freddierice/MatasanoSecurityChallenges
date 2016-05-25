package oracle

import (
	"strings"

	"../parse"
)

func profile_for(email string) string {
	m := make(map[string]string)

	email = strings.Replace(email, "&", "", -1)
	email = strings.Replace(email, "=", "", -1)

	m["email"] = email
	m["uid"] = "10"
	m["role"] = "user"

	return parse.KVEncode(m)
}

func EncryptProfile(email string) {
	profileBytes := block.Pad([]byte(profile_for(email)))
	block.ECBEncrypt(profileBytes, profileBytes)
}

func ValidateAdmin(c string) {

}
