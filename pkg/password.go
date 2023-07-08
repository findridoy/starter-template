package pkg

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func Hash(plain string) string {
	hash := sha256.New()
	_, err := hash.Write([]byte(plain))
	if err != nil {
		return ""
	}
	resultingHash := hash.Sum([]byte(os.Getenv("APP_KEY")))
	return fmt.Sprintf("%x", resultingHash)
}


