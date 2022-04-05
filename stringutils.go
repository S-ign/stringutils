package stringutils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
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

// GetAllSubstrInBetweenTwoStrs returns a slice of substrings that are
// surrounded by given start and end strings.
// Ex.
// str := "Get {{.placeholders}} within a {{.template}} string."
// substrs := GetAllSubstrInBetweenTwoStrs(str, "{{.", "}}")
// substrs = []string{"placeholders", "template"}
func GetAllSubstrInBetweenTwoStrs(str string, startS string, endS string) map[string]string {
	strS := strings.Split(str, " ")
	substrs := make(map[string]string)
	for _, ss := range strS {
		var result string
		s := strings.Index(ss, startS)
		if s == -1 {
			continue
		}
		newS := ss[s+len(startS):]
		e := strings.Index(newS, endS)
		if e == -1 {
			continue
		}
		result = newS[:e]
		substrs[result] = ""
	}
	return substrs
}
