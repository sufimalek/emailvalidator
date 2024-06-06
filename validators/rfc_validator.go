// package validators

// import "regexp"

// type RFCValidator struct{}

// var rfcEmailRegex = regexp.MustCompile(`(?i)^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// func (v *RFCValidator) Validate(email string) bool {
// 	return rfcEmailRegex.MatchString(email)
// }

package validators

import (
	"log"
	"net/mail"
)

type RFCValidator struct{}

func (v *RFCValidator) Validate(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Printf("RFCValidator: Email %s is not RFC compliant", email)
		return false
	}
	return true
}
