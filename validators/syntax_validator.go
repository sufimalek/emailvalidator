package validators

import (
	"log"
	"regexp"
)

type SyntaxValidator struct{}

func (v *SyntaxValidator) Validate(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		log.Printf("SyntaxValidator: Invalid email format for %s", email)
		return false
	}
	return true
}
