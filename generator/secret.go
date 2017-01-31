package generator

import "crypto/sha256"

func GenerateSecret() []byte {
	input := "1234"
	sha256 := sha256.New()
	sha256.Write([]byte(input))

	return sha256.Sum(nil)
}
