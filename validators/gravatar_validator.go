package validators

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type GravatarValidator struct{}

func (v *GravatarValidator) Validate(email string) bool {
	email = strings.TrimSpace(strings.ToLower(email))
	hash := md5.Sum([]byte(email))
	hashStr := hex.EncodeToString(hash[:])

	gravatarURL := fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=404", hashStr)
	resp, err := http.Get(gravatarURL)
	if err != nil {
		log.Printf("GravatarValidator: Error fetching Gravatar for %s: %v", email, err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Printf("GravatarValidator: Found Gravatar for %s", email)
		return true
	}
	log.Printf("GravatarValidator: No Gravatar found for %s", email)
	return false
}
