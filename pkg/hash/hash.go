package hash

import (
	"crypto/sha1"
	"fmt"
)

func GenerateHash(password string, salt string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(salt))), nil
}
