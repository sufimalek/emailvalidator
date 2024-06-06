package validators

import (
	"log"
	"strings"
)

type BanWordsUsernameValidator struct {
	bannedWords []string
}

func NewBanWordsUsernameValidator(bannedWords []string) *BanWordsUsernameValidator {
	return &BanWordsUsernameValidator{bannedWords: bannedWords}
}

func (v *BanWordsUsernameValidator) Validate(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		log.Printf("BanWordsUsernameValidator: Invalid email format for %s", email)
		return false
	}
	localPart := parts[0]
	for _, word := range v.bannedWords {
		if strings.Contains(localPart, word) {
			log.Printf("BanWordsUsernameValidator: Username %s contains banned word %s", localPart, word)
			return false
		}
	}
	return true
}

// func (v *BanWordsUsernameValidator) Validate(email string) bool {
// 	username := strings.Split(email, "@")[0]
// 	for _, word := range v.bannedWords {
// 		if strings.Contains(strings.ToLower(username), word) {
// 			return false
// 		}
// 	}
// 	return true
// }
