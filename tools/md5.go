package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateMD5Hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hashInBytes := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hashInBytes)

	return hashString
}
