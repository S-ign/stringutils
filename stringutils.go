package stringutils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

// GenerateHash creates a hash string from given slice of bytes and current time
func GenerateHash(b []byte) (string, error) {

	// generate hash
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", string(b), time.Now())))
	hash := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return hash, nil
}
