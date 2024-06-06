// package validators

// import "regexp"

// type NoRFCWarningsValidator struct{}

// var noRfcWarningsRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// func (v *NoRFCWarningsValidator) Validate(email string) bool {
// 	return noRfcWarningsRegex.MatchString(email)
// }

package validators

import (
	"log"
	"regexp"
)

type NoRFCWarningsValidator struct{}

func (v *NoRFCWarningsValidator) Validate(email string) bool {
	const emailRegex = `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		log.Printf("NoRFCWarningsValidator: Email %s has warnings based on stricter RFC rules", email)
		return false
	}
	return true
}
