package generator

import "crypto/sha256"

// GenerateSecret generates a new client secret for an ID
func GenerateSecret(id string) []byte {
	sha256 := sha256.New()
	sha256.Write([]byte(id))

	return sha256.Sum(nil)
}
